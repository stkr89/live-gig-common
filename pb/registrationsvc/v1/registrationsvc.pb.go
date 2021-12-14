// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/registrationsvc/v1/registrationsvc.proto

package v1

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

// register request
type RegisterRequest struct {
	// first name of user
	FirstName string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	// last name of user
	LastName string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	// email address of user
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// password of user
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// confirm password
	ConfirmPassword      string   `protobuf:"bytes,5,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ad5f245eaa35e9d, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

// register response
type RegisterResponse struct {
	// unique id of user
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string       `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string       `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string       `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Errors               []*ErrorType `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ad5f245eaa35e9d, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisterResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegisterResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterResponse) GetErrors() []*ErrorType {
	if m != nil {
		return m.Errors
	}
	return nil
}

// error type
type ErrorType struct {
	// http error code
	StatusCode int32 `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	// error detail
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorType) Reset()         { *m = ErrorType{} }
func (m *ErrorType) String() string { return proto.CompactTextString(m) }
func (*ErrorType) ProtoMessage()    {}
func (*ErrorType) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ad5f245eaa35e9d, []int{2}
}

func (m *ErrorType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorType.Unmarshal(m, b)
}
func (m *ErrorType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorType.Marshal(b, m, deterministic)
}
func (m *ErrorType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorType.Merge(m, src)
}
func (m *ErrorType) XXX_Size() int {
	return xxx_messageInfo_ErrorType.Size(m)
}
func (m *ErrorType) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorType.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorType proto.InternalMessageInfo

func (m *ErrorType) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *ErrorType) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "pb.registrationsvc.v1.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "pb.registrationsvc.v1.RegisterResponse")
	proto.RegisterType((*ErrorType)(nil), "pb.registrationsvc.v1.ErrorType")
}

func init() {
	proto.RegisterFile("pb/registrationsvc/v1/registrationsvc.proto", fileDescriptor_5ad5f245eaa35e9d)
}

var fileDescriptor_5ad5f245eaa35e9d = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0xf3, 0x30,
	0x1c, 0xc7, 0x9f, 0x76, 0xeb, 0x9e, 0xf5, 0x27, 0x38, 0x09, 0x0e, 0xcb, 0x10, 0x29, 0x3d, 0x68,
	0x41, 0xe8, 0xd8, 0xbc, 0x78, 0x56, 0x76, 0x15, 0xa9, 0x9e, 0x04, 0x0f, 0x69, 0x9b, 0x8d, 0xc0,
	0xda, 0xc4, 0xfc, 0xb2, 0x8a, 0xaf, 0xc8, 0x8b, 0x2f, 0x52, 0xd6, 0x36, 0xdb, 0x2c, 0x9b, 0x1e,
	0xbf, 0x7f, 0x42, 0x3e, 0x7c, 0x13, 0xb8, 0x96, 0xc9, 0x58, 0xb1, 0x05, 0x47, 0xad, 0xa8, 0xe6,
	0xa2, 0xc0, 0x32, 0x1d, 0x97, 0x93, 0xb6, 0x15, 0x49, 0x25, 0xb4, 0x20, 0x43, 0x99, 0x44, 0xed,
	0xa4, 0x9c, 0x04, 0x9f, 0x16, 0x0c, 0xe2, 0xca, 0x66, 0x2a, 0x66, 0x6f, 0x2b, 0x86, 0x9a, 0x9c,
	0x83, 0x3b, 0xe7, 0x0a, 0xf5, 0x03, 0xcd, 0x99, 0x67, 0xf9, 0x56, 0xe8, 0xc6, 0x5b, 0x83, 0x8c,
	0xa0, 0xbf, 0xa4, 0x4d, 0x68, 0x57, 0xe1, 0x46, 0x93, 0x53, 0x70, 0x58, 0x4e, 0xf9, 0xd2, 0xeb,
	0x54, 0x41, 0x2d, 0xd6, 0x27, 0x24, 0x45, 0x7c, 0x17, 0x2a, 0xf3, 0xba, 0xf5, 0x09, 0xa3, 0x49,
	0x08, 0x83, 0x54, 0x14, 0x73, 0xae, 0xf2, 0x47, 0x53, 0x71, 0xaa, 0x4a, 0xdb, 0x0e, 0xbe, 0x2c,
	0x38, 0xd9, 0x92, 0xa2, 0x14, 0x05, 0x32, 0x72, 0x0c, 0x36, 0xcf, 0x1a, 0x46, 0x9b, 0x67, 0x3f,
	0xd1, 0xed, 0xdf, 0xd0, 0x3b, 0x87, 0xd0, 0xbb, 0xbb, 0xe8, 0xb7, 0xd0, 0x63, 0x4a, 0x09, 0x85,
	0x9e, 0xe3, 0x77, 0xc2, 0xa3, 0xa9, 0x1f, 0xed, 0x9d, 0x31, 0x9a, 0xad, 0x4b, 0xcf, 0x1f, 0x92,
	0xc5, 0x4d, 0x3f, 0x98, 0x81, 0xbb, 0x31, 0xc9, 0x05, 0x00, 0x6a, 0xaa, 0x57, 0x78, 0x2f, 0xb2,
	0x7a, 0x52, 0x27, 0xde, 0x71, 0x88, 0x07, 0xff, 0x73, 0x86, 0x48, 0x17, 0x06, 0xda, 0xc8, 0xa9,
	0x34, 0xcf, 0x53, 0x5f, 0xf7, 0x54, 0xa6, 0xe4, 0x15, 0xfa, 0x66, 0x07, 0x72, 0x79, 0x80, 0xa7,
	0xf5, 0xa4, 0xa3, 0xab, 0x3f, 0x7b, 0xf5, 0xa0, 0xc1, 0xbf, 0xbb, 0xb3, 0x97, 0xe1, 0xde, 0x7f,
	0x95, 0xf4, 0xaa, 0x8f, 0x74, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xce, 0xd6, 0xf2, 0x5b, 0x77,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistrationSvcClient is the client API for RegistrationSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistrationSvcClient interface {
	// register new user
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type registrationSvcClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationSvcClient(cc *grpc.ClientConn) RegistrationSvcClient {
	return &registrationSvcClient{cc}
}

func (c *registrationSvcClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/pb.registrationsvc.v1.RegistrationSvc/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationSvcServer is the server API for RegistrationSvc service.
type RegistrationSvcServer interface {
	// register new user
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

// UnimplementedRegistrationSvcServer can be embedded to have forward compatible implementations.
type UnimplementedRegistrationSvcServer struct {
}

func (*UnimplementedRegistrationSvcServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterRegistrationSvcServer(s *grpc.Server, srv RegistrationSvcServer) {
	s.RegisterService(&_RegistrationSvc_serviceDesc, srv)
}

func _RegistrationSvc_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationSvcServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.registrationsvc.v1.RegistrationSvc/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationSvcServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistrationSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.registrationsvc.v1.RegistrationSvc",
	HandlerType: (*RegistrationSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RegistrationSvc_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/registrationsvc/v1/registrationsvc.proto",
}
