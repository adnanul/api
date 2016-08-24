// Code generated by protoc-gen-go.
// source: cluster.proto
// DO NOT EDIT!

/*
Package glusterfs is a generated protocol buffer package.

It is generated from these files:
	cluster.proto
	volume.proto

It has these top-level messages:
	Cluster
	ClusterListRequest
	ClusterListResponse
	ClusterDescribeRequest
	ClusterDescribeResponse
	ClusterCreateRequest
	ClusterDeleteRequest
	VolumeCreateRequest
	VolumeDeleteRequest
	VolumeListRequest
	VolumeListResponse
	Volume
*/
package glusterfs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
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

type Cluster struct {
	Phid          string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	KubeCluster   string `protobuf:"bytes,3,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	Mood          string `protobuf:"bytes,4,opt,name=mood" json:"mood,omitempty"`
	KubeNamespace string `protobuf:"bytes,5,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	Replica       int64  `protobuf:"varint,6,opt,name=replica" json:"replica,omitempty"`
	BaculaEnabled int32  `protobuf:"varint,7,opt,name=BaculaEnabled,json=baculaEnabled" json:"BaculaEnabled,omitempty"`
	Created       string `protobuf:"bytes,8,opt,name=created" json:"created,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClusterListRequest struct {
	KubeCluster string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
}

func (m *ClusterListRequest) Reset()                    { *m = ClusterListRequest{} }
func (m *ClusterListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterListRequest) ProtoMessage()               {}
func (*ClusterListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ClusterListResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Glusterfs []*Cluster     `protobuf:"bytes,2,rep,name=glusterfs" json:"glusterfs,omitempty"`
}

func (m *ClusterListResponse) Reset()                    { *m = ClusterListResponse{} }
func (m *ClusterListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterListResponse) ProtoMessage()               {}
func (*ClusterListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClusterListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterListResponse) GetGlusterfs() []*Cluster {
	if m != nil {
		return m.Glusterfs
	}
	return nil
}

type ClusterDescribeRequest struct {
	KubeCluster string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDescribeRequest) Reset()                    { *m = ClusterDescribeRequest{} }
func (m *ClusterDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeRequest) ProtoMessage()               {}
func (*ClusterDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ClusterDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *ClusterDescribeResponse) Reset()                    { *m = ClusterDescribeResponse{} }
func (m *ClusterDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeResponse) ProtoMessage()               {}
func (*ClusterDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClusterDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ClusterCreateRequest struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Node        int64    `protobuf:"varint,2,opt,name=node" json:"node,omitempty"`
	Mood        string   `protobuf:"bytes,3,opt,name=mood" json:"mood,omitempty"`
	Disks       []string `protobuf:"bytes,4,rep,name=disks" json:"disks,omitempty"`
	KubeCluster string   `protobuf:"bytes,5,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
}

