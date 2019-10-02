// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: sl_bfd_ipv6.proto

#include "sl_bfd_ipv6.pb.h"
#include "sl_bfd_ipv6.grpc.pb.h"

#include <grpc++/impl/codegen/async_stream.h>
#include <grpc++/impl/codegen/async_unary_call.h>
#include <grpc++/impl/codegen/channel_interface.h>
#include <grpc++/impl/codegen/client_unary_call.h>
#include <grpc++/impl/codegen/method_handler_impl.h>
#include <grpc++/impl/codegen/rpc_service_method.h>
#include <grpc++/impl/codegen/service_type.h>
#include <grpc++/impl/codegen/sync_stream.h>
namespace service_layer {

static const char* SLBfdv6Oper_method_names[] = {
  "/service_layer.SLBfdv6Oper/SLBfdv6RegOp",
  "/service_layer.SLBfdv6Oper/SLBfdv6Get",
  "/service_layer.SLBfdv6Oper/SLBfdv6GetStats",
  "/service_layer.SLBfdv6Oper/SLBfdv6GetNotifStream",
  "/service_layer.SLBfdv6Oper/SLBfdv6SessionOp",
  "/service_layer.SLBfdv6Oper/SLBfdv6SessionGet",
};

std::unique_ptr< SLBfdv6Oper::Stub> SLBfdv6Oper::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< SLBfdv6Oper::Stub> stub(new SLBfdv6Oper::Stub(channel));
  return stub;
}

