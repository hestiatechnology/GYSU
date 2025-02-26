// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: hestia/jobfair/events/v1/events.proto

package events

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
	EventsService_CreateEvent_FullMethodName        = "/hestia.jobfair.events.v1.EventsService/CreateEvent"
	EventsService_GetEvents_FullMethodName          = "/hestia.jobfair.events.v1.EventsService/GetEvents"
	EventsService_GetEvent_FullMethodName           = "/hestia.jobfair.events.v1.EventsService/GetEvent"
	EventsService_UpdateEvent_FullMethodName        = "/hestia.jobfair.events.v1.EventsService/UpdateEvent"
	EventsService_DeleteEvent_FullMethodName        = "/hestia.jobfair.events.v1.EventsService/DeleteEvent"
	EventsService_SubscribeToEvent_FullMethodName   = "/hestia.jobfair.events.v1.EventsService/SubscribeToEvent"
	EventsService_UnsubscribeToEvent_FullMethodName = "/hestia.jobfair.events.v1.EventsService/UnsubscribeToEvent"
)

// EventsServiceClient is the client API for EventsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventsServiceClient interface {
	CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*common.Id, error)
	GetEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EventList, error)
	GetEvent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*Event, error)
	UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteEvent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SubscribeToEvent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UnsubscribeToEvent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type eventsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsServiceClient(cc grpc.ClientConnInterface) EventsServiceClient {
	return &eventsServiceClient{cc}
}

func (c *eventsServiceClient) CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*common.Id, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Id)
	err := c.cc.Invoke(ctx, EventsService_CreateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) GetEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EventList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventList)
	err := c.cc.Invoke(ctx, EventsService_GetEvents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) GetEvent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*Event, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Event)
	err := c.cc.Invoke(ctx, EventsService_GetEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EventsService_UpdateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) DeleteEvent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EventsService_DeleteEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) SubscribeToEvent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EventsService_SubscribeToEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) UnsubscribeToEvent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EventsService_UnsubscribeToEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServiceServer is the server API for EventsService service.
// All implementations must embed UnimplementedEventsServiceServer
// for forward compatibility.
type EventsServiceServer interface {
	CreateEvent(context.Context, *Event) (*common.Id, error)
	GetEvents(context.Context, *emptypb.Empty) (*EventList, error)
	GetEvent(context.Context, *common.Id) (*Event, error)
	UpdateEvent(context.Context, *Event) (*emptypb.Empty, error)
	DeleteEvent(context.Context, *common.Id) (*emptypb.Empty, error)
	SubscribeToEvent(context.Context, *common.Id) (*emptypb.Empty, error)
	UnsubscribeToEvent(context.Context, *common.Id) (*emptypb.Empty, error)
	mustEmbedUnimplementedEventsServiceServer()
}

// UnimplementedEventsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEventsServiceServer struct{}

func (UnimplementedEventsServiceServer) CreateEvent(context.Context, *Event) (*common.Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedEventsServiceServer) GetEvents(context.Context, *emptypb.Empty) (*EventList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedEventsServiceServer) GetEvent(context.Context, *common.Id) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedEventsServiceServer) UpdateEvent(context.Context, *Event) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedEventsServiceServer) DeleteEvent(context.Context, *common.Id) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedEventsServiceServer) SubscribeToEvent(context.Context, *common.Id) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeToEvent not implemented")
}
func (UnimplementedEventsServiceServer) UnsubscribeToEvent(context.Context, *common.Id) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubscribeToEvent not implemented")
}
func (UnimplementedEventsServiceServer) mustEmbedUnimplementedEventsServiceServer() {}
func (UnimplementedEventsServiceServer) testEmbeddedByValue()                       {}

// UnsafeEventsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventsServiceServer will
// result in compilation errors.
type UnsafeEventsServiceServer interface {
	mustEmbedUnimplementedEventsServiceServer()
}

func RegisterEventsServiceServer(s grpc.ServiceRegistrar, srv EventsServiceServer) {
	// If the following call pancis, it indicates UnimplementedEventsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EventsService_ServiceDesc, srv)
}

func _EventsService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).CreateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_GetEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).GetEvents(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_GetEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).GetEvent(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_UpdateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).UpdateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_DeleteEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).DeleteEvent(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_SubscribeToEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).SubscribeToEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_SubscribeToEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).SubscribeToEvent(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_UnsubscribeToEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).UnsubscribeToEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_UnsubscribeToEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).UnsubscribeToEvent(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

// EventsService_ServiceDesc is the grpc.ServiceDesc for EventsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hestia.jobfair.events.v1.EventsService",
	HandlerType: (*EventsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _EventsService_CreateEvent_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _EventsService_GetEvents_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _EventsService_GetEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _EventsService_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _EventsService_DeleteEvent_Handler,
		},
		{
			MethodName: "SubscribeToEvent",
			Handler:    _EventsService_SubscribeToEvent_Handler,
		},
		{
			MethodName: "UnsubscribeToEvent",
			Handler:    _EventsService_UnsubscribeToEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hestia/jobfair/events/v1/events.proto",
}
