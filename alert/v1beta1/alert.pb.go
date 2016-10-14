// Code generated by protoc-gen-go.
// source: alert.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	alert.proto
	incident.proto

It has these top-level messages:
	IcingaParam
	IcingaState
	NotifierParam
	Alert
	AlertListRequest
	AlertListResponse
	AlertCreateRequest
	AlertDescribeRequest
	AlertDescribeResponse
	AlertUpdateRequest
	AlertSyncRequest
	AlertDeleteRequest
	Incident
	IncidentListRequest
	IncidentListResponse
	IncidentDescribeRequest
	IncidentDescribeResponse
	IncidentNotifyRequest
	IncidentEventCreateRequest
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IcingaParam struct {
	CheckIntervalSec int64 `protobuf:"varint,1,opt,name=check_interval_sec,json=checkIntervalSec" json:"check_interval_sec,omitempty"`
	AlertIntervalSec int64 `protobuf:"varint,2,opt,name=alert_interval_sec,json=alertIntervalSec" json:"alert_interval_sec,omitempty"`
}

func (m *IcingaParam) Reset()                    { *m = IcingaParam{} }
func (m *IcingaParam) String() string            { return proto.CompactTextString(m) }
func (*IcingaParam) ProtoMessage()               {}
func (*IcingaParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// This one is used in kubernetes/v1beta1/client.proto
// To send information if alert is set for app/pod
type IcingaState struct {
	OK       int32 `protobuf:"varint,1,opt,name=OK" json:"OK,omitempty"`
	Warning  int32 `protobuf:"varint,2,opt,name=Warning" json:"Warning,omitempty"`
	Critical int32 `protobuf:"varint,3,opt,name=Critical" json:"Critical,omitempty"`
	Unknown  int32 `protobuf:"varint,4,opt,name=Unknown" json:"Unknown,omitempty"`
}

func (m *IcingaState) Reset()                    { *m = IcingaState{} }
func (m *IcingaState) String() string            { return proto.CompactTextString(m) }
func (*IcingaState) ProtoMessage()               {}
func (*IcingaState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NotifierParam struct {
	State   string `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	UserUid string `protobuf:"bytes,2,opt,name=user_uid,json=userUid" json:"user_uid,omitempty"`
	Method  string `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
}

func (m *NotifierParam) Reset()                    { *m = NotifierParam{} }
func (m *NotifierParam) String() string            { return proto.CompactTextString(m) }
func (*NotifierParam) ProtoMessage()               {}
func (*NotifierParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Alert struct {
	Phid                 string            `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	KubernetesCluster    string            `protobuf:"bytes,2,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string            `protobuf:"bytes,3,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string            `protobuf:"bytes,4,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string            `protobuf:"bytes,5,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	IcingaService        string            `protobuf:"bytes,7,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	IcingaParam          *IcingaParam      `protobuf:"bytes,8,opt,name=icinga_param,json=icingaParam" json:"icinga_param,omitempty"`
	CheckCommand         string            `protobuf:"bytes,9,opt,name=check_command,json=checkCommand" json:"check_command,omitempty"`
	NotifierParam        []*NotifierParam  `protobuf:"bytes,11,rep,name=notifier_param,json=notifierParam" json:"notifier_param,omitempty"`
	IcingawebUrl         string            `protobuf:"bytes,12,opt,name=icingaweb_url,json=icingawebUrl" json:"icingaweb_url,omitempty"`
	Vars                 map[string]string `protobuf:"bytes,13,rep,name=vars" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Alert) Reset()                    { *m = Alert{} }
func (m *Alert) String() string            { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()               {}
func (*Alert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Alert) GetIcingaParam() *IcingaParam {
	if m != nil {
		return m.IcingaParam
	}
	return nil
}

func (m *Alert) GetNotifierParam() []*NotifierParam {
	if m != nil {
		return m.NotifierParam
	}
	return nil
}

func (m *Alert) GetVars() map[string]string {
	if m != nil {
		return m.Vars
	}
	return nil
}

// Next Id: 5
type AlertListRequest struct {
	KubernetesCluster    string `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string `protobuf:"bytes,3,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string `protobuf:"bytes,4,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
}

func (m *AlertListRequest) Reset()                    { *m = AlertListRequest{} }
func (m *AlertListRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertListRequest) ProtoMessage()               {}
func (*AlertListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type AlertListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Alerts []*Alert       `protobuf:"bytes,2,rep,name=alerts" json:"alerts,omitempty"`
}

func (m *AlertListResponse) Reset()                    { *m = AlertListResponse{} }
func (m *AlertListResponse) String() string            { return proto.CompactTextString(m) }
func (*AlertListResponse) ProtoMessage()               {}
func (*AlertListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AlertListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AlertListResponse) GetAlerts() []*Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

// Next Id: 12
type AlertCreateRequest struct {
	KubernetesCluster    string            `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string            `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string            `protobuf:"bytes,3,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string            `protobuf:"bytes,4,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	IcingaService        string            `protobuf:"bytes,6,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	IcingaParam          *IcingaParam      `protobuf:"bytes,7,opt,name=icinga_param,json=icingaParam" json:"icinga_param,omitempty"`
	CheckCommand         string            `protobuf:"bytes,8,opt,name=check_command,json=checkCommand" json:"check_command,omitempty"`
	NotifierParam        []*NotifierParam  `protobuf:"bytes,10,rep,name=notifier_param,json=notifierParam" json:"notifier_param,omitempty"`
	Vars                 map[string]string `protobuf:"bytes,11,rep,name=vars" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AlertCreateRequest) Reset()                    { *m = AlertCreateRequest{} }
func (m *AlertCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertCreateRequest) ProtoMessage()               {}
func (*AlertCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AlertCreateRequest) GetIcingaParam() *IcingaParam {
	if m != nil {
		return m.IcingaParam
	}
	return nil
}

func (m *AlertCreateRequest) GetNotifierParam() []*NotifierParam {
	if m != nil {
		return m.NotifierParam
	}
	return nil
}

func (m *AlertCreateRequest) GetVars() map[string]string {
	if m != nil {
		return m.Vars
	}
	return nil
}

// Next Id: 2
type AlertDescribeRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *AlertDescribeRequest) Reset()                    { *m = AlertDescribeRequest{} }
func (m *AlertDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertDescribeRequest) ProtoMessage()               {}
func (*AlertDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AlertDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Alerts *Alert         `protobuf:"bytes,2,opt,name=alerts" json:"alerts,omitempty"`
}

func (m *AlertDescribeResponse) Reset()                    { *m = AlertDescribeResponse{} }
func (m *AlertDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*AlertDescribeResponse) ProtoMessage()               {}
func (*AlertDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AlertDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AlertDescribeResponse) GetAlerts() *Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

// Next Id: 6
type AlertUpdateRequest struct {
	Phid          string            `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	IcingaParam   *IcingaParam      `protobuf:"bytes,2,opt,name=icinga_param,json=icingaParam" json:"icinga_param,omitempty"`
	NotifierParam []*NotifierParam  `protobuf:"bytes,4,rep,name=notifier_param,json=notifierParam" json:"notifier_param,omitempty"`
	Vars          map[string]string `protobuf:"bytes,5,rep,name=vars" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AlertUpdateRequest) Reset()                    { *m = AlertUpdateRequest{} }
func (m *AlertUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertUpdateRequest) ProtoMessage()               {}
func (*AlertUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AlertUpdateRequest) GetIcingaParam() *IcingaParam {
	if m != nil {
		return m.IcingaParam
	}
	return nil
}

func (m *AlertUpdateRequest) GetNotifierParam() []*NotifierParam {
	if m != nil {
		return m.NotifierParam
	}
	return nil
}

func (m *AlertUpdateRequest) GetVars() map[string]string {
	if m != nil {
		return m.Vars
	}
	return nil
}

// Next Id: 6
type AlertSyncRequest struct {
	KubernetesCluster    string                          `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string                          `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string                          `protobuf:"bytes,3,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string                          `protobuf:"bytes,4,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	PodAncestor          []*AlertSyncRequest_PodAncestor `protobuf:"bytes,5,rep,name=pod_ancestor,json=podAncestor" json:"pod_ancestor,omitempty"`
}

func (m *AlertSyncRequest) Reset()                    { *m = AlertSyncRequest{} }
func (m *AlertSyncRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertSyncRequest) ProtoMessage()               {}
func (*AlertSyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AlertSyncRequest) GetPodAncestor() []*AlertSyncRequest_PodAncestor {
	if m != nil {
		return m.PodAncestor
	}
	return nil
}

// Next Id: 3
type AlertSyncRequest_PodAncestor struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *AlertSyncRequest_PodAncestor) Reset()         { *m = AlertSyncRequest_PodAncestor{} }
func (m *AlertSyncRequest_PodAncestor) String() string { return proto.CompactTextString(m) }
func (*AlertSyncRequest_PodAncestor) ProtoMessage()    {}
func (*AlertSyncRequest_PodAncestor) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

// Next Id: 2
type AlertDeleteRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *AlertDeleteRequest) Reset()                    { *m = AlertDeleteRequest{} }
func (m *AlertDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertDeleteRequest) ProtoMessage()               {}
func (*AlertDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*IcingaParam)(nil), "alert.v1beta1.IcingaParam")
	proto.RegisterType((*IcingaState)(nil), "alert.v1beta1.IcingaState")
	proto.RegisterType((*NotifierParam)(nil), "alert.v1beta1.NotifierParam")
	proto.RegisterType((*Alert)(nil), "alert.v1beta1.Alert")
	proto.RegisterType((*AlertListRequest)(nil), "alert.v1beta1.AlertListRequest")
	proto.RegisterType((*AlertListResponse)(nil), "alert.v1beta1.AlertListResponse")
	proto.RegisterType((*AlertCreateRequest)(nil), "alert.v1beta1.AlertCreateRequest")
	proto.RegisterType((*AlertDescribeRequest)(nil), "alert.v1beta1.AlertDescribeRequest")
	proto.RegisterType((*AlertDescribeResponse)(nil), "alert.v1beta1.AlertDescribeResponse")
	proto.RegisterType((*AlertUpdateRequest)(nil), "alert.v1beta1.AlertUpdateRequest")
	proto.RegisterType((*AlertSyncRequest)(nil), "alert.v1beta1.AlertSyncRequest")
	proto.RegisterType((*AlertSyncRequest_PodAncestor)(nil), "alert.v1beta1.AlertSyncRequest.PodAncestor")
	proto.RegisterType((*AlertDeleteRequest)(nil), "alert.v1beta1.AlertDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Alerts service

type AlertsClient interface {
	List(ctx context.Context, in *AlertListRequest, opts ...grpc.CallOption) (*AlertListResponse, error)
	Create(ctx context.Context, in *AlertCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Describe(ctx context.Context, in *AlertDescribeRequest, opts ...grpc.CallOption) (*AlertDescribeResponse, error)
	Update(ctx context.Context, in *AlertUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Sync(ctx context.Context, in *AlertSyncRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *AlertDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type alertsClient struct {
	cc *grpc.ClientConn
}

func NewAlertsClient(cc *grpc.ClientConn) AlertsClient {
	return &alertsClient{cc}
}

func (c *alertsClient) List(ctx context.Context, in *AlertListRequest, opts ...grpc.CallOption) (*AlertListResponse, error) {
	out := new(AlertListResponse)
	err := grpc.Invoke(ctx, "/alert.v1beta1.Alerts/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Create(ctx context.Context, in *AlertCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.v1beta1.Alerts/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Describe(ctx context.Context, in *AlertDescribeRequest, opts ...grpc.CallOption) (*AlertDescribeResponse, error) {
	out := new(AlertDescribeResponse)
	err := grpc.Invoke(ctx, "/alert.v1beta1.Alerts/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Update(ctx context.Context, in *AlertUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.v1beta1.Alerts/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Sync(ctx context.Context, in *AlertSyncRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.v1beta1.Alerts/Sync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Delete(ctx context.Context, in *AlertDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.v1beta1.Alerts/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Alerts service

type AlertsServer interface {
	List(context.Context, *AlertListRequest) (*AlertListResponse, error)
	Create(context.Context, *AlertCreateRequest) (*dtypes.VoidResponse, error)
	Describe(context.Context, *AlertDescribeRequest) (*AlertDescribeResponse, error)
	Update(context.Context, *AlertUpdateRequest) (*dtypes.VoidResponse, error)
	Sync(context.Context, *AlertSyncRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *AlertDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterAlertsServer(s *grpc.Server, srv AlertsServer) {
	s.RegisterService(&_Alerts_serviceDesc, srv)
}

func _Alerts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.v1beta1.Alerts/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).List(ctx, req.(*AlertListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.v1beta1.Alerts/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Create(ctx, req.(*AlertCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.v1beta1.Alerts/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Describe(ctx, req.(*AlertDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.v1beta1.Alerts/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Update(ctx, req.(*AlertUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.v1beta1.Alerts/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Sync(ctx, req.(*AlertSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.v1beta1.Alerts/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Delete(ctx, req.(*AlertDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alerts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alert.v1beta1.Alerts",
	HandlerType: (*AlertsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Alerts_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Alerts_Create_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Alerts_Describe_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Alerts_Update_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Alerts_Sync_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Alerts_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("alert.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1028 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xdc, 0x57, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x06, 0x69, 0x89, 0x96, 0x46, 0x96, 0xe1, 0x6c, 0x9d, 0x80, 0x25, 0x82, 0xd4, 0xa5, 0xfb,
	0xe3, 0xba, 0x8e, 0x58, 0xdb, 0x08, 0x1a, 0x04, 0x2d, 0x82, 0xd4, 0xe9, 0x21, 0x48, 0x61, 0x07,
	0x74, 0x9d, 0x14, 0xbd, 0x08, 0x2b, 0x72, 0x6b, 0x6f, 0x2d, 0x2d, 0x59, 0xee, 0xca, 0x81, 0x60,
	0x04, 0x2d, 0xfc, 0x0a, 0x79, 0x83, 0x1e, 0x7a, 0xe8, 0xbd, 0x0f, 0x91, 0x73, 0xef, 0x3d, 0xb5,
	0xef, 0x51, 0xec, 0x8f, 0x24, 0x52, 0xa1, 0x6c, 0xcb, 0xc9, 0x29, 0x17, 0x63, 0xe7, 0x67, 0x67,
	0x66, 0xe7, 0xfb, 0x66, 0x64, 0x42, 0x03, 0x77, 0x49, 0x26, 0x5a, 0x69, 0x96, 0x88, 0x04, 0x35,
	0xb5, 0x70, 0xb2, 0xd9, 0x21, 0x02, 0x6f, 0x7a, 0x37, 0x0f, 0x93, 0xe4, 0xb0, 0x4b, 0x02, 0x9c,
	0xd2, 0x00, 0x33, 0x96, 0x08, 0x2c, 0x68, 0xc2, 0xb8, 0x76, 0xf6, 0x6e, 0xe1, 0x34, 0xe5, 0x51,
	0x12, 0x4f, 0xb3, 0xdf, 0x90, 0xea, 0x58, 0x0c, 0x52, 0xc2, 0x03, 0xf5, 0x57, 0xeb, 0x7d, 0x0a,
	0x8d, 0x47, 0x11, 0x65, 0x87, 0xf8, 0x09, 0xce, 0x70, 0x0f, 0x6d, 0x00, 0x8a, 0x8e, 0x48, 0x74,
	0xdc, 0xa6, 0x4c, 0x90, 0xec, 0x04, 0x77, 0xdb, 0x9c, 0x44, 0xae, 0xb5, 0x62, 0xad, 0xcd, 0x85,
	0x4b, 0xca, 0xf2, 0xc8, 0x18, 0xf6, 0x49, 0x24, 0xbd, 0x55, 0x8d, 0x45, 0x6f, 0x5b, 0x7b, 0x2b,
	0x4b, 0xce, 0xdb, 0xef, 0x0d, 0x53, 0xed, 0x0b, 0x2c, 0x08, 0x5a, 0x04, 0x7b, 0xef, 0xb1, 0x0a,
	0x5d, 0x0d, 0xed, 0xbd, 0xc7, 0xc8, 0x85, 0xf9, 0x67, 0x38, 0x63, 0x94, 0x1d, 0xaa, 0x08, 0xd5,
	0x70, 0x28, 0x22, 0x0f, 0x6a, 0x3b, 0x19, 0x15, 0x34, 0xc2, 0x5d, 0x77, 0x4e, 0x99, 0x46, 0xb2,
	0xbc, 0x75, 0xc0, 0x8e, 0x59, 0xf2, 0x9c, 0xb9, 0x15, 0x7d, 0xcb, 0x88, 0xfe, 0x0f, 0xd0, 0xdc,
	0x4d, 0x04, 0xfd, 0x89, 0x92, 0x4c, 0xbf, 0x6d, 0x19, 0xaa, 0x5c, 0x66, 0x56, 0x39, 0xeb, 0xa1,
	0x16, 0xd0, 0xfb, 0x50, 0xeb, 0x73, 0x92, 0xb5, 0xfb, 0x34, 0x56, 0x79, 0xeb, 0xe1, 0xbc, 0x94,
	0x0f, 0x68, 0x8c, 0x6e, 0x80, 0xd3, 0x23, 0xe2, 0x28, 0x89, 0x55, 0xd6, 0x7a, 0x68, 0x24, 0xff,
	0x55, 0x05, 0xaa, 0x0f, 0xe4, 0xeb, 0x10, 0x82, 0x4a, 0x7a, 0x44, 0x63, 0x13, 0x51, 0x9d, 0xd1,
	0x6d, 0x40, 0xc7, 0xfd, 0x0e, 0xc9, 0x18, 0x11, 0x84, 0xb7, 0xa3, 0x6e, 0x9f, 0x0b, 0x92, 0x99,
	0xd0, 0xd7, 0xc6, 0x96, 0x1d, 0x6d, 0x40, 0x9b, 0xb0, 0x9c, 0x73, 0x67, 0xb8, 0x47, 0x78, 0x8a,
	0x23, 0x62, 0x52, 0xbe, 0x37, 0xb6, 0xed, 0x0e, 0x4d, 0x68, 0x1b, 0xae, 0xe7, 0xae, 0x24, 0x9d,
	0x9f, 0x49, 0x24, 0xbe, 0x1f, 0xa4, 0x44, 0x75, 0xa0, 0x1e, 0xe6, 0xe2, 0xed, 0x8d, 0x6c, 0xa5,
	0x97, 0x64, 0x48, 0xb7, 0x5a, 0x7e, 0x49, 0xda, 0xd0, 0xc7, 0xb0, 0x48, 0x15, 0x64, 0x6d, 0x4e,
	0xb2, 0x13, 0x1a, 0x11, 0x77, 0x5e, 0x79, 0x37, 0xb5, 0x76, 0x5f, 0x2b, 0xd1, 0xd7, 0xb0, 0x60,
	0xdc, 0x52, 0xd9, 0x69, 0xb7, 0xb6, 0x62, 0xad, 0x35, 0xb6, 0xbc, 0x56, 0x81, 0xc0, 0xad, 0x1c,
	0xcf, 0xc2, 0x06, 0xcd, 0x91, 0x6e, 0x15, 0x9a, 0x9a, 0x74, 0x51, 0xd2, 0xeb, 0x61, 0x16, 0xbb,
	0x75, 0x95, 0x64, 0x41, 0x29, 0x77, 0xb4, 0x0e, 0xed, 0xc0, 0x22, 0x33, 0x70, 0x9a, 0x2c, 0x8d,
	0x95, 0xb9, 0xb5, 0xc6, 0xd6, 0xcd, 0x89, 0x2c, 0x05, 0xcc, 0xc3, 0x26, 0x2b, 0x50, 0x60, 0x15,
	0x4c, 0xe5, 0xcf, 0x49, 0xa7, 0xdd, 0xcf, 0xba, 0xee, 0x82, 0xce, 0x34, 0x52, 0x1e, 0x64, 0x5d,
	0xb4, 0x05, 0x95, 0x13, 0x9c, 0x71, 0xb7, 0xa9, 0xe2, 0xdf, 0x9a, 0x88, 0xaf, 0x80, 0x6f, 0x3d,
	0xc5, 0x19, 0xff, 0x96, 0x89, 0x6c, 0x10, 0x2a, 0x5f, 0xef, 0x4b, 0xa8, 0x8f, 0x54, 0x68, 0x09,
	0xe6, 0x8e, 0xc9, 0xc0, 0x90, 0x42, 0x1e, 0x25, 0xf5, 0x4e, 0x70, 0xb7, 0x4f, 0x0c, 0x0d, 0xb4,
	0x70, 0xcf, 0xbe, 0x6b, 0xf9, 0xff, 0x58, 0xb0, 0xa4, 0x42, 0x7e, 0x47, 0xb9, 0x08, 0xc9, 0x2f,
	0x7d, 0xc2, 0xc5, 0x14, 0x0a, 0x59, 0xb3, 0x52, 0xc8, 0xbe, 0x02, 0x85, 0xe6, 0xae, 0x42, 0xa1,
	0xca, 0x74, 0x0a, 0xf9, 0x14, 0xae, 0xe5, 0xde, 0xc7, 0xd3, 0x84, 0x71, 0x82, 0x3e, 0x01, 0x47,
	0x4e, 0x5f, 0x9f, 0xab, 0x47, 0x35, 0xb6, 0x16, 0x5b, 0x7a, 0x35, 0xb5, 0xf6, 0x95, 0x36, 0x34,
	0x56, 0xb4, 0x01, 0x8e, 0xea, 0x3e, 0x77, 0x6d, 0x05, 0xc6, 0x72, 0x19, 0x18, 0xa1, 0xf1, 0xf1,
	0xff, 0xac, 0x00, 0x52, 0x9a, 0x9d, 0x8c, 0x60, 0x41, 0xde, 0xc1, 0x6e, 0x96, 0x0c, 0xa4, 0x73,
	0x99, 0x81, 0x9c, 0x7f, 0xc3, 0x81, 0xac, 0x5d, 0x6a, 0x20, 0x61, 0xf6, 0x81, 0xbc, 0x6f, 0x66,
	0x4d, 0xcf, 0xf2, 0xe7, 0x65, 0xf0, 0x16, 0xc0, 0x7c, 0x7b, 0x83, 0xb7, 0x0e, 0xcb, 0x2a, 0xfc,
	0x43, 0xc2, 0xa3, 0x8c, 0x76, 0x46, 0x6c, 0x29, 0x59, 0xe9, 0x7e, 0x0f, 0xae, 0x4f, 0xf8, 0xbe,
	0x01, 0x8f, 0xad, 0x0b, 0x79, 0xfc, 0x87, 0x6d, 0x78, 0x7c, 0x90, 0xc6, 0x39, 0x1e, 0x97, 0xfd,
	0xd8, 0x4c, 0x02, 0x6d, 0xcf, 0x06, 0xf4, 0xeb, 0x18, 0x56, 0xae, 0x8e, 0x61, 0x75, 0x3a, 0x86,
	0x85, 0x87, 0xbc, 0x3d, 0x0c, 0xff, 0xb3, 0xcd, 0xf2, 0xdc, 0x1f, 0xb0, 0xe8, 0x5d, 0x1c, 0xf7,
	0x5d, 0x58, 0x48, 0x93, 0xb8, 0x8d, 0x59, 0x44, 0xb8, 0x48, 0xb2, 0xf3, 0x5a, 0x9c, 0x6b, 0x41,
	0xeb, 0x49, 0x12, 0x3f, 0x30, 0x57, 0xc2, 0x46, 0x3a, 0x16, 0xbc, 0x3b, 0xd0, 0xc8, 0xd9, 0x24,
	0xa3, 0x24, 0x5d, 0x87, 0x8c, 0x92, 0x67, 0xa9, 0x93, 0x4d, 0x30, 0xef, 0x57, 0x67, 0x7f, 0xcd,
	0xf0, 0xf1, 0x21, 0xe9, 0x92, 0x73, 0xf9, 0xb8, 0xf5, 0xca, 0x01, 0x47, 0xb9, 0x72, 0xf4, 0x9b,
	0x05, 0x15, 0xb9, 0xf4, 0xd1, 0x07, 0x65, 0xe5, 0xe6, 0x7e, 0xee, 0xbc, 0x95, 0xe9, 0x0e, 0x7a,
	0xce, 0xfc, 0x3b, 0x67, 0x7f, 0xb9, 0x76, 0xcd, 0x3a, 0xfb, 0xfb, 0xdf, 0x97, 0xf6, 0x67, 0xe8,
	0xd3, 0xa0, 0xf0, 0xcf, 0xee, 0xb8, 0x73, 0x81, 0x89, 0x10, 0xe8, 0x41, 0x42, 0xa7, 0xe0, 0xe8,
	0xed, 0x81, 0x3e, 0xbc, 0x70, 0xb3, 0x78, 0xcb, 0xc3, 0xd9, 0x7d, 0x9a, 0xd0, 0x78, 0x94, 0xf9,
	0x6e, 0x2e, 0xf3, 0x86, 0x7f, 0xd9, 0xcc, 0xf7, 0xac, 0x75, 0xf4, 0xd2, 0x82, 0xda, 0x70, 0x61,
	0xa0, 0xd5, 0xb2, 0xfc, 0x13, 0xab, 0xc7, 0xfb, 0xe8, 0x7c, 0x27, 0x53, 0xd1, 0x57, 0xb9, 0x8a,
	0xbe, 0x40, 0xad, 0x4b, 0x56, 0x14, 0x9c, 0x4a, 0x7c, 0x5e, 0xa0, 0x33, 0x0b, 0x1c, 0x3d, 0x8d,
	0xe5, 0x3d, 0x29, 0x4c, 0xea, 0x94, 0x9e, 0xdc, 0xcf, 0x55, 0xb0, 0xed, 0xcd, 0x58, 0x81, 0x6c,
	0xcd, 0xef, 0x16, 0x54, 0x24, 0x5f, 0xcb, 0xa9, 0x91, 0x63, 0xf2, 0x94, 0x02, 0xa2, 0x5c, 0x01,
	0xcf, 0xfc, 0xf0, 0xc2, 0x02, 0xcc, 0x12, 0xe0, 0xc1, 0xe9, 0xeb, 0x9b, 0xe1, 0x45, 0x80, 0x23,
	0xf5, 0xa1, 0x14, 0xf0, 0x01, 0x8b, 0x6e, 0x8f, 0xf1, 0xfb, 0x15, 0x1c, 0xcd, 0xf7, 0xf2, 0x46,
	0x15, 0x66, 0x61, 0x4a, 0x9d, 0x05, 0xa8, 0xd6, 0x67, 0x6c, 0xd4, 0x37, 0xf5, 0x1f, 0xe7, 0x8d,
	0xa1, 0xe3, 0xa8, 0x8f, 0xb5, 0xed, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xea, 0xb6, 0x17, 0xb3,
	0x20, 0x0e, 0x00, 0x00,
}
