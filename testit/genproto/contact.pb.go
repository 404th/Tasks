// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contact.proto

package genproto

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

type Contact struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Age                  uint32   `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{0}
}

func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (m *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(m, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Contact) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contact) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type Contacts struct {
	Contacts             []*Contact `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Contacts) Reset()         { *m = Contacts{} }
func (m *Contacts) String() string { return proto.CompactTextString(m) }
func (*Contacts) ProtoMessage()    {}
func (*Contacts) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1}
}

func (m *Contacts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contacts.Unmarshal(m, b)
}
func (m *Contacts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contacts.Marshal(b, m, deterministic)
}
func (m *Contacts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contacts.Merge(m, src)
}
func (m *Contacts) XXX_Size() int {
	return xxx_messageInfo_Contacts.Size(m)
}
func (m *Contacts) XXX_DiscardUnknown() {
	xxx_messageInfo_Contacts.DiscardUnknown(m)
}

var xxx_messageInfo_Contacts proto.InternalMessageInfo

func (m *Contacts) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type ContactCreateReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  uint32   `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactCreateReq) Reset()         { *m = ContactCreateReq{} }
func (m *ContactCreateReq) String() string { return proto.CompactTextString(m) }
func (*ContactCreateReq) ProtoMessage()    {}
func (*ContactCreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{2}
}

func (m *ContactCreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactCreateReq.Unmarshal(m, b)
}
func (m *ContactCreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactCreateReq.Marshal(b, m, deterministic)
}
func (m *ContactCreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactCreateReq.Merge(m, src)
}
func (m *ContactCreateReq) XXX_Size() int {
	return xxx_messageInfo_ContactCreateReq.Size(m)
}
func (m *ContactCreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactCreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_ContactCreateReq proto.InternalMessageInfo

func (m *ContactCreateReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ContactCreateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContactCreateReq) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type WithID struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithID) Reset()         { *m = WithID{} }
func (m *WithID) String() string { return proto.CompactTextString(m) }
func (*WithID) ProtoMessage()    {}
func (*WithID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{3}
}

func (m *WithID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithID.Unmarshal(m, b)
}
func (m *WithID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithID.Marshal(b, m, deterministic)
}
func (m *WithID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithID.Merge(m, src)
}
func (m *WithID) XXX_Size() int {
	return xxx_messageInfo_WithID.Size(m)
}
func (m *WithID) XXX_DiscardUnknown() {
	xxx_messageInfo_WithID.DiscardUnknown(m)
}

var xxx_messageInfo_WithID proto.InternalMessageInfo

func (m *WithID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{4}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Contact)(nil), "genproto.Contact")
	proto.RegisterType((*Contacts)(nil), "genproto.Contacts")
	proto.RegisterType((*ContactCreateReq)(nil), "genproto.ContactCreateReq")
	proto.RegisterType((*WithID)(nil), "genproto.WithID")
	proto.RegisterType((*EmptyResponse)(nil), "genproto.EmptyResponse")
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor_a5036fff2565fb15) }

var fileDescriptor_a5036fff2565fb15 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xb3, 0x49, 0xdf, 0xbe, 0x71, 0x64, 0x6b, 0x9d, 0x8b, 0x21, 0xa7, 0xb0, 0xa7, 0x5c,
	0xec, 0xa1, 0x3d, 0x88, 0x20, 0x14, 0x69, 0xa5, 0x78, 0x5d, 0x15, 0x0f, 0x9e, 0x62, 0x32, 0xd4,
	0x40, 0x9a, 0xc4, 0xec, 0x2a, 0xf8, 0xe5, 0xfc, 0x6c, 0x62, 0xfe, 0x5a, 0xd7, 0xde, 0x66, 0x9f,
	0x79, 0x66, 0xe7, 0x37, 0x0f, 0xf0, 0xb8, 0xc8, 0x75, 0x14, 0xeb, 0x59, 0x59, 0x15, 0xba, 0x40,
	0x77, 0x4b, 0x79, 0x5d, 0x89, 0x27, 0xf8, 0xbf, 0x6a, 0x5a, 0x38, 0x01, 0x3b, 0x4d, 0x3c, 0x16,
	0xb0, 0xf0, 0x9f, 0xb4, 0xd3, 0x04, 0x7d, 0x70, 0xdf, 0x14, 0x55, 0x79, 0xb4, 0x23, 0xcf, 0x0e,
	0x58, 0x78, 0x24, 0xfb, 0x37, 0x22, 0x8c, 0x6a, 0xdd, 0xa9, 0xf5, 0xba, 0xc6, 0x29, 0x38, 0xd1,
	0x96, 0xbc, 0x51, 0xc0, 0x42, 0x2e, 0xbf, 0x4b, 0x71, 0x09, 0x6e, 0xfb, 0xb9, 0xc2, 0x73, 0x70,
	0x5b, 0x06, 0xe5, 0xb1, 0xc0, 0x09, 0x8f, 0xe7, 0xa7, 0xb3, 0x8e, 0x62, 0xd6, 0xba, 0x64, 0x6f,
	0x11, 0xf7, 0x30, 0x6d, 0xc5, 0x55, 0x45, 0x91, 0x26, 0x49, 0xaf, 0x7b, 0x40, 0xec, 0x00, 0x90,
	0x6d, 0x02, 0x39, 0x03, 0x90, 0x07, 0xe3, 0xc7, 0x54, 0xbf, 0xdc, 0xae, 0x7f, 0x1f, 0x2b, 0x4e,
	0x80, 0xdf, 0xec, 0x4a, 0xfd, 0x21, 0x49, 0x95, 0x45, 0xae, 0x68, 0xfe, 0x69, 0xc3, 0xa4, 0x25,
	0xb8, 0xa3, 0xea, 0x3d, 0x8d, 0x09, 0x97, 0xc0, 0x1b, 0x98, 0x2e, 0x31, 0xdf, 0xb8, 0xa0, 0x87,
	0xf5, 0xa7, 0x43, 0xaf, 0x59, 0x29, 0x2c, 0x5c, 0x00, 0x6c, 0x48, 0x77, 0xd3, 0x86, 0xc3, 0x37,
	0x13, 0x11, 0x16, 0x2e, 0x61, 0xb2, 0x21, 0x7d, 0x9d, 0x65, 0x7d, 0x94, 0x67, 0x83, 0x6d, 0x8f,
	0xd9, 0x47, 0x63, 0x5e, 0x09, 0x0b, 0x2f, 0x80, 0x3f, 0x94, 0xc9, 0x0f, 0x6c, 0x73, 0xcd, 0xdf,
	0x9b, 0xaf, 0x80, 0xaf, 0x29, 0xa3, 0x61, 0xd0, 0x24, 0x3e, 0x84, 0x22, 0xac, 0xe7, 0x71, 0x2d,
	0x2f, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x8b, 0x76, 0xd6, 0x7b, 0x02, 0x00, 0x00,
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
	CreateContact(ctx context.Context, in *ContactCreateReq, opts ...grpc.CallOption) (*WithID, error)
	GetContact(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Contact, error)
	GetAllContacts(ctx context.Context, in *EmptyResponse, opts ...grpc.CallOption) (*Contacts, error)
	UpdateContact(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*Contact, error)
	DeleteContact(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type contactServiceClient struct {
	cc *grpc.ClientConn
}

func NewContactServiceClient(cc *grpc.ClientConn) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) CreateContact(ctx context.Context, in *ContactCreateReq, opts ...grpc.CallOption) (*WithID, error) {
	out := new(WithID)
	err := c.cc.Invoke(ctx, "/genproto.ContactService/CreateContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) GetContact(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := c.cc.Invoke(ctx, "/genproto.ContactService/GetContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) GetAllContacts(ctx context.Context, in *EmptyResponse, opts ...grpc.CallOption) (*Contacts, error) {
	out := new(Contacts)
	err := c.cc.Invoke(ctx, "/genproto.ContactService/GetAllContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) UpdateContact(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := c.cc.Invoke(ctx, "/genproto.ContactService/UpdateContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) DeleteContact(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/genproto.ContactService/DeleteContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServiceServer is the server API for ContactService service.
type ContactServiceServer interface {
	CreateContact(context.Context, *ContactCreateReq) (*WithID, error)
	GetContact(context.Context, *WithID) (*Contact, error)
	GetAllContacts(context.Context, *EmptyResponse) (*Contacts, error)
	UpdateContact(context.Context, *Contact) (*Contact, error)
	DeleteContact(context.Context, *WithID) (*EmptyResponse, error)
}

// UnimplementedContactServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContactServiceServer struct {
}

func (*UnimplementedContactServiceServer) CreateContact(ctx context.Context, req *ContactCreateReq) (*WithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContact not implemented")
}
func (*UnimplementedContactServiceServer) GetContact(ctx context.Context, req *WithID) (*Contact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContact not implemented")
}
func (*UnimplementedContactServiceServer) GetAllContacts(ctx context.Context, req *EmptyResponse) (*Contacts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllContacts not implemented")
}
func (*UnimplementedContactServiceServer) UpdateContact(ctx context.Context, req *Contact) (*Contact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContact not implemented")
}
func (*UnimplementedContactServiceServer) DeleteContact(ctx context.Context, req *WithID) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContact not implemented")
}

func RegisterContactServiceServer(s *grpc.Server, srv ContactServiceServer) {
	s.RegisterService(&_ContactService_serviceDesc, srv)
}

func _ContactService_CreateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContactCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).CreateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ContactService/CreateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).CreateContact(ctx, req.(*ContactCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_GetContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ContactService/GetContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetContact(ctx, req.(*WithID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_GetAllContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetAllContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ContactService/GetAllContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetAllContacts(ctx, req.(*EmptyResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_UpdateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).UpdateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ContactService/UpdateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).UpdateContact(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_DeleteContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).DeleteContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ContactService/DeleteContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).DeleteContact(ctx, req.(*WithID))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContact",
			Handler:    _ContactService_CreateContact_Handler,
		},
		{
			MethodName: "GetContact",
			Handler:    _ContactService_GetContact_Handler,
		},
		{
			MethodName: "GetAllContacts",
			Handler:    _ContactService_GetAllContacts_Handler,
		},
		{
			MethodName: "UpdateContact",
			Handler:    _ContactService_UpdateContact_Handler,
		},
		{
			MethodName: "DeleteContact",
			Handler:    _ContactService_DeleteContact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact.proto",
}