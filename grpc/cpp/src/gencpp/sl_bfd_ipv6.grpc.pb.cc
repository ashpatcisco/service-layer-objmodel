// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: sl_bfd_ipv6.proto

#include "sl_bfd_ipv6.pb.h"
#include "sl_bfd_ipv6.grpc.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/message_allocator.h>
#include <grpcpp/impl/codegen/method_handler.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/server_callback_handlers.h>
#include <grpcpp/impl/codegen/server_context.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>
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
  std::unique_ptr< SLBfdv6Oper::Stub> stub(new SLBfdv6Oper::Stub(channel, options));
  return stub;
}

SLBfdv6Oper::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_SLBfdv6RegOp_(SLBfdv6Oper_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SLBfdv6Get_(SLBfdv6Oper_method_names[1], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SLBfdv6GetStats_(SLBfdv6Oper_method_names[2], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SLBfdv6GetNotifStream_(SLBfdv6Oper_method_names[3], options.suffix_for_stats(),::grpc::internal::RpcMethod::SERVER_STREAMING, channel)
  , rpcmethod_SLBfdv6SessionOp_(SLBfdv6Oper_method_names[4], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_SLBfdv6SessionGet_(SLBfdv6Oper_method_names[5], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status SLBfdv6Oper::Stub::SLBfdv6RegOp(::grpc::ClientContext* context, const ::service_layer::SLBfdRegMsg& request, ::service_layer::SLBfdRegMsgRsp* response) {
  return ::grpc::internal::BlockingUnaryCall< ::service_layer::SLBfdRegMsg, ::service_layer::SLBfdRegMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SLBfdv6RegOp_, context, request, response);
}

void SLBfdv6Oper::Stub::async::SLBfdv6RegOp(::grpc::ClientContext* context, const ::service_layer::SLBfdRegMsg* request, ::service_layer::SLBfdRegMsgRsp* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::service_layer::SLBfdRegMsg, ::service_layer::SLBfdRegMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SLBfdv6RegOp_, context, request, response, std::move(f));
}

void SLBfdv6Oper::Stub::async::SLBfdv6RegOp(::grpc::ClientContext* context, const ::service_layer::SLBfdRegMsg* request, ::service_layer::SLBfdRegMsgRsp* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SLBfdv6RegOp_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdRegMsgRsp>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6RegOpRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdRegMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::service_layer::SLBfdRegMsgRsp, ::service_layer::SLBfdRegMsg, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SLBfdv6RegOp_, context, request);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdRegMsgRsp>* SLBfdv6Oper::Stub::AsyncSLBfdv6RegOpRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdRegMsg& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSLBfdv6RegOpRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status SLBfdv6Oper::Stub::SLBfdv6Get(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::service_layer::SLBfdGetMsgRsp* response) {
  return ::grpc::internal::BlockingUnaryCall< ::service_layer::SLBfdGetMsg, ::service_layer::SLBfdGetMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SLBfdv6Get_, context, request, response);
}

void SLBfdv6Oper::Stub::async::SLBfdv6Get(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg* request, ::service_layer::SLBfdGetMsgRsp* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::service_layer::SLBfdGetMsg, ::service_layer::SLBfdGetMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SLBfdv6Get_, context, request, response, std::move(f));
}

void SLBfdv6Oper::Stub::async::SLBfdv6Get(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg* request, ::service_layer::SLBfdGetMsgRsp* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SLBfdv6Get_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdGetMsgRsp>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6GetRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::service_layer::SLBfdGetMsgRsp, ::service_layer::SLBfdGetMsg, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SLBfdv6Get_, context, request);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdGetMsgRsp>* SLBfdv6Oper::Stub::AsyncSLBfdv6GetRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSLBfdv6GetRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status SLBfdv6Oper::Stub::SLBfdv6GetStats(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::service_layer::SLBfdGetStatsMsgRsp* response) {
  return ::grpc::internal::BlockingUnaryCall< ::service_layer::SLBfdGetMsg, ::service_layer::SLBfdGetStatsMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SLBfdv6GetStats_, context, request, response);
}

