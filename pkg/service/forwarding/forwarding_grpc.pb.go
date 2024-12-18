// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: service/forwarding/forwarding.proto

package forwarding

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Forwarding_Create_FullMethodName    = "/forwarding.Forwarding/Create"
	Forwarding_List_FullMethodName      = "/forwarding.Forwarding/List"
	Forwarding_Pause_FullMethodName     = "/forwarding.Forwarding/Pause"
	Forwarding_Resume_FullMethodName    = "/forwarding.Forwarding/Resume"
	Forwarding_Terminate_FullMethodName = "/forwarding.Forwarding/Terminate"
)

// ForwardingClient is the client API for Forwarding service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Forwarding manages the lifecycle of forwarding sessions.
type ForwardingClient interface {
	// Create creates a new session.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// List returns metadata for existing sessions.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Pause pauses sessions.
	Pause(ctx context.Context, in *PauseRequest, opts ...grpc.CallOption) (*PauseResponse, error)
	// Resume resumes paused or disconnected sessions.
	Resume(ctx context.Context, in *ResumeRequest, opts ...grpc.CallOption) (*ResumeResponse, error)
	// Terminate terminates sessions.
	Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error)
}

type forwardingClient struct {
	cc grpc.ClientConnInterface
}

func NewForwardingClient(cc grpc.ClientConnInterface) ForwardingClient {
	return &forwardingClient{cc}
}

func (c *forwardingClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Forwarding_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forwardingClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, Forwarding_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forwardingClient) Pause(ctx context.Context, in *PauseRequest, opts ...grpc.CallOption) (*PauseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PauseResponse)
	err := c.cc.Invoke(ctx, Forwarding_Pause_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forwardingClient) Resume(ctx context.Context, in *ResumeRequest, opts ...grpc.CallOption) (*ResumeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResumeResponse)
	err := c.cc.Invoke(ctx, Forwarding_Resume_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forwardingClient) Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TerminateResponse)
	err := c.cc.Invoke(ctx, Forwarding_Terminate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForwardingServer is the server API for Forwarding service.
// All implementations must embed UnimplementedForwardingServer
// for forward compatibility.
//
// Forwarding manages the lifecycle of forwarding sessions.
type ForwardingServer interface {
	// Create creates a new session.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// List returns metadata for existing sessions.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Pause pauses sessions.
	Pause(context.Context, *PauseRequest) (*PauseResponse, error)
	// Resume resumes paused or disconnected sessions.
	Resume(context.Context, *ResumeRequest) (*ResumeResponse, error)
	// Terminate terminates sessions.
	Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error)
	mustEmbedUnimplementedForwardingServer()
}

// UnimplementedForwardingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedForwardingServer struct{}

func (UnimplementedForwardingServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedForwardingServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedForwardingServer) Pause(context.Context, *PauseRequest) (*PauseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedForwardingServer) Resume(context.Context, *ResumeRequest) (*ResumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (UnimplementedForwardingServer) Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Terminate not implemented")
}
func (UnimplementedForwardingServer) mustEmbedUnimplementedForwardingServer() {}
func (UnimplementedForwardingServer) testEmbeddedByValue()                    {}

// UnsafeForwardingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ForwardingServer will
// result in compilation errors.
type UnsafeForwardingServer interface {
	mustEmbedUnimplementedForwardingServer()
}

func RegisterForwardingServer(s grpc.ServiceRegistrar, srv ForwardingServer) {
	// If the following call pancis, it indicates UnimplementedForwardingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Forwarding_ServiceDesc, srv)
}

func _Forwarding_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwardingServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forwarding_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwardingServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forwarding_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwardingServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forwarding_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwardingServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forwarding_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwardingServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forwarding_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwardingServer).Pause(ctx, req.(*PauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forwarding_Resume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwardingServer).Resume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forwarding_Resume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwardingServer).Resume(ctx, req.(*ResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forwarding_Terminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwardingServer).Terminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Forwarding_Terminate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwardingServer).Terminate(ctx, req.(*TerminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Forwarding_ServiceDesc is the grpc.ServiceDesc for Forwarding service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Forwarding_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "forwarding.Forwarding",
	HandlerType: (*ForwardingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Forwarding_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Forwarding_List_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _Forwarding_Pause_Handler,
		},
		{
			MethodName: "Resume",
			Handler:    _Forwarding_Resume_Handler,
		},
		{
			MethodName: "Terminate",
			Handler:    _Forwarding_Terminate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/forwarding/forwarding.proto",
}
