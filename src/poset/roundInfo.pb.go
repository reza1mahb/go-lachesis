// Code generated by protoc-gen-go. DO NOT EDIT.
// source: roundInfo.proto

package poset

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Trilean int32

const (
	Trilean_UNDEFINED Trilean = 0
	Trilean_TRUE      Trilean = 1
	Trilean_FALSE     Trilean = 2
)

var Trilean_name = map[int32]string{
	0: "UNDEFINED",
	1: "TRUE",
	2: "FALSE",
}

var Trilean_value = map[string]int32{
	"UNDEFINED": 0,
	"TRUE":      1,
	"FALSE":     2,
}

func (x Trilean) String() string {
	return proto.EnumName(Trilean_name, int32(x))
}

func (Trilean) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_06add9379b5a2b9a, []int{0}
}

type RoundEvent struct {
	Consensus            bool     `protobuf:"varint,1,opt,name=Consensus,proto3" json:"Consensus,omitempty"`
	Clotho               bool     `protobuf:"varint,2,opt,name=Clotho,proto3" json:"Clotho,omitempty"`
	Atropos              Trilean  `protobuf:"varint,3,opt,name=Atropos,proto3,enum=poset.Trilean" json:"Atropos,omitempty"`
	RoundReceived        int64    `protobuf:"varint,4,opt,name=RoundReceived,proto3" json:"RoundReceived,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoundEvent) Reset()         { *m = RoundEvent{} }
func (m *RoundEvent) String() string { return proto.CompactTextString(m) }
func (*RoundEvent) ProtoMessage()    {}
func (*RoundEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_06add9379b5a2b9a, []int{0}
}

func (m *RoundEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundEvent.Unmarshal(m, b)
}
func (m *RoundEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundEvent.Marshal(b, m, deterministic)
}
func (m *RoundEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundEvent.Merge(m, src)
}
func (m *RoundEvent) XXX_Size() int {
	return xxx_messageInfo_RoundEvent.Size(m)
}
func (m *RoundEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RoundEvent proto.InternalMessageInfo

func (m *RoundEvent) GetConsensus() bool {
	if m != nil {
		return m.Consensus
	}
	return false
}

func (m *RoundEvent) GetClotho() bool {
	if m != nil {
		return m.Clotho
	}
	return false
}

func (m *RoundEvent) GetAtropos() Trilean {
	if m != nil {
		return m.Atropos
	}
	return Trilean_UNDEFINED
}

func (m *RoundEvent) GetRoundReceived() int64 {
	if m != nil {
		return m.RoundReceived
	}
	return 0
}

type RoundCreatedMessage struct {
	Events               map[string]*RoundEvent `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Queued               bool                   `protobuf:"varint,2,opt,name=queued,proto3" json:"queued,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RoundCreatedMessage) Reset()         { *m = RoundCreatedMessage{} }
func (m *RoundCreatedMessage) String() string { return proto.CompactTextString(m) }
func (*RoundCreatedMessage) ProtoMessage()    {}
func (*RoundCreatedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_06add9379b5a2b9a, []int{1}
}

func (m *RoundCreatedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundCreatedMessage.Unmarshal(m, b)
}
func (m *RoundCreatedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundCreatedMessage.Marshal(b, m, deterministic)
}
func (m *RoundCreatedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundCreatedMessage.Merge(m, src)
}
func (m *RoundCreatedMessage) XXX_Size() int {
	return xxx_messageInfo_RoundCreatedMessage.Size(m)
}
func (m *RoundCreatedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundCreatedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RoundCreatedMessage proto.InternalMessageInfo

func (m *RoundCreatedMessage) GetEvents() map[string]*RoundEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *RoundCreatedMessage) GetQueued() bool {
	if m != nil {
		return m.Queued
	}
	return false
}

type RoundReceived struct {
	Rounds               [][]byte `protobuf:"bytes,1,rep,name=Rounds,proto3" json:"Rounds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoundReceived) Reset()         { *m = RoundReceived{} }
func (m *RoundReceived) String() string { return proto.CompactTextString(m) }
func (*RoundReceived) ProtoMessage()    {}
func (*RoundReceived) Descriptor() ([]byte, []int) {
	return fileDescriptor_06add9379b5a2b9a, []int{2}
}

func (m *RoundReceived) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundReceived.Unmarshal(m, b)
}
func (m *RoundReceived) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundReceived.Marshal(b, m, deterministic)
}
func (m *RoundReceived) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundReceived.Merge(m, src)
}
func (m *RoundReceived) XXX_Size() int {
	return xxx_messageInfo_RoundReceived.Size(m)
}
func (m *RoundReceived) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundReceived.DiscardUnknown(m)
}

var xxx_messageInfo_RoundReceived proto.InternalMessageInfo

func (m *RoundReceived) GetRounds() [][]byte {
	if m != nil {
		return m.Rounds
	}
	return nil
}

