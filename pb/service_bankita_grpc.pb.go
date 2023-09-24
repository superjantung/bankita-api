// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: service_bankita.proto

package pb

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

// BankitaClient is the client API for Bankita service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankitaClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type bankitaClient struct {
	cc grpc.ClientConnInterface
}

func NewBankitaClient(cc grpc.ClientConnInterface) BankitaClient {
	return &bankitaClient{cc}
}

func (c *bankitaClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.Bankita/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankitaClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/pb.Bankita/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankitaServer is the server API for Bankita service.
// All implementations must embed UnimplementedBankitaServer
// for forward compatibility
type BankitaServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedBankitaServer()
}

// UnimplementedBankitaServer must be embedded to have forward compatible implementations.
type UnimplementedBankitaServer struct {
}

func (UnimplementedBankitaServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedBankitaServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedBankitaServer) mustEmbedUnimplementedBankitaServer() {}

// UnsafeBankitaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankitaServer will
// result in compilation errors.
type UnsafeBankitaServer interface {
	mustEmbedUnimplementedBankitaServer()
}

func RegisterBankitaServer(s grpc.ServiceRegistrar, srv BankitaServer) {
	s.RegisterService(&Bankita_ServiceDesc, srv)
}

func _Bankita_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankitaServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Bankita/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankitaServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bankita_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankitaServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Bankita/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankitaServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bankita_ServiceDesc is the grpc.ServiceDesc for Bankita service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bankita_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Bankita",
	HandlerType: (*BankitaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Bankita_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Bankita_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_bankita.proto",
}
