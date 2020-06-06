// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package protov1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Message struct {
	Body                 string               `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
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

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Message) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "proto.v1.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x7a, 0x65, 0x86,
	0x52, 0xf2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x81, 0xa4, 0xd2, 0x34, 0xfd, 0x92,
	0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x88, 0x52, 0xa5, 0x08, 0x2e, 0x76, 0x5f, 0x88,
	0x5e, 0x21, 0x21, 0x2e, 0x96, 0xa4, 0xfc, 0x94, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x30, 0x5b, 0xc8, 0x92, 0x8b, 0x2b, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x35, 0x25, 0x3e, 0xb1, 0x44,
	0x82, 0x49, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4a, 0x0f, 0x62, 0xa8, 0x1e, 0xcc, 0x50, 0xbd, 0x10,
	0x98, 0xa1, 0x41, 0x9c, 0x50, 0xd5, 0x8e, 0x25, 0x4e, 0xce, 0x5c, 0x3c, 0xc9, 0xf9, 0xb9, 0x7a,
	0x30, 0xa7, 0x38, 0xf1, 0x40, 0xed, 0x09, 0x00, 0x09, 0x04, 0x30, 0x46, 0xb1, 0x83, 0x65, 0xca,
	0x0c, 0x17, 0x31, 0x31, 0x07, 0x44, 0x44, 0xac, 0x62, 0xe2, 0x00, 0x4b, 0xe8, 0x85, 0x19, 0x9e,
	0x82, 0x32, 0x63, 0xc2, 0x0c, 0x93, 0xd8, 0xc0, 0x8a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x6a, 0x00, 0xb2, 0xdc, 0xe1, 0x00, 0x00, 0x00,
}