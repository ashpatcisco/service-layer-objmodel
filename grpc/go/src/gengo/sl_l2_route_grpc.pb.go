// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: sl_l2_route.proto

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

// SLL2OperClient is the client API for SLL2Oper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SLL2OperClient interface {
	// Used to retrieve global L2 info from the server.
	SLL2GlobalsGet(ctx context.Context, in *SLL2GlobalsGetMsg, opts ...grpc.CallOption) (*SLL2GlobalsGetMsgRsp, error)
	// SLL2RegMsg.Oper = SL_REGOP_REGISTER:
	//
	//	Global L2 registration.
	//	A client Must Register BEFORE sending BD registration messages
	//	(to add/update/delete routes) or BEFORE requesting for L2 route
	//	notifications.
	//
	// SLL2RegMsg.Oper = SL_REGOP_UNREGISTER:
	//
	//	Global L2 un-registration.
	//	This call is used to convey that the client is no longer
	//	interested in programming L2 routes and in receiving L2 route
	//	notifications. All programmed L2 routes will be deleted on the
	//	server and the server will stop sending L2 route notifications.
	//
	// SLL2RegMsg.Oper = SL_REGOP_EOF:
	//
	//	Global L2 End Of File message.
	//	After Registration, the client is expected to send an EOF
	//	message to convey the end of replay of the client's known
	//	objects and to convey the end of requests for L2 route
	//	notifications.
	//	This is especially useful under certain restart scenarios when the
	//	client and the server are trying to synchronize their routes.
	SLL2RegOp(ctx context.Context, in *SLL2RegMsg, opts ...grpc.CallOption) (*SLL2RegMsgRsp, error)
	// SLL2BdRegMsg.Oper = SL_REGOP_REGISTER:
	//
	//	BD registration: Sends a list of BD registration messages and
	//	expects a list of registration responses.
	//	A client Must Register a BD BEFORE L2 Routes can be added/modified
	//	in that BD.
	//
	// SLL2BdRegMsg.Oper = SL_REGOP_UNREGISTER:
	//
	//	BD un-registration: Sends a list of BD un-registration messages
	//	and expects a list of un-registration responses.
	//	This can be used to convey that the client is no longer
	//	interested in programming routes in this BD. All installed L2
	//	routes will be lost.
	//
	// SLL2BdRegMsg.Oper = SL_REGOP_EOF:
	//
	//	BD End Of File message.
	//	After Registration, the client is expected to send an EOF
	//	message to convey the end of replay of the client's known objects
	//	in that BD.
	//	This is especially useful under certain restart scenarios when the
	//	client and the server are trying to synchronize their routes.
	SLL2BdRegOp(ctx context.Context, in *SLL2BdRegMsg, opts ...grpc.CallOption) (*SLL2BdRegMsgRsp, error)
	// SLL2RouteMsg.Oper = SL_OBJOP_ADD:
	// Route add. Fails if the route already exists.
	//
	// SLL2RouteMsg.Oper = SL_OBJOP_UPDATE:
	// Route update. Creates or updates the route.
	//
	// SLL2RouteMsg.Oper = SL_OBJOP_DELETE:
	// Route delete. The route path is not necessary to delete the route.
	SLL2RouteOp(ctx context.Context, in *SLL2RouteMsg, opts ...grpc.CallOption) (*SLL2RouteMsgRsp, error)
	// Stream adds/updates/deletes of L2 Routes.
	SLL2RouteOpStream(ctx context.Context, opts ...grpc.CallOption) (SLL2Oper_SLL2RouteOpStreamClient, error)
	// This call is used to get a stream of BD state and route notifications.
	// It can be used to get "push" notifications for route
	// adds/updates/deletes.
	// The caller must maintain the GRPC channel as long as there is
	// interest in route notifications.
	//
	// The call takes 3 types of notification requests:
	//  1. Request for BD state notifications only (pass only Oper and
	//     Correlator).
	//  2. Request for BD state and Route notifications in all BDs.
	//  3. Request for Route notifications per-BD.
	//     This should be sent after requesting for BD state notifications
	//     and after receiving BD-ready notification.
	//
	// The success/failure of the notification request is relayed in the
	// SLL2NotifStatusMsg followed by a Start marker, any routes if present,
	// and an End Marker.
	//
	// After all requests are sent, client should send GetNotifEof = TRUE.
	SLL2GetNotifStream(ctx context.Context, opts ...grpc.CallOption) (SLL2Oper_SLL2GetNotifStreamClient, error)
}

