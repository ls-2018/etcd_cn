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

package etcdserver

import (
	"context"
	"encoding/binary"
	"time"

	"github.com/ls-2018/etcd_cn/raft"

	"github.com/ls-2018/etcd_cn/etcd_backend/auth"
	"github.com/ls-2018/etcd_cn/etcd_backend/etcdserver/api/membership"
	"github.com/ls-2018/etcd_cn/etcd_backend/mvcc"
	pb "github.com/ls-2018/etcd_cn/offical/etcdserverpb"
	"github.com/ls-2018/etcd_cn/pkg/traceutil"
	"go.etcd.io/etcd/api/v3/membershippb"

	"github.com/gogo/protobuf/proto"
	"go.uber.org/zap"
)

const (
	// In the health case, there might backend a small gap (10s of entries) between
	// the applied index and committed index.
	// However, if the committed entries are very heavy to apply, the gap might grow.
	// We should stop accepting new proposals if the gap growing to a certain point.
	maxGapBetweenApplyAndCommitIndex = 5000
	readIndexRetryTime               = 500 * time.Millisecond
)

type Authenticator interface {
	AuthEnable(ctx context.Context, r *pb.AuthEnableRequest) (*pb.AuthEnableResponse, error)
	AuthDisable(ctx context.Context, r *pb.AuthDisableRequest) (*pb.AuthDisableResponse, error)
	AuthStatus(ctx context.Context, r *pb.AuthStatusRequest) (*pb.AuthStatusResponse, error)
	Authenticate(ctx context.Context, r *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error)
	UserAdd(ctx context.Context, r *pb.AuthUserAddRequest) (*pb.AuthUserAddResponse, error)
	UserDelete(ctx context.Context, r *pb.AuthUserDeleteRequest) (*pb.AuthUserDeleteResponse, error)
	UserChangePassword(ctx context.Context, r *pb.AuthUserChangePasswordRequest) (*pb.AuthUserChangePasswordResponse, error)
	UserGrantRole(ctx context.Context, r *pb.AuthUserGrantRoleRequest) (*pb.AuthUserGrantRoleResponse, error)
	UserGet(ctx context.Context, r *pb.AuthUserGetRequest) (*pb.AuthUserGetResponse, error)
	UserRevokeRole(ctx context.Context, r *pb.AuthUserRevokeRoleRequest) (*pb.AuthUserRevokeRoleResponse, error)
	RoleAdd(ctx context.Context, r *pb.AuthRoleAddRequest) (*pb.AuthRoleAddResponse, error)
	RoleGrantPermission(ctx context.Context, r *pb.AuthRoleGrantPermissionRequest) (*pb.AuthRoleGrantPermissionResponse, error)
	RoleGet(ctx context.Context, r *pb.AuthRoleGetRequest) (*pb.AuthRoleGetResponse, error)
	RoleRevokePermission(ctx context.Context, r *pb.AuthRoleRevokePermissionRequest) (*pb.AuthRoleRevokePermissionResponse, error)
	RoleDelete(ctx context.Context, r *pb.AuthRoleDeleteRequest) (*pb.AuthRoleDeleteResponse, error)
	UserList(ctx context.Context, r *pb.AuthUserListRequest) (*pb.AuthUserListResponse, error)
	RoleList(ctx context.Context, r *pb.AuthRoleListRequest) (*pb.AuthRoleListResponse, error)
}

func isTxnSerializable(r *pb.TxnRequest) bool {
	for _, u := range r.Success {
		if r := u.GetRequestRange(); r == nil || !r.Serializable {
			return false
		}
	}
	for _, u := range r.Failure {
		if r := u.GetRequestRange(); r == nil || !r.Serializable {
			return false
		}
	}
	return true
}

func isTxnReadonly(r *pb.TxnRequest) bool {
	for _, u := range r.Success {
		if r := u.GetRequestRange(); r == nil {
			return false
		}
	}
	for _, u := range r.Failure {
		if r := u.GetRequestRange(); r == nil {
			return false
		}
	}
	return true
}

func (s *EtcdServer) waitLeader(ctx context.Context) (*membership.Member, error) {
	leader := s.cluster.Member(s.Leader())
	for leader == nil {
		// wait an election
		dur := time.Duration(s.Cfg.ElectionTicks) * time.Duration(s.Cfg.TickMs) * time.Millisecond
		select {
		case <-time.After(dur):
			leader = s.cluster.Member(s.Leader())
		case <-s.stopping:
			return nil, ErrStopped
		case <-ctx.Done():
			return nil, ErrNoLeader
		}
	}
	if leader == nil || len(leader.PeerURLs) == 0 {
		return nil, ErrNoLeader
	}
	return leader, nil
}

