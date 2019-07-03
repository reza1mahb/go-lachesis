// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package api

import (
	context "context"
	fmt "fmt"
	wire "github.com/Fantom-foundation/go-lachesis/src/inter/wire"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KnownEvents struct {
	Lasts                map[string]uint64 `protobuf:"bytes,1,rep,name=Lasts,proto3" json:"Lasts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	SuperFrameIndex      uint64            `protobuf:"varint,2,opt,name=SuperFrameIndex,proto3" json:"SuperFrameIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *KnownEvents) Reset()         { *m = KnownEvents{} }
func (m *KnownEvents) String() string { return proto.CompactTextString(m) }
func (*KnownEvents) ProtoMessage()    {}
func (*KnownEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *KnownEvents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KnownEvents.Unmarshal(m, b)
}
func (m *KnownEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KnownEvents.Marshal(b, m, deterministic)
}
func (m *KnownEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KnownEvents.Merge(m, src)
}
func (m *KnownEvents) XXX_Size() int {
	return xxx_messageInfo_KnownEvents.Size(m)
}
func (m *KnownEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_KnownEvents.DiscardUnknown(m)
}

var xxx_messageInfo_KnownEvents proto.InternalMessageInfo

func (m *KnownEvents) GetLasts() map[string]uint64 {
	if m != nil {
		return m.Lasts
	}
	return nil
}

func (m *KnownEvents) GetSuperFrameIndex() uint64 {
	if m != nil {
		return m.SuperFrameIndex
	}
	return 0
}

type EventRequest struct {
	PeerID               string   `protobuf:"bytes,1,opt,name=PeerID,proto3" json:"PeerID,omitempty"`
	Index                uint64   `protobuf:"varint,2,opt,name=Index,proto3" json:"Index,omitempty"`
	Hash                 []byte   `protobuf:"bytes,3,opt,name=Hash,proto3" json:"Hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (m *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(m, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

func (m *EventRequest) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *EventRequest) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type PeerRequest struct {
	PeerID               string   `protobuf:"bytes,1,opt,name=PeerID,proto3" json:"PeerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerRequest) Reset()         { *m = PeerRequest{} }
func (m *PeerRequest) String() string { return proto.CompactTextString(m) }
func (*PeerRequest) ProtoMessage()    {}
func (*PeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *PeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerRequest.Unmarshal(m, b)
}
func (m *PeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerRequest.Marshal(b, m, deterministic)
}
func (m *PeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerRequest.Merge(m, src)
}
func (m *PeerRequest) XXX_Size() int {
	return xxx_messageInfo_PeerRequest.Size(m)
}
func (m *PeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeerRequest proto.InternalMessageInfo

func (m *PeerRequest) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

type PeerInfo struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PubKey               []byte   `protobuf:"bytes,2,opt,name=PubKey,proto3" json:"PubKey,omitempty"`
	Host                 string   `protobuf:"bytes,3,opt,name=Host,proto3" json:"Host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerInfo) Reset()         { *m = PeerInfo{} }
func (m *PeerInfo) String() string { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()    {}
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *PeerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerInfo.Unmarshal(m, b)
}
func (m *PeerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerInfo.Marshal(b, m, deterministic)
}
func (m *PeerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerInfo.Merge(m, src)
}
func (m *PeerInfo) XXX_Size() int {
	return xxx_messageInfo_PeerInfo.Size(m)
}
func (m *PeerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PeerInfo proto.InternalMessageInfo

func (m *PeerInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PeerInfo) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *PeerInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func init() {
	proto.RegisterType((*KnownEvents)(nil), "api.KnownEvents")
	proto.RegisterMapType((map[string]uint64)(nil), "api.KnownEvents.LastsEntry")
	proto.RegisterType((*EventRequest)(nil), "api.EventRequest")
	proto.RegisterType((*PeerRequest)(nil), "api.PeerRequest")
	proto.RegisterType((*PeerInfo)(nil), "api.PeerInfo")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4f, 0xab, 0x50,
	0x10, 0xed, 0x85, 0xb6, 0x69, 0x87, 0xf6, 0xbd, 0xbe, 0x9b, 0xa7, 0x21, 0xb8, 0x21, 0x24, 0x26,
	0x2c, 0x0c, 0x6a, 0xdd, 0x34, 0xae, 0xfb, 0x61, 0x53, 0x63, 0x1a, 0xfa, 0x0b, 0x68, 0x3b, 0x46,
	0xa2, 0x5e, 0xf0, 0xde, 0x4b, 0x2b, 0x7f, 0xc5, 0xb5, 0x3f, 0xd4, 0x30, 0x90, 0x4a, 0xea, 0xc2,
	0x0d, 0x99, 0x73, 0xce, 0x9c, 0x39, 0xc3, 0x00, 0xf4, 0x15, 0xca, 0x5d, 0xbc, 0xc1, 0x20, 0x95,
	0x89, 0x4e, 0xb8, 0x19, 0xa5, 0xb1, 0x73, 0x12, 0x0b, 0x8d, 0xf2, 0x72, 0x1f, 0x4b, 0xa4, 0x47,
	0xa9, 0x79, 0x9f, 0x0c, 0xac, 0x85, 0x48, 0xf6, 0x62, 0xb2, 0x43, 0xa1, 0x15, 0xbf, 0x86, 0xd6,
	0x7d, 0xa4, 0xb4, 0xb2, 0x99, 0x6b, 0xfa, 0xd6, 0xf0, 0x2c, 0x88, 0xd2, 0x38, 0xa8, 0x35, 0x04,
	0xa4, 0x4e, 0x84, 0x96, 0x79, 0x58, 0x76, 0x72, 0x1f, 0xfe, 0xae, 0xb2, 0x14, 0xe5, 0x54, 0x46,
	0xaf, 0x38, 0x17, 0x5b, 0x7c, 0xb7, 0x0d, 0x97, 0xf9, 0xcd, 0xf0, 0x98, 0x76, 0x46, 0x00, 0xdf,
	0x76, 0x3e, 0x00, 0xf3, 0x19, 0x73, 0x9b, 0xb9, 0xcc, 0xef, 0x86, 0x45, 0xc9, 0xff, 0x43, 0x6b,
	0x17, 0xbd, 0x64, 0x58, 0xf9, 0x4b, 0x70, 0x6b, 0x8c, 0x98, 0xb7, 0x84, 0x1e, 0xe5, 0x87, 0xf8,
	0x96, 0xa1, 0xd2, 0xfc, 0x14, 0xda, 0x4b, 0x44, 0x39, 0x1f, 0x57, 0xf6, 0x0a, 0x15, 0x13, 0xea,
	0x1b, 0x94, 0x80, 0x73, 0x68, 0xde, 0x45, 0xea, 0xc9, 0x36, 0x5d, 0xe6, 0xf7, 0x42, 0xaa, 0xbd,
	0x73, 0xb0, 0x0a, 0xcf, 0x2f, 0x03, 0xbd, 0x29, 0x74, 0xa8, 0x12, 0x8f, 0x09, 0xff, 0x03, 0xc6,
	0x41, 0x37, 0xe6, 0x63, 0xf2, 0x64, 0xeb, 0x05, 0xe6, 0x94, 0xd6, 0x0b, 0x2b, 0x44, 0x71, 0x89,
	0xd2, 0x14, 0xd7, 0x0d, 0xa9, 0x1e, 0x7e, 0x30, 0x68, 0x3e, 0x24, 0x5b, 0xe4, 0x43, 0x80, 0x55,
	0x2e, 0x36, 0xd5, 0xb9, 0x07, 0xc7, 0xf7, 0x75, 0x7e, 0x30, 0x5e, 0x83, 0x5f, 0x40, 0x67, 0x86,
	0x9a, 0x20, 0xff, 0x47, 0x7a, 0xfd, 0x18, 0x8e, 0x15, 0xd0, 0x07, 0x25, 0xce, 0x6b, 0xf0, 0x2b,
	0xb0, 0x66, 0xa8, 0x0f, 0x5b, 0x97, 0x03, 0x6b, 0xef, 0xea, 0xf4, 0x0f, 0x4c, 0xd1, 0xe0, 0x35,
	0xd6, 0x6d, 0xfa, 0x17, 0x6e, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0xd6, 0x27, 0xff, 0x38,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	SyncEvents(ctx context.Context, in *KnownEvents, opts ...grpc.CallOption) (*KnownEvents, error)
	GetEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*wire.Event, error)
	GetPeerInfo(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*PeerInfo, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) SyncEvents(ctx context.Context, in *KnownEvents, opts ...grpc.CallOption) (*KnownEvents, error) {
	out := new(KnownEvents)
	err := c.cc.Invoke(ctx, "/api.Node/SyncEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*wire.Event, error) {
	out := new(wire.Event)
	err := c.cc.Invoke(ctx, "/api.Node/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetPeerInfo(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*PeerInfo, error) {
	out := new(PeerInfo)
	err := c.cc.Invoke(ctx, "/api.Node/GetPeerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	SyncEvents(context.Context, *KnownEvents) (*KnownEvents, error)
	GetEvent(context.Context, *EventRequest) (*wire.Event, error)
	GetPeerInfo(context.Context, *PeerRequest) (*PeerInfo, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_SyncEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KnownEvents)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).SyncEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/SyncEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).SyncEvents(ctx, req.(*KnownEvents))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetPeerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetPeerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/GetPeerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetPeerInfo(ctx, req.(*PeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncEvents",
			Handler:    _Node_SyncEvents_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _Node_GetEvent_Handler,
		},
		{
			MethodName: "GetPeerInfo",
			Handler:    _Node_GetPeerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
