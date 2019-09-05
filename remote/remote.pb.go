// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote.proto

package remote

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

type AddPasswordReq struct {
	Tag                   string         `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Username              string         `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	OverrideConfiguration *Configuration `protobuf:"bytes,3,opt,name=overrideConfiguration,proto3" json:"overrideConfiguration,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}       `json:"-"`
	XXX_unrecognized      []byte         `json:"-"`
	XXX_sizecache         int32          `json:"-"`
}

func (m *AddPasswordReq) Reset()         { *m = AddPasswordReq{} }
func (m *AddPasswordReq) String() string { return proto.CompactTextString(m) }
func (*AddPasswordReq) ProtoMessage()    {}
func (*AddPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{0}
}

func (m *AddPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPasswordReq.Unmarshal(m, b)
}
func (m *AddPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPasswordReq.Marshal(b, m, deterministic)
}
func (m *AddPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPasswordReq.Merge(m, src)
}
func (m *AddPasswordReq) XXX_Size() int {
	return xxx_messageInfo_AddPasswordReq.Size(m)
}
func (m *AddPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddPasswordReq proto.InternalMessageInfo

func (m *AddPasswordReq) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *AddPasswordReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddPasswordReq) GetOverrideConfiguration() *Configuration {
	if m != nil {
		return m.OverrideConfiguration
	}
	return nil
}

type Configuration struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Seed                 string   `protobuf:"bytes,2,opt,name=seed,proto3" json:"seed,omitempty"`
	Factor               int32    `protobuf:"varint,3,opt,name=factor,proto3" json:"factor,omitempty"`
	Storage              string   `protobuf:"bytes,4,opt,name=storage,proto3" json:"storage,omitempty"`
	Output               bool     `protobuf:"varint,5,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{1}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Configuration) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *Configuration) GetFactor() int32 {
	if m != nil {
		return m.Factor
	}
	return 0
}

func (m *Configuration) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

func (m *Configuration) GetOutput() bool {
	if m != nil {
		return m.Output
	}
	return false
}

type GetPasswordReq struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPasswordReq) Reset()         { *m = GetPasswordReq{} }
func (m *GetPasswordReq) String() string { return proto.CompactTextString(m) }
func (*GetPasswordReq) ProtoMessage()    {}
func (*GetPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{2}
}

func (m *GetPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPasswordReq.Unmarshal(m, b)
}
func (m *GetPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPasswordReq.Marshal(b, m, deterministic)
}
func (m *GetPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPasswordReq.Merge(m, src)
}
func (m *GetPasswordReq) XXX_Size() int {
	return xxx_messageInfo_GetPasswordReq.Size(m)
}
func (m *GetPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPasswordReq proto.InternalMessageInfo

func (m *GetPasswordReq) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *GetPasswordReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type PasswordValue struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordValue) Reset()         { *m = PasswordValue{} }
func (m *PasswordValue) String() string { return proto.CompactTextString(m) }
func (*PasswordValue) ProtoMessage()    {}
func (*PasswordValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_eefc82927d57d89b, []int{3}
}

func (m *PasswordValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordValue.Unmarshal(m, b)
}
func (m *PasswordValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordValue.Marshal(b, m, deterministic)
}
func (m *PasswordValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordValue.Merge(m, src)
}
func (m *PasswordValue) XXX_Size() int {
	return xxx_messageInfo_PasswordValue.Size(m)
}
func (m *PasswordValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordValue.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordValue proto.InternalMessageInfo

func (m *PasswordValue) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*AddPasswordReq)(nil), "AddPasswordReq")
	proto.RegisterType((*Configuration)(nil), "Configuration")
	proto.RegisterType((*GetPasswordReq)(nil), "GetPasswordReq")
	proto.RegisterType((*PasswordValue)(nil), "PasswordValue")
}

func init() { proto.RegisterFile("remote.proto", fileDescriptor_eefc82927d57d89b) }

