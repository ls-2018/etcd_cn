package etcdserverpb

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/ls-2018/etcd_cn/offical/api/v3/membershippb"
)

type xx struct {
	Key         string
	Value       string
	Lease       int64
	PrevKv      bool
	IgnoreValue bool
	IgnoreLease bool
}
type ASD struct {
	Header                   *RequestHeader
	ID                       uint64
	V2                       *Request
	Put                      *xx
	DeleteRange              *DeleteRangeRequest
	AuthRoleRevokePermission *AuthRoleRevokePermissionRequest
	AuthRoleGet              *AuthRoleGetRequest
	AuthRoleDelete           *AuthRoleDeleteRequest
	AuthUserList             *AuthUserListRequest
	AuthUserChangePassword   *AuthUserChangePasswordRequest
	AuthStatus               *AuthStatusRequest
	LeaseCheckpoint          *LeaseCheckpointRequest
	Alarm                    *AlarmRequest
	AuthDisable              *AuthDisableRequest
	LeaseRevoke              *LeaseRevokeRequest
	AuthEnable               *AuthEnableRequest
	AuthUserDelete           *AuthUserDeleteRequest
	Authenticate             *InternalAuthenticateRequest
	AuthUserGet              *AuthUserGetRequest
	AuthUserRevokeRole       *AuthUserRevokeRoleRequest
	LeaseGrant               *LeaseGrantRequest
	Compaction               *CompactionRequest
	AuthRoleList             *AuthRoleListRequest
	AuthRoleAdd              *AuthRoleAddRequest
	AuthUserGrantRole        *AuthUserGrantRoleRequest
	AuthUserAdd              *AuthUserAddRequest
	ClusterVersionSet        *membershippb.ClusterVersionSetRequest
	ClusterMemberAttrSet     *membershippb.ClusterMemberAttrSetRequest
	DowngradeInfoSet         *membershippb.DowngradeInfoSetRequest
}

func (m *InternalRaftRequest) Marshal() (dAtA []byte, err error) {
	_ = m.Unmarshal
	a := ASD{
		Put:                      nil,
		Header:                   m.Header,
		ID:                       m.ID,
		V2:                       m.V2,
		DeleteRange:              m.DeleteRange,
		AuthRoleRevokePermission: m.AuthRoleRevokePermission,
		AuthRoleGet:              m.AuthRoleGet,
		AuthRoleDelete:           m.AuthRoleDelete,
		AuthUserList:             m.AuthUserList,
		AuthUserChangePassword:   m.AuthUserChangePassword,
		AuthStatus:               m.AuthStatus,
		LeaseCheckpoint:          m.LeaseCheckpoint,
		Alarm:                    m.Alarm,
		AuthDisable:              m.AuthDisable,
		LeaseRevoke:              m.LeaseRevoke,
		AuthEnable:               m.AuthEnable,
		AuthUserDelete:           m.AuthUserDelete,
		Authenticate:             m.Authenticate,
		AuthUserGet:              m.AuthUserGet,
		AuthUserRevokeRole:       m.AuthUserRevokeRole,
		LeaseGrant:               m.LeaseGrant,
		Compaction:               m.Compaction,
		AuthRoleList:             m.AuthRoleList,
		AuthRoleAdd:              m.AuthRoleAdd,
		AuthUserGrantRole:        m.AuthUserGrantRole,
		AuthUserAdd:              m.AuthUserAdd,
		ClusterVersionSet:        m.ClusterVersionSet,
		ClusterMemberAttrSet:     m.ClusterMemberAttrSet,
		DowngradeInfoSet:         m.DowngradeInfoSet,
	}
	if m.Put != nil {
		a.Put = &xx{
			Key:         string(m.Put.Key),
			Value:       string(m.Put.Value),
			Lease:       m.Put.Lease,
			PrevKv:      m.Put.PrevKv,
			IgnoreValue: m.Put.IgnoreValue,
			IgnoreLease: m.Put.IgnoreLease,
		}
	}

	return json.Marshal(a)
}

func (m *InternalRaftRequest) Unmarshal(dAtA []byte) error {
	// a := `{"header":{"ID":7587861231285799685},"put":{"key":"YQ==","value":"Yg=="}}`
	//	b := `{"ID":7587861231285799684,"Method":"PUT","Path":"/0/version","Val":"3.5.0","Dir":false,"PrevValue":"","PrevIndex":0,"Expiration":0,"Wait":false,"Since":0,"Recursive":false,"Sorted":false,"Quorum":false,"Time":0,"Stream":false}`
	//	fmt.Println(json.Unmarshal([]byte(a), &etcdserverpb.InternalRaftRequest{}))   // 不能反序列化成功
	//	fmt.Println(json.Unmarshal([]byte(b), &etcdserverpb.InternalRaftRequest{}))

	if strings.Contains(string(dAtA), "Method") {
		return errors.New("特殊需求,不是使其反序列化成功")
	}
	a := ASD{}
	err := json.Unmarshal(dAtA, &a)
	if a.Put != nil {
		m.Put = &PutRequest{
			Key:         a.Put.Key,
			Value:       a.Put.Value,
			Lease:       a.Put.Lease,
			PrevKv:      a.Put.PrevKv,
			IgnoreValue: a.Put.IgnoreValue,
			IgnoreLease: a.Put.IgnoreLease,
		}
	}
	m.Header = a.Header
	m.ID = a.ID
	m.V2 = a.V2
	m.DeleteRange = a.DeleteRange
	m.AuthRoleRevokePermission = a.AuthRoleRevokePermission
	m.AuthRoleGet = a.AuthRoleGet
	m.AuthRoleDelete = a.AuthRoleDelete
	m.AuthUserList = a.AuthUserList
	m.AuthUserChangePassword = a.AuthUserChangePassword
	m.AuthStatus = a.AuthStatus
	m.LeaseCheckpoint = a.LeaseCheckpoint
	m.Alarm = a.Alarm
	m.AuthDisable = a.AuthDisable
	m.LeaseRevoke = a.LeaseRevoke
	m.AuthEnable = a.AuthEnable
	m.AuthUserDelete = a.AuthUserDelete
	m.Authenticate = a.Authenticate
	m.AuthUserGet = a.AuthUserGet
	m.AuthUserRevokeRole = a.AuthUserRevokeRole
	m.LeaseGrant = a.LeaseGrant
	m.Compaction = a.Compaction
	m.AuthRoleList = a.AuthRoleList
	m.AuthRoleAdd = a.AuthRoleAdd
	m.AuthUserGrantRole = a.AuthUserGrantRole
	m.AuthUserAdd = a.AuthUserAdd
	m.ClusterVersionSet = a.ClusterVersionSet
	m.ClusterMemberAttrSet = a.ClusterMemberAttrSet
	m.DowngradeInfoSet = a.DowngradeInfoSet
	return err
}

// 不能更改
