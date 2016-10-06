// Code generated by protoc-gen-go.
// source: event.proto
// DO NOT EDIT!

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

type EventRequest struct {
	ClusterName    string                   `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	KubeNamespace  string                   `protobuf:"bytes,2,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	KubeObjectType string                   `protobuf:"bytes,3,opt,name=kube_object_type,json=kubeObjectType" json:"kube_object_type,omitempty"`
	KubeObjectName string                   `protobuf:"bytes,4,opt,name=kube_object_name,json=kubeObjectName" json:"kube_object_name,omitempty"`
	EventType      string                   `protobuf:"bytes,5,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	Metadata       *EventRequest_ObjectMeta `protobuf:"bytes,7,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *EventRequest) Reset()                    { *m = EventRequest{} }
func (m *EventRequest) String() string            { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()               {}
func (*EventRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *EventRequest) GetMetadata() *EventRequest_ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type EventRequest_ObjectMeta struct {
	Kind                  string            `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Service               []string          `protobuf:"bytes,2,rep,name=service" json:"service,omitempty"`
	ReplicationController string            `protobuf:"bytes,3,opt,name=replication_controller,json=replicationController" json:"replication_controller,omitempty"`
	Namespace             string            `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
	PodIp                 string            `protobuf:"bytes,5,opt,name=pod_ip,json=podIp" json:"pod_ip,omitempty"`
	InstanceId            string            `protobuf:"bytes,6,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	InstanceIp            string            `protobuf:"bytes,7,opt,name=instance_ip,json=instanceIp" json:"instance_ip,omitempty"`
	Labels                map[string]string `protobuf:"bytes,8,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EventRequest_ObjectMeta) Reset()                    { *m = EventRequest_ObjectMeta{} }
func (m *EventRequest_ObjectMeta) String() string            { return proto.CompactTextString(m) }
func (*EventRequest_ObjectMeta) ProtoMessage()               {}
func (*EventRequest_ObjectMeta) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

func (m *EventRequest_ObjectMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type EventResponse struct {
	Status *dtypes.Status           `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Tasks  []*EventResponse_Handler `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *EventResponse) Reset()                    { *m = EventResponse{} }
func (m *EventResponse) String() string            { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()               {}
func (*EventResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *EventResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *EventResponse) GetTasks() []*EventResponse_Handler {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type EventResponse_Handler struct {
	Action   string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Receiver string `protobuf:"bytes,2,opt,name=receiver" json:"receiver,omitempty"`
	Base     string `protobuf:"bytes,3,opt,name=base" json:"base,omitempty"`
	Suffix   string `protobuf:"bytes,4,opt,name=suffix" json:"suffix,omitempty"`
	Verb     string `protobuf:"bytes,5,opt,name=verb" json:"verb,omitempty"`
	Data     []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EventResponse_Handler) Reset()                    { *m = EventResponse_Handler{} }
func (m *EventResponse_Handler) String() string            { return proto.CompactTextString(m) }
func (*EventResponse_Handler) ProtoMessage()               {}
func (*EventResponse_Handler) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

func init() {
	proto.RegisterType((*EventRequest)(nil), "kubernetes.v1beta1.EventRequest")
	proto.RegisterType((*EventRequest_ObjectMeta)(nil), "kubernetes.v1beta1.EventRequest.ObjectMeta")
	proto.RegisterType((*EventResponse)(nil), "kubernetes.v1beta1.EventResponse")
	proto.RegisterType((*EventResponse_Handler)(nil), "kubernetes.v1beta1.EventResponse.Handler")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Events service

type EventsClient interface {
	Constructive(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	Destructive(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Constructive(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := grpc.Invoke(ctx, "/kubernetes.v1beta1.Events/Constructive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) Destructive(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := grpc.Invoke(ctx, "/kubernetes.v1beta1.Events/Destructive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Events service

type EventsServer interface {
	Constructive(context.Context, *EventRequest) (*EventResponse, error)
	Destructive(context.Context, *EventRequest) (*EventResponse, error)
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Constructive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Constructive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.v1beta1.Events/Constructive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Constructive(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_Destructive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Destructive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.v1beta1.Events/Destructive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Destructive(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kubernetes.v1beta1.Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Constructive",
			Handler:    _Events_Constructive_Handler,
		},
		{
			MethodName: "Destructive",
			Handler:    _Events_Destructive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto.RegisterFile("event.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6b, 0x13, 0x5b,
	0x14, 0x67, 0x92, 0x26, 0x69, 0xce, 0xa4, 0xa5, 0x5c, 0x5e, 0x4b, 0x18, 0xfa, 0x78, 0x69, 0xe1,
	0x3d, 0xf2, 0x14, 0x67, 0x68, 0x44, 0xd4, 0x6e, 0x04, 0x6b, 0xd1, 0x82, 0xb6, 0x32, 0xd6, 0x4d,
	0x5d, 0x84, 0x9b, 0x99, 0xd3, 0x32, 0x66, 0x7a, 0xef, 0x75, 0xee, 0x9d, 0xc1, 0x20, 0x6e, 0xdc,
	0xea, 0x46, 0xfc, 0x1e, 0xe2, 0x77, 0xf1, 0x2b, 0xb8, 0x74, 0xe9, 0x07, 0x90, 0x39, 0x73, 0xd3,
	0xc4, 0x16, 0x2c, 0xa2, 0x9b, 0xe1, 0x9c, 0xdf, 0xf9, 0xff, 0x3b, 0x67, 0x2e, 0xb8, 0x58, 0xa0,
	0x30, 0xbe, 0xca, 0xa4, 0x91, 0x8c, 0x8d, 0xf3, 0x11, 0x66, 0x02, 0x0d, 0x6a, 0xbf, 0xd8, 0x1a,
	0xa1, 0xe1, 0x5b, 0xde, 0xfa, 0x89, 0x94, 0x27, 0x29, 0x06, 0x5c, 0x25, 0x01, 0x17, 0x42, 0x1a,
	0x6e, 0x12, 0x29, 0x74, 0x15, 0xe1, 0xad, 0x95, 0x70, 0x6c, 0x26, 0x0a, 0x75, 0x40, 0xdf, 0x0a,
	0xdf, 0x7c, 0xdb, 0x80, 0xce, 0x6e, 0x99, 0x39, 0xc4, 0x17, 0x39, 0x6a, 0xc3, 0x36, 0xa0, 0x13,
	0xa5, 0xb9, 0x36, 0x98, 0x0d, 0x05, 0x3f, 0xc5, 0xae, 0xd3, 0x73, 0xfa, 0xed, 0xd0, 0xb5, 0xd8,
	0x3e, 0x3f, 0x45, 0xf6, 0x2f, 0x2c, 0x97, 0xf5, 0xc9, 0xae, 0x15, 0x8f, 0xb0, 0x5b, 0x23, 0xa7,
	0xa5, 0x12, 0xdd, 0x9f, 0x82, 0xac, 0x0f, 0x2b, 0xe4, 0x26, 0x47, 0xcf, 0x31, 0x32, 0xc3, 0xb2,
	0x6a, 0xb7, 0x4e, 0x8e, 0x14, 0x7e, 0x40, 0xf0, 0xe1, 0x44, 0x5d, 0xf0, 0xa4, 0xba, 0x0b, 0xe7,
	0x3d, 0xa9, 0xf4, 0xdf, 0x00, 0xc4, 0x43, 0x95, 0xad, 0x41, 0x3e, 0x6d, 0x42, 0x28, 0xd1, 0x7d,
	0x58, 0x3c, 0x45, 0xc3, 0x63, 0x6e, 0x78, 0xb7, 0xd5, 0x73, 0xfa, 0xee, 0xe0, 0xaa, 0x7f, 0x91,
	0x2a, 0x7f, 0x7e, 0x60, 0xbf, 0xca, 0xfe, 0x08, 0x0d, 0x0f, 0xcf, 0x82, 0xbd, 0x6f, 0x35, 0x80,
	0x99, 0x81, 0x31, 0x58, 0x18, 0x27, 0x22, 0xb6, 0x64, 0x90, 0xcc, 0xba, 0xd0, 0xd2, 0x98, 0x15,
	0x09, 0x8d, 0x5f, 0xef, 0xb7, 0xc3, 0xa9, 0xca, 0x6e, 0xc0, 0x5a, 0x86, 0x2a, 0x4d, 0x22, 0xda,
	0xc0, 0x30, 0x92, 0xc2, 0x64, 0x32, 0x4d, 0x31, 0xb3, 0xe3, 0xaf, 0xce, 0x59, 0x77, 0xce, 0x8c,
	0x6c, 0x1d, 0xda, 0x33, 0x46, 0xab, 0xf1, 0x67, 0x00, 0x5b, 0x85, 0xa6, 0x92, 0xf1, 0x30, 0x51,
	0x76, 0xea, 0x86, 0x92, 0xf1, 0x9e, 0x62, 0xff, 0x80, 0x9b, 0x08, 0x6d, 0xb8, 0x88, 0x70, 0x98,
	0xc4, 0xdd, 0x26, 0xd9, 0x60, 0x0a, 0xed, 0xc5, 0x3f, 0x3a, 0x28, 0x62, 0x65, 0xde, 0x41, 0xb1,
	0x03, 0x68, 0xa6, 0x7c, 0x84, 0xa9, 0xee, 0x2e, 0xf6, 0xea, 0x7d, 0x77, 0x70, 0xf3, 0x17, 0x18,
	0xf3, 0x1f, 0x52, 0xe4, 0xae, 0x30, 0xd9, 0x24, 0xb4, 0x69, 0xbc, 0xdb, 0xe0, 0xce, 0xc1, 0x6c,
	0x05, 0xea, 0x63, 0x9c, 0x58, 0xea, 0x4a, 0x91, 0xfd, 0x05, 0x8d, 0x82, 0xa7, 0xf9, 0xf4, 0x6c,
	0x2a, 0x65, 0xbb, 0x76, 0xcb, 0xd9, 0x7c, 0x57, 0x83, 0x25, 0x5b, 0x4a, 0x2b, 0x29, 0x34, 0xb2,
	0xff, 0xa0, 0xa9, 0x0d, 0x37, 0xb9, 0xa6, 0x04, 0xee, 0x60, 0xd9, 0xaf, 0x8e, 0xd8, 0x7f, 0x42,
	0x68, 0x68, 0xad, 0xec, 0x0e, 0x34, 0x0c, 0xd7, 0x63, 0x4d, 0xbb, 0x70, 0x07, 0xff, 0xff, 0x64,
	0x88, 0x2a, 0xb3, 0xff, 0x80, 0x8b, 0x38, 0xc5, 0x2c, 0xac, 0xe2, 0xbc, 0xf7, 0x0e, 0xb4, 0x2c,
	0xc4, 0xd6, 0xa0, 0xc9, 0xa3, 0x72, 0x3b, 0xb6, 0x6b, 0xab, 0x31, 0x0f, 0x16, 0x33, 0x8c, 0x30,
	0x29, 0x30, 0xb3, 0xbd, 0x9f, 0xe9, 0xe5, 0x89, 0x8c, 0xb8, 0x9e, 0x5e, 0x38, 0xc9, 0x65, 0x1e,
	0x9d, 0x1f, 0x1f, 0x27, 0x2f, 0xed, 0x3a, 0xad, 0x56, 0xfa, 0x16, 0x98, 0x8d, 0xec, 0x26, 0x49,
	0x2e, 0x31, 0x3a, 0xdb, 0x72, 0x83, 0x9d, 0x90, 0xe4, 0xc1, 0xd7, 0x1a, 0x34, 0xa9, 0x69, 0xcd,
	0x3e, 0x39, 0xd0, 0xd9, 0x91, 0x42, 0x9b, 0x2c, 0x8f, 0x4c, 0x52, 0x20, 0xeb, 0x5d, 0xb6, 0x26,
	0x6f, 0xe3, 0x52, 0x0e, 0x36, 0x9f, 0xbd, 0xf9, 0xfc, 0xe5, 0x43, 0xed, 0xa9, 0xf7, 0x38, 0xe0,
	0x4a, 0xe9, 0x48, 0xc6, 0xd5, 0xf3, 0x31, 0x8b, 0x0b, 0x6c, 0x5c, 0x60, 0xff, 0x7f, 0x1d, 0xbc,
	0x9a, 0x7f, 0x1d, 0x5e, 0x07, 0x15, 0x33, 0x3a, 0x88, 0xa6, 0xbd, 0x5d, 0xa3, 0xbf, 0x71, 0xdb,
	0xb9, 0xc2, 0x3e, 0x3a, 0xe0, 0xde, 0xc3, 0x3f, 0xdc, 0xf1, 0x11, 0x75, 0x7c, 0xe8, 0x1d, 0xfc,
	0x76, 0xc7, 0x31, 0x9e, 0x6b, 0xf8, 0x6e, 0xfb, 0xa8, 0x65, 0x83, 0x47, 0x4d, 0x7a, 0x1d, 0xaf,
	0x7f, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x66, 0x75, 0x44, 0x76, 0x05, 0x00, 0x00,
}