# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: multiparty/prompt/v1/notify.proto
# Protobuf Python Version: 4.25.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from multiparty.prompt.v1 import prompt_pb2 as multiparty_dot_prompt_dot_v1_dot_prompt__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!multiparty/prompt/v1/notify.proto\x12\x14multiparty.prompt.v1\x1a!multiparty/prompt/v1/prompt.proto\"\xcd\x01\n\x13WebPushSubscription\x12.\n\x04meta\x18\x01 \x01(\x0b\x32\x1a.multiparty.prompt.v1.MetaR\x04meta\x12\x1a\n\x08\x65ndpoint\x18\x02 \x01(\tR\x08\x65ndpoint\x12\'\n\x0f\x65xpiration_time\x18\x03 \x01(\x04R\x0e\x65xpirationTime\x12\x41\n\x04keys\x18\x04 \x01(\x0b\x32-.multiparty.prompt.v1.WebPushSubscriptionKeysR\x04keys\"E\n\x17WebPushSubscriptionKeys\x12\x16\n\x06p256dh\x18\x01 \x01(\tR\x06p256dh\x12\x12\n\x04\x61uth\x18\x02 \x01(\tR\x04\x61uthB\xdf\x01\n\x18\x63om.multiparty.prompt.v1B\x0bNotifyProtoP\x01ZDmultiparty.ai/api/multiparty/prompt/v1/multiparty/prompt/v1;promptv1\xa2\x02\x03MPX\xaa\x02\x14Multiparty.Prompt.V1\xca\x02\x14Multiparty\\Prompt\\V1\xe2\x02 Multiparty\\Prompt\\V1\\GPBMetadata\xea\x02\x16Multiparty::Prompt::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'multiparty.prompt.v1.notify_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\030com.multiparty.prompt.v1B\013NotifyProtoP\001ZDmultiparty.ai/api/multiparty/prompt/v1/multiparty/prompt/v1;promptv1\242\002\003MPX\252\002\024Multiparty.Prompt.V1\312\002\024Multiparty\\Prompt\\V1\342\002 Multiparty\\Prompt\\V1\\GPBMetadata\352\002\026Multiparty::Prompt::V1'
  _globals['_WEBPUSHSUBSCRIPTION']._serialized_start=95
  _globals['_WEBPUSHSUBSCRIPTION']._serialized_end=300
  _globals['_WEBPUSHSUBSCRIPTIONKEYS']._serialized_start=302
  _globals['_WEBPUSHSUBSCRIPTIONKEYS']._serialized_end=371
# @@protoc_insertion_point(module_scope)
