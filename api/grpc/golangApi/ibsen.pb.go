// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ibsen.proto

package golangApi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EntryBatch struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Offset               uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Marker               int64    `protobuf:"varint,3,opt,name=marker,proto3" json:"marker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryBatch) Reset()         { *m = EntryBatch{} }
func (m *EntryBatch) String() string { return proto.CompactTextString(m) }
func (*EntryBatch) ProtoMessage()    {}
func (*EntryBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{0}
}

func (m *EntryBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryBatch.Unmarshal(m, b)
}
func (m *EntryBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryBatch.Marshal(b, m, deterministic)
}
func (m *EntryBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryBatch.Merge(m, src)
}
func (m *EntryBatch) XXX_Size() int {
	return xxx_messageInfo_EntryBatch.Size(m)
}
func (m *EntryBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryBatch.DiscardUnknown(m)
}

var xxx_messageInfo_EntryBatch proto.InternalMessageInfo

func (m *EntryBatch) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *EntryBatch) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *EntryBatch) GetMarker() int64 {
	if m != nil {
		return m.Marker
	}
	return 0
}

type EntryBatchResponse struct {
	NextBatch            *EntryBatch `protobuf:"bytes,1,opt,name=nextBatch,proto3" json:"nextBatch,omitempty"`
	Entries              [][]byte    `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EntryBatchResponse) Reset()         { *m = EntryBatchResponse{} }
func (m *EntryBatchResponse) String() string { return proto.CompactTextString(m) }
func (*EntryBatchResponse) ProtoMessage()    {}
func (*EntryBatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{1}
}

func (m *EntryBatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryBatchResponse.Unmarshal(m, b)
}
func (m *EntryBatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryBatchResponse.Marshal(b, m, deterministic)
}
func (m *EntryBatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryBatchResponse.Merge(m, src)
}
func (m *EntryBatchResponse) XXX_Size() int {
	return xxx_messageInfo_EntryBatchResponse.Size(m)
}
func (m *EntryBatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryBatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EntryBatchResponse proto.InternalMessageInfo

func (m *EntryBatchResponse) GetNextBatch() *EntryBatch {
	if m != nil {
		return m.NextBatch
	}
	return nil
}

func (m *EntryBatchResponse) GetEntries() [][]byte {
	if m != nil {
		return m.Entries
	}
	return nil
}

type TopicStatus struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicStatus) Reset()         { *m = TopicStatus{} }
func (m *TopicStatus) String() string { return proto.CompactTextString(m) }
func (*TopicStatus) ProtoMessage()    {}
func (*TopicStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{2}
}

func (m *TopicStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicStatus.Unmarshal(m, b)
}
func (m *TopicStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicStatus.Marshal(b, m, deterministic)
}
func (m *TopicStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicStatus.Merge(m, src)
}
func (m *TopicStatus) XXX_Size() int {
	return xxx_messageInfo_TopicStatus.Size(m)
}
func (m *TopicStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TopicStatus proto.InternalMessageInfo

func (m *TopicStatus) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Offset struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Offset) Reset()         { *m = Offset{} }
func (m *Offset) String() string { return proto.CompactTextString(m) }
func (*Offset) ProtoMessage()    {}
func (*Offset) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{4}
}

func (m *Offset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Offset.Unmarshal(m, b)
}
func (m *Offset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Offset.Marshal(b, m, deterministic)
}
func (m *Offset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Offset.Merge(m, src)
}
func (m *Offset) XXX_Size() int {
	return xxx_messageInfo_Offset.Size(m)
}
func (m *Offset) XXX_DiscardUnknown() {
	xxx_messageInfo_Offset.DiscardUnknown(m)
}

var xxx_messageInfo_Offset proto.InternalMessageInfo

