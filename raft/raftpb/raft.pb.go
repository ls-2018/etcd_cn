// Code generated by protoc-gen-gogo.
// source: raft.proto

package raftpb

import (
	"encoding/json"
	fmt "fmt"
	math "math"

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

type EntryType int32

const (
	EntryNormal       EntryType = 0
	EntryConfChange   EntryType = 1
	EntryConfChangeV2 EntryType = 2
)

var EntryType_name = map[int32]string{
	0: "EntryNormal",
	1: "EntryConfChange",
	2: "EntryConfChangeV2",
}

var EntryType_value = map[string]int32{
	"EntryNormal":       0,
	"EntryConfChange":   1,
	"EntryConfChangeV2": 2,
}

func (x EntryType) Enum() *EntryType {
	p := new(EntryType)
	*p = x
	return p
}

func (x EntryType) String() string {
	return proto.EnumName(EntryType_name, int32(x))
}

func (x *EntryType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EntryType_value, data, "EntryType")
	if err != nil {
		return err
	}
	*x = EntryType(value)
	return nil
}

func (EntryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{0}
}

// For description of different message types, see:
// https://pkg.go.dev/github.com/ls-2018/etcd_cn/raft#hdr-MessageType
type MessageType int32

const (
	MsgHup            MessageType = 0
	MsgBeat           MessageType = 1
	MsgProp           MessageType = 2
	MsgApp            MessageType = 3
	MsgAppResp        MessageType = 4
	MsgVote           MessageType = 5
	MsgVoteResp       MessageType = 6
	MsgSnap           MessageType = 7
	MsgHeartbeat      MessageType = 8
	MsgHeartbeatResp  MessageType = 9
	MsgUnreachable    MessageType = 10
	MsgSnapStatus     MessageType = 11
	MsgCheckQuorum    MessageType = 12
	MsgTransferLeader MessageType = 13
	MsgTimeoutNow     MessageType = 14
	MsgReadIndex      MessageType = 15
	MsgReadIndexResp  MessageType = 16
	MsgPreVote        MessageType = 17
	MsgPreVoteResp    MessageType = 18
)

var MessageType_name = map[int32]string{
	0:  "MsgHup",
	1:  "MsgBeat",
	2:  "MsgProp",
	3:  "MsgApp",
	4:  "MsgAppResp",
	5:  "MsgVote",
	6:  "MsgVoteResp",
	7:  "MsgSnap",
	8:  "MsgHeartbeat",
	9:  "MsgHeartbeatResp",
	10: "MsgUnreachable",
	11: "MsgSnapStatus",
	12: "MsgCheckQuorum",
	13: "MsgTransferLeader",
	14: "MsgTimeoutNow",
	15: "MsgReadIndex",
	16: "MsgReadIndexResp",
	17: "MsgPreVote",
	18: "MsgPreVoteResp",
}

var MessageType_value = map[string]int32{
	"MsgHup":            0,
	"MsgBeat":           1,
	"MsgProp":           2,
	"MsgApp":            3,
	"MsgAppResp":        4,
	"MsgVote":           5,
	"MsgVoteResp":       6,
	"MsgSnap":           7,
	"MsgHeartbeat":      8,
	"MsgHeartbeatResp":  9,
	"MsgUnreachable":    10,
	"MsgSnapStatus":     11,
	"MsgCheckQuorum":    12,
	"MsgTransferLeader": 13,
	"MsgTimeoutNow":     14,
	"MsgReadIndex":      15,
	"MsgReadIndexResp":  16,
	"MsgPreVote":        17,
	"MsgPreVoteResp":    18,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{1}
}

// ConfChangeTransition specifies the behavior of a configuration change with
// respect to joint consensus.
type ConfChangeTransition int32

