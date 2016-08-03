// Code generated by protoc-gen-go.
// source: artifact.proto
// DO NOT EDIT!

/*
Package artifactory is a generated protocol buffer package.

It is generated from these files:
	artifact.proto
	version.proto

It has these top-level messages:
	ArtifactSearchRequest
	ArtifactSearchResponse
	ArtifactListRequest
	ArtifactListResponse
	Artifact
	VersionListRequest
	VersionListResponse
	VersionDescribeRequest
	VersionDescribeResponse
	ArtifactVersion
	JavaSpec
	DockerSpec
	PhpSpec
	NpmSpec
*/
package artifactory

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import dtypes "github.com/appscode/api/dtypes"

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

type ArtifactSearchRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *ArtifactSearchRequest) Reset()                    { *m = ArtifactSearchRequest{} }
func (m *ArtifactSearchRequest) String() string            { return proto.CompactTextString(m) }
func (*ArtifactSearchRequest) ProtoMessage()               {}
func (*ArtifactSearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ArtifactSearchResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Artifacts []*Artifact    `protobuf:"bytes,2,rep,name=artifacts" json:"artifacts,omitempty"`
}

func (m *ArtifactSearchResponse) Reset()                    { *m = ArtifactSearchResponse{} }
func (m *ArtifactSearchResponse) String() string            { return proto.CompactTextString(m) }
func (*ArtifactSearchResponse) ProtoMessage()               {}
func (*ArtifactSearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ArtifactSearchResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ArtifactSearchResponse) GetArtifacts() []*Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type ArtifactListRequest struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ArtifactListRequest) Reset()                    { *m = ArtifactListRequest{} }
func (m *ArtifactListRequest) String() string            { return proto.CompactTextString(m) }
func (*ArtifactListRequest) ProtoMessage()               {}
func (*ArtifactListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ArtifactListResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Artifacts []*Artifact    `protobuf:"bytes,2,rep,name=artifacts" json:"artifacts,omitempty"`
}

func (m *ArtifactListResponse) Reset()                    { *m = ArtifactListResponse{} }
func (m *ArtifactListResponse) String() string            { return proto.CompactTextString(m) }
func (*ArtifactListResponse) ProtoMessage()               {}
func (*ArtifactListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ArtifactListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ArtifactListResponse) GetArtifacts() []*Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type Artifact struct {
	Name             string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type             string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	LastModifiedTime string `protobuf:"bytes,3,opt,name=last_modified_time,json=lastModifiedTime" json:"last_modified_time,omitempty"`
}

func (m *Artifact) Reset()                    { *m = Artifact{} }
func (m *Artifact) String() string            { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()               {}
func (*Artifact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*ArtifactSearchRequest)(nil), "artifactory.ArtifactSearchRequest")
	proto.RegisterType((*ArtifactSearchResponse)(nil), "artifactory.ArtifactSearchResponse")
	proto.RegisterType((*ArtifactListRequest)(nil), "artifactory.ArtifactListRequest")
	proto.RegisterType((*ArtifactListResponse)(nil), "artifactory.ArtifactListResponse")
	proto.RegisterType((*Artifact)(nil), "artifactory.Artifact")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Artifacts service

type ArtifactsClient interface {
	Search(ctx context.Context, in *ArtifactSearchRequest, opts ...grpc.CallOption) (*ArtifactSearchResponse, error)
	List(ctx context.Context, in *ArtifactListRequest, opts ...grpc.CallOption) (*ArtifactListResponse, error)
}

type artifactsClient struct {
	cc *grpc.ClientConn
}

func NewArtifactsClient(cc *grpc.ClientConn) ArtifactsClient {
	return &artifactsClient{cc}
}