func (s *EtcdServer) Alarm(ctx context.Context, r *pb.AlarmRequest) (*pb.AlarmResponse, error) {
	resp, err := s.raftRequestOnce(ctx, pb.InternalRaftRequest{Alarm: r})
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AlarmResponse), nil
}

func (s *EtcdServer) raftRequestOnce(ctx context.Context, r pb.InternalRaftRequest) (proto.Message, error) {
	result, err := s.processInternalRaftRequestOnce(ctx, r)
	if err != nil {
		return nil, err
	}
	if result.err != nil {
		return nil, result.err
	}
	// startTime
	if startTime, ok := ctx.Value(traceutil.StartTimeKey).(time.Time); ok && result.trace != nil {
		applyStart := result.trace.GetStartTime()
		result.trace.SetStartTime(startTime)
		result.trace.InsertStep(0, applyStart, "处理raft请求")
	}
	return result.resp, nil
}

func (s *EtcdServer) raftRequest(ctx context.Context, r pb.InternalRaftRequest) (proto.Message, error) {
	return s.raftRequestOnce(ctx, r)
}

// doSerialize handles the auth logic, with permissions checked by "chk", for a serialized request "get". Returns a non-nil error on authentication failure.
func (s *EtcdServer) doSerialize(ctx context.Context, chk func(*auth.AuthInfo) error, get func()) error {
	trace := traceutil.Get(ctx)
	ai, err := s.AuthInfoFromCtx(ctx)
	if err != nil {
		return err
	}
	if ai == nil {
		// chk expects non-nil AuthInfo; use empty credentials
		ai = &auth.AuthInfo{}
	}
	if err = chk(ai); err != nil {
		return err
	}
	trace.Step("get authentication metadata")
	// fetch response for serialized request
	get()
	// check for stale token revision in case the auth store was updated while
	// the request has been handled.
	if ai.Revision != 0 && ai.Revision != s.authStore.Revision() {
		return auth.ErrAuthOldRevision
	}
	return nil
}

// 当客户端提交一条数据变更请求时
func (s *EtcdServer) processInternalRaftRequestOnce(ctx context.Context, r pb.InternalRaftRequest) (*applyResult, error) {
	// 判断已提交未apply的记录是否超过限制
	ai := s.getAppliedIndex()
	ci := s.getCommittedIndex()
	if ci > ai+maxGapBetweenApplyAndCommitIndex {
		return nil, ErrTooManyRequests
	}

	r.Header = &pb.RequestHeader{
		ID: s.reqIDGen.Next(), // 生成一个requestID
	}

	// 检查authinfo是否不是InternalAuthenticateRequest
	if r.Authenticate == nil {
		authInfo, err := s.AuthInfoFromCtx(ctx)
		if err != nil {
			return nil, err
		}
		if authInfo != nil {
			r.Header.Username = authInfo.Username
			r.Header.AuthRevision = authInfo.Revision
		}
	}
	// 反序列化请求数据

	data, err := r.Marshal()
	if err != nil {
		return nil, err
	}

	if len(data) > int(s.Cfg.MaxRequestBytes) {
		return nil, ErrRequestTooLarge
	}

	id := r.ID // 0
	if id == 0 {
		id = r.Header.ID
	}
	ch := s.w.Register(id) // 注册一个channel,等待处理完成

	cctx, cancel := context.WithTimeout(ctx, s.Cfg.ReqTimeout()) // 设置请求超时
	//cctx, cancel := context.WithTimeout(ctx, time.Second*1000) // 设置请求超时
	defer cancel()

	start := time.Now()
	_ = s.applyEntryNormal
	err = s.r.Propose(cctx, data) // 调用raft模块的Propose处理请求,存入到了待发送队列
	if err != nil {
		s.w.Trigger(id, nil)
		return nil, err
	}

	select {
	// 等待收到apply结果返回给客户端
	case x := <-ch:
		return x.(*applyResult), nil
	case <-cctx.Done():
		s.w.Trigger(id, nil)
		return nil, s.parseProposeCtxErr(cctx.Err(), start)
	case <-s.done:
		return nil, ErrStopped
	}
}

// Watchable returns a watchable interface attached to the etcdserver.
func (s *EtcdServer) Watchable() mvcc.WatchableKV { return s.KV() }

func isStopped(err error) bool {
	return err == raft.ErrStopped || err == ErrStopped
}

func uint64ToBigEndianBytes(number uint64) []byte {
	byteResult := make([]byte, 8)
	binary.BigEndian.PutUint64(byteResult, number)
	return byteResult
}

func (s *EtcdServer) sendReadIndex(requestIndex uint64) error {
	ctxToSend := uint64ToBigEndianBytes(requestIndex)

	cctx, cancel := context.WithTimeout(context.Background(), s.Cfg.ReqTimeout())
	err := s.r.ReadIndex(cctx, ctxToSend)
	cancel()
	if err == raft.ErrStopped {
		return err
	}
	if err != nil {
		lg := s.Logger()
		lg.Warn("未能从Raft获取读取索引", zap.Error(err))
		return err
	}
	return nil
}

