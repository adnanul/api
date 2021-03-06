// Code generated by protoc-gen-go. DO NOT EDIT.
// source: credential.proto

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	credential.proto

It has these top-level messages:
	CredentialCreateRequest
	CredentialUpdateRequest
	CredentialIsAuthorizedRequest
	CredentialIsAuthorizedResponse
	CredentialDeleteRequest
	CredentialDescribeRequest
	CredentialListResponse
	CredentialDescribeResponse
	Credential
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

type CredentialCreateRequest struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Data     map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CredentialCreateRequest) Reset()                    { *m = CredentialCreateRequest{} }
func (m *CredentialCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CredentialCreateRequest) ProtoMessage()               {}
func (*CredentialCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CredentialCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CredentialCreateRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *CredentialCreateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CredentialUpdateRequest struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Data     map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CredentialUpdateRequest) Reset()                    { *m = CredentialUpdateRequest{} }
func (m *CredentialUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*CredentialUpdateRequest) ProtoMessage()               {}
func (*CredentialUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CredentialUpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CredentialUpdateRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *CredentialUpdateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CredentialIsAuthorizedRequest struct {
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	GceProject string `protobuf:"bytes,2,opt,name=gce_project,json=gceProject" json:"gce_project,omitempty"`
}

func (m *CredentialIsAuthorizedRequest) Reset()                    { *m = CredentialIsAuthorizedRequest{} }
func (m *CredentialIsAuthorizedRequest) String() string            { return proto.CompactTextString(m) }
func (*CredentialIsAuthorizedRequest) ProtoMessage()               {}
func (*CredentialIsAuthorizedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CredentialIsAuthorizedRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CredentialIsAuthorizedRequest) GetGceProject() string {
	if m != nil {
		return m.GceProject
	}
	return ""
}

type CredentialIsAuthorizedResponse struct {
	Unauthorized bool   `protobuf:"varint,1,opt,name=unauthorized" json:"unauthorized,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *CredentialIsAuthorizedResponse) Reset()                    { *m = CredentialIsAuthorizedResponse{} }
func (m *CredentialIsAuthorizedResponse) String() string            { return proto.CompactTextString(m) }
func (*CredentialIsAuthorizedResponse) ProtoMessage()               {}
func (*CredentialIsAuthorizedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CredentialIsAuthorizedResponse) GetUnauthorized() bool {
	if m != nil {
		return m.Unauthorized
	}
	return false
}

func (m *CredentialIsAuthorizedResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CredentialDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CredentialDeleteRequest) Reset()                    { *m = CredentialDeleteRequest{} }
func (m *CredentialDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*CredentialDeleteRequest) ProtoMessage()               {}
func (*CredentialDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CredentialDeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CredentialDescribeRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CredentialDescribeRequest) Reset()                    { *m = CredentialDescribeRequest{} }
func (m *CredentialDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*CredentialDescribeRequest) ProtoMessage()               {}
func (*CredentialDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CredentialDescribeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CredentialListResponse struct {
	Credentials []*Credential `protobuf:"bytes,1,rep,name=credentials" json:"credentials,omitempty"`
}

func (m *CredentialListResponse) Reset()                    { *m = CredentialListResponse{} }
func (m *CredentialListResponse) String() string            { return proto.CompactTextString(m) }
func (*CredentialListResponse) ProtoMessage()               {}
func (*CredentialListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CredentialListResponse) GetCredentials() []*Credential {
	if m != nil {
		return m.Credentials
	}
	return nil
}

type CredentialDescribeResponse struct {
	Credential *Credential `protobuf:"bytes,1,opt,name=credential" json:"credential,omitempty"`
}

func (m *CredentialDescribeResponse) Reset()                    { *m = CredentialDescribeResponse{} }
func (m *CredentialDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*CredentialDescribeResponse) ProtoMessage()               {}
func (*CredentialDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CredentialDescribeResponse) GetCredential() *Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

type Credential struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider    string `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Information string `protobuf:"bytes,3,opt,name=information" json:"information,omitempty"`
	ModifiedAt  int64  `protobuf:"varint,4,opt,name=modified_at,json=modifiedAt" json:"modified_at,omitempty"`
}

func (m *Credential) Reset()                    { *m = Credential{} }
func (m *Credential) String() string            { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()               {}
func (*Credential) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Credential) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Credential) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Credential) GetInformation() string {
	if m != nil {
		return m.Information
	}
	return ""
}

