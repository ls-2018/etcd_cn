// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package raft

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/ls-2018/etcd_cn/raft/confchange"
	"github.com/ls-2018/etcd_cn/raft/quorum"
	pb "github.com/ls-2018/etcd_cn/raft/raftpb"
	"github.com/ls-2018/etcd_cn/raft/tracker"
)

// None 是一个占位的节点ID,在没有leader时使用.
const None uint64 = 0
const noLimit = math.MaxUint64

// 状态类型
const (
	StateFollower StateType = iota
	StateCandidate
	StateLeader
	StatePreCandidate
	numStates
)

type ReadOnlyOption int

const (
	ReadOnlySafe ReadOnlyOption = iota
	ReadOnlyLeaseBased
	//1、 ReadOnlySafe
	//	该线性读模式,每次 Follower 进行读请求时,需要和 Leader 同步日志提交位点信息,而 Leader ,需要向过半的 Follower 发起证明自己是 Leader 的轻量的 RPC 请求,
	//	相当于一个 Follower 读,至少需要 1 +(n/2)+ 1 次的 RPC 请求.
	//2、ReadOnlyLeaseBased
	//该线性读模式,每次 Follower 进行读请求时, Leader 只需要判断自己的 Leader 租约是否过期了,如果没有过期,直接可以回复 Follower 自己是 Leader ,
	// 但是该机制对于机器时钟要求很严格,如果有做时钟同步的话,可以考虑使用该线性读模式.
	//如果说对于配置的发布、修改操作比较频繁,可以将 Raft 快照的时间适当的进行调整,避免新节点加入或者节点重启时,由于 Raft 日志回放操作数太多导致节点可开始对外服务的时间过长.

)

const (
	campaignPreElection CampaignType = "CampaignPreElection" // 竞选类型： pre-vote模式
	campaignElection    CampaignType = "CampaignElection"    // 竞选类型：vote模式
	campaignTransfer    CampaignType = "CampaignTransfer"    // 竞选类型：leader开始转移
)

// ErrProposalDropped is returned when the proposal is ignored by some cases,
// so that the proposer can be notified and fail fast.
var ErrProposalDropped = errors.New("撤销raft提案")

// lockedRand is a small wrapper around rand.Rand to provide
// synchronization among multiple raft groups. Only the methods needed
// by the code are exposed (e.g. Intn).
type lockedRand struct {
	mu   sync.Mutex
	rand *rand.Rand
}

func (r *lockedRand) Intn(n int) int {
	r.mu.Lock()
	v := r.rand.Intn(n)
	r.mu.Unlock()
	return v
}

var globalRand = &lockedRand{
	rand: rand.New(rand.NewSource(time.Now().UnixNano())),
}

// CampaignType 竞选类型
type CampaignType string

// StateType 节点在集群中的状态
type StateType uint64

var stmap = [...]string{
	"StateFollower",
	"StateCandidate",
	"StateLeader",
	"StatePreCandidate",
}

func (st StateType) String() string {
	return stmap[uint64(st)]
}

// Config 启动raft的配置参数
type Config struct {
	// ID 是本节点raft的身份.ID不能为0.
	ID uint64

	ElectionTick int // 返回选举权检查对应多少次tick触发次数

	HeartbeatTick int // 返回心跳检查对应多少次tick触发次数

	// Storage 存储 日志项、状态
	Storage Storage //
	// Applied 提交到用户状态机的索引
	Applied uint64 // 起始为0

	// 每条消息的最大大小:math.MaxUint64表示无限制,0表示每条消息最多一个条目.
	MaxSizePerMsg uint64 // 1m
	// MaxCommittedSizePerReady 限制  commited --> apply 之间的数量
	MaxCommittedSizePerReady uint64 // MaxSizePerMsg 它们之前是同一个参数
	// MaxUncommittedEntriesSize 未提交的日志项上限
	MaxUncommittedEntriesSize uint64

	// 最大的处理中的消息数量
	MaxInflightMsgs int

	// CheckQuorum 检查需要维持的选票数,一旦小于,就会丢失leader
	CheckQuorum bool

	// PreVote 防止分区服务器[term会很大]重新加入集群时出现中断
	PreVote bool // PreVote 是否启用PreVote

	// CheckQuorum必须是enabled if ReadOnlyOption is ReadOnlyLeaseBased.
	ReadOnlyOption ReadOnlyOption

	// Logger is the logger used for raft log. For multinode which can host
	// multiple raft group, each raft group can have its own logger
	Logger Logger

	DisableProposalForwarding bool // 禁止将请求转发到leader,默认FALSE
}

// OK
func (c *Config) validate() error {
	if c.ID == None {
		return errors.New("补鞥呢使用None作为ID")
	}
	// 返回心跳检查对应多少次tick触发次数
	if c.HeartbeatTick <= 0 { //
		return errors.New("心跳间隔必须是>0")
	}

	if c.ElectionTick <= c.HeartbeatTick {
		return errors.New("选举超时必须是大于心跳间隔")
	}

	if c.Storage == nil {
		return errors.New("不能没有存储")
	}

	if c.MaxUncommittedEntriesSize == 0 {
		c.MaxUncommittedEntriesSize = noLimit
	}

	//  它们之前是同一个参数.
	if c.MaxCommittedSizePerReady == 0 {
		c.MaxCommittedSizePerReady = c.MaxSizePerMsg
	}

	if c.MaxInflightMsgs <= 0 {
		return errors.New("max inflight messages必须是>0")
	}

	if c.Logger == nil {
		c.Logger = getLogger()
	}
	// 作为leader时的检查
	if c.ReadOnlyOption == ReadOnlyLeaseBased && !c.CheckQuorum {
		return errors.New("如果ReadOnlyOption 是 ReadOnlyLeaseBased 的时候必须开启CheckQuorum.")
	}

	return nil
}

// 封装了当前节点所有的核心数据.
type raft struct {
	id uint64 // 是本节点raft的身份

	Term uint64 // 任期.如果Message的Term字段为0,则表示该消息是本地消息,例如,MsgHup、 MsgProp、 MsgReadlndex 等消息,都属于本地消息.
	Vote uint64 // 当前任期中当前节点将选票投给了哪个节点

	readStates []ReadState

	raftLog *raftLog // 当前节点的log状态信息

	maxMsgSize         uint64 // 每条消息的最大大小
	maxUncommittedSize uint64 // 每条日志最大消息体

	prs tracker.ProgressTracker // 跟踪Follower节点的状态,比如日志复制的matchIndex

	state StateType // 当前节点的状态 ,可选值分为StateFollower、StateCandidate、 StateLeader和StatePreCandidat巳四种状态.

	// isLearner 本节点是不是learner角色
	isLearner bool

	msgs []pb.Message // 缓存了当前节点等待发送的消息.

	lead uint64 // 当前leaderID
	// leader转移到的节点ID
	leadTransferee uint64
	// Only one conf change may be pending (in the log, but not yet
	// applied) at a time. This is enforced via pendingConfIndex, which
	// is set to a value >= the log index of the latest pending
	// configuration change (if any). Config changes are only allowed to
	// be proposed if the leader's applied index is greater than this
	// value.
	pendingConfIndex uint64

	uncommittedSize uint64 // 还未提交的日志条数,非准确值

	readOnly *readOnly

	checkQuorum      bool // 检查需要维持的选票数,一旦小于,就会丢失leader
	preVote          bool // PreVote 是否启用PreVote
	electionElapsed  int  // 选举计时器的指针,其单位是逻辑时钟的刻度,逻辑时钟每推进一次,该字段值就会增加1.
	heartbeatElapsed int  // 心跳计时器的指针,其单位也是逻辑时钟的刻度,逻辑时钟每推进一次,该字段值就会增加1 .

	heartbeatTimeout int // 心跳间隔    ,上限     heartbeatTimeout是当前距离上次心跳的时间
	electionTimeout  int // 选举超时时间,当electionE!apsed 宇段值到达该值时,就会触发新一轮的选举.

	randomizedElectionTimeout int  // 随机选举超时
	disableProposalForwarding bool // 禁止将请求转发到leader,默认FALSE
	// 由 r.ticker = time.NewTicker(r.heartbeat) ;触发该函数的执行  r.start

	tick func() // 逻辑计数器推进函数, 当 Leader状态时 为 tickHeartbeat.其他状态为 tickElection.

	step stepFunc // 阶段函数、在那个角色就执行那个角色的函数、处理接收到的消息

	logger Logger

	// pendingReadIndexMessages is used to store messages of type MsgReadIndex
	// that can't be answered as new leader didn't committed any log in
	// current term. Those will be handled as fast as first log is committed in
	// current term.
	pendingReadIndexMessages []pb.Message
}

