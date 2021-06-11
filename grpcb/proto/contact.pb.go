// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contact.proto

package proto

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

type Req struct {
	A                    int32    `protobuf:"varint,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{0}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Req) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type Resp struct {
	Res                  int32    `protobuf:"varint,1,opt,name=Res,proto3" json:"Res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetRes() int32 {
	if m != nil {
		return m.Res
	}
	return 0
}

func init() {
	proto.RegisterType((*Req)(nil), "proto.Req")
	proto.RegisterType((*Resp)(nil), "proto.Resp")
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor_a5036fff2565fb15) }

var fileDescriptor_a5036fff2565fb15 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0x2b,
	0x49, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x8a, 0x5c,
	0xcc, 0x41, 0xa9, 0x85, 0x42, 0x3c, 0x5c, 0x8c, 0x8e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41,
	0x8c, 0x8e, 0x20, 0x9e, 0x93, 0x04, 0x13, 0x84, 0xe7, 0xa4, 0x24, 0xc1, 0xc5, 0x12, 0x94, 0x5a,
	0x5c, 0x20, 0x24, 0x00, 0x52, 0x5a, 0x0c, 0x55, 0x05, 0x62, 0x1a, 0x4d, 0x62, 0xe4, 0xe2, 0x73,
	0x86, 0x98, 0x1a, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x24, 0xcf, 0xc5, 0x12, 0x90, 0x53,
	0x5a, 0x2c, 0xc4, 0x05, 0xb1, 0x46, 0x2f, 0x28, 0xb5, 0x50, 0x8a, 0x1b, 0xce, 0x2e, 0x2e, 0x00,
	0x29, 0xf0, 0x2d, 0xcd, 0x29, 0xc1, 0xad, 0x40, 0x8e, 0x8b, 0xd9, 0x25, 0xb3, 0x0c, 0xaf, 0xbc,
	0x6f, 0x66, 0x1e, 0x4e, 0xf9, 0x24, 0x36, 0x30, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x34,
	0x5a, 0xf6, 0x7b, 0xf0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContactServiceClient is the client API for ContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContactServiceClient interface {
	Plus(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	Mult(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	Div(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	Min(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
}

type contactServiceClient struct {
	cc *grpc.ClientConn
}

func NewContactServiceClient(cc *grpc.ClientConn) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) Plus(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/proto.ContactService/Plus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Mult(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/proto.ContactService/Mult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Div(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/proto.ContactService/Div", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Min(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/proto.ContactService/Min", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServiceServer is the server API for ContactService service.
type ContactServiceServer interface {
	Plus(context.Context, *Req) (*Resp, error)
	Mult(context.Context, *Req) (*Resp, error)
	Div(context.Context, *Req) (*Resp, error)
	Min(context.Context, *Req) (*Resp, error)
}

// UnimplementedContactServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContactServiceServer struct {
}

func (*UnimplementedContactServiceServer) Plus(ctx context.Context, req *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Plus not implemented")
}
func (*UnimplementedContactServiceServer) Mult(ctx context.Context, req *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mult not implemented")
}
func (*UnimplementedContactServiceServer) Div(ctx context.Context, req *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Div not implemented")
}
func (*UnimplementedContactServiceServer) Min(ctx context.Context, req *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Min not implemented")
}

func RegisterContactServiceServer(s *grpc.Server, srv ContactServiceServer) {
	s.RegisterService(&_ContactService_serviceDesc, srv)
}

func _ContactService_Plus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Plus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ContactService/Plus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Plus(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Mult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Mult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ContactService/Mult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Mult(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ContactService/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Div(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Min_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Min(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ContactService/Min",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Min(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Plus",
			Handler:    _ContactService_Plus_Handler,
		},
		{
			MethodName: "Mult",
			Handler:    _ContactService_Mult_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _ContactService_Div_Handler,
		},
		{
			MethodName: "Min",
			Handler:    _ContactService_Min_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact.proto",
}