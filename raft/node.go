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
	"context"
	"errors"

	pb "github.com/ls-2018/etcd_cn/raft/raftpb"
)

type SnapshotStatus int

const (
	SnapshotFinish  SnapshotStatus = 1
	SnapshotFailure SnapshotStatus = 2
)

var (
	emptyState = pb.HardState{}

	// ErrStopped is returned by methods on Nodes that have been stopped.
	ErrStopped = errors.New("raft: stopped")
)

// SoftState 提供对日志和调试有用的状态.该状态是不稳定的,不需要持久化到WAL中.
type SoftState struct {
	Lead      uint64    // 当前leader
	RaftState StateType // 节点状态
}

func (a *SoftState) equal(b *SoftState) bool {
	return a.Lead == b.Lead && a.RaftState == b.RaftState
}

// Ready 封装了准备读取、保存到稳定存储、提交或发送至其他peer的entry和message
// Ready中的所有字段都是只读的.
// 对于这种 IO 网络密集型的应用,提高吞吐最好的手段就是批量操作,ETCD 与之相关的核心抽象就是 Ready 结构体.
type Ready struct {
	// The current volatile state of a Node.
	// SoftState will be nil if there is no update.
	// It is not required to consume or store SoftState.
	*SoftState

	// The current state of a Node to be saved to stable storage BEFORE
	// Messages are sent.
	// HardState will be equal to empty state if there is no update.
	pb.HardState

	// ReadStates can be used for localNode to serve linearizable read requests locally
	// when its applied index is greater than the index in ReadState.
	// Note that the readState will be returned when raft receives msgReadIndex.
	// The returned is only valid for the request that requested to read.
	ReadStates []ReadState

	// Entries specifies entries to be saved to stable storage BEFORE
	// Messages are sent.
	Entries []pb.Entry

	// Snapshot specifies the snapshot to be saved to stable storage.
	Snapshot pb.Snapshot

	// CommittedEntries specifies entries to be committed to a
	// store/state-machine. These have previously been committed to stable
	// store.
	CommittedEntries []pb.Entry

	// Messages 日志被提交到稳定的存储.如果它包含一个MsgSnap消息,应用程序必须在收到快照或调用ReportSnapshot失败时向raft报告.
	Messages []pb.Message // 就是raft.msgs

	// MustSync indicates whether the HardState and Entries必须是synchronously
	// written to disk or if an asynchronous write is permissible.
	MustSync bool
}

func isHardStateEqual(a, b pb.HardState) bool {
	return a.Term == b.Term && a.Vote == b.Vote && a.Commit == b.Commit
}

// IsEmptyHardState 判断是不是空的
func IsEmptyHardState(st pb.HardState) bool {
	return isHardStateEqual(st, emptyState)
}

// IsEmptySnap returns true if the given Snapshot is empty.
func IsEmptySnap(sp pb.Snapshot) bool {
	return sp.Metadata.Index == 0
}

func (rd Ready) containsUpdates() bool {
	return rd.SoftState != nil || !IsEmptyHardState(rd.HardState) ||
		!IsEmptySnap(rd.Snapshot) || len(rd.Entries) > 0 ||
		len(rd.CommittedEntries) > 0 || len(rd.Messages) > 0 || len(rd.ReadStates) != 0
}

// appliedCursor extracts from the Ready the highest index the client has
// applied (once the Ready is confirmed via Advance). If no information is
// contained in the Ready, returns zero.
func (rd Ready) appliedCursor() uint64 {
	if n := len(rd.CommittedEntries); n > 0 {
		return rd.CommittedEntries[n-1].Index
	}
	if index := rd.Snapshot.Metadata.Index; index > 0 {
		return index
	}
	return 0
}

// Node raft 节点
type Node interface {
	Tick()                              // 触发一次Tick,会触发Node心跳或者选举
	Campaign(ctx context.Context) error // 触发一次选举
	// Propose 提交日志条目
	Propose(ctx context.Context, data []byte) error
	// ProposeConfChange 集群配置变更
	ProposeConfChange(ctx context.Context, cc pb.ConfChangeI) error
	Step(ctx context.Context, msg pb.Message) error // 处理msg
	// Ready 如果raft状态机有变化,会通过channel返回一个Ready的数据结构,里面包含变化信息,比如日志变化、心跳发送等.
	// 调用方在处理完后需要调用Advance()方法告诉状态机上一个Ready处理完了
	Ready() <-chan Ready
	Advance()
	// ApplyConfChange 应用集群变化到状态机
	ApplyConfChange(cc pb.ConfChangeI) *pb.ConfState
	// TransferLeadership 将Leader转给transferee.
	TransferLeadership(ctx context.Context, lead, transferee uint64)
	// ReadIndex 请求一次线性读
	ReadIndex(ctx context.Context, rctx []byte) error
	// Status raft state machine当前状态.
	Status() Status
	// ReportUnreachable 告诉状态机指定id节点不可达.
	ReportUnreachable(id uint64)
	// ReportSnapshot 告诉状态机给id节点发送snapshot的最终处理状态.
	ReportSnapshot(id uint64, status SnapshotStatus)
	// Stop 关闭节点.
	Stop()
}

