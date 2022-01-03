// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/casesvc/v1/casesvc.proto

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

// case response
type CaseResponse struct {
	// unique id of case
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	OrgId       string `protobuf:"bytes,4,opt,name=orgId,proto3" json:"orgId,omitempty"`
	CreatedBy   string `protobuf:"bytes,5,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// reference id of case
	ReferenceId          string       `protobuf:"bytes,6,opt,name=referenceId,proto3" json:"referenceId,omitempty"`
	Errors               []*ErrorType `protobuf:"bytes,7,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CaseResponse) Reset()         { *m = CaseResponse{} }
func (m *CaseResponse) String() string { return proto.CompactTextString(m) }
func (*CaseResponse) ProtoMessage()    {}
func (*CaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7670596bad00a352, []int{0}
}

func (m *CaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaseResponse.Unmarshal(m, b)
}
func (m *CaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaseResponse.Marshal(b, m, deterministic)
}
func (m *CaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaseResponse.Merge(m, src)
}
func (m *CaseResponse) XXX_Size() int {
	return xxx_messageInfo_CaseResponse.Size(m)
}
func (m *CaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CaseResponse proto.InternalMessageInfo

func (m *CaseResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CaseResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CaseResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CaseResponse) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *CaseResponse) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *CaseResponse) GetReferenceId() string {
	if m != nil {
		return m.ReferenceId
	}
	return ""
}

func (m *CaseResponse) GetErrors() []*ErrorType {
	if m != nil {
		return m.Errors
	}
	return nil
}

