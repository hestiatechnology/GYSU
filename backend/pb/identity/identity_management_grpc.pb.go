// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: identity_management.proto

package identity

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	common "hestia/jobfair/api/pb/common"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	IdentityManagement_Login_FullMethodName    = "/hestia.jobfair.v1.identity.IdentityManagement/Login"
	IdentityManagement_Register_FullMethodName = "/hestia.jobfair.v1.identity.IdentityManagement/Register"
	IdentityManagement_Alive_FullMethodName    = "/hestia.jobfair.v1.identity.IdentityManagement/Alive"
	IdentityManagement_Logout_FullMethodName   = "/hestia.jobfair.v1.identity.IdentityManagement/Logout"
)

// IdentityManagementClient is the client API for IdentityManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityManagementClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Alive(ctx context.Context, in *common.Token, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Logout(ctx context.Context, in *common.Token, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type identityManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityManagementClient(cc grpc.ClientConnInterface) IdentityManagementClient {
	return &identityManagementClient{cc}
}

func (c *identityManagementClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, IdentityManagement_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IdentityManagement_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) Alive(ctx context.Context, in *common.Token, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IdentityManagement_Alive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityManagementClient) Logout(ctx context.Context, in *common.Token, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IdentityManagement_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityManagementServer is the server API for IdentityManagement service.
// All implementations must embed UnimplementedIdentityManagementServer
// for forward compatibility.
type IdentityManagementServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*emptypb.Empty, error)
	Alive(context.Context, *common.Token) (*emptypb.Empty, error)
	Logout(context.Context, *common.Token) (*emptypb.Empty, error)
	mustEmbedUnimplementedIdentityManagementServer()
}

// UnimplementedIdentityManagementServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIdentityManagementServer struct{}

func (UnimplementedIdentityManagementServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedIdentityManagementServer) Register(context.Context, *RegisterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedIdentityManagementServer) Alive(context.Context, *common.Token) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alive not implemented")
}
func (UnimplementedIdentityManagementServer) Logout(context.Context, *common.Token) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedIdentityManagementServer) mustEmbedUnimplementedIdentityManagementServer() {}
func (UnimplementedIdentityManagementServer) testEmbeddedByValue()                            {}

// UnsafeIdentityManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityManagementServer will
// result in compilation errors.
type UnsafeIdentityManagementServer interface {
	mustEmbedUnimplementedIdentityManagementServer()
}

func RegisterIdentityManagementServer(s grpc.ServiceRegistrar, srv IdentityManagementServer) {
	// If the following call pancis, it indicates UnimplementedIdentityManagementServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&IdentityManagement_ServiceDesc, srv)
}

func _IdentityManagement_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityManagement_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityManagement_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_Alive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).Alive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityManagement_Alive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).Alive(ctx, req.(*common.Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityManagement_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityManagementServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityManagement_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityManagementServer).Logout(ctx, req.(*common.Token))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityManagement_ServiceDesc is the grpc.ServiceDesc for IdentityManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hestia.jobfair.v1.identity.IdentityManagement",
	HandlerType: (*IdentityManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _IdentityManagement_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _IdentityManagement_Register_Handler,
		},
		{
			MethodName: "Alive",
			Handler:    _IdentityManagement_Alive_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _IdentityManagement_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity_management.proto",
}