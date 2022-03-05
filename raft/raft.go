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

// None 是一个占位的节点ID,在没有领导者时使用.
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

	// ElectionTick 选举超时
	ElectionTick int
	// 心跳间隔
	HeartbeatTick int

	// Storage 存储 日志项、状态
	Storage Storage //
	// Applied 提交到用户状态机的索引
	Applied uint64 // 起始为0

	// 每条消息的最大大小 ：math.MaxUint64表示无限制,0表示每条消息最多一个条目.
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
	// 心跳间隔
	if c.HeartbeatTick <= 0 {
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

type raft struct {
	id uint64 // 是本节点raft的身份

	Term uint64 // 任期
	Vote uint64 // 上一次投票的节点,Leader等于自己的id

	readStates []ReadState

	raftLog *raftLog // 当前节点的log状态信息

	maxMsgSize         uint64 // 每条消息的最大大小
	maxUncommittedSize uint64 // 每条日志最大消息体

	prs tracker.ProgressTracker // 跟踪Follower节点的状态,比如日志复制的matchIndex

	state StateType // 当前节点的状态

	// isLearner 本节点是不是learner角色
	isLearner bool

	msgs []pb.Message // 本节点要发送出去的消息

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

	checkQuorum bool // 检查需要维持的选票数,一旦小于,就会丢失leader
	preVote     bool // PreVote 是否启用PreVote

	// 选举过期计数(electionElapsed)：主要用于follower来判断leader是不是正常工作,
	// 当follower接受到leader的心跳的时候会把electionElapsed的时候就会置为0,electionElapsed的相加是通过外部调用实现的,
	// node对外提供一个tick的接口,需要外部定时去调用,调用的周期由外部决定,每次调用就++,
	// 然后检查是否会超时,上方的tickElection就是为follower状态的定时调用函数,leader状态的定时调用函数就是向follower发送心跳.
	electionElapsed int

	// 心跳过期计数(heartbeatElapsed)：用于leader判断是不是要开始发送心跳了.
	// 只要这个值超过或等于心跳超时计数(heartbeatTimeout),就会触发leader广播heartbeat信息.
	heartbeatElapsed int

	heartbeatTimeout int // 心跳间隔    ,上限     heartbeatTimeout是当前距离上次心跳的时间
	electionTimeout  int // 选举超时	   ,上限		electionElapsed是当前距离上次选举的时间

	// 随机选举超时
	randomizedElectionTimeout int
	disableProposalForwarding bool // 禁止将请求转发到leader,默认FALSE
	// 由 r.ticker = time.NewTicker(r.heartbeat) ;触发该函数的执行  r.start
	tick func() // 定时器到期执行的函数

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
		electionTimeout:           c.ElectionTick,  // 选举超时
		heartbeatTimeout:          c.HeartbeatTick, // 心跳间隔
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
	if !IsEmptyHardState(hs) { // 判断初始状态是不是空的
		r.loadState(hs) // 更新状态索引信息
	}
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
		Vote:   r.Vote,
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

	r.msgs = append(r.msgs, m) //将消息放入队列
}

// sendHeartbeat sends a heartbeat RPC to the given peer.
func (r *raft) sendHeartbeat(to uint64, ctx []byte) {
	// Attach the commit as min(to.matched, r.committed).
	// When the leader sends out heartbeat message,
	// the receiver(follower) might not be matched with the leader
	// or it might not have all the committed entries.
	// The leader MUST NOT forward the follower's commit to
	// an unmatched index.
	commit := min(r.prs.Progress[to].Match, r.raftLog.committed)
	m := pb.Message{
		To:      to,
		Type:    pb.MsgHeartbeat,
		Commit:  commit,
		Context: ctx,
	}

	r.send(m)
}

// bcastAppend 同步日志给Follower
func (r *raft) bcastAppend() {
	//遍历所有节点,给除自己外的节点发送日志Append消息
	r.prs.Visit(func(id uint64, _ *tracker.Progress) {
		if id == r.id {
			return
		}
		r.sendAppend(id)
	})
}

// OK
func (r *raft) sendAppend(to uint64) {
	r.maybeSendAppend(to, true)
}

// maybeSendAppend 向给定的peer发送一个带有新条目的追加RPC.如果有消息被发送,返回true.
// sendIfEmpty参数控制是否发送没有条目的消息("空 "消息对于传达更新的Commit索引很有用,但当我们批量发送多条消息时就不可取).
func (r *raft) maybeSendAppend(to uint64, sendIfEmpty bool) bool {
	//1. 获取对端节点当前同步进度
	pr := r.prs.Progress[to]
	if pr.IsPaused() {
		return false
	}
	m := pb.Message{}
	m.To = to
	//2. 注意这里带的term是本次发送给follower的第一条日志条目的term
	term, errt := r.raftLog.term(pr.Next - 1)
	ents, erre := r.raftLog.entries(pr.Next, r.maxMsgSize)
	if len(ents) == 0 && !sendIfEmpty {
		return false
	}

	if errt != nil || erre != nil { // send snapshot if we failed to get term or entries
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
		m.Type = pb.MsgApp
		m.Index = pr.Next - 1
		m.LogTerm = term
		m.Entries = ents
		//6. 每次发送日志或心跳都会带上最新的commitIndex
		m.Commit = r.raftLog.committed
		if n := len(m.Entries); n != 0 {
			switch pr.State {
			// optimistically increase the next when in StateReplicate
			case tracker.StateReplicate:
				last := m.Entries[n-1].Index
				pr.OptimisticUpdate(last)
				pr.Inflights.Add(last)
			case tracker.StateProbe:
				pr.ProbeSent = true
			default:
				r.logger.Panicf("%x is sending append in unhandled state %s", r.id, pr.State)
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
		r.Vote = None // 选票
	}
	r.lead = None

	r.electionElapsed = 0
	r.heartbeatElapsed = 0
	r.resetRandomizedElectionTimeout()

	r.abortLeaderTransfer()

	r.prs.ResetVotes()
	r.prs.Visit(func(id uint64, pr *tracker.Progress) {
		*pr = tracker.Progress{
			Match:     0,
			Next:      r.raftLog.lastIndex() + 1,
			Inflights: tracker.NewInflights(r.prs.MaxInflight),
			IsLearner: pr.IsLearner,
		}
		if id == r.id {
			pr.Match = r.raftLog.lastIndex()
		}
	})

	r.pendingConfIndex = 0
	r.uncommittedSize = 0
	r.readOnly = newReadOnly(r.readOnly.option)
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

// tickElection is run by followers and candidates after r.electionTimeout.
func (r *raft) tickElection() {
	r.electionElapsed++
	// 自己可以被promote & election timeout 超时了,规定时间没有听到心跳发起选举；发送MsgHup
	if r.roleUp() && r.pastElectionTimeout() { // 选举超时
		r.electionElapsed = 0
		r.Step(pb.Message{From: r.id, Type: pb.MsgHup}) // 让自己选举
	}
}

// tickHeartbeat leader执行,在r.heartbeatTimeout之后发送一个MsgBeat.
func (r *raft) tickHeartbeat() {
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

func (r *raft) becomeCandidate() {
	// TODO(xiangli) remove the panic when the raft implementation is stable
	if r.state == StateLeader {
		panic("invalid transition [leader -> candidate]")
	}
	r.step = stepCandidate
	r.reset(r.Term + 1)
	r.tick = r.tickElection
	r.Vote = r.id
	r.state = StateCandidate
	r.logger.Infof("%x 成为Candidate 在任期: %d", r.id, r.Term)
}

func (r *raft) becomePreCandidate() {
	// TODO(xiangli) remove the panic when the raft implementation is stable
	if r.state == StateLeader {
		panic("invalid transition [leader -> pre-candidate]")
	}
	// Becoming a pre-candidate changes our step functions and state,
	// but doesn't change anything else. In particular it does not increase
	// r.Term or change r.Vote.
	// 注意：预投票不更新任期
	r.step = stepCandidate
	r.prs.ResetVotes()
	r.tick = r.tickElection
	r.lead = None
	r.state = StatePreCandidate
	r.logger.Infof("%x became pre-candidate at term %d", r.id, r.Term)
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

// 开启竞选
func (r *raft) hup(t CampaignType) {
	// 主要核对是否是合法的Follower
	if r.state == StateLeader {
		r.logger.Debugf("%x忽略MsgHup消息,因为已经是leader了", r.id)
		return
	}

	if !r.roleUp() {
		r.logger.Warningf("%x角色不可以提升,不能参与竞选", r.id)
		return
	}
	ents, err := r.raftLog.slice(r.raftLog.applied+1, r.raftLog.committed+1, noLimit)
	if err != nil {
		r.logger.Panicf("获取没有apply日志时出现错误(%v)", err)
	}
	if n := numOfPendingConf(ents); n != 0 && r.raftLog.committed > r.raftLog.applied {
		r.logger.Warningf("%x不能参与竞选在任期 %d 因为还有 %d 应用配置要更改 ", r.id, r.Term, n)
		return
	}
	// 核对完成,开始选举
	r.logger.Infof("%x开启新的任期在任期%d", r.id, r.Term)
	r.campaign(t)
}

// campaign 竞选
func (r *raft) campaign(t CampaignType) {
	if !r.roleUp() {
		// This path should not be hit (callers are supposed to check), but
		// better safe than sorry.
		r.logger.Warningf("%x is unpromotable; campaign() should have been called", r.id)
	}
	var term uint64
	var voteMsg pb.MessageType
	if t == campaignPreElection {
		r.becomePreCandidate()
		voteMsg = pb.MsgPreVote
		// PreVote RPCs are sent for the next term before we've incremented r.Term.
		term = r.Term + 1
	} else {
		r.becomeCandidate()
		voteMsg = pb.MsgVote
		term = r.Term
	}
	if _, _, res := r.poll(r.id, voteRespMsgType(voteMsg), true); res == quorum.VoteWon {
		// We won the election after voting for ourselves (which must mean that
		// this is a single-localNode cluster). Advance to the next state.
		if t == campaignPreElection {
			r.campaign(campaignElection)
		} else {
			r.becomeLeader()
		}
		return
	}
	var ids []uint64
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
			continue
		}
		r.logger.Infof("%x [logterm: %d, index: %d] sent %s request to %x at term %d",
			r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), voteMsg, id, r.Term)

		var ctx []byte
		if t == campaignTransfer {
			ctx = []byte(t)
		}
		r.send(pb.Message{Term: term, To: id, Type: voteMsg, Index: r.raftLog.lastIndex(), LogTerm: r.raftLog.lastTerm(), Context: ctx})
	}
}

func (r *raft) poll(id uint64, t pb.MessageType, v bool) (granted int, rejected int, result quorum.VoteResult) {
	if v {
		r.logger.Infof("%x received %s from %x at term %d", r.id, t, id, r.Term)
	} else {
		r.logger.Infof("%x received %s rejection from %x at term %d", r.id, t, id, r.Term)
	}
	r.prs.RecordVote(id, v)
	return r.prs.TallyVotes()
}

// 分流各种消息
func (r *raft) Step(m pb.Message) error {
	// 处理消息的期限，这可能会导致我们成为一名追随者。
	switch {
	case m.Term == 0:
		// 本地消息
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
	case pb.MsgHup: // 开始选举
		if r.preVote { // PreVote 是否启用PreVote
			r.hup(campaignPreElection)
		} else {
			r.hup(campaignElection)
		}

	case pb.MsgVote, pb.MsgPreVote:
		// 我们可以投票,如果这是在重复我们已经投过的票
		// 1. 投票情况是已经投过了
		// 2. 没投过并且没有leader
		// 3. 预投票并且term大
		canVote := r.Vote == m.From ||
			// ...we haven't voted and we don't think there's a leader yet in this term...
			(r.Vote == None && r.lead == None) ||
			// ...or this is a PreVote for a future term...
			(m.Type == pb.MsgPreVote && m.Term > r.Term)
		// ...and we believe the candidate is up to date.
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
				r.Vote = m.From
			}
		} else {
			r.logger.Infof("%x [logterm: %d, index: %d, vote: %x] rejected %s from %x [logterm: %d, index: %d] at term %d",
				r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), r.Vote, m.Type, m.From, m.LogTerm, m.Index, r.Term)
			r.send(pb.Message{To: m.From, Term: r.Term, Type: voteRespMsgType(m.Type), Reject: true})
		}

	default:
		err := r.step(r, m)
		if err != nil {
			return err
		}
	}
	return nil
}

