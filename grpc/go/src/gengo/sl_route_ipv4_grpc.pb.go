// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: sl_route_ipv4.proto

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

// SLRoutev4OperClient is the client API for SLRoutev4Oper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SLRoutev4OperClient interface {
	// Used to retrieve Global Route information
	SLRoutev4GlobalsGet(ctx context.Context, in *SLRouteGlobalsGetMsg, opts ...grpc.CallOption) (*SLRouteGlobalsGetMsgRsp, error)
	// Used to retrieve Global Route Stats
	SLRoutev4GlobalStatsGet(ctx context.Context, in *SLRouteGlobalStatsGetMsg, opts ...grpc.CallOption) (*SLRouteGlobalStatsGetMsgRsp, error)
	// SLVrfRegMsg.Oper = SL_REGOP_REGISTER:
	//
	//	 VRF registration: Sends a list of VRF registration messages
	//	 and expects a list of registration responses.
	//	 A client Must Register a VRF BEFORE routes can be added/modified in
	//	the associated VRF.
	//
	// SLVrfRegMsg.Oper = SL_REGOP_UNREGISTER:
	//
	//	VRF Un-registeration: Sends a list of VRF un-registration messages
	//	and expects a list of un-registration responses.
	//	This can be used to convey that the client is no longer interested
	//	in this VRF. All previously installed routes would be lost.
	//
	// SLVrfRegMsg.Oper = SL_REGOP_EOF:
	//
	//	VRF End Of File message.
	//	After Registration, the client is expected to send an EOF
	//	message to convey the end of replay of the client's known objects.
	//	This is especially useful under certain restart scenarios when the
	//	client and the server are trying to synchronize their Routes.
	SLRoutev4VrfRegOp(ctx context.Context, in *SLVrfRegMsg, opts ...grpc.CallOption) (*SLVrfRegMsgRsp, error)
	// VRF get. Used to retrieve VRF attributes from the server.
	SLRoutev4VrfRegGet(ctx context.Context, in *SLVrfRegGetMsg, opts ...grpc.CallOption) (*SLVrfRegGetMsgRsp, error)
	// Used to retrieve VRF Stats from the server.
	SLRoutev4VrfGetStats(ctx context.Context, in *SLVrfRegGetMsg, opts ...grpc.CallOption) (*SLVRFGetStatsMsgRsp, error)
	// SLRoutev4Msg.Oper = SL_OBJOP_ADD:
	//
	//	Route add. Fails if the route already exists.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_UPDATE:
	//
	//	Route update. Creates or updates the route.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_DELETE:
	//
	//	Route delete. The route path is not necessary to delete the route.
	SLRoutev4Op(ctx context.Context, in *SLRoutev4Msg, opts ...grpc.CallOption) (*SLRoutev4MsgRsp, error)
	// Retrieves route attributes.
	SLRoutev4Get(ctx context.Context, in *SLRoutev4GetMsg, opts ...grpc.CallOption) (*SLRoutev4GetMsgRsp, error)
	// SLRoutev4Msg.Oper = SL_OBJOP_ADD:
	//
	//	Route add. Fails if the route already exists.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_UPDATE:
	//
	//	Route update. Creates or updates the route.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_DELETE:
	//
	//	Route delete. The route path is not necessary to delete the route.
	SLRoutev4OpStream(ctx context.Context, opts ...grpc.CallOption) (SLRoutev4Oper_SLRoutev4OpStreamClient, error)
	// Retrieves route attributes.
	SLRoutev4GetStream(ctx context.Context, opts ...grpc.CallOption) (SLRoutev4Oper_SLRoutev4GetStreamClient, error)
	// This call is used to get a stream of route notifications.
	// It can be used to get "push" notifications for route
	// adds/updates/deletes.
	// The caller must maintain the GRPC channel as long as there is
	// interest in route notifications.
	//
	// The call takes a stream of per-VRF notification requests.
	// The success/failure of the notification request is relayed in the
	// SLRouteNotifStatus followed by a Start marker, any routes if present,
	// and an End Marker.
	SLRoutev4GetNotifStream(ctx context.Context, opts ...grpc.CallOption) (SLRoutev4Oper_SLRoutev4GetNotifStreamClient, error)
}