// ok
func newRaft(c *Config) *raft {
	if err := c.validate(); err != nil {
		panic(err.Error())
	}
	raftlog := newLogWithSize(c.Storage, c.Logger, c.MaxCommittedSizePerReady) // ✅
	// 搜 s = raft.NewMemoryStorage()
	hs, cs, err := c.Storage.InitialState()
	if err != nil {
		panic(err)
	}

	r := &raft{
		id:                        c.ID,                        // 是本节点raft的身份
		lead:                      None,                        // 当前leaderID
		isLearner:                 false,                       // 本节点是不是learner角色
		raftLog:                   raftlog,                     // 当前节点的log状态信息
		maxMsgSize:                c.MaxSizePerMsg,             // 每条消息的最大大小
		maxUncommittedSize:        c.MaxUncommittedEntriesSize, // 每条日志最大消息体
		prs:                       tracker.MakeProgressTracker(c.MaxInflightMsgs),
		electionTimeout:           c.ElectionTick,  // 返回选举权检查对应多少次tick触发次数
		heartbeatTimeout:          c.HeartbeatTick, // 返回心跳检查对应多少次tick触发次数
		logger:                    c.Logger,
		checkQuorum:               c.CheckQuorum,                 // 检查需要维持的选票数,一旦小于,就会丢失leader
		preVote:                   c.PreVote,                     // PreVote 是否启用PreVote
		readOnly:                  newReadOnly(c.ReadOnlyOption), // etcd_backend/etcdserver/raft.go:469    默认值0 ReadOnlySafe
		disableProposalForwarding: c.DisableProposalForwarding,   // 禁止将请求转发到leader,默认FALSE
	}
	// todo 没看懂
	// -----------------------
	cfg, prs, err := confchange.Restore(confchange.Changer{
		Tracker:   r.prs,
		LastIndex: raftlog.lastIndex(),
	}, cs)
	if err != nil {
		panic(err)
	}
	assertConfStatesEquivalent(r.logger, cs, r.switchToConfig(cfg, prs)) // 判断相不相等
	// -----------------------
	//根据从Storage中获取的HardState,初始化raftLog.committed字段,以及raft.Term和Vote字段
	if !IsEmptyHardState(hs) { // 判断初始状态是不是空的
		r.loadState(hs) // 更新状态索引信息
	}
	//如采Config中己置了Applied,则将raftLog.applied字段重直为指定的Applied值上层模块自己的控制正确的己应用位置时使用该配置
	if c.Applied > 0 {
		raftlog.appliedTo(c.Applied) // ✅
	}
	r.becomeFollower(r.Term, None) // ✅ start

	var nodesStrs []string
	for _, n := range r.prs.VoterNodes() { // 一开始没有
		nodesStrs = append(nodesStrs, fmt.Sprintf("%x", n))
	}

	r.logger.Infof("【newRaft %x [peers: [%s], term: %d, commit: %d, applied: %d, lastindex: %d, lastterm: %d]】",
		r.id, strings.Join(nodesStrs, ","), r.Term, r.raftLog.committed, r.raftLog.applied, r.raftLog.lastIndex(), r.raftLog.lastTerm())
	return r
}

func (r *raft) hasLeader() bool { return r.lead != None }

// OK
func (r *raft) softState() *SoftState { return &SoftState{Lead: r.lead, RaftState: r.state} }

// OK
func (r *raft) hardState() pb.HardState {
	return pb.HardState{
		Term:   r.Term,
		Vote:   r.Vote, // 当前节点的选票投给了谁做我Leader
		Commit: r.raftLog.committed,
	}
}

// send 将状态持久化到一个稳定的存储中,之后再发送消息
func (r *raft) send(m pb.Message) {
	if m.From == None {
		m.From = r.id
	}
	//数据校验,选举类消息必须带term属性
	// 竞选投票相关的消息类型,必须设置term
	if m.Type == pb.MsgVote || m.Type == pb.MsgVoteResp || m.Type == pb.MsgPreVote || m.Type == pb.MsgPreVoteResp {
		if m.Term == 0 {
			panic(fmt.Sprintf("任期应该被设置%s", m.Type))
		}
	} else {
		//其它类消息不能带term属性
		if m.Term != 0 {
			panic(fmt.Sprintf("任期不能被设置,当 %s (was %d)", m.Type, m.Term))
		}
		//除了MsgProp和MsgReadIndex消息外,设置term为raft当前周期
		if m.Type != pb.MsgProp && m.Type != pb.MsgReadIndex {
			m.Term = r.Term
		}
	}

	r.msgs = append(r.msgs, m) // 将消息放入队列 写
}

// 向指定的节点发送信息
func (r *raft) sendHeartbeat(to uint64, ctx []byte) {
	commit := min(r.prs.Progress[to].Match, r.raftLog.committed)
	m := pb.Message{
		To:      to,
		Type:    pb.MsgHeartbeat,
		Commit:  commit, // leader会为每个Follower都维护一个leaderCommit,表示leader认为Follower已经提交的日志条目索引值
		Context: ctx,
	}
	r.send(m)
}

// bcastAppend 向集群中其他节点广播MsgApp消息
func (r *raft) bcastAppend() {
	//遍历所有节点,给除自己外的节点发送日志Append消息
	r.prs.Visit(func(id uint64, _ *tracker.Progress) {
		if id == r.id {
			return
		}
		r.sendAppend(id)
	})
}

// OK 向集群中特定节点发送MsgApp消息
func (r *raft) sendAppend(to uint64) {
	r.maybeSendAppend(to, true)
}

// maybeSendAppend 向给定的peer发送一个带有新条目的追加RPC.如果有消息被发送,返回true.
// sendIfEmpty参数控制是否发送没有条目的消息("空 "消息对于传达更新的Commit索引很有用,但当我们批量发送多条消息时就不可取).
func (r *raft) maybeSendAppend(to uint64, sendIfEmpty bool) bool {
	// 在消息发送之前会检测当前节点的状态,然后查找待发迭的Entry记录并封装成MsgApp消息,
	// 之后根据对应节点的Progress.State值决定发送消息之后的操作

	//1. 获取对端节点当前同步进度
	pr := r.prs.Progress[to]
	if pr.IsPaused() {
		return false
	}
	m := pb.Message{}
	m.To = to
	//2. 注意这里带的term是本次发送给follower的第一条日志条目的term
	term, errt := r.raftLog.term(pr.Next - 1)              // leader认为 follower所在的任期
	ents, erre := r.raftLog.entries(pr.Next, r.maxMsgSize) // 要发给follower的日志
	if len(ents) == 0 && !sendIfEmpty {
		// 这种情况就不发了
		return false
	}

	if errt != nil || erre != nil {
		//3. 如果获取term或日志失败,说明follower落后太多,raftLog内存中日志已经做过快照后被删除了
		// 根据日志进度去取日志条目的时候发现,follower日志落后太多,这通常出现在新节点刚加入或者网络连接出现故障的情况下.
		// 那么在这种情况下,leader改为发送最近一次快照给Follower,从而提高同步效率

		if !pr.RecentActive {
			r.logger.Debugf("ignore sending snapshot to %x since it is not recently active", to)
			return false
		}
		//4. 改为发送Snapshot消息
		m.Type = pb.MsgSnap
		snapshot, err := r.raftLog.snapshot()
		if err != nil {
			if err == ErrSnapshotTemporarilyUnavailable {
				r.logger.Debugf("%x failed to send snapshot to %x because snapshot is temporarily unavailable", r.id, to)
				return false
			}
			panic(err) // TODO(bdarnell)
		}
		if IsEmptySnap(snapshot) {
			panic("need non-empty snapshot")
		}
		m.Snapshot = snapshot
		sindex, sterm := snapshot.Metadata.Index, snapshot.Metadata.Term
		r.logger.Debugf("%x [firstindex: %d, commit: %d] sent snapshot[index: %d, term: %d] to %x [%s]",
			r.id, r.raftLog.firstIndex(), r.raftLog.committed, sindex, sterm, to, pr)
		pr.BecomeSnapshot(sindex)
		r.logger.Debugf("%x paused sending replication messages to %x [%s]", r.id, to, pr)
	} else {
		//5. 发送Append消息
		m.Type = pb.MsgApp             //设置消息类型
		m.Index = pr.Next - 1          //设置MsgApp消息的Index字段
		m.LogTerm = term               //设置MsgApp消息的LogTerm字段
		m.Entries = ents               //设置消息携带的Entry记录集合
		m.Commit = r.raftLog.committed //设置消息的Commit字段,即当前节点的raftLog中最后一条已提交的记录索引值
		//6. 每次发送日志或心跳都会带上最新的commitIndex
		m.Commit = r.raftLog.committed
		if n := len(m.Entries); n != 0 {
			switch pr.State {
			// 在StateReplicate中,乐观地增加
			case tracker.StateReplicate:
				last := m.Entries[n-1].Index
				pr.OptimisticUpdate(last) // 新目标节点对应的Next值（这里不会更新Match）
				pr.Inflights.Add(last)    // 记录已发送但是未收到响应的消息
			case tracker.StateProbe:
				// 消息发送后,就将Progress.Paused字段设置成true,暂停后续消息的发送
				pr.ProbeSent = true
			default:
				r.logger.Panicf("%x 在未知的状态下发送%s", r.id, pr.State)
			}
		}
	}
	//7. 发送消息
	r.send(m)
	return true
}

// bcastHeartbeat sends RPC, without entries to all the peers.
func (r *raft) bcastHeartbeat() {
	lastCtx := r.readOnly.lastPendingRequestCtx()
	if len(lastCtx) == 0 {
		r.bcastHeartbeatWithCtx(nil)
	} else {
		r.bcastHeartbeatWithCtx([]byte(lastCtx))
	}
}

