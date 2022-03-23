package etcdserverpb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	membershippb "go.etcd.io/etcd/api/v3/membershippb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

var (
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RequestHeader struct {
	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// username is a username that is associated with an auth token of gRPC connection
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// auth_revision is a revision number of auth.authStore. It is not related to mvcc
	AuthRevision         uint64   `protobuf:"varint,3,opt,name=auth_revision,json=authRevision,proto3" json:"auth_revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4c9a9be0cfca103, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}

func (m *RequestHeader) XXX_Size() int {
	return m.Size()
}

func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

// An InternalRaftRequest is the union of all requests which can be
// sent via raft.
type InternalRaftRequest struct {
	Header                   *RequestHeader                            `protobuf:"bytes,100,opt,name=header,proto3" json:"header,omitempty"`
	ID                       uint64                                    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	V2                       *Request                                  `protobuf:"bytes,2,opt,name=v2,proto3" json:"v2,omitempty"`
	Range                    *RangeRequest                             `protobuf:"bytes,3,opt,name=range,proto3" json:"range,omitempty"`
	Put                      *PutRequest                               `protobuf:"bytes,4,opt,name=put,proto3" json:"put,omitempty"`
	DeleteRange              *DeleteRangeRequest                       `protobuf:"bytes,5,opt,name=delete_range,json=deleteRange,proto3" json:"delete_range,omitempty"`
	Txn                      *TxnRequest                               `protobuf:"bytes,6,opt,name=txn,proto3" json:"txn,omitempty"`
	Compaction               *CompactionRequest                        `protobuf:"bytes,7,opt,name=compaction,proto3" json:"compaction,omitempty"`
	LeaseGrant               *LeaseGrantRequest                        `protobuf:"bytes,8,opt,name=lease_grant,json=leaseGrant,proto3" json:"lease_grant,omitempty"`
	LeaseRevoke              *LeaseRevokeRequest                       `protobuf:"bytes,9,opt,name=lease_revoke,json=leaseRevoke,proto3" json:"lease_revoke,omitempty"`
	Alarm                    *AlarmRequest                             `protobuf:"bytes,10,opt,name=alarm,proto3" json:"alarm,omitempty"`
	LeaseCheckpoint          *LeaseCheckpointRequest                   `protobuf:"bytes,11,opt,name=lease_checkpoint,json=leaseCheckpoint,proto3" json:"lease_checkpoint,omitempty"`
	AuthEnable               *AuthEnableRequest                        `protobuf:"bytes,1000,opt,name=auth_enable,json=authEnable,proto3" json:"auth_enable,omitempty"`
	AuthDisable              *AuthDisableRequest                       `protobuf:"bytes,1011,opt,name=auth_disable,json=authDisable,proto3" json:"auth_disable,omitempty"`
	AuthStatus               *AuthStatusRequest                        `protobuf:"bytes,1013,opt,name=auth_status,json=authStatus,proto3" json:"auth_status,omitempty"`
	Authenticate             *InternalAuthenticateRequest              `protobuf:"bytes,1012,opt,name=authenticate,proto3" json:"authenticate,omitempty"`
	AuthUserAdd              *AuthUserAddRequest                       `protobuf:"bytes,1100,opt,name=auth_user_add,json=authUserAdd,proto3" json:"auth_user_add,omitempty"`
	AuthUserDelete           *AuthUserDeleteRequest                    `protobuf:"bytes,1101,opt,name=auth_user_delete,json=authUserDelete,proto3" json:"auth_user_delete,omitempty"`
	AuthUserGet              *AuthUserGetRequest                       `protobuf:"bytes,1102,opt,name=auth_user_get,json=authUserGet,proto3" json:"auth_user_get,omitempty"`
	AuthUserChangePassword   *AuthUserChangePasswordRequest            `protobuf:"bytes,1103,opt,name=auth_user_change_password,json=authUserChangePassword,proto3" json:"auth_user_change_password,omitempty"`
	AuthUserGrantRole        *AuthUserGrantRoleRequest                 `protobuf:"bytes,1104,opt,name=auth_user_grant_role,json=authUserGrantRole,proto3" json:"auth_user_grant_role,omitempty"`
	AuthUserRevokeRole       *AuthUserRevokeRoleRequest                `protobuf:"bytes,1105,opt,name=auth_user_revoke_role,json=authUserRevokeRole,proto3" json:"auth_user_revoke_role,omitempty"`
	AuthUserList             *AuthUserListRequest                      `protobuf:"bytes,1106,opt,name=auth_user_list,json=authUserList,proto3" json:"auth_user_list,omitempty"`
	AuthRoleList             *AuthRoleListRequest                      `protobuf:"bytes,1107,opt,name=auth_role_list,json=authRoleList,proto3" json:"auth_role_list,omitempty"`
	AuthRoleAdd              *AuthRoleAddRequest                       `protobuf:"bytes,1200,opt,name=auth_role_add,json=authRoleAdd,proto3" json:"auth_role_add,omitempty"`
	AuthRoleDelete           *AuthRoleDeleteRequest                    `protobuf:"bytes,1201,opt,name=auth_role_delete,json=authRoleDelete,proto3" json:"auth_role_delete,omitempty"`
	AuthRoleGet              *AuthRoleGetRequest                       `protobuf:"bytes,1202,opt,name=auth_role_get,json=authRoleGet,proto3" json:"auth_role_get,omitempty"`
	AuthRoleGrantPermission  *AuthRoleGrantPermissionRequest           `protobuf:"bytes,1203,opt,name=auth_role_grant_permission,json=authRoleGrantPermission,proto3" json:"auth_role_grant_permission,omitempty"`
	AuthRoleRevokePermission *AuthRoleRevokePermissionRequest          `protobuf:"bytes,1204,opt,name=auth_role_revoke_permission,json=authRoleRevokePermission,proto3" json:"auth_role_revoke_permission,omitempty"`
	ClusterVersionSet        *membershippb.ClusterVersionSetRequest    `protobuf:"bytes,1300,opt,name=cluster_version_set,json=clusterVersionSet,proto3" json:"cluster_version_set,omitempty"`
	ClusterMemberAttrSet     *membershippb.ClusterMemberAttrSetRequest `protobuf:"bytes,1301,opt,name=cluster_member_attr_set,json=clusterMemberAttrSet,proto3" json:"cluster_member_attr_set,omitempty"`
	DowngradeInfoSet         *membershippb.DowngradeInfoSetRequest     `protobuf:"bytes,1302,opt,name=downgrade_info_set,json=downgradeInfoSet,proto3" json:"downgrade_info_set,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                                  `json:"-"`
	XXX_unrecognized         []byte                                    `json:"-"`
	XXX_sizecache            int32                                     `json:"-"`
}

func (m *InternalRaftRequest) Reset()         { *m = InternalRaftRequest{} }
func (m *InternalRaftRequest) String() string { return proto.CompactTextString(m) }
func (*InternalRaftRequest) ProtoMessage()    {}
func (*InternalRaftRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4c9a9be0cfca103, []int{1}
}

func (m *InternalRaftRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *InternalRaftRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InternalRaftRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *InternalRaftRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalRaftRequest.Merge(m, src)
}

func (m *InternalRaftRequest) XXX_Size() int {
	return m.Size()
}

func (m *InternalRaftRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalRaftRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InternalRaftRequest proto.InternalMessageInfo

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4c9a9be0cfca103, []int{2}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}

func (m *EmptyResponse) XXX_Size() int {
	return m.Size()
}

func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

// What is the difference between AuthenticateRequest (defined in rpc.proto) and InternalAuthenticateRequest?
// InternalAuthenticateRequest has a member that is filled by etcdserver and shouldn't be user-facing.
// For avoiding misusage the field, we have an internal version of AuthenticateRequest.
type InternalAuthenticateRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// simple_token is generated in API layer (etcdserver/v3_server.go)
	SimpleToken          string   `protobuf:"bytes,3,opt,name=simple_token,json=simpleToken,proto3" json:"simple_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalAuthenticateRequest) Reset()         { *m = InternalAuthenticateRequest{} }
func (m *InternalAuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*InternalAuthenticateRequest) ProtoMessage()    {}
func (*InternalAuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4c9a9be0cfca103, []int{3}
}

func (m *InternalAuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *InternalAuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InternalAuthenticateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *InternalAuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalAuthenticateRequest.Merge(m, src)
}

func (m *InternalAuthenticateRequest) XXX_Size() int {
	return m.Size()
}

func (m *InternalAuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalAuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InternalAuthenticateRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RequestHeader)(nil), "etcdserverpb.RequestHeader")
	proto.RegisterType((*InternalRaftRequest)(nil), "etcdserverpb.InternalRaftRequest")
	proto.RegisterType((*EmptyResponse)(nil), "etcdserverpb.EmptyResponse")
	proto.RegisterType((*InternalAuthenticateRequest)(nil), "etcdserverpb.InternalAuthenticateRequest")
}

func init() { proto.RegisterFile("raft_internal.proto", fileDescriptor_b4c9a9be0cfca103) }

var fileDescriptor_b4c9a9be0cfca103 = []byte{
	// 1003 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x96, 0xd9, 0x72, 0x1b, 0x45,
	0x14, 0x86, 0x23, 0xc5, 0x71, 0xac, 0x96, 0xed, 0x38, 0x6d, 0x87, 0x34, 0x72, 0x95, 0x70, 0x1c,
	0x12, 0xcc, 0x66, 0x53, 0xca, 0x03, 0x80, 0x90, 0x5c, 0x8e, 0xab, 0x42, 0x70, 0x4d, 0xcc, 0x52,
	0xc5, 0xc5, 0xd0, 0x9a, 0x39, 0x96, 0x06, 0xcf, 0x46, 0x77, 0x4b, 0x31, 0xef, 0x11, 0x28, 0x1e,
	0x83, 0xed, 0x21, 0x72, 0xc1, 0x62, 0xe0, 0x05, 0xc0, 0xdc, 0x70, 0x0f, 0xdc, 0x53, 0xbd, 0xcc,
	0x26, 0xb5, 0x7c, 0xa7, 0xf9, 0xcf, 0x7f, 0xbe, 0x73, 0xba, 0xe7, 0xf4, 0xa8, 0xd1, 0x3a, 0xa3,
	0x27, 0xc2, 0x0d, 0x62, 0x01, 0x2c, 0xa6, 0xe1, 0x6e, 0xca, 0x12, 0x91, 0xe0, 0x65, 0x10, 0x9e,
	0xcf, 0x81, 0x4d, 0x80, 0xa5, 0x83, 0xd6, 0xc6, 0x30, 0x19, 0x26, 0x2a, 0xb0, 0x27, 0x7f, 0x69,
	0x4f, 0x6b, 0xad, 0xf0, 0x18, 0xa5, 0xc1, 0x52, 0xcf, 0xfc, 0xbc, 0x2f, 0x83, 0x7b, 0x34, 0x0d,
	0xf6, 0x22, 0x88, 0x06, 0xc0, 0xf8, 0x28, 0x48, 0xd3, 0x41, 0xe9, 0x41, 0xfb, 0xb6, 0x3f, 0x45,
	0x2b, 0x0e, 0x7c, 0x3e, 0x06, 0x2e, 0x1e, 0x02, 0xf5, 0x81, 0xe1, 0x55, 0x54, 0x3f, 0xec, 0x93,
	0xda, 0x56, 0x6d, 0x67, 0xc1, 0xa9, 0x1f, 0xf6, 0x71, 0x0b, 0x2d, 0x8d, 0xb9, 0x6c, 0x2d, 0x02,
	0x52, 0xdf, 0xaa, 0xed, 0x34, 0x9c, 0xfc, 0x19, 0xdf, 0x45, 0x2b, 0x74, 0x2c, 0x46, 0x2e, 0x83,
	0x49, 0xc0, 0x83, 0x24, 0x26, 0x57, 0x55, 0xda, 0xb2, 0x14, 0x1d, 0xa3, 0x6d, 0x3f, 0xc3, 0x68,
	0xfd, 0xd0, 0xac, 0xce, 0xa1, 0x27, 0xc2, 0x94, 0xc3, 0x0f, 0xd0, 0xe2, 0x48, 0x95, 0x24, 0xfe,
	0x56, 0x6d, 0xa7, 0xd9, 0xd9, 0xdc, 0x2d, 0xaf, 0x79, 0xb7, 0xd2, 0x95, 0x63, 0xac, 0x33, 0xdd,
	0xdd, 0x43, 0xf5, 0x49, 0x47, 0xf5, 0xd5, 0xec, 0xdc, 0xb2, 0x02, 0x9c, 0xfa, 0xa4, 0x83, 0xdf,
	0x42, 0xd7, 0x18, 0x8d, 0x87, 0xa0, 0x1a, 0x6c, 0x76, 0x5a, 0x53, 0x4e, 0x19, 0xca, 0xec, 0xda,
	0x88, 0x5f, 0x43, 0x57, 0xd3, 0xb1, 0x20, 0x0b, 0xca, 0x4f, 0xaa, 0xfe, 0xa3, 0x71, 0xb6, 0x08,
	0x47, 0x9a, 0x70, 0x0f, 0x2d, 0xfb, 0x10, 0x82, 0x00, 0x57, 0x17, 0xb9, 0xa6, 0x92, 0xb6, 0xaa,
	0x49, 0x7d, 0xe5, 0xa8, 0x94, 0x6a, 0xfa, 0x85, 0x26, 0x0b, 0x8a, 0xb3, 0x98, 0x2c, 0xda, 0x0a,
	0x1e, 0x9f, 0xc5, 0x79, 0x41, 0x71, 0x16, 0xe3, 0xb7, 0x11, 0xf2, 0x92, 0x28, 0xa5, 0x9e, 0x90,
	0x9b, 0x7e, 0x5d, 0xa5, 0xbc, 0x54, 0x4d, 0xe9, 0xe5, 0xf1, 0x2c, 0xb3, 0x94, 0x82, 0xdf, 0x41,
	0xcd, 0x10, 0x28, 0x07, 0x77, 0xc8, 0x68, 0x2c, 0xc8, 0x92, 0x8d, 0xf0, 0x48, 0x1a, 0x0e, 0x64,
	0x3c, 0x27, 0x84, 0xb9, 0x24, 0xd7, 0xac, 0x09, 0x0c, 0x26, 0xc9, 0x29, 0x90, 0x86, 0x6d, 0xcd,
	0x0a, 0xe1, 0x28, 0x43, 0xbe, 0xe6, 0xb0, 0xd0, 0xe4, 0x6b, 0xa1, 0x21, 0x65, 0x11, 0x41, 0xb6,
	0xd7, 0xd2, 0x95, 0xa1, 0xfc, 0xb5, 0x28, 0x23, 0x7e, 0x1f, 0xad, 0xe9, 0xb2, 0xde, 0x08, 0xbc,
	0xd3, 0x34, 0x09, 0x62, 0x41, 0x9a, 0x2a, 0xf9, 0x65, 0x4b, 0xe9, 0x5e, 0x6e, 0xca, 0x30, 0x37,
	0xc2, 0xaa, 0x8e, 0xbb, 0xa8, 0xa9, 0x46, 0x18, 0x62, 0x3a, 0x08, 0x81, 0xfc, 0x6d, 0xdd, 0xcc,
	0xee, 0x58, 0x8c, 0xf6, 0x95, 0x21, 0xdf, 0x0a, 0x9a, 0x4b, 0xb8, 0x8f, 0xd4, 0xc0, 0xbb, 0x7e,
	0xc0, 0x15, 0xe3, 0x9f, 0xeb, 0xb6, 0xbd, 0x90, 0x8c, 0xbe, 0x76, 0xe4, 0x7b, 0x41, 0x0b, 0x2d,
	0x6f, 0x84, 0x0b, 0x2a, 0xc6, 0x9c, 0xfc, 0x37, 0xb7, 0x91, 0x27, 0xca, 0x50, 0x69, 0x44, 0x4b,
	0xf8, 0xb1, 0x6e, 0x04, 0x62, 0x11, 0x78, 0x54, 0x00, 0xf9, 0x57, 0x33, 0x5e, 0xad, 0x32, 0xb2,
	0xb3, 0xd8, 0x2d, 0x59, 0x33, 0x5a, 0x25, 0x1f, 0xef, 0x9b, 0xe3, 0x2d, 0xcf, 0xbb, 0x4b, 0x7d,
	0x9f, 0xfc, 0xb8, 0x34, 0x6f, 0x65, 0x1f, 0x70, 0x60, 0x5d, 0xdf, 0xaf, 0xac, 0xcc, 0x68, 0xf8,
	0x31, 0x5a, 0x2b, 0x30, 0x7a, 0xe4, 0xc9, 0x4f, 0x9a, 0x74, 0xd7, 0x4e, 0x32, 0x67, 0xc5, 0xc0,
	0x56, 0x69, 0x45, 0xae, 0xb6, 0x35, 0x04, 0x41, 0x7e, 0xbe, 0xb4, 0xad, 0x03, 0x10, 0x33, 0x6d,
	0x1d, 0x80, 0xc0, 0x43, 0xf4, 0x62, 0x81, 0xf1, 0x46, 0xf2, 0x10, 0xba, 0x29, 0xe5, 0xfc, 0x69,
	0xc2, 0x7c, 0xf2, 0x8b, 0x46, 0xbe, 0x6e, 0x47, 0xf6, 0x94, 0xfb, 0xc8, 0x98, 0x33, 0xfa, 0x0b,
	0xd4, 0x1a, 0xc6, 0x1f, 0xa3, 0x8d, 0x52, 0xbf, 0xf2, 0xf4, 0xb8, 0x2c, 0x09, 0x81, 0x9c, 0xeb,
	0x1a, 0xf7, 0xe7, 0xb4, 0xad, 0x4e, 0x5e, 0x52, 0x4c, 0xcb, 0x4d, 0x3a, 0x1d, 0xc1, 0x9f, 0xa0,
	0x5b, 0x05, 0x59, 0x1f, 0x44, 0x8d, 0xfe, 0x55, 0xa3, 0x5f, 0xb1, 0xa3, 0xcd, 0x89, 0x2c, 0xb1,
	0x31, 0x9d, 0x09, 0xe1, 0x87, 0x68, 0xb5, 0x80, 0x87, 0x01, 0x17, 0xe4, 0x37, 0x4d, 0xbd, 0x63,
	0xa7, 0x3e, 0x0a, 0xb8, 0xa8, 0xcc, 0x51, 0x26, 0xe6, 0x24, 0xd9, 0x9a, 0x26, 0xfd, 0x3e, 0x97,
	0x24, 0x4b, 0xcf, 0x90, 0x32, 0x31, 0x7f, 0xf5, 0x8a, 0x24, 0x27, 0xf2, 0x9b, 0xc6, 0xbc, 0x57,
	0x2f, 0x73, 0xa6, 0x27, 0xd2, 0x68, 0xf9, 0x44, 0x2a, 0x8c, 0x99, 0xc8, 0x6f, 0x1b, 0xf3, 0x26,
	0x52, 0x66, 0x59, 0x26, 0xb2, 0x90, 0xab, 0x6d, 0xc9, 0x89, 0xfc, 0xee, 0xd2, 0xb6, 0xa6, 0x27,
	0xd2, 0x68, 0xf8, 0x33, 0xd4, 0x2a, 0x61, 0xd4, 0xa0, 0xa4, 0xc0, 0xa2, 0x80, 0xab, 0xff, 0xd6,
	0xef, 0x35, 0xf3, 0x8d, 0x39, 0x4c, 0x69, 0x3f, 0xca, 0xdd, 0x19, 0xff, 0x36, 0xb5, 0xc7, 0x71,
	0x84, 0x36, 0x8b, 0x5a, 0x66, 0x74, 0x4a, 0xc5, 0x7e, 0xd0, 0xc5, 0xde, 0xb4, 0x17, 0xd3, 0x53,
	0x32, 0x5b, 0x8d, 0xd0, 0x39, 0x06, 0xfc, 0x11, 0x5a, 0xf7, 0xc2, 0x31, 0x17, 0xc0, 0xdc, 0x09,
	0x30, 0x29, 0xb9, 0x1c, 0x04, 0x79, 0x86, 0xcc, 0x11, 0x28, 0x5f, 0x52, 0x76, 0x7b, 0xda, 0xf9,
	0xa1, 0x36, 0x3e, 0x29, 0x76, 0xeb, 0xa6, 0x37, 0x1d, 0xc1, 0x14, 0xdd, 0xce, 0xc0, 0x9a, 0xe1,
	0x52, 0x21, 0x98, 0x82, 0x7f, 0x89, 0xcc, 0xe7, 0xcf, 0x06, 0x7f, 0x4f, 0x69, 0x5d, 0x21, 0x58,
	0x89, 0xbf, 0xe1, 0x59, 0x82, 0xf8, 0x18, 0x61, 0x3f, 0x79, 0x1a, 0x0f, 0x19, 0xf5, 0xc1, 0x0d,
	0xe2, 0x93, 0x44, 0xd1, 0xbf, 0xd2, 0xf4, 0x7b, 0x55, 0x7a, 0x3f, 0x33, 0x1e, 0xc6, 0x27, 0x49,
	0x89, 0xbc, 0xe6, 0x4f, 0x05, 0xb6, 0x6f, 0xa0, 0x95, 0xfd, 0x28, 0x15, 0x5f, 0x38, 0xc0, 0xd3,
	0x24, 0xe6, 0xb0, 0x9d, 0xa2, 0xcd, 0x4b, 0x3e, 0xcd, 0x18, 0xa3, 0x05, 0x75, 0x07, 0xab, 0xa9,
	0x3b, 0x98, 0xfa, 0x2d, 0xef, 0x66, 0xf9, 0x17, 0xcb, 0xdc, 0xcd, 0xb2, 0x67, 0x7c, 0x07, 0x2d,
	0xf3, 0x20, 0x4a, 0x43, 0x70, 0x45, 0x72, 0x0a, 0xfa, 0x6a, 0xd6, 0x70, 0x9a, 0x5a, 0x3b, 0x96,
	0xd2, 0xbb, 0x1b, 0xcf, 0xff, 0x6c, 0x5f, 0x79, 0x7e, 0xd1, 0xae, 0x9d, 0x5f, 0xb4, 0x6b, 0x7f,
	0x5c, 0xb4, 0x6b, 0x5f, 0xff, 0xd5, 0xbe, 0x32, 0x58, 0x54, 0x17, 0xc3, 0x07, 0xff, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x94, 0x6f, 0x64, 0x0a, 0x98, 0x0a, 0x00, 0x00,
}

func (m *RequestHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.AuthRevision != 0 {
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRevision))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintRaftInternal(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *InternalRaftRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InternalRaftRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.DowngradeInfoSet != nil {
		{
			size, err := m.DowngradeInfoSet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x51
		i--
		dAtA[i] = 0xb2
	}
	if m.ClusterMemberAttrSet != nil {
		{
			size, err := m.ClusterMemberAttrSet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x51
		i--
		dAtA[i] = 0xaa
	}
	if m.ClusterVersionSet != nil {
		{
			size, err := m.ClusterVersionSet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x51
		i--
		dAtA[i] = 0xa2
	}
	if m.AuthRoleRevokePermission != nil {
		{
			size, err := m.AuthRoleRevokePermission.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4b
		i--
		dAtA[i] = 0xa2
	}
	if m.AuthRoleGrantPermission != nil {
		{
			size, err := m.AuthRoleGrantPermission.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4b
		i--
		dAtA[i] = 0x9a
	}
	if m.AuthRoleGet != nil {
		{
			size, err := m.AuthRoleGet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4b
		i--
		dAtA[i] = 0x92
	}
	if m.AuthRoleDelete != nil {
		{
			size, err := m.AuthRoleDelete.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4b
		i--
		dAtA[i] = 0x8a
	}
	if m.AuthRoleAdd != nil {
		{
			size, err := m.AuthRoleAdd.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4b
		i--
		dAtA[i] = 0x82
	}
	if m.AuthRoleList != nil {
		{
			size, err := m.AuthRoleList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x45
		i--
		dAtA[i] = 0x9a
	}
	if m.AuthUserList != nil {
		{
			size, err := m.AuthUserList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x45
		i--
		dAtA[i] = 0x92
	}
	if m.AuthUserRevokeRole != nil {
		{
			size, err := m.AuthUserRevokeRole.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x45
		i--
		dAtA[i] = 0x8a
	}
	if m.AuthUserGrantRole != nil {
		{
			size, err := m.AuthUserGrantRole.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x45
		i--
		dAtA[i] = 0x82
	}
	if m.AuthUserChangePassword != nil {
		{
			size, err := m.AuthUserChangePassword.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x44
		i--
		dAtA[i] = 0xfa
	}
	if m.AuthUserGet != nil {
		{
			size, err := m.AuthUserGet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x44
		i--
		dAtA[i] = 0xf2
	}
	if m.AuthUserDelete != nil {
		{
			size, err := m.AuthUserDelete.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x44
		i--
		dAtA[i] = 0xea
	}
	if m.AuthUserAdd != nil {
		{
			size, err := m.AuthUserAdd.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x44
		i--
		dAtA[i] = 0xe2
	}
	if m.AuthStatus != nil {
		{
			size, err := m.AuthStatus.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3f
		i--
		dAtA[i] = 0xaa
	}
	if m.Authenticate != nil {
		{
			size, err := m.Authenticate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3f
		i--
		dAtA[i] = 0xa2
	}
	if m.AuthDisable != nil {
		{
			size, err := m.AuthDisable.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3f
		i--
		dAtA[i] = 0x9a
	}
	if m.AuthEnable != nil {
		{
			size, err := m.AuthEnable.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3e
		i--
		dAtA[i] = 0xc2
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0xa2
	}
	if m.LeaseCheckpoint != nil {
		{
			size, err := m.LeaseCheckpoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.Alarm != nil {
		{
			size, err := m.Alarm.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.LeaseRevoke != nil {
		{
			size, err := m.LeaseRevoke.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.LeaseGrant != nil {
		{
			size, err := m.LeaseGrant.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Compaction != nil {
		{
			size, err := m.Compaction.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Txn != nil {
		{
			size, err := m.Txn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.DeleteRange != nil {
		{
			size, err := m.DeleteRange.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Put != nil {
		{
			size, err := m.Put.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Range != nil {
		{
			size, err := m.Range.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.V2 != nil {
		{
			size, err := m.V2.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRaftInternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EmptyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmptyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmptyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *InternalAuthenticateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InternalAuthenticateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InternalAuthenticateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.SimpleToken) > 0 {
		i -= len(m.SimpleToken)
		copy(dAtA[i:], m.SimpleToken)
		i = encodeVarintRaftInternal(dAtA, i, uint64(len(m.SimpleToken)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintRaftInternal(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRaftInternal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRaftInternal(dAtA []byte, offset int, v uint64) int {
	offset -= sovRaftInternal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *RequestHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRaftInternal(uint64(m.ID))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRevision != 0 {
		n += 1 + sovRaftInternal(uint64(m.AuthRevision))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InternalRaftRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRaftInternal(uint64(m.ID))
	}
	if m.V2 != nil {
		l = m.V2.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Range != nil {
		l = m.Range.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Put != nil {
		l = m.Put.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.DeleteRange != nil {
		l = m.DeleteRange.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Txn != nil {
		l = m.Txn.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Compaction != nil {
		l = m.Compaction.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.LeaseGrant != nil {
		l = m.LeaseGrant.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.LeaseRevoke != nil {
		l = m.LeaseRevoke.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Alarm != nil {
		l = m.Alarm.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.LeaseCheckpoint != nil {
		l = m.LeaseCheckpoint.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Header != nil {
		l = m.Header.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthEnable != nil {
		l = m.AuthEnable.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthDisable != nil {
		l = m.AuthDisable.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.Authenticate != nil {
		l = m.Authenticate.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthStatus != nil {
		l = m.AuthStatus.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserAdd != nil {
		l = m.AuthUserAdd.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserDelete != nil {
		l = m.AuthUserDelete.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserGet != nil {
		l = m.AuthUserGet.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserChangePassword != nil {
		l = m.AuthUserChangePassword.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserGrantRole != nil {
		l = m.AuthUserGrantRole.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserRevokeRole != nil {
		l = m.AuthUserRevokeRole.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserList != nil {
		l = m.AuthUserList.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleList != nil {
		l = m.AuthRoleList.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleAdd != nil {
		l = m.AuthRoleAdd.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleDelete != nil {
		l = m.AuthRoleDelete.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleGet != nil {
		l = m.AuthRoleGet.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleGrantPermission != nil {
		l = m.AuthRoleGrantPermission.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleRevokePermission != nil {
		l = m.AuthRoleRevokePermission.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.ClusterVersionSet != nil {
		l = m.ClusterVersionSet.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.ClusterMemberAttrSet != nil {
		l = m.ClusterMemberAttrSet.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.DowngradeInfoSet != nil {
		l = m.DowngradeInfoSet.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EmptyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InternalAuthenticateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	l = len(m.SimpleToken)
	if l > 0 {
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRaftInternal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozRaftInternal(x uint64) (n int) {
	return sovRaftInternal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *RequestHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftInternal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RequestHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRaftInternal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRevision", wireType)
			}
			m.AuthRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthRevision |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRaftInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRaftInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *EmptyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftInternal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmptyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRaftInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRaftInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *InternalAuthenticateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftInternal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InternalAuthenticateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalAuthenticateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRaftInternal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRaftInternal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SimpleToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRaftInternal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SimpleToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaftInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRaftInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skipRaftInternal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaftInternal
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRaftInternal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRaftInternal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRaftInternal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRaftInternal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaftInternal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRaftInternal = fmt.Errorf("proto: unexpected end of group")
)
