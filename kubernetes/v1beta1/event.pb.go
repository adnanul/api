// Code generated by protoc-gen-go.
// source: event.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
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
	Kind       string                           `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Parents    map[string]*EventRequest_Parents `protobuf:"bytes,9,rep,name=parents" json:"parents,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Namespace  string                           `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
	PodIp      string                           `protobuf:"bytes,5,opt,name=pod_ip,json=podIp" json:"pod_ip,omitempty"`
	InstanceId string                           `protobuf:"bytes,6,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	InstanceIp string                           `protobuf:"bytes,7,opt,name=instance_ip,json=instanceIp" json:"instance_ip,omitempty"`
	Labels     map[string]string                `protobuf:"bytes,8,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EventRequest_ObjectMeta) Reset()                    { *m = EventRequest_ObjectMeta{} }
func (m *EventRequest_ObjectMeta) String() string            { return proto.CompactTextString(m) }
func (*EventRequest_ObjectMeta) ProtoMessage()               {}
func (*EventRequest_ObjectMeta) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

func (m *EventRequest_ObjectMeta) GetParents() map[string]*EventRequest_Parents {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *EventRequest_ObjectMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type EventRequest_Parents struct {
	Name []string `protobuf:"bytes,1,rep,name=name" json:"name,omitempty"`
}

func (m *EventRequest_Parents) Reset()                    { *m = EventRequest_Parents{} }
func (m *EventRequest_Parents) String() string            { return proto.CompactTextString(m) }
func (*EventRequest_Parents) ProtoMessage()               {}
func (*EventRequest_Parents) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 1} }

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
	proto.RegisterType((*EventRequest_Parents)(nil), "kubernetes.v1beta1.EventRequest.Parents")
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
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0xc6, 0x49, 0xc6, 0x69, 0x55, 0xad, 0xa0, 0xb2, 0xac, 0x16, 0xd2, 0x4a, 0xa0,
	0x00, 0xc2, 0x56, 0xc3, 0x81, 0xd2, 0x03, 0x48, 0x40, 0x05, 0x95, 0xa0, 0xad, 0x4c, 0xb9, 0x70,
	0x09, 0x1b, 0x7b, 0x5a, 0x99, 0xa4, 0xeb, 0xc5, 0xbb, 0x8e, 0x88, 0x10, 0x97, 0x9e, 0xb9, 0xa0,
	0xfe, 0x15, 0xe0, 0x8f, 0x70, 0xe6, 0xc6, 0x8f, 0xe0, 0x88, 0x76, 0xbd, 0x69, 0xdc, 0x16, 0xa8,
	0x2a, 0xb8, 0x58, 0xe3, 0xb7, 0xf3, 0x66, 0xe6, 0xbd, 0xfd, 0x00, 0x07, 0x47, 0xc8, 0xa4, 0xcf,
	0xb3, 0x54, 0xa6, 0x84, 0x0c, 0xf2, 0x3e, 0x66, 0x0c, 0x25, 0x0a, 0x7f, 0xb4, 0xda, 0x47, 0x49,
	0x57, 0xbd, 0xc5, 0xfd, 0x34, 0xdd, 0x1f, 0x62, 0x40, 0x79, 0x12, 0x50, 0xc6, 0x52, 0x49, 0x65,
	0x92, 0x32, 0x51, 0x30, 0xbc, 0x2b, 0x94, 0x73, 0x11, 0xa5, 0xf1, 0x9f, 0xd6, 0x17, 0x14, 0x1c,
	0xcb, 0x31, 0x47, 0x11, 0xe8, 0x6f, 0x81, 0xaf, 0x1c, 0xd9, 0xd0, 0xda, 0x50, 0x9d, 0x43, 0x7c,
	0x9b, 0xa3, 0x90, 0x64, 0x19, 0x5a, 0xd1, 0x30, 0x17, 0x12, 0xb3, 0x1e, 0xa3, 0x07, 0xe8, 0x5a,
	0x6d, 0xab, 0xd3, 0x0c, 0x1d, 0x83, 0x6d, 0xd1, 0x03, 0x24, 0xd7, 0x60, 0x4e, 0xcd, 0xa7, 0xd7,
	0x05, 0xa7, 0x11, 0xba, 0x15, 0x9d, 0x34, 0xab, 0xd0, 0xad, 0x09, 0x48, 0x3a, 0x30, 0xaf, 0xd3,
	0xd2, 0xfe, 0x1b, 0x8c, 0x64, 0x4f, 0x75, 0x75, 0xab, 0x3a, 0x51, 0xd3, 0xb7, 0x35, 0xbc, 0x3b,
	0xe6, 0x67, 0x32, 0x75, 0xdf, 0x99, 0xd3, 0x99, 0xba, 0xf5, 0x12, 0x80, 0xf6, 0xa9, 0xa8, 0x56,
	0xd3, 0x39, 0x4d, 0x8d, 0xe8, 0x42, 0x4f, 0xa0, 0x71, 0x80, 0x92, 0xc6, 0x54, 0x52, 0xb7, 0xde,
	0xb6, 0x3a, 0x4e, 0xf7, 0x96, 0x7f, 0xd6, 0x4a, 0xbf, 0x2c, 0xd8, 0x2f, 0xaa, 0x3f, 0x47, 0x49,
	0xc3, 0x63, 0xb2, 0xf7, 0xbd, 0x0a, 0x30, 0x5d, 0x20, 0x04, 0x66, 0x06, 0x09, 0x8b, 0x8d, 0x19,
	0x3a, 0x26, 0x21, 0xd4, 0x39, 0xcd, 0x90, 0x49, 0xe1, 0x36, 0xdb, 0xd5, 0x8e, 0xd3, 0x5d, 0xbb,
	0x40, 0x2b, 0x7f, 0xa7, 0xa0, 0x6e, 0x30, 0x99, 0x8d, 0xc3, 0x49, 0x21, 0xb2, 0x08, 0xcd, 0xa9,
	0xa9, 0x85, 0x03, 0x53, 0x80, 0x5c, 0x06, 0x9b, 0xa7, 0x71, 0x2f, 0xe1, 0x46, 0x78, 0x8d, 0xa7,
	0xf1, 0x26, 0x27, 0x57, 0xc1, 0x49, 0x98, 0x90, 0x94, 0x45, 0xd8, 0x4b, 0x62, 0xd7, 0xd6, 0x6b,
	0x30, 0x81, 0x36, 0xe3, 0x93, 0x09, 0x5c, 0x1b, 0x53, 0x4e, 0xe0, 0x64, 0x1b, 0xec, 0x21, 0xed,
	0xe3, 0x50, 0xb8, 0x0d, 0xad, 0xe4, 0xee, 0x45, 0x94, 0x3c, 0xd3, 0xcc, 0x42, 0x88, 0x29, 0xe3,
	0xc5, 0xd0, 0x2a, 0x0b, 0x24, 0xf3, 0x50, 0x1d, 0xe0, 0xd8, 0xd8, 0xa7, 0x42, 0x72, 0x1f, 0x6a,
	0x23, 0x3a, 0xcc, 0x8b, 0xa3, 0xe3, 0x74, 0x3b, 0xe7, 0x76, 0x34, 0xf5, 0xc2, 0x82, 0xb6, 0x5e,
	0x59, 0xb3, 0xbc, 0x7b, 0xe0, 0x94, 0x9a, 0xff, 0xa6, 0xc9, 0xa5, 0x72, 0x93, 0x66, 0x99, 0xba,
	0x04, 0x75, 0x53, 0x50, 0xed, 0xad, 0x39, 0xe8, 0x55, 0xb5, 0xb7, 0x2a, 0x5e, 0xf9, 0x58, 0x81,
	0x59, 0xd3, 0x5d, 0xf0, 0x94, 0x09, 0x24, 0xd7, 0xc1, 0x16, 0x92, 0xca, 0x5c, 0xe8, 0xfa, 0x4e,
	0x77, 0xce, 0x2f, 0x2e, 0x93, 0xff, 0x42, 0xa3, 0xa1, 0x59, 0x25, 0x0f, 0xa0, 0x26, 0xa9, 0x18,
	0x08, 0xb7, 0xa2, 0x9d, 0xbc, 0xf1, 0x17, 0x5d, 0x45, 0x65, 0xff, 0x29, 0x65, 0xf1, 0x10, 0xb3,
	0xb0, 0xe0, 0x79, 0x9f, 0x2c, 0xa8, 0x1b, 0x88, 0x2c, 0x80, 0x4d, 0x23, 0x75, 0x8b, 0x8d, 0x28,
	0xf3, 0x47, 0x3c, 0x68, 0x64, 0x18, 0x61, 0x32, 0xc2, 0xcc, 0x48, 0x3b, 0xfe, 0x57, 0x72, 0xfa,
	0x54, 0x4c, 0x6e, 0x9a, 0x8e, 0x55, 0x1d, 0x91, 0xef, 0xed, 0x25, 0xef, 0xcc, 0x99, 0x32, 0x7f,
	0x2a, 0x77, 0x84, 0x59, 0xdf, 0x1c, 0x27, 0x1d, 0x2b, 0x4c, 0x5f, 0x1f, 0x75, 0x8c, 0x5a, 0xa1,
	0x8e, 0xbb, 0x3f, 0x2b, 0x60, 0xeb, 0xa1, 0x05, 0xf9, 0x6a, 0x41, 0xeb, 0x51, 0xca, 0x84, 0xcc,
	0xf2, 0x48, 0x26, 0x23, 0x24, 0xed, 0xf3, 0x76, 0xce, 0x5b, 0x3e, 0xd7, 0x83, 0x15, 0x7a, 0xf8,
	0xc5, 0xad, 0x34, 0xac, 0xc3, 0x6f, 0x3f, 0x8e, 0x2a, 0x2f, 0xbd, 0x9d, 0xe0, 0xc4, 0x63, 0x36,
	0x65, 0x07, 0x86, 0x1d, 0x98, 0xd7, 0x48, 0x04, 0xef, 0xcb, 0x6f, 0xd5, 0x87, 0xa0, 0xf0, 0x47,
	0x04, 0xd1, 0x64, 0xc2, 0xdb, 0xfa, 0x6d, 0x58, 0xb7, 0x6e, 0x92, 0xcf, 0x16, 0x38, 0x8f, 0xf1,
	0x3f, 0xcf, 0xfd, 0xba, 0x34, 0xf7, 0xae, 0xb7, 0xfd, 0xcf, 0x73, 0xc7, 0x78, 0x6a, 0xec, 0x87,
	0xcd, 0x57, 0x75, 0x43, 0xee, 0xdb, 0xfa, 0xc5, 0xbe, 0xf3, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xc6,
	0x83, 0x62, 0xaa, 0x2a, 0x06, 0x00, 0x00,
}
