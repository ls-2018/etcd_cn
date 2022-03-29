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

package mvcc

import (
	"context"

	"github.com/ls-2018/etcd_cn/etcd_backend/lease"
	"github.com/ls-2018/etcd_cn/offical/api/v3/mvccpb"
)

// RangeOptions 请求参数
type RangeOptions struct {
	Limit int64 // 用户限制的数据量
	Rev   int64 // 指定的修订版本
	Count bool  // 是否统计修订版本数
}

// RangeResult 响应
type RangeResult struct {
	KVs   []mvccpb.KeyValue
	Rev   int64 // 最新的修订版本
	Count int   // 统计当前的   修订版本数
}

type ReadView interface {
	// FirstRev
	//     before    			   cur
	//             compact
	//           rev       rev   		rev
	FirstRev() int64 // 在打开txn时返回第一个KV修订。在压实之后，第一个修订增加到压实修订。
	Rev() int64      // 在打开txn时返回KV的修订。
	Range(ctx context.Context, key, end []byte, ro RangeOptions) (r *RangeResult, err error)
}

// TxnRead 只读事务,不会锁住其他只读事务
type TxnRead interface {
	ReadView
	End() // 标记事务已完成 并且准备提交
}

type WriteView interface {
	// DeleteRange deletes the given range from the store.
	// A deleteRange increases the rev of the store if any key in the range exists.
	// The number of key deleted will be returned.
	// The returned rev is the current revision of the KV when the operation is executed.
	// It also generates one event for each key delete in the event history.
	// if the `end` is nil, deleteRange deletes the key.
	// if the `end` is not nil, deleteRange deletes the keys in range [key, range_end).
	DeleteRange(key, end []byte) (n, rev int64)

	// Put puts the given key, value into the store. Put also takes additional argument lease to
	// attach a lease to a key-value pair as meta-data. KV implementation does not validate the lease
	// id.
	// A put also increases the rev of the store, and generates one event in the event history.
	// The returned rev is the current revision of the KV when the operation is executed.
	Put(key, value []byte, lease lease.LeaseID) (rev int64)
}

type TxnWrite interface {
	TxnRead
	WriteView
	// Changes 获取打开write txn后所做的更改。
	Changes() []mvccpb.KeyValue
}

// txnReadWrite 读事务-->写事务，对任何写操作都感到恐慌。
type txnReadWrite struct {
	TxnRead
}

func (trw *txnReadWrite) DeleteRange(key, end []byte) (n, rev int64) { panic("unexpected DeleteRange") }
func (trw *txnReadWrite) Put(key, value []byte, lease lease.LeaseID) (rev int64) {
	panic("unexpected Put")
}
func (trw *txnReadWrite) Changes() []mvccpb.KeyValue { return nil }

func NewReadOnlyTxnWrite(txn TxnRead) TxnWrite { return &txnReadWrite{txn} }

type ReadTxMode uint32

const (
	ConcurrentReadTxMode = ReadTxMode(1) // 缓冲区拷贝，提高性能
	SharedBufReadTxMode  = ReadTxMode(2)
)