func (r *raft) bcastHeartbeatWithCtx(ctx []byte) {
	r.prs.Visit(func(id uint64, _ *tracker.Progress) {
		if id == r.id {
			return
		}
		r.sendHeartbeat(id, ctx)
	})
}

func (r *raft) advance(rd Ready) {
	r.reduceUncommittedSize(rd.CommittedEntries)

	// If entries were applied (or a snapshot), update our cursor for
	// the next Ready. Note that if the current HardState contains a
	// new Commit index, this does not mean that we're also applying
	// all of the new entries due to commit pagination by size.
	if newApplied := rd.appliedCursor(); newApplied > 0 {
		oldApplied := r.raftLog.applied
		r.raftLog.appliedTo(newApplied)

		if r.prs.Config.AutoLeave && oldApplied <= r.pendingConfIndex && newApplied >= r.pendingConfIndex && r.state == StateLeader {
			// If the current (and most recent, at least for this leader's term)
			// configuration should be auto-left, initiate that now. We use a
			// nil Data which unmarshals into an empty ConfChangeV2 and has the
			// benefit that appendEntry can never refuse it based on its size
			// (which registers as zero).
			ent := pb.Entry{
				Type: pb.EntryConfChangeV2,
				Data: nil,
			}
			// There's no way in which this proposal should be able to be rejected.
			if !r.appendEntry(ent) {
				panic("refused un-refusable auto-leaving ConfChangeV2")
			}
			r.pendingConfIndex = r.raftLog.lastIndex()
			r.logger.Infof("initiating automatic transition out of joint configuration %s", r.prs.Config)
		}
	}

	if len(rd.Entries) > 0 {
		e := rd.Entries[len(rd.Entries)-1]
		r.raftLog.stableTo(e.Index, e.Term)
	}
	if !IsEmptySnap(rd.Snapshot) {
		r.raftLog.stableSnapTo(rd.Snapshot.Metadata.Index)
	}
}

// maybeCommit
func (r *raft) maybeCommit() bool {
	//获取最大的超过半数确认的index
	mci := r.prs.Committed()
	//更新commitIndex
	return r.raftLog.maybeCommit(mci, r.Term)
}

func (r *raft) reset(term uint64) {
	if r.Term != term {
		r.Term = term
		r.Vote = None // 当前任期中当前节点将选票投给了哪个节点
	}
	r.lead = None

	r.electionElapsed = 0
	r.heartbeatElapsed = 0
	r.resetRandomizedElectionTimeout() // 设置随机选举超时
	r.abortLeaderTransfer()            // 置空 leader转移目标

	r.prs.ResetVotes() // 准备通过recordVote进行新一轮的计票工作
	// 重直prs, 其中每个Progress中的Next设置为raftLog.lastindex
	r.prs.Visit(func(id uint64, pr *tracker.Progress) {
		*pr = tracker.Progress{
			Match:     0,
			Next:      r.raftLog.lastIndex() + 1,
			Inflights: tracker.NewInflights(r.prs.MaxInflight),
			IsLearner: pr.IsLearner,
		}
		if id == r.id {
			pr.Match = r.raftLog.lastIndex() // 对应Follower节点当前己经成功复制的Entry记录的索引值,不知有没有同步大多数节点
		}
	})

	r.pendingConfIndex = 0
	r.uncommittedSize = 0
	r.readOnly = newReadOnly(r.readOnly.option) //只读请求的相关摄者

}

// 日志新增
func (r *raft) appendEntry(es ...pb.Entry) (accepted bool) {
	//1. 获取raft节点当前最后一条日志条目的index
	li := r.raftLog.lastIndex()
	//2. 给新的日志条目设置term和index
	for i := range es {
		es[i].Term = r.Term
		es[i].Index = li + 1 + uint64(i)
	}
	// 3. 判断未提交的日志条目是不是超过限制,是的话拒绝并返回失败
	// etcd限制了leader上最多有多少未提交的条目,防止因为leader和follower之间出现网络问题时,导致条目一直累积.
	if !r.increaseUncommittedSize(es) {
		r.logger.Debugf(
			"%x appending new entries to log would exceed uncommitted entry size limit; dropping proposal",
			r.id,
		)
		// Drop the proposal.
		return false
	}
	// 4. 将日志条目追加到raftLog中
	//将日志条目追加到raftLog内存队列中,并且返回最大一条日志的index,对于leader追加日志的情况,这里返回的li肯定等于方法第1行中获取的li
	li = r.raftLog.append(es...)
	// 5. 检查并更新日志进度
	//raft的leader节点保存了所有节点的日志同步进度,这里面也包括它自己
	r.prs.Progress[r.id].MaybeUpdate(li)
	// 6. 判断是否做一次commit
	r.maybeCommit()
	return true
}

// 非leader角色的 tick函数, 每次逻辑计时器触发就会调用
func (r *raft) tickElection() {
	r.electionElapsed++ // 收到MsgBeat消息时会重置其选举计时器,从而防止节点发起新一轮选举.
	// roleUp返回是否可以被提升为leader
	// pastElectionTimeout检测当前的候选超时间是否过期
	if r.roleUp() && r.pastElectionTimeout() {
		// 自己可以被promote & election timeout 超时了,规定时间没有听到心跳发起选举；发送MsgHup// 选举超时
		r.electionElapsed = 0                           // 避免两次计时器触发,仍然走这里
		r.Step(pb.Message{From: r.id, Type: pb.MsgHup}) // 让自己选举
	}
}

// tickHeartbeat leader执行,在r.heartbeatTimeout之后发送一个MsgBeat.
func (r *raft) tickHeartbeat() {
	// 每次tick计时器触发,会调用这个函数
	r.heartbeatElapsed++
	r.electionElapsed++
	if r.electionElapsed >= r.electionTimeout { // 如果选举计时超时
		r.electionElapsed = 0 // 重置计时器
		if r.checkQuorum {    // 给自己发送一条 MsgCheckQuorum 消息,检测是否出现网络隔离
			r.Step(pb.Message{From: r.id, Type: pb.MsgCheckQuorum})
		}
		// 判断leader是否转移; leadTransferee 为Node ,不转移
		if r.state == StateLeader && r.leadTransferee != None {
			r.abortLeaderTransfer()
		}
	}

	if r.state != StateLeader {
		return
	}

	if r.heartbeatElapsed >= r.heartbeatTimeout {
		r.heartbeatElapsed = 0
		r.Step(pb.Message{From: r.id, Type: pb.MsgBeat})
	}
}

// 变成Follower       当前任期,当前leader
func (r *raft) becomeFollower(term uint64, lead uint64) {
	r.step = stepFollower
	r.reset(term)
	r.tick = r.tickElection
	r.lead = lead
	r.state = StateFollower
	r.logger.Infof("%x 成为Follower 在任期: %d", r.id, r.Term)
}

// 变成竞选者角色
func (r *raft) becomeCandidate() {
	if r.state == StateLeader {
		panic("无效的转移 [leader -> candidate]")
	}
	r.step = stepCandidate
	r.reset(r.Term + 1)
	r.tick = r.tickElection
	r.Vote = r.id // 当前节点的选票投给了谁做我Leader
	r.state = StateCandidate
	r.logger.Infof("%x 成为Candidate 在任期: %d", r.id, r.Term)
}

// 变成预竞选者角色
func (r *raft) becomePreCandidate() {
	if r.state == StateLeader {
		panic("无效的转移 [leader -> pre-candidate]")
	}
	// 变成预竞选者更新step func和state,但绝对不能增加任期和投票
	r.step = stepCandidate
	r.prs.ResetVotes() // // 清空接收到了哪些节点的投票
	r.tick = r.tickElection
	r.lead = None
	r.state = StatePreCandidate
	r.logger.Infof("%x 成为 pre-candidate在任期 %d", r.id, r.Term)
}

func (r *raft) becomeLeader() {
	// TODO(xiangli) remove the panic when the raft implementation is stable
	if r.state == StateFollower {
		panic("invalid transition [follower -> leader]")
	}
	r.step = stepLeader
	r.reset(r.Term)
	r.tick = r.tickHeartbeat
	r.lead = r.id
	r.state = StateLeader
	// Followers enter replicate mode when they've been successfully probed
	// (perhaps after having received a snapshot as a result). The leader is
	// trivially in this state. Note that r.reset() has initialized this
	// progress with the last index already.
	r.prs.Progress[r.id].BecomeReplicate()

	// Conservatively set the pendingConfIndex to the last index in the
	// log. There may or may not be a pending config change, but it's
	// safe to delay any future proposals until we commit all our
	// pending log entries, and scanning the entire tail of the log
	// could be expensive.
	r.pendingConfIndex = r.raftLog.lastIndex()

	emptyEnt := pb.Entry{Data: nil}
	if !r.appendEntry(emptyEnt) {
		// This won't happen because we just called reset() above.
		r.logger.Panic("empty entry was dropped")
	}
	// As a special case, don't count the initial empty entry towards the
	// uncommitted log quota. This is because we want to preserve the
	// behavior of allowing one entry larger than quota if the current
	// usage is zero.
	r.reduceUncommittedSize([]pb.Entry{emptyEnt})
	r.logger.Infof("%x became leader at term %d", r.id, r.Term)
}

