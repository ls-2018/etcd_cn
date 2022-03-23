package etcdserverpb

import (
	fmt "fmt"
	io "io"
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

type Request struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID" json:"ID"`
	Method               string   `protobuf:"bytes,2,opt,name=Method" json:"Method"`
	Path                 string   `protobuf:"bytes,3,opt,name=Path" json:"Path"`
	Val                  string   `protobuf:"bytes,4,opt,name=Val" json:"Val"`
	Dir                  bool     `protobuf:"varint,5,opt,name=Dir" json:"Dir"`
	PrevValue            string   `protobuf:"bytes,6,opt,name=PrevValue" json:"PrevValue"`
	PrevIndex            uint64   `protobuf:"varint,7,opt,name=PrevIndex" json:"PrevIndex"`
	PrevExist            *bool    `protobuf:"varint,8,opt,name=PrevExist" json:"PrevExist,omitempty"`
	Expiration           int64    `protobuf:"varint,9,opt,name=Expiration" json:"Expiration"`
	Wait                 bool     `protobuf:"varint,10,opt,name=Wait" json:"Wait"`
	Since                uint64   `protobuf:"varint,11,opt,name=Since" json:"Since"`
	Recursive            bool     `protobuf:"varint,12,opt,name=Recursive" json:"Recursive"`
	Sorted               bool     `protobuf:"varint,13,opt,name=Sorted" json:"Sorted"`
	Quorum               bool     `protobuf:"varint,14,opt,name=Quorum" json:"Quorum"`
	Time                 int64    `protobuf:"varint,15,opt,name=Time" json:"Time"`
	Stream               bool     `protobuf:"varint,16,opt,name=Stream" json:"Stream"`
	Refresh              *bool    `protobuf:"varint,17,opt,name=Refresh" json:"Refresh,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ffbeb3bebbce7e, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}

func (m *Request) XXX_Size() int {
	return m.Size()
}

func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Metadata struct {
	NodeID               uint64   `protobuf:"varint,1,opt,name=NodeID" json:"NodeID"`
	ClusterID            uint64   `protobuf:"varint,2,opt,name=ClusterID" json:"ClusterID"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ffbeb3bebbce7e, []int{1}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}

func (m *Metadata) XXX_Size() int {
	return m.Size()
}

func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "etcdserverpb.Request")
	proto.RegisterType((*Metadata)(nil), "etcdserverpb.Metadata")
}

func init() { proto.RegisterFile("etcdserver.proto", fileDescriptor_09ffbeb3bebbce7e) }

