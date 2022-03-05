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

// Ready encapsulates the entries and messages that are ready to read,
// be saved to stable storage, committed or sent to other peers.
// All fields in Ready are read-only.
type Ready struct {
	// The current volatile state of a Node.
	// SoftState will be nil if there is no update.
	// It is not required to consume or store SoftState.
	*SoftState

	// The current state of a Node to be saved to stable storage BEFORE
	// Messages are sent.
	// HardState will be equal to empty state if there is no update.
	pb.HardState

	// ReadStates can be used for node to serve linearizable read requests locally
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

	// Messages specifies outbound messages to be sent AFTER Entries are
	// committed to stable storage.
	// If it contains a MsgSnap message, the application MUST report back to raft
	// when the snapshot has been received or has failed by calling ReportSnapshot.
	Messages []pb.Message

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

// Node represents a node in a raft cluster.
type Node interface {
	// Tick 触发一次心跳,raft会在触发后检查leader选举超时或发送心跳
	Tick()
	// Campaign 触发节点将自己变成候选人,开始选举
	Campaign(ctx context.Context) error
	// Propose 提交日志条目
	Propose(ctx context.Context, data []byte) error
	// ProposeConfChange 集群配置变更
	ProposeConfChange(ctx context.Context, cc pb.ConfChangeI) error
	// Step 发送一条消息给状态机,触发状态变化
	Step(ctx context.Context, msg pb.Message) error
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
func StartNode(c *Config, peers []Peer) Node {
	if len(peers) == 0 {
		panic("没有给定peers；使用RestartNode代替.")
	}
	rn, err := NewRawNode(c) // ✅
	if err != nil {
		panic(err)
	}
	rn.Bootstrap(peers) // [{"id":10276657743932975437,"peerURLs":["http://localhost:2380"],"name":"default"}]

	n := newNode(rn)

	go n.run() // ok
	return &n
}

// RestartNode   集群的当前成员将从Storage中恢复.
// 如果调用者有一个现有的状态机,请传入最后应用于它的日志索引；否则使用0.
func RestartNode(c *Config) Node {
	rn, err := NewRawNode(c)
	if err != nil {
		panic(err)
	}
	n := newNode(rn)
	go n.run()
	return &n
}

type msgWithResult struct {
	m      pb.Message
	result chan error
}

type node struct {
	rn *RawNode
	// 用于实现Propose()接口
	propc chan msgWithResult
	// 用于实现Step()接口
	recvc chan pb.Message
	// 这两个chan用于实现ApplyConfChange()接口
	confc      chan pb.ConfChangeV2
	confstatec chan pb.ConfState
	// 用于实现Ready()接口
	readyc chan Ready
	// 用于实现Advance()接口
	advancec chan struct{}
	// 用于实现Tick()接口,这个需要注意一下,创建node时tickc是有缓冲的,设计者的解释是当node
	// 忙的时候可能一个操作会超过tick的周期,这样会使得计时不准,有了缓冲就可以避免这个问题.
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
func newNode(rn *RawNode) node {
	return node{
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

func (n *node) Stop() {
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

func (n *node) run() {
	var propc chan msgWithResult
	var readyc chan Ready
	var advancec chan struct{}
	var rd Ready

	r := n.rn.raft
	// 初始状态不知道谁是leader,需要通过Ready获取
	lead := None
	for {
		if advancec != nil {
			readyc = nil
		} else if n.rn.HasReady() {
			rd = n.rn.readyWithoutAccept()
			readyc = n.readyc
		}

		if lead != r.lead {
			if r.hasLeader() {
				if lead == None {
					r.logger.Infof("raft.node: %x elected leader %x at term %d", r.id, r.lead, r.Term)
				} else {
					r.logger.Infof("raft.node: %x changed leader from %x to %x at term %d", r.id, lead, r.lead, r.Term)
				}
				propc = n.propc
			} else {
				r.logger.Infof("raft.node: %x lost leader %x at term %d", r.id, lead, r.Term)
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
		case m := <-n.recvc: //接收到readindex 请求
			// filter out response message from unknown From.
			if pr := r.prs.Progress[m.From]; pr != nil || !IsResponseMsg(m.Type) {
				r.Step(m)
			}
		case cc := <-n.confc: //配置变更
			_, okBefore := r.prs.Progress[r.id]
			cs := r.applyConfChange(cc)
			// If the node was removed, block incoming proposals. Note that we
			// only do this if the node was in the config before. Nodes may be
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
		case readyc <- rd: //数据ready
			n.rn.acceptReady(rd)
			advancec = n.advancec
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
func (n *node) Tick() {
	select {
	case n.tickc <- struct{}{}:
	case <-n.done:
	default:
		n.rn.raft.logger.Warningf("%x A tick missed to fire. Node blocks too long!", n.rn.raft.id)
	}
}

func (n *node) Campaign(ctx context.Context) error { return n.step(ctx, pb.Message{Type: pb.MsgHup}) }

func (n *node) Propose(ctx context.Context, data []byte) error {
	return n.stepWait(ctx, pb.Message{Type: pb.MsgProp, Entries: []pb.Entry{{Data: data}}})
}

func (n *node) Step(ctx context.Context, m pb.Message) error {
	// ignore unexpected local messages receiving over network
	if IsLocalMsg(m.Type) {
		// TODO: return an error?
		return nil
	}
	return n.step(ctx, m)
}

func confChangeToMsg(c pb.ConfChangeI) (pb.Message, error) {
	typ, data, err := pb.MarshalConfChange(c)
	if err != nil {
		return pb.Message{}, err
	}
	return pb.Message{Type: pb.MsgProp, Entries: []pb.Entry{{Type: typ, Data: data}}}, nil
}

func (n *node) ProposeConfChange(ctx context.Context, cc pb.ConfChangeI) error {
	msg, err := confChangeToMsg(cc)
	if err != nil {
		return err
	}
	return n.Step(ctx, msg)
}

func (n *node) step(ctx context.Context, m pb.Message) error {
	return n.stepWithWaitOption(ctx, m, false)
}

func (n *node) stepWait(ctx context.Context, m pb.Message) error {
	return n.stepWithWaitOption(ctx, m, true)
}

// Step advances the state machine using msgs. The ctx.Err() will be returned,
// if any.
func (n *node) stepWithWaitOption(ctx context.Context, m pb.Message, wait bool) error {
	if m.Type != pb.MsgProp {
		select {
		case n.recvc <- m:
			return nil
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

func (n *node) Ready() <-chan Ready { return n.readyc }

func (n *node) Advance() {
	select {
	case n.advancec <- struct{}{}:
	case <-n.done:
	}
}

func (n *node) ApplyConfChange(cc pb.ConfChangeI) *pb.ConfState {
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

func (n *node) Status() Status {
	c := make(chan Status)
	select {
	case n.status <- c:
		return <-c
	case <-n.done:
		return Status{}
	}
}

func (n *node) ReportUnreachable(id uint64) {
	select {
	case n.recvc <- pb.Message{Type: pb.MsgUnreachable, From: id}:
	case <-n.done:
	}
}

func (n *node) ReportSnapshot(id uint64, status SnapshotStatus) {
	rej := status == SnapshotFailure

	select {
	case n.recvc <- pb.Message{Type: pb.MsgSnapStatus, From: id, Reject: rej}:
	case <-n.done:
	}
}

func (n *node) TransferLeadership(ctx context.Context, lead, transferee uint64) {
	select {
	// manually set 'from' and 'to', so that leader can voluntarily transfers its leadership
	case n.recvc <- pb.Message{Type: pb.MsgTransferLeader, From: transferee, To: lead}:
	case <-n.done:
	case <-ctx.Done():
	}
}

func (n *node) ReadIndex(ctx context.Context, rctx []byte) error {
	return n.step(ctx, pb.Message{Type: pb.MsgReadIndex, Entries: []pb.Entry{{Data: rctx}}})
}

func newReady(r *raft, prevSoftSt *SoftState, prevHardSt pb.HardState) Ready {
	rd := Ready{
		Entries:          r.raftLog.unstableEntries(), //unstable中的日志交给上层持久化
		CommittedEntries: r.raftLog.nextEnts(),        //已经提交待应用的日志,交给上层应用
		Messages:         r.msgs,                      //raft要发送的消息
	}
	if softSt := r.softState(); !softSt.equal(prevSoftSt) {
		rd.SoftState = softSt
	}
	if hardSt := r.hardState(); !isHardStateEqual(hardSt, prevHardSt) {
		rd.HardState = hardSt
	}
	if r.raftLog.unstable.snapshot != nil {
		rd.Snapshot = *r.raftLog.unstable.snapshot
	}
	if len(r.readStates) != 0 {
		rd.ReadStates = r.readStates
	}
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
