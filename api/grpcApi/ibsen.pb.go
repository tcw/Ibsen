// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ibsen.proto

package grpcApi

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

type TopicsStatus struct {
	TopicStatus          []*TopicStatus `protobuf:"bytes,1,rep,name=topicStatus,proto3" json:"topicStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TopicsStatus) Reset()         { *m = TopicsStatus{} }
func (m *TopicsStatus) String() string { return proto.CompactTextString(m) }
func (*TopicsStatus) ProtoMessage()    {}
func (*TopicsStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{0}
}

func (m *TopicsStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicsStatus.Unmarshal(m, b)
}
func (m *TopicsStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicsStatus.Marshal(b, m, deterministic)
}
func (m *TopicsStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicsStatus.Merge(m, src)
}
func (m *TopicsStatus) XXX_Size() int {
	return xxx_messageInfo_TopicsStatus.Size(m)
}
func (m *TopicsStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicsStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TopicsStatus proto.InternalMessageInfo

func (m *TopicsStatus) GetTopicStatus() []*TopicStatus {
	if m != nil {
		return m.TopicStatus
	}
	return nil
}

type TopicStatus struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Blocks               int64    `protobuf:"varint,2,opt,name=blocks,proto3" json:"blocks,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	MaxBlockSize         int64    `protobuf:"varint,4,opt,name=maxBlockSize,proto3" json:"maxBlockSize,omitempty"`
	Path                 string   `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicStatus) Reset()         { *m = TopicStatus{} }
func (m *TopicStatus) String() string { return proto.CompactTextString(m) }
func (*TopicStatus) ProtoMessage()    {}
func (*TopicStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{1}
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

func (m *TopicStatus) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *TopicStatus) GetBlocks() int64 {
	if m != nil {
		return m.Blocks
	}
	return 0
}

func (m *TopicStatus) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *TopicStatus) GetMaxBlockSize() int64 {
	if m != nil {
		return m.MaxBlockSize
	}
	return 0
}

func (m *TopicStatus) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
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
	return fileDescriptor_c3e5c14afe4be539, []int{2}
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

type CreateStatus struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateStatus) Reset()         { *m = CreateStatus{} }
func (m *CreateStatus) String() string { return proto.CompactTextString(m) }
func (*CreateStatus) ProtoMessage()    {}
func (*CreateStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{3}
}

func (m *CreateStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStatus.Unmarshal(m, b)
}
func (m *CreateStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStatus.Marshal(b, m, deterministic)
}
func (m *CreateStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStatus.Merge(m, src)
}
func (m *CreateStatus) XXX_Size() int {
	return xxx_messageInfo_CreateStatus.Size(m)
}
func (m *CreateStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStatus.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStatus proto.InternalMessageInfo

func (m *CreateStatus) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

type DropStatus struct {
	Dropped              bool     `protobuf:"varint,1,opt,name=dropped,proto3" json:"dropped,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DropStatus) Reset()         { *m = DropStatus{} }
func (m *DropStatus) String() string { return proto.CompactTextString(m) }
func (*DropStatus) ProtoMessage()    {}
func (*DropStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{4}
}

func (m *DropStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DropStatus.Unmarshal(m, b)
}
func (m *DropStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DropStatus.Marshal(b, m, deterministic)
}
func (m *DropStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropStatus.Merge(m, src)
}
func (m *DropStatus) XXX_Size() int {
	return xxx_messageInfo_DropStatus.Size(m)
}
func (m *DropStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DropStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DropStatus proto.InternalMessageInfo

func (m *DropStatus) GetDropped() bool {
	if m != nil {
		return m.Dropped
	}
	return false
}

type WriteStatus struct {
	Wrote                int64    `protobuf:"varint,1,opt,name=wrote,proto3" json:"wrote,omitempty"`
	TimeNano             int64    `protobuf:"varint,2,opt,name=timeNano,proto3" json:"timeNano,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteStatus) Reset()         { *m = WriteStatus{} }
func (m *WriteStatus) String() string { return proto.CompactTextString(m) }
func (*WriteStatus) ProtoMessage()    {}
func (*WriteStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{5}
}

func (m *WriteStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteStatus.Unmarshal(m, b)
}
func (m *WriteStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteStatus.Marshal(b, m, deterministic)
}
func (m *WriteStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteStatus.Merge(m, src)
}
func (m *WriteStatus) XXX_Size() int {
	return xxx_messageInfo_WriteStatus.Size(m)
}
func (m *WriteStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteStatus.DiscardUnknown(m)
}

var xxx_messageInfo_WriteStatus proto.InternalMessageInfo

func (m *WriteStatus) GetWrote() int64 {
	if m != nil {
		return m.Wrote
	}
	return 0
}

func (m *WriteStatus) GetTimeNano() int64 {
	if m != nil {
		return m.TimeNano
	}
	return 0
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
	return fileDescriptor_c3e5c14afe4be539, []int{6}
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

type ReadParams struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Offset               uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	BatchSize            uint32   `protobuf:"varint,3,opt,name=batchSize,proto3" json:"batchSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadParams) Reset()         { *m = ReadParams{} }
func (m *ReadParams) String() string { return proto.CompactTextString(m) }
func (*ReadParams) ProtoMessage()    {}
func (*ReadParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{7}
}

func (m *ReadParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadParams.Unmarshal(m, b)
}
func (m *ReadParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadParams.Marshal(b, m, deterministic)
}
func (m *ReadParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadParams.Merge(m, src)
}
func (m *ReadParams) XXX_Size() int {
	return xxx_messageInfo_ReadParams.Size(m)
}
func (m *ReadParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadParams.DiscardUnknown(m)
}

var xxx_messageInfo_ReadParams proto.InternalMessageInfo

func (m *ReadParams) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ReadParams) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReadParams) GetBatchSize() uint32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

type InputEntries struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Entries              [][]byte `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputEntries) Reset()         { *m = InputEntries{} }
func (m *InputEntries) String() string { return proto.CompactTextString(m) }
func (*InputEntries) ProtoMessage()    {}
func (*InputEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{8}
}

func (m *InputEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputEntries.Unmarshal(m, b)
}
func (m *InputEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputEntries.Marshal(b, m, deterministic)
}
func (m *InputEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputEntries.Merge(m, src)
}
func (m *InputEntries) XXX_Size() int {
	return xxx_messageInfo_InputEntries.Size(m)
}
func (m *InputEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_InputEntries.DiscardUnknown(m)
}

var xxx_messageInfo_InputEntries proto.InternalMessageInfo

func (m *InputEntries) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *InputEntries) GetEntries() [][]byte {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Entry struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{9}
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

func (m *Entry) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type OutputEntries struct {
	Entries              []*Entry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputEntries) Reset()         { *m = OutputEntries{} }
func (m *OutputEntries) String() string { return proto.CompactTextString(m) }
func (*OutputEntries) ProtoMessage()    {}
func (*OutputEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3e5c14afe4be539, []int{10}
}

func (m *OutputEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputEntries.Unmarshal(m, b)
}
func (m *OutputEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputEntries.Marshal(b, m, deterministic)
}
func (m *OutputEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputEntries.Merge(m, src)
}
func (m *OutputEntries) XXX_Size() int {
	return xxx_messageInfo_OutputEntries.Size(m)
}
func (m *OutputEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputEntries.DiscardUnknown(m)
}

var xxx_messageInfo_OutputEntries proto.InternalMessageInfo

func (m *OutputEntries) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*TopicsStatus)(nil), "TopicsStatus")
	proto.RegisterType((*TopicStatus)(nil), "TopicStatus")
	proto.RegisterType((*Topic)(nil), "Topic")
	proto.RegisterType((*CreateStatus)(nil), "CreateStatus")
	proto.RegisterType((*DropStatus)(nil), "DropStatus")
	proto.RegisterType((*WriteStatus)(nil), "WriteStatus")
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*ReadParams)(nil), "ReadParams")
	proto.RegisterType((*InputEntries)(nil), "InputEntries")
	proto.RegisterType((*Entry)(nil), "Entry")
	proto.RegisterType((*OutputEntries)(nil), "OutputEntries")
}

func init() { proto.RegisterFile("ibsen.proto", fileDescriptor_c3e5c14afe4be539) }

var fileDescriptor_c3e5c14afe4be539 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x8e, 0xd3, 0x3a,
	0x14, 0xad, 0xdb, 0x26, 0x9d, 0xb9, 0x49, 0xdf, 0xc2, 0xaa, 0x9e, 0xa2, 0x02, 0x52, 0xf1, 0x62,
	0xc8, 0xca, 0x53, 0xca, 0x0a, 0x16, 0x20, 0x0a, 0x5d, 0x74, 0x33, 0x54, 0xe9, 0x48, 0x20, 0x76,
	0x6e, 0xea, 0x99, 0x5a, 0x90, 0xd8, 0x72, 0x5c, 0x95, 0xe1, 0x17, 0xf8, 0x0b, 0x7e, 0x90, 0x5f,
	0x40, 0xb6, 0xd3, 0x26, 0x5d, 0x00, 0x3b, 0x9f, 0x7b, 0x4f, 0xee, 0xb9, 0x39, 0xc7, 0x86, 0x48,
	0x6c, 0x2a, 0x5e, 0x52, 0xa5, 0xa5, 0x91, 0xe4, 0x35, 0xc4, 0xb7, 0x52, 0x89, 0xbc, 0x5a, 0x1b,
	0x66, 0xf6, 0x15, 0xa6, 0x10, 0x19, 0x8b, 0x3d, 0x4c, 0xd0, 0xa4, 0x97, 0x46, 0xb3, 0x98, 0xde,
	0x36, 0xb5, 0xac, 0x4d, 0x20, 0x3f, 0x10, 0x44, 0xad, 0x26, 0x1e, 0x41, 0xe0, 0xda, 0x09, 0x9a,
	0xa0, 0xf4, 0x32, 0xf3, 0x00, 0xff, 0x0f, 0xe1, 0xe6, 0xab, 0xcc, 0xbf, 0x54, 0x49, 0x77, 0x82,
	0xd2, 0x5e, 0x56, 0x23, 0x5b, 0x97, 0x77, 0x77, 0x15, 0x37, 0x49, 0xcf, 0xd7, 0x3d, 0xc2, 0x04,
	0xe2, 0x82, 0x7d, 0x9b, 0x5b, 0xd2, 0x5a, 0x7c, 0xe7, 0x49, 0xdf, 0x75, 0xcf, 0x6a, 0x18, 0x43,
	0x5f, 0x31, 0xb3, 0x4b, 0x02, 0x27, 0xe4, 0xce, 0xe4, 0x11, 0x04, 0x6e, 0x19, 0xdb, 0x2c, 0x59,
	0xc1, 0xeb, 0x2d, 0xdc, 0x99, 0xa4, 0x10, 0xbf, 0xd3, 0x9c, 0x19, 0x5e, 0xaf, 0x9a, 0xc0, 0x20,
	0x77, 0x78, 0xeb, 0x68, 0x17, 0xd9, 0x11, 0x92, 0x2b, 0x80, 0xf7, 0x5a, 0xaa, 0x86, 0xb7, 0xd5,
	0x52, 0xa9, 0x86, 0x57, 0x43, 0xf2, 0x06, 0xa2, 0x8f, 0x5a, 0x9c, 0x06, 0x8e, 0x20, 0x38, 0x68,
	0x69, 0xbc, 0x6a, 0x2f, 0xf3, 0x00, 0x8f, 0xe1, 0xc2, 0x88, 0x82, 0xdf, 0xb0, 0x52, 0xd6, 0x7f,
	0x7f, 0xc2, 0x64, 0x00, 0xc1, 0xa2, 0x50, 0xe6, 0x81, 0x7c, 0x02, 0xc8, 0x38, 0xdb, 0xae, 0x98,
	0x66, 0xc5, 0x5f, 0x4c, 0xac, 0xcd, 0xb2, 0x63, 0xfa, 0x27, 0xb3, 0x1e, 0xc3, 0xe5, 0x86, 0x99,
	0x7c, 0xe7, 0x9c, 0xb2, 0x3e, 0x0e, 0xb3, 0xa6, 0x60, 0x03, 0x5e, 0x96, 0x6a, 0x6f, 0x16, 0xa5,
	0xd1, 0x82, 0xff, 0x69, 0x76, 0x02, 0x03, 0xee, 0x09, 0x49, 0x77, 0xd2, 0x4b, 0xe3, 0xec, 0x08,
	0xc9, 0x4b, 0x08, 0xec, 0xa7, 0x0f, 0x2d, 0x79, 0x74, 0x26, 0x6f, 0x6d, 0x94, 0xa5, 0xe1, 0xa5,
	0xdf, 0x2b, 0xce, 0x8e, 0x90, 0x3c, 0x87, 0xe1, 0x87, 0xbd, 0x69, 0x69, 0x4f, 0xce, 0x55, 0xa2,
	0x59, 0x48, 0xdd, 0xec, 0x93, 0xda, 0xec, 0x17, 0x82, 0x60, 0x69, 0xaf, 0x27, 0x7e, 0x0a, 0xa1,
	0x8f, 0x03, 0x87, 0xfe, 0xf6, 0x8d, 0x87, 0xb4, 0x1d, 0x1f, 0xe9, 0xe0, 0x27, 0xd0, 0xb7, 0x49,
	0x9c, 0x08, 0x11, 0x6d, 0x52, 0x23, 0x1d, 0x3b, 0xa1, 0xf2, 0xc1, 0x84, 0xd4, 0xb9, 0x3c, 0x1e,
	0xd2, 0xf6, 0x5d, 0x27, 0x1d, 0x7c, 0x65, 0x13, 0x13, 0x86, 0xe3, 0x21, 0x6d, 0x9b, 0x34, 0x8e,
	0x69, 0x2b, 0x57, 0xd2, 0xc1, 0x53, 0x88, 0x0e, 0xbe, 0xa0, 0x39, 0x2b, 0xfe, 0xc1, 0x4e, 0xd1,
	0x14, 0xe1, 0x67, 0xd0, 0xd7, 0x9c, 0x6d, 0x71, 0x44, 0x9b, 0x5c, 0xc7, 0xff, 0xd1, 0x33, 0x3f,
	0x48, 0x67, 0x8a, 0xe6, 0xaf, 0x60, 0x74, 0x2f, 0xcc, 0x6e, 0xbf, 0xa1, 0xb9, 0x2c, 0xa8, 0xc9,
	0x0f, 0xd4, 0x3d, 0xcf, 0x39, 0x38, 0x1b, 0x56, 0xf6, 0x91, 0xae, 0xd0, 0xe7, 0xf8, 0x9a, 0x29,
	0x71, 0x7d, 0xaf, 0x55, 0xfe, 0x56, 0x89, 0x9f, 0xdd, 0x60, 0x39, 0x5f, 0x2f, 0x6e, 0x36, 0xa1,
	0x7b, 0xc3, 0x2f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x00, 0x1b, 0xd2, 0x4e, 0xd2, 0x03, 0x00,
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
	Create(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*CreateStatus, error)
	Drop(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*DropStatus, error)
	Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicsStatus, error)
	Write(ctx context.Context, in *InputEntries, opts ...grpc.CallOption) (*WriteStatus, error)
	WriteStream(ctx context.Context, opts ...grpc.CallOption) (Ibsen_WriteStreamClient, error)
	Read(ctx context.Context, in *ReadParams, opts ...grpc.CallOption) (Ibsen_ReadClient, error)
}

type ibsenClient struct {
	cc grpc.ClientConnInterface
}

func NewIbsenClient(cc grpc.ClientConnInterface) IbsenClient {
	return &ibsenClient{cc}
}

func (c *ibsenClient) Create(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*CreateStatus, error) {
	out := new(CreateStatus)
	err := c.cc.Invoke(ctx, "/Ibsen/create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ibsenClient) Drop(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*DropStatus, error) {
	out := new(DropStatus)
	err := c.cc.Invoke(ctx, "/Ibsen/drop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ibsenClient) Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicsStatus, error) {
	out := new(TopicsStatus)
	err := c.cc.Invoke(ctx, "/Ibsen/status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ibsenClient) Write(ctx context.Context, in *InputEntries, opts ...grpc.CallOption) (*WriteStatus, error) {
	out := new(WriteStatus)
	err := c.cc.Invoke(ctx, "/Ibsen/write", in, out, opts...)
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
	Send(*InputEntries) error
	Recv() (*WriteStatus, error)
	grpc.ClientStream
}

type ibsenWriteStreamClient struct {
	grpc.ClientStream
}

func (x *ibsenWriteStreamClient) Send(m *InputEntries) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ibsenWriteStreamClient) Recv() (*WriteStatus, error) {
	m := new(WriteStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ibsenClient) Read(ctx context.Context, in *ReadParams, opts ...grpc.CallOption) (Ibsen_ReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ibsen_serviceDesc.Streams[1], "/Ibsen/read", opts...)
	if err != nil {
		return nil, err
	}
	x := &ibsenReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ibsen_ReadClient interface {
	Recv() (*OutputEntries, error)
	grpc.ClientStream
}

type ibsenReadClient struct {
	grpc.ClientStream
}

func (x *ibsenReadClient) Recv() (*OutputEntries, error) {
	m := new(OutputEntries)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IbsenServer is the server API for Ibsen service.
type IbsenServer interface {
	Create(context.Context, *Topic) (*CreateStatus, error)
	Drop(context.Context, *Topic) (*DropStatus, error)
	Status(context.Context, *Empty) (*TopicsStatus, error)
	Write(context.Context, *InputEntries) (*WriteStatus, error)
	WriteStream(Ibsen_WriteStreamServer) error
	Read(*ReadParams, Ibsen_ReadServer) error
}

// UnimplementedIbsenServer can be embedded to have forward compatible implementations.
type UnimplementedIbsenServer struct {
}

func (*UnimplementedIbsenServer) Create(ctx context.Context, req *Topic) (*CreateStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedIbsenServer) Drop(ctx context.Context, req *Topic) (*DropStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Drop not implemented")
}
func (*UnimplementedIbsenServer) Status(ctx context.Context, req *Empty) (*TopicsStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedIbsenServer) Write(ctx context.Context, req *InputEntries) (*WriteStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedIbsenServer) WriteStream(srv Ibsen_WriteStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteStream not implemented")
}
func (*UnimplementedIbsenServer) Read(req *ReadParams, srv Ibsen_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
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

func _Ibsen_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IbsenServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ibsen/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IbsenServer).Status(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ibsen_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputEntries)
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
		return srv.(IbsenServer).Write(ctx, req.(*InputEntries))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ibsen_WriteStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IbsenServer).WriteStream(&ibsenWriteStreamServer{stream})
}

type Ibsen_WriteStreamServer interface {
	Send(*WriteStatus) error
	Recv() (*InputEntries, error)
	grpc.ServerStream
}

type ibsenWriteStreamServer struct {
	grpc.ServerStream
}

func (x *ibsenWriteStreamServer) Send(m *WriteStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ibsenWriteStreamServer) Recv() (*InputEntries, error) {
	m := new(InputEntries)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ibsen_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IbsenServer).Read(m, &ibsenReadServer{stream})
}

type Ibsen_ReadServer interface {
	Send(*OutputEntries) error
	grpc.ServerStream
}

type ibsenReadServer struct {
	grpc.ServerStream
}

func (x *ibsenReadServer) Send(m *OutputEntries) error {
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
			MethodName: "status",
			Handler:    _Ibsen_Status_Handler,
		},
		{
			MethodName: "write",
			Handler:    _Ibsen_Write_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "writeStream",
			Handler:       _Ibsen_WriteStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "read",
			Handler:       _Ibsen_Read_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ibsen.proto",
}
