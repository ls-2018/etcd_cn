package etcdserver

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api/membership"
	"github.com/ls-2018/etcd_cn/raft/raftpb"
	"go.uber.org/zap"

	"github.com/ls-2018/etcd_cn/client_sdk/pkg/types"
)

// LinearizableReadNotify 一致性读
func (s *EtcdServer) LinearizableReadNotify(ctx context.Context) error {
	return s.linearizeReadNotify(ctx)
}

func (s *EtcdServer) AddMember(ctx context.Context, memb membership.Member) ([]*membership.Member, error) {
	if err := s.checkMembershipOperationPermission(ctx); err != nil {
		return nil, err
	}

	b, err := json.Marshal(memb)
	if err != nil {
		return nil, err
	}

	lg := s.Logger()
	if !s.Cfg.StrictReconfigCheck {
	} else {
		// protect quorum when adding voting member
		if !memb.IsLearner && !s.cluster.IsReadyToAddVotingMember() {
			lg.Warn(
				"rejecting member add request; not enough healthy members",
				zap.String("local-member-id", s.ID().String()),
				zap.String("requested-member-add", fmt.Sprintf("%+v", memb)),
				zap.Error(ErrNotEnoughStartedMembers),
			)
			return nil, ErrNotEnoughStartedMembers
		}
		if !isConnectedFullySince(s.r.transport, time.Now().Add(-HealthInterval), s.ID(), s.cluster.VotingMembers()) {
			lg.Warn(
				"rejecting member add request; local member has not been connected to all peers, reconfigure breaks active quorum",
				zap.String("local-member-id", s.ID().String()),
				zap.String("requested-member-add", fmt.Sprintf("%+v", memb)),
				zap.Error(ErrUnhealthy),
			)
			return nil, ErrUnhealthy
		}
	}
	// by default StrictReconfigCheck is enabled; reject new members if unhealthy.
	if err != nil {
		return nil, err
	}

	cc := raftpb.ConfChangeV1{
		Type:    raftpb.ConfChangeAddNode,
		NodeID:  uint64(memb.ID),
		Context: string(b),
	}

	if memb.IsLearner {
		cc.Type = raftpb.ConfChangeAddLearnerNode
	}

	return s.configure(ctx, cc)
}

func (s *EtcdServer) RemoveMember(ctx context.Context, id uint64) ([]*membership.Member, error) {
	if err := s.checkMembershipOperationPermission(ctx); err != nil {
		return nil, err
	}

	// by default StrictReconfigCheck is enabled; reject removal if leads to quorum loss
	if err := s.mayRemoveMember(types.ID(id)); err != nil {
		return nil, err
	}

	cc := raftpb.ConfChangeV1{
		Type:   raftpb.ConfChangeRemoveNode,
		NodeID: id,
	}
	return s.configure(ctx, cc)
}

func (s *EtcdServer) UpdateMember(ctx context.Context, memb membership.Member) ([]*membership.Member, error) {
	b, merr := json.Marshal(memb)
	if merr != nil {
		return nil, merr
	}

	if err := s.checkMembershipOperationPermission(ctx); err != nil {
		return nil, err
	}
	cc := raftpb.ConfChangeV1{
		Type:    raftpb.ConfChangeUpdateNode,
		NodeID:  uint64(memb.ID),
		Context: string(b),
	}
	return s.configure(ctx, cc)
}

// PromoteMember 将learner节点提升为voter
func (s *EtcdServer) PromoteMember(ctx context.Context, id uint64) ([]*membership.Member, error) {
	// only raft leader has information on whether the to-backend-promoted learner node is ready. If promoteMember call
	// fails with ErrNotLeader, forward the request to leader node via HTTP. If promoteMember call fails with error
	// other than ErrNotLeader, return the error.
	resp, err := s.promoteMember(ctx, id)
	if err == nil {
		return resp, nil
	}
	if err != ErrNotLeader {
		return resp, err
	}

	cctx, cancel := context.WithTimeout(ctx, s.Cfg.ReqTimeout())
	defer cancel()
	// forward to leader
	for cctx.Err() == nil {
		leader, err := s.waitLeader(cctx)
		if err != nil {
			return nil, err
		}
		for _, url := range leader.PeerURLs {
			resp, err := promoteMemberHTTP(cctx, url, id, s.peerRt)
			if err == nil {
				return resp, nil
			}
			// If member promotion failed, return early. Otherwise keep retry.
			if err == ErrLearnerNotReady || err == membership.ErrIDNotFound || err == membership.ErrMemberNotLearner {
				return nil, err
			}
		}
	}

	if cctx.Err() == context.DeadlineExceeded {
		return nil, ErrTimeout
	}
	return nil, ErrCanceled
}

// promoteMember checks whether the to-backend-promoted learner node is ready before sending the promote
// request to raft.
// The function returns ErrNotLeader if the local node is not raft leader (therefore does not have
// enough information to determine if the learner node is ready), returns ErrLearnerNotReady if the
// local node is leader (therefore has enough information) but decided the learner node is not ready
// to backend promoted.
func (s *EtcdServer) promoteMember(ctx context.Context, id uint64) ([]*membership.Member, error) {
	if err := s.checkMembershipOperationPermission(ctx); err != nil {
		return nil, err
	}

	// check if we can promote this learner.
	if err := s.mayPromoteMember(types.ID(id)); err != nil {
		return nil, err
	}

	// build the context for the promote confChange. mark IsLearner to false and IsPromote to true.
	promoteChangeContext := membership.ConfigChangeContext{
		Member: membership.Member{
			ID: types.ID(id),
		},
		IsPromote: true,
	}

	b, err := json.Marshal(promoteChangeContext)
	if err != nil {
		return nil, err
	}

	cc := raftpb.ConfChangeV1{
		Type:    raftpb.ConfChangeAddNode,
		NodeID:  id,
		Context: string(b),
	}

	return s.configure(ctx, cc)
}

func (s *EtcdServer) mayPromoteMember(id types.ID) error {
	lg := s.Logger()
	err := s.isLearnerReady(uint64(id))
	if err != nil {
		return err
	}

	if !s.Cfg.StrictReconfigCheck {
		return nil
	}
	if !s.cluster.IsReadyToPromoteMember(uint64(id)) {
		lg.Warn(
			"rejecting member promote request; not enough healthy members",
			zap.String("local-member-id", s.ID().String()),
			zap.String("requested-member-remove-id", id.String()),
			zap.Error(ErrNotEnoughStartedMembers),
		)
		return ErrNotEnoughStartedMembers
	}

	return nil
}
