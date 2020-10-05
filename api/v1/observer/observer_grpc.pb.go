// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package observer

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

// ObserverClient is the client API for Observer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObserverClient interface {
	// GetFlows returning structured data, meant to eventually obsolete GetLastNFlows.
	GetFlows(ctx context.Context, in *GetFlowsRequest, opts ...grpc.CallOption) (Observer_GetFlowsClient, error)
	// GetNodes returns information about nodes in a cluster.
	GetNodes(ctx context.Context, in *GetNodesRequest, opts ...grpc.CallOption) (*GetNodesResponse, error)
	// ServerStatus returns some details about the running hubble server.
	ServerStatus(ctx context.Context, in *ServerStatusRequest, opts ...grpc.CallOption) (*ServerStatusResponse, error)
}

type observerClient struct {
	cc grpc.ClientConnInterface
}

func NewObserverClient(cc grpc.ClientConnInterface) ObserverClient {
	return &observerClient{cc}
}

func (c *observerClient) GetFlows(ctx context.Context, in *GetFlowsRequest, opts ...grpc.CallOption) (Observer_GetFlowsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Observer_ServiceDesc.Streams[0], "/observer.Observer/GetFlows", opts...)
	if err != nil {
		return nil, err
	}
	x := &observerGetFlowsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Observer_GetFlowsClient interface {
	Recv() (*GetFlowsResponse, error)
	grpc.ClientStream
}

type observerGetFlowsClient struct {
	grpc.ClientStream
}

func (x *observerGetFlowsClient) Recv() (*GetFlowsResponse, error) {
	m := new(GetFlowsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *observerClient) GetNodes(ctx context.Context, in *GetNodesRequest, opts ...grpc.CallOption) (*GetNodesResponse, error) {
	out := new(GetNodesResponse)
	err := c.cc.Invoke(ctx, "/observer.Observer/GetNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *observerClient) ServerStatus(ctx context.Context, in *ServerStatusRequest, opts ...grpc.CallOption) (*ServerStatusResponse, error) {
	out := new(ServerStatusResponse)
	err := c.cc.Invoke(ctx, "/observer.Observer/ServerStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObserverServer is the server API for Observer service.
// All implementations should embed UnimplementedObserverServer
// for forward compatibility
type ObserverServer interface {
	// GetFlows returning structured data, meant to eventually obsolete GetLastNFlows.
	GetFlows(*GetFlowsRequest, Observer_GetFlowsServer) error
	// GetNodes returns information about nodes in a cluster.
	GetNodes(context.Context, *GetNodesRequest) (*GetNodesResponse, error)
	// ServerStatus returns some details about the running hubble server.
	ServerStatus(context.Context, *ServerStatusRequest) (*ServerStatusResponse, error)
}

// UnimplementedObserverServer should be embedded to have forward compatible implementations.
type UnimplementedObserverServer struct {
}

func (UnimplementedObserverServer) GetFlows(*GetFlowsRequest, Observer_GetFlowsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFlows not implemented")
}
func (UnimplementedObserverServer) GetNodes(context.Context, *GetNodesRequest) (*GetNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodes not implemented")
}
func (UnimplementedObserverServer) ServerStatus(context.Context, *ServerStatusRequest) (*ServerStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerStatus not implemented")
}

// UnsafeObserverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObserverServer will
// result in compilation errors.
type UnsafeObserverServer interface {
	mustEmbedUnimplementedObserverServer()
}

func RegisterObserverServer(s grpc.ServiceRegistrar, srv ObserverServer) {
	s.RegisterService(&Observer_ServiceDesc, srv)
}

func _Observer_GetFlows_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFlowsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ObserverServer).GetFlows(m, &observerGetFlowsServer{stream})
}

type Observer_GetFlowsServer interface {
	Send(*GetFlowsResponse) error
	grpc.ServerStream
}

type observerGetFlowsServer struct {
	grpc.ServerStream
}

func (x *observerGetFlowsServer) Send(m *GetFlowsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Observer_GetNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObserverServer).GetNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/observer.Observer/GetNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObserverServer).GetNodes(ctx, req.(*GetNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Observer_ServerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObserverServer).ServerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/observer.Observer/ServerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObserverServer).ServerStatus(ctx, req.(*ServerStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Observer_ServiceDesc is the grpc.ServiceDesc for Observer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Observer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "observer.Observer",
	HandlerType: (*ObserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodes",
			Handler:    _Observer_GetNodes_Handler,
		},
		{
			MethodName: "ServerStatus",
			Handler:    _Observer_ServerStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFlows",
			Handler:       _Observer_GetFlows_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "observer/observer.proto",
}