type sLRoutev4OperClient struct {
	cc grpc.ClientConnInterface
}

func NewSLRoutev4OperClient(cc grpc.ClientConnInterface) SLRoutev4OperClient {
	return &sLRoutev4OperClient{cc}
}

func (c *sLRoutev4OperClient) SLRoutev4GlobalsGet(ctx context.Context, in *SLRouteGlobalsGetMsg, opts ...grpc.CallOption) (*SLRouteGlobalsGetMsgRsp, error) {
	out := new(SLRouteGlobalsGetMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4GlobalsGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4GlobalStatsGet(ctx context.Context, in *SLRouteGlobalStatsGetMsg, opts ...grpc.CallOption) (*SLRouteGlobalStatsGetMsgRsp, error) {
	out := new(SLRouteGlobalStatsGetMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4GlobalStatsGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4VrfRegOp(ctx context.Context, in *SLVrfRegMsg, opts ...grpc.CallOption) (*SLVrfRegMsgRsp, error) {
	out := new(SLVrfRegMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4VrfRegOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4VrfRegGet(ctx context.Context, in *SLVrfRegGetMsg, opts ...grpc.CallOption) (*SLVrfRegGetMsgRsp, error) {
	out := new(SLVrfRegGetMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4VrfRegGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4VrfGetStats(ctx context.Context, in *SLVrfRegGetMsg, opts ...grpc.CallOption) (*SLVRFGetStatsMsgRsp, error) {
	out := new(SLVRFGetStatsMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4VrfGetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4Op(ctx context.Context, in *SLRoutev4Msg, opts ...grpc.CallOption) (*SLRoutev4MsgRsp, error) {
	out := new(SLRoutev4MsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4Op", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4Get(ctx context.Context, in *SLRoutev4GetMsg, opts ...grpc.CallOption) (*SLRoutev4GetMsgRsp, error) {
	out := new(SLRoutev4GetMsgRsp)
	err := c.cc.Invoke(ctx, "/service_layer.SLRoutev4Oper/SLRoutev4Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLRoutev4OperClient) SLRoutev4OpStream(ctx context.Context, opts ...grpc.CallOption) (SLRoutev4Oper_SLRoutev4OpStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SLRoutev4Oper_ServiceDesc.Streams[0], "/service_layer.SLRoutev4Oper/SLRoutev4OpStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sLRoutev4OperSLRoutev4OpStreamClient{stream}
	return x, nil
}

type SLRoutev4Oper_SLRoutev4OpStreamClient interface {
	Send(*SLRoutev4Msg) error
	Recv() (*SLRoutev4MsgRsp, error)
	grpc.ClientStream
}

type sLRoutev4OperSLRoutev4OpStreamClient struct {
	grpc.ClientStream
}

func (x *sLRoutev4OperSLRoutev4OpStreamClient) Send(m *SLRoutev4Msg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sLRoutev4OperSLRoutev4OpStreamClient) Recv() (*SLRoutev4MsgRsp, error) {
	m := new(SLRoutev4MsgRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sLRoutev4OperClient) SLRoutev4GetStream(ctx context.Context, opts ...grpc.CallOption) (SLRoutev4Oper_SLRoutev4GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SLRoutev4Oper_ServiceDesc.Streams[1], "/service_layer.SLRoutev4Oper/SLRoutev4GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sLRoutev4OperSLRoutev4GetStreamClient{stream}
	return x, nil
}

type SLRoutev4Oper_SLRoutev4GetStreamClient interface {
	Send(*SLRoutev4GetMsg) error
	Recv() (*SLRoutev4GetMsgRsp, error)
	grpc.ClientStream
}

type sLRoutev4OperSLRoutev4GetStreamClient struct {
	grpc.ClientStream
}

func (x *sLRoutev4OperSLRoutev4GetStreamClient) Send(m *SLRoutev4GetMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sLRoutev4OperSLRoutev4GetStreamClient) Recv() (*SLRoutev4GetMsgRsp, error) {
	m := new(SLRoutev4GetMsgRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sLRoutev4OperClient) SLRoutev4GetNotifStream(ctx context.Context, opts ...grpc.CallOption) (SLRoutev4Oper_SLRoutev4GetNotifStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SLRoutev4Oper_ServiceDesc.Streams[2], "/service_layer.SLRoutev4Oper/SLRoutev4GetNotifStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sLRoutev4OperSLRoutev4GetNotifStreamClient{stream}
	return x, nil
}

type SLRoutev4Oper_SLRoutev4GetNotifStreamClient interface {
	Send(*SLRouteGetNotifMsg) error
	Recv() (*SLRoutev4Notif, error)
	grpc.ClientStream
}

type sLRoutev4OperSLRoutev4GetNotifStreamClient struct {
	grpc.ClientStream
}

func (x *sLRoutev4OperSLRoutev4GetNotifStreamClient) Send(m *SLRouteGetNotifMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sLRoutev4OperSLRoutev4GetNotifStreamClient) Recv() (*SLRoutev4Notif, error) {
	m := new(SLRoutev4Notif)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SLRoutev4OperServer is the server API for SLRoutev4Oper service.
// All implementations must embed UnimplementedSLRoutev4OperServer
// for forward compatibility
type SLRoutev4OperServer interface {
	// Used to retrieve Global Route information
	SLRoutev4GlobalsGet(context.Context, *SLRouteGlobalsGetMsg) (*SLRouteGlobalsGetMsgRsp, error)
	// Used to retrieve Global Route Stats
	SLRoutev4GlobalStatsGet(context.Context, *SLRouteGlobalStatsGetMsg) (*SLRouteGlobalStatsGetMsgRsp, error)
	// SLVrfRegMsg.Oper = SL_REGOP_REGISTER:
	//
	//	 VRF registration: Sends a list of VRF registration messages
	//	 and expects a list of registration responses.
	//	 A client Must Register a VRF BEFORE routes can be added/modified in
	//	the associated VRF.
	//
	// SLVrfRegMsg.Oper = SL_REGOP_UNREGISTER:
	//
	//	VRF Un-registeration: Sends a list of VRF un-registration messages
	//	and expects a list of un-registration responses.
	//	This can be used to convey that the client is no longer interested
	//	in this VRF. All previously installed routes would be lost.
	//
	// SLVrfRegMsg.Oper = SL_REGOP_EOF:
	//
	//	VRF End Of File message.
	//	After Registration, the client is expected to send an EOF
	//	message to convey the end of replay of the client's known objects.
	//	This is especially useful under certain restart scenarios when the
	//	client and the server are trying to synchronize their Routes.
	SLRoutev4VrfRegOp(context.Context, *SLVrfRegMsg) (*SLVrfRegMsgRsp, error)
	// VRF get. Used to retrieve VRF attributes from the server.
	SLRoutev4VrfRegGet(context.Context, *SLVrfRegGetMsg) (*SLVrfRegGetMsgRsp, error)
	// Used to retrieve VRF Stats from the server.
	SLRoutev4VrfGetStats(context.Context, *SLVrfRegGetMsg) (*SLVRFGetStatsMsgRsp, error)
	// SLRoutev4Msg.Oper = SL_OBJOP_ADD:
	//
	//	Route add. Fails if the route already exists.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_UPDATE:
	//
	//	Route update. Creates or updates the route.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_DELETE:
	//
	//	Route delete. The route path is not necessary to delete the route.
	SLRoutev4Op(context.Context, *SLRoutev4Msg) (*SLRoutev4MsgRsp, error)
	// Retrieves route attributes.
	SLRoutev4Get(context.Context, *SLRoutev4GetMsg) (*SLRoutev4GetMsgRsp, error)
	// SLRoutev4Msg.Oper = SL_OBJOP_ADD:
	//
	//	Route add. Fails if the route already exists.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_UPDATE:
	//
	//	Route update. Creates or updates the route.
	//
	// SLRoutev4Msg.Oper = SL_OBJOP_DELETE:
	//
	//	Route delete. The route path is not necessary to delete the route.
	SLRoutev4OpStream(SLRoutev4Oper_SLRoutev4OpStreamServer) error
	// Retrieves route attributes.
	SLRoutev4GetStream(SLRoutev4Oper_SLRoutev4GetStreamServer) error
	// This call is used to get a stream of route notifications.
	// It can be used to get "push" notifications for route
	// adds/updates/deletes.
	// The caller must maintain the GRPC channel as long as there is
	// interest in route notifications.
	//
	// The call takes a stream of per-VRF notification requests.
	// The success/failure of the notification request is relayed in the
	// SLRouteNotifStatus followed by a Start marker, any routes if present,
	// and an End Marker.
	SLRoutev4GetNotifStream(SLRoutev4Oper_SLRoutev4GetNotifStreamServer) error
	mustEmbedUnimplementedSLRoutev4OperServer()
}

// UnimplementedSLRoutev4OperServer must be embedded to have forward compatible implementations.
type UnimplementedSLRoutev4OperServer struct {
}

func (UnimplementedSLRoutev4OperServer) SLRoutev4GlobalsGet(context.Context, *SLRouteGlobalsGetMsg) (*SLRouteGlobalsGetMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLRoutev4GlobalsGet not implemented")
}
func (UnimplementedSLRoutev4OperServer) SLRoutev4GlobalStatsGet(context.Context, *SLRouteGlobalStatsGetMsg) (*SLRouteGlobalStatsGetMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLRoutev4GlobalStatsGet not implemented")
}
func (UnimplementedSLRoutev4OperServer) SLRoutev4VrfRegOp(context.Context, *SLVrfRegMsg) (*SLVrfRegMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLRoutev4VrfRegOp not implemented")
}
func (UnimplementedSLRoutev4OperServer) SLRoutev4VrfRegGet(context.Context, *SLVrfRegGetMsg) (*SLVrfRegGetMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLRoutev4VrfRegGet not implemented")
}
func (UnimplementedSLRoutev4OperServer) SLRoutev4VrfGetStats(context.Context, *SLVrfRegGetMsg) (*SLVRFGetStatsMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLRoutev4VrfGetStats not implemented")
}
func (UnimplementedSLRoutev4OperServer) SLRoutev4Op(context.Context, *SLRoutev4Msg) (*SLRoutev4MsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLRoutev4Op not implemented")
}
func (UnimplementedSLRoutev4OperServer) SLRoutev4Get(context.Context, *SLRoutev4GetMsg) (*SLRoutev4GetMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SLRoutev4Get not implemented")
}
func (UnimplementedSLRoutev4OperServer) SLRoutev4OpStream(SLRoutev4Oper_SLRoutev4OpStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SLRoutev4OpStream not implemented")
}
func (UnimplementedSLRoutev4OperServer) SLRoutev4GetStream(SLRoutev4Oper_SLRoutev4GetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SLRoutev4GetStream not implemented")
}
func (UnimplementedSLRoutev4OperServer) SLRoutev4GetNotifStream(SLRoutev4Oper_SLRoutev4GetNotifStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SLRoutev4GetNotifStream not implemented")
}
func (UnimplementedSLRoutev4OperServer) mustEmbedUnimplementedSLRoutev4OperServer() {}

// UnsafeSLRoutev4OperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SLRoutev4OperServer will
// result in compilation errors.
type UnsafeSLRoutev4OperServer interface {
	mustEmbedUnimplementedSLRoutev4OperServer()
}

func RegisterSLRoutev4OperServer(s grpc.ServiceRegistrar, srv SLRoutev4OperServer) {
	s.RegisterService(&SLRoutev4Oper_ServiceDesc, srv)
}

func _SLRoutev4Oper_SLRoutev4GlobalsGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLRouteGlobalsGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4GlobalsGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4GlobalsGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4GlobalsGet(ctx, req.(*SLRouteGlobalsGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4GlobalStatsGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLRouteGlobalStatsGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4GlobalStatsGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4GlobalStatsGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4GlobalStatsGet(ctx, req.(*SLRouteGlobalStatsGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4VrfRegOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLVrfRegMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfRegOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4VrfRegOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfRegOp(ctx, req.(*SLVrfRegMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4VrfRegGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLVrfRegGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfRegGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4VrfRegGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfRegGet(ctx, req.(*SLVrfRegGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4VrfGetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLVrfRegGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfGetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4VrfGetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4VrfGetStats(ctx, req.(*SLVrfRegGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4Op_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLRoutev4Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4Op(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4Op",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4Op(ctx, req.(*SLRoutev4Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLRoutev4GetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLRoutev4OperServer).SLRoutev4Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLRoutev4Oper/SLRoutev4Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLRoutev4OperServer).SLRoutev4Get(ctx, req.(*SLRoutev4GetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLRoutev4Oper_SLRoutev4OpStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SLRoutev4OperServer).SLRoutev4OpStream(&sLRoutev4OperSLRoutev4OpStreamServer{stream})
}

type SLRoutev4Oper_SLRoutev4OpStreamServer interface {
	Send(*SLRoutev4MsgRsp) error
	Recv() (*SLRoutev4Msg, error)
	grpc.ServerStream
}

type sLRoutev4OperSLRoutev4OpStreamServer struct {
	grpc.ServerStream
}

func (x *sLRoutev4OperSLRoutev4OpStreamServer) Send(m *SLRoutev4MsgRsp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sLRoutev4OperSLRoutev4OpStreamServer) Recv() (*SLRoutev4Msg, error) {
	m := new(SLRoutev4Msg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SLRoutev4Oper_SLRoutev4GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SLRoutev4OperServer).SLRoutev4GetStream(&sLRoutev4OperSLRoutev4GetStreamServer{stream})
}

type SLRoutev4Oper_SLRoutev4GetStreamServer interface {
	Send(*SLRoutev4GetMsgRsp) error
	Recv() (*SLRoutev4GetMsg, error)
	grpc.ServerStream
}

type sLRoutev4OperSLRoutev4GetStreamServer struct {
	grpc.ServerStream
}

func (x *sLRoutev4OperSLRoutev4GetStreamServer) Send(m *SLRoutev4GetMsgRsp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sLRoutev4OperSLRoutev4GetStreamServer) Recv() (*SLRoutev4GetMsg, error) {
	m := new(SLRoutev4GetMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SLRoutev4Oper_SLRoutev4GetNotifStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SLRoutev4OperServer).SLRoutev4GetNotifStream(&sLRoutev4OperSLRoutev4GetNotifStreamServer{stream})
}

type SLRoutev4Oper_SLRoutev4GetNotifStreamServer interface {
	Send(*SLRoutev4Notif) error
	Recv() (*SLRouteGetNotifMsg, error)
	grpc.ServerStream
}

type sLRoutev4OperSLRoutev4GetNotifStreamServer struct {
	grpc.ServerStream
}

func (x *sLRoutev4OperSLRoutev4GetNotifStreamServer) Send(m *SLRoutev4Notif) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sLRoutev4OperSLRoutev4GetNotifStreamServer) Recv() (*SLRouteGetNotifMsg, error) {
	m := new(SLRouteGetNotifMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SLRoutev4Oper_ServiceDesc is the grpc.ServiceDesc for SLRoutev4Oper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SLRoutev4Oper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_layer.SLRoutev4Oper",
	HandlerType: (*SLRoutev4OperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SLRoutev4GlobalsGet",
			Handler:    _SLRoutev4Oper_SLRoutev4GlobalsGet_Handler,
		},
		{
			MethodName: "SLRoutev4GlobalStatsGet",
			Handler:    _SLRoutev4Oper_SLRoutev4GlobalStatsGet_Handler,
		},
		{
			MethodName: "SLRoutev4VrfRegOp",
			Handler:    _SLRoutev4Oper_SLRoutev4VrfRegOp_Handler,
		},
		{
			MethodName: "SLRoutev4VrfRegGet",
			Handler:    _SLRoutev4Oper_SLRoutev4VrfRegGet_Handler,
		},
		{
			MethodName: "SLRoutev4VrfGetStats",
			Handler:    _SLRoutev4Oper_SLRoutev4VrfGetStats_Handler,
		},
		{
			MethodName: "SLRoutev4Op",
			Handler:    _SLRoutev4Oper_SLRoutev4Op_Handler,
		},
		{
			MethodName: "SLRoutev4Get",
			Handler:    _SLRoutev4Oper_SLRoutev4Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SLRoutev4OpStream",
			Handler:       _SLRoutev4Oper_SLRoutev4OpStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SLRoutev4GetStream",
			Handler:       _SLRoutev4Oper_SLRoutev4GetStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SLRoutev4GetNotifStream",
			Handler:       _SLRoutev4Oper_SLRoutev4GetNotifStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sl_route_ipv4.proto",
}
