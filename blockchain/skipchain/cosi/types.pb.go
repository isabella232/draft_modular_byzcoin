// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package cosi

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

type SignatureRequest struct {
	Message              *any.Any `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureRequest) Reset()         { *m = SignatureRequest{} }
func (m *SignatureRequest) String() string { return proto.CompactTextString(m) }
func (*SignatureRequest) ProtoMessage()    {}
func (*SignatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *SignatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureRequest.Unmarshal(m, b)
}
func (m *SignatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureRequest.Marshal(b, m, deterministic)
}
func (m *SignatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureRequest.Merge(m, src)
}
func (m *SignatureRequest) XXX_Size() int {
	return xxx_messageInfo_SignatureRequest.Size(m)
}
func (m *SignatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureRequest proto.InternalMessageInfo

func (m *SignatureRequest) GetMessage() *any.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

type SignatureResponse struct {
	Signature            *any.Any `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureResponse) Reset()         { *m = SignatureResponse{} }
func (m *SignatureResponse) String() string { return proto.CompactTextString(m) }
func (*SignatureResponse) ProtoMessage()    {}
func (*SignatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

func (m *SignatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureResponse.Unmarshal(m, b)
}
func (m *SignatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureResponse.Marshal(b, m, deterministic)
}
func (m *SignatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureResponse.Merge(m, src)
}
func (m *SignatureResponse) XXX_Size() int {
	return xxx_messageInfo_SignatureResponse.Size(m)
}
func (m *SignatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureResponse proto.InternalMessageInfo

func (m *SignatureResponse) GetSignature() *any.Any {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*SignatureRequest)(nil), "cosi.SignatureRequest")
	proto.RegisterType((*SignatureResponse)(nil), "cosi.SignatureResponse")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0xce, 0x2f, 0xce, 0x94, 0x92, 0x4c,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6, 0xe9, 0x27, 0xe6, 0x55, 0x42,
	0x14, 0x28, 0x39, 0x71, 0x09, 0x04, 0x67, 0xa6, 0xe7, 0x25, 0x96, 0x94, 0x16, 0xa5, 0x06, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xe9, 0x71, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xe8, 0x41, 0x0c, 0xd0, 0x83, 0x19, 0xa0, 0xe7,
	0x98, 0x57, 0x19, 0x04, 0x53, 0xa4, 0xe4, 0xce, 0x25, 0x88, 0x64, 0x46, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0xaa, 0x90, 0x11, 0x17, 0x67, 0x31, 0x4c, 0x10, 0xaf, 0x31, 0x08, 0x65, 0x49, 0x6c, 0x60,
	0x09, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x7c, 0x0c, 0x88, 0xc3, 0x00, 0x00, 0x00,
}
