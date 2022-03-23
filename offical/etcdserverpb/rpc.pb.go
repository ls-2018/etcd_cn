package etcdserverpb

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	authpb "go.etcd.io/etcd/api/v3/authpb"
	mvccpb "go.etcd.io/etcd/api/v3/mvccpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AlarmType int32

const (
	AlarmType_NONE    AlarmType = 0
	AlarmType_NOSPACE AlarmType = 1
	AlarmType_CORRUPT AlarmType = 2
)

var AlarmType_name = map[int32]string{
	0: "NONE",
	1: "NOSPACE",
	2: "CORRUPT",
}

var AlarmType_value = map[string]int32{
	"NONE":    0,
	"NOSPACE": 1,
	"CORRUPT": 2,
}

func (x AlarmType) String() string {
	return proto.EnumName(AlarmType_name, int32(x))
}

func (AlarmType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

type RangeRequest_SortOrder int32

const (
	RangeRequest_NONE    RangeRequest_SortOrder = 0
	RangeRequest_ASCEND  RangeRequest_SortOrder = 1
	RangeRequest_DESCEND RangeRequest_SortOrder = 2
)

var RangeRequest_SortOrder_name = map[int32]string{
	0: "NONE",
	1: "ASCEND",
	2: "DESCEND",
}

var RangeRequest_SortOrder_value = map[string]int32{
	"NONE":    0,
	"ASCEND":  1,
	"DESCEND": 2,
}

func (x RangeRequest_SortOrder) String() string {
	return proto.EnumName(RangeRequest_SortOrder_name, int32(x))
}

func (RangeRequest_SortOrder) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1, 0}
}

type RangeRequest_SortTarget int32

const (
	RangeRequest_KEY     RangeRequest_SortTarget = 0
	RangeRequest_VERSION RangeRequest_SortTarget = 1
	RangeRequest_CREATE  RangeRequest_SortTarget = 2
	RangeRequest_MOD     RangeRequest_SortTarget = 3
	RangeRequest_VALUE   RangeRequest_SortTarget = 4
)

var RangeRequest_SortTarget_name = map[int32]string{
	0: "KEY",
	1: "VERSION",
	2: "CREATE",
	3: "MOD",
	4: "VALUE",
}

var RangeRequest_SortTarget_value = map[string]int32{
	"KEY":     0,
	"VERSION": 1,
	"CREATE":  2,
	"MOD":     3,
	"VALUE":   4,
}

func (x RangeRequest_SortTarget) String() string {
	return proto.EnumName(RangeRequest_SortTarget_name, int32(x))
}

func (RangeRequest_SortTarget) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1, 1}
}

type Compare_CompareResult int32

const (
	Compare_EQUAL     Compare_CompareResult = 0
	Compare_GREATER   Compare_CompareResult = 1
	Compare_LESS      Compare_CompareResult = 2
	Compare_NOT_EQUAL Compare_CompareResult = 3
)

var Compare_CompareResult_name = map[int32]string{
	0: "EQUAL",
	1: "GREATER",
	2: "LESS",
	3: "NOT_EQUAL",
}

var Compare_CompareResult_value = map[string]int32{
	"EQUAL":     0,
	"GREATER":   1,
	"LESS":      2,
	"NOT_EQUAL": 3,
}

func (x Compare_CompareResult) String() string {
	return proto.EnumName(Compare_CompareResult_name, int32(x))
}

func (Compare_CompareResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{9, 0}
}

type Compare_CompareTarget int32

const (
	Compare_VERSION Compare_CompareTarget = 0
	Compare_CREATE  Compare_CompareTarget = 1
	Compare_MOD     Compare_CompareTarget = 2
	Compare_VALUE   Compare_CompareTarget = 3
	Compare_LEASE   Compare_CompareTarget = 4
)

var Compare_CompareTarget_name = map[int32]string{
	0: "VERSION",
	1: "CREATE",
	2: "MOD",
	3: "VALUE",
	4: "LEASE",
}

var Compare_CompareTarget_value = map[string]int32{
	"VERSION": 0,
	"CREATE":  1,
	"MOD":     2,
	"VALUE":   3,
	"LEASE":   4,
}

func (x Compare_CompareTarget) String() string {
	return proto.EnumName(Compare_CompareTarget_name, int32(x))
}

func (Compare_CompareTarget) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{9, 1}
}

type WatchCreateRequest_FilterType int32

const (
	// filter out put event.
	WatchCreateRequest_NOPUT WatchCreateRequest_FilterType = 0
	// filter out delete event.
	WatchCreateRequest_NODELETE WatchCreateRequest_FilterType = 1
)

var WatchCreateRequest_FilterType_name = map[int32]string{
	0: "NOPUT",
	1: "NODELETE",
}

var WatchCreateRequest_FilterType_value = map[string]int32{
	"NOPUT":    0,
	"NODELETE": 1,
}

func (x WatchCreateRequest_FilterType) String() string {
	return proto.EnumName(WatchCreateRequest_FilterType_name, int32(x))
}

func (WatchCreateRequest_FilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{21, 0}
}

type AlarmRequest_AlarmAction int32

const (
	AlarmRequest_GET        AlarmRequest_AlarmAction = 0
	AlarmRequest_ACTIVATE   AlarmRequest_AlarmAction = 1
	AlarmRequest_DEACTIVATE AlarmRequest_AlarmAction = 2
)

var AlarmRequest_AlarmAction_name = map[int32]string{
	0: "GET",
	1: "ACTIVATE",
	2: "DEACTIVATE",
}

var AlarmRequest_AlarmAction_value = map[string]int32{
	"GET":        0,
	"ACTIVATE":   1,
	"DEACTIVATE": 2,
}

func (x AlarmRequest_AlarmAction) String() string {
	return proto.EnumName(AlarmRequest_AlarmAction_name, int32(x))
}

func (AlarmRequest_AlarmAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{54, 0}
}

type DowngradeRequest_DowngradeAction int32

const (
	DowngradeRequest_VALIDATE DowngradeRequest_DowngradeAction = 0
	DowngradeRequest_ENABLE   DowngradeRequest_DowngradeAction = 1
	DowngradeRequest_CANCEL   DowngradeRequest_DowngradeAction = 2
)

var DowngradeRequest_DowngradeAction_name = map[int32]string{
	0: "VALIDATE",
	1: "ENABLE",
	2: "CANCEL",
}

var DowngradeRequest_DowngradeAction_value = map[string]int32{
	"VALIDATE": 0,
	"ENABLE":   1,
	"CANCEL":   2,
}

func (x DowngradeRequest_DowngradeAction) String() string {
	return proto.EnumName(DowngradeRequest_DowngradeAction_name, int32(x))
}

func (DowngradeRequest_DowngradeAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{57, 0}
}

type ResponseHeader struct {
	// cluster_id is the ID of the cluster which sent the response.
	ClusterId uint64 `protobuf:"varint,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// member_id is the ID of the member which sent the response.
	MemberId uint64 `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	// revision is the key-value store revision when the request was applied.
	// For watch progress responses, the header.revision indicates progress. All future events
	// recieved in this stream are guaranteed to have a higher revision number than the
	// header.revision number.
	Revision int64 `protobuf:"varint,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// raft_term is the raft term when the request was applied.
	RaftTerm             uint64   `protobuf:"varint,4,opt,name=raft_term,json=raftTerm,proto3" json:"raft_term,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}

func (m *ResponseHeader) XXX_Size() int {
	return m.Size()
}

func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetClusterId() uint64 {
	if m != nil {
		return m.ClusterId
	}
	return 0
}

func (m *ResponseHeader) GetMemberId() uint64 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

func (m *ResponseHeader) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *ResponseHeader) GetRaftTerm() uint64 {
	if m != nil {
		return m.RaftTerm
	}
	return 0
}

type RangeRequest struct {
	// key is the first key for the range. If range_end is not given, the request only looks up key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// range_end is the upper bound on the requested range [key, range_end).
	// If range_end is '\0', the range is all keys >= key.
	// If range_end is key plus one (e.g., "aa"+1 == "ab", "a\xff"+1 == "b"),
	// then the range request gets all keys prefixed with key.
	// If both key and range_end are '\0', then the range request returns all keys.
	RangeEnd []byte `protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	// limit is a limit on the number of keys returned for the request. When limit is set to 0,
	// it is treated as no limit.
	Limit int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// revision is the point-in-time of the key-value store to use for the range.
	// If revision is less or equal to zero, the range is over the newest key-value store.
	// If the revision has been compacted, ErrCompacted is returned as a response.
	Revision int64 `protobuf:"varint,4,opt,name=revision,proto3" json:"revision,omitempty"`
	// sort_order is the order for returned sorted results.
	SortOrder RangeRequest_SortOrder `protobuf:"varint,5,opt,name=sort_order,json=sortOrder,proto3,enum=etcdserverpb.RangeRequest_SortOrder" json:"sort_order,omitempty"`
	// sort_target is the key-value field to use for sorting.
	SortTarget RangeRequest_SortTarget `protobuf:"varint,6,opt,name=sort_target,json=sortTarget,proto3,enum=etcdserverpb.RangeRequest_SortTarget" json:"sort_target,omitempty"`
	// serializable sets the range request to use serializable member-local reads.
	// Range requests are linearizable by default; linearizable requests have higher
	// latency and lower throughput than serializable requests but reflect the current
	// consensus of the cluster. For better performance, in exchange for possible stale reads,
	// a serializable range request is served locally without needing to reach consensus
	// with other nodes in the cluster.
	Serializable bool `protobuf:"varint,7,opt,name=serializable,proto3" json:"serializable,omitempty"`
	// keys_only when set returns only the keys and not the values.
	KeysOnly bool `protobuf:"varint,8,opt,name=keys_only,json=keysOnly,proto3" json:"keys_only,omitempty"`
	// count_only when set returns only the count of the keys in the range.
	CountOnly bool `protobuf:"varint,9,opt,name=count_only,json=countOnly,proto3" json:"count_only,omitempty"`
	// min_mod_revision is the lower bound for returned key mod revisions; all keys with
	// lesser mod revisions will be filtered away.
	MinModRevision int64 `protobuf:"varint,10,opt,name=min_mod_revision,json=minModRevision,proto3" json:"min_mod_revision,omitempty"`
	// max_mod_revision is the upper bound for returned key mod revisions; all keys with
	// greater mod revisions will be filtered away.
	MaxModRevision int64 `protobuf:"varint,11,opt,name=max_mod_revision,json=maxModRevision,proto3" json:"max_mod_revision,omitempty"`
	// min_create_revision is the lower bound for returned key create revisions; all keys with
	// lesser create revisions will be filtered away.
	MinCreateRevision int64 `protobuf:"varint,12,opt,name=min_create_revision,json=minCreateRevision,proto3" json:"min_create_revision,omitempty"`
	// max_create_revision is the upper bound for returned key create revisions; all keys with
	// greater create revisions will be filtered away.
	MaxCreateRevision    int64    `protobuf:"varint,13,opt,name=max_create_revision,json=maxCreateRevision,proto3" json:"max_create_revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RangeRequest) Reset()         { *m = RangeRequest{} }
func (m *RangeRequest) String() string { return proto.CompactTextString(m) }
func (*RangeRequest) ProtoMessage()    {}
func (*RangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *RangeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *RangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RangeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *RangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RangeRequest.Merge(m, src)
}

func (m *RangeRequest) XXX_Size() int {
	return m.Size()
}

func (m *RangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RangeRequest proto.InternalMessageInfo

func (m *RangeRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RangeRequest) GetRangeEnd() []byte {
	if m != nil {
		return m.RangeEnd
	}
	return nil
}

func (m *RangeRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RangeRequest) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *RangeRequest) GetSortOrder() RangeRequest_SortOrder {
	if m != nil {
		return m.SortOrder
	}
	return RangeRequest_NONE
}

func (m *RangeRequest) GetSortTarget() RangeRequest_SortTarget {
	if m != nil {
		return m.SortTarget
	}
	return RangeRequest_KEY
}

func (m *RangeRequest) GetSerializable() bool {
	if m != nil {
		return m.Serializable
	}
	return false
}

func (m *RangeRequest) GetKeysOnly() bool {
	if m != nil {
		return m.KeysOnly
	}
	return false
}

func (m *RangeRequest) GetCountOnly() bool {
	if m != nil {
		return m.CountOnly
	}
	return false
}

func (m *RangeRequest) GetMinModRevision() int64 {
	if m != nil {
		return m.MinModRevision
	}
	return 0
}

func (m *RangeRequest) GetMaxModRevision() int64 {
	if m != nil {
		return m.MaxModRevision
	}
	return 0
}

func (m *RangeRequest) GetMinCreateRevision() int64 {
	if m != nil {
		return m.MinCreateRevision
	}
	return 0
}

func (m *RangeRequest) GetMaxCreateRevision() int64 {
	if m != nil {
		return m.MaxCreateRevision
	}
	return 0
}

type RangeResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// kvs is the list of key-value pairs matched by the range request.
	// kvs is empty when count is requested.
	Kvs []*mvccpb.KeyValue `protobuf:"bytes,2,rep,name=kvs,proto3" json:"kvs,omitempty"`
	// more indicates if there are more keys to return in the requested range.
	More bool `protobuf:"varint,3,opt,name=more,proto3" json:"more,omitempty"`
	// count is set to the number of keys within the range when requested.
	Count                int64    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RangeResponse) Reset()         { *m = RangeResponse{} }
func (m *RangeResponse) String() string { return proto.CompactTextString(m) }
func (*RangeResponse) ProtoMessage()    {}
func (*RangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *RangeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *RangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RangeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *RangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RangeResponse.Merge(m, src)
}

func (m *RangeResponse) XXX_Size() int {
	return m.Size()
}

func (m *RangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RangeResponse proto.InternalMessageInfo

func (m *RangeResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RangeResponse) GetKvs() []*mvccpb.KeyValue {
	if m != nil {
		return m.Kvs
	}
	return nil
}

func (m *RangeResponse) GetMore() bool {
	if m != nil {
		return m.More
	}
	return false
}

func (m *RangeResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type PutRequest struct {
	// key is the key, in bytes, to put into the key-value store.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value is the value, in bytes, to associate with the key in the key-value store.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// lease is the lease ID to associate with the key in the key-value store. A lease
	// value of 0 indicates no lease.
	Lease int64 `protobuf:"varint,3,opt,name=lease,proto3" json:"lease,omitempty"`
	// If prev_kv is set, etcd gets the previous key-value pair before changing it.
	// The previous key-value pair will be returned in the put response.
	PrevKv bool `protobuf:"varint,4,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
	// If ignore_value is set, etcd updates the key using its current value.
	// Returns an error if the key does not exist.
	IgnoreValue bool `protobuf:"varint,5,opt,name=ignore_value,json=ignoreValue,proto3" json:"ignore_value,omitempty"`
	// If ignore_lease is set, etcd updates the key using its current lease.
	// Returns an error if the key does not exist.
	IgnoreLease          bool     `protobuf:"varint,6,opt,name=ignore_lease,json=ignoreLease,proto3" json:"ignore_lease,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}

func (m *PutRequest) XXX_Size() int {
	return m.Size()
}

func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PutRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PutRequest) GetLease() int64 {
	if m != nil {
		return m.Lease
	}
	return 0
}

func (m *PutRequest) GetPrevKv() bool {
	if m != nil {
		return m.PrevKv
	}
	return false
}

func (m *PutRequest) GetIgnoreValue() bool {
	if m != nil {
		return m.IgnoreValue
	}
	return false
}

func (m *PutRequest) GetIgnoreLease() bool {
	if m != nil {
		return m.IgnoreLease
	}
	return false
}

type PutResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// if prev_kv is set in the request, the previous key-value pair will be returned.
	PrevKv               *mvccpb.KeyValue `protobuf:"bytes,2,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{4}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}

func (m *PutResponse) XXX_Size() int {
	return m.Size()
}

func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

func (m *PutResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PutResponse) GetPrevKv() *mvccpb.KeyValue {
	if m != nil {
		return m.PrevKv
	}
	return nil
}

type DeleteRangeRequest struct {
	// key is the first key to delete in the range.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// range_end is the key following the last key to delete for the range [key, range_end).
	// If range_end is not given, the range is defined to contain only the key argument.
	// If range_end is one bit larger than the given key, then the range is all the keys
	// with the prefix (the given key).
	// If range_end is '\0', the range is all keys greater than or equal to the key argument.
	RangeEnd []byte `protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	// If prev_kv is set, etcd gets the previous key-value pairs before deleting it.
	// The previous key-value pairs will be returned in the delete response.
	PrevKv               bool     `protobuf:"varint,3,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRangeRequest) Reset()         { *m = DeleteRangeRequest{} }
func (m *DeleteRangeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRangeRequest) ProtoMessage()    {}
func (*DeleteRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{5}
}

func (m *DeleteRangeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *DeleteRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeleteRangeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *DeleteRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRangeRequest.Merge(m, src)
}

func (m *DeleteRangeRequest) XXX_Size() int {
	return m.Size()
}

func (m *DeleteRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRangeRequest proto.InternalMessageInfo

func (m *DeleteRangeRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *DeleteRangeRequest) GetRangeEnd() []byte {
	if m != nil {
		return m.RangeEnd
	}
	return nil
}

func (m *DeleteRangeRequest) GetPrevKv() bool {
	if m != nil {
		return m.PrevKv
	}
	return false
}

type DeleteRangeResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// deleted is the number of keys deleted by the delete range request.
	Deleted int64 `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// if prev_kv is set in the request, the previous key-value pairs will be returned.
	PrevKvs              []*mvccpb.KeyValue `protobuf:"bytes,3,rep,name=prev_kvs,json=prevKvs,proto3" json:"prev_kvs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DeleteRangeResponse) Reset()         { *m = DeleteRangeResponse{} }
func (m *DeleteRangeResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteRangeResponse) ProtoMessage()    {}
func (*DeleteRangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{6}
}

func (m *DeleteRangeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *DeleteRangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeleteRangeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *DeleteRangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRangeResponse.Merge(m, src)
}

func (m *DeleteRangeResponse) XXX_Size() int {
	return m.Size()
}

func (m *DeleteRangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRangeResponse proto.InternalMessageInfo

func (m *DeleteRangeResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DeleteRangeResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func (m *DeleteRangeResponse) GetPrevKvs() []*mvccpb.KeyValue {
	if m != nil {
		return m.PrevKvs
	}
	return nil
}

type RequestOp struct {
	// request is a union of request types accepted by a transaction.
	//
	// Types that are valid to be assigned to Request:
	//	*RequestOp_RequestRange
	//	*RequestOp_RequestPut
	//	*RequestOp_RequestDeleteRange
	//	*RequestOp_RequestTxn
	Request              isRequestOp_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RequestOp) Reset()         { *m = RequestOp{} }
func (m *RequestOp) String() string { return proto.CompactTextString(m) }
func (*RequestOp) ProtoMessage()    {}
func (*RequestOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{7}
}

func (m *RequestOp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *RequestOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestOp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *RequestOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestOp.Merge(m, src)
}

func (m *RequestOp) XXX_Size() int {
	return m.Size()
}

func (m *RequestOp) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestOp.DiscardUnknown(m)
}

var xxx_messageInfo_RequestOp proto.InternalMessageInfo

type isRequestOp_Request interface {
	isRequestOp_Request()
	MarshalTo([]byte) (int, error)
	Size() int
}

type RequestOp_RequestRange struct {
	RequestRange *RangeRequest `protobuf:"bytes,1,opt,name=request_range,json=requestRange,proto3,oneof" json:"request_range,omitempty"`
}
type RequestOp_RequestPut struct {
	RequestPut *PutRequest `protobuf:"bytes,2,opt,name=request_put,json=requestPut,proto3,oneof" json:"request_put,omitempty"`
}
type RequestOp_RequestDeleteRange struct {
	RequestDeleteRange *DeleteRangeRequest `protobuf:"bytes,3,opt,name=request_delete_range,json=requestDeleteRange,proto3,oneof" json:"request_delete_range,omitempty"`
}
type RequestOp_RequestTxn struct {
	RequestTxn *TxnRequest `protobuf:"bytes,4,opt,name=request_txn,json=requestTxn,proto3,oneof" json:"request_txn,omitempty"`
}

func (*RequestOp_RequestRange) isRequestOp_Request()       {}
func (*RequestOp_RequestPut) isRequestOp_Request()         {}
func (*RequestOp_RequestDeleteRange) isRequestOp_Request() {}
func (*RequestOp_RequestTxn) isRequestOp_Request()         {}

func (m *RequestOp) GetRequest() isRequestOp_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *RequestOp) GetRequestRange() *RangeRequest {
	if x, ok := m.GetRequest().(*RequestOp_RequestRange); ok {
		return x.RequestRange
	}
	return nil
}

func (m *RequestOp) GetRequestPut() *PutRequest {
	if x, ok := m.GetRequest().(*RequestOp_RequestPut); ok {
		return x.RequestPut
	}
	return nil
}

func (m *RequestOp) GetRequestDeleteRange() *DeleteRangeRequest {
	if x, ok := m.GetRequest().(*RequestOp_RequestDeleteRange); ok {
		return x.RequestDeleteRange
	}
	return nil
}

func (m *RequestOp) GetRequestTxn() *TxnRequest {
	if x, ok := m.GetRequest().(*RequestOp_RequestTxn); ok {
		return x.RequestTxn
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RequestOp) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RequestOp_RequestRange)(nil),
		(*RequestOp_RequestPut)(nil),
		(*RequestOp_RequestDeleteRange)(nil),
		(*RequestOp_RequestTxn)(nil),
	}
}

type ResponseOp struct {
	// response is a union of response types returned by a transaction.
	//
	// Types that are valid to be assigned to Response:
	//	*ResponseOp_ResponseRange
	//	*ResponseOp_ResponsePut
	//	*ResponseOp_ResponseDeleteRange
	//	*ResponseOp_ResponseTxn
	Response             isResponseOp_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ResponseOp) Reset()         { *m = ResponseOp{} }
func (m *ResponseOp) String() string { return proto.CompactTextString(m) }
func (*ResponseOp) ProtoMessage()    {}
func (*ResponseOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{8}
}

func (m *ResponseOp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *ResponseOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseOp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *ResponseOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseOp.Merge(m, src)
}

func (m *ResponseOp) XXX_Size() int {
	return m.Size()
}

func (m *ResponseOp) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseOp.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseOp proto.InternalMessageInfo

type isResponseOp_Response interface {
	isResponseOp_Response()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ResponseOp_ResponseRange struct {
	ResponseRange *RangeResponse `protobuf:"bytes,1,opt,name=response_range,json=responseRange,proto3,oneof" json:"response_range,omitempty"`
}
type ResponseOp_ResponsePut struct {
	ResponsePut *PutResponse `protobuf:"bytes,2,opt,name=response_put,json=responsePut,proto3,oneof" json:"response_put,omitempty"`
}
type ResponseOp_ResponseDeleteRange struct {
	ResponseDeleteRange *DeleteRangeResponse `protobuf:"bytes,3,opt,name=response_delete_range,json=responseDeleteRange,proto3,oneof" json:"response_delete_range,omitempty"`
}
type ResponseOp_ResponseTxn struct {
	ResponseTxn *TxnResponse `protobuf:"bytes,4,opt,name=response_txn,json=responseTxn,proto3,oneof" json:"response_txn,omitempty"`
}

func (*ResponseOp_ResponseRange) isResponseOp_Response()       {}
func (*ResponseOp_ResponsePut) isResponseOp_Response()         {}
func (*ResponseOp_ResponseDeleteRange) isResponseOp_Response() {}
func (*ResponseOp_ResponseTxn) isResponseOp_Response()         {}

func (m *ResponseOp) GetResponse() isResponseOp_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ResponseOp) GetResponseRange() *RangeResponse {
	if x, ok := m.GetResponse().(*ResponseOp_ResponseRange); ok {
		return x.ResponseRange
	}
	return nil
}

func (m *ResponseOp) GetResponsePut() *PutResponse {
	if x, ok := m.GetResponse().(*ResponseOp_ResponsePut); ok {
		return x.ResponsePut
	}
	return nil
}

func (m *ResponseOp) GetResponseDeleteRange() *DeleteRangeResponse {
	if x, ok := m.GetResponse().(*ResponseOp_ResponseDeleteRange); ok {
		return x.ResponseDeleteRange
	}
	return nil
}

func (m *ResponseOp) GetResponseTxn() *TxnResponse {
	if x, ok := m.GetResponse().(*ResponseOp_ResponseTxn); ok {
		return x.ResponseTxn
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResponseOp) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResponseOp_ResponseRange)(nil),
		(*ResponseOp_ResponsePut)(nil),
		(*ResponseOp_ResponseDeleteRange)(nil),
		(*ResponseOp_ResponseTxn)(nil),
	}
}

type Compare struct {
	// result is logical comparison operation for this comparison.
	Result Compare_CompareResult `protobuf:"varint,1,opt,name=result,proto3,enum=etcdserverpb.Compare_CompareResult" json:"result,omitempty"`
	// target is the key-value field to inspect for the comparison.
	Target Compare_CompareTarget `protobuf:"varint,2,opt,name=target,proto3,enum=etcdserverpb.Compare_CompareTarget" json:"target,omitempty"`
	// key is the subject key for the comparison operation.
	Key []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are valid to be assigned to TargetUnion:
	//	*Compare_Version
	//	*Compare_CreateRevision
	//	*Compare_ModRevision
	//	*Compare_Value
	//	*Compare_Lease
	TargetUnion isCompare_TargetUnion `protobuf_oneof:"target_union"`
	// range_end compares the given target to all keys in the range [key, range_end).
	// See RangeRequest for more details on key ranges.
	RangeEnd             []byte   `protobuf:"bytes,64,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Compare) Reset()         { *m = Compare{} }
func (m *Compare) String() string { return proto.CompactTextString(m) }
func (*Compare) ProtoMessage()    {}
func (*Compare) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{9}
}

func (m *Compare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Compare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Compare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Compare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Compare.Merge(m, src)
}

func (m *Compare) XXX_Size() int {
	return m.Size()
}

func (m *Compare) XXX_DiscardUnknown() {
	xxx_messageInfo_Compare.DiscardUnknown(m)
}

var xxx_messageInfo_Compare proto.InternalMessageInfo

type isCompare_TargetUnion interface {
	isCompare_TargetUnion()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Compare_Version struct {
	Version int64 `protobuf:"varint,4,opt,name=version,proto3,oneof" json:"version,omitempty"`
}
type Compare_CreateRevision struct {
	CreateRevision int64 `protobuf:"varint,5,opt,name=create_revision,json=createRevision,proto3,oneof" json:"create_revision,omitempty"`
}
type Compare_ModRevision struct {
	ModRevision int64 `protobuf:"varint,6,opt,name=mod_revision,json=modRevision,proto3,oneof" json:"mod_revision,omitempty"`
}
type Compare_Value struct {
	Value []byte `protobuf:"bytes,7,opt,name=value,proto3,oneof" json:"value,omitempty"`
}
type Compare_Lease struct {
	Lease int64 `protobuf:"varint,8,opt,name=lease,proto3,oneof" json:"lease,omitempty"`
}

func (*Compare_Version) isCompare_TargetUnion()        {}
func (*Compare_CreateRevision) isCompare_TargetUnion() {}
func (*Compare_ModRevision) isCompare_TargetUnion()    {}
func (*Compare_Value) isCompare_TargetUnion()          {}
func (*Compare_Lease) isCompare_TargetUnion()          {}

func (m *Compare) GetTargetUnion() isCompare_TargetUnion {
	if m != nil {
		return m.TargetUnion
	}
	return nil
}

func (m *Compare) GetResult() Compare_CompareResult {
	if m != nil {
		return m.Result
	}
	return Compare_EQUAL
}

func (m *Compare) GetTarget() Compare_CompareTarget {
	if m != nil {
		return m.Target
	}
	return Compare_VERSION
}

func (m *Compare) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Compare) GetVersion() int64 {
	if x, ok := m.GetTargetUnion().(*Compare_Version); ok {
		return x.Version
	}
	return 0
}

func (m *Compare) GetCreateRevision() int64 {
	if x, ok := m.GetTargetUnion().(*Compare_CreateRevision); ok {
		return x.CreateRevision
	}
	return 0
}

func (m *Compare) GetModRevision() int64 {
	if x, ok := m.GetTargetUnion().(*Compare_ModRevision); ok {
		return x.ModRevision
	}
	return 0
}

func (m *Compare) GetValue() []byte {
	if x, ok := m.GetTargetUnion().(*Compare_Value); ok {
		return x.Value
	}
	return nil
}

func (m *Compare) GetLease() int64 {
	if x, ok := m.GetTargetUnion().(*Compare_Lease); ok {
		return x.Lease
	}
	return 0
}

func (m *Compare) GetRangeEnd() []byte {
	if m != nil {
		return m.RangeEnd
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Compare) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Compare_Version)(nil),
		(*Compare_CreateRevision)(nil),
		(*Compare_ModRevision)(nil),
		(*Compare_Value)(nil),
		(*Compare_Lease)(nil),
	}
}