// 指定关于联合共识的配置更改的行为.自动、显式、隐式
const (
	// Automatically use the simple protocol if possible, otherwise fall back
	// to ConfChangeJointImplicit. Most applications will want to use this.
	ConfChangeTransitionAuto ConfChangeTransition = 0
	// Use joint consensus unconditionally, and transition out of them
	// automatically (by proposing a zero configuration change).
	//
	// This option is suitable for applications that want to minimize the time
	// spent in the joint configuration and do not store the joint configuration
	// in the state machine (outside of InitialState).
	ConfChangeTransitionJointImplicit ConfChangeTransition = 1
	// Use joint consensus and remain in the joint configuration until the
	// application proposes a no-op configuration change. This is suitable for
	// applications that want to explicitly control the transitions, for example
	// to use a custom payload (via the Context field).
	ConfChangeTransitionJointExplicit ConfChangeTransition = 2
)

var ConfChangeTransition_name = map[int32]string{
	0: "ConfChangeTransitionAuto",
	1: "ConfChangeTransitionJointImplicit",
	2: "ConfChangeTransitionJointExplicit",
}

var ConfChangeTransition_value = map[string]int32{
	"ConfChangeTransitionAuto":          0,
	"ConfChangeTransitionJointImplicit": 1,
	"ConfChangeTransitionJointExplicit": 2,
}

func (x ConfChangeTransition) Enum() *ConfChangeTransition {
	p := new(ConfChangeTransition)
	*p = x
	return p
}

func (x ConfChangeTransition) String() string {
	return proto.EnumName(ConfChangeTransition_name, int32(x))
}

func (x *ConfChangeTransition) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ConfChangeTransition_value, data, "ConfChangeTransition")
	if err != nil {
		return err
	}
	*x = ConfChangeTransition(value)
	return nil
}

func (ConfChangeTransition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{2}
}

type ConfChangeType int32

const (
	ConfChangeAddNode        ConfChangeType = 0
	ConfChangeRemoveNode     ConfChangeType = 1
	ConfChangeUpdateNode     ConfChangeType = 2
	ConfChangeAddLearnerNode ConfChangeType = 3
)

var ConfChangeType_name = map[int32]string{
	0: "ConfChangeAddNode",
	1: "ConfChangeRemoveNode",
	2: "ConfChangeUpdateNode",
	3: "ConfChangeAddLearnerNode",
}

var ConfChangeType_value = map[string]int32{
	"ConfChangeAddNode":        0,
	"ConfChangeRemoveNode":     1,
	"ConfChangeUpdateNode":     2,
	"ConfChangeAddLearnerNode": 3,
}

func (x ConfChangeType) Enum() *ConfChangeType {
	p := new(ConfChangeType)
	*p = x
	return p
}

func (x ConfChangeType) String() string {
	return proto.EnumName(ConfChangeType_name, int32(x))
}

func (x *ConfChangeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ConfChangeType_value, data, "ConfChangeType")
	if err != nil {
		return err
	}
	*x = ConfChangeType(value)
	return nil
}

func (ConfChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{3}
}

