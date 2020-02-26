// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mascot.proto

package com_jsample_mascot

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

type MascotRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MascotRequest) Reset()         { *m = MascotRequest{} }
func (m *MascotRequest) String() string { return proto.CompactTextString(m) }
func (*MascotRequest) ProtoMessage()    {}
func (*MascotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cefc233ff737dc49, []int{0}
}

func (m *MascotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MascotRequest.Unmarshal(m, b)
}
func (m *MascotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MascotRequest.Marshal(b, m, deterministic)
}
func (m *MascotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MascotRequest.Merge(m, src)
}
func (m *MascotRequest) XXX_Size() int {
	return xxx_messageInfo_MascotRequest.Size(m)
}
func (m *MascotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MascotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MascotRequest proto.InternalMessageInfo

func (m *MascotRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MascotResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MascotResponse) Reset()         { *m = MascotResponse{} }
func (m *MascotResponse) String() string { return proto.CompactTextString(m) }
func (*MascotResponse) ProtoMessage()    {}
func (*MascotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cefc233ff737dc49, []int{1}
}

func (m *MascotResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MascotResponse.Unmarshal(m, b)
}
func (m *MascotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MascotResponse.Marshal(b, m, deterministic)
}
func (m *MascotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MascotResponse.Merge(m, src)
}
func (m *MascotResponse) XXX_Size() int {
	return xxx_messageInfo_MascotResponse.Size(m)
}
func (m *MascotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MascotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MascotResponse proto.InternalMessageInfo

func (m *MascotResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MascotResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MascotResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MascotRequest)(nil), "com.jsample.mascot.MascotRequest")
	proto.RegisterType((*MascotResponse)(nil), "com.jsample.mascot.MascotResponse")
}

func init() { proto.RegisterFile("mascot.proto", fileDescriptor_cefc233ff737dc49) }

var fileDescriptor_cefc233ff737dc49 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4d, 0x2c, 0x4e,
	0xce, 0x2f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0xce, 0xcf, 0xd5, 0xcb, 0x2a,
	0x4e, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x83, 0xc8, 0x28, 0xc9, 0x73, 0xf1, 0xfa, 0x82, 0x59, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xcc, 0x41, 0x4c, 0x99, 0x29, 0x4a, 0x7e, 0x5c, 0x7c, 0x30, 0x05, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0xe8, 0x2a, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18,
	0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54,
	0x09, 0x66, 0xb0, 0x30, 0x8c, 0x6b, 0x14, 0x0f, 0xb3, 0x30, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39,
	0x55, 0xc8, 0x8f, 0x8b, 0x35, 0x23, 0x35, 0x27, 0x27, 0x5f, 0x48, 0x51, 0x0f, 0xd3, 0x7d, 0x7a,
	0x28, 0x8e, 0x93, 0x52, 0xc2, 0xa7, 0x04, 0xe2, 0x3c, 0x27, 0x31, 0x2e, 0x2c, 0xfe, 0x0c, 0x60,
	0x4c, 0x62, 0x03, 0x07, 0x82, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x45, 0xdc, 0x3b, 0x03, 0x14,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MascotServiceClient is the client API for MascotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MascotServiceClient interface {
	Hello(ctx context.Context, in *MascotRequest, opts ...grpc.CallOption) (*MascotResponse, error)
}

type mascotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMascotServiceClient(cc grpc.ClientConnInterface) MascotServiceClient {
	return &mascotServiceClient{cc}
}

func (c *mascotServiceClient) Hello(ctx context.Context, in *MascotRequest, opts ...grpc.CallOption) (*MascotResponse, error) {
	out := new(MascotResponse)
	err := c.cc.Invoke(ctx, "/com.jsample.mascot.MascotService/hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MascotServiceServer is the server API for MascotService service.
type MascotServiceServer interface {
	Hello(context.Context, *MascotRequest) (*MascotResponse, error)
}

// UnimplementedMascotServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMascotServiceServer struct {
}

func (*UnimplementedMascotServiceServer) Hello(ctx context.Context, req *MascotRequest) (*MascotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

func RegisterMascotServiceServer(s *grpc.Server, srv MascotServiceServer) {
	s.RegisterService(&_MascotService_serviceDesc, srv)
}

func _MascotService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MascotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MascotServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.jsample.mascot.MascotService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MascotServiceServer).Hello(ctx, req.(*MascotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MascotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.jsample.mascot.MascotService",
	HandlerType: (*MascotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "hello",
			Handler:    _MascotService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mascot.proto",
}