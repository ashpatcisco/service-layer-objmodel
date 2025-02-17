# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: sl_global.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import sl_common_types_pb2 as sl__common__types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='sl_global.proto',
  package='service_layer',
  syntax='proto3',
  serialized_options=b'ZOgithub.com/Cisco-service-layer/service-layer-objmodel/grpc/protos;service_layer',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0fsl_global.proto\x12\rservice_layer\x1a\x15sl_common_types.proto\"?\n\tSLInitMsg\x12\x10\n\x08MajorVer\x18\x01 \x01(\r\x12\x10\n\x08MinorVer\x18\x02 \x01(\r\x12\x0e\n\x06SubVer\x18\x03 \x01(\r\"B\n\x0cSLInitMsgRsp\x12\x10\n\x08MajorVer\x18\x01 \x01(\r\x12\x10\n\x08MinorVer\x18\x02 \x01(\r\x12\x0e\n\x06SubVer\x18\x03 \x01(\r\"-\n\x1aSLVrfRouteReplayErrorNotif\x12\x0f\n\x07VrfName\x18\x01 \x01(\t\"\xfb\x01\n\rSLGlobalNotif\x12\x33\n\tEventType\x18\x01 \x01(\x0e\x32 .service_layer.SLGlobalNotifType\x12/\n\tErrStatus\x18\x02 \x01(\x0b\x32\x1c.service_layer.SLErrorStatus\x12\x31\n\nInitRspMsg\x18\x03 \x01(\x0b\x32\x1b.service_layer.SLInitMsgRspH\x00\x12H\n\x13VrfReplayErrorNotif\x18\x04 \x01(\x0b\x32).service_layer.SLVrfRouteReplayErrorNotifH\x00\x42\x07\n\x05\x45vent\"\x11\n\x0fSLGlobalsGetMsg\"\xe8\x03\n\x12SLGlobalsGetMsgRsp\x12/\n\tErrStatus\x18\x01 \x01(\x0b\x32\x1c.service_layer.SLErrorStatus\x12\x18\n\x10MaxVrfNameLength\x18\x02 \x01(\r\x12\x1e\n\x16MaxInterfaceNameLength\x18\x03 \x01(\r\x12\x18\n\x10MaxPathsPerEntry\x18\x04 \x01(\r\x12\x1e\n\x16MaxPrimaryPathPerEntry\x18\x05 \x01(\r\x12\x1d\n\x15MaxBackupPathPerEntry\x18\x06 \x01(\r\x12\x1c\n\x14MaxMplsLabelsPerPath\x18\x07 \x01(\r\x12\x1b\n\x13MinPrimaryPathIdNum\x18\x08 \x01(\r\x12\x1b\n\x13MaxPrimaryPathIdNum\x18\t \x01(\r\x12\x1a\n\x12MinBackupPathIdNum\x18\n \x01(\r\x12\x1a\n\x12MaxBackupPathIdNum\x18\x0b \x01(\r\x12\x1b\n\x13MaxRemoteAddressNum\x18\x0c \x01(\r\x12\x19\n\x11MaxL2BdNameLength\x18\r \x01(\r\x12\x1f\n\x17MaxL2PmsiTunnelIdLength\x18\x0e \x01(\r\x12%\n\x1dMaxLabelBlockClientNameLength\x18\x0f \x01(\r*\x9c\x01\n\x11SLGlobalNotifType\x12!\n\x1dSL_GLOBAL_EVENT_TYPE_RESERVED\x10\x00\x12\x1e\n\x1aSL_GLOBAL_EVENT_TYPE_ERROR\x10\x01\x12\"\n\x1eSL_GLOBAL_EVENT_TYPE_HEARTBEAT\x10\x02\x12 \n\x1cSL_GLOBAL_EVENT_TYPE_VERSION\x10\x03\x32\xac\x01\n\x08SLGlobal\x12M\n\x11SLGlobalInitNotif\x12\x18.service_layer.SLInitMsg\x1a\x1c.service_layer.SLGlobalNotif0\x01\x12Q\n\x0cSLGlobalsGet\x12\x1e.service_layer.SLGlobalsGetMsg\x1a!.service_layer.SLGlobalsGetMsgRspBQZOgithub.com/Cisco-service-layer/service-layer-objmodel/grpc/protos;service_layerb\x06proto3'
  ,
  dependencies=[sl__common__types__pb2.DESCRIPTOR,])

