// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0--rc2
// source: controller/proto/url.proto

package controller

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Url_Save_FullMethodName = "/proto.Url/Save"
	Url_Get_FullMethodName  = "/proto.Url/Get"
)

// UrlClient is the client API for Url service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlClient interface {
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type urlClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlClient(cc grpc.ClientConnInterface) UrlClient {
	return &urlClient{cc}
}

func (c *urlClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, Url_Save_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Url_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlServer is the server API for Url service.
// All implementations must embed UnimplementedUrlServer
// for forward compatibility
type UrlServer interface {
	Save(context.Context, *SaveRequest) (*SaveResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedUrlServer()
}

// UnimplementedUrlServer must be embedded to have forward compatible implementations.
type UnimplementedUrlServer struct {
}

func (UnimplementedUrlServer) Save(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedUrlServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUrlServer) mustEmbedUnimplementedUrlServer() {}

// UnsafeUrlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlServer will
// result in compilation errors.
type UnsafeUrlServer interface {
	mustEmbedUnimplementedUrlServer()
}

func RegisterUrlServer(s grpc.ServiceRegistrar, srv UrlServer) {
	s.RegisterService(&Url_ServiceDesc, srv)
}

func _Url_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Url_Save_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Url_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Url_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Url_ServiceDesc is the grpc.ServiceDesc for Url service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Url_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Url",
	HandlerType: (*UrlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Url_Save_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Url_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/proto/url.proto",
}
