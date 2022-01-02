// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/authsvc/v1/authsvc.proto

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

// sign in response
type SignInResponse struct {
	// access token
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInResponse) Reset()         { *m = SignInResponse{} }
func (m *SignInResponse) String() string { return proto.CompactTextString(m) }
func (*SignInResponse) ProtoMessage()    {}
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b7e1b8b39fd6913, []int{0}
}

func (m *SignInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResponse.Unmarshal(m, b)
}
func (m *SignInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResponse.Marshal(b, m, deterministic)
}
func (m *SignInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResponse.Merge(m, src)
}
func (m *SignInResponse) XXX_Size() int {
	return xxx_messageInfo_SignInResponse.Size(m)
}
func (m *SignInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResponse proto.InternalMessageInfo

func (m *SignInResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

// sign in request
type SignInRequest struct {
	// email address of user
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// password of user
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInRequest) Reset()         { *m = SignInRequest{} }
func (m *SignInRequest) String() string { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()    {}
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b7e1b8b39fd6913, []int{1}
}

func (m *SignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInRequest.Unmarshal(m, b)
}
func (m *SignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInRequest.Marshal(b, m, deterministic)
}
func (m *SignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInRequest.Merge(m, src)
}
func (m *SignInRequest) XXX_Size() int {
	return xxx_messageInfo_SignInRequest.Size(m)
}
func (m *SignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInRequest proto.InternalMessageInfo

func (m *SignInRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// sign up request
type SignUpRequest struct {
	// first name of user
	FirstName string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	// last name of user
	LastName string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	// email address of user
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// password of user
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// confirm password
	ConfirmPassword string `protobuf:"bytes,5,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
	// name of the org
	OrgName              string   `protobuf:"bytes,6,opt,name=orgName,proto3" json:"orgName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b7e1b8b39fd6913, []int{2}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *SignUpRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *SignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

func (m *SignUpRequest) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

// sign up response
type SignUpResponse struct {
	// aws cognito id of user
	AwsUserSub           string       `protobuf:"bytes,1,opt,name=awsUserSub,proto3" json:"awsUserSub,omitempty"`
	FirstName            string       `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string       `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string       `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	OrgName              string       `protobuf:"bytes,5,opt,name=orgName,proto3" json:"orgName,omitempty"`
	Errors               []*ErrorType `protobuf:"bytes,6,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SignUpResponse) Reset()         { *m = SignUpResponse{} }
func (m *SignUpResponse) String() string { return proto.CompactTextString(m) }
func (*SignUpResponse) ProtoMessage()    {}
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b7e1b8b39fd6913, []int{3}
}

func (m *SignUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResponse.Unmarshal(m, b)
}
func (m *SignUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResponse.Marshal(b, m, deterministic)
}
func (m *SignUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResponse.Merge(m, src)
}
func (m *SignUpResponse) XXX_Size() int {
	return xxx_messageInfo_SignUpResponse.Size(m)
}
func (m *SignUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResponse proto.InternalMessageInfo

func (m *SignUpResponse) GetAwsUserSub() string {
	if m != nil {
		return m.AwsUserSub
	}
	return ""
}

func (m *SignUpResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *SignUpResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *SignUpResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpResponse) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *SignUpResponse) GetErrors() []*ErrorType {
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
	return fileDescriptor_5b7e1b8b39fd6913, []int{4}
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
	proto.RegisterType((*SignInResponse)(nil), "pb.authsvc.v1.SignInResponse")
	proto.RegisterType((*SignInRequest)(nil), "pb.authsvc.v1.SignInRequest")
	proto.RegisterType((*SignUpRequest)(nil), "pb.authsvc.v1.SignUpRequest")
	proto.RegisterType((*SignUpResponse)(nil), "pb.authsvc.v1.SignUpResponse")
	proto.RegisterType((*ErrorType)(nil), "pb.authsvc.v1.ErrorType")
}

func init() { proto.RegisterFile("pb/authsvc/v1/authsvc.proto", fileDescriptor_5b7e1b8b39fd6913) }

var fileDescriptor_5b7e1b8b39fd6913 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xa6, 0x40, 0x8b, 0x0c, 0x41, 0x92, 0x8d, 0x87, 0x06, 0xd1, 0x90, 0x9e, 0x38, 0x15, 0xc1,
	0x27, 0x40, 0x43, 0x0c, 0x17, 0x63, 0x0a, 0x5c, 0xbc, 0x6d, 0xcb, 0x02, 0x8d, 0xb4, 0x5b, 0x77,
	0xda, 0x12, 0x9f, 0xc5, 0x87, 0xf1, 0x05, 0x7c, 0x28, 0xd3, 0x9f, 0x85, 0x96, 0x80, 0xb7, 0xce,
	0x37, 0x5f, 0xbf, 0x6f, 0xbf, 0x99, 0x5d, 0xb8, 0x0d, 0xec, 0x21, 0x8d, 0xc2, 0x2d, 0xc6, 0xce,
	0x30, 0x1e, 0xc9, 0x4f, 0x33, 0x10, 0x3c, 0xe4, 0xa4, 0x1d, 0xd8, 0xa6, 0x44, 0xe2, 0x91, 0x31,
	0x86, 0xeb, 0xb9, 0xbb, 0xf1, 0x67, 0xbe, 0xc5, 0x30, 0xe0, 0x3e, 0x32, 0xd2, 0x87, 0x16, 0x75,
	0x1c, 0x86, 0xb8, 0xe0, 0x1f, 0xcc, 0xd7, 0x95, 0xbe, 0x32, 0x68, 0x5a, 0x45, 0xc8, 0x98, 0x40,
	0x5b, 0xfe, 0xf3, 0x19, 0x31, 0x0c, 0xc9, 0x0d, 0xa8, 0xcc, 0xa3, 0xee, 0x2e, 0x27, 0x67, 0x05,
	0xe9, 0xc2, 0x55, 0x40, 0x11, 0xf7, 0x5c, 0xac, 0xf4, 0x6a, 0xda, 0x38, 0xd4, 0xc6, 0x8f, 0x92,
	0x69, 0x2c, 0x03, 0xa9, 0xd1, 0x83, 0xe6, 0xda, 0x15, 0x18, 0xbe, 0x52, 0x8f, 0xe5, 0x3a, 0x47,
	0x20, 0xd1, 0xda, 0xd1, 0xbc, 0x99, 0x6b, 0xc9, 0xfa, 0xe8, 0x5e, 0xbb, 0xe4, 0x5e, 0x2f, 0xbb,
	0x93, 0x01, 0x74, 0x1c, 0xee, 0xaf, 0x5d, 0xe1, 0xbd, 0x49, 0x8a, 0x9a, 0x52, 0x4e, 0x61, 0xa2,
	0x43, 0x83, 0x8b, 0x4d, 0x6a, 0xab, 0xa5, 0x0c, 0x59, 0x1a, 0xbf, 0x4a, 0x36, 0xb9, 0x24, 0x41,
	0x3e, 0xb9, 0x7b, 0x00, 0xba, 0xc7, 0x25, 0x32, 0x31, 0x8f, 0xec, 0x3c, 0x43, 0x01, 0x29, 0x47,
	0xac, 0xfe, 0x17, 0xb1, 0x76, 0x29, 0x62, 0xbd, 0x18, 0xb1, 0x70, 0x38, 0xb5, 0x74, 0x38, 0xf2,
	0x00, 0x1a, 0x13, 0x82, 0x0b, 0xd4, 0xb5, 0x7e, 0x6d, 0xd0, 0x1a, 0xeb, 0x66, 0x69, 0xeb, 0xe6,
	0x34, 0x69, 0x2e, 0xbe, 0x02, 0x66, 0xe5, 0x3c, 0x63, 0x0a, 0xcd, 0x03, 0x98, 0x04, 0xc1, 0x90,
	0x86, 0x11, 0x3e, 0xf3, 0x55, 0xb6, 0x0c, 0xd5, 0x2a, 0x20, 0x89, 0xb1, 0xc7, 0x10, 0xe9, 0x46,
	0xc6, 0x90, 0xe5, 0xf8, 0x5b, 0x81, 0xc6, 0x24, 0x0a, 0xb7, 0xf3, 0xd8, 0x21, 0x2f, 0xa0, 0x65,
	0x03, 0x22, 0xbd, 0x13, 0xfb, 0xd2, 0xe6, 0xbb, 0x77, 0x17, 0xba, 0xd9, 0x54, 0x8d, 0x8a, 0x14,
	0x9a, 0xf9, 0x67, 0x85, 0x0e, 0xd7, 0xf0, 0xac, 0xd0, 0xf1, 0x62, 0x1b, 0x95, 0xa7, 0xce, 0x7b,
	0xbb, 0xf4, 0x34, 0x6c, 0x2d, 0x7d, 0x13, 0x8f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0x1d,
	0xa0, 0x03, 0x32, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthSvcClient is the client API for AuthSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthSvcClient interface {
	// sign up new user
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
}

type authSvcClient struct {
	cc *grpc.ClientConn
}

func NewAuthSvcClient(cc *grpc.ClientConn) AuthSvcClient {
	return &authSvcClient{cc}
}

func (c *authSvcClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/pb.authsvc.v1.AuthSvc/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authSvcClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/pb.authsvc.v1.AuthSvc/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthSvcServer is the server API for AuthSvc service.
type AuthSvcServer interface {
	// sign up new user
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
}

// UnimplementedAuthSvcServer can be embedded to have forward compatible implementations.
type UnimplementedAuthSvcServer struct {
}

func (*UnimplementedAuthSvcServer) SignUp(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedAuthSvcServer) SignIn(ctx context.Context, req *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}

func RegisterAuthSvcServer(s *grpc.Server, srv AuthSvcServer) {
	s.RegisterService(&_AuthSvc_serviceDesc, srv)
}

func _AuthSvc_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSvcServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.authsvc.v1.AuthSvc/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSvcServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthSvc_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSvcServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.authsvc.v1.AuthSvc/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSvcServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.authsvc.v1.AuthSvc",
	HandlerType: (*AuthSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _AuthSvc_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _AuthSvc_SignIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/authsvc/v1/authsvc.proto",
}
