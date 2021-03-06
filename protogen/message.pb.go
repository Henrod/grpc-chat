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
	CreatedTime          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Title                string               `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Mentions             []string             `protobuf:"bytes,4,rep,name=mentions,proto3" json:"mentions,omitempty"`
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

func (m *Message) GetCreatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *Message) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Message) GetMentions() []string {
	if m != nil {
		return m.Mentions
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "proto.v1.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x7a, 0x65, 0x86,
	0x52, 0xf2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x81, 0xa4, 0xd2, 0x34, 0xfd, 0x92,
	0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x88, 0x52, 0xa5, 0x3e, 0x46, 0x2e, 0x76, 0x5f,
	0x88, 0x66, 0x21, 0x21, 0x2e, 0x96, 0xa4, 0xfc, 0x94, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x30, 0x5b, 0xc8, 0x96, 0x8b, 0x27, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x35, 0x25, 0x1e, 0xa4,
	0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4a, 0x0f, 0x62, 0xae, 0x1e, 0xcc, 0x5c, 0xbd,
	0x10, 0x98, 0xb9, 0x41, 0xdc, 0x50, 0xf5, 0x20, 0x11, 0x21, 0x11, 0x2e, 0xd6, 0x92, 0xcc, 0x92,
	0x9c, 0x54, 0x09, 0x66, 0xb0, 0x99, 0x10, 0x8e, 0x90, 0x14, 0x17, 0x47, 0x6e, 0x6a, 0x5e, 0x49,
	0x66, 0x7e, 0x5e, 0xb1, 0x04, 0x8b, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x9c, 0xef, 0xe4, 0xcc, 0xc5,
	0x93, 0x9c, 0x9f, 0xab, 0x07, 0xf3, 0x81, 0x13, 0x0f, 0xd4, 0x75, 0x01, 0x20, 0x81, 0x00, 0xc6,
	0x28, 0x76, 0xb0, 0x4c, 0x99, 0xe1, 0x22, 0x26, 0xe6, 0x80, 0x88, 0x88, 0x55, 0x4c, 0x1c, 0x60,
	0x09, 0xbd, 0x30, 0xc3, 0x53, 0x50, 0x66, 0x4c, 0x98, 0x61, 0x12, 0x1b, 0x58, 0x91, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xff, 0xec, 0x1d, 0xff, 0x18, 0x01, 0x00, 0x00,
}