// 开启竞选条件判断
func (r *raft) hup(t CampaignType) {
	if r.state == StateLeader {
		r.logger.Debugf("%x忽略MsgHup消息,因为已经是leader了", r.id)
		return
	}

	if !r.roleUp() {
		r.logger.Warningf("%x角色不可以提升,不能参与竞选", r.id)
		return
	}

	// 获取raftLog中已提交但未apply（ lip applied～committed） 的Entry记录
	ents, err := r.raftLog.slice(r.raftLog.applied+1, r.raftLog.committed+1, noLimit)
	if err != nil {
		r.logger.Panicf("获取没有apply日志时出现错误(%v)", err)
	}

	//检测是否有未应用的EntryConfChange记录,如果有就放弃发起选举的机会
	if n := numOfPendingConf(ents); n != 0 && r.raftLog.committed > r.raftLog.applied {
		r.logger.Warningf("%x不能参与竞选在任期 %d 因为还有 %d 应用配置要更改 ", r.id, r.Term, n)
		return
	}
	// 核对完成,开始选举
	r.logger.Infof("%x开启新的任期在任期%d", r.id, r.Term)
	r.campaign(t)
}

// campaign 开始竞选
func (r *raft) campaign(t CampaignType) {
	if !r.roleUp() {
		r.logger.Warningf("%x is 无法推动；不应该调用 campaign()", r.id)
	}
	var term uint64
	var voteMsg pb.MessageType
	if t == campaignPreElection { // pre-vote模式
		r.becomePreCandidate() // 变成预竞选者角色,更新状态、step、 但不增加任期
		voteMsg = pb.MsgPreVote
		// 在增加r.Term之前,将本节点打算增加到的任期数通过rpc发送出去
		term = r.Term + 1
	} else {
		r.becomeCandidate() // // 变成竞选者角色,更新状态、step、任期加1
		voteMsg = pb.MsgVote
		term = r.Term
	}
	// 自己给自己投票
	// pre-vote  那么Votes会置空
	//		单机 : 那么此时给自己投一票,res是VoteWon
	// 		多机:此时是VotePending
	// vote	直接给自己投票
	//		单机 : 那么此时给自己投一票,res是VoteWon
	// 		多机:此时是VotePending
	if _, _, res := r.poll(r.id, voteRespMsgType(voteMsg), true); res == quorum.VoteWon {
		// 我们在为自己投票后赢得了选举（这肯定意味着 这是一个单一的本地节点集群）.推进到下一个状态.
		if t == campaignPreElection {
			r.campaign(campaignElection)
		} else {
			r.becomeLeader()
		}
		return
	}
	// 	VotePending VoteLost 两种情况
	//	VoteLost
	var ids []uint64
	// 给节点排序
	{
		idMap := r.prs.Voters.IDs()
		ids = make([]uint64, 0, len(idMap))
		for id := range idMap {
			ids = append(ids, id)
		}
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	}
	for _, id := range ids {
		if id == r.id {
			// 不给自己投票
			continue
		}
		r.logger.Infof("%x [logterm: %d, index: %d] 发送 %s 请求到 %x在任期 %d", r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), voteMsg, id, r.Term)

		var ctx []byte
		if t == campaignTransfer { // leader开始转移
			ctx = []byte(t)
		}
		r.send(pb.Message{Term: term, To: id, Type: voteMsg, Index: r.raftLog.lastIndex(), LogTerm: r.raftLog.lastTerm(), Context: ctx})
	}
}

// 节点ID,投票响应类型,true
func (r *raft) poll(id uint64, t pb.MessageType, v bool) (granted int, rejected int, result quorum.VoteResult) {
	if v {
		r.logger.Infof("%x 收到 %s 从 %x 在任期 %d", r.id, t, id, r.Term)
	} else {
		r.logger.Infof("%x 收到 %s 拒绝消息从 %x 在任期 %d", r.id, t, id, r.Term)
	}
	r.prs.RecordVote(id, v)   // 记录投票结果
	return r.prs.TallyVotes() // 竞选情况
}

// Step 分流各种消息
func (r *raft) Step(m pb.Message) error {
	// 处理消息的期限,这可能会导致我们成为一名follower.
	switch {
	case m.Term == 0:
		// 本地消息 MsgHup、MsgProp、MsgReadindex
	case m.Term > r.Term: // 投票或预投票请求
		if m.Type == pb.MsgVote || m.Type == pb.MsgPreVote {
			force := bytes.Equal(m.Context, []byte(campaignTransfer))
			inLease := r.checkQuorum && r.lead != None && r.electionElapsed < r.electionTimeout
			if !force && inLease {
				// If a etcd receives a RequestVote request within the minimum election timeout
				// of hearing from a current leader, it does not update its term or grant its vote
				r.logger.Infof("%x [logterm: %d, index: %d, vote: %x] ignored %s from %x [logterm: %d, index: %d] at term %d: lease is not expired (remaining ticks: %d)",
					r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), r.Vote, m.Type, m.From, m.LogTerm, m.Index, r.Term, r.electionTimeout-r.electionElapsed)
				return nil
			}
		}
		switch {
		case m.Type == pb.MsgPreVote:
			// Never change our term in response to a PreVote
		case m.Type == pb.MsgPreVoteResp && !m.Reject:
			// We send pre-vote requests with a term in our future. If the
			// pre-vote is granted, we will increment our term when we get a
			// quorum. If it is not, the term comes from the localNode that
			// rejected our vote so we should become a follower at the new
			// term.
		default:
			r.logger.Infof("%x [term: %d] received a %s message with higher term from %x [term: %d]",
				r.id, r.Term, m.Type, m.From, m.Term)
			if m.Type == pb.MsgApp || m.Type == pb.MsgHeartbeat || m.Type == pb.MsgSnap {
				r.becomeFollower(m.Term, m.From)
			} else {
				r.becomeFollower(m.Term, None)
			}
		}

	case m.Term < r.Term:
		if (r.checkQuorum || r.preVote) && (m.Type == pb.MsgHeartbeat || m.Type == pb.MsgApp) {
			// We have received messages from a leader at a lower term. It is possible
			// that these messages were simply delayed in the network, but this could
			// also mean that this localNode has advanced its term number during a network
			// partition, and it is now unable to either win an election or to rejoin
			// the majority on the old term. If checkQuorum is false, this will be
			// handled by incrementing term numbers in response to MsgVote with a
			// higher term, but if checkQuorum is true we may not advance the term on
			// MsgVote and must generate other messages to advance the term. The net
			// result of these two features is to minimize the disruption caused by
			// nodes that have been removed from the cluster's configuration: a
			// removed localNode will send MsgVotes (or MsgPreVotes) which will be ignored,
			// but it will not receive MsgApp or MsgHeartbeat, so it will not create
			// disruptive term increases, by notifying leader of this localNode's activeness.
			// The above comments also true for Pre-Vote
			//
			// When follower gets isolated, it soon starts an election ending
			// up with a higher term than leader, although it won't receive enough
			// votes to win the election. When it regains connectivity, this response
			// with "pb.MsgAppResp" of higher term would force leader to step down.
			// However, this disruption is inevitable to free this stuck localNode with
			// fresh election. This can be prevented with Pre-Vote phase.
			r.send(pb.Message{To: m.From, Type: pb.MsgAppResp})
		} else if m.Type == pb.MsgPreVote {
			// Before Pre-Vote enable, there may have candidate with higher term,
			// but less log. After update to Pre-Vote, the cluster may deadlock if
			// we drop messages with a lower term.
			r.logger.Infof("%x [logterm: %d, index: %d, vote: %x] rejected %s from %x [logterm: %d, index: %d] at term %d",
				r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), r.Vote, m.Type, m.From, m.LogTerm, m.Index, r.Term)
			r.send(pb.Message{To: m.From, Term: r.Term, Type: pb.MsgPreVoteResp, Reject: true})
		} else {
			// ignore other cases
			r.logger.Infof("%x [term: %d] ignored a %s message with lower term from %x [term: %d]",
				r.id, r.Term, m.Type, m.From, m.Term)
		}
		return nil
	}

	switch m.Type {
	case pb.MsgHup: // 没有leader时会触发, 开始选举
		if r.preVote { // PreVote 是否启用PreVote
			r.hup(campaignPreElection)
		} else {
			r.hup(campaignElection)
		}

	case pb.MsgVote, pb.MsgPreVote:
		// 当前节点在参与投票时,会综合下面几个条件决定是否投票（在Raft协议的介绍中也捉到过）．
		// 1. 投票情况是已经投过了
		// 2. 没投过并且没有leader
		// 3. 预投票并且term大
		canVote := r.Vote == m.From || (r.Vote == None && r.lead == None) || (m.Type == pb.MsgPreVote && m.Term > r.Term)
		if canVote && r.raftLog.isUpToDate(m.Index, m.LogTerm) {
			// Note: it turns out that that learners必须是allowed to cast votes.
			// This seems counter- intuitive but is necessary in the situation in which
			// a learner has been promoted (i.e. is now a voter) but has not learned
			// about this yet.
			// For example, consider a group in which id=1 is a learner and id=2 and
			// id=3 are voters. A configuration change promoting 1 can be committed on
			// the quorum `{2,3}` without the config change being appended to the
			// learner's log. If the leader (say 2) fails, there are de facto two
			// voters remaining. Only 3 can win an election (due to its log containing
			// all committed entries), but to do so it will need 1 to vote. But 1
			// considers itself a learner and will continue to do so until 3 has
			// stepped up as leader, replicates the conf change to 1, and 1 applies it.
			// Ultimately, by receiving a request to vote, the learner realizes that
			// the candidate believes it to be a voter, and that it should act
			// accordingly. The candidate's config may be stale, too; but in that case
			// it won't win the election, at least in the absence of the bug discussed
			// in:
			// https://github.com/etcd-io/etcd/issues/7625#issuecomment-488798263.
			r.logger.Infof("%x [logterm: %d, index: %d, vote: %x] cast %s for %x [logterm: %d, index: %d] at term %d",
				r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), r.Vote, m.Type, m.From, m.LogTerm, m.Index, r.Term)
			// When responding to Msg{Pre,}Vote messages we include the term
			// from the message, not the local term. To see why, consider the
			// case where a single localNode was previously partitioned away and
			// it's local term is now out of date. If we include the local term
			// (recall that for pre-votes we don't update the local term), the
			// (pre-)campaigning localNode on the other end will proceed to ignore
			// the message (it ignores all out of date messages).
			// The term in the original message and current local term are the
			// same in the case of regular votes, but different for pre-votes.
			r.send(pb.Message{To: m.From, Term: m.Term, Type: voteRespMsgType(m.Type)})
			if m.Type == pb.MsgVote {
				// Only record real votes.
				r.electionElapsed = 0
				r.Vote = m.From // 当前节点的选票投给了谁做我Leader
			}
		} else {
			//不满足上述投赞同票条件时,当前节点会返回拒绝票(响应消息中的Reject字段会设立成true)
			r.logger.Infof("%x [logterm: %d, index: %d, vote: %x] 拒绝来自投票请求 %s %x [logterm: %d, index: %d] 当前任期 %d",
				r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), r.Vote, m.Type, m.From, m.LogTerm, m.Index, r.Term)
			r.send(pb.Message{To: m.From, Term: r.Term, Type: voteRespMsgType(m.Type), Reject: true})
		}

	default:
		err := r.step(r, m) // 当前节点是Follower状态,raft.step字段指向stepFollower()函数
		if err != nil {
			return err
		}
	}
	return nil
}