type Peer struct {
	ID      uint64 // 成员ID
	Context []byte // 成员信息序列化后的数据
}

// StartNode  它为每个给定的peer在初始日志中添加一个ConfChangeAddNode条目.
// Peer封装了节点的ID, peers记录了当前集群中全部节点的ID
func StartNode(c *Config, peers []Peer) Node { // ✅✈️ 🚗🚴🏻😁
	if len(peers) == 0 {
		panic("没有给定peers；使用RestartNode代替.")
	}
	rn, err := NewRawNode(c) // ✅
	if err != nil {
		panic(err)
	}
	rn.Bootstrap(peers) // [{"id":10276657743932975437,"peerURLs":["http://localhost:2380"],"name":"default"}]

	n := newLocalNode(rn) // 本机,用于接收发消息
	go n.run()            // ok

	return &n
}

// RestartNode   集群的当前成员将从Storage中恢复.
// 如果调用者有一个现有的状态机,请传入最后应用于它的日志索引；否则使用0.
func RestartNode(c *Config) Node {
	rn, err := NewRawNode(c)
	if err != nil {
		panic(err)
	}
	n := newLocalNode(rn)
	go n.run()
	return &n
}

type msgWithResult struct {
	m      pb.Message
	result chan error
}

//包含在raftNode中,是Node接口的实现.里面包含一个协程和多个队列,是状态机消息处理的入口.
type localNode struct {
	rn *RawNode
	// Propose队列,调用raftNode的Propose即把Propose消息塞到这个队列里
	propc chan msgWithResult
	// Message队列,除Propose消息以外其他消息塞到这个队列里
	recvc chan pb.Message
	// 接受配置的管道
	confc      chan pb.ConfChangeV2
	confstatec chan pb.ConfState
	// 已经准备好apply的信息队列
	readyc chan Ready
	// 每次apply好了以后往这个队列里塞个空对象.通知可以继续准备Ready消息.
	advancec chan struct{}
	//tick信息队列,用于调用心跳
	tickc chan struct{}
	// 在处理中避免不了各种chan操作,此时如果Stop()被调用了,相应的阻塞就应该被激活,否则可能
	// 面临死锁以后长时间退出后者永远无法退出.
	done chan struct{}
	// 为Stop接口实现的,应该还好理解
	stop chan struct{}
	// 一看就是为实现Status()用的,但是chan chan Status这个类型有点意思,后面分析实现函数
	// 看看如何实现的
	status chan chan Status
	// 用来写运行日志的
	logger Logger
}

// ok
func newLocalNode(rn *RawNode) localNode {
	return localNode{
		confc:      make(chan pb.ConfChangeV2), // 接收EntryConfChange类型消息比如动态添加节点
		rn:         rn,
		propc:      make(chan msgWithResult), // 接收网络层MsgProp类型消息
		recvc:      make(chan pb.Message),    // 接收网络层除MsgProp类型以外的消息
		confstatec: make(chan pb.ConfState),
		readyc:     make(chan Ready),         // 向上层返回 ready
		advancec:   make(chan struct{}),      // 上层处理往ready后返回给raft的消息
		tickc:      make(chan struct{}, 128), // 管理超时的管道,繁忙时可以处理之前的事件
		done:       make(chan struct{}),
		stop:       make(chan struct{}),
		status:     make(chan chan Status),
	}
}

func (n *localNode) Stop() {
	select {
	case n.stop <- struct{}{}:
		// Not already stopped, so trigger it
	case <-n.done:
		// Node has already been stopped - no need to do anything
		return
	}
	// Block until the stop has been acknowledged by run()
	<-n.done
}

