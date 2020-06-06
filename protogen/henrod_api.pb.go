// Code generated by protoc-gen-go. DO NOT EDIT.
// source: henrod_api.proto

package protov1

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

type EchoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e117e34666d253, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

type EchoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e117e34666d253, []int{1}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EchoRequest)(nil), "proto.v1.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "proto.v1.EchoResponse")
}

func init() { proto.RegisterFile("henrod_api.proto", fileDescriptor_a2e117e34666d253) }

var fileDescriptor_a2e117e34666d253 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0x2b,
	0xca, 0x4f, 0x89, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53,
	0x7a, 0x65, 0x86, 0x4a, 0xbc, 0x5c, 0xdc, 0xae, 0xc9, 0x19, 0xf9, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x4a, 0x7c, 0x5c, 0x3c, 0x10, 0x6e, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x91, 0x0b,
	0x17, 0xa7, 0x07, 0x58, 0xb3, 0x63, 0x80, 0xa7, 0x90, 0x39, 0x17, 0x0b, 0x48, 0x52, 0x48, 0x54,
	0x0f, 0xa6, 0x5d, 0x0f, 0x49, 0xaf, 0x94, 0x18, 0xba, 0x30, 0xc4, 0x0c, 0x25, 0x06, 0x27, 0x57,
	0x2e, 0x9e, 0xe4, 0xfc, 0x5c, 0xb8, 0xb4, 0x13, 0x1f, 0xd4, 0xcc, 0x82, 0xcc, 0x00, 0x90, 0x50,
	0x00, 0x63, 0x14, 0x3b, 0x58, 0xae, 0xcc, 0x70, 0x11, 0x13, 0x73, 0x40, 0x44, 0xc4, 0x2a, 0x26,
	0x0e, 0xb0, 0x84, 0x5e, 0x98, 0xe1, 0x29, 0x28, 0x33, 0x26, 0xcc, 0x30, 0x89, 0x0d, 0xac, 0xc8,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xae, 0xef, 0xbd, 0x90, 0xd0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HenrodAPIClient is the client API for HenrodAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HenrodAPIClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type henrodAPIClient struct {
	cc *grpc.ClientConn
}

func NewHenrodAPIClient(cc *grpc.ClientConn) HenrodAPIClient {
	return &henrodAPIClient{cc}
}

func (c *henrodAPIClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/proto.v1.HenrodAPI/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HenrodAPIServer is the server API for HenrodAPI service.
type HenrodAPIServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

// UnimplementedHenrodAPIServer can be embedded to have forward compatible implementations.
type UnimplementedHenrodAPIServer struct {
}

func (*UnimplementedHenrodAPIServer) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterHenrodAPIServer(s *grpc.Server, srv HenrodAPIServer) {
	s.RegisterService(&_HenrodAPI_serviceDesc, srv)
}

func _HenrodAPI_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HenrodAPIServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.v1.HenrodAPI/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HenrodAPIServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HenrodAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.v1.HenrodAPI",
	HandlerType: (*HenrodAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _HenrodAPI_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "henrod_api.proto",
}