type stepFunc func(r *raft, m pb.Message) error

func stepLeader(r *raft, m pb.Message) error {
	// 这些消息不会处理These message types do not require any progress for m.From.
	switch m.Type {
	case pb.MsgBeat: // leader专属
		r.bcastHeartbeat()
		return nil
		//--------------------其他消息处理----------------------
	case pb.MsgCheckQuorum: // leader专属
		// 将 leader 自己的 RecentActive 状态设置为 true
		if pr := r.prs.Progress[r.id]; pr != nil {
			pr.RecentActive = true
		}
		if !r.prs.QuorumActive() {
			// 如果当前 leader 发现其不满足 quorum 的条件,则说明该 leader 有可能处于隔离状态,step down
			r.logger.Warningf("%x stepped down to follower since quorum is not active", r.id)
			r.becomeFollower(r.Term, None)
		}
		// Mark everyone (but ourselves) as inactive in preparation for the next
		// CheckQuorum.
		r.prs.Visit(func(id uint64, pr *tracker.Progress) {
			if id != r.id {
				pr.RecentActive = false
			}
		})
		return nil
	case pb.MsgProp: // leader、Candidate、follower专属
		if len(m.Entries) == 0 {
			r.logger.Panicf("%x stepped empty MsgProp", r.id)
		}
		if r.prs.Progress[r.id] == nil {
			// 判断当前节点是不是已经被从集群中移除了
			return ErrProposalDropped
		}
		if r.leadTransferee != None {
			// 如果正在进行leader切换,拒绝写入
			r.logger.Debugf("%x [term %d]  // 如果正在进行leader切换,拒绝写入 %x  ", r.id, r.Term, r.leadTransferee)
			return ErrProposalDropped
		}

		for i := range m.Entries { //判断是否有配置变更的日志,有的话做一些特殊处理
			e := &m.Entries[i]
			var cc pb.ConfChangeI
			if e.Type == pb.EntryConfChange {
				var ccc pb.ConfChange
				if err := ccc.Unmarshal(e.Data); err != nil {
					panic(err)
				}
				cc = ccc
			} else if e.Type == pb.EntryConfChangeV2 {
				var ccc pb.ConfChangeV2
				if err := ccc.Unmarshal(e.Data); err != nil {
					panic(err)
				}
				cc = ccc
			}
			if cc != nil {
				alreadyPending := r.pendingConfIndex > r.raftLog.applied
				alreadyJoint := len(r.prs.Config.Voters[1]) > 0
				wantsLeaveJoint := len(cc.AsV2().Changes) == 0

				var refused string
				if alreadyPending {
					refused = fmt.Sprintf("possible unapplied conf change at index %d (applied to %d)", r.pendingConfIndex, r.raftLog.applied)
				} else if alreadyJoint && !wantsLeaveJoint {
					refused = "must transition out of joint config first"
				} else if !alreadyJoint && wantsLeaveJoint {
					refused = "not in joint state; refusing empty conf change"
				}

				if refused != "" {
					r.logger.Infof("%x ignoring conf change %v at config %s: %s", r.id, cc, r.prs.Config, refused)
					m.Entries[i] = pb.Entry{Type: pb.EntryNormal}
				} else {
					r.pendingConfIndex = r.raftLog.lastIndex() + uint64(i) + 1
				}
			}
		}
		//将日志追加到raft状态机中
		if !r.appendEntry(m.Entries...) {
			return ErrProposalDropped
		}
		// 发送日志给集群其它节点
		r.bcastAppend()
		return nil
	case pb.MsgReadIndex:
		// only one voting member (the leader) in the cluster
		if r.prs.IsSingleton() {
			if resp := r.responseToReadIndexReq(m, r.raftLog.committed); resp.To != None {
				r.send(resp)
			}
			return nil
		}

		// Postpone read only request when this leader has not committed
		// any log entry at its term.
		if !r.committedEntryInCurrentTerm() {
			r.pendingReadIndexMessages = append(r.pendingReadIndexMessages, m)
			return nil
		}

		sendMsgReadIndexResponse(r, m)

		return nil
	}
	// 根据消息的From字段获取对应的Progress实例,为后面的消息处理做准备

	pr := r.prs.Progress[m.From]
	if pr == nil {
		r.logger.Debugf("%x no progress available for %x", r.id, m.From)
		return nil
	}
	switch m.Type {
	case pb.MsgAppResp:
		// Leader节点在向Follower广播日志后,就一直在等待follower的MsgAppResp消息,收到后还是会进到stepLeader函数.
		// 更新对应Progress实例的RecentActive字段,从Leader节点的角度来看,MsgAppResp消息的发送节点还是存活的

		pr.RecentActive = true

		if m.Reject { //MsgApp 消息被拒绝;如果收到的是reject消息,则根据follower反馈的index重新发送日志
			_ = r.handleAppendEntries // 含有拒绝的逻辑
			r.logger.Debugf("%x 收到 MsgAppResp(rejected, hint: (index %d, term %d)) from %x for index %d", r.id, m.RejectHint, m.LogTerm, m.From, m.Index)
			// 发送的是  9 5
			nextProbeIdx := m.RejectHint // 拒绝之处的日志索引 6
			if m.LogTerm > 0 {           // 拒绝之处的日志任期 2
				//  example 1
				//   idx        1 2 3 4 5 6 7 8 9
				//              -----------------
				//   term (L)   1 3 3 3 5 5 5 5 5
				//   term (F)   1 1 1 1 2 2
				nextProbeIdx = r.raftLog.findConflictByTerm(m.RejectHint, m.LogTerm) // 下一次直接发送索引为1的消息  🐂
			}
			//通过MsgAppResp消息携带的信息及对应的Progress状态,重新设立其Next
			if pr.MaybeDecrTo(m.Index, nextProbeIdx) { // leader是否降低对该节点索引记录 ---- > 降低索引数据
				r.logger.Debugf("%x回滚进度  节点:%x to [%s]", r.id, m.From, pr)
				if pr.State == tracker.StateReplicate {
					pr.BecomeProbe()
				}
				r.sendAppend(m.From)
			}
		} else {
			//走到这说明   之前发送的MsgApp消息已经被对方的Follower节点接收（Entry记录被成功追加）
			oldPaused := pr.IsPaused()
			// m.Index: 对应Follower节点收到的raftLog中最后一条Entry记录的索引,
			if pr.MaybeUpdate(m.Index) { // 更新pr的进度

				switch {
				case pr.State == tracker.StateProbe:
					// 一旦更新了pr状态,就不再进行探测
					pr.BecomeReplicate()
				case pr.State == tracker.StateSnapshot && pr.Match >= pr.PendingSnapshot:
					// 复制完快照
					r.logger.Debugf("%x 从需要的快照中恢复,恢复发送复制信息到 %x [%s]", r.id, m.From, pr)
					pr.BecomeProbe()     // 再探测一次
					pr.BecomeReplicate() // 正常发送日志
				case pr.State == tracker.StateReplicate:
					//之前向某个Follower节点发送MsgApp消息时,会将其相关信息保存到对应的
					//Progress.ins中,在这里收到相应的MsgAppResp响应之后,会将其从ins中删除,
					//这样可以实现了限流的效采,避免网络出现延迟时,继续发送消息,从而导致网络更加拥堵
					pr.Inflights.FreeLE(m.Index)
				}
				//如果进度有更新,判断并更新commitIndex
				//收到一个Follower节点的MsgAppResp消息之后,除了修改相应的Match和Next,还会尝试更新raftLog.committed,因为有些Entry记录可能在此次复制中被保存到了
				//半数以上的节点中,raft.maybeCommit（）方法在前面已经分析过了
				if r.maybeCommit() {
					// committed index has progressed for the term, so it is safe
					// to respond to pending read index requests
					releasePendingReadIndexMessages(r)
					//向所有节点发送MsgApp消息,注意,此次MsgApp消息的Commit字段与上次MsgApp消息已经不同,raft.bcastAppend()方法前面已经讲过

					r.bcastAppend()
				} else if oldPaused {
					//之前是pause状态,现在可以任性地发消息了
					//之前Leader节点暂停向该Follower节点发送消息,收到MsgAppResp消息后,在上述代码中已经重立了相应状态,所以可以继续发送MsgApp消息
					// If we were paused before, this localNode may be missing the
					// latest commit index, so send it.
					r.sendAppend(m.From)
				}
				// We've updated flow control information above, which may
				// allow us to send multiple (size-limited) in-flight messages
				// at once (such as when transitioning from probe to
				// replicate, or when freeTo() covers multiple messages). If
				// we have more entries to send, send as many messages as we
				// can (without sending empty messages for the commit index)
				// 循环发送所有剩余的日志给follower
				for r.maybeSendAppend(m.From, false) {
				}
				// 是否正在进行leader转移
				if m.From == r.leadTransferee && pr.Match == r.raftLog.lastIndex() {
					r.logger.Infof("%x sent MsgTimeoutNow to %x after received MsgAppResp", r.id, m.From)
					r.sendTimeoutNow(m.From)
				}
			}
		}
	case pb.MsgHeartbeatResp:
		pr.RecentActive = true
		pr.ProbeSent = false

		// free one slot for the full inflights window to allow progress.
		if pr.State == tracker.StateReplicate && pr.Inflights.Full() {
			pr.Inflights.FreeFirstOne()
		}
		if pr.Match < r.raftLog.lastIndex() {
			r.sendAppend(m.From)
		}

		if r.readOnly.option != ReadOnlySafe || len(m.Context) == 0 {
			return nil
		}

		if r.prs.Voters.VoteResult(r.readOnly.recvAck(m.From, m.Context)) != quorum.VoteWon {
			return nil
		}

		rss := r.readOnly.advance(m)
		for _, rs := range rss {
			if resp := r.responseToReadIndexReq(rs.req, rs.index); resp.To != None {
				r.send(resp)
			}
		}
	case pb.MsgSnapStatus:
		if pr.State != tracker.StateSnapshot {
			return nil
		}
		// TODO(tbg): this code is very similar to the snapshot handling in
		// MsgAppResp above. In fact, the code there is more correct than the
		// code here and should likely be updated to match (or even better, the
		// logic pulled into a newly created Progress state machine handler).
		if !m.Reject {
			pr.BecomeProbe()
			r.logger.Debugf("%x snapshot succeeded, resumed sending replication messages to %x [%s]", r.id, m.From, pr)
		} else {
			// NB: the order here matters or we'll be probing erroneously from
			// the snapshot index, but the snapshot never applied.
			pr.PendingSnapshot = 0
			pr.BecomeProbe()
			r.logger.Debugf("%x snapshot failed, resumed sending replication messages to %x [%s]", r.id, m.From, pr)
		}
		// If snapshot finish, wait for the MsgAppResp from the remote localNode before sending
		// out the next MsgApp.
		// If snapshot failure, wait for a heartbeat interval before next try
		pr.ProbeSent = true
	case pb.MsgUnreachable:
		// During optimistic replication, if the remote becomes unreachable,
		// there is huge probability that a MsgApp is lost.
		if pr.State == tracker.StateReplicate {
			pr.BecomeProbe()
		}
		r.logger.Debugf("%x failed to send message to %x because it is unreachable [%s]", r.id, m.From, pr)
	case pb.MsgTransferLeader:
		if pr.IsLearner {
			r.logger.Debugf("%x is learner. Ignored transferring leadership", r.id)
			return nil
		}
		leadTransferee := m.From
		lastLeadTransferee := r.leadTransferee
		if lastLeadTransferee != None {
			if lastLeadTransferee == leadTransferee {
				r.logger.Infof("%x [term %d] transfer leadership to %x is in progress, ignores request to same localNode %x",
					r.id, r.Term, leadTransferee, leadTransferee)
				return nil
			}
			r.abortLeaderTransfer()
			r.logger.Infof("%x [term %d] abort previous transferring leadership to %x", r.id, r.Term, lastLeadTransferee)
		}
		if leadTransferee == r.id {
			r.logger.Debugf("%x is already leader. Ignored transferring leadership to self", r.id)
			return nil
		}
		// Transfer leadership to third party.
		r.logger.Infof("%x [term %d] starts to transfer leadership to %x", r.id, r.Term, leadTransferee)
		// Transfer leadership should be finished in one electionTimeout, so reset r.electionElapsed.
		r.electionElapsed = 0
		r.leadTransferee = leadTransferee
		if pr.Match == r.raftLog.lastIndex() {
			r.sendTimeoutNow(leadTransferee)
			r.logger.Infof("%x sends MsgTimeoutNow to %x immediately as %x already has up-to-date log", r.id, leadTransferee, leadTransferee)
		} else {
			r.sendAppend(leadTransferee)
		}
	}
	return nil
}

