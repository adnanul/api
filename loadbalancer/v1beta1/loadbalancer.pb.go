// Code generated by protoc-gen-go.
// source: loadbalancer.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	loadbalancer.proto

It has these top-level messages:
	ListRequest
	ListResponse
	DescribeRequest
	DescribeResponse
	CreateRequest
	UpdateRequest
	DeleteRequest
	LoadBalancer
	Spec
	Status
	LoadBalancerStatus
	LoadBalancerBackend
	LoadBalancerRule
	HTTPLoadBalancerRule
	TCPLoadBalancerRule
*/
package v1beta1

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

type ListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListResponse struct {
	Status       *dtypes.Status  `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	LoadBalancer []*LoadBalancer `protobuf:"bytes,2,rep,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListResponse) GetLoadBalancer() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type DescribeRequest struct {
	Kind      string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Cluster   string `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DescribeRequest) Reset()                    { *m = DescribeRequest{} }
func (m *DescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRequest) ProtoMessage()               {}
func (*DescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DescribeResponse struct {
	Status       *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	LoadBalancer *LoadBalancer  `protobuf:"bytes,2,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *DescribeResponse) Reset()                    { *m = DescribeResponse{} }
func (m *DescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse) ProtoMessage()               {}
func (*DescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeResponse) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type CreateRequest struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Namespace    string        `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Cluster      string        `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
	LoadBalancer *LoadBalancer `protobuf:"bytes,4,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateRequest) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type UpdateRequest struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Cluster      string        `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	LoadBalancer *LoadBalancer `protobuf:"bytes,3,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateRequest) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type DeleteRequest struct {
	Kind      string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Cluster   string `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type LoadBalancer struct {
	// 'kind' defines is it the regular kubernetes instance or the
	// appscode superset called Extended Ingress. This field will
	// strictly contains only those two values
	// 'ingress' - default kubernetes ingress object.
	// 'extendedIngress' - appscode superset of ingress.
	// when creating a Loadbalancer from UI this field will always
	// be only 'extendedIngress.' List, Describe, Update and Delete
	// will support both two modes.
	// Create will support only extendedIngress.
	// For Creating or Updating an regular ingress one must use the
	// kubectl or direct API calls directly to kubernetes.
	Kind              string            `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name              string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace         string            `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	CreationTimestamp int64             `protobuf:"varint,4,opt,name=creation_timestamp,json=creationTimestamp" json:"creation_timestamp,omitempty"`
	Options           map[string]string `protobuf:"bytes,5,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Spec              *Spec             `protobuf:"bytes,6,opt,name=spec" json:"spec,omitempty"`
	Status            *Status           `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	Json              string            `protobuf:"bytes,8,opt,name=json" json:"json,omitempty"`
}

func (m *LoadBalancer) Reset()                    { *m = LoadBalancer{} }
func (m *LoadBalancer) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancer) ProtoMessage()               {}
func (*LoadBalancer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LoadBalancer) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *LoadBalancer) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *LoadBalancer) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type Spec struct {
	Backend *HTTPLoadBalancerRule `protobuf:"bytes,1,opt,name=backend" json:"backend,omitempty"`
	Rules   []*LoadBalancerRule   `protobuf:"bytes,2,rep,name=rules" json:"rules,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Spec) GetBackend() *HTTPLoadBalancerRule {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *Spec) GetRules() []*LoadBalancerRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Status struct {
	Status []*LoadBalancerStatus `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Status) GetStatus() []*LoadBalancerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type LoadBalancerStatus struct {
	IP   string `protobuf:"bytes,1,opt,name=IP,json=iP" json:"IP,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *LoadBalancerStatus) Reset()                    { *m = LoadBalancerStatus{} }
func (m *LoadBalancerStatus) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerStatus) ProtoMessage()               {}
func (*LoadBalancerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type LoadBalancerBackend struct {
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	ServicePort string `protobuf:"bytes,2,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
}

