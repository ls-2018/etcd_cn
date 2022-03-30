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

package mvcc

import (
	"context"

	"github.com/ls-2018/etcd_cn/etcd_backend/lease"
	"github.com/ls-2018/etcd_cn/etcd_backend/mvcc/backend"
	"github.com/ls-2018/etcd_cn/etcd_backend/mvcc/buckets"
	"github.com/ls-2018/etcd_cn/offical/api/v3/mvccpb"
	"github.com/ls-2018/etcd_cn/pkg/traceutil"
	"go.uber.org/zap"
)

type storeTxnRead struct {
	s        *store
	tx       backend.ReadTx
	firstRev int64
	rev      int64 // 总的修订版本
	trace    *traceutil.Trace
}

func (tr *storeTxnRead) FirstRev() int64 { return tr.firstRev }
func (tr *storeTxnRead) Rev() int64      { return tr.rev }

// Range OK
func (tr *storeTxnRead) Range(ctx context.Context, key, end []byte, ro RangeOptions) (r *RangeResult, err error) {
	return tr.rangeKeys(ctx, key, end, tr.Rev(), ro)
}

// OK
func (tr *storeTxnRead) rangeKeys(ctx context.Context, key, end []byte, curRev int64, ro RangeOptions) (*RangeResult, error) {
	rev := ro.Rev // 指定修订版本
	if rev > curRev {
		return &RangeResult{KVs: nil, Count: -1, Rev: curRev}, ErrFutureRev
	}
	if rev <= 0 {
		rev = curRev
	}
	if rev < tr.s.compactMainRev {
		return &RangeResult{KVs: nil, Count: -1, Rev: 0}, ErrCompacted
	}
	if ro.Count {
		total := tr.s.kvindex.CountRevisions(key, end, rev)
		tr.trace.Step("从内存索引树中统计修订数")
		return &RangeResult{KVs: nil, Count: total, Rev: curRev}, nil
	}
	// 获取版本数据
	revpairs, total := tr.s.kvindex.Revisions(key, end, rev, int(ro.Limit))
	tr.trace.Step("从内存索引树中获取指定范围的keys")
	if len(revpairs) == 0 {
		return &RangeResult{KVs: nil, Count: total, Rev: curRev}, nil
	}

	limit := int(ro.Limit)
	if limit <= 0 || limit > len(revpairs) {
		limit = len(revpairs) // 实际收到的数据量
	}

	kvs := make([]mvccpb.KeyValue, limit)
	revBytes := newRevBytes() // len 为17的数组
	// 拿着索引数据去bolt.db 查数据
	for i, revpair := range revpairs[:len(kvs)] {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		revToBytes(revpair, revBytes)
		// 根据修订版本获取数据
		_, vs := tr.tx.UnsafeRange(buckets.Key, revBytes, nil, 0)
		if len(vs) != 1 {
			tr.s.lg.Fatal("Range找不到修订对", zap.Int64("revision-main", revpair.main), zap.Int64("revision-sub", revpair.sub))
		}
		if err := kvs[i].Unmarshal([]byte(vs[0])); err != nil {
			tr.s.lg.Fatal(
				"反序列失败 mvccpb.KeyValue",
				zap.Error(err),
			)
		}
	}
	tr.trace.Step("从bolt.db 中range key")
	return &RangeResult{KVs: kvs, Count: total, Rev: curRev}, nil
}

func (tr *storeTxnRead) End() {
	tr.tx.RUnlock() // RUnlock signals the end of concurrentReadTx.
	tr.s.mu.RUnlock()
}

type storeTxnWrite struct {
	storeTxnRead
	tx       backend.BatchTx
	beginRev int64 // 是TXN开始时的修订版本;它将写到下次修订。
	changes  []mvccpb.KeyValue
}

func (tw *storeTxnWrite) Rev() int64 { return tw.beginRev }

func (tw *storeTxnWrite) Range(ctx context.Context, key, end []byte, ro RangeOptions) (r *RangeResult, err error) {
	rev := tw.beginRev
	if len(tw.changes) > 0 {
		rev++
	}
	return tw.rangeKeys(ctx, key, end, rev, ro)
}

func (tw *storeTxnWrite) DeleteRange(key, end []byte) (int64, int64) {
	if n := tw.deleteRange(key, end); n != 0 || len(tw.changes) > 0 {
		return n, tw.beginRev + 1
	}
	return 0, tw.beginRev
}

func (tw *storeTxnWrite) Put(key, value []byte, lease lease.LeaseID) int64 {
	tw.put(key, value, lease)
	return tw.beginRev + 1
}