type Entry struct {
	// Term：表示该Entry所在的任期.
	Term uint64 `protobuf:"varint,2,opt,name=Term" json:"Term"`
	// Index:当前这个entry在整个raft日志中的位置索引,有了Term和Index之后,一个`log entry`就能被唯一标识.
	Index uint64 `protobuf:"varint,3,opt,name=Index" json:"Index"`
	// 当前entry的类型
	// 目前etcd支持两种类型：EntryNormal和EntryConfChange
	// EntryNormaln表示普通的数据操作
	// EntryConfChange表示集群的变更操作
	Type EntryType `protobuf:"varint,1,opt,name=Type,enum=raftpb.EntryType" json:"Type"`
	// 具体操作使用的数据
	Data []byte `protobuf:"bytes,4,opt,name=Data" json:"Data,omitempty"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{0}
}

type SnapshotMetadata struct {
	ConfState ConfState `protobuf:"bytes,1,opt,name=conf_state,json=confState" json:"conf_state"`
	Index     uint64    `protobuf:"varint,2,opt,name=index" json:"index"`
	Term      uint64    `protobuf:"varint,3,opt,name=term" json:"term"`
}

func (m *SnapshotMetadata) Reset()         { *m = SnapshotMetadata{} }
func (m *SnapshotMetadata) String() string { return proto.CompactTextString(m) }
func (*SnapshotMetadata) ProtoMessage()    {}
func (*SnapshotMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{1}
}

type Snapshot struct {
	Data     []byte           `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Metadata SnapshotMetadata `protobuf:"bytes,2,opt,name=metadata" json:"metadata"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{2}
}

// Message 消息格式
// Raft协议时提到,节点之间传递的是消息(Message), 每条消息中可以携带多条Entry记录,每条Entry记录对应一个独立的操作.
type Message struct {
	// 该字段定义了不同的消息类型,etcd-raft就是通过不同的消息类型来进行处理的,etcd中一共定义了19种类型
	Type MessageType `protobuf:"varint,1,opt,name=type,enum=raftpb.MessageType" json:"type"`
	// 消息的目标节点 ID,在急群中每个节点都有一个唯一的id作为标识
	To uint64 `protobuf:"varint,2,opt,name=to" json:"to"`
	// 发送消息的节点ID.在集群中,每个节点都拥有一个唯一ID作为标识.,或者leader转移的下一个跳
	From uint64 `protobuf:"varint,3,opt,name=from" json:"from"`
	// 发送消息的节点的Term值. 如果Term值为0,则为本地消息,在etcd刊负模块的实现中,对本地消息进行特殊处理.
	Term uint64 `protobuf:"varint,4,opt,name=term" json:"term"`
	// 该消息携带的第一条Entry记录的Term值.
	LogTerm uint64 `protobuf:"varint,5,opt,name=logTerm" json:"logTerm"`
	Index   uint64 `protobuf:"varint,6,opt,name=index" json:"index"` // 日志索引ID，用于节点向leader汇报自己已经commit的日志数据ID
	// 如果是MsgApp类型的消息,则该字段中保存了Leader节点复制到Follower节点的Entry记录.在其他类型消息中,该字段的含义后面会详细介绍.
	Entries []Entry `protobuf:"bytes,7,rep,name=entries" json:"entries"`
	// 搜 ProgressTracker
	// handleAppendEntries 处理函数
	// leader会为每个Follower都维护一个leaderCommit,表示leader认为Follower已经提交的日志条目索引值
	Commit uint64 `protobuf:"varint,8,opt,name=commit" json:"commit"`
	// 在传输快照时,该字段保存了快照数据
	Snapshot Snapshot `protobuf:"bytes,9,opt,name=snapshot" json:"snapshot"`
	// 主要用于响应类型的消息,表示是否拒绝收到的消息.
	Reject     bool   `protobuf:"varint,10,opt,name=reject" json:"reject"`
	RejectHint uint64 `protobuf:"varint,11,opt,name=rejectHint" json:"rejectHint"` // 拒绝同步日志请求时返回的当前节点日志ID，用于被拒绝方快速定位到下一次合适的同步日志位置
	// 携带的一些上下文的信息, 例如,campaignTransfer
	Context []byte `protobuf:"bytes,12,opt,name=context" json:"context,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{3}
}

// HardState 封装了raft协议中规定的需要实时持久化的状态属性：当前选举周期、投票和已提交的Index
type HardState struct {
	Term   uint64 `protobuf:"varint,1,opt,name=term" json:"term"`     // 当前任期
	Vote   uint64 `protobuf:"varint,2,opt,name=vote" json:"vote"`     // 给谁投了票
	Commit uint64 `protobuf:"varint,3,opt,name=commit" json:"commit"` // 已提交的位置
}

// Reset 重置集群状态
func (m *HardState) Reset()         { *m = HardState{} }
func (m *HardState) String() string { return proto.CompactTextString(m) }
func (*HardState) ProtoMessage()    {}
func (*HardState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{4}
}