// stepCandidate 两个阶段都会调用 StateCandidate and StatePreCandidate;
// 区别在于对投票请求的处理
func stepCandidate(r *raft, m pb.Message) error {
	// 根据当前节点的状态决定其能够处理的选举响应消息的类型
	var myVoteRespType pb.MessageType
	if r.state == StatePreCandidate {
		myVoteRespType = pb.MsgPreVoteResp
	} else {
		myVoteRespType = pb.MsgVoteResp
	}
	switch m.Type {
	case pb.MsgProp: // leader、Candidate、follower专属
		r.logger.Infof("%x no leader at term %d; dropping proposal", r.id, r.Term)
		return ErrProposalDropped
	case pb.MsgApp:
		r.becomeFollower(m.Term, m.From)
		r.handleAppendEntries(m)
	case pb.MsgHeartbeat: // √
		r.becomeFollower(m.Term, m.From)
		r.handleHeartbeat(m)
	case pb.MsgSnap:
		r.becomeFollower(m.Term, m.From)
		r.handleSnapshot(m)
	case myVoteRespType: // ✅
		// 投票、预投票
		//处理收到的选举响应消息,当前示例中处理的是MsgPreVoteResp消息
		gr, rj, res := r.poll(m.From, m.Type, !m.Reject) // 计算当前收到多少投票
		r.logger.Infof("%x 收到了 %d %s 同已投票 %d 拒绝投票", r.id, gr, m.Type, rj)
		// 投票数、拒绝数 过半判定
		switch res {
		case quorum.VoteWon: // 如果quorum 都选择了投票
			if r.state == StatePreCandidate {
				r.campaign(campaignElection) // 预投票发起正式投票
			} else {
				r.becomeLeader() // 当前节点切换成为Leader状态, 其中会重置每个节点对应的Next和Match两个索引,
				r.bcastAppend()  // 向集群中其他节点广播MsgApp消息
			}
		case quorum.VoteLost: // 集票失败,转为 follower
			r.becomeFollower(r.Term, None) // 注意,此时任期没有改变
		}
	case pb.MsgTimeoutNow:
		r.logger.Debugf("%x [term %d state %v] ignored MsgTimeoutNow from %x", r.id, r.Term, r.state, m.From)
	}
	return nil
}