_SLGLOBALNOTIFTYPE = _descriptor.EnumDescriptor(
  name='SLGlobalNotifType',
  full_name='service_layer.SLGlobalNotifType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='SL_GLOBAL_EVENT_TYPE_RESERVED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SL_GLOBAL_EVENT_TYPE_ERROR', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SL_GLOBAL_EVENT_TYPE_HEARTBEAT', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SL_GLOBAL_EVENT_TYPE_VERSION', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1002,
  serialized_end=1158,
)
_sym_db.RegisterEnumDescriptor(_SLGLOBALNOTIFTYPE)

SLGlobalNotifType = enum_type_wrapper.EnumTypeWrapper(_SLGLOBALNOTIFTYPE)
SL_GLOBAL_EVENT_TYPE_RESERVED = 0
SL_GLOBAL_EVENT_TYPE_ERROR = 1
SL_GLOBAL_EVENT_TYPE_HEARTBEAT = 2
SL_GLOBAL_EVENT_TYPE_VERSION = 3



_SLINITMSG = _descriptor.Descriptor(
  name='SLInitMsg',
  full_name='service_layer.SLInitMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='MajorVer', full_name='service_layer.SLInitMsg.MajorVer', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MinorVer', full_name='service_layer.SLInitMsg.MinorVer', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='SubVer', full_name='service_layer.SLInitMsg.SubVer', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=57,
  serialized_end=120,
)


_SLINITMSGRSP = _descriptor.Descriptor(
  name='SLInitMsgRsp',
  full_name='service_layer.SLInitMsgRsp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='MajorVer', full_name='service_layer.SLInitMsgRsp.MajorVer', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MinorVer', full_name='service_layer.SLInitMsgRsp.MinorVer', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='SubVer', full_name='service_layer.SLInitMsgRsp.SubVer', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=122,
  serialized_end=188,
)


_SLVRFROUTEREPLAYERRORNOTIF = _descriptor.Descriptor(
  name='SLVrfRouteReplayErrorNotif',
  full_name='service_layer.SLVrfRouteReplayErrorNotif',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='VrfName', full_name='service_layer.SLVrfRouteReplayErrorNotif.VrfName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=190,
  serialized_end=235,
)


_SLGLOBALNOTIF = _descriptor.Descriptor(
  name='SLGlobalNotif',
  full_name='service_layer.SLGlobalNotif',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='EventType', full_name='service_layer.SLGlobalNotif.EventType', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ErrStatus', full_name='service_layer.SLGlobalNotif.ErrStatus', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='InitRspMsg', full_name='service_layer.SLGlobalNotif.InitRspMsg', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='VrfReplayErrorNotif', full_name='service_layer.SLGlobalNotif.VrfReplayErrorNotif', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='Event', full_name='service_layer.SLGlobalNotif.Event',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=238,
  serialized_end=489,
)


_SLGLOBALSGETMSG = _descriptor.Descriptor(
  name='SLGlobalsGetMsg',
  full_name='service_layer.SLGlobalsGetMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=491,
  serialized_end=508,
)