func (m *Credential) GetModifiedAt() int64 {
	if m != nil {
		return m.ModifiedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*CredentialCreateRequest)(nil), "appscode.credential.v1beta1.CredentialCreateRequest")
	proto.RegisterType((*CredentialUpdateRequest)(nil), "appscode.credential.v1beta1.CredentialUpdateRequest")
	proto.RegisterType((*CredentialIsAuthorizedRequest)(nil), "appscode.credential.v1beta1.CredentialIsAuthorizedRequest")
	proto.RegisterType((*CredentialIsAuthorizedResponse)(nil), "appscode.credential.v1beta1.CredentialIsAuthorizedResponse")
	proto.RegisterType((*CredentialDeleteRequest)(nil), "appscode.credential.v1beta1.CredentialDeleteRequest")
	proto.RegisterType((*CredentialDescribeRequest)(nil), "appscode.credential.v1beta1.CredentialDescribeRequest")
	proto.RegisterType((*CredentialListResponse)(nil), "appscode.credential.v1beta1.CredentialListResponse")
	proto.RegisterType((*CredentialDescribeResponse)(nil), "appscode.credential.v1beta1.CredentialDescribeResponse")
	proto.RegisterType((*Credential)(nil), "appscode.credential.v1beta1.Credential")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Credentials service

