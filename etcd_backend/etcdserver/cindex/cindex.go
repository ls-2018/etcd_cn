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

package cindex

import (
	"encoding/binary"
	"sync"
	"sync/atomic"

	"github.com/ls-2018/etcd_cn/etcd_backend/mvcc/backend"
	"github.com/ls-2018/etcd_cn/etcd_backend/mvcc/buckets"
)

type Backend interface {
	BatchTx() backend.BatchTx
}

// ConsistentIndexer 是一个接口，它封装了一致性索引的Get/Set/Save方法。
type ConsistentIndexer interface {

	// ConsistentIndex 返回当前执行条目的一致索引。
	ConsistentIndex() uint64

	// SetConsistentIndex 设置当前执行条目的一致索引。
	SetConsistentIndex(v uint64, term uint64)

	// UnsafeSave 必须在持有tx上的锁的情况下被调用。 它将一致索引保存到底层稳定存储中。
	UnsafeSave(tx backend.BatchTx)

	// SetBackend 为ConsistentIndexer设置可用的backend.BatchTx。
	SetBackend(be Backend)
}

// consistentIndex implements the ConsistentIndexer interface.
type consistentIndex struct {
	// consistentIndex represents the offset of an entry in a consistent replica log.
	// It caches the "consistent_index" key's value.
	// Accessed through atomics so必须是64-bit aligned.
	consistentIndex uint64
	// term represents the RAFT term of committed entry in a consistent replica log.
	// Accessed through atomics so必须是64-bit aligned.
	// The value is being persisted in the backend since v3.5.
	term uint64

	// be is used for initial read consistentIndex
	be Backend
	// mutex is protecting be.
	mutex sync.Mutex
}

// NewConsistentIndex 返回一个一致性索引
// 如果be is nil,必须在首次调用前执行ConsistentIndex方法
func NewConsistentIndex(be Backend) ConsistentIndexer {
	return &consistentIndex{be: be}
}

func (ci *consistentIndex) ConsistentIndex() uint64 {
	if index := atomic.LoadUint64(&ci.consistentIndex); index > 0 {
		return index
	}
	ci.mutex.Lock()
	defer ci.mutex.Unlock()

	v, term := ReadConsistentIndex(ci.be.BatchTx())
	ci.SetConsistentIndex(v, term)
	return v
}

func (ci *consistentIndex) SetConsistentIndex(v uint64, term uint64) {
	atomic.StoreUint64(&ci.consistentIndex, v)
	atomic.StoreUint64(&ci.term, term)
}

func (ci *consistentIndex) UnsafeSave(tx backend.BatchTx) {
	index := atomic.LoadUint64(&ci.consistentIndex)
	term := atomic.LoadUint64(&ci.term)
	UnsafeUpdateConsistentIndex(tx, index, term, true)
}

func (ci *consistentIndex) SetBackend(be Backend) {
	ci.mutex.Lock()
	defer ci.mutex.Unlock()
	ci.be = be
	// After the backend is changed, the first access should re-read it.
	ci.SetConsistentIndex(0, 0)
}

func NewFakeConsistentIndex(index uint64) ConsistentIndexer {
	return &fakeConsistentIndex{index: index}
}

type fakeConsistentIndex struct {
	index uint64
	term  uint64
}

func (f *fakeConsistentIndex) ConsistentIndex() uint64 { return f.index }

func (f *fakeConsistentIndex) SetConsistentIndex(index uint64, term uint64) {
	atomic.StoreUint64(&f.index, index)
	atomic.StoreUint64(&f.term, term)
}

func (f *fakeConsistentIndex) UnsafeSave(_ backend.BatchTx) {}
func (f *fakeConsistentIndex) SetBackend(_ Backend)         {}

// UnsafeCreateMetaBucket creates the `meta` bucket (if it does not exists yet).
func UnsafeCreateMetaBucket(tx backend.BatchTx) {
	tx.UnsafeCreateBucket(buckets.Meta)
}

// CreateMetaBucket 创建meta bucket，如果不存在
func CreateMetaBucket(tx backend.BatchTx) {
	tx.Lock()
	defer tx.Unlock()
	tx.UnsafeCreateBucket(buckets.Meta)
}

// unsafeGetConsistentIndex loads consistent index & term from given transaction.
// returns 0,0 if the data are not found.
// Term is persisted since v3.5.
func unsafeReadConsistentIndex(tx backend.ReadTx) (uint64, uint64) {
	_, vs := tx.UnsafeRange(buckets.Meta, buckets.MetaConsistentIndexKeyName, nil, 0)
	if len(vs) == 0 {
		return 0, 0
	}
	v := binary.BigEndian.Uint64(vs[0])
	_, ts := tx.UnsafeRange(buckets.Meta, buckets.MetaTermKeyName, nil, 0)
	if len(ts) == 0 {
		return v, 0
	}
	t := binary.BigEndian.Uint64(ts[0])
	return v, t
}

// ReadConsistentIndex loads consistent index and term from given transaction.
// returns 0 if the data are not found.
func ReadConsistentIndex(tx backend.ReadTx) (uint64, uint64) {
	tx.Lock()
	defer tx.Unlock()
	return unsafeReadConsistentIndex(tx)
}

func UnsafeUpdateConsistentIndex(tx backend.BatchTx, index uint64, term uint64, onlyGrow bool) {
	if index == 0 {
		// Never save 0 as it means that we didn't loaded the real index yet.
		return
	}

	if onlyGrow {
		oldi, oldTerm := unsafeReadConsistentIndex(tx)
		if term < oldTerm {
			return
		}
		if term == oldTerm && index <= oldi {
			return
		}
	}

	bs1 := make([]byte, 8)
	binary.BigEndian.PutUint64(bs1, index)
	// put the index into the underlying backend
	// tx has been locked in TxnBegin, so there is no need to lock it again
	tx.UnsafePut(buckets.Meta, buckets.MetaConsistentIndexKeyName, bs1)
	if term > 0 {
		bs2 := make([]byte, 8)
		binary.BigEndian.PutUint64(bs2, term)
		tx.UnsafePut(buckets.Meta, buckets.MetaTermKeyName, bs2)
	}
}

func UpdateConsistentIndex(tx backend.BatchTx, index uint64, term uint64, onlyGrow bool) {
	tx.Lock()
	defer tx.Unlock()
	UnsafeUpdateConsistentIndex(tx, index, term, onlyGrow)
}