func (n *localNode) run() {
	var propc chan msgWithResult
	var readyc chan Ready
	var advancec chan struct{}
	var rd Ready

	r := n.rn.raft
	// 初始状态不知道谁是leader,需要通过Ready获取
	lead := None
	for {
		if advancec != nil { // 开始时是nil
			readyc = nil
		} else if n.rn.HasReady() { //判断是否有Ready数据:待发送的数据
			rd = n.rn.readyWithoutAccept() // 获取Ready数据
			readyc = n.readyc              // 下边有放入数据的
		}

		if lead != r.lead {
			if r.hasLeader() {
				if lead == None {
					r.logger.Infof("raft.localNode: %x elected leader %x at term %d", r.id, r.lead, r.Term)
				} else {
					r.logger.Infof("raft.localNode: %x changed leader from %x to %x at term %d", r.id, lead, r.lead, r.Term)
				}
				propc = n.propc
			} else {
				r.logger.Infof("raft.localNode: %x lost leader %x at term %d", r.id, lead, r.Term)
				propc = nil
			}
			lead = r.lead
		}

		select {
		// TODO: maybe buffer the config propose if there exists one (the way
		// described in raft dissertation)
		// Currently it is dropped in Step silently.

		case pm := <-propc: //接收到写消息
			m := pm.m
			m.From = r.id
			err := r.Step(m)
			if pm.result != nil {
				pm.result <- err
				close(pm.result)
			}
		case m := <-n.recvc: // Message队列,除Propose消息以外其他消息塞到这个队列里
			// 必须是已知节点、或者是非响应类信息
			if pr := r.prs.Progress[m.From]; pr != nil || !IsResponseMsg(m.Type) {
				r.Step(m)
			}
		case cc := <-n.confc: //配置变更
			_, okBefore := r.prs.Progress[r.id]
			cs := r.applyConfChange(cc)
			// If the localNode was removed, block incoming proposals. Note that we
			// only do this if the localNode was in the config before. Nodes may be
			// a member of the group without knowing this (when they're catching
			// up on the log and don't have the latest config) and we don't want
			// to block the proposal channel in that case.
			//
			// NB: propc is reset when the leader changes, which, if we learn
			// about it, sort of implies that we got readded, maybe? This isn't
			// very sound and likely has bugs.
			if _, okAfter := r.prs.Progress[r.id]; okBefore && !okAfter {
				var found bool
			outer:
				for _, sl := range [][]uint64{cs.Voters, cs.VotersOutgoing} {
					for _, id := range sl {
						if id == r.id {
							found = true
							break outer
						}
					}
				}
				if !found {
					propc = nil
				}
			}
			select {
			case n.confstatec <- cs:
			case <-n.done:
			}
		case <-n.tickc: //超时时间到,包括心跳超时和选举超时等
			//https://www.cnblogs.com/myd620/p/13189604.html
			n.rn.Tick()
		case readyc <- rd: // 数据放入ready channel中
			n.rn.acceptReady(rd)  // 告诉raft,ready数据已被接收
			advancec = n.advancec // 赋值Advance channel等待Ready处理完成的消息
		case <-advancec: //可以进行状态变更和日志提交
			n.rn.Advance(rd)
			rd = Ready{}
			advancec = nil
		case c := <-n.status: //节点状态信号
			c <- getStatus(r)
		case <-n.stop: //收到停止信号
			close(n.done)
			return
		}
	}
}

// Tick increments the internal logical clock for this Node. Election timeouts
// and heartbeat timeouts are in units of ticks.
func (n *localNode) Tick() {
	select {
	case n.tickc <- struct{}{}:
	case <-n.done:
	default:
		n.rn.raft.logger.Warningf("%x A tick missed to fire. Node blocks too long!", n.rn.raft.id)
	}
}

// Campaign 选举
func (n *localNode) Campaign(ctx context.Context) error {
	return n.step(ctx, pb.Message{Type: pb.MsgHup})
}

// Propose 提议
func (n *localNode) Propose(ctx context.Context, data []byte) error {
	return n.stepWait(ctx, pb.Message{Type: pb.MsgProp, Entries: []pb.Entry{{Data: data}}})
}

// Step 步骤
func (n *localNode) Step(ctx context.Context, m pb.Message) error {
	// 忽略通过网络接收的非本地信息
	if IsLocalMsg(m.Type) {
		return nil
	}
	return n.step(ctx, m)
}

// OK
func (n *localNode) step(ctx context.Context, m pb.Message) error {
	return n.stepWithWaitOption(ctx, m, false)
}

func (n *localNode) stepWait(ctx context.Context, m pb.Message) error {
	return n.stepWithWaitOption(ctx, m, true)
}