func init() {
	proto.RegisterEnum("poset.Trilean", Trilean_name, Trilean_value)
	proto.RegisterType((*RoundEvent)(nil), "poset.RoundEvent")
	proto.RegisterType((*RoundCreatedMessage)(nil), "poset.RoundCreatedMessage")
	proto.RegisterMapType((map[string]*RoundEvent)(nil), "poset.RoundCreatedMessage.EventsEntry")
	proto.RegisterType((*RoundReceived)(nil), "poset.RoundReceived")
}

func init() { proto.RegisterFile("roundInfo.proto", fileDescriptor_06add9379b5a2b9a) }

var fileDescriptor_06add9379b5a2b9a = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xdd, 0x4a, 0xc3, 0x30,
	0x14, 0x36, 0xeb, 0xfe, 0x7a, 0xe6, 0x66, 0x8d, 0x20, 0x45, 0xbc, 0x18, 0x43, 0x5c, 0x11, 0xec,
	0xc5, 0xbc, 0x11, 0x2f, 0x84, 0xb1, 0x65, 0x30, 0x98, 0xbb, 0x88, 0xdb, 0x03, 0x54, 0x7b, 0xd4,
	0x61, 0x49, 0x6a, 0x92, 0x16, 0xf6, 0x1a, 0xbe, 0x8e, 0x2f, 0x27, 0x4d, 0x3b, 0x9c, 0xe2, 0x5d,
	0xce, 0xf7, 0x93, 0xef, 0x3b, 0x09, 0x1c, 0x29, 0x99, 0x89, 0x78, 0x2e, 0x5e, 0x64, 0x98, 0x2a,
	0x69, 0x24, 0x6d, 0xa4, 0x52, 0xa3, 0x19, 0x7c, 0x12, 0x00, 0x5e, 0x50, 0x2c, 0x47, 0x61, 0xe8,
	0x39, 0xb8, 0x13, 0x29, 0x34, 0x0a, 0x9d, 0x69, 0x9f, 0xf4, 0x49, 0xd0, 0xe6, 0x3f, 0x00, 0x3d,
	0x85, 0xe6, 0x24, 0x91, 0xe6, 0x4d, 0xfa, 0x35, 0x4b, 0x55, 0x13, 0x0d, 0xa0, 0x35, 0x36, 0x4a,
	0xa6, 0x52, 0xfb, 0x4e, 0x9f, 0x04, 0xbd, 0x51, 0x2f, 0xb4, 0xb7, 0x87, 0x2b, 0xb5, 0x49, 0x30,
	0x12, 0x7c, 0x47, 0xd3, 0x0b, 0xe8, 0xda, 0x34, 0x8e, 0xcf, 0xb8, 0xc9, 0x31, 0xf6, 0xeb, 0x7d,
	0x12, 0x38, 0xfc, 0x37, 0x38, 0xf8, 0x22, 0x70, 0x62, 0x91, 0x89, 0xc2, 0xc8, 0x60, 0xfc, 0x80,
	0x5a, 0x47, 0xaf, 0x48, 0xef, 0xa1, 0x69, 0x6b, 0x16, 0xd5, 0x9c, 0xa0, 0x33, 0xba, 0xac, 0x62,
	0xfe, 0xd1, 0x86, 0xa5, 0x90, 0x09, 0xa3, 0xb6, 0xbc, 0x72, 0x15, 0xfd, 0x3f, 0x32, 0xcc, 0x30,
	0xde, 0xf5, 0x2f, 0xa7, 0xb3, 0x05, 0x74, 0xf6, 0xe4, 0xd4, 0x03, 0xe7, 0x1d, 0xb7, 0x76, 0x7d,
	0x97, 0x17, 0x47, 0x3a, 0x84, 0x46, 0x1e, 0x25, 0x19, 0x5a, 0x5f, 0x67, 0x74, 0xbc, 0x9f, 0x6b,
	0x9d, 0xbc, 0xe4, 0xef, 0x6a, 0xb7, 0x64, 0x30, 0xfc, 0xb3, 0x63, 0x11, 0x6b, 0x81, 0xb2, 0xf6,
	0x21, 0xaf, 0xa6, 0xab, 0x6b, 0x68, 0x55, 0x0f, 0x44, 0xbb, 0xe0, 0xae, 0x97, 0x53, 0x36, 0x9b,
	0x2f, 0xd9, 0xd4, 0x3b, 0xa0, 0x6d, 0xa8, 0xaf, 0xf8, 0x9a, 0x79, 0x84, 0xba, 0xd0, 0x98, 0x8d,
	0x17, 0x8f, 0xcc, 0xab, 0x3d, 0x35, 0xed, 0xc7, 0xdd, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0x8f,
	0x14, 0xf6, 0x49, 0xcb, 0x01, 0x00, 0x00,
}