_SLGLOBALSGETMSGRSP = _descriptor.Descriptor(
  name='SLGlobalsGetMsgRsp',
  full_name='service_layer.SLGlobalsGetMsgRsp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='ErrStatus', full_name='service_layer.SLGlobalsGetMsgRsp.ErrStatus', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxVrfNameLength', full_name='service_layer.SLGlobalsGetMsgRsp.MaxVrfNameLength', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxInterfaceNameLength', full_name='service_layer.SLGlobalsGetMsgRsp.MaxInterfaceNameLength', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxPathsPerEntry', full_name='service_layer.SLGlobalsGetMsgRsp.MaxPathsPerEntry', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxPrimaryPathPerEntry', full_name='service_layer.SLGlobalsGetMsgRsp.MaxPrimaryPathPerEntry', index=4,
      number=5, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxBackupPathPerEntry', full_name='service_layer.SLGlobalsGetMsgRsp.MaxBackupPathPerEntry', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxMplsLabelsPerPath', full_name='service_layer.SLGlobalsGetMsgRsp.MaxMplsLabelsPerPath', index=6,
      number=7, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MinPrimaryPathIdNum', full_name='service_layer.SLGlobalsGetMsgRsp.MinPrimaryPathIdNum', index=7,
      number=8, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxPrimaryPathIdNum', full_name='service_layer.SLGlobalsGetMsgRsp.MaxPrimaryPathIdNum', index=8,
      number=9, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MinBackupPathIdNum', full_name='service_layer.SLGlobalsGetMsgRsp.MinBackupPathIdNum', index=9,
      number=10, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxBackupPathIdNum', full_name='service_layer.SLGlobalsGetMsgRsp.MaxBackupPathIdNum', index=10,
      number=11, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxRemoteAddressNum', full_name='service_layer.SLGlobalsGetMsgRsp.MaxRemoteAddressNum', index=11,
      number=12, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxL2BdNameLength', full_name='service_layer.SLGlobalsGetMsgRsp.MaxL2BdNameLength', index=12,
      number=13, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxL2PmsiTunnelIdLength', full_name='service_layer.SLGlobalsGetMsgRsp.MaxL2PmsiTunnelIdLength', index=13,
      number=14, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='MaxLabelBlockClientNameLength', full_name='service_layer.SLGlobalsGetMsgRsp.MaxLabelBlockClientNameLength', index=14,
      number=15, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=511,
  serialized_end=999,
)

_SLGLOBALNOTIF.fields_by_name['EventType'].enum_type = _SLGLOBALNOTIFTYPE
_SLGLOBALNOTIF.fields_by_name['ErrStatus'].message_type = sl__common__types__pb2._SLERRORSTATUS
_SLGLOBALNOTIF.fields_by_name['InitRspMsg'].message_type = _SLINITMSGRSP
_SLGLOBALNOTIF.fields_by_name['VrfReplayErrorNotif'].message_type = _SLVRFROUTEREPLAYERRORNOTIF
_SLGLOBALNOTIF.oneofs_by_name['Event'].fields.append(
  _SLGLOBALNOTIF.fields_by_name['InitRspMsg'])
_SLGLOBALNOTIF.fields_by_name['InitRspMsg'].containing_oneof = _SLGLOBALNOTIF.oneofs_by_name['Event']
_SLGLOBALNOTIF.oneofs_by_name['Event'].fields.append(
  _SLGLOBALNOTIF.fields_by_name['VrfReplayErrorNotif'])
_SLGLOBALNOTIF.fields_by_name['VrfReplayErrorNotif'].containing_oneof = _SLGLOBALNOTIF.oneofs_by_name['Event']
_SLGLOBALSGETMSGRSP.fields_by_name['ErrStatus'].message_type = sl__common__types__pb2._SLERRORSTATUS
DESCRIPTOR.message_types_by_name['SLInitMsg'] = _SLINITMSG
DESCRIPTOR.message_types_by_name['SLInitMsgRsp'] = _SLINITMSGRSP
DESCRIPTOR.message_types_by_name['SLVrfRouteReplayErrorNotif'] = _SLVRFROUTEREPLAYERRORNOTIF
DESCRIPTOR.message_types_by_name['SLGlobalNotif'] = _SLGLOBALNOTIF
DESCRIPTOR.message_types_by_name['SLGlobalsGetMsg'] = _SLGLOBALSGETMSG
DESCRIPTOR.message_types_by_name['SLGlobalsGetMsgRsp'] = _SLGLOBALSGETMSGRSP
DESCRIPTOR.enum_types_by_name['SLGlobalNotifType'] = _SLGLOBALNOTIFTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SLInitMsg = _reflection.GeneratedProtocolMessageType('SLInitMsg', (_message.Message,), {
  'DESCRIPTOR' : _SLINITMSG,
  '__module__' : 'sl_global_pb2'
  # @@protoc_insertion_point(class_scope:service_layer.SLInitMsg)
  })
