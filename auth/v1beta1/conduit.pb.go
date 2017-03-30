// Code generated by protoc-gen-go.
// source: conduit.proto
// DO NOT EDIT!

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

type ConduitWhoAmIResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	User   *ConduitUser            `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *ConduitWhoAmIResponse) Reset()                    { *m = ConduitWhoAmIResponse{} }
func (m *ConduitWhoAmIResponse) String() string            { return proto.CompactTextString(m) }
func (*ConduitWhoAmIResponse) ProtoMessage()               {}
func (*ConduitWhoAmIResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ConduitWhoAmIResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConduitWhoAmIResponse) GetUser() *ConduitUser {
	if m != nil {
		return m.User
	}
	return nil
}

type ConduitUsersResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Users  []*ConduitUser          `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
}

func (m *ConduitUsersResponse) Reset()                    { *m = ConduitUsersResponse{} }
func (m *ConduitUsersResponse) String() string            { return proto.CompactTextString(m) }
func (*ConduitUsersResponse) ProtoMessage()               {}
func (*ConduitUsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ConduitUsersResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConduitUsersResponse) GetUsers() []*ConduitUser {
	if m != nil {
		return m.Users
	}
	return nil
}

type ConduitUser struct {
	Phid         string       `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	UserName     string       `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	RealName     string       `protobuf:"bytes,3,opt,name=real_name,json=realName" json:"real_name,omitempty"`
	Image        string       `protobuf:"bytes,4,opt,name=image" json:"image,omitempty"`
	Uri          string       `protobuf:"bytes,5,opt,name=uri" json:"uri,omitempty"`
	Roles        []string     `protobuf:"bytes,6,rep,name=roles" json:"roles,omitempty"`
	PrimaryEmail string       `protobuf:"bytes,7,opt,name=primary_email,json=primaryEmail" json:"primary_email,omitempty"`
	Preferences  *Preferences `protobuf:"bytes,8,opt,name=preferences" json:"preferences,omitempty"`
}

func (m *ConduitUser) Reset()                    { *m = ConduitUser{} }
func (m *ConduitUser) String() string            { return proto.CompactTextString(m) }
func (*ConduitUser) ProtoMessage()               {}
func (*ConduitUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ConduitUser) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *ConduitUser) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ConduitUser) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *ConduitUser) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ConduitUser) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *ConduitUser) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *ConduitUser) GetPrimaryEmail() string {
	if m != nil {
		return m.PrimaryEmail
	}
	return ""
}

func (m *ConduitUser) GetPreferences() *Preferences {
	if m != nil {
		return m.Preferences
	}
	return nil
}

type Preferences struct {
	TimeZone    string `protobuf:"bytes,1,opt,name=time_zone,json=timeZone" json:"time_zone,omitempty"`
	TimeFormat  string `protobuf:"bytes,2,opt,name=time_format,json=timeFormat" json:"time_format,omitempty"`
	DateFormate string `protobuf:"bytes,3,opt,name=date_formate,json=dateFormate" json:"date_formate,omitempty"`
}

func (m *Preferences) Reset()                    { *m = Preferences{} }
func (m *Preferences) String() string            { return proto.CompactTextString(m) }
func (*Preferences) ProtoMessage()               {}
func (*Preferences) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Preferences) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

func (m *Preferences) GetTimeFormat() string {
	if m != nil {
		return m.TimeFormat
	}
	return ""
}

func (m *Preferences) GetDateFormate() string {
	if m != nil {
		return m.DateFormate
	}
	return ""
}