var fileDescriptor_eefc82927d57d89b = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xd9, 0x2f, 0x6d, 0xbf, 0x3a, 0xda, 0x44, 0x16, 0x94, 0xa5, 0xa7, 0x90, 0x43, 0x09,
	0x08, 0x39, 0xd4, 0xbb, 0x50, 0x14, 0x7a, 0x12, 0x24, 0x07, 0xef, 0xab, 0x3b, 0x8d, 0x01, 0x93,
	0x89, 0x93, 0x89, 0xfe, 0x01, 0xfd, 0xdf, 0x92, 0xb0, 0x29, 0x44, 0x3c, 0x79, 0x9b, 0xe7, 0xe5,
	0x99, 0xd9, 0x61, 0x07, 0xce, 0x18, 0x2b, 0x12, 0xcc, 0x1a, 0x26, 0xa1, 0xe4, 0x53, 0x41, 0xb8,
	0x73, 0xee, 0xc1, 0xb6, 0xed, 0x07, 0xb1, 0xcb, 0xf1, 0x4d, 0x9f, 0x43, 0x20, 0xb6, 0x30, 0x2a,
	0x56, 0xe9, 0x49, 0xde, 0x97, 0x7a, 0x0d, 0xcb, 0xae, 0x45, 0xae, 0x6d, 0x85, 0xe6, 0xdf, 0x10,
	0x1f, 0x59, 0xdf, 0xc1, 0x05, 0xbd, 0x23, 0x73, 0xe9, 0xf0, 0x96, 0xea, 0x43, 0x59, 0x74, 0x6c,
	0xa5, 0xa4, 0xda, 0x04, 0xb1, 0x4a, 0x4f, 0xb7, 0x61, 0x36, 0x49, 0xf3, 0xdf, 0xe5, 0xe4, 0x4b,
	0xc1, 0x6a, 0x92, 0xe8, 0x4b, 0x58, 0x54, 0x28, 0x2f, 0xe4, 0xfc, 0x22, 0x9e, 0xb4, 0x86, 0x59,
	0x8b, 0xe8, 0xfc, 0x1e, 0x43, 0xdd, 0xbb, 0x07, 0xfb, 0x2c, 0xc4, 0xc3, 0xa3, 0xf3, 0xdc, 0x93,
	0x36, 0xf0, 0xbf, 0x15, 0x62, 0x5b, 0xa0, 0x99, 0x0d, 0xfa, 0x88, 0x7d, 0x07, 0x75, 0xd2, 0x74,
	0x62, 0xe6, 0xb1, 0x4a, 0x97, 0xb9, 0xa7, 0xe4, 0x06, 0xc2, 0x3d, 0xca, 0x9f, 0x7f, 0x23, 0xb9,
	0x82, 0xd5, 0xd8, 0xfc, 0x68, 0x5f, 0x3b, 0xec, 0xe5, 0xc6, 0x07, 0x7e, 0xc6, 0x91, 0xb7, 0x16,
	0xa2, 0x51, 0xbe, 0xb7, 0xb5, 0x2d, 0x90, 0xf5, 0x06, 0x82, 0x9d, 0x73, 0x3a, 0xca, 0xa6, 0x37,
	0x59, 0x87, 0xd9, 0x74, 0xec, 0x06, 0x82, 0x3d, 0x8a, 0x8e, 0xb2, 0xe9, 0xb6, 0x3f, 0xbd, 0xa7,
	0xc5, 0x70, 0xe5, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x9d, 0x41, 0x1f, 0xf5, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PasswordManagerClient is the client API for PasswordManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PasswordManagerClient interface {
	Add(ctx context.Context, in *AddPasswordReq, opts ...grpc.CallOption) (*PasswordValue, error)
	Get(ctx context.Context, in *GetPasswordReq, opts ...grpc.CallOption) (*PasswordValue, error)
}

type passwordManagerClient struct {
	cc *grpc.ClientConn
}

func NewPasswordManagerClient(cc *grpc.ClientConn) PasswordManagerClient {
	return &passwordManagerClient{cc}
}

func (c *passwordManagerClient) Add(ctx context.Context, in *AddPasswordReq, opts ...grpc.CallOption) (*PasswordValue, error) {
	out := new(PasswordValue)
	err := c.cc.Invoke(ctx, "/PasswordManager/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordManagerClient) Get(ctx context.Context, in *GetPasswordReq, opts ...grpc.CallOption) (*PasswordValue, error) {
	out := new(PasswordValue)
	err := c.cc.Invoke(ctx, "/PasswordManager/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PasswordManagerServer is the server API for PasswordManager service.
type PasswordManagerServer interface {
	Add(context.Context, *AddPasswordReq) (*PasswordValue, error)
	Get(context.Context, *GetPasswordReq) (*PasswordValue, error)
}

// UnimplementedPasswordManagerServer can be embedded to have forward compatible implementations.
type UnimplementedPasswordManagerServer struct {
}

func (*UnimplementedPasswordManagerServer) Add(ctx context.Context, req *AddPasswordReq) (*PasswordValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedPasswordManagerServer) Get(ctx context.Context, req *GetPasswordReq) (*PasswordValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterPasswordManagerServer(s *grpc.Server, srv PasswordManagerServer) {
	s.RegisterService(&_PasswordManager_serviceDesc, srv)
}

func _PasswordManager_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordManagerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PasswordManager/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordManagerServer).Add(ctx, req.(*AddPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PasswordManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PasswordManager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordManagerServer).Get(ctx, req.(*GetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PasswordManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PasswordManager",
	HandlerType: (*PasswordManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _PasswordManager_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PasswordManager_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote.proto",
}
