// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: sl_bfd_ipv6.proto

package service_layer

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

// SLBfdv6OperClient is the client API for SLBfdv6Oper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SLBfdv6OperClient interface {
	// SLBfdRegMsg.Oper = SL_REGOP_REGISTER:
	//
	//	Global BFD registration.
	//	A client Must Register BEFORE BFD sessions can be added/modified.
	//
	// SLBfdRegMsg.Oper = SL_REGOP_UNREGISTER:
	//
	//	Global BFD un-registration.
	//	This call is used to end all BFD notifications and unregister any
	//	interest in BFD session configuration.
	//	This call cleans up all BFD sessions previously requested.
	//
	// SLBfdRegMsg.Oper = SL_REGOP_EOF:
	//
	//	BFD End Of File.
	//	After Registration, the client is expected to send an EOF
	//	message to convey the end of replay of the client's known objects.
	//	This is especially useful under certain restart scenarios when the
	//	client and the server are trying to synchronize their BFD sessions.
	SLBfdv6RegOp(ctx context.Context, in *SLBfdRegMsg, opts ...grpc.CallOption) (*SLBfdRegMsgRsp, error)
	// Used to retrieve global BFD info from the server.
	SLBfdv6Get(ctx context.Context, in *SLBfdGetMsg, opts ...grpc.CallOption) (*SLBfdGetMsgRsp, error)
	// Used to retrieve global BFD stats from the server.
	SLBfdv6GetStats(ctx context.Context, in *SLBfdGetMsg, opts ...grpc.CallOption) (*SLBfdGetStatsMsgRsp, error)
	// This call is used to get a stream of session state notifications.
	// The caller must maintain the GRPC channel as long as
	// there is interest in BFD session notifications. Only sessions that were
	// created through this API will be notified to caller.
	// This call can be used to get "push" notifications for session states.
	// It is advised that the caller register for notifications before any
	// sessions are created to avoid any loss of notifications.
	SLBfdv6GetNotifStream(ctx context.Context, in *SLBfdGetNotifMsg, opts ...grpc.CallOption) (SLBfdv6Oper_SLBfdv6GetNotifStreamClient, error)
	// SLBfdv6Msg.Oper = SL_OBJOP_ADD:
	//
	//	Add one or multiple BFD sessions.
	//
	// SLBfdv6Msg.Oper = SL_OBJOP_UPDATE:
	//
	//	Update one or multiple BFD sessions.
	//
	// SLBfdv6Msg.Oper = SL_OBJOP_DELETE:
	//
	//	Delete one or multiple BFD sessions.
	SLBfdv6SessionOp(ctx context.Context, in *SLBfdv6Msg, opts ...grpc.CallOption) (*SLBfdv6MsgRsp, error)
	// Retrieve BFD session attributes and state.
	// This call can be used to "poll" the current state of a session.
	SLBfdv6SessionGet(ctx context.Context, in *SLBfdv6GetMsg, opts ...grpc.CallOption) (*SLBfdv6GetMsgRsp, error)
}

type sLBfdv6OperClient struct {
	cc grpc.ClientConnInterface
}

func NewSLBfdv6OperClient(cc grpc.ClientConnInterface) SLBfdv6OperClient {
	return &sLBfdv6OperClient{cc}
}

