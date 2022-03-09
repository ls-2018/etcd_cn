// Copyright 2019 The etcd Authors
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

package tracker

import (
	"fmt"
	"sort"
	"strings"
)

// Progress 在leader看来,Progress代表follower的进度.leader维护所有follower的进度,并根据follower的进度向其发送条目.
// NB(tbg).Progress基本上是一个状态机
type Progress struct {
	Match uint64 // 对应Follower节点当前己经成功复制的Entry记录的索引值.
	Next  uint64 // 对应Follower节点下一个待复制的Entry记录的索引值
	// State 对应Follower节点的复制状态
	// 当处于StateProbe状态时,leader在每个心跳间隔内最多发送一条复制消息.它也会探测follower的实际进度.
	// 当处于StateReplicate状态时,leader在发送复制消息后,乐观地增加next 索引.这是一个优化后的的状态,用于快速复制日志条目给follower.
	// 当处于StateSnapshot状态时,leader应该已经发送了快照,并停止发送任何复制消息.
	State StateType

	PendingSnapshot uint64 // 表示Leader节点正在向目标节点发送快照数据.

	RecentActive bool // 从当前Leader节点的角度来看,该Progress实例对应的Follower节点是否存活.

	StopSent bool // 是否暂停对follower的消息发送

	Inflights *Inflights // 维护着向该follower已发送,但未收到确认的消息索引 [环形队列]

	IsLearner bool // 该节点是不是learner
}

// MaybeUpdate is called when an MsgAppResp arrives from the follower, with the
// index acked by it. The method returns false if the given n index comes from
// an outdated message. Otherwise it updates the progress and returns true.
func (pr *Progress) MaybeUpdate(n uint64) bool {
	var updated bool
	if pr.Match < n {
		pr.Match = n
		updated = true
		pr.ProbeAcked()
	}
	pr.Next = max(pr.Next, n+1)
	return updated
}

// OptimisticUpdate 记录下次日志发送的起始位置,n是已发送的最新日志索引
func (pr *Progress) OptimisticUpdate(n uint64) { pr.Next = n + 1 }

// MaybeDecrTo  follower拒绝了rejected索引的日志, matchHint是应该重新调整的日志索引[leader记录的];
// leader是否降低对该节点索引记录
func (pr *Progress) MaybeDecrTo(rejected, matchHint uint64) bool {
	if pr.State == StateReplicate {
		if rejected <= pr.Match {
			return false
		}
		// Directly decrease next to match + 1.
		//
		// TODO(tbg): why not use matchHint if it's larger?
		//根据前面对MsgApp消息发送过程的分析，处于ProgressStateReplicate状态时，发送MsgApp
		//消息的同时会直接调用Progress.optimisticUpdate（）方法增加Next，这就使得Next可能会
		//比Match大很多，这里回退Next至Match位置，并在后面重新发送MsgApp消息进行尝试
		pr.Next = pr.Match + 1
		return true
	}
	// The rejection必须是stale if "rejected" does not match next - 1. This
	// is because non-replicating followers are probed one entry at a time.
	//出现过时的MsgAppResp消息直接忽略
	// 只有拒绝比当前大,才会重新      一般都是当前记录了A ,给follower发送了A+1,[leader会将Next设为A+2] ,被拒绝了   此时 rejected=A+1
	if pr.Next-1 != rejected {
		return false
	}

	//   idx        1 2 3 4 5 6 7 8 9
	//              -----------------
	//   term (L)   1 3 3 3 5 5 5 5 5
	//   term (F)   1 1 1 1 2 2
	pr.Next = max(min(rejected, matchHint+1), 1) // 此时Next就设为2
	pr.StopSent = false                             // Next重置完成，恢复消息发送，并在后面重新发送MsgApp消息
	return true
}

// -------------------------------------------------- over --------------------------------------------------

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}

// ProbeAcked 当follower接受了append消息,标志着可以继续向该节点发送消息
func (pr *Progress) ProbeAcked() {
	pr.StopSent = false
}

// IsPaused 返回发往该节点的消息是否被限流
// 当一个节点拒绝了最近的MsgApps,目前正在等待快照,或者已经达到MaxInflightMsgs限制时,就会这样做.
// 在正常操作中,这是假的.一个被节流的节点将被减少联系的频率,直到它达到能够再次接受稳定的日志条目的状态.
func (pr *Progress) IsPaused() bool {
	switch pr.State {
	case StateProbe: // 每个心跳间隔内最多发送一条复制消息,默认false
		return pr.StopSent
	case StateReplicate: // 消息复制状态
		return pr.Inflights.Full() // 根据队列是否满,判断
	case StateSnapshot:
		return true // follower接收快照时,停止发送消息
	default:
		panic("未知的状态")
	}
}

// ResetState 重置节点的跟踪状态
func (pr *Progress) ResetState(state StateType) {
	pr.StopSent = false
	pr.PendingSnapshot = 0
	pr.State = state
	pr.Inflights.reset()
}

// BecomeProbe 转变为StateProbe.下一步是重置为Match+1,或者,如果更大的话,重置为待定快照的索引.
// 恢复follower状态,以正常发送消息
func (pr *Progress) BecomeProbe() {
	// 变成探测状态,等发出去的消息响应了，再继续发消息
	if pr.State == StateSnapshot { // 当前状态是发送快照
		pendingSnapshot := pr.PendingSnapshot
		pr.Next = max(pr.Match+1, pendingSnapshot+1)
	} else { // follower 链接有问题 、网络有问题
		pr.Next = pr.Match + 1
	}
	pr.ResetState(StateProbe)
}

func (pr *Progress) BecomeReplicate() {
	pr.ResetState(StateReplicate) // 消息可以复制状态
	pr.Next = pr.Match + 1
}

// BecomeSnapshot 正在发送快照 ,snapshoti 为快照的最新日志索引
func (pr *Progress) BecomeSnapshot(snapshoti uint64) {
	pr.ResetState(StateSnapshot)
	pr.PendingSnapshot = snapshoti
}

func (pr *Progress) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s match=%d next=%d", pr.State, pr.Match, pr.Next)
	if pr.IsLearner {
		fmt.Fprint(&buf, " learner")
	}
	if pr.IsPaused() {
		fmt.Fprint(&buf, " paused")
	}
	if pr.PendingSnapshot > 0 {
		fmt.Fprintf(&buf, " 发送中的快照=%d", pr.PendingSnapshot)
	}
	if !pr.RecentActive {
		fmt.Fprintf(&buf, " 不活跃的")
	}
	if n := pr.Inflights.Count(); n > 0 {
		fmt.Fprintf(&buf, " 未确认的消息=%d", n)
		if pr.Inflights.Full() {
			fmt.Fprint(&buf, "[full]")
		}
	}
	return buf.String()
}

type ProgressMap map[uint64]*Progress

func (m ProgressMap) String() string {
	ids := make([]uint64, 0, len(m))
	for k := range m {
		ids = append(ids, k)
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	var buf strings.Builder
	for _, id := range ids {
		fmt.Fprintf(&buf, "%d: %s\n", id, m[id])
	}
	return buf.String()
}
