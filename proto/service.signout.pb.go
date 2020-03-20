// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.signout.proto

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

type SignoutRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignoutRequest) Reset()         { *m = SignoutRequest{} }
func (m *SignoutRequest) String() string { return proto.CompactTextString(m) }
func (*SignoutRequest) ProtoMessage()    {}
func (*SignoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efcf74650f51937d, []int{0}
}

func (m *SignoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignoutRequest.Unmarshal(m, b)
}
func (m *SignoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignoutRequest.Marshal(b, m, deterministic)
}
func (m *SignoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignoutRequest.Merge(m, src)
}
func (m *SignoutRequest) XXX_Size() int {
	return xxx_messageInfo_SignoutRequest.Size(m)
}
func (m *SignoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignoutRequest proto.InternalMessageInfo

func (m *SignoutRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SignoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SignoutResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignoutResponse) Reset()         { *m = SignoutResponse{} }
func (m *SignoutResponse) String() string { return proto.CompactTextString(m) }
func (*SignoutResponse) ProtoMessage()    {}
func (*SignoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_efcf74650f51937d, []int{1}
}

func (m *SignoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignoutResponse.Unmarshal(m, b)
}
func (m *SignoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignoutResponse.Marshal(b, m, deterministic)
}
func (m *SignoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignoutResponse.Merge(m, src)
}
func (m *SignoutResponse) XXX_Size() int {
	return xxx_messageInfo_SignoutResponse.Size(m)
}
func (m *SignoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignoutResponse proto.InternalMessageInfo

func (m *SignoutResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*SignoutRequest)(nil), "proto.SignoutRequest")
	proto.RegisterType((*SignoutResponse)(nil), "proto.SignoutResponse")
}

func init() {
	proto.RegisterFile("service.signout.proto", fileDescriptor_efcf74650f51937d)
}

var fileDescriptor_efcf74650f51937d = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0xce, 0x4c, 0xcf, 0xcb, 0x2f, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x66, 0x5c, 0x7c, 0xc1, 0x10, 0xf1, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20,
	0xa6, 0xcc, 0x14, 0x21, 0x11, 0x2e, 0xd6, 0x92, 0xfc, 0xec, 0xd4, 0x3c, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x08, 0x47, 0x49, 0x93, 0x8b, 0x1f, 0xae, 0xaf, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x55, 0x48, 0x8c, 0x8b, 0xad, 0xb8, 0x24, 0xb1, 0xa4, 0xb4, 0x18, 0xac, 0x99, 0x23, 0x08, 0xca,
	0x33, 0xf2, 0x82, 0x5b, 0x11, 0x0c, 0x71, 0x89, 0x90, 0x05, 0x17, 0x3b, 0x54, 0x44, 0x48, 0x14,
	0xe2, 0x1c, 0x3d, 0x54, 0x47, 0x48, 0x89, 0xa1, 0x0b, 0x43, 0xec, 0x48, 0x62, 0x03, 0x0b, 0x1b,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xf9, 0x78, 0x75, 0xd5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SignoutServiceClient is the client API for SignoutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignoutServiceClient interface {
	Signout(ctx context.Context, in *SignoutRequest, opts ...grpc.CallOption) (*SignoutResponse, error)
}

type signoutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignoutServiceClient(cc grpc.ClientConnInterface) SignoutServiceClient {
	return &signoutServiceClient{cc}
}

func (c *signoutServiceClient) Signout(ctx context.Context, in *SignoutRequest, opts ...grpc.CallOption) (*SignoutResponse, error) {
	out := new(SignoutResponse)
	err := c.cc.Invoke(ctx, "/proto.SignoutService/Signout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignoutServiceServer is the server API for SignoutService service.
type SignoutServiceServer interface {
	Signout(context.Context, *SignoutRequest) (*SignoutResponse, error)
}

// UnimplementedSignoutServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSignoutServiceServer struct {
}

func (*UnimplementedSignoutServiceServer) Signout(ctx context.Context, req *SignoutRequest) (*SignoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signout not implemented")
}

func RegisterSignoutServiceServer(s *grpc.Server, srv SignoutServiceServer) {
	s.RegisterService(&_SignoutService_serviceDesc, srv)
}

func _SignoutService_Signout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignoutServiceServer).Signout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SignoutService/Signout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignoutServiceServer).Signout(ctx, req.(*SignoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SignoutService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SignoutService",
	HandlerType: (*SignoutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signout",
			Handler:    _SignoutService_Signout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.signout.proto",
}
