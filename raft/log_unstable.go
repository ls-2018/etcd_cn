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

import pb "github.com/ls-2018/etcd_cn/raft/raftpb"

// 使用内存数组维护其中所有的Entry记录，对于Leader节点而言，它维护了客户端请求对应的Entry记录；
// 对于Follower节点而言，它维护的是从Leader节点复制来的Entry记录。
// 无论是Leader节点还是Follower节点，对于刚刚接收到的Entry记录首先都会被存储在unstable中。
// 然后按照Raft协议将unstable中缓存的这些Entry记录交给上层模块进行处理，上层模块会将这些Entry记录发送到集群其他节点或进行保存（写入Storage中）。
// 之后，上层模块会调用Advance（）方法通知底层的raft模块将unstable 中对应的Entry记录删除（因为己经保存到了Storage中）
type unstable struct {
	snapshot *pb.Snapshot // 快照数据，该快照数据也是未写入Storage中的。
	entries  []pb.Entry   // 用于保存未写入Storage中的Entry记录。
	offset   uint64       // 当前entries中第一个日志的索引,起始索引
	logger   Logger
}

// maybeFirstIndex returns the index of the first possible entry in entries
// if it has a snapshot.
// 会尝试获取unstable 的第一条Entry 记录的索引值
func (u *unstable) maybeFirstIndex() (uint64, bool) {
	if u.snapshot != nil {
		return u.snapshot.Metadata.Index + 1, true
	}
	return 0, false
}

// maybeLastIndex returns the last index if it has at least one
// unstable entry or snapshot.
// 尝试获取unstable 的最后一条Entry记录的索引值
func (u *unstable) maybeLastIndex() (uint64, bool) {
	if l := len(u.entries); l != 0 {
		return u.offset + uint64(l) - 1, true
	}
	if u.snapshot != nil {
		return u.snapshot.Metadata.Index, true
	}
	return 0, false
}

// maybeTerm 尝试获取指定Entry记录的Term值，根据条件查找指定的Entry记录的位置。
func (u *unstable) maybeTerm(i uint64) (uint64, bool) {
	// 打完快照之后,之前日志的数据就不保存了，包括任期、索引等等
	if i < u.offset {
		if u.snapshot != nil && u.snapshot.Metadata.Index == i {
			return u.snapshot.Metadata.Term, true
		}
		return 0, false
	}

	last, ok := u.maybeLastIndex()
	if !ok {
		return 0, false
	}
	if i > last {
		return 0, false
	}

	return u.entries[i-u.offset].Term, true
}

//  当unstable.entries 中的Entry记录己经被写入Storage之后，会调用unstable.stableTo（）方法清除entries 中对应的Entry记录
func (u *unstable) stableTo(i, t uint64) {
	// i: 已经持久化的日志索引; t 当前任期
	// 查找指定Entry记录的Term佳，若查找失败则表示对应的Entry不在unstable中，直接返回
	gt, ok := u.maybeTerm(i)
	if !ok {
		return
	}
	if gt == t && i >= u.offset {
		// 指定索引位之前的Entry记录都已经完成持久化，则将其之前的全部Entry记录删除
		u.entries = u.entries[i+1-u.offset:]
		u.offset = i + 1
		// 随着多次追加日志和截断日志的操作unstable.entires底层的数组会越来越大，
		// shrinkEntriesArray方法会在底层数组长度超过实际占用的两倍时，对底层数据进行缩减
		u.shrinkEntriesArray()
	}
}

// shrinkEntriesArray 释放数组无用空间
func (u *unstable) shrinkEntriesArray() {
	const lenMultiple = 2
	if len(u.entries) == 0 {
		u.entries = nil
	} else if len(u.entries)*lenMultiple < cap(u.entries) {
		// 重新创建切片,复制原有切片中的数据,重直entries字段
		newEntries := make([]pb.Entry, len(u.entries))
		copy(newEntries, u.entries)
		u.entries = newEntries
	}
}

func (u *unstable) stableSnapTo(i uint64) {
	if u.snapshot != nil && u.snapshot.Metadata.Index == i {
		u.snapshot = nil
	}
}

func (u *unstable) restore(s pb.Snapshot) {
	u.offset = s.Metadata.Index + 1
	u.entries = nil
	u.snapshot = &s
}

// 截断和追加
func (u *unstable) truncateAndAppend(ents []pb.Entry) {
	after := ents[0].Index
	switch {
	case after == u.offset+uint64(len(u.entries)):
		//若待追加的记录与e口tries中的记录正好连续，则可以直接向entries中追加
		u.entries = append(u.entries, ents...)
	case after <= u.offset:
		u.logger.Infof("直接用待追加的Entry记录替换当前的entries字段，并支新offset %d", after)
		u.offset = after
		u.entries = ents
	default:
		//after在offset～last之间，则after～last之间的Entry记录冲突。 这里会将offset~after 之间的记录保留，抛弃after之后的记录，然后完成追加操作
		//unstable.slice()方法会检测after是否合法，并返回offset~after的切片
		u.logger.Infof("截断在after之后数据 %d", after)
		u.entries = append([]pb.Entry{}, u.slice(u.offset, after)...)
		u.entries = append(u.entries, ents...)
	}
}

// 返回切片
func (u *unstable) slice(lo uint64, hi uint64) []pb.Entry {
	u.mustCheckOutOfBounds(lo, hi)
	return u.entries[lo-u.offset : hi-u.offset]
}

// 范围检查 u.offset <= lo <= hi <= u.offset+len(u.entries)
func (u *unstable) mustCheckOutOfBounds(lo, hi uint64) {
	if lo > hi {
		u.logger.Panicf("无效的切片 %d > %d", lo, hi)
	}
	upper := u.offset + uint64(len(u.entries))
	if lo < u.offset || hi > upper {
		u.logger.Panicf("unstable.slice[%d,%d) 超出范围 [%d,%d]", lo, hi, u.offset, upper)
	}
}
