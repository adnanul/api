// Code generated by protoc-gen-grpc-gateway-cors
// source: ca.proto
// DO NOT EDIT!

/*
Package v1beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ExportCAsCorsPatterns returns an array of grpc gatway mux patterns for
// CAs service to enable CORS.
func ExportCAsCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_CAs_Create_0,
	}
}