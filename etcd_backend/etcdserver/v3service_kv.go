package etcdserver

import (
	"context"
	"time"

	"github.com/ls-2018/etcd_cn/etcd_backend/auth"
	pb "github.com/ls-2018/etcd_cn/offical/etcdserverpb"
	"github.com/ls-2018/etcd_cn/pkg/traceutil"
)

type RaftKV interface {
	Range(ctx context.Context, r *pb.RangeRequest) (*pb.RangeResponse, error)
	Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error)
	DeleteRange(ctx context.Context, r *pb.DeleteRangeRequest) (*pb.DeleteRangeResponse, error)
	Txn(ctx context.Context, r *pb.TxnRequest) (*pb.TxnResponse, error)
	Compact(ctx context.Context, r *pb.CompactionRequest) (*pb.CompactionResponse, error)
}

// Range TODO 待看
func (s *EtcdServer) Range(ctx context.Context, r *pb.RangeRequest) (*pb.RangeResponse, error) {
	trace := traceutil.New("range",
		s.Logger(),
		traceutil.Field{Key: "range_begin", Value: string(r.Key)},
		traceutil.Field{Key: "range_end", Value: string(r.RangeEnd)},
	)
	ctx = context.WithValue(ctx, traceutil.TraceKey, trace)

	var resp *pb.RangeResponse
	var err error
	defer func(start time.Time) {
		if resp != nil {
			trace.AddField(
				traceutil.Field{Key: "response_count", Value: len(resp.Kvs)},
				traceutil.Field{Key: "response_revision", Value: resp.Header.Revision},
			)
		}
	}(time.Now())

	if !r.Serializable {
		err = s.linearizableReadNotify(ctx)
		trace.Step("agreement among raft nodes before linearized reading")
		if err != nil {
			return nil, err
		}
	}
	chk := func(ai *auth.AuthInfo) error {
		return s.authStore.IsRangePermitted(ai, r.Key, r.RangeEnd)
	}

	get := func() { resp, err = s.applyV3Base.Range(ctx, nil, r) }
	if serr := s.doSerialize(ctx, chk, get); serr != nil {
		err = serr
		return nil, err
	}
	return resp, err
}

func (s *EtcdServer) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	ctx = context.WithValue(ctx, traceutil.StartTimeKey, time.Now())
	resp, err := s.raftRequest(ctx, pb.InternalRaftRequest{Put: r})
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PutResponse), nil
}

func (s *EtcdServer) DeleteRange(ctx context.Context, r *pb.DeleteRangeRequest) (*pb.DeleteRangeResponse, error) {
	resp, err := s.raftRequest(ctx, pb.InternalRaftRequest{DeleteRange: r})
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DeleteRangeResponse), nil
}

func (s *EtcdServer) Txn(ctx context.Context, r *pb.TxnRequest) (*pb.TxnResponse, error) {
	if isTxnReadonly(r) {
		trace := traceutil.New("transaction",
			s.Logger(),
			traceutil.Field{Key: "read_only", Value: true},
		)
		ctx = context.WithValue(ctx, traceutil.TraceKey, trace)
		if !isTxnSerializable(r) {
			err := s.linearizableReadNotify(ctx)
			trace.Step("agreement among raft nodes before linearized reading")
			if err != nil {
				return nil, err
			}
		}
		var resp *pb.TxnResponse
		var err error
		chk := func(ai *auth.AuthInfo) error {
			return checkTxnAuth(s.authStore, ai, r)
		}

		get := func() { resp, _, err = s.applyV3Base.Txn(ctx, r) }
		if serr := s.doSerialize(ctx, chk, get); serr != nil {
			return nil, serr
		}
		return resp, err
	}

	ctx = context.WithValue(ctx, traceutil.StartTimeKey, time.Now())
	resp, err := s.raftRequest(ctx, pb.InternalRaftRequest{Txn: r})
	if err != nil {
		return nil, err
	}
	return resp.(*pb.TxnResponse), nil
}

func (s *EtcdServer) Compact(ctx context.Context, r *pb.CompactionRequest) (*pb.CompactionResponse, error) {
	startTime := time.Now()
	result, err := s.processInternalRaftRequestOnce(ctx, pb.InternalRaftRequest{Compaction: r})
	trace := traceutil.TODO()
	if result != nil && result.trace != nil {
		trace = result.trace

		applyStart := result.trace.GetStartTime()
		result.trace.SetStartTime(startTime)
		trace.InsertStep(0, applyStart, "process raft request")
	}
	if r.Physical && result != nil && result.physc != nil {
		<-result.physc
		// The compaction is done deleting keys; the hash is now settled
		// but the data is not necessarily committed. If there's a crash,
		// the hash may revert to a hash prior to compaction completing
		// if the compaction resumes. Force the finished compaction to
		// commit so it won't resume following a crash.
		s.backend.ForceCommit()
		trace.Step("physically apply compaction")
	}
	if err != nil {
		return nil, err
	}
	if result.err != nil {
		return nil, result.err
	}
	resp := result.resp.(*pb.CompactionResponse)
	if resp == nil {
		resp = &pb.CompactionResponse{}
	}
	if resp.Header == nil {
		resp.Header = &pb.ResponseHeader{}
	}
	resp.Header.Revision = s.kv.Rev()
	trace.AddField(traceutil.Field{Key: "response_revision", Value: resp.Header.Revision})
	return resp, nil
}
