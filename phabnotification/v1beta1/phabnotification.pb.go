// Code generated by protoc-gen-go. DO NOT EDIT.
// source: phabnotification.proto

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	phabnotification.proto

It has these top-level messages:
	GetResponse
	PhabNotification
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetResponse struct {
	Result []*PhabNotification `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetResponse) GetResult() []*PhabNotification {
	if m != nil {
		return m.Result
	}
	return nil
}

type PhabNotification struct {
	StoryPhid        string `protobuf:"bytes,1,opt,name=story_phid,json=storyPhid" json:"story_phid,omitempty"`
	ChronologicalKey int64  `protobuf:"varint,2,opt,name=chronological_key,json=chronologicalKey" json:"chronological_key,omitempty"`
	StoryType        string `protobuf:"bytes,3,opt,name=story_type,json=storyType" json:"story_type,omitempty"`
	StoryData        string `protobuf:"bytes,4,opt,name=story_data,json=storyData" json:"story_data,omitempty"`
	AuthorPhid       string `protobuf:"bytes,5,opt,name=author_phid,json=authorPhid" json:"author_phid,omitempty"`
	HasViewed        int32  `protobuf:"varint,6,opt,name=has_viewed,json=hasViewed" json:"has_viewed,omitempty"`
}

func (m *PhabNotification) Reset()                    { *m = PhabNotification{} }
func (m *PhabNotification) String() string            { return proto.CompactTextString(m) }
func (*PhabNotification) ProtoMessage()               {}
func (*PhabNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PhabNotification) GetStoryPhid() string {
	if m != nil {
		return m.StoryPhid
	}
	return ""
}

func (m *PhabNotification) GetChronologicalKey() int64 {
	if m != nil {
		return m.ChronologicalKey
	}
	return 0
}

func (m *PhabNotification) GetStoryType() string {
	if m != nil {
		return m.StoryType
	}
	return ""
}

func (m *PhabNotification) GetStoryData() string {
	if m != nil {
		return m.StoryData
	}
	return ""
}

func (m *PhabNotification) GetAuthorPhid() string {
	if m != nil {
		return m.AuthorPhid
	}
	return ""
}

func (m *PhabNotification) GetHasViewed() int32 {
	if m != nil {
		return m.HasViewed
	}
	return 0
}

func init() {
	proto.RegisterType((*GetResponse)(nil), "appscode.notification.v1beta1.GetResponse")
	proto.RegisterType((*PhabNotification)(nil), "appscode.notification.v1beta1.PhabNotification")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PhabNotifications service

type PhabNotificationsClient interface {
	// Gets Notifications,
	Get(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Update(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type phabNotificationsClient struct {
	cc *grpc.ClientConn
}

func NewPhabNotificationsClient(cc *grpc.ClientConn) PhabNotificationsClient {
	return &phabNotificationsClient{cc}
}

func (c *phabNotificationsClient) Get(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/appscode.notification.v1beta1.PhabNotifications/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phabNotificationsClient) Update(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.notification.v1beta1.PhabNotifications/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PhabNotifications service

type PhabNotificationsServer interface {
	// Gets Notifications,
	Get(context.Context, *appscode_dtypes.VoidRequest) (*GetResponse, error)
	Update(context.Context, *appscode_dtypes.VoidRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterPhabNotificationsServer(s *grpc.Server, srv PhabNotificationsServer) {
	s.RegisterService(&_PhabNotifications_serviceDesc, srv)
}

func _PhabNotifications_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhabNotificationsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.notification.v1beta1.PhabNotifications/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhabNotificationsServer).Get(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhabNotifications_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhabNotificationsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.notification.v1beta1.PhabNotifications/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhabNotificationsServer).Update(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PhabNotifications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.notification.v1beta1.PhabNotifications",
	HandlerType: (*PhabNotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PhabNotifications_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PhabNotifications_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "phabnotification.proto",
}

func init() { proto.RegisterFile("phabnotification.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbf, 0x8e, 0xd3, 0x40,
	0x10, 0xc6, 0xb5, 0x09, 0x17, 0xc8, 0xa6, 0xb9, 0x5b, 0x09, 0x64, 0x45, 0x17, 0x2e, 0x4a, 0x15,
	0x0e, 0x64, 0x2b, 0x47, 0x77, 0x12, 0xcd, 0x09, 0x29, 0x05, 0x02, 0x45, 0x16, 0xa4, 0xa0, 0x89,
	0xc6, 0xde, 0x21, 0xbb, 0x60, 0x3c, 0x8b, 0x77, 0x73, 0xc8, 0xed, 0x89, 0x8e, 0x92, 0x97, 0xe0,
	0x05, 0x78, 0x12, 0x5a, 0x4a, 0x1e, 0x04, 0x9d, 0x6d, 0xe4, 0x3f, 0xe8, 0x88, 0xae, 0x71, 0x31,
	0xbf, 0xfd, 0xc6, 0xdf, 0x7e, 0x33, 0xcb, 0x1f, 0x18, 0x05, 0x51, 0x4a, 0x4e, 0xbf, 0xd3, 0x31,
	0x38, 0x4d, 0xa9, 0x6f, 0x32, 0x72, 0x24, 0x26, 0x60, 0x8c, 0x8d, 0x49, 0xa2, 0xdf, 0x82, 0x97,
	0x8b, 0x08, 0x1d, 0x2c, 0xc6, 0xc7, 0x5b, 0xa2, 0x6d, 0x82, 0x01, 0x18, 0x1d, 0x40, 0x9a, 0x92,
	0x2b, 0xb0, 0x2d, 0xc5, 0xe3, 0x87, 0x7f, 0xc5, 0x37, 0xf0, 0x93, 0x16, 0x97, 0x2e, 0x37, 0x68,
	0x83, 0xe2, 0x5b, 0x1e, 0x98, 0xad, 0xf9, 0x68, 0x89, 0x2e, 0x44, 0x6b, 0x28, 0xb5, 0x28, 0x96,
	0x7c, 0x90, 0xa1, 0xdd, 0x25, 0xce, 0x63, 0xd3, 0xfe, 0x7c, 0x74, 0x16, 0xf8, 0xff, 0x75, 0xe7,
	0xaf, 0x14, 0x44, 0xaf, 0x1a, 0x20, 0xac, 0xe4, 0xb3, 0x5f, 0x8c, 0x1f, 0x76, 0xa1, 0x98, 0x70,
	0x6e, 0x1d, 0x65, 0xf9, 0xc6, 0x28, 0x2d, 0x3d, 0x36, 0x65, 0xf3, 0x61, 0x38, 0x2c, 0x2a, 0x2b,
	0xa5, 0xa5, 0x78, 0xcc, 0x8f, 0x62, 0x95, 0x51, 0x4a, 0x09, 0x6d, 0x75, 0x0c, 0xc9, 0xe6, 0x03,
	0xe6, 0x5e, 0x6f, 0xca, 0xe6, 0xfd, 0xf0, 0xb0, 0x05, 0x5e, 0x60, 0x5e, 0xf7, 0xba, 0xbe, 0x8d,
	0xd7, 0x6f, 0xf4, 0x7a, 0x9d, 0x1b, 0xac, 0xb1, 0x04, 0x07, 0xde, 0x9d, 0x06, 0x7e, 0x0e, 0x0e,
	0xc4, 0x09, 0x1f, 0xc1, 0xce, 0x29, 0xca, 0x4a, 0x2b, 0x07, 0x05, 0xe7, 0x65, 0xa9, 0xf0, 0x32,
	0xe1, 0x5c, 0x81, 0xdd, 0x5c, 0x6a, 0xfc, 0x8c, 0xd2, 0x1b, 0x4c, 0xd9, 0xfc, 0x20, 0x1c, 0x2a,
	0xb0, 0xeb, 0xa2, 0x70, 0xf6, 0xbd, 0xc7, 0x8f, 0xba, 0xd7, 0xb3, 0xe2, 0x2b, 0xe3, 0xfd, 0x25,
	0x3a, 0x71, 0x5c, 0xa7, 0x56, 0x46, 0xee, 0xaf, 0x49, 0xcb, 0x10, 0x3f, 0xed, 0xd0, 0xba, 0xf1,
	0xe9, 0x9e, 0x4c, 0x1b, 0xf3, 0x98, 0x9d, 0x5f, 0xfd, 0xf0, 0x7a, 0xf7, 0xd8, 0xd5, 0xcf, 0xdf,
	0xdf, 0x7a, 0xbe, 0x78, 0x12, 0x6c, 0x5a, 0xf3, 0xec, 0x6e, 0x54, 0x50, 0xb5, 0x08, 0xde, 0x5b,
	0x4a, 0xc5, 0x17, 0xc6, 0x07, 0x6f, 0x8c, 0x04, 0x87, 0x7b, 0x0c, 0x4d, 0x6e, 0xa0, 0x95, 0x87,
	0x67, 0x0d, 0x0f, 0x8b, 0xd9, 0xad, 0x3c, 0x9c, 0xb3, 0xd3, 0x8b, 0x97, 0xfc, 0x51, 0x4c, 0x1f,
	0xeb, 0x5f, 0x80, 0xd1, 0xfe, 0x3f, 0x4f, 0xa1, 0x12, 0x5d, 0xdc, 0xef, 0x86, 0xba, 0xba, 0xde,
	0xd2, 0x15, 0x7b, 0x7b, 0xb7, 0x3a, 0x11, 0x0d, 0x8a, 0xbd, 0x7d, 0xfa, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x48, 0x35, 0x8e, 0x27, 0x4f, 0x03, 0x00, 0x00,
}