// follower 的功能
func stepFollower(r *raft, m pb.Message) error {
	switch m.Type {
	case pb.MsgProp: // leader、Candidate、follower专属
		if r.lead == None {
			r.logger.Infof("%x no leader at term %d; dropping proposal", r.id, r.Term)
			return ErrProposalDropped
		} else if r.disableProposalForwarding {
			r.logger.Infof("%x not forwarding to leader %x at term %d; dropping proposal", r.id, r.lead, r.Term)
			return ErrProposalDropped
		}
		m.To = r.lead
		r.send(m)
	case pb.MsgApp:
		// Leader节点处理完命令后,发送日志和持久化操作都是异步进行的,但是这不代表leader已经收到回复.
		// Raft协议要求在返回leader成功的时候,日志一定已经提交了,所以Leader需要等待超过半数的Follower节点处理完日志并反馈,下面先看一下Follower的日志处理.
		// 日志消息到达Follower后,也是由EtcdServer.Process()方法来处理,最终会进到Raft模块的stepFollower()函数中.

		// 重置心跳计数
		r.electionElapsed = 0
		// 设置Leader
		r.lead = m.From
		// 处理日志条目
		r.handleAppendEntries(m)
	case pb.MsgHeartbeat:
		r.electionElapsed = 0
		r.lead = m.From
		r.handleHeartbeat(m)
	case pb.MsgSnap:
		r.electionElapsed = 0
		r.lead = m.From
		r.handleSnapshot(m)
	case pb.MsgTransferLeader:
		if r.lead == None {
			r.logger.Infof("%x no leader at term %d; dropping leader transfer msg", r.id, r.Term)
			return nil
		}
		m.To = r.lead
		r.send(m)
	case pb.MsgTimeoutNow:
		r.logger.Infof("%x [term %d] received MsgTimeoutNow from %x and starts an election to get leadership.", r.id, r.Term, m.From)
		// Leadership transfers never use pre-vote even if r.preVote is true; we
		// know we are not recovering from a partition so there is no need for the
		// extra round trip.
		r.hup(campaignTransfer)
	case pb.MsgReadIndex:
		if r.lead == None {
			r.logger.Infof("%x no leader at term %d; dropping index reading msg", r.id, r.Term)
			return nil
		}
		m.To = r.lead
		r.send(m)
	case pb.MsgReadIndexResp:
		if len(m.Entries) != 1 {
			r.logger.Errorf("%x invalid format of MsgReadIndexResp from %x, entries count: %d", r.id, m.From, len(m.Entries))
			return nil
		}
		r.readStates = append(r.readStates, ReadState{Index: m.Index, RequestCtx: m.Entries[0].Data})
	}
	return nil
}

// 处理日志
func (r *raft) handleAppendEntries(m pb.Message) {
	// 在leader在发消息时,也会将消息写入本地日志文件中,不会等待follower确认
	// 判断是否是过时的消息; 日志索引 小于本地已经commit的消息
	if m.Index < r.raftLog.committed {
		r.send(pb.Message{To: m.From, Type: pb.MsgAppResp, Index: r.raftLog.committed})
		return
	}
	// 会进行一致性检查;尝试将消息携带的Entry记录追加到raftLog中
	// m.Index:携带的日志的最小日志索引, m.LogTerm:携带的第一条日志任期, m.Commit:leader记录的本机点已经commit的日志索引
	// m.Entries... 真正的日志数据
	if mlastIndex, ok := r.raftLog.maybeAppend(m.Index, m.LogTerm, m.Commit, m.Entries...); ok {
		// 返回收到的最后一条日志的索引,这样Leader节点就可以根据此值更新其对应的Next和Match值
		r.send(pb.Message{To: m.From, Type: pb.MsgAppResp, Index: mlastIndex})
	} else {
		// 收到的日志索引任期不满足以下条件:任期一样,日志索引比lastIndex大1

		// 上面的maybeAppend() 方法只会将日志存储到RaftLog维护的内存队列中,
		// 日志的持久化是异步进行的,这个和Leader节点的存储WAL逻辑基本相同.
		// 有一点区别就是follower节点正式发送MsgAppResp消息会在wal保存成功后
		// 而leader节点是先发送消息,后保存的wal.

		//   idx        1 2 3 4 5 6 7 8 9
		//              -----------------
		//   term (L)   1 3 3 3 5 5 5 5 5
		//   term (F)   1 1 1 1 2 2
		// extern 当flower多一些未commit数据时, Leader是如何精准地找到每个Follower 与其日志条目首个不一致的那个槽位的呢
		// Follower 将之后的删除,重新同步leader之后的数据
		// 如采追加记录失败,则将失/败信息返回给Leader节点(即MsgAppResp 消息的Reject字段为true),
		// 同时返回的还有一些提示信息(RejectHint字段保存了当前节点raftLog中最后一条记录的索引)

		index, err := r.raftLog.term(m.Index) // 判断leader传过来的index在本地是否有存储
		r.logger.Debugf("%x [logterm: %d, index: %d]拒绝消息MsgApp [logterm: %d, index: %d] from %x",
			r.id, r.raftLog.zeroTermOnErrCompacted(index, err), m.Index, m.LogTerm, m.Index, m.From)
		// 向leader返回一个关于两个日志可能出现分歧关于 index 和 term 的提示.
		// if m.LogTerm >= term &&  m.Index >= index 可以跳过一些follower拥有的未提交数据
		hintIndex := min(m.Index, r.raftLog.lastIndex())               // 发来的消息最小索引与当前最新消息, 一般来说后者会比较小,6
		hintIndex = r.raftLog.findConflictByTerm(hintIndex, m.LogTerm) // 核心逻辑
		hintTerm, err := r.raftLog.term(hintIndex)
		if err != nil {
			panic(fmt.Sprintf("term(%d)必须是valid, but got %v", hintIndex, err))
		}
		r.send(pb.Message{
			To:         m.From,
			Type:       pb.MsgAppResp,
			Index:      m.Index,
			Reject:     true,
			RejectHint: hintIndex,
			LogTerm:    hintTerm,
		})
	}
}

// 处理leader发送来的心跳信息   【follower、Candidate】
func (r *raft) handleHeartbeat(m pb.Message) {
	// 把msg中的commit提交,commit是只增不减的
	r.raftLog.commitTo(m.Commit) // leader commit 了,follower再commit
	//发送Response给Leader   按照raft协议的要求带上自己日志的进度.
	r.send(pb.Message{To: m.From, Type: pb.MsgHeartbeatResp, Context: m.Context})
}

func (r *raft) handleSnapshot(m pb.Message) {
	sindex, sterm := m.Snapshot.Metadata.Index, m.Snapshot.Metadata.Term
	if r.restore(m.Snapshot) {
		r.logger.Infof("%x [commit: %d] restored snapshot [index: %d, term: %d]",
			r.id, r.raftLog.committed, sindex, sterm)
		r.send(pb.Message{To: m.From, Type: pb.MsgAppResp, Index: r.raftLog.lastIndex()})
	} else {
		r.logger.Infof("%x [commit: %d] ignored snapshot [index: %d, term: %d]",
			r.id, r.raftLog.committed, sindex, sterm)
		r.send(pb.Message{To: m.From, Type: pb.MsgAppResp, Index: r.raftLog.committed})
	}
}