SLBfdv6Oper::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_SLBfdv6RegOp_(SLBfdv6Oper_method_names[0], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SLBfdv6Get_(SLBfdv6Oper_method_names[1], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SLBfdv6GetStats_(SLBfdv6Oper_method_names[2], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SLBfdv6GetNotifStream_(SLBfdv6Oper_method_names[3], ::grpc::internal::RpcMethod::SERVER_STREAMING, channel)
  , rpcmethod_SLBfdv6SessionOp_(SLBfdv6Oper_method_names[4], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SLBfdv6SessionGet_(SLBfdv6Oper_method_names[5], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status SLBfdv6Oper::Stub::SLBfdv6RegOp(::grpc::ClientContext* context, const ::service_layer::SLBfdRegMsg& request, ::service_layer::SLBfdRegMsgRsp* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_SLBfdv6RegOp_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdRegMsgRsp>* SLBfdv6Oper::Stub::AsyncSLBfdv6RegOpRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdRegMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::service_layer::SLBfdRegMsgRsp>::Create(channel_.get(), cq, rpcmethod_SLBfdv6RegOp_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdRegMsgRsp>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6RegOpRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdRegMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::service_layer::SLBfdRegMsgRsp>::Create(channel_.get(), cq, rpcmethod_SLBfdv6RegOp_, context, request, false);
}

::grpc::Status SLBfdv6Oper::Stub::SLBfdv6Get(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::service_layer::SLBfdGetMsgRsp* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_SLBfdv6Get_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdGetMsgRsp>* SLBfdv6Oper::Stub::AsyncSLBfdv6GetRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::service_layer::SLBfdGetMsgRsp>::Create(channel_.get(), cq, rpcmethod_SLBfdv6Get_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdGetMsgRsp>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6GetRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::service_layer::SLBfdGetMsgRsp>::Create(channel_.get(), cq, rpcmethod_SLBfdv6Get_, context, request, false);
}

::grpc::Status SLBfdv6Oper::Stub::SLBfdv6GetStats(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::service_layer::SLBfdGetStatsMsgRsp* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_SLBfdv6GetStats_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdGetStatsMsgRsp>* SLBfdv6Oper::Stub::AsyncSLBfdv6GetStatsRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::service_layer::SLBfdGetStatsMsgRsp>::Create(channel_.get(), cq, rpcmethod_SLBfdv6GetStats_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdGetStatsMsgRsp>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6GetStatsRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::service_layer::SLBfdGetStatsMsgRsp>::Create(channel_.get(), cq, rpcmethod_SLBfdv6GetStats_, context, request, false);
}

::grpc::ClientReader< ::service_layer::SLBfdv6Notif>* SLBfdv6Oper::Stub::SLBfdv6GetNotifStreamRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetNotifMsg& request) {
  return ::grpc::internal::ClientReaderFactory< ::service_layer::SLBfdv6Notif>::Create(channel_.get(), rpcmethod_SLBfdv6GetNotifStream_, context, request);
}

::grpc::ClientAsyncReader< ::service_layer::SLBfdv6Notif>* SLBfdv6Oper::Stub::AsyncSLBfdv6GetNotifStreamRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetNotifMsg& request, ::grpc::CompletionQueue* cq, void* tag) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::service_layer::SLBfdv6Notif>::Create(channel_.get(), cq, rpcmethod_SLBfdv6GetNotifStream_, context, request, true, tag);
}

::grpc::ClientAsyncReader< ::service_layer::SLBfdv6Notif>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6GetNotifStreamRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetNotifMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::service_layer::SLBfdv6Notif>::Create(channel_.get(), cq, rpcmethod_SLBfdv6GetNotifStream_, context, request, false, nullptr);
}

::grpc::Status SLBfdv6Oper::Stub::SLBfdv6SessionOp(::grpc::ClientContext* context, const ::service_layer::SLBfdv6Msg& request, ::service_layer::SLBfdv6MsgRsp* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_SLBfdv6SessionOp_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdv6MsgRsp>* SLBfdv6Oper::Stub::AsyncSLBfdv6SessionOpRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdv6Msg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::service_layer::SLBfdv6MsgRsp>::Create(channel_.get(), cq, rpcmethod_SLBfdv6SessionOp_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdv6MsgRsp>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6SessionOpRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdv6Msg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::service_layer::SLBfdv6MsgRsp>::Create(channel_.get(), cq, rpcmethod_SLBfdv6SessionOp_, context, request, false);
}

::grpc::Status SLBfdv6Oper::Stub::SLBfdv6SessionGet(::grpc::ClientContext* context, const ::service_layer::SLBfdv6GetMsg& request, ::service_layer::SLBfdv6GetMsgRsp* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_SLBfdv6SessionGet_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdv6GetMsgRsp>* SLBfdv6Oper::Stub::AsyncSLBfdv6SessionGetRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdv6GetMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::service_layer::SLBfdv6GetMsgRsp>::Create(channel_.get(), cq, rpcmethod_SLBfdv6SessionGet_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdv6GetMsgRsp>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6SessionGetRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdv6GetMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::service_layer::SLBfdv6GetMsgRsp>::Create(channel_.get(), cq, rpcmethod_SLBfdv6SessionGet_, context, request, false);
}

SLBfdv6Oper::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdRegMsg, ::service_layer::SLBfdRegMsgRsp>(
          std::mem_fn(&SLBfdv6Oper::Service::SLBfdv6RegOp), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdGetMsg, ::service_layer::SLBfdGetMsgRsp>(
          std::mem_fn(&SLBfdv6Oper::Service::SLBfdv6Get), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdGetMsg, ::service_layer::SLBfdGetStatsMsgRsp>(
          std::mem_fn(&SLBfdv6Oper::Service::SLBfdv6GetStats), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[3],
      ::grpc::internal::RpcMethod::SERVER_STREAMING,
      new ::grpc::internal::ServerStreamingHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdGetNotifMsg, ::service_layer::SLBfdv6Notif>(
          std::mem_fn(&SLBfdv6Oper::Service::SLBfdv6GetNotifStream), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[4],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdv6Msg, ::service_layer::SLBfdv6MsgRsp>(
          std::mem_fn(&SLBfdv6Oper::Service::SLBfdv6SessionOp), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[5],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdv6GetMsg, ::service_layer::SLBfdv6GetMsgRsp>(
          std::mem_fn(&SLBfdv6Oper::Service::SLBfdv6SessionGet), this)));
}

SLBfdv6Oper::Service::~Service() {
}

::grpc::Status SLBfdv6Oper::Service::SLBfdv6RegOp(::grpc::ServerContext* context, const ::service_layer::SLBfdRegMsg* request, ::service_layer::SLBfdRegMsgRsp* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SLBfdv6Oper::Service::SLBfdv6Get(::grpc::ServerContext* context, const ::service_layer::SLBfdGetMsg* request, ::service_layer::SLBfdGetMsgRsp* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SLBfdv6Oper::Service::SLBfdv6GetStats(::grpc::ServerContext* context, const ::service_layer::SLBfdGetMsg* request, ::service_layer::SLBfdGetStatsMsgRsp* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SLBfdv6Oper::Service::SLBfdv6GetNotifStream(::grpc::ServerContext* context, const ::service_layer::SLBfdGetNotifMsg* request, ::grpc::ServerWriter< ::service_layer::SLBfdv6Notif>* writer) {
  (void) context;
  (void) request;
  (void) writer;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SLBfdv6Oper::Service::SLBfdv6SessionOp(::grpc::ServerContext* context, const ::service_layer::SLBfdv6Msg* request, ::service_layer::SLBfdv6MsgRsp* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status SLBfdv6Oper::Service::SLBfdv6SessionGet(::grpc::ServerContext* context, const ::service_layer::SLBfdv6GetMsg* request, ::service_layer::SLBfdv6GetMsgRsp* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace service_layer

