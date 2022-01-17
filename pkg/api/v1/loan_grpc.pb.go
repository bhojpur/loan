// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// LoanServiceClient is the client API for LoanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoanServiceClient interface {
	// StartLocalEngine starts a Loan Engine on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the swarm/config.yaml
	//   3. all bytes constituting the Engine YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalEngine(ctx context.Context, opts ...grpc.CallOption) (LoanService_StartLocalEngineClient, error)
	// StartFromPreviousEngine starts a new Engine based on a previous one.
	// If the previous Engine does not have the can-replay condition set this call will result in an error.
	StartFromPreviousEngine(ctx context.Context, in *StartFromPreviousEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error)
	// StartEngineRequest starts a new Engine based on its specification.
	StartEngine(ctx context.Context, in *StartEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error)
	// Searches for Engine(s) known to this Engine
	ListEngines(ctx context.Context, in *ListEnginesRequest, opts ...grpc.CallOption) (*ListEnginesResponse, error)
	// Subscribe listens to new Engine(s) updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (LoanService_SubscribeClient, error)
	// GetEngine retrieves details of a single Engine
	GetEngine(ctx context.Context, in *GetEngineRequest, opts ...grpc.CallOption) (*GetEngineResponse, error)
	// Listen listens to Engine updates and log output of a running Engine
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (LoanService_ListenClient, error)
	// StopEngine stops a currently running Engine
	StopEngine(ctx context.Context, in *StopEngineRequest, opts ...grpc.CallOption) (*StopEngineResponse, error)
}

type loanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoanServiceClient(cc grpc.ClientConnInterface) LoanServiceClient {
	return &loanServiceClient{cc}
}

func (c *loanServiceClient) StartLocalEngine(ctx context.Context, opts ...grpc.CallOption) (LoanService_StartLocalEngineClient, error) {
	stream, err := c.cc.NewStream(ctx, &LoanService_ServiceDesc.Streams[0], "/v1.LoanService/StartLocalEngine", opts...)
	if err != nil {
		return nil, err
	}
	x := &loanServiceStartLocalEngineClient{stream}
	return x, nil
}

type LoanService_StartLocalEngineClient interface {
	Send(*StartLocalEngineRequest) error
	CloseAndRecv() (*StartEngineResponse, error)
	grpc.ClientStream
}

type loanServiceStartLocalEngineClient struct {
	grpc.ClientStream
}