// restore recovers the state machine from a snapshot. It restores the log and the
// configuration of state machine. If this method returns false, the snapshot was
// ignored, either because it was obsolete or because of an error.
func (r *raft) restore(s pb.Snapshot) bool {
	if s.Metadata.Index <= r.raftLog.committed {
		return false
	}
	if r.state != StateFollower {
		// This is defense-in-depth: if the leader somehow ended up applying a
		// snapshot, it could move into a new term without moving into a
		// follower state. This should never fire, but if it did, we'd have
		// prevented damage by returning early, so log only a loud warning.
		//
		// At the time of writing, the instance is guaranteed to be in follower
		// state when this method is called.
		r.logger.Warningf("%x attempted to restore snapshot as leader; should never happen", r.id)
		r.becomeFollower(r.Term+1, None)
		return false
	}

	// More defense-in-depth: throw away snapshot if recipient is not in the
	// config. This shouldn't ever happen (at the time of writing) but lots of
	// code here and there assumes that r.id is in the progress tracker.
	found := false
	cs := s.Metadata.ConfState

	for _, set := range [][]uint64{
		cs.Voters,
		cs.Learners,
		cs.VotersOutgoing,
		// `LearnersNext` doesn't need to be checked. According to the rules, if a peer in
		// `LearnersNext`, it has to be in `VotersOutgoing`.
	} {
		for _, id := range set {
			if id == r.id {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		r.logger.Warningf(
			"%x attempted to restore snapshot but it is not in the ConfState %v; should never happen",
			r.id, cs,
		)
		return false
	}

	// Now go ahead and actually restore.

	if r.raftLog.matchTerm(s.Metadata.Index, s.Metadata.Term) {
		r.logger.Infof("%x [commit: %d, lastindex: %d, lastterm: %d] fast-forwarded commit to snapshot [index: %d, term: %d]",
			r.id, r.raftLog.committed, r.raftLog.lastIndex(), r.raftLog.lastTerm(), s.Metadata.Index, s.Metadata.Term)
		r.raftLog.commitTo(s.Metadata.Index)
		return false
	}

	r.raftLog.restore(s)

	// Reset the configuration and add the (potentially updated) peers in anew.
	r.prs = tracker.MakeProgressTracker(r.prs.MaxInflight)
	cfg, prs, err := confchange.Restore(confchange.Changer{
		Tracker:   r.prs,
		LastIndex: r.raftLog.lastIndex(),
	}, cs)

	if err != nil {
		// This should never happen. Either there's a bug in our config change
		// handling or the client corrupted the conf change.
		panic(fmt.Sprintf("unable to restore config %+v: %s", cs, err))
	}

	assertConfStatesEquivalent(r.logger, cs, r.switchToConfig(cfg, prs))

	pr := r.prs.Progress[r.id]
	pr.MaybeUpdate(pr.Next - 1) // TODO(tbg): this is untested and likely unneeded

	r.logger.Infof("%x [commit: %d, lastindex: %d, lastterm: %d] restored snapshot [index: %d, term: %d]",
		r.id, r.raftLog.committed, r.raftLog.lastIndex(), r.raftLog.lastTerm(), s.Metadata.Index, s.Metadata.Term)
	return true
}

// roleUp 是否可以被提升为leader.
func (r *raft) roleUp() bool {
	pr := r.prs.Progress[r.id] // 是本节点raft的身份
	// 节点不是learner 且 没有正在应用快照
	return pr != nil && !pr.IsLearner && !r.raftLog.hasPendingSnapshot()
}

// todo 看不懂
func (r *raft) applyConfChange(cc pb.ConfChangeV2) pb.ConfState {
	cfg, prs, err := func() (tracker.Config, tracker.ProgressMap, error) {
		changer := confchange.Changer{
			Tracker:   r.prs,
			LastIndex: r.raftLog.lastIndex(),
		}
		// todo
		if cc.LeaveJoint() { // 节点离开
			return changer.LeaveJoint()
		} else if autoLeave, ok := cc.EnterJoint(); ok { //
			return changer.EnterJoint(autoLeave, cc.Changes...)
		}
		return changer.Simple(cc.Changes...)
	}()

	if err != nil {
		panic(err)
	}

	return r.switchToConfig(cfg, prs)
}

// switchToConfig 重新配置这个节点以使用所提供的配置.它更新内存中的状态,并在必要时进行额外的操作,
// 如对删除节点或改变的法定人数作出反应.要求.输入通常来自于恢复一个ConfState或应用一个ConfChange.
func (r *raft) switchToConfig(cfg tracker.Config, prs tracker.ProgressMap) pb.ConfState {
	r.prs.Config = cfg
	r.prs.Progress = prs

	r.logger.Infof("%x 切换配置 %s", r.id, r.prs.Config)
	cs := r.prs.ConfState()
	pr, ok := r.prs.Progress[r.id]

	// Update whether the localNode itself is a learner, resetting to false when the
	// localNode is removed.
	r.isLearner = ok && pr.IsLearner

	if (!ok || r.isLearner) && r.state == StateLeader {
		// This localNode is leader and was removed or demoted. We prevent demotions
		// at the time writing but hypothetically we handle them the same way as
		// removing the leader: stepping down into the next Term.
		//
		// TODO(tbg): step down (for sanity) and ask follower with largest Match
		// to TimeoutNow (to avoid interruption). This might still drop some
		// proposals but it's better than nothing.
		//
		// TODO(tbg): test this branch. It is untested at the time of writing.
		return cs
	}

	// The remaining steps only make sense if this localNode is the leader and there
	// are other nodes.
	if r.state != StateLeader || len(cs.Voters) == 0 {
		return cs
	}

	if r.maybeCommit() {
		// If the configuration change means that more entries are committed now,
		// broadcast/append to everyone in the updated config.
		r.bcastAppend()
	} else {
		// Otherwise, still probe the newly added replicas; there's no reason to
		// let them wait out a heartbeat interval (or the next incoming
		// proposal).
		r.prs.Visit(func(id uint64, pr *tracker.Progress) {
			r.maybeSendAppend(id, false /* sendIfEmpty */)
		})
	}
	// If the the leadTransferee was removed or demoted, abort the leadership transfer.
	if _, tOK := r.prs.Config.Voters.IDs()[r.leadTransferee]; !tOK && r.leadTransferee != 0 {
		r.abortLeaderTransfer()
	}

	return cs
}

//   所提交的state 必须在 [r.raftLog.committed,r.raftLog.lastIndex()]之间
func (r *raft) loadState(state pb.HardState) {
	if state.Commit < r.raftLog.committed || state.Commit > r.raftLog.lastIndex() {
		r.logger.Panicf("%x state.commit %d 不再指定范围内 [%d, %d]", r.id, state.Commit, r.raftLog.committed, r.raftLog.lastIndex())
	}
	r.raftLog.committed = state.Commit
	r.Term = state.Term
	r.Vote = state.Vote // 当前节点的选票投给了谁做我Leader
}

// 判断本节点是不是重新选举,因为丢失了leader
func (r *raft) pastElectionTimeout() bool {
	// 选举过期计数(electionElapsed)：主要用于follower来判断leader是不是正常工作,
	// 当follower接受到leader的心跳的时候会把electionElapsed的时候就会置为0,electionElapsed的相加是通过外部调用实现的,
	// node对外提供一个tick的接口,需要外部定时去调用,调用的周期由外部决定,每次调用就++,
	// 然后检查是否会超时,上方的tickElection就是为follower状态的定时调用函数,leader状态的定时调用函数就是向follower发送心跳.
	// 计时次数 超过了 限定的 选举次数,   规定：在randomizedElectionTimeout次数内必须收到来自leader的消息
	return r.electionElapsed >= r.randomizedElectionTimeout
}

// 设置随机选举超时
func (r *raft) resetRandomizedElectionTimeout() {
	r.randomizedElectionTimeout = r.electionTimeout + globalRand.Intn(r.electionTimeout) // 随机选举超时
}

func (r *raft) sendTimeoutNow(to uint64) {
	r.send(pb.Message{To: to, Type: pb.MsgTimeoutNow})
}

// 置空 leader转移目标
func (r *raft) abortLeaderTransfer() {
	r.leadTransferee = None
}

// committedEntryInCurrentTerm return true if the peer has committed an entry in its term.
func (r *raft) committedEntryInCurrentTerm() bool {
	return r.raftLog.zeroTermOnErrCompacted(r.raftLog.term(r.raftLog.committed)) == r.Term
}

// responseToReadIndexReq constructs a response for `req`. If `req` comes from the peer
// itself, a blank value will be returned.
func (r *raft) responseToReadIndexReq(req pb.Message, readIndex uint64) pb.Message {
	if req.From == None || req.From == r.id {
		r.readStates = append(r.readStates, ReadState{
			Index:      readIndex,
			RequestCtx: req.Entries[0].Data,
		})
		return pb.Message{}
	}
	return pb.Message{
		Type:    pb.MsgReadIndexResp,
		To:      req.From,
		Index:   readIndex,
		Entries: req.Entries,
	}
}

// increaseUncommittedSize computes the size of the proposed entries and
// determines whether they would push leader over its maxUncommittedSize limit.
// If the new entries would exceed the limit, the method returns false. If not,
// the increase in uncommitted entry size is recorded and the method returns
// true.
//
// Empty payloads are never refused. This is used both for appending an empty
// entry at a new leader's term, as well as leaving a joint configuration.
func (r *raft) increaseUncommittedSize(ents []pb.Entry) bool {
	var s uint64
	for _, e := range ents {
		s += uint64(PayloadSize(e))
	}

	if r.uncommittedSize > 0 && s > 0 && r.uncommittedSize+s > r.maxUncommittedSize {
		// If the uncommitted tail of the Raft log is empty, allow any size
		// proposal. Otherwise, limit the size of the uncommitted tail of the
		// log and drop any proposal that would push the size over the limit.
		// Note the added requirement s>0 which is used to make sure that
		// appending single empty entries to the log always succeeds, used both
		// for replicating a new leader's initial empty entry, and for
		// auto-leaving joint configurations.
		return false
	}
	r.uncommittedSize += s
	return true
}

// reduceUncommittedSize accounts for the newly committed entries by decreasing
// the uncommitted entry size limit.
func (r *raft) reduceUncommittedSize(ents []pb.Entry) {
	if r.uncommittedSize == 0 {
		// Fast-path for followers, who do not track or enforce the limit.
		return
	}

	var s uint64
	for _, e := range ents {
		s += uint64(PayloadSize(e))
	}
	if s > r.uncommittedSize {
		// uncommittedSize may underestimate the size of the uncommitted Raft
		// log tail but will never overestimate it. Saturate at 0 instead of
		// allowing overflow.
		r.uncommittedSize = 0
	} else {
		r.uncommittedSize -= s
	}
}

// 检查是否有配置变更日志
func numOfPendingConf(ents []pb.Entry) int {
	n := 0
	for i := range ents {
		if ents[i].Type == pb.EntryConfChange || ents[i].Type == pb.EntryConfChangeV2 {
			n++
		}
	}
	return n
}

func releasePendingReadIndexMessages(r *raft) {
	if !r.committedEntryInCurrentTerm() {
		r.logger.Error("pending MsgReadIndex should be released only after first commit in current term")
		return
	}

	msgs := r.pendingReadIndexMessages
	r.pendingReadIndexMessages = nil

	for _, m := range msgs {
		sendMsgReadIndexResponse(r, m)
	}
}

func sendMsgReadIndexResponse(r *raft, m pb.Message) {
	// thinking: use an internally defined context instead of the user given context.
	// We can express this in terms of the term and index instead of a user-supplied value.
	// This would allow multiple reads to piggyback on the same message.
	/*
		Leader节点检测自身在当前任期中是否已提交Entry记录,如果没有,则无法进行读取操作
	*/
	switch r.readOnly.option {
	// If more than the local vote is needed, go through a full broadcast.
	case ReadOnlySafe:
		//记录当前节点的raftLog.committed字段值,即已提交位置
		r.readOnly.addRequest(r.raftLog.committed, m)
		// The local localNode automatically acks the request.
		r.readOnly.recvAck(r.id, m.Entries[0].Data)
		r.bcastHeartbeatWithCtx(m.Entries[0].Data) //发送心跳
	case ReadOnlyLeaseBased:
		if resp := r.responseToReadIndexReq(m, r.raftLog.committed); resp.To != None {
			r.send(resp)
		}
	}
}
