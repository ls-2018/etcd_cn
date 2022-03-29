package mvcc

import (
	"github.com/ls-2018/etcd_cn/etcd_backend/mvcc/backend"
	"github.com/ls-2018/etcd_cn/pkg/traceutil"
)

type KV interface {
	ReadView
	WriteView
	Read(mode ReadTxMode, trace *traceutil.Trace) TxnRead // 创建读事务
	Write(trace *traceutil.Trace) TxnWrite                // 创建写事务
	// Hash computes the hash of the KV's backend.
	Hash() (hash uint32, revision int64, err error)
	// HashByRev computes the hash of all MVCC revisions up to a given revision.
	HashByRev(rev int64) (hash uint32, revision int64, compactRev int64, err error)
	// Compact frees all superseded keys with revisions less than rev.
	Compact(trace *traceutil.Trace, rev int64) (<-chan struct{}, error)
	// Commit commits outstanding txns into the underlying backend.
	Commit()
	// Restore restores the KV store from a backend.
	Restore(b backend.Backend) error
	Close() error
}

type WatchableKV interface {
	KV
	Watchable
}

type Watchable interface {
	NewWatchStream() WatchStream
}