func (s *EtcdServer) linearizableReadNotify(ctx context.Context) error {
	s.readMu.RLock()
	nc := s.readNotifier
	s.readMu.RUnlock()

	//  linearizable loop 没有准备好，就发送此信号
	select {
	case s.readwaitc <- struct{}{}:
	default:
	}

	// wait for read state notification
	select {
	case <-nc.c:
		return nc.err
	case <-ctx.Done():
		return ctx.Err()
	case <-s.done:
		return ErrStopped
	}
}

// AuthInfoFromCtx 获取认证信息
func (s *EtcdServer) AuthInfoFromCtx(ctx context.Context) (*auth.AuthInfo, error) {
	authInfo, err := s.AuthStore().AuthInfoFromCtx(ctx)
	if authInfo != nil || err != nil {
		return authInfo, err
	}
	if !s.Cfg.ClientCertAuthEnabled { // 是否验证客户端证书
		return nil, nil
	}
	authInfo = s.AuthStore().AuthInfoFromTLS(ctx)
	return authInfo, nil
}

func (s *EtcdServer) Downgrade(ctx context.Context, r *pb.DowngradeRequest) (*pb.DowngradeResponse, error) {
	switch r.Action {
	case pb.DowngradeRequest_VALIDATE:
		return s.downgradeValidate(ctx, r.Version)
	case pb.DowngradeRequest_ENABLE:
		return s.downgradeEnable(ctx, r)
	case pb.DowngradeRequest_CANCEL:
		return s.downgradeCancel(ctx)
	default:
		return nil, ErrUnknownMethod
	}
}

func (s *EtcdServer) downgradeValidate(ctx context.Context, v string) (*pb.DowngradeResponse, error) {
	resp := &pb.DowngradeResponse{}

	targetVersion, err := convertToClusterVersion(v)
	if err != nil {
		return nil, err
	}

	// gets leaders commit index and wait for local store to finish applying that index
	// to avoid using stale downgrade information
	err = s.linearizableReadNotify(ctx)
	if err != nil {
		return nil, err
	}

	cv := s.ClusterVersion()
	if cv == nil {
		return nil, ErrClusterVersionUnavailable
	}
	resp.Version = cv.String()

	allowedTargetVersion := membership.AllowedDowngradeVersion(cv)
	if !targetVersion.Equal(*allowedTargetVersion) {
		return nil, ErrInvalidDowngradeTargetVersion
	}

	downgradeInfo := s.cluster.DowngradeInfo()
	if downgradeInfo.Enabled {
		// Todo: return the downgrade status along with the error msg
		return nil, ErrDowngradeInProcess
	}
	return resp, nil
}

func (s *EtcdServer) downgradeEnable(ctx context.Context, r *pb.DowngradeRequest) (*pb.DowngradeResponse, error) {
	// validate downgrade capability before starting downgrade
	v := r.Version
	lg := s.Logger()
	if resp, err := s.downgradeValidate(ctx, v); err != nil {
		lg.Warn("reject downgrade request", zap.Error(err))
		return resp, err
	}
	targetVersion, err := convertToClusterVersion(v)
	if err != nil {
		lg.Warn("reject downgrade request", zap.Error(err))
		return nil, err
	}

	raftRequest := membershippb.DowngradeInfoSetRequest{Enabled: true, Ver: targetVersion.String()}
	_, err = s.raftRequest(ctx, pb.InternalRaftRequest{DowngradeInfoSet: &raftRequest})
	if err != nil {
		lg.Warn("reject downgrade request", zap.Error(err))
		return nil, err
	}
	resp := pb.DowngradeResponse{Version: s.ClusterVersion().String()}
	return &resp, nil
}

func (s *EtcdServer) downgradeCancel(ctx context.Context) (*pb.DowngradeResponse, error) {
	// gets leaders commit index and wait for local store to finish applying that index
	// to avoid using stale downgrade information
	if err := s.linearizableReadNotify(ctx); err != nil {
		return nil, err
	}

	downgradeInfo := s.cluster.DowngradeInfo()
	if !downgradeInfo.Enabled {
		return nil, ErrNoInflightDowngrade
	}

	raftRequest := membershippb.DowngradeInfoSetRequest{Enabled: false}
	_, err := s.raftRequest(ctx, pb.InternalRaftRequest{DowngradeInfoSet: &raftRequest})
	if err != nil {
		return nil, err
	}
	resp := pb.DowngradeResponse{Version: s.ClusterVersion().String()}
	return &resp, nil
}