void SLBfdv6Oper::Stub::async::SLBfdv6GetStats(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg* request, ::service_layer::SLBfdGetStatsMsgRsp* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::service_layer::SLBfdGetMsg, ::service_layer::SLBfdGetStatsMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SLBfdv6GetStats_, context, request, response, std::move(f));
}

void SLBfdv6Oper::Stub::async::SLBfdv6GetStats(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg* request, ::service_layer::SLBfdGetStatsMsgRsp* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SLBfdv6GetStats_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdGetStatsMsgRsp>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6GetStatsRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::service_layer::SLBfdGetStatsMsgRsp, ::service_layer::SLBfdGetMsg, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SLBfdv6GetStats_, context, request);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdGetStatsMsgRsp>* SLBfdv6Oper::Stub::AsyncSLBfdv6GetStatsRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetMsg& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSLBfdv6GetStatsRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::ClientReader< ::service_layer::SLBfdv6Notif>* SLBfdv6Oper::Stub::SLBfdv6GetNotifStreamRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetNotifMsg& request) {
  return ::grpc::internal::ClientReaderFactory< ::service_layer::SLBfdv6Notif>::Create(channel_.get(), rpcmethod_SLBfdv6GetNotifStream_, context, request);
}

void SLBfdv6Oper::Stub::async::SLBfdv6GetNotifStream(::grpc::ClientContext* context, const ::service_layer::SLBfdGetNotifMsg* request, ::grpc::ClientReadReactor< ::service_layer::SLBfdv6Notif>* reactor) {
  ::grpc::internal::ClientCallbackReaderFactory< ::service_layer::SLBfdv6Notif>::Create(stub_->channel_.get(), stub_->rpcmethod_SLBfdv6GetNotifStream_, context, request, reactor);
}

::grpc::ClientAsyncReader< ::service_layer::SLBfdv6Notif>* SLBfdv6Oper::Stub::AsyncSLBfdv6GetNotifStreamRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetNotifMsg& request, ::grpc::CompletionQueue* cq, void* tag) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::service_layer::SLBfdv6Notif>::Create(channel_.get(), cq, rpcmethod_SLBfdv6GetNotifStream_, context, request, true, tag);
}

::grpc::ClientAsyncReader< ::service_layer::SLBfdv6Notif>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6GetNotifStreamRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdGetNotifMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncReaderFactory< ::service_layer::SLBfdv6Notif>::Create(channel_.get(), cq, rpcmethod_SLBfdv6GetNotifStream_, context, request, false, nullptr);
}

::grpc::Status SLBfdv6Oper::Stub::SLBfdv6SessionOp(::grpc::ClientContext* context, const ::service_layer::SLBfdv6Msg& request, ::service_layer::SLBfdv6MsgRsp* response) {
  return ::grpc::internal::BlockingUnaryCall< ::service_layer::SLBfdv6Msg, ::service_layer::SLBfdv6MsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SLBfdv6SessionOp_, context, request, response);
}

void SLBfdv6Oper::Stub::async::SLBfdv6SessionOp(::grpc::ClientContext* context, const ::service_layer::SLBfdv6Msg* request, ::service_layer::SLBfdv6MsgRsp* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::service_layer::SLBfdv6Msg, ::service_layer::SLBfdv6MsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SLBfdv6SessionOp_, context, request, response, std::move(f));
}

void SLBfdv6Oper::Stub::async::SLBfdv6SessionOp(::grpc::ClientContext* context, const ::service_layer::SLBfdv6Msg* request, ::service_layer::SLBfdv6MsgRsp* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SLBfdv6SessionOp_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdv6MsgRsp>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6SessionOpRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdv6Msg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::service_layer::SLBfdv6MsgRsp, ::service_layer::SLBfdv6Msg, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SLBfdv6SessionOp_, context, request);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdv6MsgRsp>* SLBfdv6Oper::Stub::AsyncSLBfdv6SessionOpRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdv6Msg& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSLBfdv6SessionOpRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status SLBfdv6Oper::Stub::SLBfdv6SessionGet(::grpc::ClientContext* context, const ::service_layer::SLBfdv6GetMsg& request, ::service_layer::SLBfdv6GetMsgRsp* response) {
  return ::grpc::internal::BlockingUnaryCall< ::service_layer::SLBfdv6GetMsg, ::service_layer::SLBfdv6GetMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SLBfdv6SessionGet_, context, request, response);
}

