// Code generated by protoc-gen-gogo.
// source: snap.proto

package snappb

import (
	"encoding/json"
	fmt "fmt"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type Snapshot struct {
	Crc  uint32 `protobuf:"varint,1,opt,name=crc" json:"crc"`
	Data []byte `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2e3c045ebf84d00, []int{0}
}

func init() {
	proto.RegisterType((*Snapshot)(nil), "snappb.snapshot")
}

func init() { proto.RegisterFile("snap.proto", fileDescriptor_f2e3c045ebf84d00) }

var fileDescriptor_f2e3c045ebf84d00 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xce, 0x4b, 0x2c,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0xb1, 0x0b, 0x92, 0xa4, 0x44, 0xd2, 0xf3,
	0xd3, 0xf3, 0xc1, 0x42, 0xfa, 0x20, 0x16, 0x44, 0x56, 0xc9, 0x8c, 0x8b, 0x03, 0x24, 0x5f, 0x9c,
	0x91, 0x5f, 0x22, 0x24, 0xc6, 0xc5, 0x9c, 0x5c, 0x94, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xeb,
	0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x48, 0x40, 0x48, 0x88, 0x8b, 0x25, 0x25, 0xb1, 0x24,
	0x51, 0x82, 0x49, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc, 0x76, 0x12, 0x39, 0xf1, 0x50, 0x8e, 0xe1,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e,
	0x01, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x0f, 0x32, 0xb2, 0x78, 0x00, 0x00, 0x00,
}

func (m *Snapshot) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}

func sovSnap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

var (
	ErrInvalidLengthSnap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSnap = fmt.Errorf("proto: unexpected end of group")
)
