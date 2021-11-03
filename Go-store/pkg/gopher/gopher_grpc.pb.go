// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package __

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

// GopherClient is the client API for Gopher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GopherClient interface {
	// create
	CreateQuery(ctx context.Context, in *GopherRequest, opts ...grpc.CallOption) (*GopherReply, error)
	// update
	UpdateQuery(ctx context.Context, in *GopherRequest, opts ...grpc.CallOption) (*GopherReply, error)
	// GET
	GetQuery(ctx context.Context, in *GopherRequest, opts ...grpc.CallOption) (*GopherReply, error)
}

type gopherClient struct {
	cc grpc.ClientConnInterface
}

func NewGopherClient(cc grpc.ClientConnInterface) GopherClient {
	return &gopherClient{cc}
}

func (c *gopherClient) CreateQuery(ctx context.Context, in *GopherRequest, opts ...grpc.CallOption) (*GopherReply, error) {
	out := new(GopherReply)
	err := c.cc.Invoke(ctx, "/gopher.Gopher/CreateQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gopherClient) UpdateQuery(ctx context.Context, in *GopherRequest, opts ...grpc.CallOption) (*GopherReply, error) {
	out := new(GopherReply)
	err := c.cc.Invoke(ctx, "/gopher.Gopher/UpdateQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gopherClient) GetQuery(ctx context.Context, in *GopherRequest, opts ...grpc.CallOption) (*GopherReply, error) {
	out := new(GopherReply)
	err := c.cc.Invoke(ctx, "/gopher.Gopher/GetQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GopherServer is the server API for Gopher service.
// All implementations must embed UnimplementedGopherServer
// for forward compatibility
type GopherServer interface {
	// create
	CreateQuery(context.Context, *GopherRequest) (*GopherReply, error)
	// update
	UpdateQuery(context.Context, *GopherRequest) (*GopherReply, error)
	// GET
	GetQuery(context.Context, *GopherRequest) (*GopherReply, error)
	mustEmbedUnimplementedGopherServer()
}

// UnimplementedGopherServer must be embedded to have forward compatible implementations.
type UnimplementedGopherServer struct {
}

func (UnimplementedGopherServer) CreateQuery(context.Context, *GopherRequest) (*GopherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuery not implemented")
}
func (UnimplementedGopherServer) UpdateQuery(context.Context, *GopherRequest) (*GopherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuery not implemented")
}
func (UnimplementedGopherServer) GetQuery(context.Context, *GopherRequest) (*GopherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuery not implemented")
}
func (UnimplementedGopherServer) mustEmbedUnimplementedGopherServer() {}

// UnsafeGopherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GopherServer will
// result in compilation errors.
type UnsafeGopherServer interface {
	mustEmbedUnimplementedGopherServer()
}

func RegisterGopherServer(s grpc.ServiceRegistrar, srv GopherServer) {
	s.RegisterService(&Gopher_ServiceDesc, srv)
}

func _Gopher_CreateQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GopherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GopherServer).CreateQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gopher.Gopher/CreateQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GopherServer).CreateQuery(ctx, req.(*GopherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gopher_UpdateQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GopherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GopherServer).UpdateQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gopher.Gopher/UpdateQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GopherServer).UpdateQuery(ctx, req.(*GopherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gopher_GetQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GopherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GopherServer).GetQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gopher.Gopher/GetQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GopherServer).GetQuery(ctx, req.(*GopherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gopher_ServiceDesc is the grpc.ServiceDesc for Gopher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gopher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gopher.Gopher",
	HandlerType: (*GopherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuery",
			Handler:    _Gopher_CreateQuery_Handler,
		},
		{
			MethodName: "UpdateQuery",
			Handler:    _Gopher_UpdateQuery_Handler,
		},
		{
			MethodName: "GetQuery",
			Handler:    _Gopher_GetQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gopher.proto",
}