var fileDescriptor_09ffbeb3bebbce7e = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd2, 0xdd, 0x6e, 0xda, 0x30,
	0x14, 0x07, 0x70, 0x0c, 0xe1, 0xcb, 0x63, 0x1b, 0xb3, 0xd0, 0x74, 0x84, 0xa6, 0x2c, 0x42, 0xbb,
	0xc8, 0xd5, 0xf6, 0x0e, 0x2c, 0x5c, 0x44, 0x2a, 0x15, 0x0d, 0x15, 0xbd, 0x76, 0xc9, 0x29, 0x58,
	0x02, 0x4c, 0x1d, 0x07, 0xf1, 0x06, 0x7d, 0x85, 0x3e, 0x12, 0x97, 0x7d, 0x82, 0xaa, 0xa5, 0x2f,
	0x52, 0x39, 0x24, 0xc4, 0xed, 0x5d, 0xf4, 0xfb, 0x9f, 0x1c, 0x1f, 0x7f, 0xd0, 0x2e, 0xea, 0x79,
	0x9c, 0xa0, 0xda, 0xa1, 0xfa, 0xbb, 0x55, 0x52, 0x4b, 0xd6, 0x29, 0x65, 0x7b, 0xdb, 0xef, 0x2d,
	0xe4, 0x42, 0x66, 0xc1, 0x3f, 0xf3, 0x75, 0xaa, 0x19, 0x3c, 0x38, 0xb4, 0x19, 0xe1, 0x7d, 0x8a,
	0x89, 0x66, 0x3d, 0x5a, 0x0d, 0x03, 0x20, 0x1e, 0xf1, 0x9d, 0xa1, 0x73, 0x78, 0xfe, 0x5d, 0x89,
	0xaa, 0x61, 0xc0, 0x7e, 0xd1, 0xc6, 0x18, 0xf5, 0x52, 0xc6, 0x50, 0xf5, 0x88, 0xdf, 0xce, 0x93,
	0xdc, 0x18, 0x50, 0x67, 0xc2, 0xf5, 0x12, 0x6a, 0x56, 0x96, 0x09, 0xfb, 0x49, 0x6b, 0x33, 0xbe,
	0x02, 0xc7, 0x0a, 0x0c, 0x18, 0x0f, 0x84, 0x82, 0xba, 0x47, 0xfc, 0x56, 0xe1, 0x81, 0x50, 0x6c,
	0x40, 0xdb, 0x13, 0x85, 0xbb, 0x19, 0x5f, 0xa5, 0x08, 0x0d, 0xeb, 0xaf, 0x92, 0x8b, 0x9a, 0x70,
	0x13, 0xe3, 0x1e, 0x9a, 0xd6, 0xa0, 0x25, 0x17, 0x35, 0xa3, 0xbd, 0x48, 0x34, 0xb4, 0xce, 0xab,
	0x90, 0xa8, 0x64, 0xf6, 0x87, 0xd2, 0xd1, 0x7e, 0x2b, 0x14, 0xd7, 0x42, 0x6e, 0xa0, 0xed, 0x11,
	0xbf, 0x96, 0x37, 0xb2, 0xdc, 0xec, 0xed, 0x86, 0x0b, 0x0d, 0xd4, 0x1a, 0x35, 0x13, 0xd6, 0xa7,
	0xf5, 0xa9, 0xd8, 0xcc, 0x11, 0xbe, 0x58, 0x33, 0x9c, 0xc8, 0xac, 0x1f, 0xe1, 0x3c, 0x55, 0x89,
	0xd8, 0x21, 0x74, 0xac, 0x5f, 0x4b, 0x36, 0x67, 0x3a, 0x95, 0x4a, 0x63, 0x0c, 0x5f, 0xad, 0x82,
	0xdc, 0x4c, 0x7a, 0x95, 0x4a, 0x95, 0xae, 0xe1, 0x9b, 0x9d, 0x9e, 0xcc, 0x4c, 0x75, 0x2d, 0xd6,
	0x08, 0xdf, 0xad, 0xa9, 0x33, 0xc9, 0xba, 0x6a, 0x85, 0x7c, 0x0d, 0xdd, 0x0f, 0x5d, 0x33, 0x63,
	0xae, 0xb9, 0xe8, 0x3b, 0x85, 0xc9, 0x12, 0x7e, 0x58, 0xa7, 0x52, 0xe0, 0xe0, 0x82, 0xb6, 0xc6,
	0xa8, 0x79, 0xcc, 0x35, 0x37, 0x9d, 0x2e, 0x65, 0x8c, 0x9f, 0x5e, 0x43, 0x6e, 0x66, 0x87, 0xff,
	0x57, 0x69, 0xa2, 0x51, 0x85, 0x41, 0xf6, 0x28, 0xce, 0xb7, 0x70, 0xe6, 0x61, 0xef, 0xf0, 0xea,
	0x56, 0x0e, 0x47, 0x97, 0x3c, 0x1d, 0x5d, 0xf2, 0x72, 0x74, 0xc9, 0xe3, 0x9b, 0x5b, 0x79, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0xee, 0x40, 0xba, 0xd6, 0xa4, 0x02, 0x00, 0x00,
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Refresh != nil {
		i--
		if *m.Refresh {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	i--
	if m.Stream {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x80
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.Time))
	i--
	dAtA[i] = 0x78
	i--
	if m.Quorum {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x70
	i--
	if m.Sorted {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x68
	i--
	if m.Recursive {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x60
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.Since))
	i--
	dAtA[i] = 0x58
	i--
	if m.Wait {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x50
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.Expiration))
	i--
	dAtA[i] = 0x48
	if m.PrevExist != nil {
		i--
		if *m.PrevExist {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.PrevIndex))
	i--
	dAtA[i] = 0x38
	i -= len(m.PrevValue)
	copy(dAtA[i:], m.PrevValue)
	i = encodeVarintEtcdserver(dAtA, i, uint64(len(m.PrevValue)))
	i--
	dAtA[i] = 0x32
	i--
	if m.Dir {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x28
	i -= len(m.Val)
	copy(dAtA[i:], m.Val)
	i = encodeVarintEtcdserver(dAtA, i, uint64(len(m.Val)))
	i--
	dAtA[i] = 0x22
	i -= len(m.Path)
	copy(dAtA[i:], m.Path)
	i = encodeVarintEtcdserver(dAtA, i, uint64(len(m.Path)))
	i--
	dAtA[i] = 0x1a
	i -= len(m.Method)
	copy(dAtA[i:], m.Method)
	i = encodeVarintEtcdserver(dAtA, i, uint64(len(m.Method)))
	i--
	dAtA[i] = 0x12
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.ID))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.ClusterID))
	i--
	dAtA[i] = 0x10
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.NodeID))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintEtcdserver(dAtA []byte, offset int, v uint64) int {
	offset -= sovEtcdserver(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovEtcdserver(uint64(m.ID))
	l = len(m.Method)
	n += 1 + l + sovEtcdserver(uint64(l))
	l = len(m.Path)
	n += 1 + l + sovEtcdserver(uint64(l))
	l = len(m.Val)
	n += 1 + l + sovEtcdserver(uint64(l))
	n += 2
	l = len(m.PrevValue)
	n += 1 + l + sovEtcdserver(uint64(l))
	n += 1 + sovEtcdserver(uint64(m.PrevIndex))
	if m.PrevExist != nil {
		n += 2
	}
	n += 1 + sovEtcdserver(uint64(m.Expiration))
	n += 2
	n += 1 + sovEtcdserver(uint64(m.Since))
	n += 2
	n += 2
	n += 2
	n += 1 + sovEtcdserver(uint64(m.Time))
	n += 3
	if m.Refresh != nil {
		n += 3
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovEtcdserver(uint64(m.NodeID))
	n += 1 + sovEtcdserver(uint64(m.ClusterID))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEtcdserver(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozEtcdserver(x uint64) (n int) {
	return sovEtcdserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func skipEtcdserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEtcdserver
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
					return 0, ErrIntOverflowEtcdserver
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
					return 0, ErrIntOverflowEtcdserver
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
				return 0, ErrInvalidLengthEtcdserver
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEtcdserver
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEtcdserver
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEtcdserver        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEtcdserver          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEtcdserver = fmt.Errorf("proto: unexpected end of group")
)