func init() {
	proto.RegisterType((*ConduitWhoAmIResponse)(nil), "appscode.auth.v1beta1.ConduitWhoAmIResponse")
	proto.RegisterType((*ConduitUsersResponse)(nil), "appscode.auth.v1beta1.ConduitUsersResponse")
	proto.RegisterType((*ConduitUser)(nil), "appscode.auth.v1beta1.ConduitUser")
	proto.RegisterType((*Preferences)(nil), "appscode.auth.v1beta1.Preferences")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Conduit service

type ConduitClient interface {
	// This rpc is used to check a valid user from other applications.
	WhoAmI(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ConduitWhoAmIResponse, error)
	// appctl used this to validates the user token with phabricator.
	Users(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ConduitUsersResponse, error)
}

type conduitClient struct {
	cc *grpc.ClientConn
}

func NewConduitClient(cc *grpc.ClientConn) ConduitClient {
	return &conduitClient{cc}
}

func (c *conduitClient) WhoAmI(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ConduitWhoAmIResponse, error) {
	out := new(ConduitWhoAmIResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Conduit/WhoAmI", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitClient) Users(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ConduitUsersResponse, error) {
	out := new(ConduitUsersResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Conduit/Users", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Conduit service

type ConduitServer interface {
	// This rpc is used to check a valid user from other applications.
	WhoAmI(context.Context, *appscode_dtypes.VoidRequest) (*ConduitWhoAmIResponse, error)
	// appctl used this to validates the user token with phabricator.
	Users(context.Context, *appscode_dtypes.VoidRequest) (*ConduitUsersResponse, error)
}

func RegisterConduitServer(s *grpc.Server, srv ConduitServer) {
	s.RegisterService(&_Conduit_serviceDesc, srv)
}

func _Conduit_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Conduit/WhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitServer).WhoAmI(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conduit_Users_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitServer).Users(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Conduit/Users",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitServer).Users(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Conduit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.auth.v1beta1.Conduit",
	HandlerType: (*ConduitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhoAmI",
			Handler:    _Conduit_WhoAmI_Handler,
		},
		{
			MethodName: "Users",
			Handler:    _Conduit_Users_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conduit.proto",
}

func init() { proto.RegisterFile("conduit.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdf, 0x6a, 0xd4, 0x4e,
	0x14, 0x26, 0xd9, 0xff, 0x27, 0x5b, 0xf8, 0x31, 0xb4, 0xfc, 0xc2, 0x5a, 0xed, 0x1a, 0x6f, 0x04,
	0x25, 0x71, 0x2b, 0x48, 0x2f, 0xf4, 0xc2, 0xfa, 0x07, 0xbc, 0x91, 0x25, 0xa2, 0x42, 0x6f, 0x96,
	0xe9, 0xe6, 0x74, 0x33, 0xb2, 0x99, 0x89, 0x33, 0x13, 0xa5, 0x5e, 0x69, 0xf1, 0x01, 0x84, 0xbe,
	0x89, 0xe0, 0x93, 0xf8, 0x0a, 0x3e, 0x88, 0xcc, 0x64, 0xb6, 0x66, 0x51, 0xdb, 0xe2, 0x4d, 0xc8,
	0x7c, 0xdf, 0x77, 0xe6, 0x7c, 0x73, 0xfe, 0xc0, 0xc6, 0x5c, 0xf0, 0xac, 0x62, 0x3a, 0x2e, 0xa5,
	0xd0, 0x82, 0x6c, 0xd1, 0xb2, 0x54, 0x73, 0x91, 0x61, 0x4c, 0x2b, 0x9d, 0xc7, 0xef, 0x26, 0x87,
	0xa8, 0xe9, 0x64, 0xb4, 0xbd, 0x10, 0x62, 0xb1, 0xc4, 0x84, 0x96, 0x2c, 0xa1, 0x9c, 0x0b, 0x4d,
	0x35, 0x13, 0x5c, 0xd5, 0x41, 0xa3, 0x6b, 0xab, 0xa0, 0xbf, 0xf0, 0x3b, 0x6b, 0x7c, 0xa6, 0x8f,
	0x4b, 0x54, 0x89, 0xfd, 0xd6, 0x82, 0xe8, 0xa3, 0x07, 0x5b, 0x8f, 0x6a, 0x1f, 0xaf, 0x73, 0xf1,
	0xb0, 0x78, 0x96, 0xa2, 0x2a, 0x05, 0x57, 0x48, 0x12, 0xe8, 0x2a, 0x4d, 0x75, 0xa5, 0x42, 0x6f,
	0xec, 0xdd, 0x0c, 0x76, 0xff, 0x8f, 0xcf, 0x0c, 0xd6, 0xf7, 0xc4, 0x2f, 0x2c, 0x9d, 0x3a, 0x19,
	0xb9, 0x07, 0xed, 0x4a, 0xa1, 0x0c, 0x7d, 0x2b, 0x8f, 0xe2, 0x3f, 0xbe, 0x27, 0x76, 0xc9, 0x5e,
	0x2a, 0x94, 0xa9, 0xd5, 0x47, 0x9f, 0x3c, 0xd8, 0x6c, 0xa0, 0xea, 0xdf, 0x1d, 0xec, 0x41, 0xc7,
	0xdc, 0xa8, 0x42, 0x7f, 0xdc, 0xba, 0xa4, 0x85, 0x3a, 0x20, 0xfa, 0xec, 0x43, 0xd0, 0x80, 0x09,
	0x81, 0x76, 0x99, 0xb3, 0xcc, 0x26, 0x1e, 0xa4, 0xf6, 0x9f, 0x5c, 0x81, 0x81, 0x11, 0xcf, 0x38,
	0x2d, 0xd0, 0x3e, 0x72, 0x90, 0xf6, 0x0d, 0xf0, 0x9c, 0x16, 0x68, 0x48, 0x89, 0x74, 0x59, 0x93,
	0xad, 0x9a, 0x34, 0x80, 0x25, 0x37, 0xa1, 0xc3, 0x0a, 0xba, 0xc0, 0xb0, 0x6d, 0x89, 0xfa, 0x40,
	0xfe, 0x83, 0x56, 0x25, 0x59, 0xd8, 0xb1, 0x98, 0xf9, 0x35, 0x3a, 0x29, 0x96, 0xa8, 0xc2, 0xee,
	0xb8, 0x65, 0x74, 0xf6, 0x40, 0x6e, 0xc0, 0x46, 0x29, 0x59, 0x41, 0xe5, 0xf1, 0x0c, 0x0b, 0xca,
	0x96, 0x61, 0xcf, 0x46, 0x0c, 0x1d, 0xf8, 0xc4, 0x60, 0xe4, 0x31, 0x04, 0xa5, 0xc4, 0x23, 0x94,
	0xc8, 0xe7, 0xa8, 0xc2, 0xfe, 0xb9, 0x3d, 0x98, 0xfe, 0x52, 0xa6, 0xcd, 0xb0, 0x88, 0x43, 0xd0,
	0xe0, 0xcc, 0xa3, 0x34, 0x2b, 0x70, 0xf6, 0x41, 0x70, 0x74, 0xa5, 0xe8, 0x1b, 0xe0, 0x40, 0x70,
	0x24, 0x3b, 0x10, 0x58, 0xf2, 0x48, 0xc8, 0x82, 0x6a, 0x57, 0x10, 0x30, 0xd0, 0x53, 0x8b, 0x90,
	0xeb, 0x30, 0xcc, 0xa8, 0x5e, 0x09, 0x56, 0x55, 0x09, 0x0c, 0x56, 0x2b, 0x70, 0xf7, 0xab, 0x0f,
	0x3d, 0x57, 0x76, 0x72, 0xea, 0x41, 0xb7, 0x1e, 0x41, 0xb2, 0xfd, 0x5b, 0xa3, 0x5f, 0x09, 0x96,
	0xa5, 0xf8, 0xb6, 0x42, 0xa5, 0x47, 0xb7, 0xcf, 0x6f, 0xeb, 0xfa, 0x18, 0x47, 0x0f, 0x4e, 0xbe,
	0x85, 0x7e, 0xdf, 0x3b, 0xf9, 0xfe, 0xe3, 0xd4, 0x9f, 0x90, 0x24, 0x99, 0xad, 0x6f, 0x4c, 0xa5,
	0xf3, 0xc4, 0x5d, 0x90, 0xb8, 0x7d, 0x4c, 0xde, 0xe7, 0x82, 0x16, 0x2c, 0x79, 0xa3, 0x04, 0x27,
	0x5f, 0x3c, 0xe8, 0xd8, 0xa9, 0xbc, 0xc0, 0xd4, 0xad, 0x8b, 0x67, 0xed, 0x6c, 0xb0, 0xa3, 0xfb,
	0x0d, 0x4f, 0x77, 0x48, 0x7c, 0x19, 0x4f, 0x76, 0x4c, 0xad, 0xa5, 0xfd, 0x3d, 0xb8, 0x3a, 0x17,
	0x45, 0x23, 0x5f, 0xc9, 0xd6, 0x72, 0xee, 0x0f, 0x5d, 0xd2, 0xa9, 0xd9, 0xf0, 0xa9, 0x77, 0xd0,
	0x73, 0xc4, 0x61, 0xd7, 0xee, 0xfc, 0xdd, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0xa6, 0xb5,
	0x3f, 0x7a, 0x04, 0x00, 0x00,
}