// create case request
type CreateRequest struct {
	// title of case
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description of case
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// org id of case
	OrgId string `protobuf:"bytes,3,opt,name=orgId,proto3" json:"orgId,omitempty"`
	// id of user creating this case
	CreatedBy            string   `protobuf:"bytes,4,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7670596bad00a352, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateRequest) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *CreateRequest) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
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
	return fileDescriptor_7670596bad00a352, []int{2}
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
	proto.RegisterType((*CaseResponse)(nil), "pb.casesvc.v1.CaseResponse")
	proto.RegisterType((*CreateRequest)(nil), "pb.casesvc.v1.CreateRequest")
	proto.RegisterType((*ErrorType)(nil), "pb.casesvc.v1.ErrorType")
}

func init() { proto.RegisterFile("pb/casesvc/v1/casesvc.proto", fileDescriptor_7670596bad00a352) }

var fileDescriptor_7670596bad00a352 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xf2, 0x30,
	0x0c, 0xc7, 0xbf, 0x16, 0x28, 0xc2, 0x7c, 0x6c, 0x52, 0xb4, 0x43, 0x34, 0xd0, 0x84, 0x7a, 0xe2,
	0x54, 0x06, 0x7b, 0x03, 0x10, 0x07, 0x6e, 0x53, 0xb7, 0xd3, 0x6e, 0x6d, 0xe2, 0xa1, 0x4a, 0x5b,
	0x93, 0xc5, 0xa1, 0x12, 0xda, 0xbb, 0xee, 0x59, 0xa6, 0xa6, 0xeb, 0x68, 0x2b, 0x8d, 0x5b, 0xfc,
	0xff, 0xc7, 0x76, 0x7e, 0xb1, 0x61, 0xaa, 0xd3, 0xa5, 0x48, 0x08, 0xa9, 0x10, 0xcb, 0x62, 0x55,
	0x1f, 0x23, 0x6d, 0x94, 0x55, 0x6c, 0xa2, 0xd3, 0xa8, 0x56, 0x8a, 0x55, 0xf8, 0xe5, 0xc1, 0xff,
	0x6d, 0x42, 0x18, 0x23, 0x69, 0x95, 0x13, 0xb2, 0x2b, 0xf0, 0x33, 0xc9, 0xbd, 0xb9, 0xb7, 0x18,
	0xc5, 0x7e, 0x26, 0xd9, 0x0d, 0x0c, 0x6c, 0x66, 0xdf, 0x90, 0xfb, 0x4e, 0xaa, 0x02, 0x36, 0x87,
	0xb1, 0x44, 0x12, 0x26, 0xd3, 0x36, 0x53, 0x39, 0xef, 0x39, 0xaf, 0x29, 0x95, 0x79, 0xca, 0x1c,
	0xf6, 0x92, 0xf7, 0xab, 0x3c, 0x17, 0xb0, 0x19, 0x8c, 0x84, 0xc1, 0xc4, 0xa2, 0xdc, 0x9c, 0xf8,
	0xc0, 0x39, 0x67, 0xa1, 0xac, 0x6a, 0xf0, 0x15, 0x0d, 0xe6, 0x02, 0xf7, 0x92, 0x07, 0x55, 0xd5,
	0x86, 0xc4, 0xee, 0x21, 0x40, 0x63, 0x94, 0x21, 0x3e, 0x9c, 0xf7, 0x16, 0xe3, 0x35, 0x8f, 0x5a,
	0x38, 0xd1, 0xae, 0x34, 0x9f, 0x4f, 0x1a, 0xe3, 0x9f, 0x7b, 0xe1, 0x27, 0x4c, 0xb6, 0xae, 0x41,
	0x8c, 0x1f, 0x47, 0x24, 0x7b, 0x06, 0xf2, 0x2e, 0x00, 0xf9, 0x17, 0x80, 0x7a, 0x7f, 0x02, 0xf5,
	0x3b, 0x40, 0xe1, 0x0e, 0x46, 0xbf, 0x2f, 0x62, 0x77, 0x00, 0x64, 0x13, 0x7b, 0xa4, 0xad, 0x92,
	0x55, 0xf7, 0x41, 0xdc, 0x50, 0x18, 0x87, 0xe1, 0x3b, 0x12, 0x25, 0x87, 0xfa, 0xaf, 0xeb, 0x70,
	0xfd, 0x08, 0xc3, 0x72, 0x46, 0x4f, 0x85, 0x60, 0x3b, 0x08, 0x2a, 0x1c, 0x36, 0xeb, 0xa0, 0xb7,
	0x28, 0x6f, 0xa7, 0x5d, 0xb7, 0x31, 0xe3, 0xf0, 0xdf, 0xe6, 0xfa, 0x65, 0xd2, 0x5a, 0x92, 0x34,
	0x70, 0xdb, 0xf1, 0xf0, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x79, 0x44, 0xa8, 0x3c, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CaseSvcClient is the client API for CaseSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CaseSvcClient interface {
	// create new case
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CaseResponse, error)
}

type caseSvcClient struct {
	cc *grpc.ClientConn
}

func NewCaseSvcClient(cc *grpc.ClientConn) CaseSvcClient {
	return &caseSvcClient{cc}
}

func (c *caseSvcClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CaseResponse, error) {
	out := new(CaseResponse)
	err := c.cc.Invoke(ctx, "/pb.casesvc.v1.CaseSvc/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaseSvcServer is the server API for CaseSvc service.
type CaseSvcServer interface {
	// create new case
	Create(context.Context, *CreateRequest) (*CaseResponse, error)
}

// UnimplementedCaseSvcServer can be embedded to have forward compatible implementations.
type UnimplementedCaseSvcServer struct {
}

func (*UnimplementedCaseSvcServer) Create(ctx context.Context, req *CreateRequest) (*CaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterCaseSvcServer(s *grpc.Server, srv CaseSvcServer) {
	s.RegisterService(&_CaseSvc_serviceDesc, srv)
}

func _CaseSvc_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseSvcServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.casesvc.v1.CaseSvc/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseSvcServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CaseSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.casesvc.v1.CaseSvc",
	HandlerType: (*CaseSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CaseSvc_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/casesvc/v1/casesvc.proto",
}
