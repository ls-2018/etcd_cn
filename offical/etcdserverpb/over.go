package etcdserverpb

import (
	"encoding/json"
	"errors"
	"strings"

	"go.etcd.io/etcd/api/v3/membershippb"
)

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *Metadata) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *Request) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

type ASD struct {
	Header *RequestHeader
	ID     uint64
	V2     *Request
	Put    *struct {
		Key         string
		Value       string
		Lease       int64
		PrevKv      bool
		IgnoreValue bool
		IgnoreLease bool
	}
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
		Put: &struct {
			Key         string
			Value       string
			Lease       int64
			PrevKv      bool
			IgnoreValue bool
			IgnoreLease bool
		}{
			Key:         string(m.Put.Key),
			Value:       string(m.Put.Value),
			Lease:       m.Put.Lease,
			PrevKv:      m.Put.PrevKv,
			IgnoreValue: m.Put.IgnoreValue,
			IgnoreLease: m.Put.IgnoreLease,
		},
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
	m.Put = &PutRequest{
		Key:         []byte(a.Put.Key),
		Value:       []byte(a.Put.Value),
		Lease:       a.Put.Lease,
		PrevKv:      a.Put.PrevKv,
		IgnoreValue: a.Put.IgnoreValue,
		IgnoreLease: a.Put.IgnoreLease,
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

//func (m *InternalRaftRequest) Marshal() (dAtA []byte, err error) {
//	size := m.Size()
//	dAtA = make([]byte, size)
//	n, err := m.MarshalToSizedBuffer(dAtA[:size])
//	if err != nil {
//		return nil, err
//	}
//	return dAtA[:n], nil
//}
//func (m *InternalRaftRequest) Unmarshal(dAtA []byte) error {
//	l := len(dAtA)
//	iNdEx := 0
//	for iNdEx < l {
//		preIndex := iNdEx
//		var wire uint64
//		for shift := uint(0); ; shift += 7 {
//			if shift >= 64 {
//				return ErrIntOverflowRaftInternal
//			}
//			if iNdEx >= l {
//				return io.ErrUnexpectedEOF
//			}
//			b := dAtA[iNdEx]
//			iNdEx++
//			wire |= uint64(b&0x7F) << shift
//			if b < 0x80 {
//				break
//			}
//		}
//		fieldNum := int32(wire >> 3)
//		wireType := int(wire & 0x7)
//		if wireType == 4 {
//			return fmt.Errorf("proto: InternalRaftRequest: wiretype end group for non-group")
//		}
//		if fieldNum <= 0 {
//			return fmt.Errorf("proto: InternalRaftRequest: illegal tag %d (wire type %d)", fieldNum, wire)
//		}
//		switch fieldNum {
//		case 1:
//			if wireType != 0 {
//				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
//			}
//			m.ID = 0
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				m.ID |= uint64(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//		case 2:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field V2", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.V2 == nil {
//				m.V2 = &Request{}
//			}
//			if err := m.V2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 3:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Range", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.Range == nil {
//				m.Range = &RangeRequest{}
//			}
//			if err := m.Range.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 4:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Put", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.Put == nil {
//				m.Put = &PutRequest{}
//			}
//			if err := m.Put.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 5:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field DeleteRange", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.DeleteRange == nil {
//				m.DeleteRange = &DeleteRangeRequest{}
//			}
//			if err := m.DeleteRange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 6:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.Txn == nil {
//				m.Txn = &TxnRequest{}
//			}
//			if err := m.Txn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 7:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Compaction", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.Compaction == nil {
//				m.Compaction = &CompactionRequest{}
//			}
//			if err := m.Compaction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 8:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field LeaseGrant", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.LeaseGrant == nil {
//				m.LeaseGrant = &LeaseGrantRequest{}
//			}
//			if err := m.LeaseGrant.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 9:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field LeaseRevoke", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.LeaseRevoke == nil {
//				m.LeaseRevoke = &LeaseRevokeRequest{}
//			}
//			if err := m.LeaseRevoke.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 10:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Alarm", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.Alarm == nil {
//				m.Alarm = &AlarmRequest{}
//			}
//			if err := m.Alarm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 11:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field LeaseCheckpoint", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.LeaseCheckpoint == nil {
//				m.LeaseCheckpoint = &LeaseCheckpointRequest{}
//			}
//			if err := m.LeaseCheckpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 100:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.Header == nil {
//				m.Header = &RequestHeader{}
//			}
//			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1000:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthEnable", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthEnable == nil {
//				m.AuthEnable = &AuthEnableRequest{}
//			}
//			if err := m.AuthEnable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1011:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthDisable", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthDisable == nil {
//				m.AuthDisable = &AuthDisableRequest{}
//			}
//			if err := m.AuthDisable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1012:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field Authenticate", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.Authenticate == nil {
//				m.Authenticate = &InternalAuthenticateRequest{}
//			}
//			if err := m.Authenticate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1013:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthStatus", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthStatus == nil {
//				m.AuthStatus = &AuthStatusRequest{}
//			}
//			if err := m.AuthStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1100:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserAdd", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthUserAdd == nil {
//				m.AuthUserAdd = &AuthUserAddRequest{}
//			}
//			if err := m.AuthUserAdd.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1101:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserDelete", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthUserDelete == nil {
//				m.AuthUserDelete = &AuthUserDeleteRequest{}
//			}
//			if err := m.AuthUserDelete.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1102:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserGet", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthUserGet == nil {
//				m.AuthUserGet = &AuthUserGetRequest{}
//			}
//			if err := m.AuthUserGet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1103:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserChangePassword", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthUserChangePassword == nil {
//				m.AuthUserChangePassword = &AuthUserChangePasswordRequest{}
//			}
//			if err := m.AuthUserChangePassword.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1104:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserGrantRole", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthUserGrantRole == nil {
//				m.AuthUserGrantRole = &AuthUserGrantRoleRequest{}
//			}
//			if err := m.AuthUserGrantRole.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1105:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserRevokeRole", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthUserRevokeRole == nil {
//				m.AuthUserRevokeRole = &AuthUserRevokeRoleRequest{}
//			}
//			if err := m.AuthUserRevokeRole.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1106:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserList", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthUserList == nil {
//				m.AuthUserList = &AuthUserListRequest{}
//			}
//			if err := m.AuthUserList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1107:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleList", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthRoleList == nil {
//				m.AuthRoleList = &AuthRoleListRequest{}
//			}
//			if err := m.AuthRoleList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1200:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleAdd", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthRoleAdd == nil {
//				m.AuthRoleAdd = &AuthRoleAddRequest{}
//			}
//			if err := m.AuthRoleAdd.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1201:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleDelete", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthRoleDelete == nil {
//				m.AuthRoleDelete = &AuthRoleDeleteRequest{}
//			}
//			if err := m.AuthRoleDelete.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1202:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleGet", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthRoleGet == nil {
//				m.AuthRoleGet = &AuthRoleGetRequest{}
//			}
//			if err := m.AuthRoleGet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1203:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleGrantPermission", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthRoleGrantPermission == nil {
//				m.AuthRoleGrantPermission = &AuthRoleGrantPermissionRequest{}
//			}
//			if err := m.AuthRoleGrantPermission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1204:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleRevokePermission", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.AuthRoleRevokePermission == nil {
//				m.AuthRoleRevokePermission = &AuthRoleRevokePermissionRequest{}
//			}
//			if err := m.AuthRoleRevokePermission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1300:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field ClusterVersionSet", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.ClusterVersionSet == nil {
//				m.ClusterVersionSet = &membershippb.ClusterVersionSetRequest{}
//			}
//			if err := m.ClusterVersionSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1301:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field ClusterMemberAttrSet", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.ClusterMemberAttrSet == nil {
//				m.ClusterMemberAttrSet = &membershippb.ClusterMemberAttrSetRequest{}
//			}
//			if err := m.ClusterMemberAttrSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		case 1302:
//			if wireType != 2 {
//				return fmt.Errorf("proto: wrong wireType = %d for field DowngradeInfoSet", wireType)
//			}
//			var msglen int
//			for shift := uint(0); ; shift += 7 {
//				if shift >= 64 {
//					return ErrIntOverflowRaftInternal
//				}
//				if iNdEx >= l {
//					return io.ErrUnexpectedEOF
//				}
//				b := dAtA[iNdEx]
//				iNdEx++
//				msglen |= int(b&0x7F) << shift
//				if b < 0x80 {
//					break
//				}
//			}
//			if msglen < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			postIndex := iNdEx + msglen
//			if postIndex < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if postIndex > l {
//				return io.ErrUnexpectedEOF
//			}
//			if m.DowngradeInfoSet == nil {
//				m.DowngradeInfoSet = &membershippb.DowngradeInfoSetRequest{}
//			}
//			if err := m.DowngradeInfoSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
//				return err
//			}
//			iNdEx = postIndex
//		default:
//			iNdEx = preIndex
//			skippy, err := skipRaftInternal(dAtA[iNdEx:])
//			if err != nil {
//				return err
//			}
//			if (skippy < 0) || (iNdEx+skippy) < 0 {
//				return ErrInvalidLengthRaftInternal
//			}
//			if (iNdEx + skippy) > l {
//				return io.ErrUnexpectedEOF
//			}
//			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
//			iNdEx += skippy
//		}
//	}
//
//	if iNdEx > l {
//		return io.ErrUnexpectedEOF
//	}
//	return nil
//}