func (m *ClusterCreateRequest) Reset()                    { *m = ClusterCreateRequest{} }
func (m *ClusterCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterCreateRequest) ProtoMessage()               {}
func (*ClusterCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ClusterDeleteRequest struct {
	KubeCluster string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDeleteRequest) Reset()                    { *m = ClusterDeleteRequest{} }
func (m *ClusterDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDeleteRequest) ProtoMessage()               {}
func (*ClusterDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Cluster)(nil), "glusterfs.Cluster")
	proto.RegisterType((*ClusterListRequest)(nil), "glusterfs.ClusterListRequest")
	proto.RegisterType((*ClusterListResponse)(nil), "glusterfs.ClusterListResponse")
	proto.RegisterType((*ClusterDescribeRequest)(nil), "glusterfs.ClusterDescribeRequest")
	proto.RegisterType((*ClusterDescribeResponse)(nil), "glusterfs.ClusterDescribeResponse")
	proto.RegisterType((*ClusterCreateRequest)(nil), "glusterfs.ClusterCreateRequest")
	proto.RegisterType((*ClusterDeleteRequest)(nil), "glusterfs.ClusterDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Clusters service

type ClustersClient interface {
	// Glusterfs cluster list. Needs to work with two modes.
	// First is to list all the glusterfs cluster through the
	// space with out considering the kubernetes cluster. if the
	// cluster_name is provided then list all the glusterfs cluster
	// with respect to the provided kube cluster space.
	List(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error)
	Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error)
	Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type clustersClient struct {
	cc *grpc.ClientConn
}

func NewClustersClient(cc *grpc.ClientConn) ClustersClient {
	return &clustersClient{cc}
}

func (c *clustersClient) List(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error) {
	out := new(ClusterListResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error) {
	out := new(ClusterDescribeResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Clusters service

type ClustersServer interface {
	// Glusterfs cluster list. Needs to work with two modes.
	// First is to list all the glusterfs cluster through the
	// space with out considering the kubernetes cluster. if the
	// cluster_name is provided then list all the glusterfs cluster
	// with respect to the provided kube cluster space.
	List(context.Context, *ClusterListRequest) (*ClusterListResponse, error)
	Describe(context.Context, *ClusterDescribeRequest) (*ClusterDescribeResponse, error)
	Create(context.Context, *ClusterCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *ClusterDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterClustersServer(s *grpc.Server, srv ClustersServer) {
	s.RegisterService(&_Clusters_serviceDesc, srv)
}

func _Clusters_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Clusters/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).List(ctx, req.(*ClusterListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Clusters/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Describe(ctx, req.(*ClusterDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Clusters/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Create(ctx, req.(*ClusterCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Clusters/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Delete(ctx, req.(*ClusterDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Clusters_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glusterfs.Clusters",
	HandlerType: (*ClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Clusters_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Clusters_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Clusters_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Clusters_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0xd6, 0xf9, 0x9d, 0x90, 0x1e, 0x86, 0xa8, 0xac, 0x2c, 0x7e, 0x52, 0x0b, 0x50, 0xd4,
	0x83, 0x5d, 0x02, 0x02, 0xf1, 0x27, 0x04, 0x29, 0x37, 0x7e, 0x24, 0x23, 0x71, 0x45, 0x1b, 0x7b,
	0x09, 0x56, 0x5d, 0xaf, 0xc9, 0x6e, 0x90, 0x50, 0x55, 0x09, 0x01, 0x37, 0x8e, 0x48, 0xbc, 0x08,
	0x8f, 0xc2, 0x2b, 0x70, 0xe1, 0x2d, 0xd0, 0xee, 0xda, 0xa9, 0x83, 0x5b, 0x04, 0xed, 0x25, 0xda,
	0xf9, 0x66, 0x67, 0xf6, 0x9b, 0x6f, 0xbe, 0x18, 0xfa, 0x51, 0xba, 0x90, 0x8a, 0xcf, 0xfd, 0x7c,
	0x2e, 0x94, 0xc0, 0xee, 0xcc, 0x86, 0xaf, 0xa5, 0x7b, 0x7e, 0x26, 0xc4, 0x2c, 0xe5, 0x01, 0xcb,
	0x93, 0x80, 0x65, 0x99, 0x50, 0x4c, 0x25, 0x22, 0x93, 0xf6, 0xa2, 0xbb, 0xa1, 0xe1, 0x58, 0xbd,
	0xcf, 0xb9, 0x0c, 0xcc, 0xaf, 0xc5, 0xbd, 0x5f, 0x04, 0xda, 0x13, 0xdb, 0x03, 0x11, 0x1a, 0xf9,
	0x9b, 0x24, 0xa6, 0x64, 0x48, 0x46, 0xdd, 0xd0, 0x9c, 0x35, 0x96, 0xb1, 0x3d, 0x4e, 0xd7, 0x2c,
	0xa6, 0xcf, 0xb8, 0x09, 0x67, 0x76, 0x17, 0x53, 0xfe, 0xaa, 0xa0, 0x42, 0x1d, 0x93, 0xeb, 0x69,
	0xac, 0xd2, 0x6a, 0x4f, 0x88, 0x98, 0x36, 0x6c, 0x99, 0x3e, 0xe3, 0x15, 0x58, 0x37, 0x65, 0xba,
	0x87, 0xcc, 0x59, 0xc4, 0x69, 0xd3, 0x64, 0xfb, 0x1a, 0x7d, 0x56, 0x82, 0x48, 0xa1, 0x3d, 0xe7,
	0x79, 0x9a, 0x44, 0x8c, 0xb6, 0x86, 0x64, 0xe4, 0x84, 0x65, 0x88, 0x97, 0xa1, 0xff, 0x88, 0x45,
	0x8b, 0x94, 0x3d, 0xce, 0xd8, 0x34, 0xe5, 0x31, 0x6d, 0x0f, 0xc9, 0xa8, 0x19, 0xf6, 0xa7, 0x55,
	0x50, 0xd7, 0x47, 0x73, 0xce, 0x14, 0x8f, 0x69, 0xc7, 0xf4, 0x2f, 0x43, 0xef, 0x16, 0x60, 0xc1,
	0xef, 0x49, 0x22, 0x55, 0xc8, 0xdf, 0x2e, 0xb8, 0x54, 0xb5, 0x69, 0x48, 0x6d, 0x1a, 0x4f, 0xc0,
	0xd9, 0x95, 0x42, 0x99, 0x8b, 0x4c, 0x72, 0xbc, 0x0a, 0x2d, 0xa9, 0x98, 0x5a, 0x48, 0x53, 0xd3,
	0x1b, 0xaf, 0xfb, 0x56, 0x60, 0xff, 0x85, 0x41, 0xc3, 0x22, 0x8b, 0xdb, 0x70, 0xb8, 0x26, 0xba,
	0x36, 0x74, 0x46, 0xbd, 0x31, 0xfa, 0x4b, 0xc4, 0x2f, 0x5a, 0x87, 0x87, 0x97, 0xbc, 0xe7, 0xb0,
	0x51, 0xa0, 0x3b, 0x5c, 0x46, 0xf3, 0x64, 0xca, 0xff, 0x9d, 0xed, 0x51, 0x2b, 0xf3, 0x1e, 0xc2,
	0xb9, 0x5a, 0xc3, 0xff, 0x9b, 0xc2, 0xfb, 0x42, 0x60, 0x50, 0xf4, 0x98, 0x18, 0x41, 0x4b, 0x4a,
	0xe5, 0x7b, 0xa4, 0x62, 0x11, 0x8d, 0x89, 0xd8, 0x72, 0x70, 0x42, 0x73, 0x5e, 0x7a, 0xc2, 0xa9,
	0x78, 0x62, 0x00, 0xcd, 0x38, 0x91, 0xbb, 0x92, 0x36, 0x86, 0xce, 0xa8, 0x1b, 0xda, 0xa0, 0x36,
	0x64, 0xb3, 0xbe, 0x92, 0xa7, 0x4b, 0x32, 0x3b, 0x3c, 0xe5, 0xea, 0x94, 0xfa, 0x8c, 0xbf, 0x37,
	0xa0, 0x53, 0xe4, 0x25, 0x7e, 0x22, 0xd0, 0xd0, 0x8b, 0xc6, 0x0b, 0xf5, 0x2d, 0x55, 0x9c, 0xe3,
	0x5e, 0x3c, 0x2e, 0x6d, 0x95, 0xf5, 0xee, 0x7d, 0xfc, 0xf1, 0xf3, 0xeb, 0xda, 0x4d, 0xbc, 0x11,
	0xb0, 0x3c, 0x97, 0x91, 0x88, 0xed, 0x9f, 0x73, 0x59, 0x14, 0xbc, 0xdb, 0xf6, 0xaf, 0x05, 0x05,
	0x53, 0x19, 0xec, 0x57, 0x89, 0x1f, 0xe0, 0x37, 0x02, 0x9d, 0x72, 0x59, 0xb8, 0x59, 0x7f, 0xea,
	0x0f, 0x67, 0xb8, 0xde, 0xdf, 0xae, 0x14, 0x8c, 0x26, 0x86, 0xd1, 0x7d, 0xbc, 0x7b, 0x12, 0x46,
	0xc1, 0xbe, 0x96, 0xea, 0x00, 0x3f, 0x10, 0x68, 0x59, 0x07, 0xe0, 0xa5, 0xfa, 0x9b, 0x2b, 0xde,
	0x70, 0x07, 0xa5, 0x99, 0x5e, 0x8a, 0x24, 0x5e, 0xd2, 0x78, 0x60, 0x68, 0xdc, 0x76, 0x4f, 0x24,
	0xcc, 0x1d, 0xb2, 0x85, 0x9f, 0x09, 0xb4, 0xec, 0xde, 0x8f, 0xa2, 0xb0, 0xe2, 0x88, 0x63, 0x28,
	0x14, 0x4a, 0x6c, 0x9d, 0x46, 0x89, 0x69, 0xcb, 0x7c, 0x43, 0xaf, 0xff, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0xeb, 0x87, 0xb1, 0x95, 0x05, 0x00, 0x00,
}
