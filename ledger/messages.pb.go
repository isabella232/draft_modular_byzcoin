// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package ledger

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type TransactionInput struct {
	ContractID           string   `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Body                 *any.Any `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionInput) Reset()         { *m = TransactionInput{} }
func (m *TransactionInput) String() string { return proto.CompactTextString(m) }
func (*TransactionInput) ProtoMessage()    {}
func (*TransactionInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *TransactionInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionInput.Unmarshal(m, b)
}
func (m *TransactionInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionInput.Marshal(b, m, deterministic)
}
func (m *TransactionInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInput.Merge(m, src)
}
func (m *TransactionInput) XXX_Size() int {
	return xxx_messageInfo_TransactionInput.Size(m)
}
func (m *TransactionInput) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInput.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInput proto.InternalMessageInfo

func (m *TransactionInput) GetContractID() string {
	if m != nil {
		return m.ContractID
	}
	return ""
}

func (m *TransactionInput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *TransactionInput) GetBody() *any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

type TransactionResult struct {
	Transaction          *TransactionInput `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Accepted             bool              `protobuf:"varint,2,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Instances            [][]byte          `protobuf:"bytes,3,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TransactionResult) Reset()         { *m = TransactionResult{} }
func (m *TransactionResult) String() string { return proto.CompactTextString(m) }
func (*TransactionResult) ProtoMessage()    {}
func (*TransactionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *TransactionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResult.Unmarshal(m, b)
}
func (m *TransactionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResult.Marshal(b, m, deterministic)
}
func (m *TransactionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResult.Merge(m, src)
}
func (m *TransactionResult) XXX_Size() int {
	return xxx_messageInfo_TransactionResult.Size(m)
}
func (m *TransactionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResult.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResult proto.InternalMessageInfo

func (m *TransactionResult) GetTransaction() *TransactionInput {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *TransactionResult) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *TransactionResult) GetInstances() [][]byte {
	if m != nil {
		return m.Instances
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionInput)(nil), "ledger.TransactionInput")
	proto.RegisterType((*TransactionResult)(nil), "ledger.TransactionResult")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x95, 0xb2, 0x3b, 0x15, 0xd1, 0x20, 0x12, 0x17, 0x91, 0xb2, 0xa7, 0x9e, 0xb2,
	0xb0, 0xde, 0xbc, 0x09, 0x5e, 0xf6, 0x1a, 0x7c, 0x81, 0x34, 0x1d, 0xc3, 0x42, 0x9d, 0x94, 0x64,
	0x7a, 0xe8, 0x23, 0xf8, 0xd6, 0x42, 0xa2, 0xb6, 0x78, 0x9c, 0x2f, 0x1f, 0x7c, 0x7f, 0xe0, 0xfa,
	0x13, 0x63, 0x34, 0x0e, 0xa3, 0x1a, 0x83, 0x67, 0x2f, 0xaa, 0x01, 0x7b, 0x87, 0x61, 0xf7, 0xe0,
	0xbc, 0x77, 0x03, 0x1e, 0x12, 0xed, 0xa6, 0x8f, 0x83, 0xa1, 0x39, 0x2b, 0x7b, 0x86, 0x9b, 0xf7,
	0x60, 0x28, 0x1a, 0xcb, 0x67, 0x4f, 0x27, 0x1a, 0x27, 0x16, 0x4f, 0x00, 0xd6, 0x13, 0x07, 0x63,
	0xf9, 0xf4, 0x26, 0x8b, 0xa6, 0x68, 0xb7, 0x7a, 0x45, 0xc4, 0x3d, 0x54, 0x59, 0x97, 0x17, 0xe9,
	0xed, 0xe7, 0x12, 0x2d, 0x5c, 0x76, 0xbe, 0x9f, 0x65, 0xd9, 0x14, 0x6d, 0x7d, 0xbc, 0x53, 0xb9,
	0xaa, 0x7e, 0xab, 0xea, 0x95, 0x66, 0x9d, 0x8c, 0xfd, 0x57, 0x01, 0xb7, 0xab, 0xac, 0xc6, 0x38,
	0x0d, 0x2c, 0x5e, 0xa0, 0xe6, 0x05, 0xa6, 0x70, 0x7d, 0x94, 0x2a, 0x7f, 0x42, 0xfd, 0x9f, 0xa9,
	0xd7, 0xb2, 0xd8, 0xc1, 0xc6, 0x58, 0x8b, 0x23, 0x63, 0x9f, 0x56, 0x6d, 0xf4, 0xdf, 0x2d, 0x1e,
	0x61, 0x7b, 0xa6, 0xc8, 0x86, 0x2c, 0x46, 0x59, 0x36, 0x65, 0x7b, 0xa5, 0x17, 0xd0, 0x55, 0x69,
	0xdf, 0xf3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xaa, 0xea, 0xa1, 0x3d, 0x01, 0x00, 0x00,
}