// Step 使用msgs推进状态机.如果有的话,ctx.Err()将被返回.
func (n *localNode) stepWithWaitOption(ctx context.Context, m pb.Message, wait bool) error {
	if m.Type != pb.MsgProp { // pb.MsgProp  本地：Propose -----> MsgApp
		select {
		case n.recvc <- m:
			return nil // 一般都会走这里
		case <-ctx.Done():
			return ctx.Err()
		case <-n.done:
			return ErrStopped
		}
	}
	ch := n.propc
	pm := msgWithResult{m: m}
	if wait {
		pm.result = make(chan error, 1)
	}
	select {
	case ch <- pm:
		if !wait {
			return nil
		}
	case <-ctx.Done():
		return ctx.Err()
	case <-n.done:
		return ErrStopped
	}
	select {
	case err := <-pm.result:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	case <-n.done:
		return ErrStopped
	}
	return nil
}

func confChangeToMsg(c pb.ConfChangeI) (pb.Message, error) {
	typ, data, err := pb.MarshalConfChange(c)
	if err != nil {
		return pb.Message{}, err
	}
	return pb.Message{Type: pb.MsgProp, Entries: []pb.Entry{{Type: typ, Data: data}}}, nil
}

func (n *localNode) ProposeConfChange(ctx context.Context, cc pb.ConfChangeI) error {
	msg, err := confChangeToMsg(cc)
	if err != nil {
		return err
	}
	return n.Step(ctx, msg)
}

// Ready 如果raft状态机有变化,会通过channel返回一个Ready的数据结构,里面包含变化信息,比如日志变化、心跳发送等.
func (n *localNode) Ready() <-chan Ready { return n.readyc }

// Advance ready消息处理完后,发送一个通知消息
func (n *localNode) Advance() {
	select {
	case n.advancec <- struct{}{}:
	case <-n.done:
	}
}

func (n *localNode) ApplyConfChange(cc pb.ConfChangeI) *pb.ConfState {
	var cs pb.ConfState
	select {
	case n.confc <- cc.AsV2():
	case <-n.done:
	}
	select {
	case cs = <-n.confstatec:
	case <-n.done:
	}
	return &cs
}

func (n *localNode) Status() Status {
	c := make(chan Status)
	select {
	case n.status <- c:
		return <-c
	case <-n.done:
		return Status{}
	}
}

func (n *localNode) ReportUnreachable(id uint64) {
	select {
	case n.recvc <- pb.Message{Type: pb.MsgUnreachable, From: id}:
	case <-n.done:
	}
}

func (n *localNode) ReportSnapshot(id uint64, status SnapshotStatus) {
	rej := status == SnapshotFailure

	select {
	case n.recvc <- pb.Message{Type: pb.MsgSnapStatus, From: id, Reject: rej}:
	case <-n.done:
	}
}

func (n *localNode) TransferLeadership(ctx context.Context, lead, transferee uint64) {
	select {
	// manually set 'from' and 'to', so that leader can voluntarily transfers its leadership
	case n.recvc <- pb.Message{Type: pb.MsgTransferLeader, From: transferee, To: lead}:
	case <-n.done:
	case <-ctx.Done():
	}
}

func (n *localNode) ReadIndex(ctx context.Context, rctx []byte) error {
	return n.step(ctx, pb.Message{Type: pb.MsgReadIndex, Entries: []pb.Entry{{Data: rctx}}})
}

func newReady(r *raft, prevSoftSt *SoftState, prevHardSt pb.HardState) Ready {
	rd := Ready{
		Entries:          r.raftLog.unstableEntries(), // unstable中的日志交给上层持久化
		CommittedEntries: r.raftLog.nextEnts(),        // 已经提交待应用的日志,交给上层应用
		Messages:         r.msgs,                      // raft要发送的消息   ,为了之后读
	}
	//判断softState有没有变化,有则赋值
	if softSt := r.softState(); !softSt.equal(prevSoftSt) {
		rd.SoftState = softSt
	}
	//判断hardState有没有变化,有则赋值
	if hardSt := r.hardState(); !isHardStateEqual(hardSt, prevHardSt) {
		rd.HardState = hardSt
	}
	//判断是不是收到snapshot
	if r.raftLog.unstable.snapshot != nil {
		rd.Snapshot = *r.raftLog.unstable.snapshot
	}
	if len(r.readStates) != 0 {
		rd.ReadStates = r.readStates
	}
	//处理该Ready后是否需要做fsync,将数据强制刷盘
	rd.MustSync = MustSync(r.hardState(), prevHardSt, len(rd.Entries))
	return rd
}

// MustSync returns true if the hard state and count of Raft entries indicate
// that a synchronous write to persistent storage is required.
func MustSync(st, prevst pb.HardState, entsnum int) bool {
	// Persistent state on all servers:
	// (Updated on stable storage before responding to RPCs)
	// currentTerm
	// votedFor
	// log entries[]
	return entsnum != 0 || st.Vote != prevst.Vote || st.Term != prevst.Term
}