type CredentialsClient interface {
	List(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*CredentialListResponse, error)
	Describe(ctx context.Context, in *CredentialDescribeRequest, opts ...grpc.CallOption) (*CredentialDescribeResponse, error)
	Create(ctx context.Context, in *CredentialCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Update(ctx context.Context, in *CredentialUpdateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	IsAuthorized(ctx context.Context, in *CredentialIsAuthorizedRequest, opts ...grpc.CallOption) (*CredentialIsAuthorizedResponse, error)
	Delete(ctx context.Context, in *CredentialDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type credentialsClient struct {
	cc *grpc.ClientConn
}

func NewCredentialsClient(cc *grpc.ClientConn) CredentialsClient {
	return &credentialsClient{cc}
}

func (c *credentialsClient) List(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*CredentialListResponse, error) {
	out := new(CredentialListResponse)
	err := grpc.Invoke(ctx, "/appscode.credential.v1beta1.Credentials/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsClient) Describe(ctx context.Context, in *CredentialDescribeRequest, opts ...grpc.CallOption) (*CredentialDescribeResponse, error) {
	out := new(CredentialDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.credential.v1beta1.Credentials/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsClient) Create(ctx context.Context, in *CredentialCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.credential.v1beta1.Credentials/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsClient) Update(ctx context.Context, in *CredentialUpdateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.credential.v1beta1.Credentials/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsClient) IsAuthorized(ctx context.Context, in *CredentialIsAuthorizedRequest, opts ...grpc.CallOption) (*CredentialIsAuthorizedResponse, error) {
	out := new(CredentialIsAuthorizedResponse)
	err := grpc.Invoke(ctx, "/appscode.credential.v1beta1.Credentials/IsAuthorized", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsClient) Delete(ctx context.Context, in *CredentialDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.credential.v1beta1.Credentials/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Credentials service

type CredentialsServer interface {
	List(context.Context, *appscode_dtypes.VoidRequest) (*CredentialListResponse, error)
	Describe(context.Context, *CredentialDescribeRequest) (*CredentialDescribeResponse, error)
	Create(context.Context, *CredentialCreateRequest) (*appscode_dtypes.VoidResponse, error)
	Update(context.Context, *CredentialUpdateRequest) (*appscode_dtypes.VoidResponse, error)
	IsAuthorized(context.Context, *CredentialIsAuthorizedRequest) (*CredentialIsAuthorizedResponse, error)
	Delete(context.Context, *CredentialDeleteRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterCredentialsServer(s *grpc.Server, srv CredentialsServer) {
	s.RegisterService(&_Credentials_serviceDesc, srv)
}

func _Credentials_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.credential.v1beta1.Credentials/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).List(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credentials_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.credential.v1beta1.Credentials/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Describe(ctx, req.(*CredentialDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credentials_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.credential.v1beta1.Credentials/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Create(ctx, req.(*CredentialCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credentials_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.credential.v1beta1.Credentials/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Update(ctx, req.(*CredentialUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credentials_IsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialIsAuthorizedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).IsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.credential.v1beta1.Credentials/IsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).IsAuthorized(ctx, req.(*CredentialIsAuthorizedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credentials_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.credential.v1beta1.Credentials/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServer).Delete(ctx, req.(*CredentialDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Credentials_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.credential.v1beta1.Credentials",
	HandlerType: (*CredentialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Credentials_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Credentials_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Credentials_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Credentials_Update_Handler,
		},
		{
			MethodName: "IsAuthorized",
			Handler:    _Credentials_IsAuthorized_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Credentials_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credential.proto",
}

func init() { proto.RegisterFile("credential.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0x66, 0x92, 0x98, 0xa6, 0x27, 0x05, 0xcb, 0x20, 0xba, 0xae, 0xfd, 0x09, 0xeb, 0x85, 0xa5,
	0xd0, 0x2c, 0xfd, 0xa1, 0xd5, 0x16, 0x0a, 0x4d, 0x5b, 0x4a, 0xc1, 0x8b, 0xb2, 0xfe, 0x5c, 0x78,
	0x61, 0x99, 0xee, 0x4e, 0xe3, 0xd4, 0x64, 0x67, 0xdd, 0x99, 0x14, 0xaa, 0x88, 0xd0, 0x57, 0x10,
	0xbc, 0xf3, 0x42, 0xf0, 0x09, 0x04, 0xaf, 0x7d, 0x08, 0xc1, 0x2b, 0x2f, 0x7d, 0x10, 0xd9, 0xff,
	0xd9, 0x92, 0xa4, 0x9b, 0x20, 0xde, 0x84, 0xcc, 0x39, 0x73, 0xce, 0xf9, 0xbe, 0xf9, 0x0e, 0x1f,
	0x0b, 0xd3, 0xb6, 0x4f, 0x1d, 0xea, 0x4a, 0x46, 0x3a, 0x4d, 0xcf, 0xe7, 0x92, 0xe3, 0x7b, 0xc4,
	0xf3, 0x84, 0xcd, 0x1d, 0xda, 0x54, 0x52, 0xe7, 0xcb, 0x27, 0x54, 0x92, 0x65, 0x7d, 0xa6, 0xcd,
	0x79, 0xbb, 0x43, 0x4d, 0xe2, 0x31, 0x93, 0xb8, 0x2e, 0x97, 0x44, 0x32, 0xee, 0x8a, 0xa8, 0x54,
	0x9f, 0x4b, 0x4a, 0x07, 0xe4, 0xe7, 0x73, 0x79, 0x47, 0x5e, 0x78, 0x54, 0x98, 0xe1, 0x6f, 0x74,
	0xc1, 0xf8, 0x85, 0xe0, 0xce, 0x6e, 0x3a, 0x75, 0xd7, 0xa7, 0x44, 0x52, 0x8b, 0xbe, 0xe9, 0x51,
	0x21, 0x31, 0x86, 0x8a, 0x4b, 0xba, 0x54, 0x43, 0x0d, 0xb4, 0x30, 0x69, 0x85, 0xff, 0xb1, 0x0e,
	0x35, 0xcf, 0xe7, 0xe7, 0xcc, 0xa1, 0xbe, 0x56, 0x0a, 0xe3, 0xe9, 0x19, 0x5b, 0x50, 0x71, 0x88,
	0x24, 0x5a, 0xb9, 0x51, 0x5e, 0xa8, 0xaf, 0x6c, 0x37, 0x87, 0xd0, 0x6a, 0x0e, 0x98, 0xd9, 0xdc,
	0x23, 0x92, 0xec, 0xbb, 0xd2, 0xbf, 0xb0, 0xc2, 0x5e, 0xfa, 0x06, 0x4c, 0xa6, 0x21, 0x3c, 0x0d,
	0xe5, 0xd7, 0xf4, 0x22, 0xc6, 0x13, 0xfc, 0xc5, 0xb7, 0xe0, 0xc6, 0x39, 0xe9, 0xf4, 0x68, 0x8c,
	0x25, 0x3a, 0x6c, 0x96, 0x1e, 0xa2, 0x2b, 0xc4, 0x9e, 0x79, 0xce, 0x7f, 0x27, 0x96, 0x9b, 0xf9,
	0xef, 0x88, 0x3d, 0x85, 0xd9, 0x6c, 0xc6, 0xa1, 0xd8, 0xe9, 0xc9, 0x57, 0xdc, 0x67, 0x6f, 0xa9,
	0x33, 0x8c, 0xdd, 0x3c, 0xd4, 0xdb, 0x36, 0x3d, 0xf6, 0x7c, 0x7e, 0x46, 0x6d, 0x19, 0x37, 0x85,
	0xb6, 0x4d, 0x8f, 0xa2, 0x88, 0xf1, 0x12, 0xe6, 0x06, 0x75, 0x15, 0x1e, 0x77, 0x05, 0xc5, 0x06,
	0x4c, 0xf5, 0x5c, 0x92, 0xc6, 0xc3, 0xf6, 0x35, 0x2b, 0x17, 0xc3, 0x1a, 0x4c, 0x74, 0xa9, 0x10,
	0xa4, 0x9d, 0xe0, 0x4e, 0x8e, 0xc6, 0x92, 0xaa, 0xc6, 0x1e, 0xed, 0xd0, 0xa1, 0x6a, 0x18, 0x26,
	0xdc, 0x55, 0xaf, 0x0b, 0xdb, 0x67, 0x27, 0x43, 0x0b, 0x6c, 0xb8, 0x9d, 0x15, 0x3c, 0x66, 0x42,
	0xa6, 0xb8, 0x0f, 0xa1, 0x9e, 0xc9, 0x24, 0x34, 0x14, 0x6a, 0xf8, 0xa0, 0xa0, 0x86, 0x96, 0x5a,
	0x6b, 0x50, 0xd0, 0xfb, 0xa1, 0x8a, 0x07, 0x1d, 0x00, 0x64, 0x97, 0x43, 0x70, 0x23, 0xcc, 0x51,
	0x4a, 0x8d, 0x0f, 0x00, 0x59, 0x66, 0xe4, 0x65, 0x6d, 0x40, 0x9d, 0xb9, 0xa7, 0xdc, 0xef, 0x86,
	0x46, 0xa0, 0x95, 0xc3, 0xb4, 0x1a, 0x0a, 0x96, 0xa1, 0xcb, 0x1d, 0x76, 0xca, 0xa8, 0x73, 0x4c,
	0xa4, 0x56, 0x69, 0xa0, 0x85, 0xb2, 0x05, 0x49, 0x68, 0x47, 0xae, 0x7c, 0xab, 0x41, 0x3d, 0x43,
	0x20, 0xf0, 0x27, 0x04, 0x95, 0xe0, 0x4d, 0xf1, 0x4c, 0x46, 0x27, 0xf2, 0x92, 0xe6, 0x73, 0xce,
	0x92, 0xc5, 0xd3, 0x57, 0x0b, 0x92, 0x55, 0xe5, 0x31, 0xb6, 0x2e, 0xbf, 0x6b, 0xa5, 0x1a, 0xba,
	0xfc, 0xf9, 0xe7, 0x63, 0xc9, 0xc4, 0x4b, 0xe6, 0x71, 0xce, 0xb1, 0xec, 0x0e, 0xef, 0x39, 0x66,
	0xdc, 0xc2, 0x54, 0xb4, 0x30, 0xcf, 0x04, 0x77, 0xf1, 0x0f, 0x04, 0xb5, 0x44, 0x07, 0xbc, 0x5e,
	0x70, 0xfc, 0x95, 0x75, 0xd2, 0x37, 0x46, 0xae, 0x8b, 0xa1, 0xb7, 0x14, 0xe8, 0xeb, 0x78, 0xad,
	0x30, 0xf4, 0x77, 0x81, 0x86, 0xef, 0x23, 0x06, 0x9f, 0x11, 0x54, 0x23, 0x07, 0xc4, 0x6b, 0xe3,
	0x18, 0xa6, 0x3e, 0x3b, 0x40, 0x92, 0x18, 0xe3, 0xb6, 0x82, 0x71, 0xc5, 0x18, 0xed, 0x79, 0x37,
	0xd1, 0x22, 0xfe, 0x8a, 0xa0, 0x1a, 0x19, 0x59, 0x61, 0x7c, 0x39, 0xdf, 0xbb, 0x0e, 0xdf, 0xbe,
	0x82, 0xef, 0x91, 0x3e, 0xd6, 0x1b, 0x06, 0x30, 0x7f, 0x23, 0x98, 0x52, 0x5d, 0x0b, 0x6f, 0x16,
	0x04, 0xdb, 0xc7, 0x40, 0xf5, 0xad, 0xb1, 0x6a, 0x63, 0x42, 0x4f, 0x14, 0x42, 0x07, 0x46, 0x6b,
	0x54, 0x42, 0x4c, 0x2c, 0x65, 0x7e, 0x9a, 0xd2, 0xfb, 0x82, 0xa0, 0x1a, 0x99, 0x66, 0x61, 0x15,
	0x72, 0x1e, 0x7b, 0x9d, 0x0a, 0xb9, 0x4d, 0x5e, 0x1c, 0x4b, 0x85, 0xd6, 0x0e, 0xdc, 0xb7, 0x79,
	0x37, 0x9b, 0x43, 0x3c, 0xd6, 0x07, 0x61, 0xeb, 0x66, 0x06, 0xf1, 0x28, 0xf8, 0x02, 0x39, 0x42,
	0x2f, 0x26, 0xe2, 0xdc, 0x49, 0x35, 0xfc, 0x26, 0x59, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x36,
	0x86, 0x95, 0xea, 0x23, 0x09, 0x00, 0x00,
}