type stepFunc func(r *raft, m pb.Message) error

func stepLeader(r *raft, m pb.Message) error {
	// These message types do not require any progress for m.From.
	switch m.Type {
	case pb.MsgBeat:
		r.bcastHeartbeat()
		return nil
		//--------------------其他消息处理----------------------
	case pb.MsgCheckQuorum:
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
	case pb.MsgProp:
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

	// All other message types require a progress for m.From (pr).
	pr := r.prs.Progress[m.From]
	if pr == nil {
		r.logger.Debugf("%x no progress available for %x", r.id, m.From)
		return nil
	}
	switch m.Type {
	case pb.MsgAppResp:
		// Leader节点在向Follower广播日志后,就一直在等待follower的MsgAppResp消息,收到后还是会进到stepLeader函数.
		pr.RecentActive = true

		if m.Reject {
			//如果收到的是reject消息,则根据follower反馈的index重新发送日志
			r.logger.Debugf("%x 收到 MsgAppResp(rejected, hint: (index %d, term %d)) from %x for index %d",
				r.id, m.RejectHint, m.LogTerm, m.From, m.Index)
			nextProbeIdx := m.RejectHint
			if m.LogTerm > 0 {
				// If the follower has an uncommitted log tail, we would end up
				// probing one by one until we hit the common prefix.
				//
				// For example, if the leader has:
				//
				//   idx        1 2 3 4 5 6 7 8 9
				//              -----------------
				//   term (L)   1 3 3 3 5 5 5 5 5
				//   term (F)   1 1 1 1 2 2
				//
				// Then, after sending an append anchored at (idx=9,term=5) we
				// would receive a RejectHint of 6 and LogTerm of 2. Without the
				// code below, we would try an append at index 6, which would
				// fail again.
				//
				// However, looking only at what the leader knows about its own
				// log and the rejection hint, it is clear that a probe at index
				// 6, 5, 4, 3, and 2 must fail as well:
				//
				// For all of these indexes, the leader's log term is larger than
				// the rejection's log term. If a probe at one of these indexes
				// succeeded, its log term at that index would match the leader's,
				// i.e. 3 or 5 in this example. But the follower already told the
				// leader that it is still at term 2 at index 9, and since the
				// log term only ever goes up (within a log), this is a contradiction.
				//
				// At index 1, however, the leader can draw no such conclusion,
				// as its term 1 is not larger than the term 2 from the
				// follower's rejection. We thus probe at 1, which will succeed
				// in this example. In general, with this approach we probe at
				// most once per term found in the leader's log.
				//
				// There is a similar mechanism on the follower (implemented in
				// handleAppendEntries via a call to findConflictByTerm) that is
				// useful if the follower has a large divergent uncommitted log
				// tail[1], as in this example:
				//
				//   idx        1 2 3 4 5 6 7 8 9
				//              -----------------
				//   term (L)   1 3 3 3 3 3 3 3 7
				//   term (F)   1 3 3 4 4 5 5 5 6
				//
				// Naively, the leader would probe at idx=9, receive a rejection
				// revealing the log term of 6 at the follower. Since the leader's
				// term at the previous index is already smaller than 6, the leader-
				// side optimization discussed above is ineffective. The leader thus
				// probes at index 8 and, naively, receives a rejection for the same
				// index and log term 5. Again, the leader optimization does not improve
				// over linear probing as term 5 is above the leader's term 3 for that
				// and many preceding indexes; the leader would have to probe linearly
				// until it would finally hit index 3, where the probe would succeed.
				//
				// Instead, we apply a similar optimization on the follower. When the
				// follower receives the probe at index 8 (log term 3), it concludes
				// that all of the leader's log preceding that index has log terms of
				// 3 or below. The largest index in the follower's log with a log term
				// of 3 or below is index 3. The follower will thus return a rejection
				// for index=3, log term=3 instead. The leader's next probe will then
				// succeed at that index.
				//
				// [1]: more precisely, if the log terms in the large uncommitted
				// tail on the follower are larger than the leader's. At first,
				// it may seem unintuitive that a follower could even have such
				// a large tail, but it can happen:
				//
				// 1. Leader appends (but does not commit) entries 2 and 3, crashes.
				//   idx        1 2 3 4 5 6 7 8 9
				//              -----------------
				//   term (L)   1 2 2     [crashes]
				//   term (F)   1
				//   term (F)   1
				//
				// 2. a follower becomes leader and appends entries at term 3.
				//              -----------------
				//   term (x)   1 2 2     [down]
				//   term (F)   1 3 3 3 3
				//   term (F)   1
				//
				// 3. term 3 leader goes down, term 2 leader returns as term 4
				//    leader. It commits the log & entries at term 4.
				//
				//              -----------------
				//   term (L)   1 2 2 2
				//   term (x)   1 3 3 3 3 [down]
				//   term (F)   1
				//              -----------------
				//   term (L)   1 2 2 2 4 4 4
				//   term (F)   1 3 3 3 3 [gets probed]
				//   term (F)   1 2 2 2 4 4 4
				//
				// 4. the leader will now probe the returning follower at index
				//    7, the rejection points it at the end of the follower's log
				//    which is at a higher log term than the actually committed
				//    log.
				nextProbeIdx = r.raftLog.findConflictByTerm(m.RejectHint, m.LogTerm)
			}
			if pr.MaybeDecrTo(m.Index, nextProbeIdx) {
				r.logger.Debugf("%x decreased progress of %x to [%s]", r.id, m.From, pr)
				if pr.State == tracker.StateReplicate {
					pr.BecomeProbe()
				}
				r.sendAppend(m.From)
			}
		} else {
			//更新缓存的日志同步进度
			oldPaused := pr.IsPaused()
			if pr.MaybeUpdate(m.Index) {
				switch {
				case pr.State == tracker.StateProbe:
					pr.BecomeReplicate()
				case pr.State == tracker.StateSnapshot && pr.Match >= pr.PendingSnapshot:
					// TODO(tbg): we should also enter this branch if a snapshot is
					// received that is below pr.PendingSnapshot but which makes it
					// possible to use the log again.
					r.logger.Debugf("%x recovered from needing snapshot, resumed sending replication messages to %x [%s]", r.id, m.From, pr)
					// Transition back to replicating state via probing state
					// (which takes the snapshot into account). If we didn't
					// move to replicating state, that would only happen with
					// the next round of appends (but there may not be a next
					// round for a while, exposing an inconsistent RaftStatus).
					pr.BecomeProbe()
					pr.BecomeReplicate()
				case pr.State == tracker.StateReplicate:
					pr.Inflights.FreeLE(m.Index)
				}
				//如果进度有更新,判断并更新commitIndex
				if r.maybeCommit() {
					// committed index has progressed for the term, so it is safe
					// to respond to pending read index requests
					releasePendingReadIndexMessages(r)
					r.bcastAppend()
				} else if oldPaused {
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
	// Only handle vote responses corresponding to our candidacy (while in
	// StateCandidate, we may get stale MsgPreVoteResp messages in this term from
	// our pre-candidate state).
	var myVoteRespType pb.MessageType
	if r.state == StatePreCandidate {
		myVoteRespType = pb.MsgPreVoteResp
	} else {
		myVoteRespType = pb.MsgVoteResp
	}
	switch m.Type {
	case pb.MsgProp:
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
	case myVoteRespType:
		gr, rj, res := r.poll(m.From, m.Type, !m.Reject) // 计算当前收到多少投票
		r.logger.Infof("%x has received %d %s votes and %d vote rejections", r.id, gr, m.Type, rj)
		switch res {
		case quorum.VoteWon: // 如果quorum 都选择了投票
			if r.state == StatePreCandidate {
				r.campaign(campaignElection) // 预投票发起正式投票
			} else {
				r.becomeLeader() // 变成 leader
				r.bcastAppend()  // 发送 AppendEntries RPC
			}
		case quorum.VoteLost: // 集票失败,转为 follower
			// pb.MsgPreVoteResp contains future term of pre-candidate
			// m.Term > r.Term; reuse r.Term
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
	case pb.MsgProp:
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
	// 判断是否是过时的消息
	if m.Index < r.raftLog.committed {
		// 日志索引 小于本地已经commit的消息
		r.send(pb.Message{To: m.From, Type: pb.MsgAppResp, Index: r.raftLog.committed})
		return
	}
	//进行一致性检查,即搜索自己的日志文件中是否存在这样的日志条目,如果不存在,就向Leader
	//返回AppendEntriesRPC 失败.如果返回失败信息,就意味着Follower 发
	//现自己的日志与领导人的不一致.在失败之后,领导人会将nextlndex 递减( nextlndex －－),
	//然后重试AppendEntriesRPC,直到AppendEntriesRPC
	//返回成功为止.这才表明在nextlndex 位置的日志条目中领导人与追随者的保持一致.
	//	【当然第一次拒绝时,可以直接将差值返回,nextlndex －－x;省略n次调用】

	//当找到对应的槽位后随后, Leader 就从nextlndex号位置开始把余下的所有日志条目一次性推送给b

	if mlastIndex, ok := r.raftLog.maybeAppend(m.Index, m.LogTerm, m.Commit, m.Entries...); ok {
		// 处理成功,发送MsgAppResp给Leader
		r.send(pb.Message{To: m.From, Type: pb.MsgAppResp, Index: mlastIndex})
	} else {
		// 日志的index和Follower的lastIndex不匹配,返回reject消息; 出现原因

		//  一种是日志条目中带的term和follower的term不一致,
		//  一种是日志列表中最小的index大于follower的最大的日志index.
		// 上面的maybeAppend() 方法只会将日志存储到RaftLog维护的内存队列中,
		// 日志的持久化是异步进行的,这个和Leader节点的存储WAL逻辑基本相同.
		// 有一点区别就是follower节点正式发送MsgAppResp消息会在wal保存成功后,而leader节点是先发送消息,后保存的wal.

		// extern 当flower多一些无用数据时, Leader是如何精准地找到每个Follower 与其日志条目不一致的那个槽位的呢
		// Follower 将之后的删除,重新同步leader之后的数据

		r.logger.Debugf("%x [logterm: %d, index: %d] rejected MsgApp [logterm: %d, index: %d] from %x",
			r.id, r.raftLog.zeroTermOnErrCompacted(r.raftLog.term(m.Index)), m.Index, m.LogTerm, m.Index, m.From)

		// Return a hint to the leader about the maximum index and term that the
		// two logs could be divergent at. Do this by searching through the
		// follower's log for the maximum (index, term) pair with a term <= the
		// MsgApp's LogTerm and an index <= the MsgApp's Index. This can help
		// skip all indexes in the follower's uncommitted tail with terms
		// greater than the MsgApp's LogTerm.
		//
		// See the other caller for findConflictByTerm (in stepLeader) for a much
		// more detailed explanation of this mechanism.
		hintIndex := min(m.Index, r.raftLog.lastIndex())
		hintIndex = r.raftLog.findConflictByTerm(hintIndex, m.LogTerm)
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

// roleUp 是否可以被提升为领导者.
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
	r.Vote = state.Vote
}

//判断是不是丢失了leader
func (r *raft) pastElectionTimeout() bool {
	// 选举过期计数(electionElapsed)：主要用于follower来判断leader是不是正常工作,
	// 当follower接受到leader的心跳的时候会把electionElapsed的时候就会置为0,electionElapsed的相加是通过外部调用实现的,
	// node对外提供一个tick的接口,需要外部定时去调用,调用的周期由外部决定,每次调用就++,
	// 然后检查是否会超时,上方的tickElection就是为follower状态的定时调用函数,leader状态的定时调用函数就是向follower发送心跳.
	return r.electionElapsed >= r.randomizedElectionTimeout
}

func (r *raft) resetRandomizedElectionTimeout() {
	r.randomizedElectionTimeout = r.electionTimeout + globalRand.Intn(r.electionTimeout)
}

func (r *raft) sendTimeoutNow(to uint64) {
	r.send(pb.Message{To: to, Type: pb.MsgTimeoutNow})
}

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
