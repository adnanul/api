// Code generated by protoc-gen-grpc-gateway-cors
// source: credential.proto
// DO NOT EDIT!

/*
Package v1beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ExportCredentialsCorsPatterns returns an array of grpc gatway mux patterns for
// Credentials service to enable CORS.
func ExportCredentialsCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Credentials_List_0,
		pattern_Credentials_Create_0,
		pattern_Credentials_Update_0,
		pattern_Credentials_IsAuthorized_0,
		pattern_Credentials_Delete_0,
	}
}