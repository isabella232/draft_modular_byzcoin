// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package dkg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	onet "go.dedis.ch/phoenix/onet"
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

type Init struct {
	Addresses            []*onet.Address `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Init) Reset()         { *m = Init{} }
func (m *Init) String() string { return proto.CompactTextString(m) }
func (*Init) ProtoMessage()    {}
func (*Init) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Init) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Init.Unmarshal(m, b)
}
func (m *Init) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Init.Marshal(b, m, deterministic)
}
func (m *Init) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Init.Merge(m, src)
}
func (m *Init) XXX_Size() int {
	return xxx_messageInfo_Init.Size(m)
}
func (m *Init) XXX_DiscardUnknown() {
	xxx_messageInfo_Init.DiscardUnknown(m)
}

var xxx_messageInfo_Init proto.InternalMessageInfo

func (m *Init) GetAddresses() []*onet.Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type EncryptedDeal struct {
	DHKey                []byte   `protobuf:"bytes,1,opt,name=DHKey,proto3" json:"DHKey,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Nonce                []byte   `protobuf:"bytes,3,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Cipher               []byte   `protobuf:"bytes,4,opt,name=Cipher,proto3" json:"Cipher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptedDeal) Reset()         { *m = EncryptedDeal{} }
func (m *EncryptedDeal) String() string { return proto.CompactTextString(m) }
func (*EncryptedDeal) ProtoMessage()    {}
func (*EncryptedDeal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *EncryptedDeal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptedDeal.Unmarshal(m, b)
}
func (m *EncryptedDeal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptedDeal.Marshal(b, m, deterministic)
}
func (m *EncryptedDeal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedDeal.Merge(m, src)
}
func (m *EncryptedDeal) XXX_Size() int {
	return xxx_messageInfo_EncryptedDeal.Size(m)
}
func (m *EncryptedDeal) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedDeal.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedDeal proto.InternalMessageInfo

func (m *EncryptedDeal) GetDHKey() []byte {
	if m != nil {
		return m.DHKey
	}
	return nil
}

func (m *EncryptedDeal) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *EncryptedDeal) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *EncryptedDeal) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
}

type Deal struct {
	Index                uint32         `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Deal                 *EncryptedDeal `protobuf:"bytes,2,opt,name=deal,proto3" json:"deal,omitempty"`
	Signature            []byte         `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Deal) Reset()         { *m = Deal{} }
func (m *Deal) String() string { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()    {}
func (*Deal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *Deal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deal.Unmarshal(m, b)
}
func (m *Deal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deal.Marshal(b, m, deterministic)
}
func (m *Deal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deal.Merge(m, src)
}
func (m *Deal) XXX_Size() int {
	return xxx_messageInfo_Deal.Size(m)
}
func (m *Deal) XXX_DiscardUnknown() {
	xxx_messageInfo_Deal.DiscardUnknown(m)
}

var xxx_messageInfo_Deal proto.InternalMessageInfo

func (m *Deal) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Deal) GetDeal() *EncryptedDeal {
	if m != nil {
		return m.Deal
	}
	return nil
}

func (m *Deal) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Ack struct {
	Index                uint32        `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Response             *Ack_Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ack.Unmarshal(m, b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
}
func (m *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(m, src)
}
func (m *Ack) XXX_Size() int {
	return xxx_messageInfo_Ack.Size(m)
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func (m *Ack) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Ack) GetResponse() *Ack_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type Ack_Response struct {
	SessionID            []byte   `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	Index                uint32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Status               bool     `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ack_Response) Reset()         { *m = Ack_Response{} }
func (m *Ack_Response) String() string { return proto.CompactTextString(m) }
func (*Ack_Response) ProtoMessage()    {}
func (*Ack_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3, 0}
}

func (m *Ack_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ack_Response.Unmarshal(m, b)
}
func (m *Ack_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ack_Response.Marshal(b, m, deterministic)
}
func (m *Ack_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack_Response.Merge(m, src)
}
func (m *Ack_Response) XXX_Size() int {
	return xxx_messageInfo_Ack_Response.Size(m)
}
func (m *Ack_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Ack_Response proto.InternalMessageInfo

func (m *Ack_Response) GetSessionID() []byte {
	if m != nil {
		return m.SessionID
	}
	return nil
}

func (m *Ack_Response) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Ack_Response) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Ack_Response) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Done struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Done) Reset()         { *m = Done{} }
func (m *Done) String() string { return proto.CompactTextString(m) }
func (*Done) ProtoMessage()    {}
func (*Done) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}

func (m *Done) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Done.Unmarshal(m, b)
}
func (m *Done) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Done.Marshal(b, m, deterministic)
}
func (m *Done) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Done.Merge(m, src)
}
func (m *Done) XXX_Size() int {
	return xxx_messageInfo_Done.Size(m)
}
func (m *Done) XXX_DiscardUnknown() {
	xxx_messageInfo_Done.DiscardUnknown(m)
}

