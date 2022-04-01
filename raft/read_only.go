// Copyright 2016 The etcd Authors
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
	pb "github.com/ls-2018/etcd_cn/raft/raftpb"
	"github.com/ls-2018/etcd_cn/raft/tracker"
)

// ReadState provides state for read only query.
// It's caller's responsibility to call ReadIndex first before getting
// this state from ready, it's also caller's duty to differentiate if this
// state is what it requests through RequestCtx, eg. given a unique id as
// RequestCtx
// 这个参数就是Node.ReadIndex()的结果回调.
type ReadState struct {
	Index      uint64 // leader 节点已经committed的索引
	RequestCtx []byte // 递增ID
}

type readIndexStatus struct {
	req   pb.Message // 记录了对应的MsgReadIndex请求
	index uint64     // 该MsgReadIndex请求到达时,对应的已提交位置
	// NB: this never records 'false', but it's more convenient to use this
	// instead of a map[uint64]struct{} due to the API of quorum.VoteResult. If
	// this becomes performance sensitive enough (doubtful), quorum.VoteResult
	// can change to an API that is closer to that of CommittedIndex.
	acks map[uint64]bool // 记录了该MsgReadIndex相关的MsgHeartbeatResp响应的信息
}

type readOnly struct {
	option ReadOnlyOption // 当前只读请求的处理模式,ReadOnlySafe ReadOnlyOpt 和	ReadOnlyLeaseBased两种模式
	/*
		在etcd服务端收到MsgReadIndex消息时,会为其创建一个唯一的消息ID,并作为MsgReadIndex消息的第一条Entry记录.
		在pendingReadIndex维护了消息ID与对应请求readIndexStatus实例的映射
	*/
	pendingReadIndex map[string]*readIndexStatus // id -->readIndexStatus
	readIndexQueue   []string                    // 记录了MsgReadIndex请求对应的消息ID,这样可以保证MsgReadIndex的顺序
}

func newReadOnly(option ReadOnlyOption) *readOnly {
	return &readOnly{
		option:           option,
		pendingReadIndex: make(map[string]*readIndexStatus),
	}
}

// addRequest adds a read only request into readonly struct.
// `index` is the commit index of the raft state machine when it received
// the read only request.
// `m` is the original read only request message from the local or remote localNode.
//将已提交的位置(raftLog.committed)以及MsgReadIndex消息的相关信息存到readOnly中
/*
1.获取消息ID,在ReadIndex消息的第一个记录中记录了消息ID
2.判断该消息是否已经记录在pendingReadIndex中,如果已存在则直接返回
3.如果不存在,则维护到pendingReadIndex中,index是当前Leader已提交的位置,m是请求的消息
4.并将消息ID追加到readIndexQueue队列中
*/
func (ro *readOnly) addRequest(index uint64, m pb.Message) {
	s := string(m.Entries[0].Data)
	if _, ok := ro.pendingReadIndex[s]; ok {
		return
	}
	ro.pendingReadIndex[s] = &readIndexStatus{index: index, req: m, acks: make(map[uint64]bool)}
	ro.readIndexQueue = append(ro.readIndexQueue, s)
}

// recvAck notifies the readonly struct that the raft state machine received
// an acknowledgment of the heartbeat that attached with the read only request
// context.
/*
recvAck通知readonly结构,即raft状态机接受了对只读请求上下文附加的心跳的确认.
1.消息的Context即消息ID,根据消息id获取对应的readIndexStatus
2.如果获取不到则返回0
3.记录了该Follower节点返回的MsgHeartbeatResp响应的信息
4.返回Follower响应的数量
*/
func (ro *readOnly) recvAck(id uint64, context []byte) map[uint64]bool {
	rs, ok := ro.pendingReadIndex[string(context)]
	if !ok {
		return nil
	}

	rs.acks[id] = true
	return rs.acks
}

// advance advances the read only request queue kept by the readonly struct.
// It dequeues the requests until it finds the read only request that has
// the same context as the given `m`.
/*
1.遍历readIndexQueue队列,如果能找到该消息的Context,则返回该消息及之前的所有记录rss,
	并删除readIndexQueue队列和pendingReadIndex中对应的记录
2.如果没有Context对应的消息ID,则返回nil
*/
func (ro *readOnly) advance(m pb.Message) []*readIndexStatus {
	var (
		i     int
		found bool
	)

	ctx := string(m.Context)
	var rss []*readIndexStatus

	for _, okctx := range ro.readIndexQueue {
		i++
		rs, ok := ro.pendingReadIndex[okctx]
		if !ok {
			panic("cannot find corresponding read state from pending map")
		}
		rss = append(rss, rs)
		if okctx == ctx {
			found = true
			break
		}
	}

	if found {
		ro.readIndexQueue = ro.readIndexQueue[i:]
		for _, rs := range rss {
			delete(ro.pendingReadIndex, string(rs.req.Entries[0].Data))
		}
		return rss
	}

	return nil
}

// lastPendingRequestCtx returns the context of the last pending read only
// request in readonly struct.
// 返回记录中最后一个消息ID
func (ro *readOnly) lastPendingRequestCtx() string {
	if len(ro.readIndexQueue) == 0 {
		return ""
	}
	return ro.readIndexQueue[len(ro.readIndexQueue)-1]
}

// bcastHeartbeat 发送RPC,没有日志给所有对等体.
func (r *raft) bcastHeartbeat() {
	lastCtx := r.readOnly.lastPendingRequestCtx()
	if len(lastCtx) == 0 {
		r.bcastHeartbeatWithCtx(nil)
	} else {
		r.bcastHeartbeatWithCtx([]byte(lastCtx))
	}
}

func (r *raft) bcastHeartbeatWithCtx(ctx []byte) {
	r.prstrack.Visit(func(id uint64, _ *tracker.Progress) {
		if id == r.id {
			return
		}
		r.sendHeartbeat(id, ctx)
	})
}