_sym_db.RegisterMessage(SLInitMsg)

SLInitMsgRsp = _reflection.GeneratedProtocolMessageType('SLInitMsgRsp', (_message.Message,), {
  'DESCRIPTOR' : _SLINITMSGRSP,
  '__module__' : 'sl_global_pb2'
  # @@protoc_insertion_point(class_scope:service_layer.SLInitMsgRsp)
  })
_sym_db.RegisterMessage(SLInitMsgRsp)

SLVrfRouteReplayErrorNotif = _reflection.GeneratedProtocolMessageType('SLVrfRouteReplayErrorNotif', (_message.Message,), {
  'DESCRIPTOR' : _SLVRFROUTEREPLAYERRORNOTIF,
  '__module__' : 'sl_global_pb2'
  # @@protoc_insertion_point(class_scope:service_layer.SLVrfRouteReplayErrorNotif)
  })
_sym_db.RegisterMessage(SLVrfRouteReplayErrorNotif)

SLGlobalNotif = _reflection.GeneratedProtocolMessageType('SLGlobalNotif', (_message.Message,), {
  'DESCRIPTOR' : _SLGLOBALNOTIF,
  '__module__' : 'sl_global_pb2'
  # @@protoc_insertion_point(class_scope:service_layer.SLGlobalNotif)
  })
_sym_db.RegisterMessage(SLGlobalNotif)

SLGlobalsGetMsg = _reflection.GeneratedProtocolMessageType('SLGlobalsGetMsg', (_message.Message,), {
  'DESCRIPTOR' : _SLGLOBALSGETMSG,
  '__module__' : 'sl_global_pb2'
  # @@protoc_insertion_point(class_scope:service_layer.SLGlobalsGetMsg)
  })
_sym_db.RegisterMessage(SLGlobalsGetMsg)

SLGlobalsGetMsgRsp = _reflection.GeneratedProtocolMessageType('SLGlobalsGetMsgRsp', (_message.Message,), {
  'DESCRIPTOR' : _SLGLOBALSGETMSGRSP,
  '__module__' : 'sl_global_pb2'
  # @@protoc_insertion_point(class_scope:service_layer.SLGlobalsGetMsgRsp)
  })
_sym_db.RegisterMessage(SLGlobalsGetMsgRsp)


DESCRIPTOR._options = None

_SLGLOBAL = _descriptor.ServiceDescriptor(
  name='SLGlobal',
  full_name='service_layer.SLGlobal',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1161,
  serialized_end=1333,
  methods=[
  _descriptor.MethodDescriptor(
    name='SLGlobalInitNotif',
    full_name='service_layer.SLGlobal.SLGlobalInitNotif',
    index=0,
    containing_service=None,
    input_type=_SLINITMSG,
    output_type=_SLGLOBALNOTIF,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='SLGlobalsGet',
    full_name='service_layer.SLGlobal.SLGlobalsGet',
    index=1,
    containing_service=None,
    input_type=_SLGLOBALSGETMSG,
    output_type=_SLGLOBALSGETMSGRSP,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_SLGLOBAL)

DESCRIPTOR.services_by_name['SLGlobal'] = _SLGLOBAL

# @@protoc_insertion_point(module_scope)
