// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_apiservice_grpc

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

// RouteCourseClient is the client API for RouteCourse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteCourseClient interface {
	AddCourse(ctx context.Context, in *NewCourse, opts ...grpc.CallOption) (*Course, error)
	DeleteCourse(ctx context.Context, in *NewCourse, opts ...grpc.CallOption) (*Course, error)
	GetCourseById(ctx context.Context, in *GetCourseByIdRequest, opts ...grpc.CallOption) (*Course, error)
	GetOverview(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RouteCourse_GetOverviewClient, error)
	UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*Course, error)
}

type routeCourseClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteCourseClient(cc grpc.ClientConnInterface) RouteCourseClient {
	return &routeCourseClient{cc}
}

func (c *routeCourseClient) AddCourse(ctx context.Context, in *NewCourse, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/apiservice.RouteCourse/AddCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeCourseClient) DeleteCourse(ctx context.Context, in *NewCourse, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/apiservice.RouteCourse/DeleteCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeCourseClient) GetCourseById(ctx context.Context, in *GetCourseByIdRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/apiservice.RouteCourse/GetCourseById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeCourseClient) GetOverview(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RouteCourse_GetOverviewClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteCourse_ServiceDesc.Streams[0], "/apiservice.RouteCourse/GetOverview", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeCourseGetOverviewClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteCourse_GetOverviewClient interface {
	Recv() (*Course, error)
	grpc.ClientStream
}

type routeCourseGetOverviewClient struct {
	grpc.ClientStream
}

func (x *routeCourseGetOverviewClient) Recv() (*Course, error) {
	m := new(Course)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeCourseClient) UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/apiservice.RouteCourse/UpdateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteCourseServer is the server API for RouteCourse service.
// All implementations must embed UnimplementedRouteCourseServer
// for forward compatibility
type RouteCourseServer interface {
	AddCourse(context.Context, *NewCourse) (*Course, error)
	DeleteCourse(context.Context, *NewCourse) (*Course, error)
	GetCourseById(context.Context, *GetCourseByIdRequest) (*Course, error)
	GetOverview(*emptypb.Empty, RouteCourse_GetOverviewServer) error
	UpdateCourse(context.Context, *UpdateCourseRequest) (*Course, error)
	mustEmbedUnimplementedRouteCourseServer()
}

// UnimplementedRouteCourseServer must be embedded to have forward compatible implementations.
type UnimplementedRouteCourseServer struct {
}

func (UnimplementedRouteCourseServer) AddCourse(context.Context, *NewCourse) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCourse not implemented")
}
func (UnimplementedRouteCourseServer) DeleteCourse(context.Context, *NewCourse) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}
func (UnimplementedRouteCourseServer) GetCourseById(context.Context, *GetCourseByIdRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourseById not implemented")
}
func (UnimplementedRouteCourseServer) GetOverview(*emptypb.Empty, RouteCourse_GetOverviewServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOverview not implemented")
}
func (UnimplementedRouteCourseServer) UpdateCourse(context.Context, *UpdateCourseRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourse not implemented")
}
func (UnimplementedRouteCourseServer) mustEmbedUnimplementedRouteCourseServer() {}

// UnsafeRouteCourseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteCourseServer will
// result in compilation errors.
type UnsafeRouteCourseServer interface {
	mustEmbedUnimplementedRouteCourseServer()
}

func RegisterRouteCourseServer(s grpc.ServiceRegistrar, srv RouteCourseServer) {
	s.RegisterService(&RouteCourse_ServiceDesc, srv)
}

func _RouteCourse_AddCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCourse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteCourseServer).AddCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiservice.RouteCourse/AddCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteCourseServer).AddCourse(ctx, req.(*NewCourse))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteCourse_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCourse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteCourseServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiservice.RouteCourse/DeleteCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteCourseServer).DeleteCourse(ctx, req.(*NewCourse))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteCourse_GetCourseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteCourseServer).GetCourseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiservice.RouteCourse/GetCourseById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteCourseServer).GetCourseById(ctx, req.(*GetCourseByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteCourse_GetOverview_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteCourseServer).GetOverview(m, &routeCourseGetOverviewServer{stream})
}

type RouteCourse_GetOverviewServer interface {
	Send(*Course) error
	grpc.ServerStream
}

type routeCourseGetOverviewServer struct {
	grpc.ServerStream
}

func (x *routeCourseGetOverviewServer) Send(m *Course) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteCourse_UpdateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteCourseServer).UpdateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiservice.RouteCourse/UpdateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteCourseServer).UpdateCourse(ctx, req.(*UpdateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RouteCourse_ServiceDesc is the grpc.ServiceDesc for RouteCourse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteCourse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiservice.RouteCourse",
	HandlerType: (*RouteCourseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCourse",
			Handler:    _RouteCourse_AddCourse_Handler,
		},
		{
			MethodName: "DeleteCourse",
			Handler:    _RouteCourse_DeleteCourse_Handler,
		},
		{
			MethodName: "GetCourseById",
			Handler:    _RouteCourse_GetCourseById_Handler,
		},
		{
			MethodName: "UpdateCourse",
			Handler:    _RouteCourse_UpdateCourse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetOverview",
			Handler:       _RouteCourse_GetOverview_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "apiservice/apiservice.proto",
}
