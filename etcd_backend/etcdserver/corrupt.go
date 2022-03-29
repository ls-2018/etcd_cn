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

package etcdserver

import (
	"context"
	"fmt"
	"time"

	"github.com/ls-2018/etcd_cn/client_sdk/pkg/types"
	"github.com/ls-2018/etcd_cn/etcd_backend/mvcc"
	"github.com/ls-2018/etcd_cn/offical/api/v3/v3rpc/rpctypes"
	pb "github.com/ls-2018/etcd_cn/offical/etcdserverpb"
	"github.com/ls-2018/etcd_cn/pkg/traceutil"

	"go.uber.org/zap"
)

// CheckInitialHashKV compares initial hash values with its peers
// before serving any peer/client traffic. Only mismatch when hashes
// are different at requested revision, with same compact revision.
func (s *EtcdServer) CheckInitialHashKV() error {
	if !s.Cfg.InitialCorruptCheck {
		return nil
	}

	lg := s.Logger()

	lg.Info(
		"starting initial corruption check",
		zap.String("local-member-id", s.ID().String()),
		zap.Duration("timeout", s.Cfg.ReqTimeout()),
	)

	h, rev, crev, err := s.kv.HashByRev(0)
	if err != nil {
		return fmt.Errorf("%s failed to fetch hash (%v)", s.ID(), err)
	}
	peers := s.getPeerHashKVs(rev)
	mismatch := 0
	for _, p := range peers {
		if p.resp != nil {
			peerID := types.ID(p.resp.Header.MemberId)
			fields := []zap.Field{
				zap.String("local-member-id", s.ID().String()),
				zap.Int64("local-member-revision", rev),
				zap.Int64("local-member-compact-revision", crev),
				zap.Uint32("local-member-hash", h),
				zap.String("remote-peer-id", peerID.String()),
				zap.Strings("remote-peer-endpoints", p.eps),
				zap.Int64("remote-peer-revision", p.resp.Header.Revision),
				zap.Int64("remote-peer-compact-revision", p.resp.CompactRevision),
				zap.Uint32("remote-peer-hash", p.resp.Hash),
			}

			if h != p.resp.Hash {
				if crev == p.resp.CompactRevision {
					lg.Warn("found different hash values from remote peer", fields...)
					mismatch++
				} else {
					lg.Warn("found different compact revision values from remote peer", fields...)
				}
			}

			continue
		}

		if p.err != nil {
			switch p.err {
			case rpctypes.ErrFutureRev:
				lg.Warn(
					"cannot fetch hash from slow remote peer",
					zap.String("local-member-id", s.ID().String()),
					zap.Int64("local-member-revision", rev),
					zap.Int64("local-member-compact-revision", crev),
					zap.Uint32("local-member-hash", h),
					zap.String("remote-peer-id", p.id.String()),
					zap.Strings("remote-peer-endpoints", p.eps),
					zap.Error(err),
				)
			case rpctypes.ErrCompacted:
				lg.Warn(
					"cannot fetch hash from remote peer; local member is behind",
					zap.String("local-member-id", s.ID().String()),
					zap.Int64("local-member-revision", rev),
					zap.Int64("local-member-compact-revision", crev),
					zap.Uint32("local-member-hash", h),
					zap.String("remote-peer-id", p.id.String()),
					zap.Strings("remote-peer-endpoints", p.eps),
					zap.Error(err),
				)
			}
		}
	}
	if mismatch > 0 {
		return fmt.Errorf("%s found data inconsistency with peers", s.ID())
	}

	lg.Info(
		"initial corruption checking passed; no corruption",
		zap.String("local-member-id", s.ID().String()),
	)
	return nil
}

func (s *EtcdServer) monitorKVHash() {
	t := s.Cfg.CorruptCheckTime
	if t == 0 {
		return
	}

	lg := s.Logger()
	lg.Info("启用损坏检查", zap.String("local-member-id", s.ID().String()), zap.Duration("interval", t))

	for {
		select {
		case <-s.stopping:
			return
		case <-time.After(t):
		}
		if !s.isLeader() {
			continue
		}
		if err := s.checkHashKV(); err != nil {
			lg.Warn("failed to check hash KV", zap.Error(err))
		}
	}
}

