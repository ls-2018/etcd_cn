// Copyright 2016 The etcd Authors
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
	"sync"

	"github.com/ls-2018/etcd_cn/etcd/auth"
	"github.com/ls-2018/etcd_cn/etcd/etcdserver/api/membership"
	"github.com/ls-2018/etcd_cn/etcd/lease"
	"github.com/ls-2018/etcd_cn/etcd/mvcc"
	pb "github.com/ls-2018/etcd_cn/offical/etcdserverpb"
	"github.com/ls-2018/etcd_cn/pkg/traceutil"
)

type authApplierV3 struct {
	applierV3                // applierV3backend
	as        auth.AuthStore // 内循环时,提供认证token
	lessor    lease.Lessor   // 租约管理者
	// mu serializes Apply so that user isn't corrupted and so that
	// serialized requests don't leak data from TOCTOU errors
	mu       sync.Mutex
	authInfo auth.AuthInfo
}

func newAuthApplierV3(as auth.AuthStore, base applierV3, lessor lease.Lessor) *authApplierV3 {
	return &authApplierV3{applierV3: base, as: as, lessor: lessor}
}

func (aa *authApplierV3) Apply(r *pb.InternalRaftRequest, shouldApplyV3 membership.ShouldApplyV3) *applyResult {
	aa.mu.Lock()
	defer aa.mu.Unlock()
	if r.Header != nil {
		// 当internalRaftRequest没有header时,向后兼容3.0之前的版本
		aa.authInfo.Username = r.Header.Username
		aa.authInfo.Revision = r.Header.AuthRevision
	}
	if needAdminPermission(r) {
		if err := aa.as.IsAdminPermitted(&aa.authInfo); err != nil {
			aa.authInfo.Username = ""
			aa.authInfo.Revision = 0
			return &applyResult{err: err}
		}
	}
	ret := aa.applierV3.Apply(r, shouldApplyV3)
	aa.authInfo.Username = ""
	aa.authInfo.Revision = 0
	return ret
}

func (aa *authApplierV3) Put(ctx context.Context, txn mvcc.TxnWrite, r *pb.PutRequest) (*pb.PutResponse, *traceutil.Trace, error) {
	if err := aa.as.IsPutPermitted(&aa.authInfo, []byte(r.Key)); err != nil {
		return nil, nil, err
	}

	if err := aa.checkLeasePuts(lease.LeaseID(r.Lease)); err != nil {
		// The specified lease is already attached with a key that cannot
		// backend written by this user. It means the user cannot revoke the
		// lease so attaching the lease to the newly written key should
		// backend forbidden.
		return nil, nil, err
	}

	if r.PrevKv {
		err := aa.as.IsRangePermitted(&aa.authInfo, []byte(r.Key), nil)
		if err != nil {
			return nil, nil, err
		}
	}
	return aa.applierV3.Put(ctx, txn, r)
}

func (aa *authApplierV3) Range(ctx context.Context, txn mvcc.TxnRead, r *pb.RangeRequest) (*pb.RangeResponse, error) {
	if err := aa.as.IsRangePermitted(&aa.authInfo, []byte(r.Key), []byte(r.RangeEnd)); err != nil {
		return nil, err
	}
	return aa.applierV3.Range(ctx, txn, r)
}

func (aa *authApplierV3) DeleteRange(txn mvcc.TxnWrite, r *pb.DeleteRangeRequest) (*pb.DeleteRangeResponse, error) {
	if err := aa.as.IsDeleteRangePermitted(&aa.authInfo, []byte(r.Key), []byte(r.RangeEnd)); err != nil {
		return nil, err
	}
	if r.PrevKv {
		err := aa.as.IsRangePermitted(&aa.authInfo, []byte(r.Key), []byte(r.RangeEnd)) // {a,b true}
		if err != nil {
			return nil, err
		}
	}

	return aa.applierV3.DeleteRange(txn, r)
}