var xxx_messageInfo_Done proto.InternalMessageInfo

func (m *Done) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func init() {
	proto.RegisterType((*Init)(nil), "dkg.Init")
	proto.RegisterType((*EncryptedDeal)(nil), "dkg.EncryptedDeal")
	proto.RegisterType((*Deal)(nil), "dkg.Deal")
	proto.RegisterType((*Ack)(nil), "dkg.Ack")
	proto.RegisterType((*Ack_Response)(nil), "dkg.Ack.Response")
	proto.RegisterType((*Done)(nil), "dkg.Done")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xe1, 0x4a, 0xfb, 0x30,
	0x14, 0xc5, 0xe9, 0xda, 0xff, 0xd8, 0xee, 0xfe, 0x13, 0x8c, 0x32, 0xca, 0xf0, 0xc3, 0x28, 0x22,
	0x03, 0xb1, 0xc2, 0xf6, 0x04, 0xc3, 0x0a, 0x0e, 0xc1, 0x0f, 0xf1, 0x09, 0xba, 0xe6, 0x52, 0x43,
	0x67, 0x52, 0x73, 0x33, 0x70, 0x6f, 0xe7, 0xa3, 0xc9, 0xd2, 0x6e, 0xb1, 0x82, 0x1f, 0xcf, 0x39,
	0xb9, 0xfc, 0xce, 0xcd, 0x85, 0xb3, 0x77, 0x24, 0xca, 0x4b, 0xa4, 0xb4, 0x36, 0xda, 0x6a, 0x16,
	0x8a, 0xaa, 0x9c, 0x5e, 0x68, 0x85, 0xf6, 0xbe, 0x9b, 0x24, 0x4b, 0x88, 0xd6, 0x4a, 0x5a, 0x76,
	0x0b, 0xc3, 0x5c, 0x08, 0x83, 0x44, 0x48, 0x71, 0x30, 0x0b, 0xe7, 0xa3, 0xc5, 0x38, 0x3d, 0x0c,
	0xa4, 0xab, 0xc6, 0xe6, 0x3e, 0x4f, 0x3e, 0x60, 0xfc, 0xa8, 0x0a, 0xb3, 0xaf, 0x2d, 0x8a, 0x0c,
	0xf3, 0x2d, 0xbb, 0x84, 0x7f, 0xd9, 0xd3, 0x33, 0xee, 0xe3, 0x60, 0x16, 0xcc, 0xff, 0xf3, 0x46,
	0xb0, 0x2b, 0x18, 0xbe, 0xca, 0x52, 0xe5, 0x76, 0x67, 0x30, 0xee, 0xb9, 0xc4, 0x1b, 0x87, 0x99,
	0x17, 0xad, 0x0a, 0x8c, 0xc3, 0x66, 0xc6, 0x09, 0x36, 0x81, 0xfe, 0x83, 0xac, 0xdf, 0xd0, 0xc4,
	0x91, 0xb3, 0x5b, 0x95, 0x6c, 0x20, 0x3a, 0x92, 0xa4, 0x12, 0xf8, 0xe9, 0x48, 0x63, 0xde, 0x08,
	0x76, 0x03, 0x91, 0xc0, 0x7c, 0xeb, 0x20, 0xa3, 0x05, 0x4b, 0x45, 0x55, 0xa6, 0x9d, 0x86, 0xdc,
	0xe5, 0x87, 0x46, 0x74, 0x6a, 0xd4, 0x70, 0xbd, 0x91, 0x7c, 0x05, 0x10, 0xae, 0x8a, 0xea, 0x0f,
	0xc6, 0x1d, 0x0c, 0x0c, 0x52, 0xad, 0x15, 0x61, 0xcb, 0x39, 0x77, 0x9c, 0x55, 0x51, 0xa5, 0xbc,
	0x0d, 0xf8, 0xe9, 0xc9, 0xd4, 0xc2, 0xe0, 0xe8, 0x3a, 0x2c, 0x12, 0x49, 0xad, 0xd6, 0x59, 0xfb,
	0x45, 0xde, 0xf0, 0xb8, 0xde, 0x4f, 0xdc, 0x04, 0xfa, 0x64, 0x73, 0xbb, 0x23, 0xd7, 0x73, 0xc0,
	0x5b, 0xd5, 0x5d, 0x21, 0xfa, 0xbd, 0xc2, 0x35, 0x44, 0x99, 0x56, 0x8e, 0x58, 0xef, 0x36, 0x5b,
	0x59, 0xf8, 0xa3, 0x78, 0x63, 0xd3, 0x77, 0xb7, 0x5f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4a,
	0x5f, 0x9e, 0x11, 0x27, 0x02, 0x00, 0x00,
}