func (m *Offset) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Topic struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{5}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TopicOffset struct {
	TopicName            string   `protobuf:"bytes,1,opt,name=topicName,proto3" json:"topicName,omitempty"`
	Offset               uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicOffset) Reset()         { *m = TopicOffset{} }
func (m *TopicOffset) String() string { return proto.CompactTextString(m) }
func (*TopicOffset) ProtoMessage()    {}
func (*TopicOffset) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{6}
}

func (m *TopicOffset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicOffset.Unmarshal(m, b)
}
func (m *TopicOffset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicOffset.Marshal(b, m, deterministic)
}
func (m *TopicOffset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicOffset.Merge(m, src)
}
func (m *TopicOffset) XXX_Size() int {
	return xxx_messageInfo_TopicOffset.Size(m)
}
func (m *TopicOffset) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicOffset.DiscardUnknown(m)
}

var xxx_messageInfo_TopicOffset proto.InternalMessageInfo

func (m *TopicOffset) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *TopicOffset) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type TopicMessage struct {
	TopicName            string   `protobuf:"bytes,1,opt,name=topicName,proto3" json:"topicName,omitempty"`
	MessagePayload       []byte   `protobuf:"bytes,2,opt,name=messagePayload,proto3" json:"messagePayload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicMessage) Reset()         { *m = TopicMessage{} }
func (m *TopicMessage) String() string { return proto.CompactTextString(m) }
func (*TopicMessage) ProtoMessage()    {}
func (*TopicMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{7}
}

func (m *TopicMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicMessage.Unmarshal(m, b)
}
func (m *TopicMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicMessage.Marshal(b, m, deterministic)
}
func (m *TopicMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicMessage.Merge(m, src)
}
func (m *TopicMessage) XXX_Size() int {
	return xxx_messageInfo_TopicMessage.Size(m)
}
func (m *TopicMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TopicMessage proto.InternalMessageInfo

func (m *TopicMessage) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *TopicMessage) GetMessagePayload() []byte {
	if m != nil {
		return m.MessagePayload
	}
	return nil
}

type TopicBatchMessage struct {
	TopicName            string   `protobuf:"bytes,1,opt,name=topicName,proto3" json:"topicName,omitempty"`
	MessagePayload       [][]byte `protobuf:"bytes,2,rep,name=messagePayload,proto3" json:"messagePayload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicBatchMessage) Reset()         { *m = TopicBatchMessage{} }
func (m *TopicBatchMessage) String() string { return proto.CompactTextString(m) }
func (*TopicBatchMessage) ProtoMessage()    {}
func (*TopicBatchMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{8}
}

func (m *TopicBatchMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicBatchMessage.Unmarshal(m, b)
}
func (m *TopicBatchMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicBatchMessage.Marshal(b, m, deterministic)
}
func (m *TopicBatchMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicBatchMessage.Merge(m, src)
}
func (m *TopicBatchMessage) XXX_Size() int {
	return xxx_messageInfo_TopicBatchMessage.Size(m)
}
func (m *TopicBatchMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicBatchMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TopicBatchMessage proto.InternalMessageInfo

func (m *TopicBatchMessage) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *TopicBatchMessage) GetMessagePayload() [][]byte {
	if m != nil {
		return m.MessagePayload
	}
	return nil
}

type Message struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{9}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Entry struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{10}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Entry) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Status struct {
	Entries              int32    `protobuf:"varint,1,opt,name=entries,proto3" json:"entries,omitempty"`
	Current              *Offset  `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{11}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetEntries() int32 {
	if m != nil {
		return m.Entries
	}
	return 0
}

func (m *Status) GetCurrent() *Offset {
	if m != nil {
		return m.Current
	}
	return nil
}

func init() {
	proto.RegisterType((*EntryBatch)(nil), "EntryBatch")
	proto.RegisterType((*EntryBatchResponse)(nil), "EntryBatchResponse")
	proto.RegisterType((*TopicStatus)(nil), "TopicStatus")
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*Offset)(nil), "Offset")
	proto.RegisterType((*Topic)(nil), "Topic")
	proto.RegisterType((*TopicOffset)(nil), "TopicOffset")
	proto.RegisterType((*TopicMessage)(nil), "TopicMessage")
	proto.RegisterType((*TopicBatchMessage)(nil), "TopicBatchMessage")
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Entry)(nil), "Entry")
	proto.RegisterType((*Status)(nil), "Status")
}

func init() { proto.RegisterFile("ibsen.proto", fileDescriptor_c3e5c14afe4be539) }

var fileDescriptor_c3e5c14afe4be539 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x97, 0xb6, 0x49, 0xe8, 0xa5, 0x54, 0x9a, 0x37, 0x50, 0x54, 0x10, 0xca, 0x8c, 0xc4,
	0x3a, 0x90, 0xd2, 0x51, 0x9e, 0x78, 0xa4, 0xa8, 0x48, 0x3c, 0xf0, 0x43, 0xd9, 0x24, 0x34, 0xde,
	0xbc, 0xc6, 0xcb, 0x2c, 0x1a, 0x27, 0x72, 0x3c, 0x41, 0xff, 0x1b, 0xfe, 0x54, 0x94, 0xb3, 0xd3,
	0xa4, 0x40, 0x85, 0xc4, 0x5b, 0xbf, 0xe7, 0xef, 0x7d, 0x7a, 0xbe, 0x3b, 0x07, 0x02, 0x71, 0x5d,
	0x71, 0x19, 0x97, 0xaa, 0xd0, 0x05, 0x4d, 0x00, 0x96, 0x52, 0xab, 0xcd, 0x82, 0xe9, 0xd5, 0x2d,
	0x39, 0x06, 0x57, 0x17, 0xa5, 0x58, 0x85, 0x4e, 0xe4, 0x4c, 0x87, 0x89, 0x11, 0xe4, 0x21, 0x78,
	0xc5, 0xcd, 0x4d, 0xc5, 0x75, 0xd8, 0x8b, 0x9c, 0xe9, 0x20, 0xb1, 0xaa, 0x8e, 0xe7, 0x4c, 0x7d,
	0xe3, 0x2a, 0xec, 0x47, 0xce, 0xb4, 0x9f, 0x58, 0x45, 0xaf, 0x80, 0xb4, 0xcc, 0x84, 0x57, 0x65,
	0x21, 0x2b, 0x4e, 0xce, 0x60, 0x28, 0xf9, 0x0f, 0x8d, 0x41, 0xe4, 0x07, 0xf3, 0x20, 0xee, 0xf8,
	0xda, 0x53, 0x12, 0x82, 0xcf, 0xa5, 0x56, 0x82, 0x57, 0x61, 0x2f, 0xea, 0x4f, 0x47, 0x49, 0x23,
	0xe9, 0x29, 0x04, 0x97, 0x75, 0x4d, 0x17, 0x9a, 0xe9, 0xbb, 0xaa, 0x36, 0xae, 0x14, 0x67, 0x9a,
	0xa7, 0x48, 0xbc, 0x97, 0x34, 0x92, 0xfa, 0xe0, 0x2e, 0xf3, 0x52, 0x6f, 0x68, 0x08, 0xde, 0x27,
	0x53, 0xee, 0x18, 0x7a, 0xc2, 0xf8, 0x06, 0x49, 0x4f, 0xa4, 0xf4, 0x11, 0xb8, 0xc8, 0x22, 0x04,
	0x06, 0x92, 0xe5, 0xdc, 0x5e, 0x1a, 0x7f, 0xd3, 0xb7, 0xf6, 0x8f, 0x6c, 0xee, 0x63, 0x18, 0x62,
	0x2f, 0x3e, 0xb6, 0xbe, 0x36, 0xb0, 0xaf, 0x41, 0xf4, 0x12, 0x46, 0x08, 0xf9, 0xc0, 0xab, 0x8a,
	0x65, 0xfc, 0x1f, 0x94, 0x67, 0x30, 0xce, 0x8d, 0xf1, 0x33, 0xdb, 0xac, 0x0b, 0x96, 0x22, 0x6d,
	0x94, 0xfc, 0x16, 0xa5, 0x57, 0x70, 0x88, 0x54, 0xec, 0xd5, 0xff, 0xa3, 0xfb, 0x7f, 0x41, 0x3f,
	0x05, 0xbf, 0x01, 0x86, 0xe0, 0x97, 0xd6, 0xeb, 0x60, 0x19, 0x8d, 0xa4, 0xaf, 0xc1, 0xc5, 0xb1,
	0x75, 0xae, 0xed, 0xec, 0xec, 0x45, 0x27, 0xb5, 0xb7, 0x9b, 0xba, 0x04, 0xaf, 0x9d, 0x5c, 0x33,
	0xe2, 0x3a, 0xd9, 0xdd, 0x8e, 0x98, 0x9c, 0x80, 0xbf, 0xba, 0x53, 0x8a, 0x4b, 0xd3, 0xcd, 0x60,
	0xee, 0xc7, 0x66, 0x08, 0x49, 0x13, 0x9f, 0xff, 0xec, 0x83, 0xfb, 0xbe, 0x5e, 0x62, 0x12, 0x81,
	0x67, 0x26, 0x4e, 0xbc, 0x18, 0x9b, 0x32, 0x19, 0xc5, 0x9d, 0x05, 0xa1, 0x07, 0xe4, 0x09, 0x0c,
	0x52, 0x55, 0x94, 0x7b, 0xcf, 0x4f, 0xc0, 0xfd, 0xae, 0x84, 0xe6, 0xe4, 0x7e, 0xdc, 0x9d, 0xd5,
	0xc4, 0x8f, 0xb7, 0x96, 0x17, 0x00, 0x68, 0x31, 0xcb, 0x49, 0xe2, 0x3f, 0xba, 0xdf, 0x35, 0x9f,
	0x41, 0x80, 0xe6, 0x0b, 0xad, 0x38, 0xcb, 0xf7, 0x53, 0xa7, 0x0e, 0x39, 0x85, 0x43, 0xc5, 0x59,
	0xfa, 0x4e, 0x15, 0xf9, 0x82, 0x67, 0x42, 0x4a, 0x21, 0xb3, 0x6d, 0x9d, 0x9e, 0x79, 0x1b, 0xf4,
	0xe0, 0xdc, 0x21, 0xcf, 0x61, 0xdc, 0x18, 0xed, 0x3e, 0xda, 0x5b, 0x18, 0xb5, 0xe3, 0x7d, 0x09,
	0x47, 0xb5, 0x17, 0xcb, 0xeb, 0x24, 0x74, 0x9f, 0xda, 0xa4, 0x2b, 0xe8, 0x01, 0x89, 0x00, 0xd6,
	0xa2, 0xd2, 0x48, 0xac, 0x88, 0x17, 0xe3, 0xc3, 0x99, 0xd8, 0x42, 0x10, 0x7a, 0x0e, 0xc7, 0xad,
	0xe3, 0x8b, 0xd0, 0xb7, 0x96, 0xda, 0x78, 0x77, 0xca, 0xa9, 0x33, 0x16, 0x0f, 0xbe, 0x1e, 0xcd,
	0x58, 0x29, 0x66, 0x99, 0x2a, 0x57, 0xb3, 0xac, 0x58, 0x33, 0x99, 0xbd, 0x29, 0xc5, 0xb5, 0x87,
	0x5f, 0x9d, 0x57, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xa3, 0x0a, 0x63, 0x84, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IbsenClient is the client API for Ibsen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IbsenClient interface {
	Create(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*TopicStatus, error)
	Drop(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*TopicStatus, error)
	Write(ctx context.Context, in *TopicMessage, opts ...grpc.CallOption) (*Status, error)
	WriteBatch(ctx context.Context, in *TopicBatchMessage, opts ...grpc.CallOption) (*Status, error)
	WriteStream(ctx context.Context, opts ...grpc.CallOption) (Ibsen_WriteStreamClient, error)
	ReadFromBeginning(ctx context.Context, in *Topic, opts ...grpc.CallOption) (Ibsen_ReadFromBeginningClient, error)
	ReadFromOffset(ctx context.Context, in *TopicOffset, opts ...grpc.CallOption) (Ibsen_ReadFromOffsetClient, error)
	ReadBatchFromOffset(ctx context.Context, in *EntryBatch, opts ...grpc.CallOption) (*EntryBatch, error)
	ListTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Ibsen_ListTopicsClient, error)
	ListTopicsWithOffset(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Ibsen_ListTopicsWithOffsetClient, error)
}

type ibsenClient struct {
	cc grpc.ClientConnInterface
}

func NewIbsenClient(cc grpc.ClientConnInterface) IbsenClient {
	return &ibsenClient{cc}
}

func (c *ibsenClient) Create(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*TopicStatus, error) {
	out := new(TopicStatus)
	err := c.cc.Invoke(ctx, "/Ibsen/create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ibsenClient) Drop(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*TopicStatus, error) {
	out := new(TopicStatus)
	err := c.cc.Invoke(ctx, "/Ibsen/drop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ibsenClient) Write(ctx context.Context, in *TopicMessage, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Ibsen/write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ibsenClient) WriteBatch(ctx context.Context, in *TopicBatchMessage, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Ibsen/writeBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ibsenClient) WriteStream(ctx context.Context, opts ...grpc.CallOption) (Ibsen_WriteStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ibsen_serviceDesc.Streams[0], "/Ibsen/writeStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &ibsenWriteStreamClient{stream}
	return x, nil
}

type Ibsen_WriteStreamClient interface {
	Send(*TopicMessage) error
	CloseAndRecv() (*Status, error)
	grpc.ClientStream
}

type ibsenWriteStreamClient struct {
	grpc.ClientStream
}

func (x *ibsenWriteStreamClient) Send(m *TopicMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ibsenWriteStreamClient) CloseAndRecv() (*Status, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ibsenClient) ReadFromBeginning(ctx context.Context, in *Topic, opts ...grpc.CallOption) (Ibsen_ReadFromBeginningClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ibsen_serviceDesc.Streams[1], "/Ibsen/readFromBeginning", opts...)
	if err != nil {
		return nil, err
	}
	x := &ibsenReadFromBeginningClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ibsen_ReadFromBeginningClient interface {
	Recv() (*Entry, error)
	grpc.ClientStream
}

type ibsenReadFromBeginningClient struct {
	grpc.ClientStream
}

func (x *ibsenReadFromBeginningClient) Recv() (*Entry, error) {
	m := new(Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ibsenClient) ReadFromOffset(ctx context.Context, in *TopicOffset, opts ...grpc.CallOption) (Ibsen_ReadFromOffsetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ibsen_serviceDesc.Streams[2], "/Ibsen/readFromOffset", opts...)
	if err != nil {
		return nil, err
	}
	x := &ibsenReadFromOffsetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ibsen_ReadFromOffsetClient interface {
	Recv() (*Entry, error)
	grpc.ClientStream
}

type ibsenReadFromOffsetClient struct {
	grpc.ClientStream
}

func (x *ibsenReadFromOffsetClient) Recv() (*Entry, error) {
	m := new(Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ibsenClient) ReadBatchFromOffset(ctx context.Context, in *EntryBatch, opts ...grpc.CallOption) (*EntryBatch, error) {
	out := new(EntryBatch)
	err := c.cc.Invoke(ctx, "/Ibsen/readBatchFromOffset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ibsenClient) ListTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Ibsen_ListTopicsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ibsen_serviceDesc.Streams[3], "/Ibsen/listTopics", opts...)
	if err != nil {
		return nil, err
	}
	x := &ibsenListTopicsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ibsen_ListTopicsClient interface {
	Recv() (*Topic, error)
	grpc.ClientStream
}

type ibsenListTopicsClient struct {
	grpc.ClientStream
}

func (x *ibsenListTopicsClient) Recv() (*Topic, error) {
	m := new(Topic)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ibsenClient) ListTopicsWithOffset(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Ibsen_ListTopicsWithOffsetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ibsen_serviceDesc.Streams[4], "/Ibsen/listTopicsWithOffset", opts...)
	if err != nil {
		return nil, err
	}
	x := &ibsenListTopicsWithOffsetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ibsen_ListTopicsWithOffsetClient interface {
	Recv() (*TopicOffset, error)
	grpc.ClientStream
}

type ibsenListTopicsWithOffsetClient struct {
	grpc.ClientStream
}

func (x *ibsenListTopicsWithOffsetClient) Recv() (*TopicOffset, error) {
	m := new(TopicOffset)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IbsenServer is the server API for Ibsen service.
type IbsenServer interface {
	Create(context.Context, *Topic) (*TopicStatus, error)
	Drop(context.Context, *Topic) (*TopicStatus, error)
	Write(context.Context, *TopicMessage) (*Status, error)
	WriteBatch(context.Context, *TopicBatchMessage) (*Status, error)
	WriteStream(Ibsen_WriteStreamServer) error
	ReadFromBeginning(*Topic, Ibsen_ReadFromBeginningServer) error
	ReadFromOffset(*TopicOffset, Ibsen_ReadFromOffsetServer) error
	ReadBatchFromOffset(context.Context, *EntryBatch) (*EntryBatch, error)
	ListTopics(*Empty, Ibsen_ListTopicsServer) error
	ListTopicsWithOffset(*Empty, Ibsen_ListTopicsWithOffsetServer) error
}

// UnimplementedIbsenServer can be embedded to have forward compatible implementations.
type UnimplementedIbsenServer struct {
}

func (*UnimplementedIbsenServer) Create(ctx context.Context, req *Topic) (*TopicStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedIbsenServer) Drop(ctx context.Context, req *Topic) (*TopicStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Drop not implemented")
}
func (*UnimplementedIbsenServer) Write(ctx context.Context, req *TopicMessage) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedIbsenServer) WriteBatch(ctx context.Context, req *TopicBatchMessage) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteBatch not implemented")
}
func (*UnimplementedIbsenServer) WriteStream(srv Ibsen_WriteStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteStream not implemented")
}
func (*UnimplementedIbsenServer) ReadFromBeginning(req *Topic, srv Ibsen_ReadFromBeginningServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadFromBeginning not implemented")
}
func (*UnimplementedIbsenServer) ReadFromOffset(req *TopicOffset, srv Ibsen_ReadFromOffsetServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadFromOffset not implemented")
}
func (*UnimplementedIbsenServer) ReadBatchFromOffset(ctx context.Context, req *EntryBatch) (*EntryBatch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBatchFromOffset not implemented")
}
func (*UnimplementedIbsenServer) ListTopics(req *Empty, srv Ibsen_ListTopicsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}
func (*UnimplementedIbsenServer) ListTopicsWithOffset(req *Empty, srv Ibsen_ListTopicsWithOffsetServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTopicsWithOffset not implemented")
}

func RegisterIbsenServer(s *grpc.Server, srv IbsenServer) {
	s.RegisterService(&_Ibsen_serviceDesc, srv)
}

func _Ibsen_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IbsenServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ibsen/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IbsenServer).Create(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ibsen_Drop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IbsenServer).Drop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ibsen/Drop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IbsenServer).Drop(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ibsen_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IbsenServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ibsen/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IbsenServer).Write(ctx, req.(*TopicMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ibsen_WriteBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicBatchMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IbsenServer).WriteBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ibsen/WriteBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IbsenServer).WriteBatch(ctx, req.(*TopicBatchMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ibsen_WriteStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IbsenServer).WriteStream(&ibsenWriteStreamServer{stream})
}

type Ibsen_WriteStreamServer interface {
	SendAndClose(*Status) error
	Recv() (*TopicMessage, error)
	grpc.ServerStream
}

type ibsenWriteStreamServer struct {
	grpc.ServerStream
}

func (x *ibsenWriteStreamServer) SendAndClose(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ibsenWriteStreamServer) Recv() (*TopicMessage, error) {
	m := new(TopicMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ibsen_ReadFromBeginning_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Topic)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IbsenServer).ReadFromBeginning(m, &ibsenReadFromBeginningServer{stream})
}

type Ibsen_ReadFromBeginningServer interface {
	Send(*Entry) error
	grpc.ServerStream
}

type ibsenReadFromBeginningServer struct {
	grpc.ServerStream
}

func (x *ibsenReadFromBeginningServer) Send(m *Entry) error {
	return x.ServerStream.SendMsg(m)
}

func _Ibsen_ReadFromOffset_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TopicOffset)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IbsenServer).ReadFromOffset(m, &ibsenReadFromOffsetServer{stream})
}

type Ibsen_ReadFromOffsetServer interface {
	Send(*Entry) error
	grpc.ServerStream
}

type ibsenReadFromOffsetServer struct {
	grpc.ServerStream
}

func (x *ibsenReadFromOffsetServer) Send(m *Entry) error {
	return x.ServerStream.SendMsg(m)
}

func _Ibsen_ReadBatchFromOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntryBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IbsenServer).ReadBatchFromOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ibsen/ReadBatchFromOffset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IbsenServer).ReadBatchFromOffset(ctx, req.(*EntryBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ibsen_ListTopics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IbsenServer).ListTopics(m, &ibsenListTopicsServer{stream})
}

type Ibsen_ListTopicsServer interface {
	Send(*Topic) error
	grpc.ServerStream
}

type ibsenListTopicsServer struct {
	grpc.ServerStream
}

func (x *ibsenListTopicsServer) Send(m *Topic) error {
	return x.ServerStream.SendMsg(m)
}

func _Ibsen_ListTopicsWithOffset_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IbsenServer).ListTopicsWithOffset(m, &ibsenListTopicsWithOffsetServer{stream})
}

type Ibsen_ListTopicsWithOffsetServer interface {
	Send(*TopicOffset) error
	grpc.ServerStream
}

type ibsenListTopicsWithOffsetServer struct {
	grpc.ServerStream
}

func (x *ibsenListTopicsWithOffsetServer) Send(m *TopicOffset) error {
	return x.ServerStream.SendMsg(m)
}

var _Ibsen_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Ibsen",
	HandlerType: (*IbsenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Ibsen_Create_Handler,
		},
		{
			MethodName: "drop",
			Handler:    _Ibsen_Drop_Handler,
		},
		{
			MethodName: "write",
			Handler:    _Ibsen_Write_Handler,
		},
		{
			MethodName: "writeBatch",
			Handler:    _Ibsen_WriteBatch_Handler,
		},
		{
			MethodName: "readBatchFromOffset",
			Handler:    _Ibsen_ReadBatchFromOffset_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "writeStream",
			Handler:       _Ibsen_WriteStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "readFromBeginning",
			Handler:       _Ibsen_ReadFromBeginning_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "readFromOffset",
			Handler:       _Ibsen_ReadFromOffset_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listTopics",
			Handler:       _Ibsen_ListTopics_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listTopicsWithOffset",
			Handler:       _Ibsen_ListTopicsWithOffset_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ibsen.proto",
}
