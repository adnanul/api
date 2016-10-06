// Code generated by protoc-gen-go.
// source: cloudcredential.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	cloudcredential.proto

It has these top-level messages:
	CloudCredentialCreateRequest
	CloudCredentialUpdateRequest
	CloudCredentialDeleteRequest
	CloudCredentialListRequest
	CloudCredentialListResponse
	Credential
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

type CloudCredentialCreateRequest struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Data     map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CloudCredentialCreateRequest) Reset()                    { *m = CloudCredentialCreateRequest{} }
func (m *CloudCredentialCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialCreateRequest) ProtoMessage()               {}
func (*CloudCredentialCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CloudCredentialCreateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CloudCredentialUpdateRequest struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Data     map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CloudCredentialUpdateRequest) Reset()                    { *m = CloudCredentialUpdateRequest{} }
func (m *CloudCredentialUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialUpdateRequest) ProtoMessage()               {}
func (*CloudCredentialUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CloudCredentialUpdateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CloudCredentialDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CloudCredentialDeleteRequest) Reset()                    { *m = CloudCredentialDeleteRequest{} }
func (m *CloudCredentialDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialDeleteRequest) ProtoMessage()               {}
func (*CloudCredentialDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CloudCredentialListRequest struct {
}

func (m *CloudCredentialListRequest) Reset()                    { *m = CloudCredentialListRequest{} }
func (m *CloudCredentialListRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialListRequest) ProtoMessage()               {}
func (*CloudCredentialListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CloudCredentialListResponse struct {
	Status      *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Credentials []*Credential  `protobuf:"bytes,2,rep,name=credentials" json:"credentials,omitempty"`
}

func (m *CloudCredentialListResponse) Reset()                    { *m = CloudCredentialListResponse{} }
func (m *CloudCredentialListResponse) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialListResponse) ProtoMessage()               {}
func (*CloudCredentialListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CloudCredentialListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CloudCredentialListResponse) GetCredentials() []*Credential {
	if m != nil {
		return m.Credentials
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
func (*Credential) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*CloudCredentialCreateRequest)(nil), "credential.v1beta1.CloudCredentialCreateRequest")
	proto.RegisterType((*CloudCredentialUpdateRequest)(nil), "credential.v1beta1.CloudCredentialUpdateRequest")
	proto.RegisterType((*CloudCredentialDeleteRequest)(nil), "credential.v1beta1.CloudCredentialDeleteRequest")
	proto.RegisterType((*CloudCredentialListRequest)(nil), "credential.v1beta1.CloudCredentialListRequest")
	proto.RegisterType((*CloudCredentialListResponse)(nil), "credential.v1beta1.CloudCredentialListResponse")
	proto.RegisterType((*Credential)(nil), "credential.v1beta1.Credential")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for CloudCredentials service

type CloudCredentialsClient interface {
	List(ctx context.Context, in *CloudCredentialListRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error)
	Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type cloudCredentialsClient struct {
	cc *grpc.ClientConn
}

func NewCloudCredentialsClient(cc *grpc.ClientConn) CloudCredentialsClient {
	return &cloudCredentialsClient{cc}
}

func (c *cloudCredentialsClient) List(ctx context.Context, in *CloudCredentialListRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error) {
	out := new(CloudCredentialListResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialsClient) Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialsClient) Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialsClient) Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CloudCredentials service

type CloudCredentialsServer interface {
	List(context.Context, *CloudCredentialListRequest) (*CloudCredentialListResponse, error)
	Create(context.Context, *CloudCredentialCreateRequest) (*dtypes.VoidResponse, error)
	Update(context.Context, *CloudCredentialUpdateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *CloudCredentialDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterCloudCredentialsServer(s *grpc.Server, srv CloudCredentialsServer) {
	s.RegisterService(&_CloudCredentials_serviceDesc, srv)
}

func _CloudCredentials_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).List(ctx, req.(*CloudCredentialListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredentials_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).Create(ctx, req.(*CloudCredentialCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredentials_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).Update(ctx, req.(*CloudCredentialUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredentials_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).Delete(ctx, req.(*CloudCredentialDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudCredentials_serviceDesc = grpc.ServiceDesc{
	ServiceName: "credential.v1beta1.CloudCredentials",
	HandlerType: (*CloudCredentialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CloudCredentials_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CloudCredentials_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CloudCredentials_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CloudCredentials_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("cloudcredential.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x94, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x99, 0x36, 0x46, 0xfb, 0x02, 0xb2, 0x0c, 0xab, 0x84, 0x58, 0xb4, 0xe4, 0xa0, 0x75,
	0x91, 0xcc, 0x6e, 0x16, 0xfc, 0xd1, 0x93, 0xda, 0xf5, 0x26, 0x1e, 0x22, 0x7a, 0xf0, 0x22, 0xb3,
	0x9d, 0xd9, 0x65, 0x30, 0xcd, 0xc4, 0xcc, 0xb4, 0x50, 0x44, 0x04, 0x11, 0x14, 0x3c, 0x7a, 0xf1,
	0x9f, 0xf2, 0xe4, 0xc1, 0x7f, 0xc0, 0x3f, 0x44, 0x32, 0x93, 0xcd, 0xa6, 0xbb, 0xdd, 0x25, 0x05,
	0xd9, 0x4b, 0x49, 0xbe, 0xaf, 0xef, 0xcd, 0xe7, 0x7d, 0xe7, 0xbd, 0xc0, 0xb5, 0x49, 0x2a, 0x67,
	0x6c, 0x52, 0x70, 0xc6, 0x33, 0x2d, 0x68, 0x1a, 0xe5, 0x85, 0xd4, 0x12, 0xe3, 0x86, 0x32, 0xdf,
	0xd9, 0xe7, 0x9a, 0xee, 0x04, 0xfd, 0x43, 0x29, 0x0f, 0x53, 0x4e, 0x68, 0x2e, 0x08, 0xcd, 0x32,
	0xa9, 0xa9, 0x16, 0x32, 0x53, 0x36, 0x23, 0xb8, 0x5e, 0xca, 0x4c, 0x2f, 0x72, 0xae, 0x88, 0xf9,
	0xb5, 0x7a, 0xf8, 0x07, 0x41, 0x7f, 0x5c, 0x9e, 0x31, 0xae, 0x2b, 0x8e, 0x0b, 0x4e, 0x35, 0x4f,
	0xf8, 0xfb, 0x19, 0x57, 0x1a, 0x63, 0x70, 0x32, 0x3a, 0xe5, 0x3e, 0x1a, 0xa0, 0x61, 0x2f, 0x31,
	0xcf, 0x38, 0x80, 0x2b, 0x79, 0x21, 0xe7, 0x82, 0xf1, 0xc2, 0xef, 0x18, 0xbd, 0x7e, 0xc7, 0x2f,
	0xc0, 0x61, 0x54, 0x53, 0xbf, 0x3b, 0xe8, 0x0e, 0xbd, 0x78, 0x14, 0x9d, 0x26, 0x8d, 0xce, 0x3b,
	0x2f, 0xda, 0xa3, 0x9a, 0x3e, 0xcb, 0x74, 0xb1, 0x48, 0x4c, 0x9d, 0xe0, 0x01, 0xf4, 0x6a, 0x09,
	0x6f, 0x40, 0xf7, 0x1d, 0x5f, 0x54, 0x2c, 0xe5, 0x23, 0xde, 0x84, 0x4b, 0x73, 0x9a, 0xce, 0x78,
	0xc5, 0x61, 0x5f, 0x46, 0x9d, 0x87, 0x68, 0x55, 0x67, 0xaf, 0x72, 0x76, 0xa1, 0x9d, 0x2d, 0x9d,
	0xf7, 0xff, 0x3a, 0x8b, 0x4f, 0x35, 0xb6, 0xc7, 0x53, 0x7e, 0x6e, 0x63, 0x61, 0x1f, 0x82, 0x13,
	0x39, 0xcf, 0x85, 0xd2, 0x55, 0x46, 0xf8, 0x15, 0xc1, 0x8d, 0x95, 0x61, 0x95, 0xcb, 0x4c, 0x71,
	0x7c, 0x1b, 0x5c, 0xa5, 0xa9, 0x9e, 0x29, 0x53, 0xd3, 0x8b, 0xaf, 0x46, 0x76, 0x94, 0xa2, 0x97,
	0x46, 0x4d, 0xaa, 0x28, 0x7e, 0x0c, 0xde, 0xb1, 0x2b, 0xca, 0xef, 0x18, 0xa7, 0x6e, 0xae, 0x74,
	0xaa, 0x96, 0x92, 0x66, 0x4a, 0xf8, 0x09, 0xe0, 0x38, 0xb4, 0xf6, 0x15, 0x0d, 0xc0, 0x13, 0xd9,
	0x81, 0x2c, 0xa6, 0x66, 0xf6, 0xfd, 0xae, 0x09, 0x37, 0x25, 0x7c, 0x0b, 0xbc, 0xa9, 0x64, 0xe2,
	0x40, 0x70, 0xf6, 0x96, 0x6a, 0xdf, 0x19, 0xa0, 0x61, 0x37, 0x81, 0x23, 0xe9, 0x89, 0x8e, 0x7f,
	0x39, 0xb0, 0x71, 0xc2, 0x0a, 0x85, 0x7f, 0x22, 0x70, 0x4a, 0x43, 0x70, 0xd4, 0xe2, 0xd6, 0x1b,
	0xc6, 0x06, 0xa4, 0xf5, 0xff, 0xad, 0xd3, 0x21, 0xf9, 0xfc, 0xfb, 0xef, 0x8f, 0xce, 0x5d, 0x7c,
	0x87, 0xd0, 0x3c, 0x57, 0x13, 0xc9, 0xec, 0x42, 0x9b, 0xcf, 0x00, 0xa9, 0x0a, 0x90, 0x86, 0x61,
	0xf8, 0x0b, 0x02, 0xd7, 0x6e, 0x10, 0xde, 0x5e, 0x77, 0xd9, 0x82, 0xcd, 0xa3, 0x7b, 0x7c, 0x2d,
	0x05, 0xab, 0x19, 0x62, 0xc3, 0x70, 0x2f, 0x6c, 0xcb, 0x30, 0x42, 0x5b, 0xf8, 0x3b, 0x02, 0xd7,
	0x8e, 0x7b, 0x2b, 0x8c, 0xa5, 0xcd, 0x38, 0x03, 0xe3, 0x91, 0xc1, 0xd8, 0x0d, 0xa2, 0x96, 0x18,
	0xe4, 0x43, 0x39, 0x1f, 0x1f, 0x4b, 0x9a, 0x6f, 0x08, 0x5c, 0xbb, 0x13, 0xad, 0x68, 0x96, 0xd6,
	0xe7, 0x0c, 0x9a, 0xfb, 0x86, 0x66, 0x7b, 0x6b, 0x4d, 0x9a, 0xa7, 0xbd, 0x37, 0x97, 0xab, 0xe8,
	0xbe, 0x6b, 0x3e, 0xb9, 0xbb, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x54, 0x28, 0x56, 0xd5,
	0x05, 0x00, 0x00,
}