func (tw *storeTxnWrite) End() {
	// only update index if the txn modifies the mvcc state.
	if len(tw.changes) != 0 {
		// hold revMu lock to prevent new read txns from opening until writeback.
		tw.s.revMu.Lock()
		tw.s.currentRev++
	}
	tw.tx.Unlock()
	if len(tw.changes) != 0 {
		tw.s.revMu.Unlock()
	}
	tw.s.mu.RUnlock()
}

func (tw *storeTxnWrite) put(key, value []byte, leaseID lease.LeaseID) {
	rev := tw.beginRev + 1 // 事务开始时有一个ID，写这个操作，对应的ID应+1
	c := rev
	oldLease := lease.NoLease

	// 如果该键之前存在，使用它之前创建的并获取它之前的leaseID
	_, created, beforeVersion, err := tw.s.kvindex.Get(key, rev) // 0,0,nil  <= rev的最新修改
	if err == nil {
		c = created.main
		oldLease = tw.s.le.GetLease(lease.LeaseItem{Key: string(key)})
		tw.trace.Step("获取键先前的created_revision和leaseID")
	}
	indexBytes := newRevBytes()
	idxRev := revision{main: rev, sub: int64(len(tw.changes))}
	revToBytes(idxRev, indexBytes)

	kv := mvccpb.KeyValue{
		Key:            string(key),
		Value:          string(value),
		CreateRevision: c,                 // 当前代，创建时的修订版本
		ModRevision:    rev,               // 修订版本
		Version:        beforeVersion + 1, // Version是key的版本。删除键会将该键的版本重置为0，对键的任何修改都会增加它的版本。
		Lease:          int64(leaseID),    // 租约ID
	}

	d, err := kv.Marshal()
	if err != nil {
		tw.storeTxnRead.s.lg.Fatal("序列化失败 mvccpb.KeyValue", zap.Error(err))
	}

	tw.trace.Step("序列化 mvccpb.KeyValue")
	tw.tx.UnsafeSeqPut(buckets.Key, indexBytes, d) // ✅ 写入db,buf
	tw.s.kvindex.Put(key, idxRev)
	tw.changes = append(tw.changes, kv)
	tw.trace.Step("存储键值对到bolt.db")

	if oldLease != lease.NoLease {
		if tw.s.le == nil {
			panic("no lessor to detach lease")
		}
		err = tw.s.le.Detach(oldLease, []lease.LeaseItem{{Key: string(key)}})
		if err != nil {
			tw.storeTxnRead.s.lg.Error("从key中分离旧的租约失败", zap.Error(err))
		}
	}
	if leaseID != lease.NoLease {
		if tw.s.le == nil {
			panic("租约控制方为空")
		}
		err = tw.s.le.Attach(leaseID, []lease.LeaseItem{{Key: string(key)}})
		if err != nil {
			panic("租约附加失败")
		}
	}
	tw.trace.Step("附加租约到key")
}

func (tw *storeTxnWrite) deleteRange(key, end []byte) int64 {
	rrev := tw.beginRev
	if len(tw.changes) > 0 {
		rrev++
	}
	keys, _ := tw.s.kvindex.Range(key, end, rrev)
	if len(keys) == 0 {
		return 0
	}
	for _, key := range keys {
		tw.delete(key)
	}
	return int64(len(keys))
}

func (tw *storeTxnWrite) delete(key []byte) {
	indexBytes := newRevBytes()
	idxRev := revision{main: tw.beginRev + 1, sub: int64(len(tw.changes))}
	revToBytes(idxRev, indexBytes)

	indexBytes = appendMarkTombstone(tw.storeTxnRead.s.lg, indexBytes)

	kv := mvccpb.KeyValue{Key: string(key)}

	d, err := kv.Marshal()
	if err != nil {
		tw.storeTxnRead.s.lg.Fatal(
			"failed to marshal mvccpb.KeyValue",
			zap.Error(err),
		)
	}

	tw.tx.UnsafeSeqPut(buckets.Key, indexBytes, d)
	err = tw.s.kvindex.Tombstone(key, idxRev)
	if err != nil {
		tw.storeTxnRead.s.lg.Fatal(
			"failed to tombstone an existing key",
			zap.String("key", string(key)),
			zap.Error(err),
		)
	}
	tw.changes = append(tw.changes, kv)

	item := lease.LeaseItem{Key: string(key)}
	leaseID := tw.s.le.GetLease(item)

	if leaseID != lease.NoLease {
		err = tw.s.le.Detach(leaseID, []lease.LeaseItem{item})
		if err != nil {
			tw.storeTxnRead.s.lg.Error(
				"failed to detach old lease from a key",
				zap.Error(err),
			)
		}
	}
}

func (tw *storeTxnWrite) Changes() []mvccpb.KeyValue { return tw.changes }
