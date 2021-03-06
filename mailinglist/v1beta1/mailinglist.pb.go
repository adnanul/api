// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mailinglist.proto

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	mailinglist.proto

It has these top-level messages:
	SendEmailRequest
	SubscribeRequest
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

type SendEmailRequest struct {
	SenderName    string `protobuf:"bytes,1,opt,name=sender_name,json=senderName" json:"sender_name,omitempty"`
	SenderEmail   string `protobuf:"bytes,2,opt,name=sender_email,json=senderEmail" json:"sender_email,omitempty"`
	Subject       string `protobuf:"bytes,3,opt,name=subject" json:"subject,omitempty"`
	Body          string `protobuf:"bytes,4,opt,name=body" json:"body,omitempty"`
	ReceiverEmail string `protobuf:"bytes,5,opt,name=receiver_email,json=receiverEmail" json:"receiver_email,omitempty"`
}

func (m *SendEmailRequest) Reset()                    { *m = SendEmailRequest{} }
func (m *SendEmailRequest) String() string            { return proto.CompactTextString(m) }
func (*SendEmailRequest) ProtoMessage()               {}
func (*SendEmailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SendEmailRequest) GetSenderName() string {
	if m != nil {
		return m.SenderName
	}
	return ""
}

func (m *SendEmailRequest) GetSenderEmail() string {
	if m != nil {
		return m.SenderEmail
	}
	return ""
}

func (m *SendEmailRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SendEmailRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SendEmailRequest) GetReceiverEmail() string {
	if m != nil {
		return m.ReceiverEmail
	}
	return ""
}

type SubscribeRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SubscribeRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*SendEmailRequest)(nil), "appscode.mailinglist.v1beta1.SendEmailRequest")
	proto.RegisterType((*SubscribeRequest)(nil), "appscode.mailinglist.v1beta1.SubscribeRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MailingList service

type MailingListClient interface {
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type mailingListClient struct {
	cc *grpc.ClientConn
}

func NewMailingListClient(cc *grpc.ClientConn) MailingListClient {
	return &mailingListClient{cc}
}

func (c *mailingListClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.mailinglist.v1beta1.MailingList/SendEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailingListClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.mailinglist.v1beta1.MailingList/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MailingList service

type MailingListServer interface {
	SendEmail(context.Context, *SendEmailRequest) (*appscode_dtypes.VoidResponse, error)
	Subscribe(context.Context, *SubscribeRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterMailingListServer(s *grpc.Server, srv MailingListServer) {
	s.RegisterService(&_MailingList_serviceDesc, srv)
}

func _MailingList_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailingListServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.mailinglist.v1beta1.MailingList/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailingList_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailingListServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.mailinglist.v1beta1.MailingList/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MailingList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.mailinglist.v1beta1.MailingList",
	HandlerType: (*MailingListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _MailingList_SendEmail_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _MailingList_Subscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mailinglist.proto",
}

func init() { proto.RegisterFile("mailinglist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x49, 0x6e, 0x7b, 0x7b, 0x3b, 0xbd, 0xf7, 0xd2, 0x3b, 0xdc, 0x45, 0x08, 0xd5, 0xd6,
	0x60, 0xa1, 0xb8, 0x48, 0x88, 0x82, 0x0b, 0x97, 0x15, 0x77, 0x2a, 0xa5, 0x05, 0x17, 0x6e, 0xca,
	0x24, 0x39, 0x94, 0x29, 0xcd, 0x4c, 0xcc, 0x4c, 0x0b, 0xdd, 0x76, 0xef, 0xca, 0xbd, 0x2f, 0xe0,
	0xda, 0xad, 0x2f, 0xe1, 0x2b, 0xf8, 0x20, 0x92, 0x99, 0xa4, 0xff, 0xa8, 0x76, 0x13, 0x32, 0xf3,
	0x9d, 0xf3, 0x9d, 0xdf, 0xf9, 0x18, 0xf4, 0x2f, 0x26, 0x74, 0x42, 0xd9, 0x68, 0x42, 0x85, 0x74,
	0x93, 0x94, 0x4b, 0x8e, 0x1b, 0x24, 0x49, 0x44, 0xc8, 0x23, 0x70, 0xd7, 0xb5, 0x99, 0x1f, 0x80,
	0x24, 0xbe, 0xdd, 0x18, 0x71, 0x3e, 0x9a, 0x80, 0x47, 0x12, 0xea, 0x11, 0xc6, 0xb8, 0x24, 0x92,
	0x72, 0x26, 0x74, 0xaf, 0x7d, 0x58, 0xf4, 0x7e, 0xa1, 0x37, 0x37, 0xf4, 0x48, 0xce, 0x13, 0x10,
	0x9e, 0xfa, 0xea, 0x02, 0xe7, 0xc5, 0x40, 0xf5, 0x01, 0xb0, 0xe8, 0x2a, 0x9b, 0xdd, 0x87, 0x87,
	0x29, 0x08, 0x89, 0x9b, 0xa8, 0x26, 0x80, 0x45, 0x90, 0x0e, 0x19, 0x89, 0xc1, 0x32, 0x5a, 0x46,
	0xa7, 0xda, 0x47, 0xfa, 0xea, 0x96, 0xc4, 0x80, 0x8f, 0xd0, 0xef, 0xbc, 0x00, 0xb2, 0x3e, 0xcb,
	0x54, 0x15, 0x79, 0x93, 0xb2, 0xc2, 0x16, 0xaa, 0x88, 0x69, 0x30, 0x86, 0x50, 0x5a, 0x3f, 0x94,
	0x5a, 0x1c, 0x31, 0x46, 0xa5, 0x80, 0x47, 0x73, 0xab, 0xa4, 0xae, 0xd5, 0x3f, 0x6e, 0xa3, 0xbf,
	0x29, 0x84, 0x40, 0x67, 0x4b, 0xcb, 0xb2, 0x52, 0xff, 0x14, 0xb7, 0xca, 0xd4, 0xe9, 0xa0, 0xfa,
	0x60, 0x1a, 0x88, 0x30, 0xa5, 0x01, 0x14, 0xb0, 0xff, 0x51, 0x59, 0x77, 0x68, 0x4c, 0x7d, 0x38,
	0x7d, 0x33, 0x51, 0xed, 0x46, 0xc7, 0x79, 0x4d, 0x85, 0xc4, 0x8f, 0x06, 0xaa, 0x2e, 0xf7, 0xc4,
	0xae, 0xfb, 0x5d, 0xe6, 0xee, 0x76, 0x20, 0xf6, 0xc1, 0xaa, 0x5e, 0x67, 0xe8, 0xde, 0x71, 0x1a,
	0xf5, 0x41, 0x24, 0x9c, 0x09, 0x70, 0xfc, 0xc5, 0xab, 0x65, 0xfe, 0x32, 0x16, 0xef, 0x1f, 0x4f,
	0x66, 0xdb, 0x69, 0x79, 0xc3, 0x8d, 0xd4, 0x33, 0x23, 0x2f, 0xb7, 0xf6, 0xc6, 0x82, 0xb3, 0x0b,
	0xe3, 0x04, 0x3f, 0x67, 0x3c, 0xc5, 0x2a, 0x7b, 0x79, 0xb6, 0x76, 0xde, 0xc7, 0xd3, 0x5d, 0xe3,
	0x39, 0xb7, 0xfd, 0x1d, 0x3c, 0xf9, 0x84, 0x25, 0x96, 0x28, 0x26, 0x14, 0x80, 0xdd, 0x4b, 0x74,
	0x1c, 0xf2, 0x78, 0x35, 0x87, 0x24, 0x74, 0x17, 0x5b, 0xb7, 0xbe, 0x96, 0x72, 0x2f, 0x7b, 0x52,
	0x3d, 0xe3, 0xbe, 0x92, 0x8b, 0xc1, 0x4f, 0xf5, 0xc8, 0xce, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x58, 0xaa, 0xa2, 0x71, 0xf6, 0x02, 0x00, 0x00,
}