func (m *LoadBalancerBackend) Reset()                    { *m = LoadBalancerBackend{} }
func (m *LoadBalancerBackend) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerBackend) ProtoMessage()               {}
func (*LoadBalancerBackend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type LoadBalancerRule struct {
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	// ssl secret name to enable https on the host.
	// ssl secret must contain data with the certs pem file.
	SSLSecretName string                  `protobuf:"bytes,5,opt,name=SSL_secret_name,json=sSLSecretName" json:"SSL_secret_name,omitempty"`
	Http          []*HTTPLoadBalancerRule `protobuf:"bytes,2,rep,name=http" json:"http,omitempty"`
	Tcp           []*TCPLoadBalancerRule  `protobuf:"bytes,3,rep,name=tcp" json:"tcp,omitempty"`
}

func (m *LoadBalancerRule) Reset()                    { *m = LoadBalancerRule{} }
func (m *LoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerRule) ProtoMessage()               {}
func (*LoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *LoadBalancerRule) GetHttp() []*HTTPLoadBalancerRule {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *LoadBalancerRule) GetTcp() []*TCPLoadBalancerRule {
	if m != nil {
		return m.Tcp
	}
	return nil
}

type HTTPLoadBalancerRule struct {
	Path        string               `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Backend     *LoadBalancerBackend `protobuf:"bytes,2,opt,name=backend" json:"backend,omitempty"`
	HeaderRule  []string             `protobuf:"bytes,3,rep,name=header_rule,json=headerRule" json:"header_rule,omitempty"`
	RewriteRule []string             `protobuf:"bytes,4,rep,name=rewrite_rule,json=rewriteRule" json:"rewrite_rule,omitempty"`
}

func (m *HTTPLoadBalancerRule) Reset()                    { *m = HTTPLoadBalancerRule{} }
func (m *HTTPLoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*HTTPLoadBalancerRule) ProtoMessage()               {}
func (*HTTPLoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *HTTPLoadBalancerRule) GetBackend() *LoadBalancerBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

type TCPLoadBalancerRule struct {
	Port          string               `protobuf:"bytes,1,opt,name=port" json:"port,omitempty"`
	Backend       *LoadBalancerBackend `protobuf:"bytes,2,opt,name=backend" json:"backend,omitempty"`
	SSLSecretName string               `protobuf:"bytes,3,opt,name=SSL_secret_name,json=sSLSecretName" json:"SSL_secret_name,omitempty"`
	SecretPemName string               `protobuf:"bytes,4,opt,name=secret_pem_name,json=secretPemName" json:"secret_pem_name,omitempty"`
}

func (m *TCPLoadBalancerRule) Reset()                    { *m = TCPLoadBalancerRule{} }
func (m *TCPLoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*TCPLoadBalancerRule) ProtoMessage()               {}
func (*TCPLoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *TCPLoadBalancerRule) GetBackend() *LoadBalancerBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "loadbalancer.v1beta1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "loadbalancer.v1beta1.ListResponse")
	proto.RegisterType((*DescribeRequest)(nil), "loadbalancer.v1beta1.DescribeRequest")
	proto.RegisterType((*DescribeResponse)(nil), "loadbalancer.v1beta1.DescribeResponse")
	proto.RegisterType((*CreateRequest)(nil), "loadbalancer.v1beta1.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "loadbalancer.v1beta1.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "loadbalancer.v1beta1.DeleteRequest")
	proto.RegisterType((*LoadBalancer)(nil), "loadbalancer.v1beta1.LoadBalancer")
	proto.RegisterType((*Spec)(nil), "loadbalancer.v1beta1.Spec")
	proto.RegisterType((*Status)(nil), "loadbalancer.v1beta1.Status")
	proto.RegisterType((*LoadBalancerStatus)(nil), "loadbalancer.v1beta1.LoadBalancerStatus")
	proto.RegisterType((*LoadBalancerBackend)(nil), "loadbalancer.v1beta1.LoadBalancerBackend")
	proto.RegisterType((*LoadBalancerRule)(nil), "loadbalancer.v1beta1.LoadBalancerRule")
	proto.RegisterType((*HTTPLoadBalancerRule)(nil), "loadbalancer.v1beta1.HTTPLoadBalancerRule")
	proto.RegisterType((*TCPLoadBalancerRule)(nil), "loadbalancer.v1beta1.TCPLoadBalancerRule")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for LoadBalancers service

type LoadBalancersClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type loadBalancersClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancersClient(cc *grpc.ClientConn) LoadBalancersClient {
	return &loadBalancersClient{cc}
}

func (c *loadBalancersClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/loadbalancer.v1beta1.LoadBalancers/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error) {
	out := new(DescribeResponse)
	err := grpc.Invoke(ctx, "/loadbalancer.v1beta1.LoadBalancers/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/loadbalancer.v1beta1.LoadBalancers/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/loadbalancer.v1beta1.LoadBalancers/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/loadbalancer.v1beta1.LoadBalancers/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoadBalancers service

type LoadBalancersServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Describe(context.Context, *DescribeRequest) (*DescribeResponse, error)
	Create(context.Context, *CreateRequest) (*dtypes.VoidResponse, error)
	Update(context.Context, *UpdateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *DeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterLoadBalancersServer(s *grpc.Server, srv LoadBalancersServer) {
	s.RegisterService(&_LoadBalancers_serviceDesc, srv)
}

func _LoadBalancers_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loadbalancer.v1beta1.LoadBalancers/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loadbalancer.v1beta1.LoadBalancers/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Describe(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loadbalancer.v1beta1.LoadBalancers/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loadbalancer.v1beta1.LoadBalancers/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loadbalancer.v1beta1.LoadBalancers/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadBalancers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loadbalancer.v1beta1.LoadBalancers",
	HandlerType: (*LoadBalancersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _LoadBalancers_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _LoadBalancers_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LoadBalancers_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LoadBalancers_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LoadBalancers_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("loadbalancer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 955 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x96, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0x35, 0xde, 0x8d, 0xdd, 0x3c, 0xdb, 0x6d, 0x98, 0x46, 0xc8, 0xb2, 0x22, 0xd1, 0x2c,
	0xc2, 0xa4, 0x91, 0xf0, 0xaa, 0x81, 0x43, 0x15, 0x10, 0x82, 0x24, 0x40, 0x53, 0x19, 0xb0, 0xec,
	0xc0, 0x01, 0x0e, 0xd6, 0x78, 0xfd, 0xd4, 0x2c, 0x59, 0xef, 0x6c, 0x77, 0xc6, 0x41, 0x51, 0x55,
	0x21, 0x01, 0xe2, 0xc0, 0x0d, 0x21, 0x24, 0x6e, 0x5c, 0xb8, 0x71, 0xe4, 0xce, 0x3f, 0xc0, 0x91,
	0x7f, 0x81, 0x3f, 0x04, 0xcd, 0x8f, 0x8d, 0x77, 0x9b, 0x75, 0x0d, 0x49, 0xc4, 0x25, 0x99, 0x7d,
	0xf3, 0xde, 0xcc, 0xe7, 0xfb, 0xde, 0xcc, 0xf8, 0x01, 0x8d, 0x38, 0x9b, 0x8c, 0x59, 0xc4, 0xe2,
	0x00, 0xd3, 0x6e, 0x92, 0x72, 0xc9, 0xe9, 0x7a, 0xc1, 0x76, 0x7a, 0x6f, 0x8c, 0x92, 0xdd, 0x6b,
	0x6f, 0x3c, 0xe2, 0xfc, 0x51, 0x84, 0x3e, 0x4b, 0x42, 0x9f, 0xc5, 0x31, 0x97, 0x4c, 0x86, 0x3c,
	0x16, 0x26, 0xa6, 0xfd, 0xa2, 0x32, 0x4f, 0xe4, 0x59, 0x82, 0xc2, 0xd7, 0x7f, 0x8d, 0xdd, 0x7b,
	0x15, 0xea, 0xbd, 0x50, 0xc8, 0x01, 0x3e, 0x9e, 0xa1, 0x90, 0xb4, 0x05, 0xb5, 0x20, 0x9a, 0x09,
	0x89, 0x69, 0x8b, 0xdc, 0x21, 0x5b, 0xab, 0x83, 0xec, 0xd3, 0xfb, 0x0a, 0x1a, 0xc6, 0x51, 0x24,
	0x3c, 0x16, 0x48, 0x3b, 0x50, 0x15, 0x92, 0xc9, 0x99, 0xd0, 0x8e, 0xf5, 0x9d, 0x9b, 0x5d, 0xb3,
	0x7a, 0x77, 0xa8, 0xad, 0x03, 0x3b, 0x4b, 0x3f, 0x80, 0xa6, 0xc2, 0x1d, 0x65, 0xbc, 0xad, 0xca,
	0x1d, 0x67, 0xab, 0xbe, 0xe3, 0x75, 0xcb, 0x44, 0x74, 0x7b, 0x9c, 0x4d, 0xf6, 0xac, 0x71, 0xd0,
	0x88, 0x72, 0x5f, 0xde, 0x63, 0xb8, 0x75, 0x80, 0x22, 0x48, 0xc3, 0x31, 0x66, 0xb4, 0x14, 0xdc,
	0x93, 0x30, 0x9e, 0x58, 0x54, 0x3d, 0x56, 0xb6, 0x98, 0x4d, 0xb1, 0x55, 0x31, 0x36, 0x35, 0xa6,
	0x1b, 0xb0, 0xaa, 0xfe, 0x8b, 0x84, 0x05, 0xd8, 0x72, 0xf4, 0xc4, 0xdc, 0x90, 0xd7, 0xec, 0x16,
	0x35, 0x7f, 0x43, 0x60, 0x6d, 0xbe, 0xe7, 0xd5, 0x85, 0x93, 0x4b, 0x09, 0xff, 0x95, 0x40, 0x73,
	0x3f, 0x45, 0x26, 0xf3, 0xba, 0xb5, 0x46, 0xb2, 0x48, 0x63, 0xe5, 0x39, 0x1a, 0x9d, 0x82, 0xc6,
	0x8b, 0x98, 0xee, 0x25, 0x31, 0xbf, 0x23, 0xd0, 0xfc, 0x24, 0x99, 0x2c, 0xc1, 0xcc, 0x81, 0x54,
	0x96, 0x80, 0x38, 0x97, 0x04, 0xe1, 0xd0, 0x3c, 0xc0, 0x08, 0xe5, 0xff, 0x76, 0x4c, 0xbe, 0x75,
	0xa0, 0x91, 0xe7, 0xb9, 0xa6, 0x0d, 0x5f, 0x03, 0x1a, 0xa8, 0xb2, 0x87, 0x3c, 0x1e, 0xc9, 0x70,
	0x8a, 0x42, 0xb2, 0x69, 0xa2, 0xf7, 0x76, 0x06, 0x2f, 0x64, 0x33, 0x47, 0xd9, 0x04, 0x3d, 0x84,
	0x1a, 0x4f, 0xf4, 0x95, 0x6f, 0xad, 0xe8, 0x2b, 0xe6, 0x2f, 0xcf, 0x5c, 0xf7, 0x63, 0x13, 0xf1,
	0x5e, 0x2c, 0xd3, 0xb3, 0x41, 0x16, 0x4f, 0xbb, 0xe0, 0x8a, 0x04, 0x83, 0x56, 0x55, 0x57, 0xa0,
	0x5d, 0xbe, 0xce, 0x30, 0xc1, 0x60, 0xa0, 0xfd, 0xe8, 0x1b, 0xe7, 0x57, 0xa2, 0xa6, 0x23, 0x36,
	0x16, 0x44, 0x14, 0x2f, 0x08, 0x05, 0xf7, 0x0b, 0xc1, 0xe3, 0xd6, 0x0d, 0x93, 0x11, 0x35, 0x6e,
	0xef, 0x42, 0x23, 0x8f, 0x44, 0xd7, 0xc0, 0x39, 0xc1, 0x33, 0x9b, 0x48, 0x35, 0xa4, 0xeb, 0xb0,
	0x72, 0xca, 0xa2, 0x59, 0x96, 0x48, 0xf3, 0xb1, 0x5b, 0xb9, 0x4f, 0xbc, 0xef, 0x09, 0xb8, 0x0a,
	0x8a, 0x1e, 0x40, 0x6d, 0xcc, 0x82, 0x13, 0xb4, 0x15, 0xa8, 0xef, 0x6c, 0x97, 0xf3, 0x3c, 0x38,
	0x3a, 0xea, 0x17, 0xce, 0xd1, 0x2c, 0xc2, 0x41, 0x16, 0x4a, 0xdf, 0x82, 0x95, 0x74, 0x16, 0xa1,
	0xb0, 0x0f, 0x56, 0xe7, 0x5f, 0x9c, 0x43, 0x15, 0x6f, 0x82, 0xbc, 0x87, 0x50, 0x35, 0x72, 0xe9,
	0x3b, 0xb9, 0xf7, 0x42, 0x2d, 0xb4, 0xb5, 0x7c, 0xa1, 0x62, 0xa2, 0xbc, 0xfb, 0x40, 0x2f, 0xce,
	0xd2, 0x9b, 0x50, 0x39, 0xec, 0xdb, 0xcc, 0x54, 0xc2, 0xbe, 0x4a, 0xe7, 0x31, 0x17, 0x32, 0x3b,
	0x60, 0x6a, 0xec, 0x7d, 0x0e, 0xb7, 0xf3, 0x91, 0x7b, 0x56, 0xda, 0x26, 0x34, 0x04, 0xa6, 0xa7,
	0x61, 0x80, 0xa3, 0xdc, 0x05, 0xad, 0x5b, 0xdb, 0x47, 0xea, 0x68, 0xe6, 0x5c, 0x12, 0x9e, 0x66,
	0xab, 0x66, 0x2e, 0x7d, 0x9e, 0x4a, 0xef, 0x4f, 0x02, 0x6b, 0xcf, 0xca, 0x3f, 0xa7, 0x20, 0x73,
	0x0a, 0xda, 0x81, 0x5b, 0xc3, 0x61, 0x6f, 0x24, 0x30, 0x48, 0x51, 0x9a, 0x1d, 0x57, 0xf4, 0x74,
	0x53, 0x0c, 0x7b, 0x43, 0x6d, 0xd5, 0x7b, 0xbe, 0x0d, 0xee, 0xb1, 0x94, 0x89, 0x4d, 0xf8, 0x7f,
	0x29, 0x9a, 0x8e, 0xa3, 0x6f, 0x82, 0x23, 0x83, 0xa4, 0xe5, 0xe8, 0xf0, 0xbb, 0xe5, 0xe1, 0x47,
	0xfb, 0x17, 0xa3, 0x55, 0x94, 0xf7, 0x3b, 0x81, 0xf5, 0xb2, 0xb5, 0x95, 0xa2, 0x84, 0xc9, 0xe3,
	0x4c, 0x91, 0x1a, 0xd3, 0xfd, 0xf9, 0x09, 0x33, 0xaf, 0xfa, 0xdd, 0xe5, 0x45, 0xb5, 0xc9, 0x9f,
	0x1f, 0xb0, 0x97, 0xa0, 0x7e, 0x8c, 0x6c, 0x82, 0xe9, 0x48, 0x1d, 0x19, 0x8d, 0xbd, 0x3a, 0x00,
	0x63, 0xd2, 0x3b, 0x6f, 0x42, 0x23, 0xc5, 0x2f, 0xd3, 0x50, 0xa2, 0xf1, 0x70, 0xb5, 0x47, 0xdd,
	0xda, 0x94, 0x8b, 0xf7, 0x07, 0x81, 0xdb, 0x25, 0x92, 0x34, 0xb4, 0x2a, 0x5b, 0x06, 0xcd, 0x53,
	0x79, 0x3d, 0xd0, 0x25, 0xb5, 0x74, 0xca, 0x6a, 0xd9, 0x81, 0x5b, 0xd6, 0x27, 0xc1, 0xa9, 0xf1,
	0x73, 0xad, 0x9f, 0x36, 0xf7, 0x71, 0xaa, 0xfc, 0x76, 0x7e, 0xa9, 0x42, 0x33, 0xbf, 0xa1, 0xa0,
	0x3f, 0x13, 0x70, 0x55, 0xa7, 0x41, 0x37, 0x17, 0xe0, 0xcd, 0xdb, 0x95, 0xb6, 0xf7, 0x3c, 0x17,
	0xf3, 0x7b, 0xed, 0x1d, 0x7e, 0xfd, 0xd7, 0xdf, 0x3f, 0x56, 0xf6, 0xe9, 0xbb, 0x3e, 0x4b, 0x12,
	0x11, 0xf0, 0x89, 0x69, 0x91, 0x4e, 0x66, 0x63, 0x4c, 0x63, 0x94, 0x28, 0x7c, 0x1b, 0xe6, 0xdb,
	0xf7, 0x5c, 0xf8, 0x4f, 0xec, 0xe8, 0xa9, 0x9f, 0x5f, 0x5e, 0xd0, 0xdf, 0x08, 0xdc, 0xc8, 0xfa,
	0x01, 0xfa, 0x4a, 0xf9, 0xde, 0xcf, 0xf4, 0x28, 0xed, 0xce, 0x32, 0x37, 0x8b, 0xd9, 0xd7, 0x98,
	0x0f, 0xe9, 0x83, 0x2b, 0x63, 0xfa, 0x4f, 0x54, 0xba, 0x9f, 0xd2, 0x1f, 0x08, 0x54, 0x4d, 0xdf,
	0x40, 0x5f, 0x2e, 0x87, 0x28, 0x74, 0x15, 0xed, 0xf5, 0xac, 0x91, 0xf9, 0x94, 0x87, 0x93, 0x73,
	0xae, 0x9e, 0xe6, 0x7a, 0xdf, 0xbb, 0x7a, 0xfa, 0x76, 0xc9, 0xb6, 0x2a, 0x6e, 0xd5, 0x34, 0x09,
	0x8b, 0x98, 0x0a, 0x2d, 0xc4, 0x02, 0xa6, 0xa1, 0x66, 0xfa, 0xb0, 0x7d, 0x6d, 0xb9, 0x52, 0x68,
	0x3f, 0x11, 0xa8, 0x9a, 0xbe, 0x61, 0x11, 0x5a, 0xa1, 0xab, 0x58, 0x80, 0x66, 0xcb, 0xb8, 0x7d,
	0x6d, 0x68, 0x7b, 0xab, 0x9f, 0xd5, 0x6c, 0xc8, 0xb8, 0xaa, 0x7b, 0xf6, 0xd7, 0xff, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x67, 0x02, 0x6c, 0xe8, 0x15, 0x0c, 0x00, 0x00,
}