func (s *EtcdServer) checkHashKV() error {
	lg := s.Logger()

	h, rev, crev, err := s.kv.HashByRev(0)
	if err != nil {
		return err
	}
	peers := s.getPeerHashKVs(rev)

	ctx, cancel := context.WithTimeout(context.Background(), s.Cfg.ReqTimeout())
	err = s.linearizableReadNotify(ctx)
	cancel()
	if err != nil {
		return err
	}

	h2, rev2, crev2, err := s.kv.HashByRev(0)
	if err != nil {
		return err
	}

	alarmed := false
	mismatch := func(id uint64) {
		if alarmed {
			return
		}
		alarmed = true
		a := &pb.AlarmRequest{
			MemberID: id,
			Action:   pb.AlarmRequest_ACTIVATE,
			Alarm:    pb.AlarmType_CORRUPT,
		}
		s.GoAttach(func() {
			s.raftRequest(s.ctx, pb.InternalRaftRequest{Alarm: a})
		})
	}

	if h2 != h && rev2 == rev && crev == crev2 {
		lg.Warn(
			"found hash mismatch",
			zap.Int64("revision-1", rev),
			zap.Int64("compact-revision-1", crev),
			zap.Uint32("hash-1", h),
			zap.Int64("revision-2", rev2),
			zap.Int64("compact-revision-2", crev2),
			zap.Uint32("hash-2", h2),
		)
		mismatch(uint64(s.ID()))
	}

	checkedCount := 0
	for _, p := range peers {
		if p.resp == nil {
			continue
		}
		checkedCount++
		id := p.resp.Header.MemberId

		// leader expects follower's latest revision less than or equal to leader's
		if p.resp.Header.Revision > rev2 {
			lg.Warn(
				"revision from follower必须是less than or equal to leader's",
				zap.Int64("leader-revision", rev2),
				zap.Int64("follower-revision", p.resp.Header.Revision),
				zap.String("follower-peer-id", types.ID(id).String()),
			)
			mismatch(id)
		}

		// leader expects follower's latest compact revision less than or equal to leader's
		if p.resp.CompactRevision > crev2 {
			lg.Warn(
				"compact revision from follower必须是less than or equal to leader's",
				zap.Int64("leader-compact-revision", crev2),
				zap.Int64("follower-compact-revision", p.resp.CompactRevision),
				zap.String("follower-peer-id", types.ID(id).String()),
			)
			mismatch(id)
		}

		// follower's compact revision is leader's old one, then hashes must match
		if p.resp.CompactRevision == crev && p.resp.Hash != h {
			lg.Warn(
				"same compact revision then hashes must match",
				zap.Int64("leader-compact-revision", crev2),
				zap.Uint32("leader-hash", h),
				zap.Int64("follower-compact-revision", p.resp.CompactRevision),
				zap.Uint32("follower-hash", p.resp.Hash),
				zap.String("follower-peer-id", types.ID(id).String()),
			)
			mismatch(id)
		}
	}
	lg.Info("finished peer corruption check", zap.Int("number-of-peers-checked", checkedCount))
	return nil
}

type peerInfo struct {
	id  types.ID
	eps []string
}

type peerHashKVResp struct {
	peerInfo
	resp *pb.HashKVResponse
	err  error
}

func (s *EtcdServer) getPeerHashKVs(rev int64) []*peerHashKVResp {
	// TODO: handle the case when "s.cluster.Members" have not
	// been populated (e.g. no snapshot to load from disk)
	members := s.cluster.Members()
	peers := make([]peerInfo, 0, len(members))
	for _, m := range members {
		if m.ID == s.ID() {
			continue
		}
		peers = append(peers, peerInfo{id: m.ID, eps: m.PeerURLs})
	}

	lg := s.Logger()

	var resps []*peerHashKVResp
	for _, p := range peers {
		if len(p.eps) == 0 {
			continue
		}

		respsLen := len(resps)
		var lastErr error
		for _, ep := range p.eps {
			ctx, cancel := context.WithTimeout(context.Background(), s.Cfg.ReqTimeout())
			resp, lastErr := s.getPeerHashKVHTTP(ctx, ep, rev)
			cancel()
			if lastErr == nil {
				resps = append(resps, &peerHashKVResp{peerInfo: p, resp: resp, err: nil})
				break
			}
			lg.Warn(
				"failed hash kv request",
				zap.String("local-member-id", s.ID().String()),
				zap.Int64("requested-revision", rev),
				zap.String("remote-peer-endpoint", ep),
				zap.Error(lastErr),
			)
		}

		// failed to get hashKV from all endpoints of this peer
		if respsLen == len(resps) {
			resps = append(resps, &peerHashKVResp{peerInfo: p, resp: nil, err: lastErr})
		}
	}
	return resps
}

type applierV3Corrupt struct {
	applierV3
}

func newApplierV3Corrupt(a applierV3) *applierV3Corrupt { return &applierV3Corrupt{a} }

func (a *applierV3Corrupt) Put(ctx context.Context, txn mvcc.TxnWrite, p *pb.PutRequest) (*pb.PutResponse, *traceutil.Trace, error) {
	return nil, nil, ErrCorrupt
}

func (a *applierV3Corrupt) Range(ctx context.Context, txn mvcc.TxnRead, p *pb.RangeRequest) (*pb.RangeResponse, error) {
	return nil, ErrCorrupt
}

func (a *applierV3Corrupt) DeleteRange(txn mvcc.TxnWrite, p *pb.DeleteRangeRequest) (*pb.DeleteRangeResponse, error) {
	return nil, ErrCorrupt
}

func (a *applierV3Corrupt) Txn(ctx context.Context, rt *pb.TxnRequest) (*pb.TxnResponse, *traceutil.Trace, error) {
	return nil, nil, ErrCorrupt
}

func (a *applierV3Corrupt) Compaction(compaction *pb.CompactionRequest) (*pb.CompactionResponse, <-chan struct{}, *traceutil.Trace, error) {
	return nil, nil, nil, ErrCorrupt
}

func (a *applierV3Corrupt) LeaseGrant(lc *pb.LeaseGrantRequest) (*pb.LeaseGrantResponse, error) {
	return nil, ErrCorrupt
}

func (a *applierV3Corrupt) LeaseRevoke(lc *pb.LeaseRevokeRequest) (*pb.LeaseRevokeResponse, error) {
	return nil, ErrCorrupt
}