func (x *loanServiceStartLocalEngineClient) Send(m *StartLocalEngineRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loanServiceStartLocalEngineClient) CloseAndRecv() (*StartEngineResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartEngineResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *loanServiceClient) StartFromPreviousEngine(ctx context.Context, in *StartFromPreviousEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error) {
	out := new(StartEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.LoanService/StartFromPreviousEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loanServiceClient) StartEngine(ctx context.Context, in *StartEngineRequest, opts ...grpc.CallOption) (*StartEngineResponse, error) {
	out := new(StartEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.LoanService/StartEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loanServiceClient) ListEngines(ctx context.Context, in *ListEnginesRequest, opts ...grpc.CallOption) (*ListEnginesResponse, error) {
	out := new(ListEnginesResponse)
	err := c.cc.Invoke(ctx, "/v1.LoanService/ListEngines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loanServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (LoanService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &LoanService_ServiceDesc.Streams[1], "/v1.LoanService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &loanServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LoanService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type loanServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *loanServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *loanServiceClient) GetEngine(ctx context.Context, in *GetEngineRequest, opts ...grpc.CallOption) (*GetEngineResponse, error) {
	out := new(GetEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.LoanService/GetEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loanServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (LoanService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &LoanService_ServiceDesc.Streams[2], "/v1.LoanService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &loanServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LoanService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type loanServiceListenClient struct {
	grpc.ClientStream
}

func (x *loanServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *loanServiceClient) StopEngine(ctx context.Context, in *StopEngineRequest, opts ...grpc.CallOption) (*StopEngineResponse, error) {
	out := new(StopEngineResponse)
	err := c.cc.Invoke(ctx, "/v1.LoanService/StopEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoanServiceServer is the server API for LoanService service.
// All implementations must embed UnimplementedLoanServiceServer
// for forward compatibility
type LoanServiceServer interface {
	// StartLocalEngine starts a Loan Engine on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the swarm/config.yaml
	//   3. all bytes constituting the Engine YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalEngine(LoanService_StartLocalEngineServer) error
	// StartFromPreviousEngine starts a new Engine based on a previous one.
	// If the previous Engine does not have the can-replay condition set this call will result in an error.
	StartFromPreviousEngine(context.Context, *StartFromPreviousEngineRequest) (*StartEngineResponse, error)
	// StartEngineRequest starts a new Engine based on its specification.
	StartEngine(context.Context, *StartEngineRequest) (*StartEngineResponse, error)
	// Searches for Engine(s) known to this Engine
	ListEngines(context.Context, *ListEnginesRequest) (*ListEnginesResponse, error)
	// Subscribe listens to new Engine(s) updates
	Subscribe(*SubscribeRequest, LoanService_SubscribeServer) error
	// GetEngine retrieves details of a single Engine
	GetEngine(context.Context, *GetEngineRequest) (*GetEngineResponse, error)
	// Listen listens to Engine updates and log output of a running Engine
	Listen(*ListenRequest, LoanService_ListenServer) error
	// StopEngine stops a currently running Engine
	StopEngine(context.Context, *StopEngineRequest) (*StopEngineResponse, error)
	mustEmbedUnimplementedLoanServiceServer()
}

// UnimplementedLoanServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoanServiceServer struct {
}

func (UnimplementedLoanServiceServer) StartLocalEngine(LoanService_StartLocalEngineServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalEngine not implemented")
}
func (UnimplementedLoanServiceServer) StartFromPreviousEngine(context.Context, *StartFromPreviousEngineRequest) (*StartEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousEngine not implemented")
}
func (UnimplementedLoanServiceServer) StartEngine(context.Context, *StartEngineRequest) (*StartEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartEngine not implemented")
}
func (UnimplementedLoanServiceServer) ListEngines(context.Context, *ListEnginesRequest) (*ListEnginesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEngines not implemented")
}
func (UnimplementedLoanServiceServer) Subscribe(*SubscribeRequest, LoanService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedLoanServiceServer) GetEngine(context.Context, *GetEngineRequest) (*GetEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEngine not implemented")
}
func (UnimplementedLoanServiceServer) Listen(*ListenRequest, LoanService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedLoanServiceServer) StopEngine(context.Context, *StopEngineRequest) (*StopEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopEngine not implemented")
}
func (UnimplementedLoanServiceServer) mustEmbedUnimplementedLoanServiceServer() {}

// UnsafeLoanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoanServiceServer will
// result in compilation errors.
type UnsafeLoanServiceServer interface {
	mustEmbedUnimplementedLoanServiceServer()
}

func RegisterLoanServiceServer(s grpc.ServiceRegistrar, srv LoanServiceServer) {
	s.RegisterService(&LoanService_ServiceDesc, srv)
}

func _LoanService_StartLocalEngine_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoanServiceServer).StartLocalEngine(&loanServiceStartLocalEngineServer{stream})
}

type LoanService_StartLocalEngineServer interface {
	SendAndClose(*StartEngineResponse) error
	Recv() (*StartLocalEngineRequest, error)
	grpc.ServerStream
}

type loanServiceStartLocalEngineServer struct {
	grpc.ServerStream
}

func (x *loanServiceStartLocalEngineServer) SendAndClose(m *StartEngineResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loanServiceStartLocalEngineServer) Recv() (*StartLocalEngineRequest, error) {
	m := new(StartLocalEngineRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LoanService_StartFromPreviousEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoanServiceServer).StartFromPreviousEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.LoanService/StartFromPreviousEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoanServiceServer).StartFromPreviousEngine(ctx, req.(*StartFromPreviousEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoanService_StartEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoanServiceServer).StartEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.LoanService/StartEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoanServiceServer).StartEngine(ctx, req.(*StartEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoanService_ListEngines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnginesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoanServiceServer).ListEngines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.LoanService/ListEngines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoanServiceServer).ListEngines(ctx, req.(*ListEnginesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoanService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LoanServiceServer).Subscribe(m, &loanServiceSubscribeServer{stream})
}

type LoanService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type loanServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *loanServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LoanService_GetEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoanServiceServer).GetEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.LoanService/GetEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoanServiceServer).GetEngine(ctx, req.(*GetEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoanService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LoanServiceServer).Listen(m, &loanServiceListenServer{stream})
}

type LoanService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type loanServiceListenServer struct {
	grpc.ServerStream
}

func (x *loanServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LoanService_StopEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoanServiceServer).StopEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.LoanService/StopEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoanServiceServer).StopEngine(ctx, req.(*StopEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoanService_ServiceDesc is the grpc.ServiceDesc for LoanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.LoanService",
	HandlerType: (*LoanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousEngine",
			Handler:    _LoanService_StartFromPreviousEngine_Handler,
		},
		{
			MethodName: "StartEngine",
			Handler:    _LoanService_StartEngine_Handler,
		},
		{
			MethodName: "ListEngines",
			Handler:    _LoanService_ListEngines_Handler,
		},
		{
			MethodName: "GetEngine",
			Handler:    _LoanService_GetEngine_Handler,
		},
		{
			MethodName: "StopEngine",
			Handler:    _LoanService_StopEngine_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalEngine",
			Handler:       _LoanService_StartLocalEngine_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _LoanService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _LoanService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "loan.proto",
}