void SLBfdv6Oper::Stub::async::SLBfdv6SessionGet(::grpc::ClientContext* context, const ::service_layer::SLBfdv6GetMsg* request, ::service_layer::SLBfdv6GetMsgRsp* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::service_layer::SLBfdv6GetMsg, ::service_layer::SLBfdv6GetMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SLBfdv6SessionGet_, context, request, response, std::move(f));
}

void SLBfdv6Oper::Stub::async::SLBfdv6SessionGet(::grpc::ClientContext* context, const ::service_layer::SLBfdv6GetMsg* request, ::service_layer::SLBfdv6GetMsgRsp* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SLBfdv6SessionGet_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdv6GetMsgRsp>* SLBfdv6Oper::Stub::PrepareAsyncSLBfdv6SessionGetRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdv6GetMsg& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::service_layer::SLBfdv6GetMsgRsp, ::service_layer::SLBfdv6GetMsg, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SLBfdv6SessionGet_, context, request);
}

::grpc::ClientAsyncResponseReader< ::service_layer::SLBfdv6GetMsgRsp>* SLBfdv6Oper::Stub::AsyncSLBfdv6SessionGetRaw(::grpc::ClientContext* context, const ::service_layer::SLBfdv6GetMsg& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSLBfdv6SessionGetRaw(context, request, cq);
  result->StartCall();
  return result;
}

SLBfdv6Oper::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdRegMsg, ::service_layer::SLBfdRegMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SLBfdv6Oper::Service* service,
             ::grpc::ServerContext* ctx,
             const ::service_layer::SLBfdRegMsg* req,
             ::service_layer::SLBfdRegMsgRsp* resp) {
               return service->SLBfdv6RegOp(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdGetMsg, ::service_layer::SLBfdGetMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SLBfdv6Oper::Service* service,
             ::grpc::ServerContext* ctx,
             const ::service_layer::SLBfdGetMsg* req,
             ::service_layer::SLBfdGetMsgRsp* resp) {
               return service->SLBfdv6Get(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdGetMsg, ::service_layer::SLBfdGetStatsMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SLBfdv6Oper::Service* service,
             ::grpc::ServerContext* ctx,
             const ::service_layer::SLBfdGetMsg* req,
             ::service_layer::SLBfdGetStatsMsgRsp* resp) {
               return service->SLBfdv6GetStats(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[3],
      ::grpc::internal::RpcMethod::SERVER_STREAMING,
      new ::grpc::internal::ServerStreamingHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdGetNotifMsg, ::service_layer::SLBfdv6Notif>(
          [](SLBfdv6Oper::Service* service,
             ::grpc::ServerContext* ctx,
             const ::service_layer::SLBfdGetNotifMsg* req,
             ::grpc::ServerWriter<::service_layer::SLBfdv6Notif>* writer) {
               return service->SLBfdv6GetNotifStream(ctx, req, writer);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[4],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdv6Msg, ::service_layer::SLBfdv6MsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SLBfdv6Oper::Service* service,
             ::grpc::ServerContext* ctx,
             const ::service_layer::SLBfdv6Msg* req,
             ::service_layer::SLBfdv6MsgRsp* resp) {
               return service->SLBfdv6SessionOp(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SLBfdv6Oper_method_names[5],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SLBfdv6Oper::Service, ::service_layer::SLBfdv6GetMsg, ::service_layer::SLBfdv6GetMsgRsp, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SLBfdv6Oper::Service* service,
             ::grpc::ServerContext* ctx,
             const ::service_layer::SLBfdv6GetMsg* req,
             ::service_layer::SLBfdv6GetMsgRsp* resp) {
               return service->SLBfdv6SessionGet(ctx, req, resp);
             }, this)));
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

