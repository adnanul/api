// Code generated by protoc-gen-go.
// source: longrunning.proto
// DO NOT EDIT!

/*
Package testlongrunning is a generated protocol buffer package.

It is generated from these files:
	longrunning.proto

It has these top-level messages:
	LongRunningRequest
*/
package testlongrunning

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LongRunningRequest struct {
	Duration int32 `protobuf:"varint,1,opt,name=duration" json:"duration,omitempty"`
}

func (m *LongRunningRequest) Reset()                    { *m = LongRunningRequest{} }
func (m *LongRunningRequest) String() string            { return proto.CompactTextString(m) }
func (*LongRunningRequest) ProtoMessage()               {}
func (*LongRunningRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LongRunningRequest) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func init() {
	proto.RegisterType((*LongRunningRequest)(nil), "appscode.testlongrunning.LongRunningRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestLongRunning service

type TestLongRunningClient interface {
	Run(ctx context.Context, in *LongRunningRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error)
}

type testLongRunningClient struct {
	cc *grpc.ClientConn
}

func NewTestLongRunningClient(cc *grpc.ClientConn) TestLongRunningClient {
	return &testLongRunningClient{cc}
}

func (c *testLongRunningClient) Run(ctx context.Context, in *LongRunningRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error) {
	out := new(appscode_dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/appscode.testlongrunning.TestLongRunning/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestLongRunning service

type TestLongRunningServer interface {
	Run(context.Context, *LongRunningRequest) (*appscode_dtypes.LongRunningResponse, error)
}

func RegisterTestLongRunningServer(s *grpc.Server, srv TestLongRunningServer) {
	s.RegisterService(&_TestLongRunning_serviceDesc, srv)
}

func _TestLongRunning_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LongRunningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestLongRunningServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.testlongrunning.TestLongRunning/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestLongRunningServer).Run(ctx, req.(*LongRunningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestLongRunning_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.testlongrunning.TestLongRunning",
	HandlerType: (*TestLongRunningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _TestLongRunning_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "longrunning.proto",
}

func init() { proto.RegisterFile("longrunning.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xc9, 0xcf, 0x4b,
	0x2f, 0x2a, 0xcd, 0xcb, 0xcb, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x48,
	0x2c, 0x28, 0x28, 0x4e, 0xce, 0x4f, 0x49, 0xd5, 0x2b, 0x49, 0x2d, 0x2e, 0x41, 0x92, 0x97, 0x92,
	0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f,
	0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xe8, 0x93, 0x92, 0x83, 0xe9, 0xc3, 0x21, 0x2f, 0x8f,
	0x22, 0x9f, 0x52, 0x52, 0x59, 0x90, 0x5a, 0xac, 0x0f, 0x26, 0x21, 0x0a, 0x94, 0x0c, 0xb8, 0x84,
	0x7c, 0xf2, 0xf3, 0xd2, 0x83, 0x20, 0xb6, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49,
	0x71, 0x71, 0xa4, 0x94, 0x16, 0x81, 0x4d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf3,
	0x8d, 0x16, 0x31, 0x72, 0xf1, 0x87, 0xa4, 0x16, 0x97, 0x20, 0x69, 0x13, 0xea, 0x63, 0xe4, 0x62,
	0x0e, 0x2a, 0xcd, 0x13, 0xd2, 0xd1, 0xc3, 0xe5, 0x0f, 0x3d, 0x4c, 0x5b, 0xa4, 0x54, 0x10, 0xaa,
	0x21, 0x2e, 0x43, 0x55, 0x54, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0xaa, 0x64, 0xdc, 0xb4, 0x55, 0x82,
	0x89, 0x83, 0xb1, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xea, 0x42, 0xaa, 0xfa, 0xf1, 0x28, 0x5e, 0x42,
	0xb2, 0x42, 0xbf, 0x1a, 0xe6, 0xc6, 0x5a, 0x27, 0x37, 0x2e, 0x85, 0xe4, 0xfc, 0x5c, 0x84, 0xf9,
	0x89, 0x05, 0x99, 0xe8, 0x2e, 0x72, 0x42, 0xf7, 0x45, 0x00, 0x63, 0x14, 0x3f, 0x9a, 0x9a, 0x24,
	0x36, 0x70, 0x28, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x16, 0xf8, 0x9b, 0x9e, 0xb3, 0x01,
	0x00, 0x00,
}