type sLL2OperClient struct {
	cc grpc.ClientConnInterface
}

func NewSLL2OperClient(cc grpc.ClientConnInterface) SLL2OperClient {
	return &sLL2OperClient{cc}
}

func (c *sLL2OperClient) SLL2GlobalsGet(ctx context.Context, in *SLL2GlobalsGetMsg, opts ...grpc.CallOption) (*SLL2GlobalsGetMsgRsp, error) {
	out := new(SLL2GlobalsGetMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLL2Oper/SLL2GlobalsGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLL2OperClient) SLL2RegOp(ctx context.Context, in *SLL2RegMsg, opts ...grpc.CallOption) (*SLL2RegMsgRsp, error) {
	out := new(SLL2RegMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLL2Oper/SLL2RegOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLL2OperClient) SLL2BdRegOp(ctx context.Context, in *SLL2BdRegMsg, opts ...grpc.CallOption) (*SLL2BdRegMsgRsp, error) {
	out := new(SLL2BdRegMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLL2Oper/SLL2BdRegOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLL2OperClient) SLL2RouteOp(ctx context.Context, in *SLL2RouteMsg, opts ...grpc.CallOption) (*SLL2RouteMsgRsp, error) {
	out := new(SLL2RouteMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLL2Oper/SLL2RouteOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLL2OperClient) SLL2RouteOpStream(ctx context.Context, opts ...grpc.CallOption) (SLL2Oper_SLL2RouteOpStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SLL2Oper_ServiceDesc.Streams[0], "/service_layer.SLL2Oper/SLL2RouteOpStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sLL2OperSLL2RouteOpStreamClient{stream}
	return x, nil
}

type SLL2Oper_SLL2RouteOpStreamClient interface {
	Send(*SLL2RouteMsg) error
	Recv() (*SLL2RouteMsgRsp, error)
	grpc.ClientStream
}

type sLL2OperSLL2RouteOpStreamClient struct {
	grpc.ClientStream
}

func (x *sLL2OperSLL2RouteOpStreamClient) Send(m *SLL2RouteMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sLL2OperSLL2RouteOpStreamClient) Recv() (*SLL2RouteMsgRsp, error) {
	m := new(SLL2RouteMsgRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sLL2OperClient) SLL2GetNotifStream(ctx context.Context, opts ...grpc.CallOption) (SLL2Oper_SLL2GetNotifStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SLL2Oper_ServiceDesc.Streams[1], "/service_layer.SLL2Oper/SLL2GetNotifStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sLL2OperSLL2GetNotifStreamClient{stream}
	return x, nil
}

type SLL2Oper_SLL2GetNotifStreamClient interface {
	Send(*SLL2GetNotifMsg) error
	Recv() (*SLL2Notif, error)
	grpc.ClientStream
}

type sLL2OperSLL2GetNotifStreamClient struct {
	grpc.ClientStream
}

func (x *sLL2OperSLL2GetNotifStreamClient) Send(m *SLL2GetNotifMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sLL2OperSLL2GetNotifStreamClient) Recv() (*SLL2Notif, error) {
	m := new(SLL2Notif)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SLL2OperServer is the server API for SLL2Oper service.
// All implementations must embed UnimplementedSLL2OperServer
// for forward compatibility
type SLL2OperServer interface {
	// Used to retrieve global L2 info from the server.
	SLL2GlobalsGet(context.Context, *SLL2GlobalsGetMsg) (*SLL2GlobalsGetMsgRsp, error)
	// SLL2RegMsg.Oper = SL_REGOP_REGISTER:
	//
	//	Global L2 registration.
	//	A client Must Register BEFORE sending BD registration messages
	//	(to add/update/delete routes) or BEFORE requesting for L2 route
	//	notifications.
	//
	// SLL2RegMsg.Oper = SL_REGOP_UNREGISTER:
	//
	//	Global L2 un-registration.
	//	This call is used to convey that the client is no longer
	//	interested in programming L2 routes and in receiving L2 route
	//	notifications. All programmed L2 routes will be deleted on the
	//	server and the server will stop sending L2 route notifications.
	//
	// SLL2RegMsg.Oper = SL_REGOP_EOF:
	//
	//	Global L2 End Of File message.
	//	After Registration, the client is expected to send an EOF
	//	message to convey the end of replay of the client's known
	//	objects and to convey the end of requests for L2 route
	//	notifications.
	//	This is especially useful under certain restart scenarios when the
	//	client and the server are trying to synchronize their routes.
	SLL2RegOp(context.Context, *SLL2RegMsg) (*SLL2RegMsgRsp, error)
	// SLL2BdRegMsg.Oper = SL_REGOP_REGISTER:
	//
	//	BD registration: Sends a list of BD registration messages and
	//	expects a list of registration responses.
	//	A client Must Register a BD BEFORE L2 Routes can be added/modified
	//	in that BD.
	//
	// SLL2BdRegMsg.Oper = SL_REGOP_UNREGISTER:
	//
	//	BD un-registration: Sends a list of BD un-registration messages
	//	and expects a list of un-registration responses.
	//	This can be used to convey that the client is no longer
	//	interested in programming routes in this BD. All installed L2
	//	routes will be lost.
	//
	// SLL2BdRegMsg.Oper = SL_REGOP_EOF:
	//
	//	BD End Of File message.
	//	After Registration, the client is expected to send an EOF
	//	message to convey the end of replay of the client's known objects
	//	in that BD.
	//	This is especially useful under certain restart scenarios when the
	//	client and the server are trying to synchronize their routes.
	SLL2BdRegOp(context.Context, *SLL2BdRegMsg) (*SLL2BdRegMsgRsp, error)
	// SLL2RouteMsg.Oper = SL_OBJOP_ADD:
	// Route add. Fails if the route already exists.
	//
	// SLL2RouteMsg.Oper = SL_OBJOP_UPDATE:
	// Route update. Creates or updates the route.
	//
	// SLL2RouteMsg.Oper = SL_OBJOP_DELETE:
	// Route delete. The route path is not necessary to delete the route.
	SLL2RouteOp(context.Context, *SLL2RouteMsg) (*SLL2RouteMsgRsp, error)
	// Stream adds/updates/deletes of L2 Routes.
	SLL2RouteOpStream(SLL2Oper_SLL2RouteOpStreamServer) error
	// This call is used to get a stream of BD state and route notifications.
	// It can be used to get "push" notifications for route
	// adds/updates/deletes.
	// The caller must maintain the GRPC channel as long as there is
	// interest in route notifications.
	//
	// The call takes 3 types of notification requests:
	//  1. Request for BD state notifications only (pass only Oper and
	//     Correlator).
	//  2. Request for BD state and Route notifications in all BDs.
	//  3. Request for Route notifications per-BD.
	//     This should be sent after requesting for BD state notifications
	//     and after receiving BD-ready notification.
	//
	// The success/failure of the notification request is relayed in the
	// SLL2NotifStatusMsg followed by a Start marker, any routes if present,
	// and an End Marker.
	//
	// After all requests are sent, client should send GetNotifEof = TRUE.
	SLL2GetNotifStream(SLL2Oper_SLL2GetNotifStreamServer) error
	mustEmbedUnimplementedSLL2OperServer()
}

// UnimplementedSLL2OperServer must be embedded to have forward compatible implementations.
type UnimplementedSLL2OperServer struct {
}

func (UnimplementedSLL2OperServer) SLL2GlobalsGet(context.Context, *SLL2GlobalsGetMsg) (*SLL2GlobalsGetMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLL2GlobalsGet not implemented")
}
func (UnimplementedSLL2OperServer) SLL2RegOp(context.Context, *SLL2RegMsg) (*SLL2RegMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLL2RegOp not implemented")
}
func (UnimplementedSLL2OperServer) SLL2BdRegOp(context.Context, *SLL2BdRegMsg) (*SLL2BdRegMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLL2BdRegOp not implemented")
}
func (UnimplementedSLL2OperServer) SLL2RouteOp(context.Context, *SLL2RouteMsg) (*SLL2RouteMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLL2RouteOp not implemented")
}
func (UnimplementedSLL2OperServer) SLL2RouteOpStream(SLL2Oper_SLL2RouteOpStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SLL2RouteOpStream not implemented")
}
func (UnimplementedSLL2OperServer) SLL2GetNotifStream(SLL2Oper_SLL2GetNotifStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SLL2GetNotifStream not implemented")
}
func (UnimplementedSLL2OperServer) mustEmbedUnimplementedSLL2OperServer() {}

// UnsafeSLL2OperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SLL2OperServer will
// result in compilation errors.
type UnsafeSLL2OperServer interface {
	mustEmbedUnimplementedSLL2OperServer()
}

func RegisterSLL2OperServer(s grpc.ServiceRegistrar, srv SLL2OperServer) {
	s.RegisterService(&SLL2Oper_ServiceDesc, srv)
}

func _SLL2Oper_SLL2GlobalsGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLL2GlobalsGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLL2OperServer).SLL2GlobalsGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLL2Oper/SLL2GlobalsGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLL2OperServer).SLL2GlobalsGet(ctx, req.(*SLL2GlobalsGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLL2Oper_SLL2RegOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLL2RegMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLL2OperServer).SLL2RegOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLL2Oper/SLL2RegOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLL2OperServer).SLL2RegOp(ctx, req.(*SLL2RegMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLL2Oper_SLL2BdRegOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLL2BdRegMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLL2OperServer).SLL2BdRegOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLL2Oper/SLL2BdRegOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLL2OperServer).SLL2BdRegOp(ctx, req.(*SLL2BdRegMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLL2Oper_SLL2RouteOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLL2RouteMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLL2OperServer).SLL2RouteOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLL2Oper/SLL2RouteOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLL2OperServer).SLL2RouteOp(ctx, req.(*SLL2RouteMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLL2Oper_SLL2RouteOpStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SLL2OperServer).SLL2RouteOpStream(&sLL2OperSLL2RouteOpStreamServer{stream})
}

type SLL2Oper_SLL2RouteOpStreamServer interface {
	Send(*SLL2RouteMsgRsp) error
	Recv() (*SLL2RouteMsg, error)
	grpc.ServerStream
}

type sLL2OperSLL2RouteOpStreamServer struct {
	grpc.ServerStream
}

func (x *sLL2OperSLL2RouteOpStreamServer) Send(m *SLL2RouteMsgRsp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sLL2OperSLL2RouteOpStreamServer) Recv() (*SLL2RouteMsg, error) {
	m := new(SLL2RouteMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SLL2Oper_SLL2GetNotifStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SLL2OperServer).SLL2GetNotifStream(&sLL2OperSLL2GetNotifStreamServer{stream})
}

type SLL2Oper_SLL2GetNotifStreamServer interface {
	Send(*SLL2Notif) error
	Recv() (*SLL2GetNotifMsg, error)
	grpc.ServerStream
}

type sLL2OperSLL2GetNotifStreamServer struct {
	grpc.ServerStream
}

func (x *sLL2OperSLL2GetNotifStreamServer) Send(m *SLL2Notif) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sLL2OperSLL2GetNotifStreamServer) Recv() (*SLL2GetNotifMsg, error) {
	m := new(SLL2GetNotifMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SLL2Oper_ServiceDesc is the grpc.ServiceDesc for SLL2Oper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SLL2Oper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_layer.SLL2Oper",
	HandlerType: (*SLL2OperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SLL2GlobalsGet",
			Handler:    _SLL2Oper_SLL2GlobalsGet_Handler,
		},
		{
			MethodName: "SLL2RegOp",
			Handler:    _SLL2Oper_SLL2RegOp_Handler,
		},
		{
			MethodName: "SLL2BdRegOp",
			Handler:    _SLL2Oper_SLL2BdRegOp_Handler,
		},
		{
			MethodName: "SLL2RouteOp",
			Handler:    _SLL2Oper_SLL2RouteOp_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SLL2RouteOpStream",
			Handler:       _SLL2Oper_SLL2RouteOpStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SLL2GetNotifStream",
			Handler:       _SLL2Oper_SLL2GetNotifStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sl_l2_route.proto",
}