// From google paxosdb paper:
// Our implementation hinges around a powerful primitive which we call MultiOp. All other database
// operations except for iteration are implemented as a single call to MultiOp. A MultiOp is applied atomically
// and consists of three components:
// 1. A list of tests called guard. Each test in guard checks a single entry in the database. It may check
// for the absence or presence of a value, or compare with a given value. Two different tests in the guard
// may apply to the same or different entries in the database. All tests in the guard are applied and
// MultiOp returns the results. If all tests are true, MultiOp executes t op (see item 2 below), otherwise
// it executes f op (see item 3 below).
// 2. A list of database operations called t op. Each operation in the list is either an insert, delete, or
// lookup operation, and applies to a single database entry. Two different operations in the list may apply
// to the same or different entries in the database. These operations are executed
// if guard evaluates to
// true.
// 3. A list of database operations called f op. Like t op, but executed if guard evaluates to false.
type TxnRequest struct {
	// compare is a list of predicates representing a conjunction of terms.
	// If the comparisons succeed, then the success requests will be processed in order,
	// and the response will contain their respective responses in order.
	// If the comparisons fail, then the failure requests will be processed in order,
	// and the response will contain their respective responses in order.
	Compare []*Compare `protobuf:"bytes,1,rep,name=compare,proto3" json:"compare,omitempty"`
	// success is a list of requests which will be applied when compare evaluates to true.
	Success []*RequestOp `protobuf:"bytes,2,rep,name=success,proto3" json:"success,omitempty"`
	// failure is a list of requests which will be applied when compare evaluates to false.
	Failure              []*RequestOp `protobuf:"bytes,3,rep,name=failure,proto3" json:"failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TxnRequest) Reset()         { *m = TxnRequest{} }
func (m *TxnRequest) String() string { return proto.CompactTextString(m) }
func (*TxnRequest) ProtoMessage()    {}
func (*TxnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{10}
}

func (m *TxnRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *TxnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxnRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *TxnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxnRequest.Merge(m, src)
}

func (m *TxnRequest) XXX_Size() int {
	return m.Size()
}

func (m *TxnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TxnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TxnRequest proto.InternalMessageInfo

func (m *TxnRequest) GetCompare() []*Compare {
	if m != nil {
		return m.Compare
	}
	return nil
}

func (m *TxnRequest) GetSuccess() []*RequestOp {
	if m != nil {
		return m.Success
	}
	return nil
}

func (m *TxnRequest) GetFailure() []*RequestOp {
	if m != nil {
		return m.Failure
	}
	return nil
}

type TxnResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// succeeded is set to true if the compare evaluated to true or false otherwise.
	Succeeded bool `protobuf:"varint,2,opt,name=succeeded,proto3" json:"succeeded,omitempty"`
	// responses is a list of responses corresponding to the results from applying
	// success if succeeded is true or failure if succeeded is false.
	Responses            []*ResponseOp `protobuf:"bytes,3,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TxnResponse) Reset()         { *m = TxnResponse{} }
func (m *TxnResponse) String() string { return proto.CompactTextString(m) }
func (*TxnResponse) ProtoMessage()    {}
func (*TxnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{11}
}

func (m *TxnResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *TxnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxnResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *TxnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxnResponse.Merge(m, src)
}

func (m *TxnResponse) XXX_Size() int {
	return m.Size()
}

func (m *TxnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TxnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TxnResponse proto.InternalMessageInfo

func (m *TxnResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TxnResponse) GetSucceeded() bool {
	if m != nil {
		return m.Succeeded
	}
	return false
}

func (m *TxnResponse) GetResponses() []*ResponseOp {
	if m != nil {
		return m.Responses
	}
	return nil
}

// CompactionRequest compacts the key-value store up to a given revision. All superseded keys
// with a revision less than the compaction revision will be removed.
type CompactionRequest struct {
	// revision is the key-value store revision for the compaction operation.
	Revision int64 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	// physical is set so the RPC will wait until the compaction is physically
	// applied to the local database such that compacted entries are totally
	// removed from the backend database.
	Physical             bool     `protobuf:"varint,2,opt,name=physical,proto3" json:"physical,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompactionRequest) Reset()         { *m = CompactionRequest{} }
func (m *CompactionRequest) String() string { return proto.CompactTextString(m) }
func (*CompactionRequest) ProtoMessage()    {}
func (*CompactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{12}
}

func (m *CompactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *CompactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *CompactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactionRequest.Merge(m, src)
}

func (m *CompactionRequest) XXX_Size() int {
	return m.Size()
}

func (m *CompactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompactionRequest proto.InternalMessageInfo

func (m *CompactionRequest) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *CompactionRequest) GetPhysical() bool {
	if m != nil {
		return m.Physical
	}
	return false
}

type CompactionResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CompactionResponse) Reset()         { *m = CompactionResponse{} }
func (m *CompactionResponse) String() string { return proto.CompactTextString(m) }
func (*CompactionResponse) ProtoMessage()    {}
func (*CompactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{13}
}

func (m *CompactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *CompactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *CompactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactionResponse.Merge(m, src)
}

func (m *CompactionResponse) XXX_Size() int {
	return m.Size()
}

func (m *CompactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CompactionResponse proto.InternalMessageInfo

func (m *CompactionResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HashRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashRequest) Reset()         { *m = HashRequest{} }
func (m *HashRequest) String() string { return proto.CompactTextString(m) }
func (*HashRequest) ProtoMessage()    {}
func (*HashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{14}
}

func (m *HashRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *HashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HashRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *HashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashRequest.Merge(m, src)
}

func (m *HashRequest) XXX_Size() int {
	return m.Size()
}

func (m *HashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HashRequest proto.InternalMessageInfo

type HashKVRequest struct {
	// revision is the key-value store revision for the hash operation.
	Revision             int64    `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashKVRequest) Reset()         { *m = HashKVRequest{} }
func (m *HashKVRequest) String() string { return proto.CompactTextString(m) }
func (*HashKVRequest) ProtoMessage()    {}
func (*HashKVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{15}
}

func (m *HashKVRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *HashKVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HashKVRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *HashKVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashKVRequest.Merge(m, src)
}

func (m *HashKVRequest) XXX_Size() int {
	return m.Size()
}

func (m *HashKVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HashKVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HashKVRequest proto.InternalMessageInfo

func (m *HashKVRequest) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

type HashKVResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// hash is the hash value computed from the responding member's MVCC keys up to a given revision.
	Hash uint32 `protobuf:"varint,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// compact_revision is the compacted revision of key-value store when hash begins.
	CompactRevision      int64    `protobuf:"varint,3,opt,name=compact_revision,json=compactRevision,proto3" json:"compact_revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashKVResponse) Reset()         { *m = HashKVResponse{} }
func (m *HashKVResponse) String() string { return proto.CompactTextString(m) }
func (*HashKVResponse) ProtoMessage()    {}
func (*HashKVResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{16}
}

func (m *HashKVResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *HashKVResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HashKVResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *HashKVResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashKVResponse.Merge(m, src)
}

func (m *HashKVResponse) XXX_Size() int {
	return m.Size()
}

func (m *HashKVResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HashKVResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HashKVResponse proto.InternalMessageInfo

func (m *HashKVResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HashKVResponse) GetHash() uint32 {
	if m != nil {
		return m.Hash
	}
	return 0
}

func (m *HashKVResponse) GetCompactRevision() int64 {
	if m != nil {
		return m.CompactRevision
	}
	return 0
}

type HashResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// hash is the hash value computed from the responding member's KV's backend.
	Hash                 uint32   `protobuf:"varint,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashResponse) Reset()         { *m = HashResponse{} }
func (m *HashResponse) String() string { return proto.CompactTextString(m) }
func (*HashResponse) ProtoMessage()    {}
func (*HashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{17}
}

func (m *HashResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *HashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HashResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *HashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashResponse.Merge(m, src)
}

func (m *HashResponse) XXX_Size() int {
	return m.Size()
}

func (m *HashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HashResponse proto.InternalMessageInfo

func (m *HashResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HashResponse) GetHash() uint32 {
	if m != nil {
		return m.Hash
	}
	return 0
}

type SnapshotRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotRequest) Reset()         { *m = SnapshotRequest{} }
func (m *SnapshotRequest) String() string { return proto.CompactTextString(m) }
func (*SnapshotRequest) ProtoMessage()    {}
func (*SnapshotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{18}
}

func (m *SnapshotRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *SnapshotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *SnapshotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotRequest.Merge(m, src)
}

func (m *SnapshotRequest) XXX_Size() int {
	return m.Size()
}

func (m *SnapshotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotRequest proto.InternalMessageInfo

type SnapshotResponse struct {
	// header has the current key-value store information. The first header in the snapshot
	// stream indicates the point in time of the snapshot.
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// remaining_bytes is the number of blob bytes to be sent after this message
	RemainingBytes uint64 `protobuf:"varint,2,opt,name=remaining_bytes,json=remainingBytes,proto3" json:"remaining_bytes,omitempty"`
	// blob contains the next chunk of the snapshot in the snapshot stream.
	Blob                 []byte   `protobuf:"bytes,3,opt,name=blob,proto3" json:"blob,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotResponse) Reset()         { *m = SnapshotResponse{} }
func (m *SnapshotResponse) String() string { return proto.CompactTextString(m) }
func (*SnapshotResponse) ProtoMessage()    {}
func (*SnapshotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{19}
}

func (m *SnapshotResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *SnapshotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *SnapshotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotResponse.Merge(m, src)
}

func (m *SnapshotResponse) XXX_Size() int {
	return m.Size()
}

func (m *SnapshotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotResponse proto.InternalMessageInfo

func (m *SnapshotResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SnapshotResponse) GetRemainingBytes() uint64 {
	if m != nil {
		return m.RemainingBytes
	}
	return 0
}

func (m *SnapshotResponse) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

type WatchRequest struct {
	// request_union is a request to either create a new watcher or cancel an existing watcher.
	//
	// Types that are valid to be assigned to RequestUnion:
	//	*WatchRequest_CreateRequest
	//	*WatchRequest_CancelRequest
	//	*WatchRequest_ProgressRequest
	RequestUnion         isWatchRequest_RequestUnion `protobuf_oneof:"request_union"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{20}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}

func (m *WatchRequest) XXX_Size() int {
	return m.Size()
}

func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

type isWatchRequest_RequestUnion interface {
	isWatchRequest_RequestUnion()
	MarshalTo([]byte) (int, error)
	Size() int
}

type WatchRequest_CreateRequest struct {
	CreateRequest *WatchCreateRequest `protobuf:"bytes,1,opt,name=create_request,json=createRequest,proto3,oneof" json:"create_request,omitempty"`
}
type WatchRequest_CancelRequest struct {
	CancelRequest *WatchCancelRequest `protobuf:"bytes,2,opt,name=cancel_request,json=cancelRequest,proto3,oneof" json:"cancel_request,omitempty"`
}
type WatchRequest_ProgressRequest struct {
	ProgressRequest *WatchProgressRequest `protobuf:"bytes,3,opt,name=progress_request,json=progressRequest,proto3,oneof" json:"progress_request,omitempty"`
}

func (*WatchRequest_CreateRequest) isWatchRequest_RequestUnion()   {}
func (*WatchRequest_CancelRequest) isWatchRequest_RequestUnion()   {}
func (*WatchRequest_ProgressRequest) isWatchRequest_RequestUnion() {}

func (m *WatchRequest) GetRequestUnion() isWatchRequest_RequestUnion {
	if m != nil {
		return m.RequestUnion
	}
	return nil
}

func (m *WatchRequest) GetCreateRequest() *WatchCreateRequest {
	if x, ok := m.GetRequestUnion().(*WatchRequest_CreateRequest); ok {
		return x.CreateRequest
	}
	return nil
}

func (m *WatchRequest) GetCancelRequest() *WatchCancelRequest {
	if x, ok := m.GetRequestUnion().(*WatchRequest_CancelRequest); ok {
		return x.CancelRequest
	}
	return nil
}

func (m *WatchRequest) GetProgressRequest() *WatchProgressRequest {
	if x, ok := m.GetRequestUnion().(*WatchRequest_ProgressRequest); ok {
		return x.ProgressRequest
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WatchRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WatchRequest_CreateRequest)(nil),
		(*WatchRequest_CancelRequest)(nil),
		(*WatchRequest_ProgressRequest)(nil),
	}
}

type WatchCreateRequest struct {
	// key is the key to register for watching.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// range_end is the end of the range [key, range_end) to watch. If range_end is not given,
	// only the key argument is watched. If range_end is equal to '\0', all keys greater than
	// or equal to the key argument are watched.
	// If the range_end is one bit larger than the given key,
	// then all keys with the prefix (the given key) will be watched.
	RangeEnd []byte `protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	// start_revision is an optional revision to watch from (inclusive). No start_revision is "now".
	StartRevision int64 `protobuf:"varint,3,opt,name=start_revision,json=startRevision,proto3" json:"start_revision,omitempty"`
	// progress_notify is set so that the etcd server will periodically send a WatchResponse with
	// no events to the new watcher if there are no recent events. It is useful when clients
	// wish to recover a disconnected watcher starting from a recent known revision.
	// The etcd server may decide how often it will send notifications based on current load.
	ProgressNotify bool `protobuf:"varint,4,opt,name=progress_notify,json=progressNotify,proto3" json:"progress_notify,omitempty"`
	// filters filter the events at server side before it sends back to the watcher.
	Filters []WatchCreateRequest_FilterType `protobuf:"varint,5,rep,packed,name=filters,proto3,enum=etcdserverpb.WatchCreateRequest_FilterType" json:"filters,omitempty"`
	// If prev_kv is set, created watcher gets the previous KV before the event happens.
	// If the previous KV is already compacted, nothing will be returned.
	PrevKv bool `protobuf:"varint,6,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
	// If watch_id is provided and non-zero, it will be assigned to this watcher.
	// Since creating a watcher in etcd is not a synchronous operation,
	// this can be used ensure that ordering is correct when creating multiple
	// watchers on the same stream. Creating a watcher with an ID already in
	// use on the stream will cause an error to be returned.
	WatchId int64 `protobuf:"varint,7,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
	// fragment enables splitting large revisions into multiple watch responses.
	Fragment             bool     `protobuf:"varint,8,opt,name=fragment,proto3" json:"fragment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchCreateRequest) Reset()         { *m = WatchCreateRequest{} }
func (m *WatchCreateRequest) String() string { return proto.CompactTextString(m) }
func (*WatchCreateRequest) ProtoMessage()    {}
func (*WatchCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{21}
}

func (m *WatchCreateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *WatchCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WatchCreateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *WatchCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchCreateRequest.Merge(m, src)
}

func (m *WatchCreateRequest) XXX_Size() int {
	return m.Size()
}

func (m *WatchCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchCreateRequest proto.InternalMessageInfo

func (m *WatchCreateRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *WatchCreateRequest) GetRangeEnd() []byte {
	if m != nil {
		return m.RangeEnd
	}
	return nil
}

func (m *WatchCreateRequest) GetStartRevision() int64 {
	if m != nil {
		return m.StartRevision
	}
	return 0
}

func (m *WatchCreateRequest) GetProgressNotify() bool {
	if m != nil {
		return m.ProgressNotify
	}
	return false
}

func (m *WatchCreateRequest) GetFilters() []WatchCreateRequest_FilterType {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *WatchCreateRequest) GetPrevKv() bool {
	if m != nil {
		return m.PrevKv
	}
	return false
}

func (m *WatchCreateRequest) GetWatchId() int64 {
	if m != nil {
		return m.WatchId
	}
	return 0
}

func (m *WatchCreateRequest) GetFragment() bool {
	if m != nil {
		return m.Fragment
	}
	return false
}

type WatchCancelRequest struct {
	// watch_id is the watcher id to cancel so that no more events are transmitted.
	WatchId              int64    `protobuf:"varint,1,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchCancelRequest) Reset()         { *m = WatchCancelRequest{} }
func (m *WatchCancelRequest) String() string { return proto.CompactTextString(m) }
func (*WatchCancelRequest) ProtoMessage()    {}
func (*WatchCancelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{22}
}

func (m *WatchCancelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *WatchCancelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WatchCancelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *WatchCancelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchCancelRequest.Merge(m, src)
}

func (m *WatchCancelRequest) XXX_Size() int {
	return m.Size()
}

func (m *WatchCancelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchCancelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchCancelRequest proto.InternalMessageInfo

func (m *WatchCancelRequest) GetWatchId() int64 {
	if m != nil {
		return m.WatchId
	}
	return 0
}

// Requests the a watch stream progress status be sent in the watch response stream as soon as
// possible.
type WatchProgressRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchProgressRequest) Reset()         { *m = WatchProgressRequest{} }
func (m *WatchProgressRequest) String() string { return proto.CompactTextString(m) }
func (*WatchProgressRequest) ProtoMessage()    {}
func (*WatchProgressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{23}
}

func (m *WatchProgressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *WatchProgressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WatchProgressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *WatchProgressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchProgressRequest.Merge(m, src)
}

func (m *WatchProgressRequest) XXX_Size() int {
	return m.Size()
}

func (m *WatchProgressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchProgressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchProgressRequest proto.InternalMessageInfo

type WatchResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// watch_id is the ID of the watcher that corresponds to the response.
	WatchId int64 `protobuf:"varint,2,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
	// created is set to true if the response is for a create watch request.
	// The client should record the watch_id and expect to receive events for
	// the created watcher from the same stream.
	// All events sent to the created watcher will attach with the same watch_id.
	Created bool `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	// canceled is set to true if the response is for a cancel watch request.
	// No further events will be sent to the canceled watcher.
	Canceled bool `protobuf:"varint,4,opt,name=canceled,proto3" json:"canceled,omitempty"`
	// compact_revision is set to the minimum index if a watcher tries to watch
	// at a compacted index.
	//
	// This happens when creating a watcher at a compacted revision or the watcher cannot
	// catch up with the progress of the key-value store.
	//
	// The client should treat the watcher as canceled and should not try to create any
	// watcher with the same start_revision again.
	CompactRevision int64 `protobuf:"varint,5,opt,name=compact_revision,json=compactRevision,proto3" json:"compact_revision,omitempty"`
	// cancel_reason indicates the reason for canceling the watcher.
	CancelReason string `protobuf:"bytes,6,opt,name=cancel_reason,json=cancelReason,proto3" json:"cancel_reason,omitempty"`
	// framgment is true if large watch response was split over multiple responses.
	Fragment             bool            `protobuf:"varint,7,opt,name=fragment,proto3" json:"fragment,omitempty"`
	Events               []*mvccpb.Event `protobuf:"bytes,11,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WatchResponse) Reset()         { *m = WatchResponse{} }
func (m *WatchResponse) String() string { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()    {}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{24}
}

func (m *WatchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *WatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WatchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *WatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchResponse.Merge(m, src)
}

func (m *WatchResponse) XXX_Size() int {
	return m.Size()
}

func (m *WatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchResponse proto.InternalMessageInfo

func (m *WatchResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *WatchResponse) GetWatchId() int64 {
	if m != nil {
		return m.WatchId
	}
	return 0
}

func (m *WatchResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *WatchResponse) GetCanceled() bool {
	if m != nil {
		return m.Canceled
	}
	return false
}

func (m *WatchResponse) GetCompactRevision() int64 {
	if m != nil {
		return m.CompactRevision
	}
	return 0
}

func (m *WatchResponse) GetCancelReason() string {
	if m != nil {
		return m.CancelReason
	}
	return ""
}

func (m *WatchResponse) GetFragment() bool {
	if m != nil {
		return m.Fragment
	}
	return false
}

func (m *WatchResponse) GetEvents() []*mvccpb.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type LeaseGrantRequest struct {
	// TTL is the advisory time-to-live in seconds. Expired lease will return -1.
	TTL int64 `protobuf:"varint,1,opt,name=TTL,proto3" json:"TTL,omitempty"`
	// ID is the requested ID for the lease. If ID is set to 0, the lessor chooses an ID.
	ID                   int64    `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseGrantRequest) Reset()         { *m = LeaseGrantRequest{} }
func (m *LeaseGrantRequest) String() string { return proto.CompactTextString(m) }
func (*LeaseGrantRequest) ProtoMessage()    {}
func (*LeaseGrantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{25}
}

func (m *LeaseGrantRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseGrantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseGrantRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseGrantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseGrantRequest.Merge(m, src)
}

func (m *LeaseGrantRequest) XXX_Size() int {
	return m.Size()
}

func (m *LeaseGrantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseGrantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseGrantRequest proto.InternalMessageInfo

func (m *LeaseGrantRequest) GetTTL() int64 {
	if m != nil {
		return m.TTL
	}
	return 0
}

func (m *LeaseGrantRequest) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type LeaseGrantResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// ID is the lease ID for the granted lease.
	ID int64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	// TTL is the server chosen lease time-to-live in seconds.
	TTL                  int64    `protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
	Error                string   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseGrantResponse) Reset()         { *m = LeaseGrantResponse{} }
func (m *LeaseGrantResponse) String() string { return proto.CompactTextString(m) }
func (*LeaseGrantResponse) ProtoMessage()    {}
func (*LeaseGrantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{26}
}

func (m *LeaseGrantResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseGrantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseGrantResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseGrantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseGrantResponse.Merge(m, src)
}

func (m *LeaseGrantResponse) XXX_Size() int {
	return m.Size()
}

func (m *LeaseGrantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseGrantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseGrantResponse proto.InternalMessageInfo

func (m *LeaseGrantResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LeaseGrantResponse) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *LeaseGrantResponse) GetTTL() int64 {
	if m != nil {
		return m.TTL
	}
	return 0
}

func (m *LeaseGrantResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LeaseRevokeRequest struct {
	// ID is the lease ID to revoke. When the ID is revoked, all associated keys will be deleted.
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseRevokeRequest) Reset()         { *m = LeaseRevokeRequest{} }
func (m *LeaseRevokeRequest) String() string { return proto.CompactTextString(m) }
func (*LeaseRevokeRequest) ProtoMessage()    {}
func (*LeaseRevokeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{27}
}

func (m *LeaseRevokeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseRevokeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseRevokeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseRevokeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseRevokeRequest.Merge(m, src)
}

func (m *LeaseRevokeRequest) XXX_Size() int {
	return m.Size()
}

func (m *LeaseRevokeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseRevokeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseRevokeRequest proto.InternalMessageInfo

func (m *LeaseRevokeRequest) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type LeaseRevokeResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LeaseRevokeResponse) Reset()         { *m = LeaseRevokeResponse{} }
func (m *LeaseRevokeResponse) String() string { return proto.CompactTextString(m) }
func (*LeaseRevokeResponse) ProtoMessage()    {}
func (*LeaseRevokeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{28}
}

func (m *LeaseRevokeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseRevokeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseRevokeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseRevokeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseRevokeResponse.Merge(m, src)
}

func (m *LeaseRevokeResponse) XXX_Size() int {
	return m.Size()
}

func (m *LeaseRevokeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseRevokeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseRevokeResponse proto.InternalMessageInfo

func (m *LeaseRevokeResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type LeaseCheckpoint struct {
	// ID is the lease ID to checkpoint.
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Remaining_TTL is the remaining time until expiry of the lease.
	Remaining_TTL        int64    `protobuf:"varint,2,opt,name=remaining_TTL,json=remainingTTL,proto3" json:"remaining_TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseCheckpoint) Reset()         { *m = LeaseCheckpoint{} }
func (m *LeaseCheckpoint) String() string { return proto.CompactTextString(m) }
func (*LeaseCheckpoint) ProtoMessage()    {}
func (*LeaseCheckpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{29}
}

func (m *LeaseCheckpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseCheckpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseCheckpoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseCheckpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseCheckpoint.Merge(m, src)
}

func (m *LeaseCheckpoint) XXX_Size() int {
	return m.Size()
}

func (m *LeaseCheckpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseCheckpoint.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseCheckpoint proto.InternalMessageInfo

func (m *LeaseCheckpoint) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *LeaseCheckpoint) GetRemaining_TTL() int64 {
	if m != nil {
		return m.Remaining_TTL
	}
	return 0
}

type LeaseCheckpointRequest struct {
	Checkpoints          []*LeaseCheckpoint `protobuf:"bytes,1,rep,name=checkpoints,proto3" json:"checkpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LeaseCheckpointRequest) Reset()         { *m = LeaseCheckpointRequest{} }
func (m *LeaseCheckpointRequest) String() string { return proto.CompactTextString(m) }
func (*LeaseCheckpointRequest) ProtoMessage()    {}
func (*LeaseCheckpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{30}
}

func (m *LeaseCheckpointRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseCheckpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseCheckpointRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseCheckpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseCheckpointRequest.Merge(m, src)
}

func (m *LeaseCheckpointRequest) XXX_Size() int {
	return m.Size()
}

func (m *LeaseCheckpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseCheckpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseCheckpointRequest proto.InternalMessageInfo

func (m *LeaseCheckpointRequest) GetCheckpoints() []*LeaseCheckpoint {
	if m != nil {
		return m.Checkpoints
	}
	return nil
}

type LeaseCheckpointResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LeaseCheckpointResponse) Reset()         { *m = LeaseCheckpointResponse{} }
func (m *LeaseCheckpointResponse) String() string { return proto.CompactTextString(m) }
func (*LeaseCheckpointResponse) ProtoMessage()    {}
func (*LeaseCheckpointResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{31}
}

func (m *LeaseCheckpointResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseCheckpointResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseCheckpointResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseCheckpointResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseCheckpointResponse.Merge(m, src)
}

func (m *LeaseCheckpointResponse) XXX_Size() int {
	return m.Size()
}

func (m *LeaseCheckpointResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseCheckpointResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseCheckpointResponse proto.InternalMessageInfo

func (m *LeaseCheckpointResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type LeaseKeepAliveRequest struct {
	// ID is the lease ID for the lease to keep alive.
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseKeepAliveRequest) Reset()         { *m = LeaseKeepAliveRequest{} }
func (m *LeaseKeepAliveRequest) String() string { return proto.CompactTextString(m) }
func (*LeaseKeepAliveRequest) ProtoMessage()    {}
func (*LeaseKeepAliveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{32}
}

func (m *LeaseKeepAliveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseKeepAliveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseKeepAliveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseKeepAliveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseKeepAliveRequest.Merge(m, src)
}

func (m *LeaseKeepAliveRequest) XXX_Size() int {
	return m.Size()
}

func (m *LeaseKeepAliveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseKeepAliveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseKeepAliveRequest proto.InternalMessageInfo

func (m *LeaseKeepAliveRequest) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type LeaseKeepAliveResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// ID is the lease ID from the keep alive request.
	ID int64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	// TTL is the new time-to-live for the lease.
	TTL                  int64    `protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseKeepAliveResponse) Reset()         { *m = LeaseKeepAliveResponse{} }
func (m *LeaseKeepAliveResponse) String() string { return proto.CompactTextString(m) }
func (*LeaseKeepAliveResponse) ProtoMessage()    {}
func (*LeaseKeepAliveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{33}
}

func (m *LeaseKeepAliveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseKeepAliveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseKeepAliveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseKeepAliveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseKeepAliveResponse.Merge(m, src)
}

func (m *LeaseKeepAliveResponse) XXX_Size() int {
	return m.Size()
}

func (m *LeaseKeepAliveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseKeepAliveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseKeepAliveResponse proto.InternalMessageInfo

func (m *LeaseKeepAliveResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LeaseKeepAliveResponse) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *LeaseKeepAliveResponse) GetTTL() int64 {
	if m != nil {
		return m.TTL
	}
	return 0
}

type LeaseTimeToLiveRequest struct {
	// ID is the lease ID for the lease.
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// keys is true to query all the keys attached to this lease.
	Keys                 bool     `protobuf:"varint,2,opt,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseTimeToLiveRequest) Reset()         { *m = LeaseTimeToLiveRequest{} }
func (m *LeaseTimeToLiveRequest) String() string { return proto.CompactTextString(m) }
func (*LeaseTimeToLiveRequest) ProtoMessage()    {}
func (*LeaseTimeToLiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{34}
}

func (m *LeaseTimeToLiveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseTimeToLiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseTimeToLiveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseTimeToLiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseTimeToLiveRequest.Merge(m, src)
}

func (m *LeaseTimeToLiveRequest) XXX_Size() int {
	return m.Size()
}

func (m *LeaseTimeToLiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseTimeToLiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseTimeToLiveRequest proto.InternalMessageInfo

func (m *LeaseTimeToLiveRequest) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *LeaseTimeToLiveRequest) GetKeys() bool {
	if m != nil {
		return m.Keys
	}
	return false
}

type LeaseTimeToLiveResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// ID is the lease ID from the keep alive request.
	ID int64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	// TTL is the remaining TTL in seconds for the lease; the lease will expire in under TTL+1 seconds.
	TTL int64 `protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
	// GrantedTTL is the initial granted time in seconds upon lease creation/renewal.
	GrantedTTL int64 `protobuf:"varint,4,opt,name=grantedTTL,proto3" json:"grantedTTL,omitempty"`
	// Keys is the list of keys attached to this lease.
	Keys                 [][]byte `protobuf:"bytes,5,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseTimeToLiveResponse) Reset()         { *m = LeaseTimeToLiveResponse{} }
func (m *LeaseTimeToLiveResponse) String() string { return proto.CompactTextString(m) }
func (*LeaseTimeToLiveResponse) ProtoMessage()    {}
func (*LeaseTimeToLiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{35}
}

func (m *LeaseTimeToLiveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseTimeToLiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseTimeToLiveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseTimeToLiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseTimeToLiveResponse.Merge(m, src)
}

func (m *LeaseTimeToLiveResponse) XXX_Size() int {
	return m.Size()
}

func (m *LeaseTimeToLiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseTimeToLiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseTimeToLiveResponse proto.InternalMessageInfo

func (m *LeaseTimeToLiveResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LeaseTimeToLiveResponse) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *LeaseTimeToLiveResponse) GetTTL() int64 {
	if m != nil {
		return m.TTL
	}
	return 0
}

func (m *LeaseTimeToLiveResponse) GetGrantedTTL() int64 {
	if m != nil {
		return m.GrantedTTL
	}
	return 0
}

func (m *LeaseTimeToLiveResponse) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

type LeaseLeasesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseLeasesRequest) Reset()         { *m = LeaseLeasesRequest{} }
func (m *LeaseLeasesRequest) String() string { return proto.CompactTextString(m) }
func (*LeaseLeasesRequest) ProtoMessage()    {}
func (*LeaseLeasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{36}
}

func (m *LeaseLeasesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseLeasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseLeasesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseLeasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseLeasesRequest.Merge(m, src)
}

func (m *LeaseLeasesRequest) XXX_Size() int {
	return m.Size()
}

func (m *LeaseLeasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseLeasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseLeasesRequest proto.InternalMessageInfo

type LeaseStatus struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaseStatus) Reset()         { *m = LeaseStatus{} }
func (m *LeaseStatus) String() string { return proto.CompactTextString(m) }
func (*LeaseStatus) ProtoMessage()    {}
func (*LeaseStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{37}
}

func (m *LeaseStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseStatus.Merge(m, src)
}

func (m *LeaseStatus) XXX_Size() int {
	return m.Size()
}

func (m *LeaseStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseStatus.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseStatus proto.InternalMessageInfo

func (m *LeaseStatus) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type LeaseLeasesResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Leases               []*LeaseStatus  `protobuf:"bytes,2,rep,name=leases,proto3" json:"leases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LeaseLeasesResponse) Reset()         { *m = LeaseLeasesResponse{} }
func (m *LeaseLeasesResponse) String() string { return proto.CompactTextString(m) }
func (*LeaseLeasesResponse) ProtoMessage()    {}
func (*LeaseLeasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{38}
}

func (m *LeaseLeasesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *LeaseLeasesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseLeasesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *LeaseLeasesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseLeasesResponse.Merge(m, src)
}

func (m *LeaseLeasesResponse) XXX_Size() int {
	return m.Size()
}

func (m *LeaseLeasesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseLeasesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseLeasesResponse proto.InternalMessageInfo

func (m *LeaseLeasesResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LeaseLeasesResponse) GetLeases() []*LeaseStatus {
	if m != nil {
		return m.Leases
	}
	return nil
}

type Member struct {
	// ID is the member ID for this member.
	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// name is the human-readable name of the member. If the member is not started, the name will be an empty string.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// peerURLs is the list of URLs the member exposes to the cluster for communication.
	PeerURLs []string `protobuf:"bytes,3,rep,name=peerURLs,proto3" json:"peerURLs,omitempty"`
	// clientURLs is the list of URLs the member exposes to clients for communication. If the member is not started, clientURLs will be empty.
	ClientURLs []string `protobuf:"bytes,4,rep,name=clientURLs,proto3" json:"clientURLs,omitempty"`
	// isLearner indicates if the member is raft learner.
	IsLearner            bool     `protobuf:"varint,5,opt,name=isLearner,proto3" json:"isLearner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{39}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Member.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}

func (m *Member) XXX_Size() int {
	return m.Size()
}

func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Member) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Member) GetPeerURLs() []string {
	if m != nil {
		return m.PeerURLs
	}
	return nil
}

func (m *Member) GetClientURLs() []string {
	if m != nil {
		return m.ClientURLs
	}
	return nil
}

func (m *Member) GetIsLearner() bool {
	if m != nil {
		return m.IsLearner
	}
	return false
}

type MemberAddRequest struct {
	// peerURLs是新增成员用来与集群通信的URL列表。
	PeerURLs             []string `protobuf:"bytes,1,rep,name=peerURLs,proto3" json:"peerURLs,omitempty"`
	IsLearner            bool     `protobuf:"varint,2,opt,name=isLearner,proto3" json:"isLearner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberAddRequest) Reset()         { *m = MemberAddRequest{} }
func (m *MemberAddRequest) String() string { return proto.CompactTextString(m) }
func (*MemberAddRequest) ProtoMessage()    {}
func (*MemberAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{40}
}

func (m *MemberAddRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MemberAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberAddRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MemberAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberAddRequest.Merge(m, src)
}

func (m *MemberAddRequest) XXX_Size() int {
	return m.Size()
}

func (m *MemberAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberAddRequest proto.InternalMessageInfo

func (m *MemberAddRequest) GetPeerURLs() []string {
	if m != nil {
		return m.PeerURLs
	}
	return nil
}

func (m *MemberAddRequest) GetIsLearner() bool {
	if m != nil {
		return m.IsLearner
	}
	return false
}

type MemberAddResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// member is the member information for the added member.
	Member *Member `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
	// members is a list of all members after adding the new member.
	Members              []*Member `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MemberAddResponse) Reset()         { *m = MemberAddResponse{} }
func (m *MemberAddResponse) String() string { return proto.CompactTextString(m) }
func (*MemberAddResponse) ProtoMessage()    {}
func (*MemberAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{41}
}

func (m *MemberAddResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MemberAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberAddResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MemberAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberAddResponse.Merge(m, src)
}

func (m *MemberAddResponse) XXX_Size() int {
	return m.Size()
}

func (m *MemberAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberAddResponse proto.InternalMessageInfo

func (m *MemberAddResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MemberAddResponse) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *MemberAddResponse) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type MemberRemoveRequest struct {
	// ID is the member ID of the member to remove.
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberRemoveRequest) Reset()         { *m = MemberRemoveRequest{} }
func (m *MemberRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*MemberRemoveRequest) ProtoMessage()    {}
func (*MemberRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{42}
}

func (m *MemberRemoveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MemberRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberRemoveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MemberRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberRemoveRequest.Merge(m, src)
}

func (m *MemberRemoveRequest) XXX_Size() int {
	return m.Size()
}

func (m *MemberRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberRemoveRequest proto.InternalMessageInfo

func (m *MemberRemoveRequest) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type MemberRemoveResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// members is a list of all members after removing the member.
	Members              []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MemberRemoveResponse) Reset()         { *m = MemberRemoveResponse{} }
func (m *MemberRemoveResponse) String() string { return proto.CompactTextString(m) }
func (*MemberRemoveResponse) ProtoMessage()    {}
func (*MemberRemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{43}
}

func (m *MemberRemoveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MemberRemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberRemoveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MemberRemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberRemoveResponse.Merge(m, src)
}

func (m *MemberRemoveResponse) XXX_Size() int {
	return m.Size()
}

func (m *MemberRemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberRemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberRemoveResponse proto.InternalMessageInfo

func (m *MemberRemoveResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MemberRemoveResponse) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type MemberUpdateRequest struct {
	// ID is the member ID of the member to update.
	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// peerURLs is the new list of URLs the member will use to communicate with the cluster.
	PeerURLs             []string `protobuf:"bytes,2,rep,name=peerURLs,proto3" json:"peerURLs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberUpdateRequest) Reset()         { *m = MemberUpdateRequest{} }
func (m *MemberUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*MemberUpdateRequest) ProtoMessage()    {}
func (*MemberUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{44}
}

func (m *MemberUpdateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MemberUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberUpdateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MemberUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberUpdateRequest.Merge(m, src)
}

func (m *MemberUpdateRequest) XXX_Size() int {
	return m.Size()
}

func (m *MemberUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberUpdateRequest proto.InternalMessageInfo

func (m *MemberUpdateRequest) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *MemberUpdateRequest) GetPeerURLs() []string {
	if m != nil {
		return m.PeerURLs
	}
	return nil
}

type MemberUpdateResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// members is a list of all members after updating the member.
	Members              []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MemberUpdateResponse) Reset()         { *m = MemberUpdateResponse{} }
func (m *MemberUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*MemberUpdateResponse) ProtoMessage()    {}
func (*MemberUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{45}
}

func (m *MemberUpdateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MemberUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberUpdateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MemberUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberUpdateResponse.Merge(m, src)
}

func (m *MemberUpdateResponse) XXX_Size() int {
	return m.Size()
}

func (m *MemberUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberUpdateResponse proto.InternalMessageInfo

func (m *MemberUpdateResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MemberUpdateResponse) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type MemberListRequest struct {
	Linearizable         bool     `protobuf:"varint,1,opt,name=linearizable,proto3" json:"linearizable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberListRequest) Reset()         { *m = MemberListRequest{} }
func (m *MemberListRequest) String() string { return proto.CompactTextString(m) }
func (*MemberListRequest) ProtoMessage()    {}
func (*MemberListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{46}
}

func (m *MemberListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MemberListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MemberListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberListRequest.Merge(m, src)
}

func (m *MemberListRequest) XXX_Size() int {
	return m.Size()
}

func (m *MemberListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberListRequest proto.InternalMessageInfo

func (m *MemberListRequest) GetLinearizable() bool {
	if m != nil {
		return m.Linearizable
	}
	return false
}

type MemberListResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// members is a list of all members associated with the cluster.
	Members              []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MemberListResponse) Reset()         { *m = MemberListResponse{} }
func (m *MemberListResponse) String() string { return proto.CompactTextString(m) }
func (*MemberListResponse) ProtoMessage()    {}
func (*MemberListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{47}
}

func (m *MemberListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MemberListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MemberListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberListResponse.Merge(m, src)
}

func (m *MemberListResponse) XXX_Size() int {
	return m.Size()
}

func (m *MemberListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberListResponse proto.InternalMessageInfo

func (m *MemberListResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MemberListResponse) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type MemberPromoteRequest struct {
	// ID is the member ID of the member to promote.
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberPromoteRequest) Reset()         { *m = MemberPromoteRequest{} }
func (m *MemberPromoteRequest) String() string { return proto.CompactTextString(m) }
func (*MemberPromoteRequest) ProtoMessage()    {}
func (*MemberPromoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{48}
}

func (m *MemberPromoteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MemberPromoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberPromoteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MemberPromoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberPromoteRequest.Merge(m, src)
}

func (m *MemberPromoteRequest) XXX_Size() int {
	return m.Size()
}

func (m *MemberPromoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberPromoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberPromoteRequest proto.InternalMessageInfo

func (m *MemberPromoteRequest) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type MemberPromoteResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// members is a list of all members after promoting the member.
	Members              []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MemberPromoteResponse) Reset()         { *m = MemberPromoteResponse{} }
func (m *MemberPromoteResponse) String() string { return proto.CompactTextString(m) }
func (*MemberPromoteResponse) ProtoMessage()    {}
func (*MemberPromoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{49}
}

func (m *MemberPromoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MemberPromoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberPromoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MemberPromoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberPromoteResponse.Merge(m, src)
}

func (m *MemberPromoteResponse) XXX_Size() int {
	return m.Size()
}

func (m *MemberPromoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberPromoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberPromoteResponse proto.InternalMessageInfo

func (m *MemberPromoteResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MemberPromoteResponse) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type DefragmentRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefragmentRequest) Reset()         { *m = DefragmentRequest{} }
func (m *DefragmentRequest) String() string { return proto.CompactTextString(m) }
func (*DefragmentRequest) ProtoMessage()    {}
func (*DefragmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{50}
}

func (m *DefragmentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *DefragmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DefragmentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *DefragmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefragmentRequest.Merge(m, src)
}

func (m *DefragmentRequest) XXX_Size() int {
	return m.Size()
}

func (m *DefragmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DefragmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DefragmentRequest proto.InternalMessageInfo

type DefragmentResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DefragmentResponse) Reset()         { *m = DefragmentResponse{} }
func (m *DefragmentResponse) String() string { return proto.CompactTextString(m) }
func (*DefragmentResponse) ProtoMessage()    {}
func (*DefragmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{51}
}

func (m *DefragmentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *DefragmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DefragmentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *DefragmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefragmentResponse.Merge(m, src)
}

func (m *DefragmentResponse) XXX_Size() int {
	return m.Size()
}

func (m *DefragmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DefragmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DefragmentResponse proto.InternalMessageInfo

func (m *DefragmentResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type MoveLeaderRequest struct {
	// targetID is the node ID for the new leader.
	TargetID             uint64   `protobuf:"varint,1,opt,name=targetID,proto3" json:"targetID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveLeaderRequest) Reset()         { *m = MoveLeaderRequest{} }
func (m *MoveLeaderRequest) String() string { return proto.CompactTextString(m) }
func (*MoveLeaderRequest) ProtoMessage()    {}
func (*MoveLeaderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{52}
}

func (m *MoveLeaderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MoveLeaderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MoveLeaderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MoveLeaderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveLeaderRequest.Merge(m, src)
}

func (m *MoveLeaderRequest) XXX_Size() int {
	return m.Size()
}

func (m *MoveLeaderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveLeaderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MoveLeaderRequest proto.InternalMessageInfo

func (m *MoveLeaderRequest) GetTargetID() uint64 {
	if m != nil {
		return m.TargetID
	}
	return 0
}

type MoveLeaderResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MoveLeaderResponse) Reset()         { *m = MoveLeaderResponse{} }
func (m *MoveLeaderResponse) String() string { return proto.CompactTextString(m) }
func (*MoveLeaderResponse) ProtoMessage()    {}
func (*MoveLeaderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{53}
}

func (m *MoveLeaderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MoveLeaderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MoveLeaderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MoveLeaderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveLeaderResponse.Merge(m, src)
}

func (m *MoveLeaderResponse) XXX_Size() int {
	return m.Size()
}

func (m *MoveLeaderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveLeaderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MoveLeaderResponse proto.InternalMessageInfo

func (m *MoveLeaderResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AlarmRequest struct {
	// action is the kind of alarm request to issue. The action
	// may GET alarm statuses, ACTIVATE an alarm, or DEACTIVATE a
	// raised alarm.
	Action AlarmRequest_AlarmAction `protobuf:"varint,1,opt,name=action,proto3,enum=etcdserverpb.AlarmRequest_AlarmAction" json:"action,omitempty"`
	// memberID is the ID of the member associated with the alarm. If memberID is 0, the
	// alarm request covers all members.
	MemberID uint64 `protobuf:"varint,2,opt,name=memberID,proto3" json:"memberID,omitempty"`
	// alarm is the type of alarm to consider for this request.
	Alarm                AlarmType `protobuf:"varint,3,opt,name=alarm,proto3,enum=etcdserverpb.AlarmType" json:"alarm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AlarmRequest) Reset()         { *m = AlarmRequest{} }
func (m *AlarmRequest) String() string { return proto.CompactTextString(m) }
func (*AlarmRequest) ProtoMessage()    {}
func (*AlarmRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{54}
}

func (m *AlarmRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AlarmRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AlarmRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AlarmRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmRequest.Merge(m, src)
}

func (m *AlarmRequest) XXX_Size() int {
	return m.Size()
}

func (m *AlarmRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmRequest proto.InternalMessageInfo

func (m *AlarmRequest) GetAction() AlarmRequest_AlarmAction {
	if m != nil {
		return m.Action
	}
	return AlarmRequest_GET
}

func (m *AlarmRequest) GetMemberID() uint64 {
	if m != nil {
		return m.MemberID
	}
	return 0
}

func (m *AlarmRequest) GetAlarm() AlarmType {
	if m != nil {
		return m.Alarm
	}
	return AlarmType_NONE
}

type AlarmMember struct {
	// memberID is the ID of the member associated with the raised alarm.
	MemberID uint64 `protobuf:"varint,1,opt,name=memberID,proto3" json:"memberID,omitempty"`
	// alarm is the type of alarm which has been raised.
	Alarm                AlarmType `protobuf:"varint,2,opt,name=alarm,proto3,enum=etcdserverpb.AlarmType" json:"alarm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AlarmMember) Reset()         { *m = AlarmMember{} }
func (m *AlarmMember) String() string { return proto.CompactTextString(m) }
func (*AlarmMember) ProtoMessage()    {}
func (*AlarmMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{55}
}

func (m *AlarmMember) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AlarmMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AlarmMember.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AlarmMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMember.Merge(m, src)
}

func (m *AlarmMember) XXX_Size() int {
	return m.Size()
}

func (m *AlarmMember) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMember.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMember proto.InternalMessageInfo

func (m *AlarmMember) GetMemberID() uint64 {
	if m != nil {
		return m.MemberID
	}
	return 0
}

func (m *AlarmMember) GetAlarm() AlarmType {
	if m != nil {
		return m.Alarm
	}
	return AlarmType_NONE
}

type AlarmResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// alarms is a list of alarms associated with the alarm request.
	Alarms               []*AlarmMember `protobuf:"bytes,2,rep,name=alarms,proto3" json:"alarms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AlarmResponse) Reset()         { *m = AlarmResponse{} }
func (m *AlarmResponse) String() string { return proto.CompactTextString(m) }
func (*AlarmResponse) ProtoMessage()    {}
func (*AlarmResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{56}
}

func (m *AlarmResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AlarmResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AlarmResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AlarmResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmResponse.Merge(m, src)
}

func (m *AlarmResponse) XXX_Size() int {
	return m.Size()
}

func (m *AlarmResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmResponse proto.InternalMessageInfo

func (m *AlarmResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AlarmResponse) GetAlarms() []*AlarmMember {
	if m != nil {
		return m.Alarms
	}
	return nil
}

type DowngradeRequest struct {
	// action is the kind of downgrade request to issue. The action may
	// VALIDATE the target version, DOWNGRADE the cluster version,
	// or CANCEL the current downgrading job.
	Action DowngradeRequest_DowngradeAction `protobuf:"varint,1,opt,name=action,proto3,enum=etcdserverpb.DowngradeRequest_DowngradeAction" json:"action,omitempty"`
	// version is the target version to downgrade.
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DowngradeRequest) Reset()         { *m = DowngradeRequest{} }
func (m *DowngradeRequest) String() string { return proto.CompactTextString(m) }
func (*DowngradeRequest) ProtoMessage()    {}
func (*DowngradeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{57}
}

func (m *DowngradeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *DowngradeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DowngradeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *DowngradeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DowngradeRequest.Merge(m, src)
}

func (m *DowngradeRequest) XXX_Size() int {
	return m.Size()
}

func (m *DowngradeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DowngradeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DowngradeRequest proto.InternalMessageInfo

func (m *DowngradeRequest) GetAction() DowngradeRequest_DowngradeAction {
	if m != nil {
		return m.Action
	}
	return DowngradeRequest_VALIDATE
}

func (m *DowngradeRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type DowngradeResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// version is the current cluster version.
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DowngradeResponse) Reset()         { *m = DowngradeResponse{} }
func (m *DowngradeResponse) String() string { return proto.CompactTextString(m) }
func (*DowngradeResponse) ProtoMessage()    {}
func (*DowngradeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{58}
}

func (m *DowngradeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *DowngradeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DowngradeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *DowngradeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DowngradeResponse.Merge(m, src)
}

func (m *DowngradeResponse) XXX_Size() int {
	return m.Size()
}

func (m *DowngradeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DowngradeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DowngradeResponse proto.InternalMessageInfo

func (m *DowngradeResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DowngradeResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{59}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}

func (m *StatusRequest) XXX_Size() int {
	return m.Size()
}

func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// version is the cluster protocol version used by the responding member.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// dbSize is the size of the backend database physically allocated, in bytes, of the responding member.
	DbSize int64 `protobuf:"varint,3,opt,name=dbSize,proto3" json:"dbSize,omitempty"`
	// leader is the member ID which the responding member believes is the current leader.
	Leader uint64 `protobuf:"varint,4,opt,name=leader,proto3" json:"leader,omitempty"`
	// raftIndex is the current raft committed index of the responding member.
	RaftIndex uint64 `protobuf:"varint,5,opt,name=raftIndex,proto3" json:"raftIndex,omitempty"`
	// raftTerm is the current raft term of the responding member.
	RaftTerm uint64 `protobuf:"varint,6,opt,name=raftTerm,proto3" json:"raftTerm,omitempty"`
	// raftAppliedIndex is the current raft applied index of the responding member.
	RaftAppliedIndex uint64 `protobuf:"varint,7,opt,name=raftAppliedIndex,proto3" json:"raftAppliedIndex,omitempty"`
	// errors contains alarm/health information and status.
	Errors []string `protobuf:"bytes,8,rep,name=errors,proto3" json:"errors,omitempty"`
	// dbSizeInUse is the size of the backend database logically in use, in bytes, of the responding member.
	DbSizeInUse int64 `protobuf:"varint,9,opt,name=dbSizeInUse,proto3" json:"dbSizeInUse,omitempty"`
	// isLearner indicates if the member is raft learner.
	IsLearner            bool     `protobuf:"varint,10,opt,name=isLearner,proto3" json:"isLearner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{60}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}

func (m *StatusResponse) XXX_Size() int {
	return m.Size()
}

func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *StatusResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *StatusResponse) GetDbSize() int64 {
	if m != nil {
		return m.DbSize
	}
	return 0
}

func (m *StatusResponse) GetLeader() uint64 {
	if m != nil {
		return m.Leader
	}
	return 0
}

func (m *StatusResponse) GetRaftIndex() uint64 {
	if m != nil {
		return m.RaftIndex
	}
	return 0
}

func (m *StatusResponse) GetRaftTerm() uint64 {
	if m != nil {
		return m.RaftTerm
	}
	return 0
}

func (m *StatusResponse) GetRaftAppliedIndex() uint64 {
	if m != nil {
		return m.RaftAppliedIndex
	}
	return 0
}

func (m *StatusResponse) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *StatusResponse) GetDbSizeInUse() int64 {
	if m != nil {
		return m.DbSizeInUse
	}
	return 0
}

func (m *StatusResponse) GetIsLearner() bool {
	if m != nil {
		return m.IsLearner
	}
	return false
}

type AuthEnableRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthEnableRequest) Reset()         { *m = AuthEnableRequest{} }
func (m *AuthEnableRequest) String() string { return proto.CompactTextString(m) }
func (*AuthEnableRequest) ProtoMessage()    {}
func (*AuthEnableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{61}
}

func (m *AuthEnableRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthEnableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthEnableRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthEnableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthEnableRequest.Merge(m, src)
}

func (m *AuthEnableRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthEnableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthEnableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthEnableRequest proto.InternalMessageInfo

type AuthDisableRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthDisableRequest) Reset()         { *m = AuthDisableRequest{} }
func (m *AuthDisableRequest) String() string { return proto.CompactTextString(m) }
func (*AuthDisableRequest) ProtoMessage()    {}
func (*AuthDisableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{62}
}

func (m *AuthDisableRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthDisableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthDisableRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthDisableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthDisableRequest.Merge(m, src)
}

func (m *AuthDisableRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthDisableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthDisableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthDisableRequest proto.InternalMessageInfo

type AuthStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthStatusRequest) Reset()         { *m = AuthStatusRequest{} }
func (m *AuthStatusRequest) String() string { return proto.CompactTextString(m) }
func (*AuthStatusRequest) ProtoMessage()    {}
func (*AuthStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{63}
}

func (m *AuthStatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthStatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthStatusRequest.Merge(m, src)
}

func (m *AuthStatusRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthStatusRequest proto.InternalMessageInfo

type AuthenticateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{64}
}

func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(m, src)
}

func (m *AuthenticateRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

func (m *AuthenticateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthenticateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthUserAddRequest struct {
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Options              *authpb.UserAddOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	HashedPassword       string                 `protobuf:"bytes,4,opt,name=hashedPassword,proto3" json:"hashedPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AuthUserAddRequest) Reset()         { *m = AuthUserAddRequest{} }
func (m *AuthUserAddRequest) String() string { return proto.CompactTextString(m) }
func (*AuthUserAddRequest) ProtoMessage()    {}
func (*AuthUserAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{65}
}

func (m *AuthUserAddRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserAddRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserAddRequest.Merge(m, src)
}

func (m *AuthUserAddRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserAddRequest proto.InternalMessageInfo

func (m *AuthUserAddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthUserAddRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthUserAddRequest) GetOptions() *authpb.UserAddOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *AuthUserAddRequest) GetHashedPassword() string {
	if m != nil {
		return m.HashedPassword
	}
	return ""
}

type AuthUserGetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUserGetRequest) Reset()         { *m = AuthUserGetRequest{} }
func (m *AuthUserGetRequest) String() string { return proto.CompactTextString(m) }
func (*AuthUserGetRequest) ProtoMessage()    {}
func (*AuthUserGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{66}
}

func (m *AuthUserGetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserGetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserGetRequest.Merge(m, src)
}

func (m *AuthUserGetRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserGetRequest proto.InternalMessageInfo

func (m *AuthUserGetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AuthUserDeleteRequest struct {
	// name is the name of the user to delete.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUserDeleteRequest) Reset()         { *m = AuthUserDeleteRequest{} }
func (m *AuthUserDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*AuthUserDeleteRequest) ProtoMessage()    {}
func (*AuthUserDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{67}
}

func (m *AuthUserDeleteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserDeleteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserDeleteRequest.Merge(m, src)
}

func (m *AuthUserDeleteRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserDeleteRequest proto.InternalMessageInfo

func (m *AuthUserDeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AuthUserChangePasswordRequest struct {
	// name is the name of the user whose password is being changed.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// password is the new password for the user. Note that this field will be removed in the API layer.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// hashedPassword is the new password for the user. Note that this field will be initialized in the API layer.
	HashedPassword       string   `protobuf:"bytes,3,opt,name=hashedPassword,proto3" json:"hashedPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUserChangePasswordRequest) Reset()         { *m = AuthUserChangePasswordRequest{} }
func (m *AuthUserChangePasswordRequest) String() string { return proto.CompactTextString(m) }
func (*AuthUserChangePasswordRequest) ProtoMessage()    {}
func (*AuthUserChangePasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{68}
}

func (m *AuthUserChangePasswordRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserChangePasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserChangePasswordRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserChangePasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserChangePasswordRequest.Merge(m, src)
}

func (m *AuthUserChangePasswordRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserChangePasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserChangePasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserChangePasswordRequest proto.InternalMessageInfo

func (m *AuthUserChangePasswordRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthUserChangePasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthUserChangePasswordRequest) GetHashedPassword() string {
	if m != nil {
		return m.HashedPassword
	}
	return ""
}

type AuthUserGrantRoleRequest struct {
	// user is the name of the user which should be granted a given role.
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// role is the name of the role to grant to the user.
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUserGrantRoleRequest) Reset()         { *m = AuthUserGrantRoleRequest{} }
func (m *AuthUserGrantRoleRequest) String() string { return proto.CompactTextString(m) }
func (*AuthUserGrantRoleRequest) ProtoMessage()    {}
func (*AuthUserGrantRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{69}
}

func (m *AuthUserGrantRoleRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserGrantRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserGrantRoleRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserGrantRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserGrantRoleRequest.Merge(m, src)
}

func (m *AuthUserGrantRoleRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserGrantRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserGrantRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserGrantRoleRequest proto.InternalMessageInfo

func (m *AuthUserGrantRoleRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *AuthUserGrantRoleRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type AuthUserRevokeRoleRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUserRevokeRoleRequest) Reset()         { *m = AuthUserRevokeRoleRequest{} }
func (m *AuthUserRevokeRoleRequest) String() string { return proto.CompactTextString(m) }
func (*AuthUserRevokeRoleRequest) ProtoMessage()    {}
func (*AuthUserRevokeRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{70}
}

func (m *AuthUserRevokeRoleRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserRevokeRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserRevokeRoleRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserRevokeRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserRevokeRoleRequest.Merge(m, src)
}

func (m *AuthUserRevokeRoleRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserRevokeRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserRevokeRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserRevokeRoleRequest proto.InternalMessageInfo

func (m *AuthUserRevokeRoleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthUserRevokeRoleRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type AuthRoleAddRequest struct {
	// name is the name of the role to add to the authentication system.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRoleAddRequest) Reset()         { *m = AuthRoleAddRequest{} }
func (m *AuthRoleAddRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRoleAddRequest) ProtoMessage()    {}
func (*AuthRoleAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{71}
}

func (m *AuthRoleAddRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleAddRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleAddRequest.Merge(m, src)
}

func (m *AuthRoleAddRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleAddRequest proto.InternalMessageInfo

func (m *AuthRoleAddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AuthRoleGetRequest struct {
	Role                 string   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRoleGetRequest) Reset()         { *m = AuthRoleGetRequest{} }
func (m *AuthRoleGetRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRoleGetRequest) ProtoMessage()    {}
func (*AuthRoleGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{72}
}

func (m *AuthRoleGetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleGetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleGetRequest.Merge(m, src)
}

func (m *AuthRoleGetRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleGetRequest proto.InternalMessageInfo

func (m *AuthRoleGetRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type AuthUserListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUserListRequest) Reset()         { *m = AuthUserListRequest{} }
func (m *AuthUserListRequest) String() string { return proto.CompactTextString(m) }
func (*AuthUserListRequest) ProtoMessage()    {}
func (*AuthUserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{73}
}

func (m *AuthUserListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserListRequest.Merge(m, src)
}

func (m *AuthUserListRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserListRequest proto.InternalMessageInfo

type AuthRoleListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRoleListRequest) Reset()         { *m = AuthRoleListRequest{} }
func (m *AuthRoleListRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRoleListRequest) ProtoMessage()    {}
func (*AuthRoleListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{74}
}

func (m *AuthRoleListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleListRequest.Merge(m, src)
}

func (m *AuthRoleListRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleListRequest proto.InternalMessageInfo

type AuthRoleDeleteRequest struct {
	Role                 string   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRoleDeleteRequest) Reset()         { *m = AuthRoleDeleteRequest{} }
func (m *AuthRoleDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRoleDeleteRequest) ProtoMessage()    {}
func (*AuthRoleDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{75}
}

func (m *AuthRoleDeleteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleDeleteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleDeleteRequest.Merge(m, src)
}

func (m *AuthRoleDeleteRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleDeleteRequest proto.InternalMessageInfo

func (m *AuthRoleDeleteRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type AuthRoleGrantPermissionRequest struct {
	// name is the name of the role which will be granted the permission.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// perm is the permission to grant to the role.
	Perm                 *authpb.Permission `protobuf:"bytes,2,opt,name=perm,proto3" json:"perm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AuthRoleGrantPermissionRequest) Reset()         { *m = AuthRoleGrantPermissionRequest{} }
func (m *AuthRoleGrantPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRoleGrantPermissionRequest) ProtoMessage()    {}
func (*AuthRoleGrantPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{76}
}

func (m *AuthRoleGrantPermissionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleGrantPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleGrantPermissionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleGrantPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleGrantPermissionRequest.Merge(m, src)
}

func (m *AuthRoleGrantPermissionRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleGrantPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleGrantPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleGrantPermissionRequest proto.InternalMessageInfo

func (m *AuthRoleGrantPermissionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthRoleGrantPermissionRequest) GetPerm() *authpb.Permission {
	if m != nil {
		return m.Perm
	}
	return nil
}

type AuthRoleRevokePermissionRequest struct {
	Role                 string   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	RangeEnd             []byte   `protobuf:"bytes,3,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRoleRevokePermissionRequest) Reset()         { *m = AuthRoleRevokePermissionRequest{} }
func (m *AuthRoleRevokePermissionRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRoleRevokePermissionRequest) ProtoMessage()    {}
func (*AuthRoleRevokePermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{77}
}

func (m *AuthRoleRevokePermissionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleRevokePermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleRevokePermissionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleRevokePermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleRevokePermissionRequest.Merge(m, src)
}

func (m *AuthRoleRevokePermissionRequest) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleRevokePermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleRevokePermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleRevokePermissionRequest proto.InternalMessageInfo

func (m *AuthRoleRevokePermissionRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *AuthRoleRevokePermissionRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *AuthRoleRevokePermissionRequest) GetRangeEnd() []byte {
	if m != nil {
		return m.RangeEnd
	}
	return nil
}

type AuthEnableResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthEnableResponse) Reset()         { *m = AuthEnableResponse{} }
func (m *AuthEnableResponse) String() string { return proto.CompactTextString(m) }
func (*AuthEnableResponse) ProtoMessage()    {}
func (*AuthEnableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{78}
}

func (m *AuthEnableResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthEnableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthEnableResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthEnableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthEnableResponse.Merge(m, src)
}

func (m *AuthEnableResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthEnableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthEnableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthEnableResponse proto.InternalMessageInfo

func (m *AuthEnableResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthDisableResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthDisableResponse) Reset()         { *m = AuthDisableResponse{} }
func (m *AuthDisableResponse) String() string { return proto.CompactTextString(m) }
func (*AuthDisableResponse) ProtoMessage()    {}
func (*AuthDisableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{79}
}

func (m *AuthDisableResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthDisableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthDisableResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthDisableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthDisableResponse.Merge(m, src)
}

func (m *AuthDisableResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthDisableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthDisableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthDisableResponse proto.InternalMessageInfo

func (m *AuthDisableResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthStatusResponse struct {
	Header  *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Enabled bool            `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// authRevision is the current revision of auth store
	AuthRevision         uint64   `protobuf:"varint,3,opt,name=authRevision,proto3" json:"authRevision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthStatusResponse) Reset()         { *m = AuthStatusResponse{} }
func (m *AuthStatusResponse) String() string { return proto.CompactTextString(m) }
func (*AuthStatusResponse) ProtoMessage()    {}
func (*AuthStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{80}
}

func (m *AuthStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthStatusResponse.Merge(m, src)
}

func (m *AuthStatusResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthStatusResponse proto.InternalMessageInfo

func (m *AuthStatusResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AuthStatusResponse) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *AuthStatusResponse) GetAuthRevision() uint64 {
	if m != nil {
		return m.AuthRevision
	}
	return 0
}

type AuthenticateResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// token is an authorized token that can be used in succeeding RPCs
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{81}
}

func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(m, src)
}

func (m *AuthenticateResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AuthenticateResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthUserAddResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthUserAddResponse) Reset()         { *m = AuthUserAddResponse{} }
func (m *AuthUserAddResponse) String() string { return proto.CompactTextString(m) }
func (*AuthUserAddResponse) ProtoMessage()    {}
func (*AuthUserAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{82}
}

func (m *AuthUserAddResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserAddResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserAddResponse.Merge(m, src)
}

func (m *AuthUserAddResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserAddResponse proto.InternalMessageInfo

func (m *AuthUserAddResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthUserGetResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Roles                []string        `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthUserGetResponse) Reset()         { *m = AuthUserGetResponse{} }
func (m *AuthUserGetResponse) String() string { return proto.CompactTextString(m) }
func (*AuthUserGetResponse) ProtoMessage()    {}
func (*AuthUserGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{83}
}

func (m *AuthUserGetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserGetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserGetResponse.Merge(m, src)
}

func (m *AuthUserGetResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserGetResponse proto.InternalMessageInfo

func (m *AuthUserGetResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AuthUserGetResponse) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type AuthUserDeleteResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthUserDeleteResponse) Reset()         { *m = AuthUserDeleteResponse{} }
func (m *AuthUserDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*AuthUserDeleteResponse) ProtoMessage()    {}
func (*AuthUserDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{84}
}

func (m *AuthUserDeleteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserDeleteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserDeleteResponse.Merge(m, src)
}

func (m *AuthUserDeleteResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserDeleteResponse proto.InternalMessageInfo

func (m *AuthUserDeleteResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthUserChangePasswordResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthUserChangePasswordResponse) Reset()         { *m = AuthUserChangePasswordResponse{} }
func (m *AuthUserChangePasswordResponse) String() string { return proto.CompactTextString(m) }
func (*AuthUserChangePasswordResponse) ProtoMessage()    {}
func (*AuthUserChangePasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{85}
}

func (m *AuthUserChangePasswordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserChangePasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserChangePasswordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserChangePasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserChangePasswordResponse.Merge(m, src)
}

func (m *AuthUserChangePasswordResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserChangePasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserChangePasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserChangePasswordResponse proto.InternalMessageInfo

func (m *AuthUserChangePasswordResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthUserGrantRoleResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthUserGrantRoleResponse) Reset()         { *m = AuthUserGrantRoleResponse{} }
func (m *AuthUserGrantRoleResponse) String() string { return proto.CompactTextString(m) }
func (*AuthUserGrantRoleResponse) ProtoMessage()    {}
func (*AuthUserGrantRoleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{86}
}

func (m *AuthUserGrantRoleResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserGrantRoleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserGrantRoleResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserGrantRoleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserGrantRoleResponse.Merge(m, src)
}

func (m *AuthUserGrantRoleResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserGrantRoleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserGrantRoleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserGrantRoleResponse proto.InternalMessageInfo

func (m *AuthUserGrantRoleResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthUserRevokeRoleResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthUserRevokeRoleResponse) Reset()         { *m = AuthUserRevokeRoleResponse{} }
func (m *AuthUserRevokeRoleResponse) String() string { return proto.CompactTextString(m) }
func (*AuthUserRevokeRoleResponse) ProtoMessage()    {}
func (*AuthUserRevokeRoleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{87}
}

func (m *AuthUserRevokeRoleResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserRevokeRoleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserRevokeRoleResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserRevokeRoleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserRevokeRoleResponse.Merge(m, src)
}

func (m *AuthUserRevokeRoleResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserRevokeRoleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserRevokeRoleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserRevokeRoleResponse proto.InternalMessageInfo

func (m *AuthUserRevokeRoleResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthRoleAddResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthRoleAddResponse) Reset()         { *m = AuthRoleAddResponse{} }
func (m *AuthRoleAddResponse) String() string { return proto.CompactTextString(m) }
func (*AuthRoleAddResponse) ProtoMessage()    {}
func (*AuthRoleAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{88}
}

func (m *AuthRoleAddResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleAddResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleAddResponse.Merge(m, src)
}

func (m *AuthRoleAddResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleAddResponse proto.InternalMessageInfo

func (m *AuthRoleAddResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthRoleGetResponse struct {
	Header               *ResponseHeader      `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Perm                 []*authpb.Permission `protobuf:"bytes,2,rep,name=perm,proto3" json:"perm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AuthRoleGetResponse) Reset()         { *m = AuthRoleGetResponse{} }
func (m *AuthRoleGetResponse) String() string { return proto.CompactTextString(m) }
func (*AuthRoleGetResponse) ProtoMessage()    {}
func (*AuthRoleGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{89}
}

func (m *AuthRoleGetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleGetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleGetResponse.Merge(m, src)
}

func (m *AuthRoleGetResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleGetResponse proto.InternalMessageInfo

func (m *AuthRoleGetResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AuthRoleGetResponse) GetPerm() []*authpb.Permission {
	if m != nil {
		return m.Perm
	}
	return nil
}

type AuthRoleListResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Roles                []string        `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthRoleListResponse) Reset()         { *m = AuthRoleListResponse{} }
func (m *AuthRoleListResponse) String() string { return proto.CompactTextString(m) }
func (*AuthRoleListResponse) ProtoMessage()    {}
func (*AuthRoleListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{90}
}

func (m *AuthRoleListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleListResponse.Merge(m, src)
}

func (m *AuthRoleListResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleListResponse proto.InternalMessageInfo

func (m *AuthRoleListResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AuthRoleListResponse) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type AuthUserListResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Users                []string        `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthUserListResponse) Reset()         { *m = AuthUserListResponse{} }
func (m *AuthUserListResponse) String() string { return proto.CompactTextString(m) }
func (*AuthUserListResponse) ProtoMessage()    {}
func (*AuthUserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{91}
}

func (m *AuthUserListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthUserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthUserListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthUserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserListResponse.Merge(m, src)
}

func (m *AuthUserListResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthUserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserListResponse proto.InternalMessageInfo

func (m *AuthUserListResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AuthUserListResponse) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

type AuthRoleDeleteResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthRoleDeleteResponse) Reset()         { *m = AuthRoleDeleteResponse{} }
func (m *AuthRoleDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*AuthRoleDeleteResponse) ProtoMessage()    {}
func (*AuthRoleDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{92}
}

func (m *AuthRoleDeleteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleDeleteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleDeleteResponse.Merge(m, src)
}

func (m *AuthRoleDeleteResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleDeleteResponse proto.InternalMessageInfo

func (m *AuthRoleDeleteResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthRoleGrantPermissionResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthRoleGrantPermissionResponse) Reset()         { *m = AuthRoleGrantPermissionResponse{} }
func (m *AuthRoleGrantPermissionResponse) String() string { return proto.CompactTextString(m) }
func (*AuthRoleGrantPermissionResponse) ProtoMessage()    {}
func (*AuthRoleGrantPermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{93}
}

func (m *AuthRoleGrantPermissionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleGrantPermissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleGrantPermissionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleGrantPermissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleGrantPermissionResponse.Merge(m, src)
}

func (m *AuthRoleGrantPermissionResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleGrantPermissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleGrantPermissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleGrantPermissionResponse proto.InternalMessageInfo

func (m *AuthRoleGrantPermissionResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthRoleRevokePermissionResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthRoleRevokePermissionResponse) Reset()         { *m = AuthRoleRevokePermissionResponse{} }
func (m *AuthRoleRevokePermissionResponse) String() string { return proto.CompactTextString(m) }
func (*AuthRoleRevokePermissionResponse) ProtoMessage()    {}
func (*AuthRoleRevokePermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{94}
}

func (m *AuthRoleRevokePermissionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *AuthRoleRevokePermissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthRoleRevokePermissionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *AuthRoleRevokePermissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRoleRevokePermissionResponse.Merge(m, src)
}

func (m *AuthRoleRevokePermissionResponse) XXX_Size() int {
	return m.Size()
}

func (m *AuthRoleRevokePermissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRoleRevokePermissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRoleRevokePermissionResponse proto.InternalMessageInfo

func (m *AuthRoleRevokePermissionResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterEnum("etcdserverpb.AlarmType", AlarmType_name, AlarmType_value)
	proto.RegisterEnum("etcdserverpb.RangeRequest_SortOrder", RangeRequest_SortOrder_name, RangeRequest_SortOrder_value)
	proto.RegisterEnum("etcdserverpb.RangeRequest_SortTarget", RangeRequest_SortTarget_name, RangeRequest_SortTarget_value)
	proto.RegisterEnum("etcdserverpb.Compare_CompareResult", Compare_CompareResult_name, Compare_CompareResult_value)
	proto.RegisterEnum("etcdserverpb.Compare_CompareTarget", Compare_CompareTarget_name, Compare_CompareTarget_value)
	proto.RegisterEnum("etcdserverpb.WatchCreateRequest_FilterType", WatchCreateRequest_FilterType_name, WatchCreateRequest_FilterType_value)
	proto.RegisterEnum("etcdserverpb.AlarmRequest_AlarmAction", AlarmRequest_AlarmAction_name, AlarmRequest_AlarmAction_value)
	proto.RegisterEnum("etcdserverpb.DowngradeRequest_DowngradeAction", DowngradeRequest_DowngradeAction_name, DowngradeRequest_DowngradeAction_value)
	proto.RegisterType((*ResponseHeader)(nil), "etcdserverpb.ResponseHeader")
	proto.RegisterType((*RangeRequest)(nil), "etcdserverpb.RangeRequest")
	proto.RegisterType((*RangeResponse)(nil), "etcdserverpb.RangeResponse")
	proto.RegisterType((*PutRequest)(nil), "etcdserverpb.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "etcdserverpb.PutResponse")
	proto.RegisterType((*DeleteRangeRequest)(nil), "etcdserverpb.DeleteRangeRequest")
	proto.RegisterType((*DeleteRangeResponse)(nil), "etcdserverpb.DeleteRangeResponse")
	proto.RegisterType((*RequestOp)(nil), "etcdserverpb.RequestOp")
	proto.RegisterType((*ResponseOp)(nil), "etcdserverpb.ResponseOp")
	proto.RegisterType((*Compare)(nil), "etcdserverpb.Compare")
	proto.RegisterType((*TxnRequest)(nil), "etcdserverpb.TxnRequest")
	proto.RegisterType((*TxnResponse)(nil), "etcdserverpb.TxnResponse")
	proto.RegisterType((*CompactionRequest)(nil), "etcdserverpb.CompactionRequest")
	proto.RegisterType((*CompactionResponse)(nil), "etcdserverpb.CompactionResponse")
	proto.RegisterType((*HashRequest)(nil), "etcdserverpb.HashRequest")
	proto.RegisterType((*HashKVRequest)(nil), "etcdserverpb.HashKVRequest")
	proto.RegisterType((*HashKVResponse)(nil), "etcdserverpb.HashKVResponse")
	proto.RegisterType((*HashResponse)(nil), "etcdserverpb.HashResponse")
	proto.RegisterType((*SnapshotRequest)(nil), "etcdserverpb.SnapshotRequest")
	proto.RegisterType((*SnapshotResponse)(nil), "etcdserverpb.SnapshotResponse")
	proto.RegisterType((*WatchRequest)(nil), "etcdserverpb.WatchRequest")
	proto.RegisterType((*WatchCreateRequest)(nil), "etcdserverpb.WatchCreateRequest")
	proto.RegisterType((*WatchCancelRequest)(nil), "etcdserverpb.WatchCancelRequest")
	proto.RegisterType((*WatchProgressRequest)(nil), "etcdserverpb.WatchProgressRequest")
	proto.RegisterType((*WatchResponse)(nil), "etcdserverpb.WatchResponse")
	proto.RegisterType((*LeaseGrantRequest)(nil), "etcdserverpb.LeaseGrantRequest")
	proto.RegisterType((*LeaseGrantResponse)(nil), "etcdserverpb.LeaseGrantResponse")
	proto.RegisterType((*LeaseRevokeRequest)(nil), "etcdserverpb.LeaseRevokeRequest")
	proto.RegisterType((*LeaseRevokeResponse)(nil), "etcdserverpb.LeaseRevokeResponse")
	proto.RegisterType((*LeaseCheckpoint)(nil), "etcdserverpb.LeaseCheckpoint")
	proto.RegisterType((*LeaseCheckpointRequest)(nil), "etcdserverpb.LeaseCheckpointRequest")
	proto.RegisterType((*LeaseCheckpointResponse)(nil), "etcdserverpb.LeaseCheckpointResponse")
	proto.RegisterType((*LeaseKeepAliveRequest)(nil), "etcdserverpb.LeaseKeepAliveRequest")
	proto.RegisterType((*LeaseKeepAliveResponse)(nil), "etcdserverpb.LeaseKeepAliveResponse")
	proto.RegisterType((*LeaseTimeToLiveRequest)(nil), "etcdserverpb.LeaseTimeToLiveRequest")
	proto.RegisterType((*LeaseTimeToLiveResponse)(nil), "etcdserverpb.LeaseTimeToLiveResponse")
	proto.RegisterType((*LeaseLeasesRequest)(nil), "etcdserverpb.LeaseLeasesRequest")
	proto.RegisterType((*LeaseStatus)(nil), "etcdserverpb.LeaseStatus")
	proto.RegisterType((*LeaseLeasesResponse)(nil), "etcdserverpb.LeaseLeasesResponse")
	proto.RegisterType((*Member)(nil), "etcdserverpb.Member")
	proto.RegisterType((*MemberAddRequest)(nil), "etcdserverpb.MemberAddRequest")
	proto.RegisterType((*MemberAddResponse)(nil), "etcdserverpb.MemberAddResponse")
	proto.RegisterType((*MemberRemoveRequest)(nil), "etcdserverpb.MemberRemoveRequest")
	proto.RegisterType((*MemberRemoveResponse)(nil), "etcdserverpb.MemberRemoveResponse")
	proto.RegisterType((*MemberUpdateRequest)(nil), "etcdserverpb.MemberUpdateRequest")
	proto.RegisterType((*MemberUpdateResponse)(nil), "etcdserverpb.MemberUpdateResponse")
	proto.RegisterType((*MemberListRequest)(nil), "etcdserverpb.MemberListRequest")
	proto.RegisterType((*MemberListResponse)(nil), "etcdserverpb.MemberListResponse")
	proto.RegisterType((*MemberPromoteRequest)(nil), "etcdserverpb.MemberPromoteRequest")
	proto.RegisterType((*MemberPromoteResponse)(nil), "etcdserverpb.MemberPromoteResponse")
	proto.RegisterType((*DefragmentRequest)(nil), "etcdserverpb.DefragmentRequest")
	proto.RegisterType((*DefragmentResponse)(nil), "etcdserverpb.DefragmentResponse")
	proto.RegisterType((*MoveLeaderRequest)(nil), "etcdserverpb.MoveLeaderRequest")
	proto.RegisterType((*MoveLeaderResponse)(nil), "etcdserverpb.MoveLeaderResponse")
	proto.RegisterType((*AlarmRequest)(nil), "etcdserverpb.AlarmRequest")
	proto.RegisterType((*AlarmMember)(nil), "etcdserverpb.AlarmMember")
	proto.RegisterType((*AlarmResponse)(nil), "etcdserverpb.AlarmResponse")
	proto.RegisterType((*DowngradeRequest)(nil), "etcdserverpb.DowngradeRequest")
	proto.RegisterType((*DowngradeResponse)(nil), "etcdserverpb.DowngradeResponse")
	proto.RegisterType((*StatusRequest)(nil), "etcdserverpb.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "etcdserverpb.StatusResponse")
	proto.RegisterType((*AuthEnableRequest)(nil), "etcdserverpb.AuthEnableRequest")
	proto.RegisterType((*AuthDisableRequest)(nil), "etcdserverpb.AuthDisableRequest")
	proto.RegisterType((*AuthStatusRequest)(nil), "etcdserverpb.AuthStatusRequest")
	proto.RegisterType((*AuthenticateRequest)(nil), "etcdserverpb.AuthenticateRequest")
	proto.RegisterType((*AuthUserAddRequest)(nil), "etcdserverpb.AuthUserAddRequest")
	proto.RegisterType((*AuthUserGetRequest)(nil), "etcdserverpb.AuthUserGetRequest")
	proto.RegisterType((*AuthUserDeleteRequest)(nil), "etcdserverpb.AuthUserDeleteRequest")
	proto.RegisterType((*AuthUserChangePasswordRequest)(nil), "etcdserverpb.AuthUserChangePasswordRequest")
	proto.RegisterType((*AuthUserGrantRoleRequest)(nil), "etcdserverpb.AuthUserGrantRoleRequest")
	proto.RegisterType((*AuthUserRevokeRoleRequest)(nil), "etcdserverpb.AuthUserRevokeRoleRequest")
	proto.RegisterType((*AuthRoleAddRequest)(nil), "etcdserverpb.AuthRoleAddRequest")
	proto.RegisterType((*AuthRoleGetRequest)(nil), "etcdserverpb.AuthRoleGetRequest")
	proto.RegisterType((*AuthUserListRequest)(nil), "etcdserverpb.AuthUserListRequest")
	proto.RegisterType((*AuthRoleListRequest)(nil), "etcdserverpb.AuthRoleListRequest")
	proto.RegisterType((*AuthRoleDeleteRequest)(nil), "etcdserverpb.AuthRoleDeleteRequest")
	proto.RegisterType((*AuthRoleGrantPermissionRequest)(nil), "etcdserverpb.AuthRoleGrantPermissionRequest")
	proto.RegisterType((*AuthRoleRevokePermissionRequest)(nil), "etcdserverpb.AuthRoleRevokePermissionRequest")
	proto.RegisterType((*AuthEnableResponse)(nil), "etcdserverpb.AuthEnableResponse")
	proto.RegisterType((*AuthDisableResponse)(nil), "etcdserverpb.AuthDisableResponse")
	proto.RegisterType((*AuthStatusResponse)(nil), "etcdserverpb.AuthStatusResponse")
	proto.RegisterType((*AuthenticateResponse)(nil), "etcdserverpb.AuthenticateResponse")
	proto.RegisterType((*AuthUserAddResponse)(nil), "etcdserverpb.AuthUserAddResponse")
	proto.RegisterType((*AuthUserGetResponse)(nil), "etcdserverpb.AuthUserGetResponse")
	proto.RegisterType((*AuthUserDeleteResponse)(nil), "etcdserverpb.AuthUserDeleteResponse")
	proto.RegisterType((*AuthUserChangePasswordResponse)(nil), "etcdserverpb.AuthUserChangePasswordResponse")
	proto.RegisterType((*AuthUserGrantRoleResponse)(nil), "etcdserverpb.AuthUserGrantRoleResponse")
	proto.RegisterType((*AuthUserRevokeRoleResponse)(nil), "etcdserverpb.AuthUserRevokeRoleResponse")
	proto.RegisterType((*AuthRoleAddResponse)(nil), "etcdserverpb.AuthRoleAddResponse")
	proto.RegisterType((*AuthRoleGetResponse)(nil), "etcdserverpb.AuthRoleGetResponse")
	proto.RegisterType((*AuthRoleListResponse)(nil), "etcdserverpb.AuthRoleListResponse")
	proto.RegisterType((*AuthUserListResponse)(nil), "etcdserverpb.AuthUserListResponse")
	proto.RegisterType((*AuthRoleDeleteResponse)(nil), "etcdserverpb.AuthRoleDeleteResponse")
	proto.RegisterType((*AuthRoleGrantPermissionResponse)(nil), "etcdserverpb.AuthRoleGrantPermissionResponse")
	proto.RegisterType((*AuthRoleRevokePermissionResponse)(nil), "etcdserverpb.AuthRoleRevokePermissionResponse")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 4107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x5b, 0x5b, 0x73, 0x1b, 0xc9,
	0x75, 0xe6, 0x00, 0xc4, 0xed, 0xe0, 0x42, 0xb0, 0x79, 0x11, 0x84, 0x95, 0x28, 0x6e, 0x6b, 0xa5,
	0xe5, 0x4a, 0xbb, 0xc4, 0x9a, 0xb6, 0xb3, 0x55, 0x4a, 0xe2, 0x18, 0x22, 0xb1, 0x12, 0x97, 0x14,
	0xc9, 0x1d, 0x42, 0xda, 0x4b, 0xb9, 0xc2, 0x1a, 0x02, 0x2d, 0x72, 0x42, 0x60, 0x06, 0x9e, 0x19,
	0x40, 0xe4, 0xe6, 0xe2, 0x94, 0xcb, 0x71, 0x25, 0xaf, 0x76, 0x55, 0x2a, 0x79, 0x48, 0x5e, 0x52,
	0x29, 0x97, 0x1f, 0xfc, 0x9c, 0xbf, 0x90, 0xa7, 0x5c, 0x2a, 0x7f, 0x20, 0xb5, 0xf1, 0x4b, 0xf2,
	0x23, 0x52, 0xae, 0xbe, 0xcd, 0xf4, 0xdc, 0x40, 0xd9, 0xd8, 0xdd, 0x17, 0x11, 0x7d, 0xfa, 0xf4,
	0xf9, 0x4e, 0x9f, 0xee, 0x3e, 0xe7, 0xf4, 0xe9, 0x11, 0x94, 0x9c, 0x51, 0x6f, 0x73, 0xe4, 0xd8,
	0x9e, 0x8d, 0x2a, 0xc4, 0xeb, 0xf5, 0x5d, 0xe2, 0x4c, 0x88, 0x33, 0x3a, 0x6d, 0x2e, 0x9f, 0xd9,
	0x67, 0x36, 0xeb, 0x68, 0xd1, 0x5f, 0x9c, 0xa7, 0xd9, 0xa0, 0x3c, 0x2d, 0x63, 0x64, 0xb6, 0x86,
	0x93, 0x5e, 0x6f, 0x74, 0xda, 0xba, 0x98, 0x88, 0x9e, 0xa6, 0xdf, 0x63, 0x8c, 0xbd, 0xf3, 0xd1,
	0x29, 0xfb, 0x23, 0xfa, 0x6e, 0x9d, 0xd9, 0xf6, 0xd9, 0x80, 0xf0, 0x5e, 0xcb, 0xb2, 0x3d, 0xc3,
	0x33, 0x6d, 0xcb, 0xe5, 0xbd, 0xf8, 0xaf, 0x34, 0xa8, 0xe9, 0xc4, 0x1d, 0xd9, 0x96, 0x4b, 0x9e,
	0x12, 0xa3, 0x4f, 0x1c, 0x74, 0x1b, 0xa0, 0x37, 0x18, 0xbb, 0x1e, 0x71, 0x4e, 0xcc, 0x7e, 0x43,
	0x5b, 0xd7, 0x36, 0xe6, 0xf5, 0x92, 0xa0, 0xec, 0xf6, 0xd1, 0x1b, 0x50, 0x1a, 0x92, 0xe1, 0x29,
	0xef, 0xcd, 0xb0, 0xde, 0x22, 0x27, 0xec, 0xf6, 0x51, 0x13, 0x8a, 0x0e, 0x99, 0x98, 0xae, 0x69,
	0x5b, 0x8d, 0xec, 0xba, 0xb6, 0x91, 0xd5, 0xfd, 0x36, 0x1d, 0xe8, 0x18, 0x2f, 0xbd, 0x13, 0x8f,
	0x38, 0xc3, 0xc6, 0x3c, 0x1f, 0x48, 0x09, 0x5d, 0xe2, 0x0c, 0xf1, 0x4f, 0x72, 0x50, 0xd1, 0x0d,
	0xeb, 0x8c, 0xe8, 0xe4, 0x87, 0x63, 0xe2, 0x7a, 0xa8, 0x0e, 0xd9, 0x0b, 0x72, 0xc5, 0xe0, 0x2b,
	0x3a, 0xfd, 0xc9, 0xc7, 0x5b, 0x67, 0xe4, 0x84, 0x58, 0x1c, 0xb8, 0x42, 0xc7, 0x5b, 0x67, 0xa4,
	0x63, 0xf5, 0xd1, 0x32, 0xe4, 0x06, 0xe6, 0xd0, 0xf4, 0x04, 0x2a, 0x6f, 0x84, 0xd4, 0x99, 0x8f,
	0xa8, 0xb3, 0x0d, 0xe0, 0xda, 0x8e, 0x77, 0x62, 0x3b, 0x7d, 0xe2, 0x34, 0x72, 0xeb, 0xda, 0x46,
	0x6d, 0xeb, 0xad, 0x4d, 0x75, 0x19, 0x36, 0x55, 0x85, 0x36, 0x8f, 0x6d, 0xc7, 0x3b, 0xa4, 0xbc,
	0x7a, 0xc9, 0x95, 0x3f, 0xd1, 0x87, 0x50, 0x66, 0x42, 0x3c, 0xc3, 0x39, 0x23, 0x5e, 0x23, 0xcf,
	0xa4, 0xdc, 0xbb, 0x46, 0x4a, 0x97, 0x31, 0xeb, 0x0c, 0x9e, 0xff, 0x46, 0x18, 0x2a, 0x2e, 0x71,
	0x4c, 0x63, 0x60, 0x7e, 0x61, 0x9c, 0x0e, 0x48, 0xa3, 0xb0, 0xae, 0x6d, 0x14, 0xf5, 0x10, 0x8d,
	0xce, 0xff, 0x82, 0x5c, 0xb9, 0x27, 0xb6, 0x35, 0xb8, 0x6a, 0x14, 0x19, 0x43, 0x91, 0x12, 0x0e,
	0xad, 0xc1, 0x15, 0x5b, 0x34, 0x7b, 0x6c, 0x79, 0xbc, 0xb7, 0xc4, 0x7a, 0x4b, 0x8c, 0xc2, 0xba,
	0x37, 0xa0, 0x3e, 0x34, 0xad, 0x93, 0xa1, 0xdd, 0x3f, 0xf1, 0x0d, 0x02, 0xcc, 0x20, 0xb5, 0xa1,
	0x69, 0x3d, 0xb3, 0xfb, 0xba, 0x34, 0x0b, 0xe5, 0x34, 0x2e, 0xc3, 0x9c, 0x65, 0xc1, 0x69, 0x5c,
	0xaa, 0x9c, 0x9b, 0xb0, 0x44, 0x65, 0xf6, 0x1c, 0x62, 0x78, 0x24, 0x60, 0xae, 0x30, 0xe6, 0xc5,
	0xa1, 0x69, 0x6d, 0xb3, 0x9e, 0x10, 0xbf, 0x71, 0x19, 0xe3, 0xaf, 0x0a, 0x7e, 0xe3, 0x32, 0xcc,
	0x8f, 0x37, 0xa1, 0xe4, 0xdb, 0x1c, 0x15, 0x61, 0xfe, 0xe0, 0xf0, 0xa0, 0x53, 0x9f, 0x43, 0x00,
	0xf9, 0xf6, 0xf1, 0x76, 0xe7, 0x60, 0xa7, 0xae, 0xa1, 0x32, 0x14, 0x76, 0x3a, 0xbc, 0x91, 0xc1,
	0x8f, 0x01, 0x02, 0xeb, 0xa2, 0x02, 0x64, 0xf7, 0x3a, 0x9f, 0xd5, 0xe7, 0x28, 0xcf, 0x8b, 0x8e,
	0x7e, 0xbc, 0x7b, 0x78, 0x50, 0xd7, 0xe8, 0xe0, 0x6d, 0xbd, 0xd3, 0xee, 0x76, 0xea, 0x19, 0xca,
	0xf1, 0xec, 0x70, 0xa7, 0x9e, 0x45, 0x25, 0xc8, 0xbd, 0x68, 0xef, 0x3f, 0xef, 0xd4, 0xe7, 0xf1,
	0xcf, 0x35, 0xa8, 0x8a, 0xf5, 0xe2, 0x67, 0x02, 0x7d, 0x07, 0xf2, 0xe7, 0xec, 0x5c, 0xb0, 0xad,
	0x58, 0xde, 0xba, 0x15, 0x59, 0xdc, 0xd0, 0xd9, 0xd1, 0x05, 0x2f, 0xc2, 0x90, 0xbd, 0x98, 0xb8,
	0x8d, 0xcc, 0x7a, 0x76, 0xa3, 0xbc, 0x55, 0xdf, 0xe4, 0xe7, 0x75, 0x73, 0x8f, 0x5c, 0xbd, 0x30,
	0x06, 0x63, 0xa2, 0xd3, 0x4e, 0x84, 0x60, 0x7e, 0x68, 0x3b, 0x84, 0xed, 0xd8, 0xa2, 0xce, 0x7e,
	0xd3, 0x6d, 0xcc, 0x16, 0x4d, 0xec, 0x56, 0xde, 0xc0, 0xbf, 0xd4, 0x00, 0x8e, 0xc6, 0x5e, 0xfa,
	0xd1, 0x58, 0x86, 0xdc, 0x84, 0x0a, 0x16, 0xc7, 0x82, 0x37, 0xd8, 0x99, 0x20, 0x86, 0x4b, 0xfc,
	0x33, 0x41, 0x1b, 0xe8, 0x06, 0x14, 0x46, 0x0e, 0x99, 0x9c, 0x5c, 0x4c, 0x18, 0x48, 0x51, 0xcf,
	0xd3, 0xe6, 0xde, 0x04, 0xbd, 0x09, 0x15, 0xf3, 0xcc, 0xb2, 0x1d, 0x72, 0xc2, 0x65, 0xe5, 0x58,
	0x6f, 0x99, 0xd3, 0x98, 0xde, 0x0a, 0x0b, 0x17, 0x9c, 0x57, 0x59, 0xf6, 0x29, 0x09, 0x5b, 0x50,
	0x66, 0xaa, 0xce, 0x64, 0xbe, 0x77, 0x02, 0x1d, 0x33, 0x6c, 0x58, 0xdc, 0x84, 0x42, 0x6b, 0xfc,
	0x03, 0x40, 0x3b, 0x64, 0x40, 0x3c, 0x32, 0x8b, 0xf7, 0x50, 0x6c, 0x92, 0x55, 0x6d, 0x82, 0x7f,
	0xa6, 0xc1, 0x52, 0x48, 0xfc, 0x4c, 0xd3, 0x6a, 0x40, 0xa1, 0xcf, 0x84, 0x71, 0x0d, 0xb2, 0xba,
	0x6c, 0xa2, 0x87, 0x50, 0x14, 0x0a, 0xb8, 0x8d, 0x6c, 0xca, 0xa6, 0x29, 0x70, 0x9d, 0x5c, 0xfc,
	0xcb, 0x0c, 0x94, 0xc4, 0x44, 0x0f, 0x47, 0xa8, 0x0d, 0x55, 0x87, 0x37, 0x4e, 0xd8, 0x7c, 0x84,
	0x46, 0xcd, 0x74, 0x27, 0xf4, 0x74, 0x4e, 0xaf, 0x88, 0x21, 0x8c, 0x8c, 0x7e, 0x1f, 0xca, 0x52,
	0xc4, 0x68, 0xec, 0x09, 0x93, 0x37, 0xc2, 0x02, 0x82, 0xfd, 0xf7, 0x74, 0x4e, 0x07, 0xc1, 0x7e,
	0x34, 0xf6, 0x50, 0x17, 0x96, 0xe5, 0x60, 0x3e, 0x1b, 0xa1, 0x46, 0x96, 0x49, 0x59, 0x0f, 0x4b,
	0x89, 0x2f, 0xd5, 0xd3, 0x39, 0x1d, 0x89, 0xf1, 0x4a, 0xa7, 0xaa, 0x92, 0x77, 0xc9, 0x9d, 0x77,
	0x4c, 0xa5, 0xee, 0xa5, 0x15, 0x57, 0xa9, 0x7b, 0x69, 0x3d, 0x2e, 0x41, 0x41, 0xb4, 0xf0, 0xbf,
	0x64, 0x00, 0xe4, 0x6a, 0x1c, 0x8e, 0xd0, 0x0e, 0xd4, 0x1c, 0xd1, 0x0a, 0x59, 0xeb, 0x8d, 0x44,
	0x6b, 0x89, 0x45, 0x9c, 0xd3, 0xab, 0x72, 0x10, 0x57, 0xee, 0x7b, 0x50, 0xf1, 0xa5, 0x04, 0x06,
	0xbb, 0x99, 0x60, 0x30, 0x5f, 0x42, 0x59, 0x0e, 0xa0, 0x26, 0xfb, 0x04, 0x56, 0xfc, 0xf1, 0x09,
	0x36, 0x7b, 0x73, 0x8a, 0xcd, 0x7c, 0x81, 0x4b, 0x52, 0x82, 0x6a, 0x35, 0x55, 0xb1, 0xc0, 0x6c,
	0x37, 0x13, 0xcc, 0x16, 0x57, 0x8c, 0x1a, 0x0e, 0x68, 0xbc, 0xe4, 0x4d, 0xfc, 0xbf, 0x59, 0x28,
	0x6c, 0xdb, 0xc3, 0x91, 0xe1, 0xd0, 0xd5, 0xc8, 0x3b, 0xc4, 0x1d, 0x0f, 0x3c, 0x66, 0xae, 0xda,
	0xd6, 0xdd, 0xb0, 0x44, 0xc1, 0x26, 0xff, 0xea, 0x8c, 0x55, 0x17, 0x43, 0xe8, 0x60, 0x11, 0x1e,
	0x33, 0xaf, 0x31, 0x58, 0x04, 0x47, 0x31, 0x44, 0x1e, 0xe4, 0x6c, 0x70, 0x90, 0x9b, 0x50, 0x98,
	0x10, 0x27, 0x08, 0xe9, 0x4f, 0xe7, 0x74, 0x49, 0x40, 0xef, 0xc0, 0x42, 0x34, 0xbc, 0xe4, 0x04,
	0x4f, 0xad, 0x17, 0x8e, 0x46, 0x77, 0xa1, 0x12, 0x8a, 0x71, 0x79, 0xc1, 0x57, 0x1e, 0x2a, 0x21,
	0x6e, 0x55, 0xfa, 0x55, 0x1a, 0x8f, 0x2b, 0x4f, 0xe7, 0xa4, 0x67, 0x5d, 0x95, 0x9e, 0xb5, 0x28,
	0x46, 0x09, 0xdf, 0x1a, 0x72, 0x32, 0xdf, 0x0f, 0x3b, 0x19, 0xfc, 0x7d, 0xa8, 0x86, 0x0c, 0x44,
	0xe3, 0x4e, 0xe7, 0xe3, 0xe7, 0xed, 0x7d, 0x1e, 0xa4, 0x9e, 0xb0, 0xb8, 0xa4, 0xd7, 0x35, 0x1a,
	0xeb, 0xf6, 0x3b, 0xc7, 0xc7, 0xf5, 0x0c, 0xaa, 0x42, 0xe9, 0xe0, 0xb0, 0x7b, 0xc2, 0xb9, 0xb2,
	0xf8, 0x89, 0x2f, 0x41, 0x04, 0x39, 0x25, 0xb6, 0xcd, 0x29, 0xb1, 0x4d, 0x93, 0xb1, 0x2d, 0x13,
	0xc4, 0x36, 0x16, 0xe6, 0xf6, 0x3b, 0xed, 0xe3, 0x4e, 0x7d, 0xfe, 0x71, 0x0d, 0x2a, 0xdc, 0xbe,
	0x27, 0x63, 0x8b, 0x86, 0xda, 0x7f, 0xd2, 0x00, 0x82, 0xd3, 0x84, 0x5a, 0x50, 0xe8, 0x71, 0x9c,
	0x86, 0xc6, 0x9c, 0xd1, 0x4a, 0xe2, 0x92, 0xe9, 0x92, 0x0b, 0x7d, 0x0b, 0x0a, 0xee, 0xb8, 0xd7,
	0x23, 0xae, 0x0c, 0x79, 0x37, 0xa2, 0xfe, 0x50, 0x78, 0x2b, 0x5d, 0xf2, 0xd1, 0x21, 0x2f, 0x0d,
	0x73, 0x30, 0x66, 0x01, 0x70, 0xfa, 0x10, 0xc1, 0x87, 0xff, 0x5e, 0x83, 0xb2, 0xb2, 0x79, 0x7f,
	0x47, 0x27, 0x7c, 0x0b, 0x4a, 0x4c, 0x07, 0xd2, 0x17, 0x6e, 0xb8, 0xa8, 0x07, 0x04, 0xf4, 0x7b,
	0x50, 0x92, 0x27, 0x40, 0x7a, 0xe2, 0x46, 0xb2, 0xd8, 0xc3, 0x91, 0x1e, 0xb0, 0xe2, 0x3d, 0x58,
	0x64, 0x56, 0xe9, 0xd1, 0xe4, 0x5a, 0xda, 0x51, 0x4d, 0x3f, 0xb5, 0x48, 0xfa, 0xd9, 0x84, 0xe2,
	0xe8, 0xfc, 0xca, 0x35, 0x7b, 0xc6, 0x40, 0x68, 0xe1, 0xb7, 0xf1, 0x47, 0x80, 0x54, 0x61, 0xb3,
	0x4c, 0x17, 0x57, 0xa1, 0xfc, 0xd4, 0x70, 0xcf, 0x85, 0x4a, 0xf8, 0x21, 0x54, 0x69, 0x73, 0xef,
	0xc5, 0x6b, 0xe8, 0xc8, 0x2e, 0x07, 0x92, 0x7b, 0x26, 0x9b, 0x23, 0x98, 0x3f, 0x37, 0xdc, 0x73,
	0x36, 0xd1, 0xaa, 0xce, 0x7e, 0xa3, 0x77, 0xa0, 0xde, 0xe3, 0x93, 0x3c, 0x89, 0x5c, 0x19, 0x16,
	0x04, 0xdd, 0xcf, 0x04, 0x3f, 0x85, 0x0a, 0x9f, 0xc3, 0x57, 0xad, 0x04, 0x5e, 0x84, 0x85, 0x63,
	0xcb, 0x18, 0xb9, 0xe7, 0xb6, 0x8c, 0x6e, 0x74, 0xd2, 0xf5, 0x80, 0x36, 0x13, 0xe2, 0xdb, 0xb0,
	0xe0, 0x90, 0xa1, 0x61, 0x5a, 0xa6, 0x75, 0x76, 0x72, 0x7a, 0xe5, 0x11, 0x57, 0x5c, 0x98, 0x6a,
	0x3e, 0xf9, 0x31, 0xa5, 0x52, 0xd5, 0x4e, 0x07, 0xf6, 0xa9, 0x70, 0x73, 0xec, 0x37, 0xfe, 0x69,
	0x06, 0x2a, 0x9f, 0x18, 0x5e, 0x4f, 0x2e, 0x1d, 0xda, 0x85, 0x9a, 0xef, 0xdc, 0x18, 0x45, 0xe8,
	0x12, 0x09, 0xb1, 0x6c, 0x8c, 0x4c, 0xa5, 0x65, 0x74, 0xac, 0xf6, 0x54, 0x02, 0x13, 0x65, 0x58,
	0x3d, 0x32, 0xf0, 0x45, 0x65, 0xd2, 0x45, 0x31, 0x46, 0x55, 0x94, 0x4a, 0x40, 0x87, 0x50, 0x1f,
	0x39, 0xf6, 0x99, 0x43, 0x5c, 0xd7, 0x17, 0xc6, 0xc3, 0x18, 0x4e, 0x10, 0x76, 0x24, 0x58, 0x03,
	0x71, 0x0b, 0xa3, 0x30, 0xe9, 0xf1, 0x42, 0x90, 0xcf, 0x70, 0xe7, 0xf4, 0x9f, 0x19, 0x40, 0xf1,
	0x49, 0xfd, 0xb6, 0x29, 0xde, 0x3d, 0xa8, 0xb9, 0x9e, 0xe1, 0xc4, 0x36, 0x5b, 0x95, 0x51, 0x7d,
	0x8f, 0xff, 0x36, 0xf8, 0x0a, 0x9d, 0x58, 0xb6, 0x67, 0xbe, 0xbc, 0x12, 0x59, 0x72, 0x4d, 0x92,
	0x0f, 0x18, 0x15, 0x75, 0xa0, 0xf0, 0xd2, 0x1c, 0x78, 0xc4, 0x71, 0x1b, 0xb9, 0xf5, 0xec, 0x46,
	0x6d, 0xeb, 0xe1, 0x75, 0xcb, 0xb0, 0xf9, 0x21, 0xe3, 0xef, 0x5e, 0x8d, 0x88, 0x2e, 0xc7, 0xaa,
	0x99, 0x67, 0x3e, 0x94, 0x8d, 0xdf, 0x84, 0xe2, 0x2b, 0x2a, 0x82, 0xde, 0xb2, 0x0b, 0x3c, 0x59,
	0x64, 0x6d, 0x7e, 0xc9, 0x7e, 0xe9, 0x18, 0x67, 0x43, 0x62, 0x79, 0xf2, 0x1e, 0x28, 0xdb, 0xf8,
	0x1e, 0x40, 0x00, 0x43, 0x5d, 0xfe, 0xc1, 0xe1, 0xd1, 0xf3, 0x6e, 0x7d, 0x0e, 0x55, 0xa0, 0x78,
	0x70, 0xb8, 0xd3, 0xd9, 0xef, 0xd0, 0xf8, 0x80, 0x5b, 0xd2, 0xa4, 0xa1, 0xb5, 0x54, 0x31, 0xb5,
	0x10, 0x26, 0x5e, 0x85, 0xe5, 0xa4, 0x05, 0xa4, 0xb9, 0x68, 0x55, 0xec, 0xd2, 0x99, 0x8e, 0x8a,
	0x0a, 0x9d, 0x09, 0x4f, 0xb7, 0x01, 0x05, 0xbe, 0x7b, 0xfb, 0x22, 0x39, 0x97, 0x4d, 0x6a, 0x08,
	0xbe, 0x19, 0x49, 0x5f, 0xac, 0x92, 0xdf, 0x4e, 0x74, 0x2f, 0xb9, 0x44, 0xf7, 0x82, 0xee, 0x42,
	0xd5, 0x3f, 0x0d, 0x86, 0x2b, 0x72, 0x81, 0x92, 0x5e, 0x91, 0x1b, 0x9d, 0xd2, 0x42, 0x46, 0x2f,
	0x84, 0x8d, 0x8e, 0xee, 0x41, 0x9e, 0x4c, 0x88, 0xe5, 0xb9, 0x8d, 0x32, 0x8b, 0x18, 0x55, 0x99,
	0xbb, 0x77, 0x28, 0x55, 0x17, 0x9d, 0xf8, 0xbb, 0xb0, 0xc8, 0xee, 0x48, 0x4f, 0x1c, 0xc3, 0x52,
	0x2f, 0x73, 0xdd, 0xee, 0xbe, 0x30, 0x37, 0xfd, 0x89, 0x6a, 0x90, 0xd9, 0xdd, 0x11, 0x46, 0xc8,
	0xec, 0xee, 0xe0, 0x1f, 0x6b, 0x80, 0xd4, 0x71, 0x33, 0xd9, 0x39, 0x22, 0x5c, 0xc2, 0x67, 0x03,
	0xf8, 0x65, 0xc8, 0x11, 0xc7, 0xb1, 0x1d, 0x66, 0xd1, 0x92, 0xce, 0x1b, 0xf8, 0x2d, 0xa1, 0x83,
	0x4e, 0x26, 0xf6, 0x85, 0x7f, 0x06, 0xb9, 0x34, 0xcd, 0x57, 0x75, 0x0f, 0x96, 0x42, 0x5c, 0x33,
	0x45, 0xae, 0x0f, 0x61, 0x81, 0x09, 0xdb, 0x3e, 0x27, 0xbd, 0x8b, 0x91, 0x6d, 0x5a, 0x31, 0x3c,
	0xba, 0x72, 0x81, 0x83, 0xa5, 0xf3, 0xe0, 0x13, 0xab, 0xf8, 0xc4, 0x6e, 0x77, 0x1f, 0x7f, 0x06,
	0xab, 0x11, 0x39, 0x52, 0xfd, 0x3f, 0x82, 0x72, 0xcf, 0x27, 0xba, 0x22, 0xd7, 0xb9, 0x1d, 0x56,
	0x2e, 0x3a, 0x54, 0x1d, 0x81, 0x0f, 0xe1, 0x46, 0x4c, 0xf4, 0x4c, 0x73, 0x7e, 0x1b, 0x56, 0x98,
	0xc0, 0x3d, 0x42, 0x46, 0xed, 0x81, 0x39, 0x49, 0xb5, 0xf4, 0x48, 0x4c, 0x4a, 0x61, 0xfc, 0x7a,
	0xf7, 0x05, 0xfe, 0x03, 0x81, 0xd8, 0x35, 0x87, 0xa4, 0x6b, 0xef, 0xa7, 0xeb, 0x46, 0xa3, 0xd9,
	0x05, 0xb9, 0x72, 0x45, 0x5a, 0xc3, 0x7e, 0xe3, 0x7f, 0xd6, 0x84, 0xa9, 0xd4, 0xe1, 0x5f, 0xf3,
	0x4e, 0x5e, 0x03, 0x38, 0xa3, 0x47, 0x86, 0xf4, 0x69, 0x07, 0xaf, 0xa8, 0x28, 0x14, 0x5f, 0x4f,
	0xea, 0xbf, 0x2b, 0x42, 0xcf, 0x65, 0xb1, 0xcf, 0xd9, 0x3f, 0xbe, 0x97, 0xbb, 0x0d, 0x65, 0x46,
	0x38, 0xf6, 0x0c, 0x6f, 0xec, 0xc6, 0x16, 0xe3, 0x2f, 0xc4, 0xb6, 0x97, 0x83, 0x66, 0x9a, 0xd7,
	0xb7, 0x20, 0xcf, 0x2e, 0x13, 0x32, 0x95, 0xbe, 0x99, 0xb0, 0x1f, 0xb9, 0x1e, 0xba, 0x60, 0xc4,
	0x3f, 0xd5, 0x20, 0xff, 0x8c, 0x95, 0x60, 0x15, 0xd5, 0xe6, 0xe5, 0x5a, 0x58, 0xc6, 0x90, 0x17,
	0x86, 0x4a, 0x3a, 0xfb, 0xcd, 0x52, 0x4f, 0x42, 0x9c, 0xe7, 0xfa, 0x3e, 0x4f, 0x71, 0x4b, 0xba,
	0xdf, 0xa6, 0x36, 0xeb, 0x0d, 0x4c, 0x62, 0x79, 0xac, 0x77, 0x9e, 0xf5, 0x2a, 0x14, 0x9a, 0x3d,
	0x9b, 0xee, 0x3e, 0x31, 0x1c, 0x4b, 0x14, 0x4d, 0x8b, 0x7a, 0x40, 0xc0, 0xfb, 0x50, 0xe7, 0x7a,
	0xb4, 0xfb, 0x7d, 0x25, 0xc1, 0xf4, 0xd1, 0xb4, 0x08, 0x5a, 0x48, 0x5a, 0x26, 0x2a, 0xed, 0x17,
	0x1a, 0x2c, 0x2a, 0xe2, 0x66, 0xb2, 0xea, 0xbb, 0x90, 0xe7, 0x45, 0x6a, 0x91, 0xe9, 0x2c, 0x87,
	0x47, 0x71, 0x18, 0x5d, 0xf0, 0xa0, 0x4d, 0x28, 0xf0, 0x5f, 0xf2, 0x0e, 0x90, 0xcc, 0x2e, 0x99,
	0xf0, 0x3d, 0x58, 0x12, 0x24, 0x32, 0xb4, 0x93, 0x0e, 0x06, 0x5b, 0x0c, 0xfc, 0x67, 0xb0, 0x1c,
	0x66, 0x9b, 0x69, 0x4a, 0x8a, 0x92, 0x99, 0xd7, 0x51, 0xb2, 0x2d, 0x95, 0x7c, 0x3e, 0xea, 0x2b,
	0x79, 0x54, 0x74, 0xc7, 0xa8, 0xeb, 0x95, 0x09, 0xaf, 0x57, 0x30, 0x01, 0x29, 0xe2, 0x1b, 0x9d,
	0xc0, 0x07, 0x72, 0x3b, 0xec, 0x9b, 0xae, 0xef, 0xc3, 0x31, 0x54, 0x06, 0xa6, 0x45, 0x0c, 0x47,
	0x54, 0xce, 0x35, 0x5e, 0x39, 0x57, 0x69, 0xf8, 0x0b, 0x40, 0xea, 0xc0, 0x6f, 0x54, 0xe9, 0xfb,
	0xd2, 0x64, 0x47, 0x8e, 0x3d, 0xb4, 0x53, 0xcd, 0x8e, 0xff, 0x1c, 0x56, 0x22, 0x7c, 0xdf, 0xa8,
	0x9a, 0x4b, 0xb0, 0xb8, 0x43, 0x64, 0x42, 0x23, 0xdd, 0xde, 0x47, 0x80, 0x54, 0xe2, 0x4c, 0x91,
	0xad, 0x05, 0x8b, 0xcf, 0xec, 0x09, 0x75, 0x91, 0x94, 0x1a, 0xf8, 0x06, 0x5e, 0x87, 0xf0, 0x4d,
	0xe1, 0xb7, 0x29, 0xb8, 0x3a, 0x60, 0x26, 0xf0, 0x7f, 0xd7, 0xa0, 0xd2, 0x1e, 0x18, 0xce, 0x50,
	0x02, 0x7f, 0x0f, 0xf2, 0xfc, 0x76, 0x2d, 0x0a, 0x5a, 0xf7, 0xc3, 0x62, 0x54, 0x5e, 0xde, 0x68,
	0xf3, 0xbb, 0xb8, 0x18, 0x45, 0x15, 0x17, 0x6f, 0x5e, 0x3b, 0x91, 0x37, 0xb0, 0x1d, 0xf4, 0x1e,
	0xe4, 0x0c, 0x3a, 0x84, 0x85, 0xa2, 0x5a, 0xb4, 0xae, 0xc1, 0xa4, 0xb1, 0x3b, 0x00, 0xe7, 0xc2,
	0xdf, 0x81, 0xb2, 0x82, 0x80, 0x0a, 0x90, 0x7d, 0xd2, 0x11, 0x09, 0x7b, 0x7b, 0xbb, 0xbb, 0xfb,
	0x82, 0x17, 0x74, 0x6a, 0x00, 0x3b, 0x1d, 0xbf, 0x9d, 0xc1, 0x9f, 0x8a, 0x51, 0xc2, 0xed, 0xab,
	0xfa, 0x68, 0x69, 0xfa, 0x64, 0x5e, 0x4b, 0x9f, 0x4b, 0xa8, 0x8a, 0xe9, 0xcf, 0x1a, 0xc6, 0x98,
	0xbc, 0x94, 0x30, 0xa6, 0x28, 0xaf, 0x0b, 0x46, 0xfc, 0x2b, 0x0d, 0xea, 0x3b, 0xf6, 0x2b, 0xeb,
	0xcc, 0x31, 0xfa, 0xfe, 0x39, 0xf9, 0x30, 0xb2, 0x52, 0x9b, 0x91, 0xe2, 0x68, 0x84, 0x3f, 0x20,
	0x44, 0x56, 0xac, 0x11, 0x94, 0x0d, 0x79, 0x2c, 0x94, 0x4d, 0xfc, 0x01, 0x2c, 0x44, 0x06, 0x51,
	0xdb, 0xbf, 0x68, 0xef, 0xef, 0xee, 0x50, 0x5b, 0xb3, 0xc2, 0x5a, 0xe7, 0xa0, 0xfd, 0x78, 0xbf,
	0x23, 0x1e, 0x90, 0xda, 0x07, 0xdb, 0x9d, 0xfd, 0x7a, 0x06, 0xf7, 0x60, 0x51, 0x81, 0x9f, 0xf5,
	0x65, 0x20, 0x45, 0xbb, 0x05, 0xa8, 0x8a, 0x68, 0x2f, 0x0e, 0xe5, 0xbf, 0x65, 0xa0, 0x26, 0x29,
	0x5f, 0x0f, 0x26, 0x5a, 0x85, 0x7c, 0xff, 0xf4, 0xd8, 0xfc, 0x42, 0xbe, 0x1c, 0x89, 0x16, 0xa5,
	0x0f, 0x38, 0x0e, 0x7f, 0xbe, 0x15, 0x2d, 0x1a, 0xc6, 0x1d, 0xe3, 0xa5, 0xb7, 0x6b, 0xf5, 0xc9,
	0x25, 0x4b, 0x0a, 0xe6, 0xf5, 0x80, 0xc0, 0x2a, 0x4c, 0xe2, 0x99, 0x97, 0xdd, 0xac, 0x94, 0x67,
	0x5f, 0xf4, 0x00, 0xea, 0xf4, 0x77, 0x7b, 0x34, 0x1a, 0x98, 0xa4, 0xcf, 0x05, 0x14, 0x18, 0x4f,
	0x8c, 0x4e, 0xd1, 0xd9, 0x5d, 0xc4, 0x6d, 0x14, 0x59, 0x58, 0x12, 0x2d, 0xb4, 0x0e, 0x65, 0xae,
	0xdf, 0xae, 0xf5, 0xdc, 0x25, 0xec, 0xed, 0x33, 0xab, 0xab, 0xa4, 0x70, 0x9a, 0x01, 0xd1, 0x34,
	0x63, 0x09, 0x16, 0xdb, 0x63, 0xef, 0xbc, 0x63, 0xd1, 0x58, 0x21, 0xad, 0xbc, 0x0c, 0x88, 0x12,
	0x77, 0x4c, 0x57, 0xa5, 0x0a, 0xd6, 0xf0, 0x82, 0x74, 0x60, 0x89, 0x12, 0x89, 0xe5, 0x99, 0x3d,
	0x25, 0xae, 0xca, 0xcc, 0x4b, 0x8b, 0x64, 0x5e, 0x86, 0xeb, 0xbe, 0xb2, 0x9d, 0xbe, 0xb0, 0xb9,
	0xdf, 0xc6, 0xff, 0xa8, 0x71, 0xc8, 0xe7, 0x6e, 0x28, 0x7d, 0xfa, 0x2d, 0xc5, 0xa0, 0xf7, 0xa1,
	0x60, 0x8f, 0xd8, 0x0b, 0xbf, 0x28, 0xc3, 0xac, 0x6e, 0xf2, 0x6f, 0x02, 0x36, 0x85, 0xe0, 0x43,
	0xde, 0xab, 0x4b, 0x36, 0x74, 0x1f, 0x6a, 0xe7, 0x86, 0x7b, 0x4e, 0xfa, 0x47, 0x52, 0x26, 0xbf,
	0xf9, 0x45, 0xa8, 0x78, 0x23, 0xd0, 0xef, 0x09, 0xf1, 0xa6, 0xe8, 0x87, 0x1f, 0xc2, 0x8a, 0xe4,
	0x14, 0xaf, 0x13, 0x53, 0x98, 0x5f, 0xc1, 0x6d, 0xc9, 0xbc, 0x7d, 0x6e, 0x58, 0x67, 0x44, 0x02,
	0xfe, 0xae, 0x16, 0x88, 0xcf, 0x27, 0x9b, 0x38, 0x9f, 0xc7, 0xd0, 0xf0, 0xe7, 0xc3, 0x6e, 0xd6,
	0xf6, 0x40, 0x55, 0x74, 0xec, 0x8a, 0xf3, 0x54, 0xd2, 0xd9, 0x6f, 0x4a, 0x73, 0xec, 0x81, 0x9f,
	0x4a, 0xd3, 0xdf, 0x78, 0x1b, 0x6e, 0x4a, 0x19, 0xe2, 0xce, 0x1b, 0x16, 0x12, 0x53, 0x3c, 0x49,
	0x88, 0x30, 0x2c, 0x1d, 0x3a, 0x7d, 0xe1, 0x55, 0xce, 0xf0, 0x12, 0x30, 0x99, 0x9a, 0x22, 0x73,
	0x85, 0x6f, 0x4a, 0xaa, 0x98, 0x92, 0x2d, 0x49, 0x32, 0x15, 0xa0, 0x92, 0xc5, 0x82, 0x51, 0x72,
	0x6c, 0xc1, 0x62, 0xa2, 0x7f, 0x00, 0x6b, 0xbe, 0x12, 0xd4, 0x6e, 0x47, 0xc4, 0x19, 0x9a, 0xae,
	0xab, 0xd4, 0xbd, 0x93, 0x26, 0x7e, 0x1f, 0xe6, 0x47, 0x44, 0x04, 0xa1, 0xf2, 0x16, 0x92, 0x9b,
	0x52, 0x19, 0xcc, 0xfa, 0x71, 0x1f, 0xee, 0x48, 0xe9, 0xdc, 0xa2, 0x89, 0xe2, 0xa3, 0x4a, 0xc9,
	0x6a, 0x60, 0x26, 0xa5, 0x1a, 0x98, 0x8d, 0xbc, 0xc5, 0x7c, 0xc4, 0x0d, 0x29, 0xcf, 0xfc, 0x4c,
	0xc9, 0xc5, 0x1e, 0xb7, 0xa9, 0xef, 0x2a, 0x66, 0x12, 0xf6, 0xd7, 0xc2, 0x0b, 0x7c, 0x55, 0x1e,
	0x9e, 0xb0, 0x19, 0xca, 0x87, 0x0e, 0xd9, 0xa4, 0x59, 0x33, 0x5d, 0x00, 0x5d, 0xad, 0x85, 0xce,
	0xeb, 0x21, 0x1a, 0x3e, 0x85, 0xe5, 0xb0, 0x5f, 0x9b, 0x49, 0x97, 0x65, 0xc8, 0x79, 0xf6, 0x05,
	0x91, 0xb1, 0x86, 0x37, 0xa4, 0xed, 0x7c, 0x9f, 0x37, 0x93, 0xed, 0x8c, 0x40, 0x18, 0x3b, 0x1d,
	0xb3, 0xea, 0x4b, 0x37, 0x96, 0xbc, 0x03, 0xf1, 0x06, 0x3e, 0x80, 0xd5, 0xa8, 0x67, 0x9b, 0x49,
	0xe5, 0x17, 0xfc, 0x2c, 0x25, 0x39, 0xbf, 0x99, 0xe4, 0x7e, 0x1c, 0xf8, 0x25, 0xc5, 0xb7, 0xcd,
	0x24, 0x52, 0x87, 0x66, 0x92, 0xab, 0xfb, 0x2a, 0x8e, 0x8e, 0xef, 0xf9, 0x66, 0x12, 0xe6, 0x06,
	0xc2, 0x66, 0x5f, 0xfe, 0xc0, 0x5d, 0x65, 0xa7, 0xba, 0x2b, 0x71, 0x48, 0x02, 0x87, 0xfa, 0x35,
	0x6c, 0x3a, 0x81, 0x11, 0xf8, 0xf2, 0x59, 0x31, 0x68, 0x38, 0xf3, 0x31, 0x58, 0x43, 0x6e, 0x6c,
	0x35, 0x02, 0xcc, 0xb4, 0x18, 0x9f, 0x04, 0x6e, 0x3c, 0x16, 0x24, 0x66, 0x12, 0xfc, 0x29, 0xac,
	0xa7, 0xc7, 0x87, 0x59, 0x24, 0x3f, 0x68, 0x41, 0xc9, 0xbf, 0x0c, 0x29, 0xdf, 0x9b, 0x95, 0xa1,
	0x70, 0x70, 0x78, 0x7c, 0xd4, 0xde, 0xee, 0xf0, 0x0f, 0xce, 0xb6, 0x0f, 0x75, 0xfd, 0xf9, 0x51,
	0xb7, 0x9e, 0xd9, 0xfa, 0x75, 0x16, 0x32, 0x7b, 0x2f, 0xd0, 0x67, 0x90, 0xe3, 0x5f, 0x5f, 0x4c,
	0xf9, 0xe4, 0xa6, 0x39, 0xed, 0x03, 0x13, 0x7c, 0xe3, 0xc7, 0xff, 0xf5, 0xeb, 0x9f, 0x67, 0x16,
	0x71, 0xa5, 0x35, 0xf9, 0x76, 0xeb, 0x62, 0xd2, 0x62, 0x61, 0xea, 0x91, 0xf6, 0x00, 0x7d, 0x0c,
	0xd9, 0xa3, 0xb1, 0x87, 0x52, 0x3f, 0xc5, 0x69, 0xa6, 0x7f, 0x73, 0x82, 0x57, 0x98, 0xd0, 0x05,
	0x0c, 0x42, 0xe8, 0x68, 0xec, 0x51, 0x91, 0x3f, 0x84, 0xb2, 0xfa, 0xc5, 0xc8, 0xb5, 0xdf, 0xe7,
	0x34, 0xaf, 0xff, 0x1a, 0x05, 0xdf, 0x66, 0x50, 0x37, 0x30, 0x12, 0x50, 0xfc, 0x9b, 0x16, 0x75,
	0x16, 0xdd, 0x4b, 0x0b, 0xa5, 0x7e, 0xbd, 0xd3, 0x4c, 0xff, 0x40, 0x25, 0x36, 0x0b, 0xef, 0xd2,
	0xa2, 0x22, 0xff, 0x44, 0x7c, 0x9b, 0xd2, 0xf3, 0xd0, 0x9d, 0x84, 0x6f, 0x13, 0xd4, 0x57, 0xf8,
	0xe6, 0x7a, 0x3a, 0x83, 0x00, 0xb9, 0xc5, 0x40, 0x56, 0xf1, 0xa2, 0x00, 0xe9, 0xf9, 0x2c, 0x8f,
	0xb4, 0x07, 0x5b, 0x3d, 0xc8, 0xb1, 0x17, 0x2e, 0xf4, 0xb9, 0xfc, 0xd1, 0x4c, 0x78, 0xea, 0x4b,
	0x59, 0xe8, 0xd0, 0xdb, 0x18, 0x5e, 0x66, 0x40, 0x35, 0x5c, 0xa2, 0x40, 0xec, 0x7d, 0xeb, 0x91,
	0xf6, 0x60, 0x43, 0x7b, 0x5f, 0xdb, 0xfa, 0x55, 0x0e, 0x72, 0xac, 0xb4, 0x8b, 0x2e, 0x00, 0x82,
	0xd7, 0x9e, 0xe8, 0xec, 0x62, 0xef, 0x47, 0xd1, 0xd9, 0xc5, 0x1f, 0x8a, 0x70, 0x93, 0x81, 0x2e,
	0xe3, 0x05, 0x0a, 0xca, 0x2a, 0xc6, 0x2d, 0x56, 0x04, 0xa7, 0x76, 0xfc, 0x1b, 0x4d, 0x54, 0xb6,
	0xf9, 0x59, 0x42, 0x49, 0xd2, 0x42, 0x4f, 0x3e, 0xd1, 0xed, 0x90, 0xf0, 0xdc, 0x83, 0xbf, 0xcb,
	0x00, 0x5b, 0xb8, 0x1e, 0x00, 0x3a, 0x8c, 0xe3, 0x91, 0xf6, 0xe0, 0xf3, 0x06, 0x5e, 0x12, 0x56,
	0x8e, 0xf4, 0xa0, 0x1f, 0x41, 0x2d, 0xfc, 0xa4, 0x81, 0xee, 0x26, 0x60, 0x45, 0x5f, 0x46, 0x9a,
	0x6f, 0x4d, 0x67, 0x12, 0x3a, 0xad, 0x31, 0x9d, 0x04, 0x38, 0x47, 0xbe, 0x20, 0x64, 0x64, 0x50,
	0x26, 0xb1, 0x06, 0xe8, 0x1f, 0x34, 0xf1, 0xe2, 0x14, 0xbc, 0x51, 0xa0, 0x24, 0xe9, 0xb1, 0x17,
	0x90, 0xe6, 0xbd, 0x6b, 0xb8, 0x84, 0x12, 0x7f, 0xc8, 0x94, 0xf8, 0x00, 0x2f, 0x07, 0x4a, 0x78,
	0xe6, 0x90, 0x78, 0xb6, 0xd0, 0xe2, 0xf3, 0x5b, 0xf8, 0x46, 0xc8, 0x38, 0xa1, 0xde, 0x60, 0xb1,
	0xf8, 0x3b, 0x43, 0xe2, 0x62, 0x85, 0xde, 0x2d, 0x12, 0x17, 0x2b, 0xfc, 0x48, 0x91, 0xb4, 0x58,
	0xfc, 0x55, 0x21, 0x69, 0xb1, 0xfc, 0x9e, 0xad, 0xff, 0x9b, 0x87, 0xc2, 0x36, 0xff, 0x26, 0x1c,
	0xd9, 0x50, 0xf2, 0xcb, 0xf4, 0x68, 0x2d, 0xa9, 0xce, 0x18, 0x5c, 0x6b, 0x9a, 0x77, 0x52, 0xfb,
	0x85, 0x42, 0x6f, 0x32, 0x85, 0xde, 0xc0, 0xab, 0x14, 0x59, 0x7c, 0x76, 0xde, 0xe2, 0xc5, 0xac,
	0x96, 0xd1, 0xef, 0x53, 0x43, 0xfc, 0x29, 0x54, 0xd4, 0x3a, 0x3a, 0x7a, 0x33, 0xb1, 0xb6, 0xa9,
	0x96, 0xe2, 0x9b, 0x78, 0x1a, 0x8b, 0x40, 0x7e, 0x8b, 0x21, 0xaf, 0xe1, 0x9b, 0x09, 0xc8, 0x0e,
	0x63, 0x0d, 0x81, 0xf3, 0x1a, 0x78, 0x32, 0x78, 0xa8, 0xc4, 0x9e, 0x0c, 0x1e, 0x2e, 0xa1, 0x4f,
	0x05, 0x1f, 0x33, 0x56, 0x0a, 0xee, 0x02, 0x04, 0x95, 0x6c, 0x94, 0x68, 0x4b, 0xe5, 0x5e, 0x17,
	0x75, 0x0e, 0xf1, 0x22, 0x38, 0xc6, 0x0c, 0x56, 0xec, 0xbb, 0x08, 0xec, 0xc0, 0x74, 0x3d, 0x7e,
	0x30, 0xab, 0xa1, 0xd2, 0x34, 0x4a, 0x9c, 0x4f, 0xb8, 0xbe, 0xdd, 0xbc, 0x3b, 0x95, 0x47, 0xa0,
	0xdf, 0x63, 0xe8, 0x77, 0x70, 0x33, 0x01, 0x7d, 0xc4, 0x79, 0xe9, 0x66, 0xfb, 0xff, 0x3c, 0x94,
	0x9f, 0x19, 0xa6, 0xe5, 0x11, 0xcb, 0xb0, 0x7a, 0x04, 0x9d, 0x42, 0x8e, 0x45, 0xea, 0xa8, 0x23,
	0x56, 0xcb, 0xb6, 0x51, 0x47, 0x1c, 0xaa, 0x69, 0xe2, 0x75, 0x06, 0xdc, 0xc4, 0x2b, 0x14, 0x78,
	0x18, 0x88, 0x6e, 0xb1, 0x52, 0x24, 0x9d, 0xf4, 0x4b, 0xc8, 0x8b, 0xd7, 0xbe, 0x88, 0xa0, 0x50,
	0xf1, 0xa7, 0x79, 0x2b, 0xb9, 0x33, 0x69, 0x2f, 0xab, 0x30, 0x2e, 0xe3, 0xa3, 0x38, 0x13, 0x80,
	0xa0, 0xc6, 0x1e, 0x5d, 0xd1, 0x58, 0x49, 0xbe, 0xb9, 0x9e, 0xce, 0x90, 0x64, 0x53, 0x15, 0xb3,
	0xef, 0xf3, 0x52, 0xdc, 0x3f, 0x86, 0xf9, 0xa7, 0x86, 0x7b, 0x8e, 0x22, 0xb1, 0x57, 0xf9, 0x56,
	0xac, 0xd9, 0x4c, 0xea, 0x12, 0x28, 0x77, 0x18, 0xca, 0x4d, 0xee, 0xca, 0x54, 0x94, 0x73, 0xc3,
	0xa5, 0x41, 0x0d, 0xf5, 0x21, 0xcf, 0x3f, 0x1d, 0x8b, 0xda, 0x2f, 0xf4, 0xf9, 0x59, 0xd4, 0x7e,
	0xe1, 0xaf, 0xcd, 0xae, 0x47, 0x19, 0x41, 0x51, 0x7e, 0xab, 0x85, 0x22, 0x0f, 0xf7, 0x91, 0xef,
	0xba, 0x9a, 0x6b, 0x69, 0xdd, 0x02, 0xeb, 0x2e, 0xc3, 0xba, 0x8d, 0x1b, 0xb1, 0xb5, 0x12, 0x9c,
	0x8f, 0xb4, 0x07, 0xef, 0x6b, 0xe8, 0x47, 0x00, 0xc1, 0xb3, 0x44, 0xec, 0x04, 0x46, 0x5f, 0x38,
	0x62, 0x27, 0x30, 0xf6, 0xa2, 0x81, 0x37, 0x19, 0xee, 0x06, 0xbe, 0x1b, 0xc5, 0xf5, 0x1c, 0xc3,
	0x72, 0x5f, 0x12, 0xe7, 0x3d, 0x5e, 0x65, 0x75, 0xcf, 0xcd, 0x11, 0x9d, 0xb2, 0x03, 0x25, 0xbf,
	0xea, 0x1c, 0xf5, 0xb6, 0xd1, 0x6a, 0x78, 0xd4, 0xdb, 0xc6, 0xca, 0xd5, 0x61, 0xb7, 0x13, 0xda,
	0x2d, 0x92, 0x95, 0x1e, 0xc0, 0x5f, 0xd4, 0x61, 0x9e, 0x66, 0xdd, 0x34, 0x39, 0x09, 0xea, 0x26,
	0xd1, 0xd9, 0xc7, 0xaa, 0xa8, 0xd1, 0xd9, 0xc7, 0x4b, 0x2e, 0xe1, 0xe4, 0x84, 0x5e, 0xb2, 0x5a,
	0xbc, 0x44, 0x41, 0x67, 0x6a, 0x43, 0x59, 0x29, 0xac, 0xa0, 0x04, 0x61, 0xe1, 0xf2, 0x6c, 0x34,
	0xdc, 0x25, 0x54, 0x65, 0xf0, 0x1b, 0x0c, 0x6f, 0x85, 0x87, 0x3b, 0x86, 0xd7, 0xe7, 0x1c, 0x14,
	0x50, 0xcc, 0x4e, 0x9c, 0xfb, 0x84, 0xd9, 0x85, 0xcf, 0xfe, 0x7a, 0x3a, 0x43, 0xea, 0xec, 0x82,
	0x83, 0xff, 0x0a, 0x2a, 0x6a, 0x79, 0x05, 0x25, 0x28, 0x1f, 0x29, 0x29, 0x47, 0xe3, 0x48, 0x52,
	0x75, 0x26, 0xec, 0xd9, 0x18, 0xa4, 0xa1, 0xb0, 0x51, 0xe0, 0x01, 0x14, 0x44, 0xbd, 0x25, 0xc9,
	0xa4, 0xe1, 0xf2, 0x73, 0x92, 0x49, 0x23, 0xc5, 0x9a, 0x70, 0xf6, 0xcc, 0x10, 0xe9, 0x95, 0x52,
	0xc6, 0x6a, 0x81, 0xf6, 0x84, 0x78, 0x69, 0x68, 0x41, 0x25, 0x33, 0x0d, 0x4d, 0xb9, 0xce, 0xa7,
	0xa1, 0x9d, 0x11, 0x4f, 0xf8, 0x03, 0x79, 0x4d, 0x46, 0x29, 0xc2, 0xd4, 0xf8, 0x88, 0xa7, 0xb1,
	0x24, 0x5d, 0x6e, 0x02, 0x40, 0x19, 0x1c, 0x2f, 0x01, 0x82, 0x6a, 0x50, 0x34, 0x63, 0x4d, 0xac,
	0x82, 0x47, 0x33, 0xd6, 0xe4, 0x82, 0x52, 0xd8, 0xf7, 0x05, 0xb8, 0xfc, 0x6e, 0x45, 0x91, 0x7f,
	0xa6, 0x01, 0x8a, 0x17, 0x8e, 0xd0, 0xc3, 0x64, 0xe9, 0x89, 0xb5, 0xf5, 0xe6, 0xbb, 0xaf, 0xc7,
	0x9c, 0x14, 0xce, 0x02, 0x95, 0x7a, 0x8c, 0x7b, 0xf4, 0x8a, 0x2a, 0xf5, 0x97, 0x1a, 0x54, 0x43,
	0x55, 0x27, 0x74, 0x3f, 0x65, 0x4d, 0x23, 0x25, 0xf7, 0xe6, 0xdb, 0xd7, 0xf2, 0x25, 0xa5, 0xf2,
	0xca, 0x0e, 0x90, 0x77, 0x9a, 0x9f, 0x68, 0x50, 0x0b, 0x57, 0xa9, 0x50, 0x8a, 0xec, 0x58, 0xc9,
	0xbe, 0xb9, 0x71, 0x3d, 0xe3, 0xf4, 0xe5, 0x09, 0xae, 0x33, 0x03, 0x28, 0x88, 0xba, 0x56, 0xd2,
	0xc6, 0x0f, 0x17, 0xfb, 0x93, 0x36, 0x7e, 0xa4, 0x28, 0x96, 0xb0, 0xf1, 0x1d, 0x7b, 0x40, 0x94,
	0x63, 0x26, 0x0a, 0x5f, 0x69, 0x68, 0xd3, 0x8f, 0x59, 0xa4, 0x6a, 0x96, 0x86, 0x16, 0x1c, 0x33,
	0x59, 0xf1, 0x42, 0x29, 0xc2, 0xae, 0x39, 0x66, 0xd1, 0x82, 0x59, 0xc2, 0x31, 0x63, 0x80, 0xca,
	0x31, 0x0b, 0x6a, 0x53, 0x49, 0xc7, 0x2c, 0xf6, 0x76, 0x91, 0x74, 0xcc, 0xe2, 0xe5, 0xad, 0x84,
	0x75, 0x64, 0xb8, 0xa1, 0x63, 0xb6, 0x94, 0x50, 0xc6, 0x42, 0xef, 0xa6, 0x18, 0x31, 0xf1, 0x49,
	0xa4, 0xf9, 0xde, 0x6b, 0x72, 0xa7, 0xee, 0x71, 0x6e, 0x7e, 0xb9, 0xc7, 0xff, 0x56, 0x83, 0xe5,
	0xa4, 0x12, 0x18, 0x4a, 0xc1, 0x49, 0x79, 0x4a, 0x69, 0x6e, 0xbe, 0x2e, 0xfb, 0x74, 0x6b, 0xf9,
	0xbb, 0xfe, 0x71, 0xfd, 0x5f, 0xbf, 0x5c, 0xd3, 0xfe, 0xe3, 0xcb, 0x35, 0xed, 0xbf, 0xbf, 0x5c,
	0xd3, 0xfe, 0xee, 0x7f, 0xd6, 0xe6, 0x4e, 0xf3, 0xec, 0x3f, 0x1a, 0x7f, 0xfb, 0x37, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xee, 0x4f, 0x63, 0x90, 0xed, 0x3c, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVClient interface {
	// Range gets the keys in the range from the key-value store.
	Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*RangeResponse, error)
	// Put puts the given key into the key-value store.
	// A put request increments the revision of the key-value store
	// and generates one event in the event history.
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// DeleteRange deletes the given range from the key-value store.
	// A delete request increments the revision of the key-value store
	// and generates a delete event in the event history for every deleted key.
	DeleteRange(ctx context.Context, in *DeleteRangeRequest, opts ...grpc.CallOption) (*DeleteRangeResponse, error)
	// Txn processes multiple requests in a single transaction.
	// A txn request increments the revision of the key-value store
	// and generates events with the same revision for every completed request.
	// It is not allowed to modify the same key several times within one txn.
	Txn(ctx context.Context, in *TxnRequest, opts ...grpc.CallOption) (*TxnResponse, error)
	// Compact compacts the event history in the etcd key-value store. The key-value
	// store should be periodically compacted or the event history will continue to grow
	// indefinitely.
	Compact(ctx context.Context, in *CompactionRequest, opts ...grpc.CallOption) (*CompactionResponse, error)
}

type kVClient struct {
	cc *grpc.ClientConn
}

func NewKVClient(cc *grpc.ClientConn) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*RangeResponse, error) {
	out := new(RangeResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.KV/Range", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.KV/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) DeleteRange(ctx context.Context, in *DeleteRangeRequest, opts ...grpc.CallOption) (*DeleteRangeResponse, error) {
	out := new(DeleteRangeResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.KV/DeleteRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Txn(ctx context.Context, in *TxnRequest, opts ...grpc.CallOption) (*TxnResponse, error) {
	out := new(TxnResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.KV/Txn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Compact(ctx context.Context, in *CompactionRequest, opts ...grpc.CallOption) (*CompactionResponse, error) {
	out := new(CompactionResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.KV/Compact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVServer is the server API for KV service.
type KVServer interface {
	// Range gets the keys in the range from the key-value store.
	Range(context.Context, *RangeRequest) (*RangeResponse, error)
	// Put puts the given key into the key-value store.
	// A put request increments the revision of the key-value store
	// and generates one event in the event history.
	Put(context.Context, *PutRequest) (*PutResponse, error)
	// DeleteRange deletes the given range from the key-value store.
	// A delete request increments the revision of the key-value store
	// and generates a delete event in the event history for every deleted key.
	DeleteRange(context.Context, *DeleteRangeRequest) (*DeleteRangeResponse, error)
	// Txn processes multiple requests in a single transaction.
	// A txn request increments the revision of the key-value store
	// and generates events with the same revision for every completed request.
	// It is not allowed to modify the same key several times within one txn.
	Txn(context.Context, *TxnRequest) (*TxnResponse, error)
	// Compact compacts the event history in the etcd key-value store. The key-value
	// store should be periodically compacted or the event history will continue to grow
	// indefinitely.
	Compact(context.Context, *CompactionRequest) (*CompactionResponse, error)
}

// UnimplementedKVServer can be embedded to have forward compatible implementations.
type UnimplementedKVServer struct{}

func (*UnimplementedKVServer) Range(ctx context.Context, req *RangeRequest) (*RangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Range not implemented")
}

func (*UnimplementedKVServer) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}

func (*UnimplementedKVServer) DeleteRange(ctx context.Context, req *DeleteRangeRequest) (*DeleteRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRange not implemented")
}

func (*UnimplementedKVServer) Txn(ctx context.Context, req *TxnRequest) (*TxnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Txn not implemented")
}

func (*UnimplementedKVServer) Compact(ctx context.Context, req *CompactionRequest) (*CompactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compact not implemented")
}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	s.RegisterService(&_KV_serviceDesc, srv)
}

func _KV_Range_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Range(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.KV/Range",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Range(ctx, req.(*RangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.KV/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_DeleteRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).DeleteRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.KV/DeleteRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).DeleteRange(ctx, req.(*DeleteRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Txn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Txn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.KV/Txn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Txn(ctx, req.(*TxnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Compact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Compact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.KV/Compact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Compact(ctx, req.(*CompactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "etcdserverpb.KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Range",
			Handler:    _KV_Range_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KV_Put_Handler,
		},
		{
			MethodName: "DeleteRange",
			Handler:    _KV_DeleteRange_Handler,
		},
		{
			MethodName: "Txn",
			Handler:    _KV_Txn_Handler,
		},
		{
			MethodName: "Compact",
			Handler:    _KV_Compact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// WatchClient is the client API for Watch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WatchClient interface {
	// Watch watches for events happening or that have happened. Both input and output
	// are streams; the input stream is for creating and canceling watchers and the output
	// stream sends events. One watch RPC can watch on multiple key ranges, streaming events
	// for several watches at once. The entire event history can be watched starting from the
	// last compaction revision.
	Watch(ctx context.Context, opts ...grpc.CallOption) (Watch_WatchClient, error)
}

type watchClient struct {
	cc *grpc.ClientConn
}

func NewWatchClient(cc *grpc.ClientConn) WatchClient {
	return &watchClient{cc}
}

func (c *watchClient) Watch(ctx context.Context, opts ...grpc.CallOption) (Watch_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Watch_serviceDesc.Streams[0], "/etcdserverpb.Watch/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchWatchClient{stream}
	return x, nil
}

type Watch_WatchClient interface {
	Send(*WatchRequest) error
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type watchWatchClient struct {
	grpc.ClientStream
}

func (x *watchWatchClient) Send(m *WatchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *watchWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatchServer is the server API for Watch service.
type WatchServer interface {
	// Watch watches for events happening or that have happened. Both input and output
	// are streams; the input stream is for creating and canceling watchers and the output
	// stream sends events. One watch RPC can watch on multiple key ranges, streaming events
	// for several watches at once. The entire event history can be watched starting from the
	// last compaction revision.
	Watch(Watch_WatchServer) error
}

// UnimplementedWatchServer can be embedded to have forward compatible implementations.
type UnimplementedWatchServer struct{}

func (*UnimplementedWatchServer) Watch(srv Watch_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterWatchServer(s *grpc.Server, srv WatchServer) {
	s.RegisterService(&_Watch_serviceDesc, srv)
}

func _Watch_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WatchServer).Watch(&watchWatchServer{stream})
}

type Watch_WatchServer interface {
	Send(*WatchResponse) error
	Recv() (*WatchRequest, error)
	grpc.ServerStream
}

type watchWatchServer struct {
	grpc.ServerStream
}

func (x *watchWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *watchWatchServer) Recv() (*WatchRequest, error) {
	m := new(WatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Watch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "etcdserverpb.Watch",
	HandlerType: (*WatchServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Watch_Watch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

// LeaseClient is the client API for Lease service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LeaseClient interface {
	// LeaseGrant creates a lease which expires if the server does not receive a keepAlive
	// within a given time to live period. All keys attached to the lease will be expired and
	// deleted if the lease expires. Each expired key generates a delete event in the event history.
	LeaseGrant(ctx context.Context, in *LeaseGrantRequest, opts ...grpc.CallOption) (*LeaseGrantResponse, error)
	// LeaseRevoke revokes a lease. All keys attached to the lease will expire and be deleted.
	LeaseRevoke(ctx context.Context, in *LeaseRevokeRequest, opts ...grpc.CallOption) (*LeaseRevokeResponse, error)
	// LeaseKeepAlive keeps the lease alive by streaming keep alive requests from the client
	// to the server and streaming keep alive responses from the server to the client.
	LeaseKeepAlive(ctx context.Context, opts ...grpc.CallOption) (Lease_LeaseKeepAliveClient, error)
	// LeaseTimeToLive retrieves lease information.
	LeaseTimeToLive(ctx context.Context, in *LeaseTimeToLiveRequest, opts ...grpc.CallOption) (*LeaseTimeToLiveResponse, error)
	// LeaseLeases lists all existing leases.
	LeaseLeases(ctx context.Context, in *LeaseLeasesRequest, opts ...grpc.CallOption) (*LeaseLeasesResponse, error)
}

type leaseClient struct {
	cc *grpc.ClientConn
}

func NewLeaseClient(cc *grpc.ClientConn) LeaseClient {
	return &leaseClient{cc}
}

func (c *leaseClient) LeaseGrant(ctx context.Context, in *LeaseGrantRequest, opts ...grpc.CallOption) (*LeaseGrantResponse, error) {
	out := new(LeaseGrantResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Lease/LeaseGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaseClient) LeaseRevoke(ctx context.Context, in *LeaseRevokeRequest, opts ...grpc.CallOption) (*LeaseRevokeResponse, error) {
	out := new(LeaseRevokeResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Lease/LeaseRevoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaseClient) LeaseKeepAlive(ctx context.Context, opts ...grpc.CallOption) (Lease_LeaseKeepAliveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Lease_serviceDesc.Streams[0], "/etcdserverpb.Lease/LeaseKeepAlive", opts...)
	if err != nil {
		return nil, err
	}
	x := &leaseLeaseKeepAliveClient{stream}
	return x, nil
}

type Lease_LeaseKeepAliveClient interface {
	Send(*LeaseKeepAliveRequest) error
	Recv() (*LeaseKeepAliveResponse, error)
	grpc.ClientStream
}

type leaseLeaseKeepAliveClient struct {
	grpc.ClientStream
}

func (x *leaseLeaseKeepAliveClient) Send(m *LeaseKeepAliveRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *leaseLeaseKeepAliveClient) Recv() (*LeaseKeepAliveResponse, error) {
	m := new(LeaseKeepAliveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *leaseClient) LeaseTimeToLive(ctx context.Context, in *LeaseTimeToLiveRequest, opts ...grpc.CallOption) (*LeaseTimeToLiveResponse, error) {
	out := new(LeaseTimeToLiveResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Lease/LeaseTimeToLive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaseClient) LeaseLeases(ctx context.Context, in *LeaseLeasesRequest, opts ...grpc.CallOption) (*LeaseLeasesResponse, error) {
	out := new(LeaseLeasesResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Lease/LeaseLeases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeaseServer is the server API for Lease service.
type LeaseServer interface {
	// LeaseGrant creates a lease which expires if the server does not receive a keepAlive
	// within a given time to live period. All keys attached to the lease will be expired and
	// deleted if the lease expires. Each expired key generates a delete event in the event history.
	LeaseGrant(context.Context, *LeaseGrantRequest) (*LeaseGrantResponse, error)
	// LeaseRevoke revokes a lease. All keys attached to the lease will expire and be deleted.
	LeaseRevoke(context.Context, *LeaseRevokeRequest) (*LeaseRevokeResponse, error)
	// LeaseKeepAlive keeps the lease alive by streaming keep alive requests from the client
	// to the server and streaming keep alive responses from the server to the client.
	LeaseKeepAlive(Lease_LeaseKeepAliveServer) error
	// LeaseTimeToLive retrieves lease information.
	LeaseTimeToLive(context.Context, *LeaseTimeToLiveRequest) (*LeaseTimeToLiveResponse, error)
	// LeaseLeases lists all existing leases.
	LeaseLeases(context.Context, *LeaseLeasesRequest) (*LeaseLeasesResponse, error)
}

// UnimplementedLeaseServer can be embedded to have forward compatible implementations.
type UnimplementedLeaseServer struct{}

func (*UnimplementedLeaseServer) LeaseGrant(ctx context.Context, req *LeaseGrantRequest) (*LeaseGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaseGrant not implemented")
}

func (*UnimplementedLeaseServer) LeaseRevoke(ctx context.Context, req *LeaseRevokeRequest) (*LeaseRevokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaseRevoke not implemented")
}

func (*UnimplementedLeaseServer) LeaseKeepAlive(srv Lease_LeaseKeepAliveServer) error {
	return status.Errorf(codes.Unimplemented, "method LeaseKeepAlive not implemented")
}

func (*UnimplementedLeaseServer) LeaseTimeToLive(ctx context.Context, req *LeaseTimeToLiveRequest) (*LeaseTimeToLiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaseTimeToLive not implemented")
}

func (*UnimplementedLeaseServer) LeaseLeases(ctx context.Context, req *LeaseLeasesRequest) (*LeaseLeasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaseLeases not implemented")
}

func RegisterLeaseServer(s *grpc.Server, srv LeaseServer) {
	s.RegisterService(&_Lease_serviceDesc, srv)
}

func _Lease_LeaseGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaseGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServer).LeaseGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Lease/LeaseGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServer).LeaseGrant(ctx, req.(*LeaseGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lease_LeaseRevoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaseRevokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServer).LeaseRevoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Lease/LeaseRevoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServer).LeaseRevoke(ctx, req.(*LeaseRevokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lease_LeaseKeepAlive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LeaseServer).LeaseKeepAlive(&leaseLeaseKeepAliveServer{stream})
}

type Lease_LeaseKeepAliveServer interface {
	Send(*LeaseKeepAliveResponse) error
	Recv() (*LeaseKeepAliveRequest, error)
	grpc.ServerStream
}

type leaseLeaseKeepAliveServer struct {
	grpc.ServerStream
}

func (x *leaseLeaseKeepAliveServer) Send(m *LeaseKeepAliveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *leaseLeaseKeepAliveServer) Recv() (*LeaseKeepAliveRequest, error) {
	m := new(LeaseKeepAliveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Lease_LeaseTimeToLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaseTimeToLiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServer).LeaseTimeToLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Lease/LeaseTimeToLive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServer).LeaseTimeToLive(ctx, req.(*LeaseTimeToLiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lease_LeaseLeases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaseLeasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServer).LeaseLeases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Lease/LeaseLeases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServer).LeaseLeases(ctx, req.(*LeaseLeasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lease_serviceDesc = grpc.ServiceDesc{
	ServiceName: "etcdserverpb.Lease",
	HandlerType: (*LeaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LeaseGrant",
			Handler:    _Lease_LeaseGrant_Handler,
		},
		{
			MethodName: "LeaseRevoke",
			Handler:    _Lease_LeaseRevoke_Handler,
		},
		{
			MethodName: "LeaseTimeToLive",
			Handler:    _Lease_LeaseTimeToLive_Handler,
		},
		{
			MethodName: "LeaseLeases",
			Handler:    _Lease_LeaseLeases_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LeaseKeepAlive",
			Handler:       _Lease_LeaseKeepAlive_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

// ClusterClient is the client API for Cluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterClient interface {
	// MemberAdd adds a member into the cluster.
	MemberAdd(ctx context.Context, in *MemberAddRequest, opts ...grpc.CallOption) (*MemberAddResponse, error)
	// MemberRemove removes an existing member from the cluster.
	MemberRemove(ctx context.Context, in *MemberRemoveRequest, opts ...grpc.CallOption) (*MemberRemoveResponse, error)
	// MemberUpdate updates the member configuration.
	MemberUpdate(ctx context.Context, in *MemberUpdateRequest, opts ...grpc.CallOption) (*MemberUpdateResponse, error)
	// MemberList lists all the members in the cluster.
	MemberList(ctx context.Context, in *MemberListRequest, opts ...grpc.CallOption) (*MemberListResponse, error)
	// MemberPromote promotes a member from raft learner (non-voting) to raft voting member.
	MemberPromote(ctx context.Context, in *MemberPromoteRequest, opts ...grpc.CallOption) (*MemberPromoteResponse, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) MemberAdd(ctx context.Context, in *MemberAddRequest, opts ...grpc.CallOption) (*MemberAddResponse, error) {
	out := new(MemberAddResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Cluster/MemberAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) MemberRemove(ctx context.Context, in *MemberRemoveRequest, opts ...grpc.CallOption) (*MemberRemoveResponse, error) {
	out := new(MemberRemoveResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Cluster/MemberRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) MemberUpdate(ctx context.Context, in *MemberUpdateRequest, opts ...grpc.CallOption) (*MemberUpdateResponse, error) {
	out := new(MemberUpdateResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Cluster/MemberUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) MemberList(ctx context.Context, in *MemberListRequest, opts ...grpc.CallOption) (*MemberListResponse, error) {
	out := new(MemberListResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Cluster/MemberList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) MemberPromote(ctx context.Context, in *MemberPromoteRequest, opts ...grpc.CallOption) (*MemberPromoteResponse, error) {
	out := new(MemberPromoteResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Cluster/MemberPromote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServer is the server API for Cluster service.
type ClusterServer interface {
	// MemberAdd adds a member into the cluster.
	MemberAdd(context.Context, *MemberAddRequest) (*MemberAddResponse, error)
	// MemberRemove removes an existing member from the cluster.
	MemberRemove(context.Context, *MemberRemoveRequest) (*MemberRemoveResponse, error)
	// MemberUpdate updates the member configuration.
	MemberUpdate(context.Context, *MemberUpdateRequest) (*MemberUpdateResponse, error)
	// MemberList lists all the members in the cluster.
	MemberList(context.Context, *MemberListRequest) (*MemberListResponse, error)
	// MemberPromote promotes a member from raft learner (non-voting) to raft voting member.
	MemberPromote(context.Context, *MemberPromoteRequest) (*MemberPromoteResponse, error)
}

// UnimplementedClusterServer can be embedded to have forward compatible implementations.
type UnimplementedClusterServer struct{}

func (*UnimplementedClusterServer) MemberAdd(ctx context.Context, req *MemberAddRequest) (*MemberAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberAdd not implemented")
}

func (*UnimplementedClusterServer) MemberRemove(ctx context.Context, req *MemberRemoveRequest) (*MemberRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberRemove not implemented")
}

func (*UnimplementedClusterServer) MemberUpdate(ctx context.Context, req *MemberUpdateRequest) (*MemberUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberUpdate not implemented")
}

func (*UnimplementedClusterServer) MemberList(ctx context.Context, req *MemberListRequest) (*MemberListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberList not implemented")
}

func (*UnimplementedClusterServer) MemberPromote(ctx context.Context, req *MemberPromoteRequest) (*MemberPromoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberPromote not implemented")
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_MemberAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Cluster/MemberAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberAdd(ctx, req.(*MemberAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_MemberRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Cluster/MemberRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberRemove(ctx, req.(*MemberRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_MemberUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Cluster/MemberUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberUpdate(ctx, req.(*MemberUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_MemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Cluster/MemberList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberList(ctx, req.(*MemberListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_MemberPromote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberPromoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberPromote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Cluster/MemberPromote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberPromote(ctx, req.(*MemberPromoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "etcdserverpb.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MemberAdd",
			Handler:    _Cluster_MemberAdd_Handler,
		},
		{
			MethodName: "MemberRemove",
			Handler:    _Cluster_MemberRemove_Handler,
		},
		{
			MethodName: "MemberUpdate",
			Handler:    _Cluster_MemberUpdate_Handler,
		},
		{
			MethodName: "MemberList",
			Handler:    _Cluster_MemberList_Handler,
		},
		{
			MethodName: "MemberPromote",
			Handler:    _Cluster_MemberPromote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// MaintenanceClient is the client API for Maintenance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MaintenanceClient interface {
	// Alarm activates, deactivates, and queries alarms regarding cluster health.
	Alarm(ctx context.Context, in *AlarmRequest, opts ...grpc.CallOption) (*AlarmResponse, error)
	// Status gets the status of the member.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	// Defragment defragments a member's backend database to recover storage space.
	Defragment(ctx context.Context, in *DefragmentRequest, opts ...grpc.CallOption) (*DefragmentResponse, error)
	// Hash computes the hash of whole backend keyspace,
	// including key, lease, and other buckets in storage.
	// This is designed for testing ONLY!
	// Do not rely on this in production with ongoing transactions,
	// since Hash operation does not hold MVCC locks.
	// Use "HashKV" API instead for "key" bucket consistency checks.
	Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error)
	// HashKV computes the hash of all MVCC keys up to a given revision.
	// It only iterates "key" bucket in backend storage.
	HashKV(ctx context.Context, in *HashKVRequest, opts ...grpc.CallOption) (*HashKVResponse, error)
	// Snapshot sends a snapshot of the entire backend from a member over a stream to a client.
	Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (Maintenance_SnapshotClient, error)
	// MoveLeader requests current leader node to transfer its leadership to transferee.
	MoveLeader(ctx context.Context, in *MoveLeaderRequest, opts ...grpc.CallOption) (*MoveLeaderResponse, error)
	// Downgrade requests downgrades, verifies feasibility or cancels downgrade
	// on the cluster version.
	// Supported since etcd 3.5.
	Downgrade(ctx context.Context, in *DowngradeRequest, opts ...grpc.CallOption) (*DowngradeResponse, error)
}

type maintenanceClient struct {
	cc *grpc.ClientConn
}

func NewMaintenanceClient(cc *grpc.ClientConn) MaintenanceClient {
	return &maintenanceClient{cc}
}

func (c *maintenanceClient) Alarm(ctx context.Context, in *AlarmRequest, opts ...grpc.CallOption) (*AlarmResponse, error) {
	out := new(AlarmResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Maintenance/Alarm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintenanceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Maintenance/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintenanceClient) Defragment(ctx context.Context, in *DefragmentRequest, opts ...grpc.CallOption) (*DefragmentResponse, error) {
	out := new(DefragmentResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Maintenance/Defragment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintenanceClient) Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Maintenance/Hash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintenanceClient) HashKV(ctx context.Context, in *HashKVRequest, opts ...grpc.CallOption) (*HashKVResponse, error) {
	out := new(HashKVResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Maintenance/HashKV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintenanceClient) Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (Maintenance_SnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Maintenance_serviceDesc.Streams[0], "/etcdserverpb.Maintenance/Snapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &maintenanceSnapshotClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Maintenance_SnapshotClient interface {
	Recv() (*SnapshotResponse, error)
	grpc.ClientStream
}

type maintenanceSnapshotClient struct {
	grpc.ClientStream
}

func (x *maintenanceSnapshotClient) Recv() (*SnapshotResponse, error) {
	m := new(SnapshotResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *maintenanceClient) MoveLeader(ctx context.Context, in *MoveLeaderRequest, opts ...grpc.CallOption) (*MoveLeaderResponse, error) {
	out := new(MoveLeaderResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Maintenance/MoveLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintenanceClient) Downgrade(ctx context.Context, in *DowngradeRequest, opts ...grpc.CallOption) (*DowngradeResponse, error) {
	out := new(DowngradeResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Maintenance/Downgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MaintenanceServer is the server API for Maintenance service.
type MaintenanceServer interface {
	// Alarm activates, deactivates, and queries alarms regarding cluster health.
	Alarm(context.Context, *AlarmRequest) (*AlarmResponse, error)
	// Status gets the status of the member.
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	// Defragment defragments a member's backend database to recover storage space.
	Defragment(context.Context, *DefragmentRequest) (*DefragmentResponse, error)
	// Hash computes the hash of whole backend keyspace,
	// including key, lease, and other buckets in storage.
	// This is designed for testing ONLY!
	// Do not rely on this in production with ongoing transactions,
	// since Hash operation does not hold MVCC locks.
	// Use "HashKV" API instead for "key" bucket consistency checks.
	Hash(context.Context, *HashRequest) (*HashResponse, error)
	// HashKV computes the hash of all MVCC keys up to a given revision.
	// It only iterates "key" bucket in backend storage.
	HashKV(context.Context, *HashKVRequest) (*HashKVResponse, error)
	// Snapshot sends a snapshot of the entire backend from a member over a stream to a client.
	Snapshot(*SnapshotRequest, Maintenance_SnapshotServer) error
	// MoveLeader requests current leader node to transfer its leadership to transferee.
	MoveLeader(context.Context, *MoveLeaderRequest) (*MoveLeaderResponse, error)
	// Downgrade requests downgrades, verifies feasibility or cancels downgrade
	// on the cluster version.
	// Supported since etcd 3.5.
	Downgrade(context.Context, *DowngradeRequest) (*DowngradeResponse, error)
}

// UnimplementedMaintenanceServer can be embedded to have forward compatible implementations.
type UnimplementedMaintenanceServer struct{}

func (*UnimplementedMaintenanceServer) Alarm(ctx context.Context, req *AlarmRequest) (*AlarmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alarm not implemented")
}

func (*UnimplementedMaintenanceServer) Status(ctx context.Context, req *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}

func (*UnimplementedMaintenanceServer) Defragment(ctx context.Context, req *DefragmentRequest) (*DefragmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Defragment not implemented")
}

func (*UnimplementedMaintenanceServer) Hash(ctx context.Context, req *HashRequest) (*HashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hash not implemented")
}

func (*UnimplementedMaintenanceServer) HashKV(ctx context.Context, req *HashKVRequest) (*HashKVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HashKV not implemented")
}

func (*UnimplementedMaintenanceServer) Snapshot(req *SnapshotRequest, srv Maintenance_SnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method Snapshot not implemented")
}

func (*UnimplementedMaintenanceServer) MoveLeader(ctx context.Context, req *MoveLeaderRequest) (*MoveLeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveLeader not implemented")
}

func (*UnimplementedMaintenanceServer) Downgrade(ctx context.Context, req *DowngradeRequest) (*DowngradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Downgrade not implemented")
}

func RegisterMaintenanceServer(s *grpc.Server, srv MaintenanceServer) {
	s.RegisterService(&_Maintenance_serviceDesc, srv)
}

func _Maintenance_Alarm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Alarm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Maintenance/Alarm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Alarm(ctx, req.(*AlarmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maintenance_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Maintenance/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maintenance_Defragment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefragmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Defragment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Maintenance/Defragment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Defragment(ctx, req.(*DefragmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maintenance_Hash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Hash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Maintenance/Hash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Hash(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maintenance_HashKV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashKVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).HashKV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Maintenance/HashKV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).HashKV(ctx, req.(*HashKVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maintenance_Snapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SnapshotRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MaintenanceServer).Snapshot(m, &maintenanceSnapshotServer{stream})
}

type Maintenance_SnapshotServer interface {
	Send(*SnapshotResponse) error
	grpc.ServerStream
}

type maintenanceSnapshotServer struct {
	grpc.ServerStream
}

func (x *maintenanceSnapshotServer) Send(m *SnapshotResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Maintenance_MoveLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).MoveLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Maintenance/MoveLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).MoveLeader(ctx, req.(*MoveLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Maintenance_Downgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DowngradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Downgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Maintenance/Downgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Downgrade(ctx, req.(*DowngradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Maintenance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "etcdserverpb.Maintenance",
	HandlerType: (*MaintenanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Alarm",
			Handler:    _Maintenance_Alarm_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Maintenance_Status_Handler,
		},
		{
			MethodName: "Defragment",
			Handler:    _Maintenance_Defragment_Handler,
		},
		{
			MethodName: "Hash",
			Handler:    _Maintenance_Hash_Handler,
		},
		{
			MethodName: "HashKV",
			Handler:    _Maintenance_HashKV_Handler,
		},
		{
			MethodName: "MoveLeader",
			Handler:    _Maintenance_MoveLeader_Handler,
		},
		{
			MethodName: "Downgrade",
			Handler:    _Maintenance_Downgrade_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Snapshot",
			Handler:       _Maintenance_Snapshot_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	// AuthEnable enables authentication.
	AuthEnable(ctx context.Context, in *AuthEnableRequest, opts ...grpc.CallOption) (*AuthEnableResponse, error)
	// AuthDisable disables authentication.
	AuthDisable(ctx context.Context, in *AuthDisableRequest, opts ...grpc.CallOption) (*AuthDisableResponse, error)
	// AuthStatus displays authentication status.
	AuthStatus(ctx context.Context, in *AuthStatusRequest, opts ...grpc.CallOption) (*AuthStatusResponse, error)
	// Authenticate processes an authenticate request.
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	// UserAdd adds a new user. User name cannot be empty.
	UserAdd(ctx context.Context, in *AuthUserAddRequest, opts ...grpc.CallOption) (*AuthUserAddResponse, error)
	// UserGet gets detailed user information.
	UserGet(ctx context.Context, in *AuthUserGetRequest, opts ...grpc.CallOption) (*AuthUserGetResponse, error)
	// UserList gets a list of all users.
	UserList(ctx context.Context, in *AuthUserListRequest, opts ...grpc.CallOption) (*AuthUserListResponse, error)
	// UserDelete deletes a specified user.
	UserDelete(ctx context.Context, in *AuthUserDeleteRequest, opts ...grpc.CallOption) (*AuthUserDeleteResponse, error)
	// UserChangePassword changes the password of a specified user.
	UserChangePassword(ctx context.Context, in *AuthUserChangePasswordRequest, opts ...grpc.CallOption) (*AuthUserChangePasswordResponse, error)
	// UserGrant grants a role to a specified user.
	UserGrantRole(ctx context.Context, in *AuthUserGrantRoleRequest, opts ...grpc.CallOption) (*AuthUserGrantRoleResponse, error)
	// UserRevokeRole revokes a role of specified user.
	UserRevokeRole(ctx context.Context, in *AuthUserRevokeRoleRequest, opts ...grpc.CallOption) (*AuthUserRevokeRoleResponse, error)
	// RoleAdd adds a new role. Role name cannot be empty.
	RoleAdd(ctx context.Context, in *AuthRoleAddRequest, opts ...grpc.CallOption) (*AuthRoleAddResponse, error)
	// RoleGet gets detailed role information.
	RoleGet(ctx context.Context, in *AuthRoleGetRequest, opts ...grpc.CallOption) (*AuthRoleGetResponse, error)
	// RoleList gets lists of all roles.
	RoleList(ctx context.Context, in *AuthRoleListRequest, opts ...grpc.CallOption) (*AuthRoleListResponse, error)
	// RoleDelete deletes a specified role.
	RoleDelete(ctx context.Context, in *AuthRoleDeleteRequest, opts ...grpc.CallOption) (*AuthRoleDeleteResponse, error)
	// RoleGrantPermission grants a permission of a specified key or range to a specified role.
	RoleGrantPermission(ctx context.Context, in *AuthRoleGrantPermissionRequest, opts ...grpc.CallOption) (*AuthRoleGrantPermissionResponse, error)
	// RoleRevokePermission revokes a key or range permission of a specified role.
	RoleRevokePermission(ctx context.Context, in *AuthRoleRevokePermissionRequest, opts ...grpc.CallOption) (*AuthRoleRevokePermissionResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) AuthEnable(ctx context.Context, in *AuthEnableRequest, opts ...grpc.CallOption) (*AuthEnableResponse, error) {
	out := new(AuthEnableResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/AuthEnable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthDisable(ctx context.Context, in *AuthDisableRequest, opts ...grpc.CallOption) (*AuthDisableResponse, error) {
	out := new(AuthDisableResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/AuthDisable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthStatus(ctx context.Context, in *AuthStatusRequest, opts ...grpc.CallOption) (*AuthStatusResponse, error) {
	out := new(AuthStatusResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/AuthStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserAdd(ctx context.Context, in *AuthUserAddRequest, opts ...grpc.CallOption) (*AuthUserAddResponse, error) {
	out := new(AuthUserAddResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/UserAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserGet(ctx context.Context, in *AuthUserGetRequest, opts ...grpc.CallOption) (*AuthUserGetResponse, error) {
	out := new(AuthUserGetResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/UserGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserList(ctx context.Context, in *AuthUserListRequest, opts ...grpc.CallOption) (*AuthUserListResponse, error) {
	out := new(AuthUserListResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserDelete(ctx context.Context, in *AuthUserDeleteRequest, opts ...grpc.CallOption) (*AuthUserDeleteResponse, error) {
	out := new(AuthUserDeleteResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserChangePassword(ctx context.Context, in *AuthUserChangePasswordRequest, opts ...grpc.CallOption) (*AuthUserChangePasswordResponse, error) {
	out := new(AuthUserChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/UserChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserGrantRole(ctx context.Context, in *AuthUserGrantRoleRequest, opts ...grpc.CallOption) (*AuthUserGrantRoleResponse, error) {
	out := new(AuthUserGrantRoleResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/UserGrantRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserRevokeRole(ctx context.Context, in *AuthUserRevokeRoleRequest, opts ...grpc.CallOption) (*AuthUserRevokeRoleResponse, error) {
	out := new(AuthUserRevokeRoleResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/UserRevokeRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RoleAdd(ctx context.Context, in *AuthRoleAddRequest, opts ...grpc.CallOption) (*AuthRoleAddResponse, error) {
	out := new(AuthRoleAddResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/RoleAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RoleGet(ctx context.Context, in *AuthRoleGetRequest, opts ...grpc.CallOption) (*AuthRoleGetResponse, error) {
	out := new(AuthRoleGetResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/RoleGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RoleList(ctx context.Context, in *AuthRoleListRequest, opts ...grpc.CallOption) (*AuthRoleListResponse, error) {
	out := new(AuthRoleListResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/RoleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RoleDelete(ctx context.Context, in *AuthRoleDeleteRequest, opts ...grpc.CallOption) (*AuthRoleDeleteResponse, error) {
	out := new(AuthRoleDeleteResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/RoleDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RoleGrantPermission(ctx context.Context, in *AuthRoleGrantPermissionRequest, opts ...grpc.CallOption) (*AuthRoleGrantPermissionResponse, error) {
	out := new(AuthRoleGrantPermissionResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/RoleGrantPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RoleRevokePermission(ctx context.Context, in *AuthRoleRevokePermissionRequest, opts ...grpc.CallOption) (*AuthRoleRevokePermissionResponse, error) {
	out := new(AuthRoleRevokePermissionResponse)
	err := c.cc.Invoke(ctx, "/etcdserverpb.Auth/RoleRevokePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	// AuthEnable enables authentication.
	AuthEnable(context.Context, *AuthEnableRequest) (*AuthEnableResponse, error)
	// AuthDisable disables authentication.
	AuthDisable(context.Context, *AuthDisableRequest) (*AuthDisableResponse, error)
	// AuthStatus displays authentication status.
	AuthStatus(context.Context, *AuthStatusRequest) (*AuthStatusResponse, error)
	// Authenticate processes an authenticate request.
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	// UserAdd adds a new user. User name cannot be empty.
	UserAdd(context.Context, *AuthUserAddRequest) (*AuthUserAddResponse, error)
	// UserGet gets detailed user information.
	UserGet(context.Context, *AuthUserGetRequest) (*AuthUserGetResponse, error)
	// UserList gets a list of all users.
	UserList(context.Context, *AuthUserListRequest) (*AuthUserListResponse, error)
	// UserDelete deletes a specified user.
	UserDelete(context.Context, *AuthUserDeleteRequest) (*AuthUserDeleteResponse, error)
	// UserChangePassword changes the password of a specified user.
	UserChangePassword(context.Context, *AuthUserChangePasswordRequest) (*AuthUserChangePasswordResponse, error)
	// UserGrant grants a role to a specified user.
	UserGrantRole(context.Context, *AuthUserGrantRoleRequest) (*AuthUserGrantRoleResponse, error)
	// UserRevokeRole revokes a role of specified user.
	UserRevokeRole(context.Context, *AuthUserRevokeRoleRequest) (*AuthUserRevokeRoleResponse, error)
	// RoleAdd adds a new role. Role name cannot be empty.
	RoleAdd(context.Context, *AuthRoleAddRequest) (*AuthRoleAddResponse, error)
	// RoleGet gets detailed role information.
	RoleGet(context.Context, *AuthRoleGetRequest) (*AuthRoleGetResponse, error)
	// RoleList gets lists of all roles.
	RoleList(context.Context, *AuthRoleListRequest) (*AuthRoleListResponse, error)
	// RoleDelete deletes a specified role.
	RoleDelete(context.Context, *AuthRoleDeleteRequest) (*AuthRoleDeleteResponse, error)
	// RoleGrantPermission grants a permission of a specified key or range to a specified role.
	RoleGrantPermission(context.Context, *AuthRoleGrantPermissionRequest) (*AuthRoleGrantPermissionResponse, error)
	// RoleRevokePermission revokes a key or range permission of a specified role.
	RoleRevokePermission(context.Context, *AuthRoleRevokePermissionRequest) (*AuthRoleRevokePermissionResponse, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct{}

func (*UnimplementedAuthServer) AuthEnable(ctx context.Context, req *AuthEnableRequest) (*AuthEnableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthEnable not implemented")
}

func (*UnimplementedAuthServer) AuthDisable(ctx context.Context, req *AuthDisableRequest) (*AuthDisableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthDisable not implemented")
}

func (*UnimplementedAuthServer) AuthStatus(ctx context.Context, req *AuthStatusRequest) (*AuthStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthStatus not implemented")
}

func (*UnimplementedAuthServer) Authenticate(ctx context.Context, req *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

func (*UnimplementedAuthServer) UserAdd(ctx context.Context, req *AuthUserAddRequest) (*AuthUserAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}

func (*UnimplementedAuthServer) UserGet(ctx context.Context, req *AuthUserGetRequest) (*AuthUserGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGet not implemented")
}

func (*UnimplementedAuthServer) UserList(ctx context.Context, req *AuthUserListRequest) (*AuthUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}

func (*UnimplementedAuthServer) UserDelete(ctx context.Context, req *AuthUserDeleteRequest) (*AuthUserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}

func (*UnimplementedAuthServer) UserChangePassword(ctx context.Context, req *AuthUserChangePasswordRequest) (*AuthUserChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserChangePassword not implemented")
}

func (*UnimplementedAuthServer) UserGrantRole(ctx context.Context, req *AuthUserGrantRoleRequest) (*AuthUserGrantRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGrantRole not implemented")
}

func (*UnimplementedAuthServer) UserRevokeRole(ctx context.Context, req *AuthUserRevokeRoleRequest) (*AuthUserRevokeRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRevokeRole not implemented")
}

func (*UnimplementedAuthServer) RoleAdd(ctx context.Context, req *AuthRoleAddRequest) (*AuthRoleAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleAdd not implemented")
}

func (*UnimplementedAuthServer) RoleGet(ctx context.Context, req *AuthRoleGetRequest) (*AuthRoleGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleGet not implemented")
}

func (*UnimplementedAuthServer) RoleList(ctx context.Context, req *AuthRoleListRequest) (*AuthRoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleList not implemented")
}

func (*UnimplementedAuthServer) RoleDelete(ctx context.Context, req *AuthRoleDeleteRequest) (*AuthRoleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleDelete not implemented")
}

func (*UnimplementedAuthServer) RoleGrantPermission(ctx context.Context, req *AuthRoleGrantPermissionRequest) (*AuthRoleGrantPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleGrantPermission not implemented")
}

func (*UnimplementedAuthServer) RoleRevokePermission(ctx context.Context, req *AuthRoleRevokePermissionRequest) (*AuthRoleRevokePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleRevokePermission not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_AuthEnable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthEnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthEnable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/AuthEnable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthEnable(ctx, req.(*AuthEnableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthDisable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthDisableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthDisable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/AuthDisable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthDisable(ctx, req.(*AuthDisableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/AuthStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthStatus(ctx, req.(*AuthStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/UserAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserAdd(ctx, req.(*AuthUserAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UserGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/UserGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserGet(ctx, req.(*AuthUserGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserList(ctx, req.(*AuthUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserDelete(ctx, req.(*AuthUserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UserChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/UserChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserChangePassword(ctx, req.(*AuthUserChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UserGrantRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserGrantRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserGrantRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/UserGrantRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserGrantRole(ctx, req.(*AuthUserGrantRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UserRevokeRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserRevokeRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserRevokeRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/UserRevokeRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserRevokeRole(ctx, req.(*AuthUserRevokeRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RoleAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRoleAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/RoleAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleAdd(ctx, req.(*AuthRoleAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RoleGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRoleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/RoleGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleGet(ctx, req.(*AuthRoleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/RoleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleList(ctx, req.(*AuthRoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RoleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/RoleDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleDelete(ctx, req.(*AuthRoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RoleGrantPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRoleGrantPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleGrantPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/RoleGrantPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleGrantPermission(ctx, req.(*AuthRoleGrantPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RoleRevokePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRoleRevokePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleRevokePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etcdserverpb.Auth/RoleRevokePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleRevokePermission(ctx, req.(*AuthRoleRevokePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "etcdserverpb.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthEnable",
			Handler:    _Auth_AuthEnable_Handler,
		},
		{
			MethodName: "AuthDisable",
			Handler:    _Auth_AuthDisable_Handler,
		},
		{
			MethodName: "AuthStatus",
			Handler:    _Auth_AuthStatus_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _Auth_Authenticate_Handler,
		},
		{
			MethodName: "UserAdd",
			Handler:    _Auth_UserAdd_Handler,
		},
		{
			MethodName: "UserGet",
			Handler:    _Auth_UserGet_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _Auth_UserList_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _Auth_UserDelete_Handler,
		},
		{
			MethodName: "UserChangePassword",
			Handler:    _Auth_UserChangePassword_Handler,
		},
		{
			MethodName: "UserGrantRole",
			Handler:    _Auth_UserGrantRole_Handler,
		},
		{
			MethodName: "UserRevokeRole",
			Handler:    _Auth_UserRevokeRole_Handler,
		},
		{
			MethodName: "RoleAdd",
			Handler:    _Auth_RoleAdd_Handler,
		},
		{
			MethodName: "RoleGet",
			Handler:    _Auth_RoleGet_Handler,
		},
		{
			MethodName: "RoleList",
			Handler:    _Auth_RoleList_Handler,
		},
		{
			MethodName: "RoleDelete",
			Handler:    _Auth_RoleDelete_Handler,
		},
		{
			MethodName: "RoleGrantPermission",
			Handler:    _Auth_RoleGrantPermission_Handler,
		},
		{
			MethodName: "RoleRevokePermission",
			Handler:    _Auth_RoleRevokePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func (m *ResponseHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RaftTerm != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.RaftTerm))
		i--
		dAtA[i] = 0x20
	}
	if m.Revision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Revision))
		i--
		dAtA[i] = 0x18
	}
	if m.MemberId != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MemberId))
		i--
		dAtA[i] = 0x10
	}
	if m.ClusterId != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ClusterId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RangeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RangeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RangeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxCreateRevision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MaxCreateRevision))
		i--
		dAtA[i] = 0x68
	}
	if m.MinCreateRevision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MinCreateRevision))
		i--
		dAtA[i] = 0x60
	}
	if m.MaxModRevision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MaxModRevision))
		i--
		dAtA[i] = 0x58
	}
	if m.MinModRevision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MinModRevision))
		i--
		dAtA[i] = 0x50
	}
	if m.CountOnly {
		i--
		if m.CountOnly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.KeysOnly {
		i--
		if m.KeysOnly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.Serializable {
		i--
		if m.Serializable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.SortTarget != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.SortTarget))
		i--
		dAtA[i] = 0x30
	}
	if m.SortOrder != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.SortOrder))
		i--
		dAtA[i] = 0x28
	}
	if m.Revision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Revision))
		i--
		dAtA[i] = 0x20
	}
	if m.Limit != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x18
	}
	if len(m.RangeEnd) > 0 {
		i -= len(m.RangeEnd)
		copy(dAtA[i:], m.RangeEnd)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RangeEnd)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RangeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RangeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RangeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Count != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x20
	}
	if m.More {
		i--
		if m.More {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Kvs) > 0 {
		for iNdEx := len(m.Kvs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Kvs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PutRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PutRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PutRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IgnoreLease {
		i--
		if m.IgnoreLease {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.IgnoreValue {
		i--
		if m.IgnoreValue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.PrevKv {
		i--
		if m.PrevKv {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Lease != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Lease))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PutResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PutResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PutResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PrevKv != nil {
		{
			size, err := m.PrevKv.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeleteRangeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteRangeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeleteRangeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PrevKv {
		i--
		if m.PrevKv {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.RangeEnd) > 0 {
		i -= len(m.RangeEnd)
		copy(dAtA[i:], m.RangeEnd)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RangeEnd)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeleteRangeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteRangeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeleteRangeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PrevKvs) > 0 {
		for iNdEx := len(m.PrevKvs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PrevKvs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Deleted != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Deleted))
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestOp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestOp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestOp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Request != nil {
		{
			size := m.Request.Size()
			i -= size
			if _, err := m.Request.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *RequestOp_RequestRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestOp_RequestRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RequestRange != nil {
		{
			size, err := m.RequestRange.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestOp_RequestPut) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestOp_RequestPut) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RequestPut != nil {
		{
			size, err := m.RequestPut.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *RequestOp_RequestDeleteRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestOp_RequestDeleteRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RequestDeleteRange != nil {
		{
			size, err := m.RequestDeleteRange.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}

func (m *RequestOp_RequestTxn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestOp_RequestTxn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RequestTxn != nil {
		{
			size, err := m.RequestTxn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}

func (m *ResponseOp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseOp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseOp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Response != nil {
		{
			size := m.Response.Size()
			i -= size
			if _, err := m.Response.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ResponseOp_ResponseRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseOp_ResponseRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ResponseRange != nil {
		{
			size, err := m.ResponseRange.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseOp_ResponsePut) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseOp_ResponsePut) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ResponsePut != nil {
		{
			size, err := m.ResponsePut.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *ResponseOp_ResponseDeleteRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseOp_ResponseDeleteRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ResponseDeleteRange != nil {
		{
			size, err := m.ResponseDeleteRange.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}

func (m *ResponseOp_ResponseTxn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseOp_ResponseTxn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ResponseTxn != nil {
		{
			size, err := m.ResponseTxn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}

func (m *Compare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Compare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Compare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RangeEnd) > 0 {
		i -= len(m.RangeEnd)
		copy(dAtA[i:], m.RangeEnd)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RangeEnd)))
		i--
		dAtA[i] = 0x4
		i--
		dAtA[i] = 0x82
	}
	if m.TargetUnion != nil {
		{
			size := m.TargetUnion.Size()
			i -= size
			if _, err := m.TargetUnion.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Target != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Target))
		i--
		dAtA[i] = 0x10
	}
	if m.Result != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Compare_Version) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Compare_Version) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintRpc(dAtA, i, uint64(m.Version))
	i--
	dAtA[i] = 0x20
	return len(dAtA) - i, nil
}

func (m *Compare_CreateRevision) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Compare_CreateRevision) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintRpc(dAtA, i, uint64(m.CreateRevision))
	i--
	dAtA[i] = 0x28
	return len(dAtA) - i, nil
}

func (m *Compare_ModRevision) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Compare_ModRevision) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintRpc(dAtA, i, uint64(m.ModRevision))
	i--
	dAtA[i] = 0x30
	return len(dAtA) - i, nil
}

func (m *Compare_Value) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Compare_Value) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Value != nil {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}

func (m *Compare_Lease) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Compare_Lease) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintRpc(dAtA, i, uint64(m.Lease))
	i--
	dAtA[i] = 0x40
	return len(dAtA) - i, nil
}

func (m *TxnRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxnRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxnRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Failure) > 0 {
		for iNdEx := len(m.Failure) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Failure[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Success) > 0 {
		for iNdEx := len(m.Success) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Success[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Compare) > 0 {
		for iNdEx := len(m.Compare) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Compare[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TxnResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxnResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxnResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Responses) > 0 {
		for iNdEx := len(m.Responses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Responses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Succeeded {
		i--
		if m.Succeeded {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CompactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Physical {
		i--
		if m.Physical {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Revision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Revision))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CompactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HashRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HashRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *HashKVRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashKVRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HashKVRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Revision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Revision))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HashKVResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashKVResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HashKVResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CompactRevision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.CompactRevision))
		i--
		dAtA[i] = 0x18
	}
	if m.Hash != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Hash))
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HashResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HashResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Hash != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Hash))
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SnapshotRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *SnapshotResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Blob) > 0 {
		i -= len(m.Blob)
		copy(dAtA[i:], m.Blob)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Blob)))
		i--
		dAtA[i] = 0x1a
	}
	if m.RemainingBytes != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.RemainingBytes))
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WatchRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WatchRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RequestUnion != nil {
		{
			size := m.RequestUnion.Size()
			i -= size
			if _, err := m.RequestUnion.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *WatchRequest_CreateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WatchRequest_CreateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CreateRequest != nil {
		{
			size, err := m.CreateRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WatchRequest_CancelRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WatchRequest_CancelRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CancelRequest != nil {
		{
			size, err := m.CancelRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *WatchRequest_ProgressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WatchRequest_ProgressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ProgressRequest != nil {
		{
			size, err := m.ProgressRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}

func (m *WatchCreateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchCreateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WatchCreateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Fragment {
		i--
		if m.Fragment {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.WatchId != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.WatchId))
		i--
		dAtA[i] = 0x38
	}
	if m.PrevKv {
		i--
		if m.PrevKv {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Filters) > 0 {
		dAtA22 := make([]byte, len(m.Filters)*10)
		var j21 int
		for _, num := range m.Filters {
			for num >= 1<<7 {
				dAtA22[j21] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j21++
			}
			dAtA22[j21] = uint8(num)
			j21++
		}
		i -= j21
		copy(dAtA[i:], dAtA22[:j21])
		i = encodeVarintRpc(dAtA, i, uint64(j21))
		i--
		dAtA[i] = 0x2a
	}
	if m.ProgressNotify {
		i--
		if m.ProgressNotify {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.StartRevision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.StartRevision))
		i--
		dAtA[i] = 0x18
	}
	if len(m.RangeEnd) > 0 {
		i -= len(m.RangeEnd)
		copy(dAtA[i:], m.RangeEnd)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RangeEnd)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WatchCancelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchCancelRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WatchCancelRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.WatchId != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.WatchId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *WatchProgressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchProgressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WatchProgressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *WatchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WatchResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if m.Fragment {
		i--
		if m.Fragment {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.CancelReason) > 0 {
		i -= len(m.CancelReason)
		copy(dAtA[i:], m.CancelReason)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.CancelReason)))
		i--
		dAtA[i] = 0x32
	}
	if m.CompactRevision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.CompactRevision))
		i--
		dAtA[i] = 0x28
	}
	if m.Canceled {
		i--
		if m.Canceled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Created {
		i--
		if m.Created {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.WatchId != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.WatchId))
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LeaseGrantRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseGrantRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseGrantRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if m.TTL != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.TTL))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LeaseGrantResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseGrantResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseGrantResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x22
	}
	if m.TTL != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.TTL))
		i--
		dAtA[i] = 0x18
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LeaseRevokeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseRevokeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseRevokeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LeaseRevokeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseRevokeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseRevokeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LeaseCheckpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseCheckpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseCheckpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Remaining_TTL != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Remaining_TTL))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LeaseCheckpointRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseCheckpointRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseCheckpointRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Checkpoints) > 0 {
		for iNdEx := len(m.Checkpoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Checkpoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LeaseCheckpointResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseCheckpointResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseCheckpointResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LeaseKeepAliveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseKeepAliveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseKeepAliveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LeaseKeepAliveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseKeepAliveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseKeepAliveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TTL != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.TTL))
		i--
		dAtA[i] = 0x18
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LeaseTimeToLiveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseTimeToLiveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseTimeToLiveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Keys {
		i--
		if m.Keys {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LeaseTimeToLiveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseTimeToLiveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseTimeToLiveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Keys[iNdEx])
			copy(dAtA[i:], m.Keys[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.Keys[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.GrantedTTL != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.GrantedTTL))
		i--
		dAtA[i] = 0x20
	}
	if m.TTL != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.TTL))
		i--
		dAtA[i] = 0x18
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LeaseLeasesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseLeasesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseLeasesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *LeaseStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LeaseLeasesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseLeasesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseLeasesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Leases) > 0 {
		for iNdEx := len(m.Leases) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Leases[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Member) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Member) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Member) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IsLearner {
		i--
		if m.IsLearner {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.ClientURLs) > 0 {
		for iNdEx := len(m.ClientURLs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ClientURLs[iNdEx])
			copy(dAtA[i:], m.ClientURLs[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.ClientURLs[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PeerURLs) > 0 {
		for iNdEx := len(m.PeerURLs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PeerURLs[iNdEx])
			copy(dAtA[i:], m.PeerURLs[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerURLs[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MemberAddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberAddRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberAddRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IsLearner {
		i--
		if m.IsLearner {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.PeerURLs) > 0 {
		for iNdEx := len(m.PeerURLs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PeerURLs[iNdEx])
			copy(dAtA[i:], m.PeerURLs[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerURLs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MemberAddResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberAddResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberAddResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Member != nil {
		{
			size, err := m.Member.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MemberRemoveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberRemoveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberRemoveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MemberRemoveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberRemoveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberRemoveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MemberUpdateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberUpdateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberUpdateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PeerURLs) > 0 {
		for iNdEx := len(m.PeerURLs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PeerURLs[iNdEx])
			copy(dAtA[i:], m.PeerURLs[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerURLs[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MemberUpdateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberUpdateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberUpdateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MemberListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Linearizable {
		i--
		if m.Linearizable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MemberListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MemberPromoteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberPromoteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberPromoteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MemberPromoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberPromoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberPromoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DefragmentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DefragmentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefragmentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *DefragmentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DefragmentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefragmentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MoveLeaderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MoveLeaderRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MoveLeaderRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TargetID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.TargetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MoveLeaderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MoveLeaderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MoveLeaderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AlarmRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AlarmRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AlarmRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Alarm != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Alarm))
		i--
		dAtA[i] = 0x18
	}
	if m.MemberID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MemberID))
		i--
		dAtA[i] = 0x10
	}
	if m.Action != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AlarmMember) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AlarmMember) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AlarmMember) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Alarm != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Alarm))
		i--
		dAtA[i] = 0x10
	}
	if m.MemberID != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.MemberID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AlarmResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AlarmResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AlarmResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Alarms) > 0 {
		for iNdEx := len(m.Alarms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Alarms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DowngradeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DowngradeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DowngradeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if m.Action != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DowngradeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DowngradeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DowngradeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *StatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IsLearner {
		i--
		if m.IsLearner {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.DbSizeInUse != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.DbSizeInUse))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Errors) > 0 {
		for iNdEx := len(m.Errors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Errors[iNdEx])
			copy(dAtA[i:], m.Errors[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.Errors[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if m.RaftAppliedIndex != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.RaftAppliedIndex))
		i--
		dAtA[i] = 0x38
	}
	if m.RaftTerm != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.RaftTerm))
		i--
		dAtA[i] = 0x30
	}
	if m.RaftIndex != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.RaftIndex))
		i--
		dAtA[i] = 0x28
	}
	if m.Leader != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.Leader))
		i--
		dAtA[i] = 0x20
	}
	if m.DbSize != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.DbSize))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthEnableRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthEnableRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthEnableRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *AuthDisableRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthDisableRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthDisableRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *AuthStatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthStatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthStatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *AuthenticateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthenticateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserAddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserAddRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserAddRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.HashedPassword) > 0 {
		i -= len(m.HashedPassword)
		copy(dAtA[i:], m.HashedPassword)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.HashedPassword)))
		i--
		dAtA[i] = 0x22
	}
	if m.Options != nil {
		{
			size, err := m.Options.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserGetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserGetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserGetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserDeleteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserDeleteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserDeleteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserChangePasswordRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserChangePasswordRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserChangePasswordRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.HashedPassword) > 0 {
		i -= len(m.HashedPassword)
		copy(dAtA[i:], m.HashedPassword)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.HashedPassword)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserGrantRoleRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserGrantRoleRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserGrantRoleRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserRevokeRoleRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserRevokeRoleRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserRevokeRoleRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthRoleAddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleAddRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleAddRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthRoleGetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleGetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleGetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *AuthRoleListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *AuthRoleDeleteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleDeleteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleDeleteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthRoleGrantPermissionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleGrantPermissionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleGrantPermissionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Perm != nil {
		{
			size, err := m.Perm.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthRoleRevokePermissionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleRevokePermissionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleRevokePermissionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RangeEnd) > 0 {
		i -= len(m.RangeEnd)
		copy(dAtA[i:], m.RangeEnd)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RangeEnd)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthEnableResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthEnableResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthEnableResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthDisableResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthDisableResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthDisableResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.AuthRevision != 0 {
		i = encodeVarintRpc(dAtA, i, uint64(m.AuthRevision))
		i--
		dAtA[i] = 0x18
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthenticateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthenticateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserAddResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserAddResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserAddResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserGetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserGetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserGetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Roles) > 0 {
		for iNdEx := len(m.Roles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Roles[iNdEx])
			copy(dAtA[i:], m.Roles[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.Roles[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserDeleteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserDeleteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserDeleteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserChangePasswordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserChangePasswordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserChangePasswordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserGrantRoleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserGrantRoleResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserGrantRoleResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserRevokeRoleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserRevokeRoleResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserRevokeRoleResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthRoleAddResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleAddResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleAddResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthRoleGetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleGetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleGetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Perm) > 0 {
		for iNdEx := len(m.Perm) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Perm[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthRoleListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Roles) > 0 {
		for iNdEx := len(m.Roles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Roles[iNdEx])
			copy(dAtA[i:], m.Roles[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.Roles[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthUserListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthUserListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthUserListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Users) > 0 {
		for iNdEx := len(m.Users) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Users[iNdEx])
			copy(dAtA[i:], m.Users[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.Users[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthRoleDeleteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleDeleteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleDeleteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthRoleGrantPermissionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleGrantPermissionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleGrantPermissionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthRoleRevokePermissionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthRoleRevokePermissionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthRoleRevokePermissionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovRpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *ResponseHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClusterId != 0 {
		n += 1 + sovRpc(uint64(m.ClusterId))
	}
	if m.MemberId != 0 {
		n += 1 + sovRpc(uint64(m.MemberId))
	}
	if m.Revision != 0 {
		n += 1 + sovRpc(uint64(m.Revision))
	}
	if m.RaftTerm != 0 {
		n += 1 + sovRpc(uint64(m.RaftTerm))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RangeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.RangeEnd)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Limit != 0 {
		n += 1 + sovRpc(uint64(m.Limit))
	}
	if m.Revision != 0 {
		n += 1 + sovRpc(uint64(m.Revision))
	}
	if m.SortOrder != 0 {
		n += 1 + sovRpc(uint64(m.SortOrder))
	}
	if m.SortTarget != 0 {
		n += 1 + sovRpc(uint64(m.SortTarget))
	}
	if m.Serializable {
		n += 2
	}
	if m.KeysOnly {
		n += 2
	}
	if m.CountOnly {
		n += 2
	}
	if m.MinModRevision != 0 {
		n += 1 + sovRpc(uint64(m.MinModRevision))
	}
	if m.MaxModRevision != 0 {
		n += 1 + sovRpc(uint64(m.MaxModRevision))
	}
	if m.MinCreateRevision != 0 {
		n += 1 + sovRpc(uint64(m.MinCreateRevision))
	}
	if m.MaxCreateRevision != 0 {
		n += 1 + sovRpc(uint64(m.MaxCreateRevision))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RangeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Kvs) > 0 {
		for _, e := range m.Kvs {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.More {
		n += 2
	}
	if m.Count != 0 {
		n += 1 + sovRpc(uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PutRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Lease != 0 {
		n += 1 + sovRpc(uint64(m.Lease))
	}
	if m.PrevKv {
		n += 2
	}
	if m.IgnoreValue {
		n += 2
	}
	if m.IgnoreLease {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PutResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.PrevKv != nil {
		l = m.PrevKv.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeleteRangeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.RangeEnd)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.PrevKv {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeleteRangeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Deleted != 0 {
		n += 1 + sovRpc(uint64(m.Deleted))
	}
	if len(m.PrevKvs) > 0 {
		for _, e := range m.PrevKvs {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestOp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Request != nil {
		n += m.Request.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RequestOp_RequestRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestRange != nil {
		l = m.RequestRange.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *RequestOp_RequestPut) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestPut != nil {
		l = m.RequestPut.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *RequestOp_RequestDeleteRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestDeleteRange != nil {
		l = m.RequestDeleteRange.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *RequestOp_RequestTxn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestTxn != nil {
		l = m.RequestTxn.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *ResponseOp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Response != nil {
		n += m.Response.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseOp_ResponseRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResponseRange != nil {
		l = m.ResponseRange.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *ResponseOp_ResponsePut) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResponsePut != nil {
		l = m.ResponsePut.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *ResponseOp_ResponseDeleteRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResponseDeleteRange != nil {
		l = m.ResponseDeleteRange.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *ResponseOp_ResponseTxn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResponseTxn != nil {
		l = m.ResponseTxn.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *Compare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovRpc(uint64(m.Result))
	}
	if m.Target != 0 {
		n += 1 + sovRpc(uint64(m.Target))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.TargetUnion != nil {
		n += m.TargetUnion.Size()
	}
	l = len(m.RangeEnd)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Compare_Version) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovRpc(uint64(m.Version))
	return n
}

func (m *Compare_CreateRevision) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovRpc(uint64(m.CreateRevision))
	return n
}

func (m *Compare_ModRevision) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovRpc(uint64(m.ModRevision))
	return n
}

func (m *Compare_Value) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		l = len(m.Value)
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *Compare_Lease) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovRpc(uint64(m.Lease))
	return n
}

func (m *TxnRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Compare) > 0 {
		for _, e := range m.Compare {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if len(m.Success) > 0 {
		for _, e := range m.Success {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if len(m.Failure) > 0 {
		for _, e := range m.Failure {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TxnResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Succeeded {
		n += 2
	}
	if len(m.Responses) > 0 {
		for _, e := range m.Responses {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CompactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Revision != 0 {
		n += 1 + sovRpc(uint64(m.Revision))
	}
	if m.Physical {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CompactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HashRequest) Size() (n int) {
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

func (m *HashKVRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Revision != 0 {
		n += 1 + sovRpc(uint64(m.Revision))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HashKVResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Hash != 0 {
		n += 1 + sovRpc(uint64(m.Hash))
	}
	if m.CompactRevision != 0 {
		n += 1 + sovRpc(uint64(m.CompactRevision))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HashResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Hash != 0 {
		n += 1 + sovRpc(uint64(m.Hash))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SnapshotRequest) Size() (n int) {
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

func (m *SnapshotResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.RemainingBytes != 0 {
		n += 1 + sovRpc(uint64(m.RemainingBytes))
	}
	l = len(m.Blob)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WatchRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestUnion != nil {
		n += m.RequestUnion.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WatchRequest_CreateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreateRequest != nil {
		l = m.CreateRequest.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *WatchRequest_CancelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CancelRequest != nil {
		l = m.CancelRequest.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *WatchRequest_ProgressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProgressRequest != nil {
		l = m.ProgressRequest.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *WatchCreateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.RangeEnd)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.StartRevision != 0 {
		n += 1 + sovRpc(uint64(m.StartRevision))
	}
	if m.ProgressNotify {
		n += 2
	}
	if len(m.Filters) > 0 {
		l = 0
		for _, e := range m.Filters {
			l += sovRpc(uint64(e))
		}
		n += 1 + sovRpc(uint64(l)) + l
	}
	if m.PrevKv {
		n += 2
	}
	if m.WatchId != 0 {
		n += 1 + sovRpc(uint64(m.WatchId))
	}
	if m.Fragment {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WatchCancelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.WatchId != 0 {
		n += 1 + sovRpc(uint64(m.WatchId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WatchProgressRequest) Size() (n int) {
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

func (m *WatchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.WatchId != 0 {
		n += 1 + sovRpc(uint64(m.WatchId))
	}
	if m.Created {
		n += 2
	}
	if m.Canceled {
		n += 2
	}
	if m.CompactRevision != 0 {
		n += 1 + sovRpc(uint64(m.CompactRevision))
	}
	l = len(m.CancelReason)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Fragment {
		n += 2
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseGrantRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TTL != 0 {
		n += 1 + sovRpc(uint64(m.TTL))
	}
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseGrantResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.TTL != 0 {
		n += 1 + sovRpc(uint64(m.TTL))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseRevokeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseRevokeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseCheckpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.Remaining_TTL != 0 {
		n += 1 + sovRpc(uint64(m.Remaining_TTL))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseCheckpointRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Checkpoints) > 0 {
		for _, e := range m.Checkpoints {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseCheckpointResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseKeepAliveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseKeepAliveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.TTL != 0 {
		n += 1 + sovRpc(uint64(m.TTL))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseTimeToLiveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.Keys {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseTimeToLiveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.TTL != 0 {
		n += 1 + sovRpc(uint64(m.TTL))
	}
	if m.GrantedTTL != 0 {
		n += 1 + sovRpc(uint64(m.GrantedTTL))
	}
	if len(m.Keys) > 0 {
		for _, b := range m.Keys {
			l = len(b)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseLeasesRequest) Size() (n int) {
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

func (m *LeaseStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseLeasesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Leases) > 0 {
		for _, e := range m.Leases {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Member) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.PeerURLs) > 0 {
		for _, s := range m.PeerURLs {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if len(m.ClientURLs) > 0 {
		for _, s := range m.ClientURLs {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.IsLearner {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MemberAddRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PeerURLs) > 0 {
		for _, s := range m.PeerURLs {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.IsLearner {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MemberAddResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Member != nil {
		l = m.Member.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MemberRemoveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MemberRemoveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MemberUpdateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if len(m.PeerURLs) > 0 {
		for _, s := range m.PeerURLs {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MemberUpdateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MemberListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Linearizable {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MemberListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MemberPromoteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MemberPromoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DefragmentRequest) Size() (n int) {
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

func (m *DefragmentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MoveLeaderRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TargetID != 0 {
		n += 1 + sovRpc(uint64(m.TargetID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MoveLeaderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AlarmRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Action != 0 {
		n += 1 + sovRpc(uint64(m.Action))
	}
	if m.MemberID != 0 {
		n += 1 + sovRpc(uint64(m.MemberID))
	}
	if m.Alarm != 0 {
		n += 1 + sovRpc(uint64(m.Alarm))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AlarmMember) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MemberID != 0 {
		n += 1 + sovRpc(uint64(m.MemberID))
	}
	if m.Alarm != 0 {
		n += 1 + sovRpc(uint64(m.Alarm))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AlarmResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Alarms) > 0 {
		for _, e := range m.Alarms {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DowngradeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Action != 0 {
		n += 1 + sovRpc(uint64(m.Action))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DowngradeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StatusRequest) Size() (n int) {
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

func (m *StatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.DbSize != 0 {
		n += 1 + sovRpc(uint64(m.DbSize))
	}
	if m.Leader != 0 {
		n += 1 + sovRpc(uint64(m.Leader))
	}
	if m.RaftIndex != 0 {
		n += 1 + sovRpc(uint64(m.RaftIndex))
	}
	if m.RaftTerm != 0 {
		n += 1 + sovRpc(uint64(m.RaftTerm))
	}
	if m.RaftAppliedIndex != 0 {
		n += 1 + sovRpc(uint64(m.RaftAppliedIndex))
	}
	if len(m.Errors) > 0 {
		for _, s := range m.Errors {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.DbSizeInUse != 0 {
		n += 1 + sovRpc(uint64(m.DbSizeInUse))
	}
	if m.IsLearner {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthEnableRequest) Size() (n int) {
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

func (m *AuthDisableRequest) Size() (n int) {
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

func (m *AuthStatusRequest) Size() (n int) {
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

func (m *AuthenticateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserAddRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Options != nil {
		l = m.Options.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.HashedPassword)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserGetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserDeleteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserChangePasswordRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.HashedPassword)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserGrantRoleRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserRevokeRoleRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthRoleAddRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthRoleGetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserListRequest) Size() (n int) {
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

func (m *AuthRoleListRequest) Size() (n int) {
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

func (m *AuthRoleDeleteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthRoleGrantPermissionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Perm != nil {
		l = m.Perm.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthRoleRevokePermissionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.RangeEnd)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthEnableResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthDisableResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if m.AuthRevision != 0 {
		n += 1 + sovRpc(uint64(m.AuthRevision))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthenticateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserAddResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserGetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserDeleteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserChangePasswordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserGrantRoleResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserRevokeRoleResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthRoleAddResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthRoleGetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Perm) > 0 {
		for _, e := range m.Perm {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthRoleListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthUserListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Users) > 0 {
		for _, s := range m.Users {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthRoleDeleteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthRoleGrantPermissionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthRoleRevokePermissionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozRpc(x uint64) (n int) {
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *ResponseHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: ResponseHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			m.ClusterId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClusterId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberId", wireType)
			}
			m.MemberId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemberId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftTerm", wireType)
			}
			m.RaftTerm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftTerm |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *RangeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: RangeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeEnd = append(m.RangeEnd[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeEnd == nil {
				m.RangeEnd = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortOrder", wireType)
			}
			m.SortOrder = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SortOrder |= RangeRequest_SortOrder(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortTarget", wireType)
			}
			m.SortTarget = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SortTarget |= RangeRequest_SortTarget(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Serializable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Serializable = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeysOnly", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.KeysOnly = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountOnly", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CountOnly = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinModRevision", wireType)
			}
			m.MinModRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinModRevision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxModRevision", wireType)
			}
			m.MaxModRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxModRevision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCreateRevision", wireType)
			}
			m.MinCreateRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinCreateRevision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCreateRevision", wireType)
			}
			m.MaxCreateRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxCreateRevision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *RangeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: RangeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kvs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kvs = append(m.Kvs, &mvccpb.KeyValue{})
			if err := m.Kvs[len(m.Kvs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field More", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.More = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *PutRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: PutRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			m.Lease = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lease |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKv", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PrevKv = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoreValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IgnoreValue = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoreLease", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IgnoreLease = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *PutResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: PutResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrevKv == nil {
				m.PrevKv = &mvccpb.KeyValue{}
			}
			if err := m.PrevKv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *DeleteRangeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: DeleteRangeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteRangeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeEnd = append(m.RangeEnd[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeEnd == nil {
				m.RangeEnd = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKv", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PrevKv = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *DeleteRangeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: DeleteRangeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteRangeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			m.Deleted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Deleted |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKvs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevKvs = append(m.PrevKvs, &mvccpb.KeyValue{})
			if err := m.PrevKvs[len(m.PrevKvs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *RequestOp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: RequestOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RangeRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Request = &RequestOp_RequestRange{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestPut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PutRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Request = &RequestOp_RequestPut{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestDeleteRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DeleteRangeRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Request = &RequestOp_RequestDeleteRange{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTxn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TxnRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Request = &RequestOp_RequestTxn{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *ResponseOp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: ResponseOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RangeResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Response = &ResponseOp_ResponseRange{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponsePut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PutResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Response = &ResponseOp_ResponsePut{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseDeleteRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DeleteRangeResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Response = &ResponseOp_ResponseDeleteRange{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseTxn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TxnResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Response = &ResponseOp_ResponseTxn{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *Compare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: Compare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Compare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= Compare_CompareResult(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			m.Target = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Target |= Compare_CompareTarget(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TargetUnion = &Compare_Version{v}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateRevision", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TargetUnion = &Compare_CreateRevision{v}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModRevision", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TargetUnion = &Compare_ModRevision{v}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.TargetUnion = &Compare_Value{v}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TargetUnion = &Compare_Lease{v}
		case 64:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeEnd = append(m.RangeEnd[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeEnd == nil {
				m.RangeEnd = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *TxnRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: TxnRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxnRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compare", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Compare = append(m.Compare, &Compare{})
			if err := m.Compare[len(m.Compare)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Success = append(m.Success, &RequestOp{})
			if err := m.Success[len(m.Success)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failure", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Failure = append(m.Failure, &RequestOp{})
			if err := m.Failure[len(m.Failure)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *TxnResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: TxnResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxnResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Succeeded", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Succeeded = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Responses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Responses = append(m.Responses, &ResponseOp{})
			if err := m.Responses[len(m.Responses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *CompactionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: CompactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Physical", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Physical = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *CompactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: CompactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *HashRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: HashRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *HashKVRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: HashKVRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashKVRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *HashKVResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: HashKVResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashKVResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			m.Hash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hash |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompactRevision", wireType)
			}
			m.CompactRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompactRevision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *HashResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: HashResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			m.Hash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hash |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *SnapshotRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: SnapshotRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *SnapshotResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: SnapshotResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingBytes", wireType)
			}
			m.RemainingBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemainingBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blob = append(m.Blob[:0], dAtA[iNdEx:postIndex]...)
			if m.Blob == nil {
				m.Blob = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *WatchRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: WatchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &WatchCreateRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.RequestUnion = &WatchRequest_CreateRequest{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CancelRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &WatchCancelRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.RequestUnion = &WatchRequest_CancelRequest{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgressRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &WatchProgressRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.RequestUnion = &WatchRequest_ProgressRequest{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *WatchCreateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: WatchCreateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchCreateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeEnd = append(m.RangeEnd[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeEnd == nil {
				m.RangeEnd = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartRevision", wireType)
			}
			m.StartRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartRevision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgressNotify", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ProgressNotify = bool(v != 0)
		case 5:
			if wireType == 0 {
				var v WatchCreateRequest_FilterType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRpc
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= WatchCreateRequest_FilterType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Filters = append(m.Filters, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRpc
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthRpc
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthRpc
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Filters) == 0 {
					m.Filters = make([]WatchCreateRequest_FilterType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v WatchCreateRequest_FilterType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRpc
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= WatchCreateRequest_FilterType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Filters = append(m.Filters, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKv", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PrevKv = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WatchId", wireType)
			}
			m.WatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WatchId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fragment", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Fragment = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *WatchCancelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: WatchCancelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchCancelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WatchId", wireType)
			}
			m.WatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WatchId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *WatchProgressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: WatchProgressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchProgressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *WatchResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: WatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WatchId", wireType)
			}
			m.WatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WatchId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Created = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Canceled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Canceled = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompactRevision", wireType)
			}
			m.CompactRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompactRevision |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CancelReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CancelReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fragment", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Fragment = bool(v != 0)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &mvccpb.Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseGrantRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseGrantRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseGrantRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseGrantResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseGrantResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseGrantResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseRevokeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseRevokeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseRevokeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseRevokeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseRevokeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseRevokeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseCheckpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseCheckpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseCheckpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Remaining_TTL", wireType)
			}
			m.Remaining_TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Remaining_TTL |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseCheckpointRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseCheckpointRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseCheckpointRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checkpoints = append(m.Checkpoints, &LeaseCheckpoint{})
			if err := m.Checkpoints[len(m.Checkpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseCheckpointResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseCheckpointResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseCheckpointResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseKeepAliveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseKeepAliveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseKeepAliveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseKeepAliveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseKeepAliveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseKeepAliveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseTimeToLiveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseTimeToLiveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseTimeToLiveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Keys = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseTimeToLiveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseTimeToLiveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseTimeToLiveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrantedTTL", wireType)
			}
			m.GrantedTTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GrantedTTL |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, make([]byte, postIndex-iNdEx))
			copy(m.Keys[len(m.Keys)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseLeasesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseLeasesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseLeasesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *LeaseLeasesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseLeasesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseLeasesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leases", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Leases = append(m.Leases, &LeaseStatus{})
			if err := m.Leases[len(m.Leases)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *Member) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: Member: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Member: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerURLs = append(m.PeerURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientURLs = append(m.ClientURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLearner", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsLearner = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MemberAddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerURLs = append(m.PeerURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLearner", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsLearner = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MemberAddResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberAddResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberAddResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Member == nil {
				m.Member = &Member{}
			}
			if err := m.Member.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MemberRemoveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberRemoveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberRemoveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MemberRemoveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberRemoveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberRemoveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MemberUpdateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberUpdateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberUpdateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return fmt.Errorf("proto: wrong wireType = %d for field PeerURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerURLs = append(m.PeerURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MemberUpdateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberUpdateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberUpdateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MemberListRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Linearizable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Linearizable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MemberListResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MemberPromoteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberPromoteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberPromoteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MemberPromoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberPromoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberPromoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *DefragmentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: DefragmentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefragmentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *DefragmentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: DefragmentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefragmentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MoveLeaderRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MoveLeaderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MoveLeaderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetID", wireType)
			}
			m.TargetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *MoveLeaderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MoveLeaderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MoveLeaderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AlarmRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AlarmRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AlarmRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= AlarmRequest_AlarmAction(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberID", wireType)
			}
			m.MemberID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemberID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alarm", wireType)
			}
			m.Alarm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Alarm |= AlarmType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AlarmMember) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AlarmMember: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AlarmMember: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberID", wireType)
			}
			m.MemberID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemberID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alarm", wireType)
			}
			m.Alarm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Alarm |= AlarmType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AlarmResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AlarmResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AlarmResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alarms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alarms = append(m.Alarms, &AlarmMember{})
			if err := m.Alarms[len(m.Alarms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *DowngradeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: DowngradeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DowngradeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= DowngradeRequest_DowngradeAction(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *DowngradeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: DowngradeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DowngradeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *StatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: StatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *StatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: StatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DbSize", wireType)
			}
			m.DbSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DbSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leader", wireType)
			}
			m.Leader = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Leader |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftIndex", wireType)
			}
			m.RaftIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftTerm", wireType)
			}
			m.RaftTerm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftTerm |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftAppliedIndex", wireType)
			}
			m.RaftAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftAppliedIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Errors = append(m.Errors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DbSizeInUse", wireType)
			}
			m.DbSizeInUse = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DbSizeInUse |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLearner", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsLearner = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthEnableRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthEnableRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthEnableRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthDisableRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthDisableRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthDisableRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthStatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthStatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthStatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthenticateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthenticateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
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
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserAddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
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
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Options == nil {
				m.Options = &authpb.UserAddOptions{}
			}
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashedPassword", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashedPassword = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserGetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserGetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserGetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserDeleteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserDeleteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserDeleteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserChangePasswordRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserChangePasswordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserChangePasswordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
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
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashedPassword", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashedPassword = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserGrantRoleRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserGrantRoleRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserGrantRoleRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserRevokeRoleRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserRevokeRoleRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserRevokeRoleRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleAddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleGetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleGetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleGetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserListRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleListRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleDeleteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleDeleteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleDeleteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleGrantPermissionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleGrantPermissionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleGrantPermissionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Perm", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Perm == nil {
				m.Perm = &authpb.Permission{}
			}
			if err := m.Perm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleRevokePermissionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleRevokePermissionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleRevokePermissionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeEnd = append(m.RangeEnd[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeEnd == nil {
				m.RangeEnd = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthEnableResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthEnableResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthEnableResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthDisableResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthDisableResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthDisableResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRevision", wireType)
			}
			m.AuthRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthenticateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthenticateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserAddResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserAddResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserAddResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserGetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserDeleteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserDeleteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserDeleteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserChangePasswordResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserChangePasswordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserChangePasswordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserGrantRoleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserGrantRoleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserGrantRoleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserRevokeRoleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserRevokeRoleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserRevokeRoleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleAddResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleAddResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleAddResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleGetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Perm", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Perm = append(m.Perm, &authpb.Permission{})
			if err := m.Perm[len(m.Perm)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleListResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthUserListResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleDeleteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleDeleteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleDeleteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleGrantPermissionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleGrantPermissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleGrantPermissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func (m *AuthRoleRevokePermissionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleRevokePermissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleRevokePermissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
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

func skipRpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
				return 0, ErrInvalidLengthRpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRpc = fmt.Errorf("proto: unexpected end of group")
)
