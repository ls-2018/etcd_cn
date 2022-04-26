// Code generated by protoc-gen-gogo.
// source: lease.proto

package leasepb

import (
	"encoding/json"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	etcdserverpb "github.com/ls-2018/etcd_cn/offical/etcdserverpb"
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

type Lease struct {
	ID           int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TTL          int64 `protobuf:"varint,2,opt,name=TTL,proto3" json:"TTL,omitempty"`
	RemainingTTL int64 `protobuf:"varint,3,opt,name=RemainingTTL,proto3" json:"RemainingTTL,omitempty"`
}

func (m *Lease) Reset()         { *m = Lease{} }
func (m *Lease) String() string { return proto.CompactTextString(m) }
func (*Lease) ProtoMessage()    {}
func (*Lease) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dd57e402472b33a, []int{0}
}

type LeaseInternalRequest struct {
	LeaseTimeToLiveRequest *etcdserverpb.LeaseTimeToLiveRequest `protobuf:"bytes,1,opt,name=LeaseTimeToLiveRequest,proto3" json:"LeaseTimeToLiveRequest,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                             `json:"-"`
	XXX_unrecognized       []byte                               `json:"-"`
	XXX_sizecache          int32                                `json:"-"`
}

func (m *LeaseInternalRequest) Reset()         { *m = LeaseInternalRequest{} }
func (m *LeaseInternalRequest) String() string { return proto.CompactTextString(m) }
func (*LeaseInternalRequest) ProtoMessage()    {}
func (*LeaseInternalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dd57e402472b33a, []int{1}
}

type LeaseInternalResponse struct {
	LeaseTimeToLiveResponse *etcdserverpb.LeaseTimeToLiveResponse `protobuf:"bytes,1,opt,name=LeaseTimeToLiveResponse,proto3" json:"LeaseTimeToLiveResponse,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                              `json:"-"`
	XXX_unrecognized        []byte                                `json:"-"`
	XXX_sizecache           int32                                 `json:"-"`
}

func (m *LeaseInternalResponse) Reset()         { *m = LeaseInternalResponse{} }
func (m *LeaseInternalResponse) String() string { return proto.CompactTextString(m) }
func (*LeaseInternalResponse) ProtoMessage()    {}
func (*LeaseInternalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dd57e402472b33a, []int{2}
}

func init() {
	proto.RegisterType((*Lease)(nil), "leasepb.Lease")
	proto.RegisterType((*LeaseInternalRequest)(nil), "leasepb.LeaseInternalRequest")
	proto.RegisterType((*LeaseInternalResponse)(nil), "leasepb.LeaseInternalResponse")
}

func init() { proto.RegisterFile("lease.proto", fileDescriptor_3dd57e402472b33a) }

var fileDescriptor_3dd57e402472b33a = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x49, 0x4d, 0x2c,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0x73, 0x0a, 0x92, 0xa4, 0x44, 0xd2,
	0xf3, 0xd3, 0xf3, 0xc1, 0x62, 0xfa, 0x20, 0x16, 0x44, 0x5a, 0x4a, 0x3e, 0xb5, 0x24, 0x39, 0x45,
	0x3f, 0xb1, 0x20, 0x53, 0x1f, 0xc4, 0x28, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x2a, 0x48, 0xd2, 0x2f,
	0x2a, 0x48, 0x86, 0x28, 0x50, 0xf2, 0xe5, 0x62, 0xf5, 0x01, 0x99, 0x20, 0xc4, 0xc7, 0xc5, 0xe4,
	0xe9, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1c, 0xc4, 0xe4, 0xe9, 0x22, 0x24, 0xc0, 0xc5, 0x1c,
	0x12, 0xe2, 0x23, 0xc1, 0x04, 0x16, 0x00, 0x31, 0x85, 0x94, 0xb8, 0x78, 0x82, 0x52, 0x73, 0x13,
	0x33, 0xf3, 0x32, 0xf3, 0xd2, 0x41, 0x52, 0xcc, 0x60, 0x29, 0x14, 0x31, 0xa5, 0x12, 0x2e, 0x11,
	0xb0, 0x71, 0x9e, 0x79, 0x25, 0xa9, 0x45, 0x79, 0x89, 0x39, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0x31, 0x5c, 0x62, 0x60, 0xf1, 0x90, 0xcc, 0xdc, 0xd4, 0x90, 0x7c, 0x9f, 0xcc, 0xb2,
	0x54, 0xa8, 0x0c, 0xd8, 0x46, 0x6e, 0x23, 0x15, 0x3d, 0x64, 0xf7, 0xe9, 0x61, 0x57, 0x1b, 0x84,
	0xc3, 0x0c, 0xa5, 0x0a, 0x2e, 0x51, 0x34, 0x5b, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xe2,
	0xb9, 0xc4, 0x31, 0xb4, 0x40, 0xa4, 0xa0, 0xf6, 0xaa, 0x12, 0xb0, 0x17, 0xa2, 0x38, 0x08, 0x97,
	0x29, 0x4e, 0x12, 0x27, 0x1e, 0xca, 0x31, 0x5c, 0x78, 0x28, 0xc7, 0x70, 0xe2, 0x91, 0x1c, 0xe3,
	0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0x0e,
	0x5f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x8a, 0x94, 0xb9, 0xae, 0x01, 0x00, 0x00,
}

func (m *Lease) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *LeaseInternalRequest) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *LeaseInternalResponse) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

func (m *Lease) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}

func (m *LeaseInternalRequest) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}

func (m *LeaseInternalResponse) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}

func (m *Lease) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

func (m *LeaseInternalRequest) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

func (m *LeaseInternalResponse) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

var (
	ErrInvalidLengthLease        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLease          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLease = fmt.Errorf("proto: unexpected end of group")
)