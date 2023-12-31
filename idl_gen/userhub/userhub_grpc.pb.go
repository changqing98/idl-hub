// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.9.2
// source: userhub/userhub.proto

package userhub

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

// UserhubServiceClient is the client API for UserhubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserhubServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type userhubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserhubServiceClient(cc grpc.ClientConnInterface) UserhubServiceClient {
	return &userhubServiceClient{cc}
}

func (c *userhubServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/UserhubService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserhubServiceServer is the server API for UserhubService service.
// All implementations must embed UnimplementedUserhubServiceServer
// for forward compatibility
type UserhubServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedUserhubServiceServer()
}

// UnimplementedUserhubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserhubServiceServer struct {
}

func (UnimplementedUserhubServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserhubServiceServer) mustEmbedUnimplementedUserhubServiceServer() {}

// UnsafeUserhubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserhubServiceServer will
// result in compilation errors.
type UnsafeUserhubServiceServer interface {
	mustEmbedUnimplementedUserhubServiceServer()
}

func RegisterUserhubServiceServer(s grpc.ServiceRegistrar, srv UserhubServiceServer) {
	s.RegisterService(&UserhubService_ServiceDesc, srv)
}

func _UserhubService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserhubServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserhubService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserhubServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserhubService_ServiceDesc is the grpc.ServiceDesc for UserhubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserhubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserhubService",
	HandlerType: (*UserhubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserhubService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userhub/userhub.proto",
}
