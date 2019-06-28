// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package state

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

type Borrow struct {
	Recs                 map[uint64]uint64 `protobuf:"bytes,1,rep,name=Recs,proto3" json:"Recs,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Borrow) Reset()         { *m = Borrow{} }
func (m *Borrow) String() string { return proto.CompactTextString(m) }
func (*Borrow) ProtoMessage()    {}
func (*Borrow) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *Borrow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Borrow.Unmarshal(m, b)
}
func (m *Borrow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Borrow.Marshal(b, m, deterministic)
}
func (m *Borrow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Borrow.Merge(m, src)
}
func (m *Borrow) XXX_Size() int {
	return xxx_messageInfo_Borrow.Size(m)
}
func (m *Borrow) XXX_DiscardUnknown() {
	xxx_messageInfo_Borrow.DiscardUnknown(m)
}

var xxx_messageInfo_Borrow proto.InternalMessageInfo

func (m *Borrow) GetRecs() map[uint64]uint64 {
	if m != nil {
		return m.Recs
	}
	return nil
}

type Account struct {
	Balance              uint64             `protobuf:"varint,1,opt,name=Balance,proto3" json:"Balance,omitempty"`
	RawRoot              []byte             `protobuf:"bytes,2,opt,name=RawRoot,proto3" json:"RawRoot,omitempty"`
	DelegatedFrom        uint64             `protobuf:"varint,3,opt,name=DelegatedFrom,proto3" json:"DelegatedFrom,omitempty"`
	DelegatingFrom       map[string]*Borrow `protobuf:"bytes,4,rep,name=DelegatingFrom,proto3" json:"DelegatingFrom,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DelegatedTo          uint64             `protobuf:"varint,5,opt,name=DelegatedTo,proto3" json:"DelegatedTo,omitempty"`
	DelegatingTo         map[string]*Borrow `protobuf:"bytes,6,rep,name=DelegatingTo,proto3" json:"DelegatingTo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetRawRoot() []byte {
	if m != nil {
		return m.RawRoot
	}
	return nil
}

func (m *Account) GetDelegatedFrom() uint64 {
	if m != nil {
		return m.DelegatedFrom
	}
	return 0
}

func (m *Account) GetDelegatingFrom() map[string]*Borrow {
	if m != nil {
		return m.DelegatingFrom
	}
	return nil
}

func (m *Account) GetDelegatedTo() uint64 {
	if m != nil {
		return m.DelegatedTo
	}
	return 0
}

func (m *Account) GetDelegatingTo() map[string]*Borrow {
	if m != nil {
		return m.DelegatingTo
	}
	return nil
}

func init() {
	proto.RegisterType((*Borrow)(nil), "state.Borrow")
	proto.RegisterMapType((map[uint64]uint64)(nil), "state.Borrow.RecsEntry")
	proto.RegisterType((*Account)(nil), "state.Account")
	proto.RegisterMapType((map[string]*Borrow)(nil), "state.Account.DelegatingFromEntry")
	proto.RegisterMapType((map[string]*Borrow)(nil), "state.Account.DelegatingToEntry")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0xcd, 0x47, 0xe9, 0xa4, 0x11, 0x5d, 0x05, 0x97, 0x9c, 0x42, 0xf4, 0x50, 0x10, 0x72,
	0xa8, 0x07, 0xc5, 0x9b, 0xa5, 0x7a, 0xf0, 0x20, 0xb2, 0xe4, 0x0f, 0xac, 0x71, 0x29, 0x62, 0xdc,
	0x91, 0xed, 0xd6, 0xd2, 0x9f, 0xe6, 0xbf, 0x93, 0xcc, 0xa6, 0x35, 0xf1, 0xe3, 0xe4, 0x2d, 0xf3,
	0xde, 0xbc, 0x79, 0x33, 0x2f, 0x0b, 0x89, 0xac, 0x2a, 0x5c, 0x69, 0x5b, 0xbc, 0x19, 0xb4, 0xc8,
	0xc2, 0xa5, 0x95, 0x56, 0xe5, 0x1a, 0xa2, 0x19, 0x1a, 0x83, 0x6b, 0x76, 0x06, 0x81, 0x50, 0xd5,
	0x92, 0x7b, 0x99, 0x3f, 0x89, 0xa7, 0xc7, 0x05, 0xf1, 0x85, 0x23, 0x8b, 0x86, 0xb9, 0xd1, 0xd6,
	0x6c, 0x04, 0x35, 0xa5, 0x17, 0x30, 0xda, 0x41, 0x6c, 0x1f, 0xfc, 0x17, 0xb5, 0xe1, 0x5e, 0xe6,
	0x4d, 0x02, 0xd1, 0x7c, 0xb2, 0x23, 0x08, 0xdf, 0x65, 0xbd, 0x52, 0x7c, 0x40, 0x98, 0x2b, 0xae,
	0x06, 0x97, 0x5e, 0xfe, 0xe1, 0xc3, 0xf0, 0xda, 0x2d, 0xc2, 0x38, 0x0c, 0x67, 0xb2, 0x96, 0xba,
	0x52, 0xad, 0x76, 0x5b, 0x36, 0x8c, 0x90, 0x6b, 0x81, 0x68, 0x69, 0xc2, 0x58, 0x6c, 0x4b, 0x76,
	0x0a, 0xc9, 0x5c, 0xd5, 0x6a, 0x21, 0xad, 0x7a, 0xba, 0x35, 0xf8, 0xca, 0x7d, 0x52, 0xf6, 0x41,
	0x76, 0x07, 0x7b, 0x2d, 0xf0, 0xac, 0x17, 0xd4, 0x16, 0xd0, 0x55, 0x79, 0x7b, 0x55, 0xbb, 0x41,
	0xd1, 0x6f, 0x72, 0x07, 0x7e, 0x53, 0xb2, 0x0c, 0xe2, 0xdd, 0xf0, 0x12, 0x79, 0x48, 0x7e, 0x5d,
	0x88, 0xcd, 0x61, 0xfc, 0xa5, 0x29, 0x91, 0x47, 0xe4, 0x95, 0xfd, 0xe9, 0x55, 0xa2, 0x73, 0xea,
	0xa9, 0xd2, 0x07, 0x38, 0xfc, 0x65, 0x9d, 0x6e, 0xb8, 0x23, 0x17, 0xee, 0x49, 0x37, 0xdc, 0x78,
	0x9a, 0xf4, 0xfe, 0x54, 0x27, 0xeb, 0xf4, 0x1e, 0x0e, 0x7e, 0x98, 0xfe, 0x63, 0xde, 0x63, 0x44,
	0x2f, 0xe7, 0xfc, 0x33, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xf1, 0x3b, 0x5d, 0x4a, 0x02, 0x00, 0x00,
}
