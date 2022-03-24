package etcdserver

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/ls-2018/etcd_cn/etcd_backend/etcdserver/api/membership"
	"github.com/ls-2018/etcd_cn/pkg/traceutil"
	"github.com/ls-2018/etcd_cn/raft/raftpb"
	"go.uber.org/zap"

	"github.com/ls-2018/etcd_cn/client_sdk/pkg/types"
)

// LinearizableReadNotify 一致性读
func (s *EtcdServer) LinearizableReadNotify(ctx context.Context) error {
	return s.linearizableReadNotify(ctx)
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

func (s *EtcdServer) linearizableReadLoop() {
	for {
		requestId := s.reqIDGen.Next()
		leaderChangedNotifier := s.LeaderChangedNotify()
		select {
		case <-leaderChangedNotifier:
			continue
		case <-s.readwaitc:
		case <-s.stopping:
			return
		}

		// as a single loop is can unlock multiple reads, it is not very useful
		// to propagate the trace from Txn or Range.
		trace := traceutil.New("linearizableReadLoop", s.Logger())

		nextnr := newNotifier()
		s.readMu.Lock()
		nr := s.readNotifier
		s.readNotifier = nextnr
		s.readMu.Unlock()

		confirmedIndex, err := s.requestCurrentIndex(leaderChangedNotifier, requestId)
		if isStopped(err) {
			return
		}
		if err != nil {
			nr.notify(err)
			continue
		}

		trace.Step("read index received")

		trace.AddField(traceutil.Field{Key: "readStateIndex", Value: confirmedIndex})

		appliedIndex := s.getAppliedIndex()
		trace.AddField(traceutil.Field{Key: "appliedIndex", Value: strconv.FormatUint(appliedIndex, 10)})

		if appliedIndex < confirmedIndex {
			select {
			case <-s.applyWait.Wait(confirmedIndex):
			case <-s.stopping:
				return
			}
		}
		// unblock all l-reads requested at indices before confirmedIndex
		nr.notify(nil)
		trace.Step("applied index is now lower than readState.Index")

	}
}

func (s *EtcdServer) requestCurrentIndex(leaderChangedNotifier <-chan struct{}, requestId uint64) (uint64, error) {
	err := s.sendReadIndex(requestId)
	if err != nil {
		return 0, err
	}

	lg := s.Logger()
	errorTimer := time.NewTimer(s.Cfg.ReqTimeout())
	defer errorTimer.Stop()
	retryTimer := time.NewTimer(readIndexRetryTime)
	defer retryTimer.Stop()

	firstCommitInTermNotifier := s.FirstCommitInTermNotify()

	for {
		select {
		case rs := <-s.r.readStateC:
			requestIdBytes := uint64ToBigEndianBytes(requestId)
			gotOwnResponse := bytes.Equal(rs.RequestCtx, requestIdBytes)
			if !gotOwnResponse {
				// a previous request might time out. now we should ignore the response of it and
				// continue waiting for the response of the current requests.
				responseId := uint64(0)
				if len(rs.RequestCtx) == 8 {
					responseId = binary.BigEndian.Uint64(rs.RequestCtx)
				}
				lg.Warn(
					"ignored out-of-date read index response; local node read indexes queueing up and waiting to backend in sync with leader",
					zap.Uint64("sent-request-id", requestId),
					zap.Uint64("received-request-id", responseId),
				)
				continue
			}
			return rs.Index, nil
		case <-leaderChangedNotifier:
			// return a retryable error.
			return 0, ErrLeaderChanged
		case <-firstCommitInTermNotifier:
			firstCommitInTermNotifier = s.FirstCommitInTermNotify()
			lg.Info("first commit in current term: resending ReadIndex request")
			err := s.sendReadIndex(requestId)
			if err != nil {
				return 0, err
			}
			retryTimer.Reset(readIndexRetryTime)
			continue
		case <-retryTimer.C:
			lg.Warn(
				"waiting for ReadIndex response took too long, retrying",
				zap.Uint64("sent-request-id", requestId),
				zap.Duration("retry-timeout", readIndexRetryTime),
			)
			err := s.sendReadIndex(requestId)
			if err != nil {
				return 0, err
			}
			retryTimer.Reset(readIndexRetryTime)
			continue
		case <-errorTimer.C:
			lg.Warn(
				"timed out waiting for read index response (local node might have slow network)",
				zap.Duration("timeout", s.Cfg.ReqTimeout()),
			)
			return 0, ErrTimeout
		case <-s.stopping:
			return 0, ErrStopped
		}
	}
}