// ConfState tracker.Config的另一种体现形式
type ConfState struct {
	// The voters in the incoming config. (If the configuration is not joint,
	// then the outgoing config is empty).
	Voters []uint64 `protobuf:"varint,1,rep,name=voters" json:"voters,omitempty"`
	// The learners in the incoming config.
	Learners []uint64 `protobuf:"varint,2,rep,name=learners" json:"learners,omitempty"`
	// The voters in the outgoing config.
	VotersOutgoing []uint64 `protobuf:"varint,3,rep,name=voters_outgoing,json=votersOutgoing" json:"voters_outgoing,omitempty"`
	// The nodes that will become learners when the outgoing config is removed.
	// These nodes are necessarily currently in nodes_joint (or they would have
	// been added to the incoming config right away).
	LearnersNext []uint64 `protobuf:"varint,4,rep,name=learners_next,json=learnersNext" json:"learners_next,omitempty"`
	// If set, the config is joint and Raft will automatically transition into
	// the final config (i.e. remove the outgoing config) when this is safe.
	AutoLeave bool `protobuf:"varint,5,opt,name=auto_leave,json=autoLeave" json:"auto_leave"`
}

// ConfChangeV1 只能传递一个节点变更
type ConfChangeV1 struct {
	Type    ConfChangeType `protobuf:"varint,2,opt,name=type,enum=raftpb.ConfChangeType" json:"type"`
	NodeID  uint64         `protobuf:"varint,3,opt,name=node_id,json=nodeId" json:"node_id"` // NodeID: 变更节点的id
	Context string         `protobuf:"string,4,opt,name=context" json:"context,omitempty"`
	// 这个字段只被etcd用来传播一个唯一的标识符.理想情况下,它应该真正使用Context来代替.在ConfChangeV2中没有对应的字段存在.
	ID uint64 `protobuf:"varint,1,opt,name=id" json:"id"` // 节点变更的次数
}

// ConfChangeV2 可以传递多个
type ConfChangeV2 struct {
	Transition ConfChangeTransition `protobuf:"varint,1,opt,name=transition,enum=raftpb.ConfChangeTransition" json:"transition"`
	Changes    []ConfChangeSingle   `protobuf:"bytes,2,rep,name=changes" json:"changes"`
	Context    string               `protobuf:"bytes,3,opt,name=context" json:"context,omitempty"`
}

func (m *ConfState) Reset()         { *m = ConfState{} }
func (m *ConfState) String() string { return proto.CompactTextString(m) }
func (*ConfState) ProtoMessage()    {}
func (*ConfState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{5}
}

func (m *ConfChangeV1) Reset()         { *m = ConfChangeV1{} }
func (m *ConfChangeV1) String() string { return proto.CompactTextString(m) }
func (*ConfChangeV1) ProtoMessage()    {}
func (*ConfChangeV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{6}
}

// ConfChangeSingle is an individual configuration change operation. Multiple
// such operations can be carried out atomically via a ConfChangeV2.
type ConfChangeSingle struct {
	Type   ConfChangeType `protobuf:"varint,1,opt,name=type,enum=raftpb.ConfChangeType" json:"type"`
	NodeID uint64         `protobuf:"varint,2,opt,name=node_id,json=nodeId" json:"node_id"`
}

func (m *ConfChangeSingle) Reset()         { *m = ConfChangeSingle{} }
func (m *ConfChangeSingle) String() string { return proto.CompactTextString(m) }
func (*ConfChangeSingle) ProtoMessage()    {}
func (*ConfChangeSingle) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{7}
}

// ConfChangeV2 messages initiate configuration changes. They support both the
// simple "one at a time" membership change protocol and full Joint Consensus
// allowing for arbitrary changes in membership.
//
// The supplied context is treated as an opaque payload and can be used to
// attach an action on the state machine to the application of the config change
// proposal. Note that contrary to Joint Consensus as outlined in the Raft
// paper[1], configuration changes become active when they are *applied* to the
// state machine (not when they are appended to the log).
//
// The simple protocol can be used whenever only a single change is made.
//
// Non-simple changes require the use of Joint Consensus, for which two
// configuration changes are run. The first configuration change specifies the
// desired changes and transitions the Raft group into the joint configuration,
// in which quorum requires a majority of both the pre-changes and post-changes
// configuration. Joint Consensus avoids entering fragile intermediate
// configurations that could compromise survivability. For example, without the
// use of Joint Consensus and running across three availability zones with a
// replication factor of three, it is not possible to replace a voter without
// entering an intermediate configuration that does not survive the outage of
// one availability zone.
//
// The provided ConfChangeTransition specifies how (and whether) Joint Consensus
// is used, and assigns the task of leaving the joint configuration either to
// Raft or the application. Leaving the joint configuration is accomplished by
// proposing a ConfChangeV2 with only and optionally the Context field
// populated.
//
// For details on Raft membership changes, see:
//
// [1]: https://github.com/ongardie/dissertation/blob/master/online-trim.pdf

