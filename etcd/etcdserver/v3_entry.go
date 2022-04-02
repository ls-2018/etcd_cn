package etcdserver

import (
	"context"
	"time"

	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api/membership"
	"github.com/ls-2018/etcd_cn/etcd/mvcc"
	pb "github.com/ls-2018/etcd_cn/offical/etcdserverpb"
	"go.uber.org/zap"
)

// Apply 入口函数,负责处理内部的消息
func (a *applierV3backend) Apply(r *pb.InternalRaftRequest, shouldApplyV3 membership.ShouldApplyV3) *applyResult {
	ar := &applyResult{}
	defer func(start time.Time) {
		success := ar.err == nil || ar.err == mvcc.ErrCompacted
		if !success {
			warnOfFailedRequest(a.s.Logger(), start, &pb.InternalRaftStringer{Request: r}, ar.resp, ar.err)
		}
	}(time.Now())

	switch {
	case r.ClusterVersionSet != nil: //     3.5.x 实现
		// 设置集群版本
		a.s.applyV3Internal.ClusterVersionSet(r.ClusterVersionSet, shouldApplyV3)
		return nil
	case r.ClusterMemberAttrSet != nil:
		// 集群成员属性
		a.s.applyV3Internal.ClusterMemberAttrSet(r.ClusterMemberAttrSet, shouldApplyV3)
		return nil
	case r.DowngradeInfoSet != nil:
		// 成员降级
		a.s.applyV3Internal.DowngradeInfoSet(r.DowngradeInfoSet, shouldApplyV3)
		return nil
	}

	if !shouldApplyV3 {
		return nil
	}

	switch {
	case r.Range != nil:
		ar.resp, ar.err = a.s.applyV3.Range(context.TODO(), nil, r.Range) // ✅
	case r.Put != nil:
		ar.resp, ar.trace, ar.err = a.s.applyV3.Put(context.TODO(), nil, r.Put) // ✅
	case r.DeleteRange != nil:
		ar.resp, ar.err = a.s.applyV3.DeleteRange(nil, r.DeleteRange) // ✅
	case r.Txn != nil:
		ar.resp, ar.trace, ar.err = a.s.applyV3.Txn(context.TODO(), r.Txn)
	case r.Compaction != nil:
		ar.resp, ar.physc, ar.trace, ar.err = a.s.applyV3.Compaction(r.Compaction)
	case r.LeaseGrant != nil:
		ar.resp, ar.err = a.s.applyV3.LeaseGrant(r.LeaseGrant)
	case r.LeaseRevoke != nil:
		ar.resp, ar.err = a.s.applyV3.LeaseRevoke(r.LeaseRevoke)
	case r.LeaseCheckpoint != nil:
		ar.resp, ar.err = a.s.applyV3.LeaseCheckpoint(r.LeaseCheckpoint)
	case r.Alarm != nil:
		ar.resp, ar.err = a.s.applyV3.Alarm(r.Alarm) // ✅
	case r.Authenticate != nil:
		ar.resp, ar.err = a.s.applyV3.Authenticate(r.Authenticate)
	case r.AuthEnable != nil:
		ar.resp, ar.err = a.s.applyV3.AuthEnable()
	case r.AuthDisable != nil:
		ar.resp, ar.err = a.s.applyV3.AuthDisable()
	case r.AuthStatus != nil:
		ar.resp, ar.err = a.s.applyV3.AuthStatus()
	case r.AuthUserAdd != nil:
		ar.resp, ar.err = a.s.applyV3.UserAdd(r.AuthUserAdd)
	case r.AuthUserDelete != nil:
		ar.resp, ar.err = a.s.applyV3.UserDelete(r.AuthUserDelete)
	case r.AuthUserChangePassword != nil:
		ar.resp, ar.err = a.s.applyV3.UserChangePassword(r.AuthUserChangePassword)
	case r.AuthUserGrantRole != nil:
		ar.resp, ar.err = a.s.applyV3.UserGrantRole(r.AuthUserGrantRole)
	case r.AuthUserGet != nil:
		ar.resp, ar.err = a.s.applyV3.UserGet(r.AuthUserGet)
	case r.AuthUserRevokeRole != nil:
		ar.resp, ar.err = a.s.applyV3.UserRevokeRole(r.AuthUserRevokeRole)
	case r.AuthRoleAdd != nil:
		ar.resp, ar.err = a.s.applyV3.RoleAdd(r.AuthRoleAdd) // ✅
	case r.AuthRoleGrantPermission != nil:
		ar.resp, ar.err = a.s.applyV3.RoleGrantPermission(r.AuthRoleGrantPermission)
	case r.AuthRoleGet != nil:
		ar.resp, ar.err = a.s.applyV3.RoleGet(r.AuthRoleGet)
	case r.AuthRoleRevokePermission != nil:
		ar.resp, ar.err = a.s.applyV3.RoleRevokePermission(r.AuthRoleRevokePermission)
	case r.AuthRoleDelete != nil:
		ar.resp, ar.err = a.s.applyV3.RoleDelete(r.AuthRoleDelete)
	case r.AuthUserList != nil:
		ar.resp, ar.err = a.s.applyV3.UserList(r.AuthUserList)
	case r.AuthRoleList != nil:
		ar.resp, ar.err = a.s.applyV3.RoleList(r.AuthRoleList)
	default:
		a.s.lg.Panic("没有实现应用", zap.Stringer("raft-request", r))
	}
	return ar
}
