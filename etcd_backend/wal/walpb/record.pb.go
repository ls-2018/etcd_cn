// Code generated by protoc-gen-gogo.
// source: record.proto

package walpb

import (
	"encoding/json"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	raftpb "github.com/ls-2018/etcd_cn/raft/raftpb"
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

type Record struct {
	Type int64  `protobuf:"varint,1,opt,name=type" json:"type"`
	Crc  uint32 `protobuf:"varint,2,opt,name=crc" json:"crc"`
	Data []byte `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{0}
}

// Keep in sync with raftpb.SnapshotMetadata.
type Snapshot struct {
	Index uint64 `protobuf:"varint,1,opt,name=index" json:"index"`
	Term  uint64 `protobuf:"varint,2,opt,name=term" json:"term"`
	// Field populated since >=etcd-3.5.0.
	ConfState            *raftpb.ConfState `protobuf:"bytes,3,opt,name=conf_state,json=confState" json:"conf_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{1}
}

func init() {
	proto.RegisterType((*Record)(nil), "walpb.Record")
	proto.RegisterType((*Snapshot)(nil), "walpb.Snapshot")
}

func init() { proto.RegisterFile("record.proto", fileDescriptor_bf94fd919e302a1d) }

var fileDescriptor_bf94fd919e302a1d = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0x45, 0x63, 0xe2, 0x22, 0x18, 0xca, 0x02, 0xab, 0xaa, 0xa2, 0x2c, 0x4c, 0xd4, 0x55, 0x56,
	0x29, 0xe2, 0x08, 0x65, 0xcf, 0x22, 0x3d, 0x00, 0x72, 0x1d, 0xa7, 0x20, 0xd1, 0x8c, 0x35, 0xb5,
	0x04, 0xdc, 0x84, 0x23, 0x65, 0xc9, 0x09, 0x10, 0x84, 0x8b, 0xa0, 0x8c, 0x03, 0x1b, 0xfb, 0xeb,
	0x7d, 0xf9, 0x7d, 0xc3, 0x9c, 0x9c, 0x45, 0x6a, 0x2a, 0x4f, 0x18, 0x50, 0xcd, 0x5e, 0xcc, 0xb3,
	0xdf, 0xe5, 0x8b, 0x3d, 0xee, 0x91, 0xc9, 0x7a, 0x4c, 0xb1, 0xcc, 0x97, 0x64, 0xda, 0xb0, 0x1e,
	0x0f, 0xbf, 0xe3, 0x2b, 0xf2, 0xd5, 0x3d, 0x9c, 0xd6, 0x2c, 0x51, 0x19, 0xc8, 0xf0, 0xe6, 0x5d,
	0x26, 0x0a, 0x51, 0xa6, 0x1b, 0xd9, 0x7f, 0x5e, 0x27, 0x35, 0x13, 0xb5, 0x84, 0xd4, 0x92, 0xcd,
	0x4e, 0x0a, 0x51, 0x5e, 0x4e, 0xc5, 0x08, 0x94, 0x02, 0xd9, 0x98, 0x60, 0xb2, 0xb4, 0x10, 0xe5,
	0xbc, 0xe6, 0xbc, 0x22, 0x38, 0xdb, 0x76, 0xc6, 0x1f, 0x1f, 0x31, 0xa8, 0x1c, 0x66, 0x4f, 0x5d,
	0xe3, 0x5e, 0x59, 0x29, 0xa7, 0x97, 0x11, 0xf1, 0x9a, 0xa3, 0x03, 0x4b, 0xe5, 0xff, 0x9a, 0xa3,
	0x83, 0xba, 0x01, 0xb0, 0xd8, 0xb5, 0x0f, 0xc7, 0x60, 0x82, 0x63, 0xf7, 0xc5, 0xed, 0x55, 0x15,
	0x7f, 0x5e, 0xdd, 0x61, 0xd7, 0x6e, 0xc7, 0xa2, 0x3e, 0xb7, 0x7f, 0x71, 0xb3, 0xe8, 0xbf, 0x75,
	0xd2, 0x0f, 0x5a, 0x7c, 0x0c, 0x5a, 0x7c, 0x0d, 0x5a, 0xbc, 0xff, 0xe8, 0xe4, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xc3, 0x36, 0x0c, 0xad, 0x1d, 0x01, 0x00, 0x00,
}

func (m *Record) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}

func (m *Snapshot) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}
