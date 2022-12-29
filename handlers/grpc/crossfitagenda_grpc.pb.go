// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: handlers/grpc/crossfitagenda.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DefaultServiceClient is the client API for DefaultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DefaultServiceClient interface {
	StartCrossfitAgenda(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
}

type defaultServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDefaultServiceClient(cc grpc.ClientConnInterface) DefaultServiceClient {
	return &defaultServiceClient{cc}
}

func (c *defaultServiceClient) StartCrossfitAgenda(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/crossfitagenda.DefaultService/StartCrossfitAgenda", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/crossfitagenda.DefaultService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DefaultServiceServer is the server API for DefaultService service.
// All implementations should embed UnimplementedDefaultServiceServer
// for forward compatibility
type DefaultServiceServer interface {
	StartCrossfitAgenda(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Status(context.Context, *emptypb.Empty) (*StatusResponse, error)
}

// UnimplementedDefaultServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDefaultServiceServer struct {
}

func (UnimplementedDefaultServiceServer) StartCrossfitAgenda(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartCrossfitAgenda not implemented")
}
func (UnimplementedDefaultServiceServer) Status(context.Context, *emptypb.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}

// UnsafeDefaultServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DefaultServiceServer will
// result in compilation errors.
type UnsafeDefaultServiceServer interface {
	mustEmbedUnimplementedDefaultServiceServer()
}

func RegisterDefaultServiceServer(s grpc.ServiceRegistrar, srv DefaultServiceServer) {
	s.RegisterService(&DefaultService_ServiceDesc, srv)
}

func _DefaultService_StartCrossfitAgenda_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).StartCrossfitAgenda(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crossfitagenda.DefaultService/StartCrossfitAgenda",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).StartCrossfitAgenda(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crossfitagenda.DefaultService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DefaultService_ServiceDesc is the grpc.ServiceDesc for DefaultService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DefaultService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crossfitagenda.DefaultService",
	HandlerType: (*DefaultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartCrossfitAgenda",
			Handler:    _DefaultService_StartCrossfitAgenda_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _DefaultService_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "handlers/grpc/crossfitagenda.proto",
}