func (c *sLBfdv6OperClient) SLBfdv6RegOp(ctx context.Context, in *SLBfdRegMsg, opts ...grpc.CallOption) (*SLBfdRegMsgRsp, error) {
	out := new(SLBfdRegMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLBfdv6Oper/SLBfdv6RegOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLBfdv6OperClient) SLBfdv6Get(ctx context.Context, in *SLBfdGetMsg, opts ...grpc.CallOption) (*SLBfdGetMsgRsp, error) {
	out := new(SLBfdGetMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLBfdv6Oper/SLBfdv6Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLBfdv6OperClient) SLBfdv6GetStats(ctx context.Context, in *SLBfdGetMsg, opts ...grpc.CallOption) (*SLBfdGetStatsMsgRsp, error) {
	out := new(SLBfdGetStatsMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLBfdv6Oper/SLBfdv6GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLBfdv6OperClient) SLBfdv6GetNotifStream(ctx context.Context, in *SLBfdGetNotifMsg, opts ...grpc.CallOption) (SLBfdv6Oper_SLBfdv6GetNotifStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SLBfdv6Oper_ServiceDesc.Streams[0], "/service_layer.SLBfdv6Oper/SLBfdv6GetNotifStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sLBfdv6OperSLBfdv6GetNotifStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SLBfdv6Oper_SLBfdv6GetNotifStreamClient interface {
	Recv() (*SLBfdv6Notif, error)
	grpc.ClientStream
}

type sLBfdv6OperSLBfdv6GetNotifStreamClient struct {
	grpc.ClientStream
}

func (x *sLBfdv6OperSLBfdv6GetNotifStreamClient) Recv() (*SLBfdv6Notif, error) {
	m := new(SLBfdv6Notif)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sLBfdv6OperClient) SLBfdv6SessionOp(ctx context.Context, in *SLBfdv6Msg, opts ...grpc.CallOption) (*SLBfdv6MsgRsp, error) {
	out := new(SLBfdv6MsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLBfdv6Oper/SLBfdv6SessionOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLBfdv6OperClient) SLBfdv6SessionGet(ctx context.Context, in *SLBfdv6GetMsg, opts ...grpc.CallOption) (*SLBfdv6GetMsgRsp, error) {
	out := new(SLBfdv6GetMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLBfdv6Oper/SLBfdv6SessionGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SLBfdv6OperServer is the server API for SLBfdv6Oper service.
// All implementations must embed UnimplementedSLBfdv6OperServer
// for forward compatibility
type SLBfdv6OperServer interface {
	// SLBfdRegMsg.Oper = SL_REGOP_REGISTER:
	//
	//	Global BFD registration.
	//	A client Must Register BEFORE BFD sessions can be added/modified.
	//
	// SLBfdRegMsg.Oper = SL_REGOP_UNREGISTER:
	//
	//	Global BFD un-registration.
	//	This call is used to end all BFD notifications and unregister any
	//	interest in BFD session configuration.
	//	This call cleans up all BFD sessions previously requested.
	//
	// SLBfdRegMsg.Oper = SL_REGOP_EOF:
	//
	//	BFD End Of File.
	//	After Registration, the client is expected to send an EOF
	//	message to convey the end of replay of the client's known objects.
	//	This is especially useful under certain restart scenarios when the
	//	client and the server are trying to synchronize their BFD sessions.
	SLBfdv6RegOp(context.Context, *SLBfdRegMsg) (*SLBfdRegMsgRsp, error)
	// Used to retrieve global BFD info from the server.
	SLBfdv6Get(context.Context, *SLBfdGetMsg) (*SLBfdGetMsgRsp, error)
	// Used to retrieve global BFD stats from the server.
	SLBfdv6GetStats(context.Context, *SLBfdGetMsg) (*SLBfdGetStatsMsgRsp, error)
	// This call is used to get a stream of session state notifications.
	// The caller must maintain the GRPC channel as long as
	// there is interest in BFD session notifications. Only sessions that were
	// created through this API will be notified to caller.
	// This call can be used to get "push" notifications for session states.
	// It is advised that the caller register for notifications before any
	// sessions are created to avoid any loss of notifications.
	SLBfdv6GetNotifStream(*SLBfdGetNotifMsg, SLBfdv6Oper_SLBfdv6GetNotifStreamServer) error
	// SLBfdv6Msg.Oper = SL_OBJOP_ADD:
	//
	//	Add one or multiple BFD sessions.
	//
	// SLBfdv6Msg.Oper = SL_OBJOP_UPDATE:
	//
	//	Update one or multiple BFD sessions.
	//
	// SLBfdv6Msg.Oper = SL_OBJOP_DELETE:
	//
	//	Delete one or multiple BFD sessions.
	SLBfdv6SessionOp(context.Context, *SLBfdv6Msg) (*SLBfdv6MsgRsp, error)
	// Retrieve BFD session attributes and state.
	// This call can be used to "poll" the current state of a session.
	SLBfdv6SessionGet(context.Context, *SLBfdv6GetMsg) (*SLBfdv6GetMsgRsp, error)
	mustEmbedUnimplementedSLBfdv6OperServer()
}

// UnimplementedSLBfdv6OperServer must be embedded to have forward compatible implementations.
type UnimplementedSLBfdv6OperServer struct {
}

func (UnimplementedSLBfdv6OperServer) SLBfdv6RegOp(context.Context, *SLBfdRegMsg) (*SLBfdRegMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLBfdv6RegOp not implemented")
}
func (UnimplementedSLBfdv6OperServer) SLBfdv6Get(context.Context, *SLBfdGetMsg) (*SLBfdGetMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLBfdv6Get not implemented")
}
func (UnimplementedSLBfdv6OperServer) SLBfdv6GetStats(context.Context, *SLBfdGetMsg) (*SLBfdGetStatsMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLBfdv6GetStats not implemented")
}
func (UnimplementedSLBfdv6OperServer) SLBfdv6GetNotifStream(*SLBfdGetNotifMsg, SLBfdv6Oper_SLBfdv6GetNotifStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SLBfdv6GetNotifStream not implemented")
}
func (UnimplementedSLBfdv6OperServer) SLBfdv6SessionOp(context.Context, *SLBfdv6Msg) (*SLBfdv6MsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLBfdv6SessionOp not implemented")
}
func (UnimplementedSLBfdv6OperServer) SLBfdv6SessionGet(context.Context, *SLBfdv6GetMsg) (*SLBfdv6GetMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLBfdv6SessionGet not implemented")
}
func (UnimplementedSLBfdv6OperServer) mustEmbedUnimplementedSLBfdv6OperServer() {}

// UnsafeSLBfdv6OperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SLBfdv6OperServer will
// result in compilation errors.
type UnsafeSLBfdv6OperServer interface {
	mustEmbedUnimplementedSLBfdv6OperServer()
}

func RegisterSLBfdv6OperServer(s grpc.ServiceRegistrar, srv SLBfdv6OperServer) {
	s.RegisterService(&SLBfdv6Oper_ServiceDesc, srv)
}

func _SLBfdv6Oper_SLBfdv6RegOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLBfdRegMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLBfdv6OperServer).SLBfdv6RegOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLBfdv6Oper/SLBfdv6RegOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLBfdv6OperServer).SLBfdv6RegOp(ctx, req.(*SLBfdRegMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLBfdv6Oper_SLBfdv6Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLBfdGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLBfdv6OperServer).SLBfdv6Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLBfdv6Oper/SLBfdv6Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLBfdv6OperServer).SLBfdv6Get(ctx, req.(*SLBfdGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLBfdv6Oper_SLBfdv6GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLBfdGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLBfdv6OperServer).SLBfdv6GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLBfdv6Oper/SLBfdv6GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLBfdv6OperServer).SLBfdv6GetStats(ctx, req.(*SLBfdGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLBfdv6Oper_SLBfdv6GetNotifStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SLBfdGetNotifMsg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SLBfdv6OperServer).SLBfdv6GetNotifStream(m, &sLBfdv6OperSLBfdv6GetNotifStreamServer{stream})
}

type SLBfdv6Oper_SLBfdv6GetNotifStreamServer interface {
	Send(*SLBfdv6Notif) error
	grpc.ServerStream
}

type sLBfdv6OperSLBfdv6GetNotifStreamServer struct {
	grpc.ServerStream
}

func (x *sLBfdv6OperSLBfdv6GetNotifStreamServer) Send(m *SLBfdv6Notif) error {
	return x.ServerStream.SendMsg(m)
}

func _SLBfdv6Oper_SLBfdv6SessionOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLBfdv6Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLBfdv6OperServer).SLBfdv6SessionOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLBfdv6Oper/SLBfdv6SessionOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLBfdv6OperServer).SLBfdv6SessionOp(ctx, req.(*SLBfdv6Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLBfdv6Oper_SLBfdv6SessionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLBfdv6GetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLBfdv6OperServer).SLBfdv6SessionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLBfdv6Oper/SLBfdv6SessionGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLBfdv6OperServer).SLBfdv6SessionGet(ctx, req.(*SLBfdv6GetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// SLBfdv6Oper_ServiceDesc is the grpc.ServiceDesc for SLBfdv6Oper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SLBfdv6Oper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_layer.SLBfdv6Oper",
	HandlerType: (*SLBfdv6OperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SLBfdv6RegOp",
			Handler:    _SLBfdv6Oper_SLBfdv6RegOp_Handler,
		},
		{
			MethodName: "SLBfdv6Get",
			Handler:    _SLBfdv6Oper_SLBfdv6Get_Handler,
		},
		{
			MethodName: "SLBfdv6GetStats",
			Handler:    _SLBfdv6Oper_SLBfdv6GetStats_Handler,
		},
		{
			MethodName: "SLBfdv6SessionOp",
			Handler:    _SLBfdv6Oper_SLBfdv6SessionOp_Handler,
		},
		{
			MethodName: "SLBfdv6SessionGet",
			Handler:    _SLBfdv6Oper_SLBfdv6SessionGet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SLBfdv6GetNotifStream",
			Handler:       _SLBfdv6Oper_SLBfdv6GetNotifStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sl_bfd_ipv6.proto",
}