func checkTxnReqsPermission(as auth.AuthStore, ai *auth.AuthInfo, reqs []*pb.RequestOp) error {
	for _, requ := range reqs {
		if requ.RequestOp_RequestRange != nil {
			tv := requ.RequestOp_RequestRange
			if tv.RequestRange == nil {
				continue
			}
			if err := as.IsRangePermitted(ai, []byte(tv.RequestRange.Key), []byte(tv.RequestRange.RangeEnd)); err != nil {
				return err
			}

		}
		if requ.RequestOp_RequestPut != nil {
			tv := requ.RequestOp_RequestPut
			if tv.RequestPut == nil {
				continue
			}

			if err := as.IsPutPermitted(ai, []byte(tv.RequestPut.Key)); err != nil {
				return err
			}

		}
		if requ.RequestOp_RequestDeleteRange != nil {
			tv := requ.RequestOp_RequestDeleteRange
			if tv.RequestDeleteRange == nil {
				continue
			}

			if tv.RequestDeleteRange.PrevKv {
				err := as.IsRangePermitted(ai, []byte(tv.RequestDeleteRange.Key), []byte(tv.RequestDeleteRange.RangeEnd))
				if err != nil {
					return err
				}
			}

			err := as.IsDeleteRangePermitted(ai, []byte(tv.RequestDeleteRange.Key), []byte(tv.RequestDeleteRange.RangeEnd))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkTxnAuth(as auth.AuthStore, ai *auth.AuthInfo, rt *pb.TxnRequest) error {
	for _, c := range rt.Compare {
		if err := as.IsRangePermitted(ai, []byte(c.Key), []byte(c.RangeEnd)); err != nil {
			return err
		}
	}
	if err := checkTxnReqsPermission(as, ai, rt.Success); err != nil {
		return err
	}
	return checkTxnReqsPermission(as, ai, rt.Failure)
}

func (aa *authApplierV3) Txn(ctx context.Context, rt *pb.TxnRequest) (*pb.TxnResponse, *traceutil.Trace, error) {
	if err := checkTxnAuth(aa.as, &aa.authInfo, rt); err != nil {
		return nil, nil, err
	}
	return aa.applierV3.Txn(ctx, rt)
}

func (aa *authApplierV3) LeaseRevoke(lc *pb.LeaseRevokeRequest) (*pb.LeaseRevokeResponse, error) {
	if err := aa.checkLeasePuts(lease.LeaseID(lc.ID)); err != nil { // 检查租约是否存在
		return nil, err
	}
	return aa.applierV3.LeaseRevoke(lc)
}

// 检查租约更新的key是否有权限操作
func (aa *authApplierV3) checkLeasePuts(leaseID lease.LeaseID) error {
	lease := aa.lessor.Lookup(leaseID)
	if lease != nil {
		for _, key := range lease.Keys() {
			if err := aa.as.IsPutPermitted(&aa.authInfo, []byte(key)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (aa *authApplierV3) UserGet(r *pb.AuthUserGetRequest) (*pb.AuthUserGetResponse, error) {
	err := aa.as.IsAdminPermitted(&aa.authInfo)
	if err != nil && r.Name != aa.authInfo.Username {
		aa.authInfo.Username = ""
		aa.authInfo.Revision = 0
		return &pb.AuthUserGetResponse{}, err
	}

	return aa.applierV3.UserGet(r)
}

func (aa *authApplierV3) RoleGet(r *pb.AuthRoleGetRequest) (*pb.AuthRoleGetResponse, error) {
	err := aa.as.IsAdminPermitted(&aa.authInfo)
	if err != nil && !aa.as.UserHasRole(aa.authInfo.Username, r.Role) {
		aa.authInfo.Username = ""
		aa.authInfo.Revision = 0
		return &pb.AuthRoleGetResponse{}, err
	}

	return aa.applierV3.RoleGet(r)
}

func needAdminPermission(r *pb.InternalRaftRequest) bool {
	switch {
	case r.AuthEnable != nil:
		return true
	case r.AuthDisable != nil:
		return true
	case r.AuthStatus != nil:
		return true
	case r.AuthUserAdd != nil:
		return true
	case r.AuthUserDelete != nil:
		return true
	case r.AuthUserChangePassword != nil:
		return true
	case r.AuthUserGrantRole != nil:
		return true
	case r.AuthUserRevokeRole != nil:
		return true
	case r.AuthRoleAdd != nil:
		return true
	case r.AuthRoleGrantPermission != nil:
		return true
	case r.AuthRoleRevokePermission != nil:
		return true
	case r.AuthRoleDelete != nil:
		return true
	case r.AuthUserList != nil:
		return true
	case r.AuthRoleList != nil:
		return true
	default:
		return false
	}
}