func (m *ConfChangeV2) Reset()         { *m = ConfChangeV2{} }
func (m *ConfChangeV2) String() string { return proto.CompactTextString(m) }
func (*ConfChangeV2) ProtoMessage()    {}
func (*ConfChangeV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{8}
}

func init() {
	proto.RegisterEnum("raftpb.EntryType", EntryType_name, EntryType_value)
	proto.RegisterEnum("raftpb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("raftpb.ConfChangeTransition", ConfChangeTransition_name, ConfChangeTransition_value)
	proto.RegisterEnum("raftpb.ConfChangeType", ConfChangeType_name, ConfChangeType_value)
	proto.RegisterType((*Entry)(nil), "raftpb.Entry")
	proto.RegisterType((*SnapshotMetadata)(nil), "raftpb.SnapshotMetadata")
	proto.RegisterType((*Snapshot)(nil), "raftpb.Snapshot")
	proto.RegisterType((*Message)(nil), "raftpb.Message")
	proto.RegisterType((*HardState)(nil), "raftpb.HardState")
	proto.RegisterType((*ConfState)(nil), "raftpb.ConfState")
	proto.RegisterType((*ConfChangeV1)(nil), "raftpb.ConfChange")
	proto.RegisterType((*ConfChangeSingle)(nil), "raftpb.ConfChangeSingle")
	proto.RegisterType((*ConfChangeV2)(nil), "raftpb.ConfChangeV2")
}

func init() { proto.RegisterFile("raft.proto", fileDescriptor_b042552c306ae59b) }

var fileDescriptor_b042552c306ae59b = []byte{
	// 1026 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x17, 0x25, 0x29, 0x5a, 0x3f, 0x57, 0xb2, 0x3c, 0xbe, 0xf1, 0x17, 0x10, 0x86, 0xc1, 0xe8, 0x53,
	0x52, 0x44, 0x70, 0x11, 0xb7, 0xd0, 0xa2, 0x28, 0xba, 0xf3, 0x4f, 0x00, 0xab, 0xb0, 0xdc, 0x54,
	0x76, 0xbc, 0x28, 0x50, 0x08, 0x63, 0x71, 0x44, 0xb3, 0x15, 0x39, 0x04, 0x39, 0x72, 0xed, 0x4d,
	0x51, 0xf4, 0x09, 0xba, 0xec, 0x26, 0xdb, 0x3e, 0x40, 0x9f, 0xc2, 0x4b, 0x03, 0xdd, 0x74, 0x15,
	0x34, 0xf6, 0x8b, 0x14, 0x33, 0x1c, 0x4a, 0x94, 0x6c, 0x64, 0xd1, 0xdd, 0xcc, 0xb9, 0x67, 0xee,
	0x9c, 0x73, 0xef, 0xe5, 0x10, 0x20, 0xa1, 0x63, 0xb1, 0x13, 0x27, 0x5c, 0x70, 0x2c, 0xcb, 0x75,
	0x7c, 0xbe, 0xb9, 0xe1, 0x73, 0x9f, 0x2b, 0xe8, 0x33, 0xb9, 0xca, 0xa2, 0xed, 0x9f, 0x61, 0xe5,
	0x75, 0x24, 0x92, 0x6b, 0x74, 0xc0, 0x3e, 0x65, 0x49, 0xe8, 0x58, 0x2d, 0xb3, 0x63, 0xef, 0xd9,
	0x37, 0xef, 0x9f, 0x19, 0x03, 0x85, 0xe0, 0x26, 0xac, 0xf4, 0x22, 0x8f, 0x5d, 0x39, 0xa5, 0x42,
	0x28, 0x83, 0xf0, 0x53, 0xb0, 0x4f, 0xaf, 0x63, 0xe6, 0x98, 0x2d, 0xb3, 0xd3, 0xec, 0xae, 0xef,
	0x64, 0x77, 0xed, 0xa8, 0x94, 0x32, 0x30, 0x4b, 0x74, 0x1d, 0x33, 0x44, 0xb0, 0x0f, 0xa8, 0xa0,
	0x8e, 0xdd, 0x32, 0x3b, 0x8d, 0x81, 0x5a, 0xb7, 0x7f, 0x31, 0x81, 0x9c, 0x44, 0x34, 0x4e, 0x2f,
	0xb8, 0xe8, 0x33, 0x41, 0x3d, 0x2a, 0x28, 0x7e, 0x01, 0x30, 0xe2, 0xd1, 0x78, 0x98, 0x0a, 0x2a,
	0xb2, 0xdc, 0xf5, 0x79, 0xee, 0x7d, 0x1e, 0x8d, 0x4f, 0x64, 0x40, 0xe7, 0xae, 0x8d, 0x72, 0x40,
	0x2a, 0x0d, 0x94, 0xd2, 0xa2, 0x89, 0x0c, 0x92, 0xfe, 0x84, 0xf4, 0x57, 0x34, 0xa1, 0x90, 0xf6,
	0x77, 0x50, 0xcd, 0x15, 0x48, 0x89, 0x52, 0x81, 0xba, 0xb3, 0x31, 0x50, 0x6b, 0xfc, 0x0a, 0xaa,
	0xa1, 0x56, 0xa6, 0x12, 0xd7, 0xbb, 0x4e, 0xae, 0x65, 0x59, 0xb9, 0xce, 0x3b, 0xe3, 0xb7, 0xdf,
	0x95, 0xa0, 0xd2, 0x67, 0x69, 0x4a, 0x7d, 0x86, 0xaf, 0xc0, 0x16, 0xf3, 0x5a, 0x3d, 0xc9, 0x73,
	0xe8, 0x70, 0xb1, 0x5a, 0x92, 0x86, 0x1b, 0x60, 0x09, 0xbe, 0xe0, 0xc4, 0x12, 0x5c, 0xda, 0x18,
	0x27, 0x7c, 0xc9, 0x86, 0x44, 0x66, 0x06, 0xed, 0x65, 0x83, 0xe8, 0x42, 0x65, 0xc2, 0x7d, 0xd5,
	0xdd, 0x95, 0x42, 0x30, 0x07, 0xe7, 0x65, 0x2b, 0x3f, 0x2c, 0xdb, 0x2b, 0xa8, 0xb0, 0x48, 0x24,
	0x01, 0x4b, 0x9d, 0x4a, 0xab, 0xd4, 0xa9, 0x77, 0x57, 0x17, 0x7a, 0x9c, 0xa7, 0xd2, 0x1c, 0xdc,
	0x82, 0xf2, 0x88, 0x87, 0x61, 0x20, 0x9c, 0x6a, 0x21, 0x97, 0xc6, 0xb0, 0x0b, 0xd5, 0x54, 0x57,
	0xcc, 0xa9, 0xa9, 0x4a, 0x92, 0xe5, 0x4a, 0xe6, 0x15, 0xcc, 0x79, 0x32, 0x63, 0xc2, 0x7e, 0x60,
	0x23, 0xe1, 0x40, 0xcb, 0xec, 0x54, 0xf3, 0x8c, 0x19, 0x86, 0x2f, 0x00, 0xb2, 0xd5, 0x61, 0x10,
	0x09, 0xa7, 0x5e, 0xb8, 0xb3, 0x80, 0xa3, 0x03, 0x95, 0x11, 0x8f, 0x04, 0xbb, 0x12, 0x4e, 0x43,
	0x35, 0x36, 0xdf, 0xb6, 0xbf, 0x87, 0xda, 0x21, 0x4d, 0xbc, 0x6c, 0x7c, 0xf2, 0x0a, 0x9a, 0x0f,
	0x2a, 0xe8, 0x80, 0x7d, 0xc9, 0x05, 0x5b, 0xfc, 0x38, 0x24, 0x52, 0x30, 0x5c, 0x7a, 0x68, 0xb8,
	0xfd, 0xa7, 0x09, 0xb5, 0xd9, 0xbc, 0xe2, 0x53, 0x28, 0xcb, 0x33, 0x49, 0xea, 0x98, 0xad, 0x52,
	0xc7, 0x1e, 0xe8, 0x1d, 0x6e, 0x42, 0x75, 0xc2, 0x68, 0x12, 0xc9, 0x88, 0xa5, 0x22, 0xb3, 0x3d,
	0xbe, 0x84, 0xb5, 0x8c, 0x35, 0xe4, 0x53, 0xe1, 0xf3, 0x20, 0xf2, 0x9d, 0x92, 0xa2, 0x34, 0x33,
	0xf8, 0x1b, 0x8d, 0xe2, 0x73, 0x58, 0xcd, 0x0f, 0x0d, 0x23, 0xe9, 0xd4, 0x56, 0xb4, 0x46, 0x0e,
	0x1e, 0xb3, 0x2b, 0x81, 0xcf, 0x01, 0xe8, 0x54, 0xf0, 0xe1, 0x84, 0xd1, 0x4b, 0xa6, 0x86, 0x21,
	0x2f, 0x68, 0x4d, 0xe2, 0x47, 0x12, 0x6e, 0xbf, 0x33, 0x01, 0xa4, 0xe8, 0xfd, 0x0b, 0x1a, 0xf9,
	0x0c, 0x3f, 0xd7, 0x63, 0x6b, 0xa9, 0xb1, 0x7d, 0x5a, 0xfc, 0x0c, 0x33, 0xc6, 0x83, 0xc9, 0x7d,
	0x09, 0x95, 0x88, 0x7b, 0x6c, 0x18, 0x78, 0xba, 0x28, 0x4d, 0x19, 0xbc, 0x7b, 0xff, 0xac, 0x7c,
	0xcc, 0x3d, 0xd6, 0x3b, 0x18, 0x94, 0x65, 0xb8, 0xe7, 0x15, 0xfb, 0x62, 0x2f, 0xf4, 0x05, 0x37,
	0xc1, 0x0a, 0x3c, 0xdd, 0x08, 0xd0, 0xa7, 0xad, 0xde, 0xc1, 0xc0, 0x0a, 0xbc, 0x76, 0x08, 0x64,
	0x7e, 0xf9, 0x49, 0x10, 0xf9, 0x93, 0xb9, 0x48, 0xf3, 0xbf, 0x88, 0xb4, 0x3e, 0x26, 0xb2, 0xfd,
	0x87, 0x09, 0x8d, 0x79, 0x9e, 0xb3, 0x2e, 0xee, 0x01, 0x88, 0x84, 0x46, 0x69, 0x20, 0x02, 0x1e,
	0xe9, 0x1b, 0xb7, 0x1e, 0xb9, 0x71, 0xc6, 0xc9, 0x27, 0x72, 0x7e, 0x0a, 0xbf, 0x84, 0xca, 0x48,
	0xb1, 0xb2, 0x8e, 0x17, 0x9e, 0x94, 0x65, 0x6b, 0xf9, 0x17, 0xa6, 0xe9, 0xc5, 0x9a, 0x95, 0x16,
	0x6a, 0xb6, 0x7d, 0x08, 0xb5, 0xd9, 0xbb, 0x8b, 0x6b, 0x50, 0x57, 0x9b, 0x63, 0x9e, 0x84, 0x74,
	0x42, 0x0c, 0x7c, 0x02, 0x6b, 0x0a, 0x98, 0xe7, 0x27, 0x26, 0xfe, 0x0f, 0xd6, 0x97, 0xc0, 0xb3,
	0x2e, 0xb1, 0xb6, 0xff, 0xb2, 0xa0, 0x5e, 0x78, 0x96, 0x10, 0xa0, 0xdc, 0x4f, 0xfd, 0xc3, 0x69,
	0x4c, 0x0c, 0xac, 0x43, 0xa5, 0x9f, 0xfa, 0x7b, 0x8c, 0x0a, 0x62, 0xea, 0xcd, 0x9b, 0x84, 0xc7,
	0xc4, 0xd2, 0xac, 0xdd, 0x38, 0x26, 0x25, 0x6c, 0x02, 0x64, 0xeb, 0x01, 0x4b, 0x63, 0x62, 0x6b,
	0xe2, 0x19, 0x17, 0x8c, 0xac, 0x48, 0x6d, 0x7a, 0xa3, 0xa2, 0x65, 0x1d, 0x95, 0x4f, 0x00, 0xa9,
	0x20, 0x81, 0x86, 0xbc, 0x8c, 0xd1, 0x44, 0x9c, 0xcb, 0x5b, 0xaa, 0xb8, 0x01, 0xa4, 0x88, 0xa8,
	0x43, 0x35, 0x44, 0x68, 0xf6, 0x53, 0xff, 0x6d, 0x94, 0x30, 0x3a, 0xba, 0xa0, 0xe7, 0x13, 0x46,
	0x00, 0xd7, 0x61, 0x55, 0x27, 0x92, 0x5f, 0xdc, 0x34, 0x25, 0x75, 0x4d, 0xdb, 0xbf, 0x60, 0xa3,
	0x1f, 0xbf, 0x9d, 0xf2, 0x64, 0x1a, 0x92, 0x86, 0xb4, 0xdd, 0x4f, 0x7d, 0xd5, 0xa0, 0x31, 0x4b,
	0x8e, 0x18, 0xf5, 0x58, 0x42, 0x56, 0xf5, 0xe9, 0xd3, 0x20, 0x64, 0x7c, 0x2a, 0x8e, 0xf9, 0x4f,
	0xa4, 0xa9, 0xc5, 0x0c, 0x18, 0xf5, 0xd4, 0xff, 0x8e, 0xac, 0x69, 0x31, 0x33, 0x44, 0x89, 0x21,
	0xda, 0xef, 0x9b, 0x84, 0x29, 0x8b, 0xeb, 0xfa, 0x56, 0xbd, 0x57, 0x1c, 0xdc, 0xfe, 0xd5, 0x84,
	0x8d, 0xc7, 0xc6, 0x03, 0xb7, 0xc0, 0x79, 0x0c, 0xdf, 0x9d, 0x0a, 0x4e, 0x0c, 0xfc, 0x04, 0xfe,
	0xff, 0x58, 0xf4, 0x6b, 0x1e, 0x44, 0xa2, 0x17, 0xc6, 0x93, 0x60, 0x14, 0xc8, 0x56, 0x7c, 0x8c,
	0xf6, 0xfa, 0x4a, 0xd3, 0xac, 0xed, 0x6b, 0x68, 0x2e, 0x7e, 0x14, 0xb2, 0x18, 0x73, 0x64, 0xd7,
	0xf3, 0xe4, 0xf8, 0x13, 0x03, 0x9d, 0xa2, 0xd8, 0x01, 0x0b, 0xf9, 0x25, 0x53, 0x11, 0x73, 0x31,
	0xf2, 0x36, 0xf6, 0xa8, 0xc8, 0x22, 0xd6, 0xa2, 0x91, 0x5d, 0xcf, 0x3b, 0xca, 0xde, 0x1e, 0x15,
	0x2d, 0xed, 0xbd, 0xb8, 0xf9, 0xe0, 0x1a, 0xb7, 0x1f, 0x5c, 0xe3, 0xe6, 0xce, 0x35, 0x6f, 0xef,
	0x5c, 0xf3, 0x9f, 0x3b, 0xd7, 0xfc, 0xed, 0xde, 0x35, 0x7e, 0xbf, 0x77, 0x8d, 0xdb, 0x7b, 0xd7,
	0xf8, 0xfb, 0xde, 0x35, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xee, 0xe3, 0x39, 0x8b, 0xbb, 0x08,
	0x00, 0x00,
}

func (m *Entry) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}

func (m *SnapshotMetadata) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}

func (m *Snapshot) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}

func (m *Message) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}

func (m *HardState) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}

func (m *ConfState) Size() (n int) {
	marshal, _ := json.Marshal(m)
	return len(marshal)
}

var (
	ErrInvalidLengthRaft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRaft = fmt.Errorf("proto: unexpected end of group")
)
