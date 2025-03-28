// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: api/proto/sentry.proto

package proto

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
	SentryService_StartJob_FullMethodName      = "/sentry.SentryService/StartJob"
	SentryService_KillJob_FullMethodName       = "/sentry.SentryService/KillJob"
	SentryService_GetJobStatus_FullMethodName  = "/sentry.SentryService/GetJobStatus"
	SentryService_StreamJobLogs_FullMethodName = "/sentry.SentryService/StreamJobLogs"
	SentryService_ListJobs_FullMethodName      = "/sentry.SentryService/ListJobs"
)

// SentryServiceClient is the client API for SentryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SentryServiceClient interface {
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error)
	KillJob(ctx context.Context, in *KillJobRequest, opts ...grpc.CallOption) (*KillJobResponse, error)
	GetJobStatus(ctx context.Context, in *JobStatusRequest, opts ...grpc.CallOption) (*JobStatusResponse, error)
	StreamJobLogs(ctx context.Context, in *JobLogsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[JobOutput], error)
	ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error)
}

type sentryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSentryServiceClient(cc grpc.ClientConnInterface) SentryServiceClient {
	return &sentryServiceClient{cc}
}

func (c *sentryServiceClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartJobResponse)
	err := c.cc.Invoke(ctx, SentryService_StartJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryServiceClient) KillJob(ctx context.Context, in *KillJobRequest, opts ...grpc.CallOption) (*KillJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KillJobResponse)
	err := c.cc.Invoke(ctx, SentryService_KillJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryServiceClient) GetJobStatus(ctx context.Context, in *JobStatusRequest, opts ...grpc.CallOption) (*JobStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobStatusResponse)
	err := c.cc.Invoke(ctx, SentryService_GetJobStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryServiceClient) StreamJobLogs(ctx context.Context, in *JobLogsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[JobOutput], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SentryService_ServiceDesc.Streams[0], SentryService_StreamJobLogs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[JobLogsRequest, JobOutput]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SentryService_StreamJobLogsClient = grpc.ServerStreamingClient[JobOutput]

func (c *sentryServiceClient) ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListJobsResponse)
	err := c.cc.Invoke(ctx, SentryService_ListJobs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SentryServiceServer is the server API for SentryService service.
// All implementations must embed UnimplementedSentryServiceServer
// for forward compatibility.
type SentryServiceServer interface {
	StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error)
	KillJob(context.Context, *KillJobRequest) (*KillJobResponse, error)
	GetJobStatus(context.Context, *JobStatusRequest) (*JobStatusResponse, error)
	StreamJobLogs(*JobLogsRequest, grpc.ServerStreamingServer[JobOutput]) error
	ListJobs(context.Context, *ListJobsRequest) (*ListJobsResponse, error)
	mustEmbedUnimplementedSentryServiceServer()
}

// UnimplementedSentryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSentryServiceServer struct{}

func (UnimplementedSentryServiceServer) StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartJob not implemented")
}
func (UnimplementedSentryServiceServer) KillJob(context.Context, *KillJobRequest) (*KillJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillJob not implemented")
}
func (UnimplementedSentryServiceServer) GetJobStatus(context.Context, *JobStatusRequest) (*JobStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobStatus not implemented")
}
func (UnimplementedSentryServiceServer) StreamJobLogs(*JobLogsRequest, grpc.ServerStreamingServer[JobOutput]) error {
	return status.Errorf(codes.Unimplemented, "method StreamJobLogs not implemented")
}
func (UnimplementedSentryServiceServer) ListJobs(context.Context, *ListJobsRequest) (*ListJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobs not implemented")
}
func (UnimplementedSentryServiceServer) mustEmbedUnimplementedSentryServiceServer() {}
func (UnimplementedSentryServiceServer) testEmbeddedByValue()                       {}

// UnsafeSentryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SentryServiceServer will
// result in compilation errors.
type UnsafeSentryServiceServer interface {
	mustEmbedUnimplementedSentryServiceServer()
}

func RegisterSentryServiceServer(s grpc.ServiceRegistrar, srv SentryServiceServer) {
	// If the following call pancis, it indicates UnimplementedSentryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SentryService_ServiceDesc, srv)
}

func _SentryService_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServiceServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SentryService_StartJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServiceServer).StartJob(ctx, req.(*StartJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SentryService_KillJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServiceServer).KillJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SentryService_KillJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServiceServer).KillJob(ctx, req.(*KillJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SentryService_GetJobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServiceServer).GetJobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SentryService_GetJobStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServiceServer).GetJobStatus(ctx, req.(*JobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SentryService_StreamJobLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JobLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SentryServiceServer).StreamJobLogs(m, &grpc.GenericServerStream[JobLogsRequest, JobOutput]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SentryService_StreamJobLogsServer = grpc.ServerStreamingServer[JobOutput]

func _SentryService_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServiceServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SentryService_ListJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServiceServer).ListJobs(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SentryService_ServiceDesc is the grpc.ServiceDesc for SentryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SentryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sentry.SentryService",
	HandlerType: (*SentryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _SentryService_StartJob_Handler,
		},
		{
			MethodName: "KillJob",
			Handler:    _SentryService_KillJob_Handler,
		},
		{
			MethodName: "GetJobStatus",
			Handler:    _SentryService_GetJobStatus_Handler,
		},
		{
			MethodName: "ListJobs",
			Handler:    _SentryService_ListJobs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamJobLogs",
			Handler:       _SentryService_StreamJobLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/proto/sentry.proto",
}
