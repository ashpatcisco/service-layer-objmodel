# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import sl_interface_pb2 as sl__interface__pb2


class SLInterfaceOperStub(object):
    """@defgroup Interfaces
    @brief Interface service definitions.
    Defines the RPC for getting interface status(es).
    @{
    @addtogroup Interfaces
    @{
    ;
    Interface Registration Operations.

    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SLInterfaceGlobalsRegOp = channel.unary_unary(
                '/service_layer.SLInterfaceOper/SLInterfaceGlobalsRegOp',
                request_serializer=sl__interface__pb2.SLInterfaceGlobalsRegMsg.SerializeToString,
                response_deserializer=sl__interface__pb2.SLInterfaceGlobalsRegMsgRsp.FromString,
                )
        self.SLInterfaceGlobalsGet = channel.unary_unary(
                '/service_layer.SLInterfaceOper/SLInterfaceGlobalsGet',
                request_serializer=sl__interface__pb2.SLInterfaceGlobalsGetMsg.SerializeToString,
                response_deserializer=sl__interface__pb2.SLInterfaceGlobalsGetMsgRsp.FromString,
                )
        self.SLInterfaceGlobalsGetStats = channel.unary_unary(
                '/service_layer.SLInterfaceOper/SLInterfaceGlobalsGetStats',
                request_serializer=sl__interface__pb2.SLInterfaceGlobalsGetMsg.SerializeToString,
                response_deserializer=sl__interface__pb2.SLInterfaceGlobalsGetStatsMsgRsp.FromString,
                )
        self.SLInterfaceGet = channel.unary_unary(
                '/service_layer.SLInterfaceOper/SLInterfaceGet',
                request_serializer=sl__interface__pb2.SLInterfaceGetMsg.SerializeToString,
                response_deserializer=sl__interface__pb2.SLInterfaceGetMsgRsp.FromString,
                )
        self.SLInterfaceGetNotifStream = channel.unary_stream(
                '/service_layer.SLInterfaceOper/SLInterfaceGetNotifStream',
                request_serializer=sl__interface__pb2.SLInterfaceGetNotifMsg.SerializeToString,
                response_deserializer=sl__interface__pb2.SLInterfaceNotif.FromString,
                )
        self.SLInterfaceNotifOp = channel.unary_unary(
                '/service_layer.SLInterfaceOper/SLInterfaceNotifOp',
                request_serializer=sl__interface__pb2.SLInterfaceNotifMsg.SerializeToString,
                response_deserializer=sl__interface__pb2.SLInterfaceNotifMsgRsp.FromString,
                )


class SLInterfaceOperServicer(object):
    """@defgroup Interfaces
    @brief Interface service definitions.
    Defines the RPC for getting interface status(es).
    @{
    @addtogroup Interfaces
    @{
    ;
    Interface Registration Operations.

    """

    def SLInterfaceGlobalsRegOp(self, request, context):
        """SLInterfaceGlobalsRegMsg.Oper = SL_REGOP_REGISTER:
        Global Interface registration.
        A client Must Register BEFORE interfaces can be modified/queried.

        SLInterfaceGlobalsRegMsg.Oper = SL_REGOP_UNREGISTER:
        Global Interface un-registration.
        This call is used to end all interface notifications.
        This call cleans up all interface notifications previously requested.

        SLInterfaceGlobalsRegMsg.Oper = SL_REGOP_EOF:
        Interface End Of File.
        After Registration, the client is expected to send an EOF
        message to convey the end of replay of the client's known objects.
        This is especially useful under certain restart scenarios when the
        client and the server are trying to synchronize their interfaces.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SLInterfaceGlobalsGet(self, request, context):
        """Used to retrieve global Interface info from the server.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SLInterfaceGlobalsGetStats(self, request, context):
        """Used to retrieve global Interface stats from the server.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SLInterfaceGet(self, request, context):
        """Retrieve interface attributes and state.
        This call can be used to "poll" the current state of interfaces.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SLInterfaceGetNotifStream(self, request, context):
        """This call is used to get a stream of interface notifications.
        The caller must maintain the GRPC channel as long as
        there is interest in interface notifications.
        This call can be used to get "push" notifications for interface info.
        It is advised that the caller register for notifications before any
        interfaces are used to avoid any loss of notifications.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SLInterfaceNotifOp(self, request, context):
        """Used to enable/disable event notifications for a certain interface.
        By default, all interface events are disabled. The user must enable
        notifications for the interested interfaces.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SLInterfaceOperServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SLInterfaceGlobalsRegOp': grpc.unary_unary_rpc_method_handler(
                    servicer.SLInterfaceGlobalsRegOp,
                    request_deserializer=sl__interface__pb2.SLInterfaceGlobalsRegMsg.FromString,
                    response_serializer=sl__interface__pb2.SLInterfaceGlobalsRegMsgRsp.SerializeToString,
            ),
            'SLInterfaceGlobalsGet': grpc.unary_unary_rpc_method_handler(
                    servicer.SLInterfaceGlobalsGet,
                    request_deserializer=sl__interface__pb2.SLInterfaceGlobalsGetMsg.FromString,
                    response_serializer=sl__interface__pb2.SLInterfaceGlobalsGetMsgRsp.SerializeToString,
            ),
            'SLInterfaceGlobalsGetStats': grpc.unary_unary_rpc_method_handler(
                    servicer.SLInterfaceGlobalsGetStats,
                    request_deserializer=sl__interface__pb2.SLInterfaceGlobalsGetMsg.FromString,
                    response_serializer=sl__interface__pb2.SLInterfaceGlobalsGetStatsMsgRsp.SerializeToString,
            ),
            'SLInterfaceGet': grpc.unary_unary_rpc_method_handler(
                    servicer.SLInterfaceGet,
                    request_deserializer=sl__interface__pb2.SLInterfaceGetMsg.FromString,
                    response_serializer=sl__interface__pb2.SLInterfaceGetMsgRsp.SerializeToString,
            ),
            'SLInterfaceGetNotifStream': grpc.unary_stream_rpc_method_handler(
                    servicer.SLInterfaceGetNotifStream,
                    request_deserializer=sl__interface__pb2.SLInterfaceGetNotifMsg.FromString,
                    response_serializer=sl__interface__pb2.SLInterfaceNotif.SerializeToString,
            ),
            'SLInterfaceNotifOp': grpc.unary_unary_rpc_method_handler(
                    servicer.SLInterfaceNotifOp,
                    request_deserializer=sl__interface__pb2.SLInterfaceNotifMsg.FromString,
                    response_serializer=sl__interface__pb2.SLInterfaceNotifMsgRsp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'service_layer.SLInterfaceOper', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SLInterfaceOper(object):
    """@defgroup Interfaces
    @brief Interface service definitions.
    Defines the RPC for getting interface status(es).
    @{
    @addtogroup Interfaces
    @{
    ;
    Interface Registration Operations.

    """

    @staticmethod
    def SLInterfaceGlobalsRegOp(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/service_layer.SLInterfaceOper/SLInterfaceGlobalsRegOp',
            sl__interface__pb2.SLInterfaceGlobalsRegMsg.SerializeToString,
            sl__interface__pb2.SLInterfaceGlobalsRegMsgRsp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SLInterfaceGlobalsGet(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/service_layer.SLInterfaceOper/SLInterfaceGlobalsGet',
            sl__interface__pb2.SLInterfaceGlobalsGetMsg.SerializeToString,
            sl__interface__pb2.SLInterfaceGlobalsGetMsgRsp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SLInterfaceGlobalsGetStats(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/service_layer.SLInterfaceOper/SLInterfaceGlobalsGetStats',
            sl__interface__pb2.SLInterfaceGlobalsGetMsg.SerializeToString,
            sl__interface__pb2.SLInterfaceGlobalsGetStatsMsgRsp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SLInterfaceGet(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/service_layer.SLInterfaceOper/SLInterfaceGet',
            sl__interface__pb2.SLInterfaceGetMsg.SerializeToString,
            sl__interface__pb2.SLInterfaceGetMsgRsp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SLInterfaceGetNotifStream(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/service_layer.SLInterfaceOper/SLInterfaceGetNotifStream',
            sl__interface__pb2.SLInterfaceGetNotifMsg.SerializeToString,
            sl__interface__pb2.SLInterfaceNotif.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SLInterfaceNotifOp(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/service_layer.SLInterfaceOper/SLInterfaceNotifOp',
            sl__interface__pb2.SLInterfaceNotifMsg.SerializeToString,
            sl__interface__pb2.SLInterfaceNotifMsgRsp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
