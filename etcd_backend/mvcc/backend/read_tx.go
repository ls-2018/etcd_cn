// Copyright 2017 The etcd Authors
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

package backend

import (
	"math"
	"sync"

	bolt "go.etcd.io/bbolt"
)

// IsSafeRangeBucket is a hack to avoid inadvertently reading duplicate keys;
// overwrites on a bucket should only fetch with limit=1, but IsSafeRangeBucket
// is known to never overwrite any key so range is safe.
// IsSafeRangeBucket是一个黑科技,用来避免无意中读取重复的键.
// 对一个桶的覆盖应该只在limit=1的情况下获取,但IsSafeRangeBucket是已知的,永远不会覆盖任何键,所以范围是安全的.

// 负责读请求
type ReadTx interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
	UnsafeRange(bucket Bucket, key, endKey []byte, limit int64) (keys [][]byte, vals [][]byte)
	UnsafeForEach(bucket Bucket, visitor func(k, v []byte) error) error // 对指定的桶,所有k,v遍历
}

// baseReadTx的访问是并发的所以需要读写锁来保护.
type baseReadTx struct {
	// 写事务执行End时候需要获取这个写锁然后把写事务的更新写到 baseReadTx 的buffer里面；
	// 创建 concurrentReadTx 时候需要获取读锁因为需要拷贝buffer
	mu      sync.RWMutex  // 保护 txReadBuffer 的访问
	buf     txReadBuffer  // 用于加速读效率的缓存
	txMu    *sync.RWMutex // 保护下面的tx和buckets
	tx      *bolt.Tx
	buckets map[BucketID]*bolt.Bucket
	txWg    *sync.WaitGroup // txWg 保护 tx 在批处理间隔结束时不会被回滚直到使用此 tx 的所有读取完成.
}

func (baseReadTx *baseReadTx) UnsafeForEach(bucket Bucket, visitor func(k, v []byte) error) error {
	dups := make(map[string]struct{})
	getDups := func(k, v []byte) error {
		dups[string(k)] = struct{}{}
		return nil
	}
	visitNoDup := func(k, v []byte) error {
		if _, ok := dups[string(k)]; ok {
			return nil
		}
		return visitor(k, v)
	}
	if err := baseReadTx.buf.ForEach(bucket, getDups); err != nil {
		return err
	}
	baseReadTx.txMu.Lock()
	err := unsafeForEach(baseReadTx.tx, bucket, visitNoDup)
	baseReadTx.txMu.Unlock()
	if err != nil {
		return err
	}
	return baseReadTx.buf.ForEach(bucket, visitor)
}

func (baseReadTx *baseReadTx) UnsafeRange(bucketType Bucket, key, endKey []byte, limit int64) ([][]byte, [][]byte) {
	if endKey == nil {
		// forbid duplicates for single keys
		limit = 1
	}
	if limit <= 0 {
		limit = math.MaxInt64
	}
	if limit > 1 && !bucketType.IsSafeRangeBucket() {
		panic("do not use unsafeRange on non-keys bucket")
	}
	// 首先从缓存中查询键值对
	keys, vals := baseReadTx.buf.Range(bucketType, key, endKey, limit)
	// 检测缓存中返回的键值对是否达到Limit的要求如果达到Limit的指定上限直接返回缓存的查询结果
	if int64(len(keys)) == limit {
		return keys, vals
	}

	// find/cache bucket
	bn := bucketType.ID()
	baseReadTx.txMu.RLock()
	bucket, ok := baseReadTx.buckets[bn]
	baseReadTx.txMu.RUnlock()
	lockHeld := false
	if !ok {
		baseReadTx.txMu.Lock()
		lockHeld = true
		bucket = baseReadTx.tx.Bucket(bucketType.Name())
		baseReadTx.buckets[bn] = bucket
	}

	// ignore missing bucket since may have been created in this batch
	if bucket == nil {
		if lockHeld {
			baseReadTx.txMu.Unlock()
		}
		return keys, vals
	}
	if !lockHeld {
		baseReadTx.txMu.Lock()
	}
	c := bucket.Cursor()
	baseReadTx.txMu.Unlock()
	// 将查询缓存的结采与查询 BlotDB 的结果合并 然后返回
	k2, v2 := unsafeRange(c, key, endKey, limit-int64(len(keys)))
	return append(k2, keys...), append(v2, vals...)
}

// 负责读请求
type readTx struct {
	baseReadTx
}

func (rt *readTx) Lock()    { rt.mu.Lock() }
func (rt *readTx) Unlock()  { rt.mu.Unlock() }
func (rt *readTx) RLock()   { rt.mu.RLock() }
func (rt *readTx) RUnlock() { rt.mu.RUnlock() }

func (rt *readTx) reset() {
	rt.buf.reset()
	rt.buckets = make(map[BucketID]*bolt.Bucket)
	rt.tx = nil
	rt.txWg = new(sync.WaitGroup)
}

type concurrentReadTx struct {
	baseReadTx
}

func (rt *concurrentReadTx) Lock()   {}
func (rt *concurrentReadTx) Unlock() {}

// RLock is no-op. concurrentReadTx does not need to be locked after it is created.
func (rt *concurrentReadTx) RLock() {}

// RUnlock signals the end of concurrentReadTx.
func (rt *concurrentReadTx) RUnlock() { rt.txWg.Done() }