func (c *artifactsClient) Search(ctx context.Context, in *ArtifactSearchRequest, opts ...grpc.CallOption) (*ArtifactSearchResponse, error) {
	out := new(ArtifactSearchResponse)
	err := grpc.Invoke(ctx, "/artifactory.Artifacts/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactsClient) List(ctx context.Context, in *ArtifactListRequest, opts ...grpc.CallOption) (*ArtifactListResponse, error) {
	out := new(ArtifactListResponse)
	err := grpc.Invoke(ctx, "/artifactory.Artifacts/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Artifacts service

type ArtifactsServer interface {
	Search(context.Context, *ArtifactSearchRequest) (*ArtifactSearchResponse, error)
	List(context.Context, *ArtifactListRequest) (*ArtifactListResponse, error)
}

func RegisterArtifactsServer(s *grpc.Server, srv ArtifactsServer) {
	s.RegisterService(&_Artifacts_serviceDesc, srv)
}

func _Artifacts_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactsServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artifactory.Artifacts/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactsServer).Search(ctx, req.(*ArtifactSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Artifacts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artifactory.Artifacts/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactsServer).List(ctx, req.(*ArtifactListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Artifacts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "artifactory.Artifacts",
	HandlerType: (*ArtifactsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Artifacts_Search_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Artifacts_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("artifact.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x52, 0xcd, 0x4e, 0x2a, 0x31,
	0x18, 0xcd, 0x00, 0x97, 0x5c, 0x3e, 0x12, 0x72, 0xd3, 0x0b, 0x84, 0x4c, 0xee, 0x82, 0x3b, 0xc6,
	0xbf, 0x44, 0xa7, 0x02, 0x0b, 0xd7, 0xec, 0x75, 0x33, 0xb8, 0xc7, 0xca, 0x14, 0x6c, 0x02, 0xd3,
	0x71, 0x5a, 0x4c, 0x88, 0x31, 0x31, 0xc6, 0x37, 0x30, 0x3e, 0x99, 0xaf, 0xe0, 0x83, 0xd8, 0x9f,
	0x29, 0x3f, 0x86, 0xe0, 0xce, 0xcd, 0xa4, 0x73, 0x7a, 0x7a, 0xce, 0xe9, 0xe9, 0x07, 0x35, 0x92,
	0x49, 0x36, 0x26, 0x23, 0x19, 0xa6, 0x19, 0x97, 0x1c, 0x55, 0xdd, 0x3f, 0xcf, 0x16, 0xfe, 0xbf,
	0x09, 0xe7, 0x93, 0x29, 0xc5, 0x24, 0x65, 0x98, 0x24, 0x09, 0x97, 0x44, 0x32, 0x9e, 0x08, 0x4b,
	0xf5, 0x9b, 0x1a, 0x8e, 0xe5, 0x22, 0xa5, 0x02, 0x9b, 0xaf, 0xc5, 0x83, 0x3e, 0x34, 0xfa, 0xb9,
	0xc8, 0x80, 0x92, 0x6c, 0x74, 0x1b, 0xd1, 0xbb, 0x39, 0x15, 0x12, 0xd5, 0xe1, 0x97, 0x5a, 0x64,
	0x8b, 0x96, 0xd7, 0xf6, 0x8e, 0x2a, 0x91, 0xfd, 0x41, 0x08, 0x4a, 0xfa, 0x74, 0xab, 0x60, 0x40,
	0xb3, 0x0e, 0xe6, 0xd0, 0xfc, 0x2a, 0x21, 0x52, 0xe5, 0x4c, 0xd1, 0x01, 0x94, 0x85, 0x8a, 0x31,
	0x17, 0x46, 0xa4, 0xda, 0xad, 0x85, 0x36, 0x41, 0x38, 0x30, 0x68, 0x94, 0xef, 0xa2, 0x1e, 0x54,
	0xdc, 0x4d, 0x84, 0x92, 0x2e, 0x2a, 0x6a, 0x23, 0x5c, 0xbb, 0x5b, 0xe8, 0xf4, 0xa3, 0x15, 0x2f,
	0x38, 0x86, 0xbf, 0x0e, 0xbe, 0x60, 0x42, 0xba, 0xdc, 0x2e, 0xa1, 0xb7, 0x96, 0x50, 0x40, 0x7d,
	0x93, 0xfa, 0x13, 0xf9, 0xae, 0xe1, 0xb7, 0x83, 0x75, 0xa8, 0x84, 0xcc, 0x96, 0xa1, 0xf4, 0x7a,
	0x5b, 0x95, 0xe8, 0x04, 0xd0, 0x94, 0x08, 0x39, 0x9c, 0xf1, 0x98, 0x8d, 0x19, 0x8d, 0x87, 0x92,
	0xa9, 0x53, 0x45, 0xc3, 0xf8, 0xa3, 0x77, 0x2e, 0xf3, 0x8d, 0x2b, 0x85, 0x77, 0xdf, 0x0a, 0x50,
	0x71, 0x16, 0x02, 0x3d, 0x79, 0x50, 0xb6, 0xfd, 0xa3, 0x60, 0x6b, 0xb8, 0x8d, 0xf7, 0xf5, 0xf7,
	0x76, 0x72, 0x6c, 0x41, 0xc1, 0xe9, 0xf3, 0xfb, 0xc7, 0x6b, 0xe1, 0x10, 0xed, 0xab, 0xa9, 0x4a,
	0xc5, 0x88, 0xc7, 0xf9, 0x78, 0xad, 0x4e, 0xe2, 0xfb, 0xb3, 0xb0, 0x83, 0x85, 0xf5, 0x7d, 0xf1,
	0xa0, 0xa4, 0x0b, 0x46, 0xed, 0xad, 0xe2, 0x6b, 0xcf, 0xe4, 0xff, 0xdf, 0xc1, 0xc8, 0xcd, 0xcf,
	0x8d, 0x79, 0x07, 0xe1, 0x6f, 0xcc, 0x97, 0x95, 0xe3, 0x07, 0x5d, 0xe2, 0xe3, 0x4d, 0xd9, 0x8c,
	0x76, 0xef, 0x33, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x02, 0x07, 0x0f, 0x2f, 0x03, 0x00, 0x00,
}
