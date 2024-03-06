# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: multiparty/prompt/v1/prompt.proto
# Protobuf Python Version: 4.25.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!multiparty/prompt/v1/prompt.proto\x12\x14multiparty.prompt.v1\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xf4\x01\n\x04Meta\x12\x14\n\x05space\x18\x01 \x01(\tR\x05space\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12 \n\x0b\x64\x65scription\x18\x03 \x01(\tR\x0b\x64\x65scription\x12\x34\n\x07\x63reated\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07\x63reated\x12\x34\n\x07updated\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07updated\x12\x34\n\x07\x65xpires\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07\x65xpires\"\xd0\x02\n\x06Prompt\x12.\n\x04meta\x18\x01 \x01(\x0b\x32\x1a.multiparty.prompt.v1.MetaR\x04meta\x12\x14\n\x05title\x18\x02 \x01(\tR\x05title\x12;\n\x06status\x18\x03 \x01(\x0e\x32#.multiparty.prompt.v1.Prompt.StatusR\x06status\x12&\n\x02in\x18\x04 \x01(\x0b\x32\x16.google.protobuf.ValueR\x02in\x12,\n\x05value\x18\x06 \x01(\x0b\x32\x16.google.protobuf.ValueR\x05value\x12\x35\n\x07replies\x18\x08 \x03(\x0b\x32\x1b.multiparty.prompt.v1.ReplyR\x07replies\"6\n\x06Status\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04OPEN\x10\x01\x12\n\n\x06\x43LOSED\x10\x02\x12\t\n\x05\x45RROR\x10\x03\"\x86\x01\n\x05Reply\x12.\n\x04meta\x18\x01 \x01(\x0b\x32\x1a.multiparty.prompt.v1.MetaR\x04meta\x12\x1f\n\x0bprompt_name\x18\x02 \x01(\tR\npromptName\x12,\n\x05value\x18\x03 \x01(\x0b\x32\x16.google.protobuf.ValueR\x05value\"j\n\x13\x43reatePromptRequest\x12\x1d\n\nspace_name\x18\x01 \x01(\tR\tspaceName\x12\x34\n\x06prompt\x18\x02 \x01(\x0b\x32\x1c.multiparty.prompt.v1.PromptR\x06prompt\"L\n\x14\x43reatePromptResponse\x12\x34\n\x06prompt\x18\x01 \x01(\x0b\x32\x1c.multiparty.prompt.v1.PromptR\x06prompt\"3\n\x12ListPromptsRequest\x12\x1d\n\nspace_name\x18\x01 \x01(\tR\tspaceName\"M\n\x13ListPromptsResponse\x12\x36\n\x07prompts\x18\x01 \x03(\x0b\x32\x1c.multiparty.prompt.v1.PromptR\x07prompts\"E\n\x10GetPromptRequest\x12\x1d\n\nspace_name\x18\x01 \x01(\tR\tspaceName\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\"I\n\x11GetPromptResponse\x12\x34\n\x06prompt\x18\x01 \x01(\x0b\x32\x1c.multiparty.prompt.v1.PromptR\x06prompt\"\x87\x01\n\x12\x43reateReplyRequest\x12\x1d\n\nspace_name\x18\x01 \x01(\tR\tspaceName\x12\x1f\n\x0bprompt_name\x18\x02 \x01(\tR\npromptName\x12\x31\n\x05reply\x18\x03 \x01(\x0b\x32\x1b.multiparty.prompt.v1.ReplyR\x05reply\"H\n\x13\x43reateReplyResponse\x12\x31\n\x05reply\x18\x01 \x01(\x0b\x32\x1b.multiparty.prompt.v1.ReplyR\x05replyB\xdf\x01\n\x18\x63om.multiparty.prompt.v1B\x0bPromptProtoP\x01ZDmultiparty.ai/api/multiparty/prompt/v1/multiparty/prompt/v1;promptv1\xa2\x02\x03MPX\xaa\x02\x14Multiparty.Prompt.V1\xca\x02\x14Multiparty\\Prompt\\V1\xe2\x02 Multiparty\\Prompt\\V1\\GPBMetadata\xea\x02\x16Multiparty::Prompt::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'multiparty.prompt.v1.prompt_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\030com.multiparty.prompt.v1B\013PromptProtoP\001ZDmultiparty.ai/api/multiparty/prompt/v1/multiparty/prompt/v1;promptv1\242\002\003MPX\252\002\024Multiparty.Prompt.V1\312\002\024Multiparty\\Prompt\\V1\342\002 Multiparty\\Prompt\\V1\\GPBMetadata\352\002\026Multiparty::Prompt::V1'
  _globals['_META']._serialized_start=123
  _globals['_META']._serialized_end=367
  _globals['_PROMPT']._serialized_start=370
  _globals['_PROMPT']._serialized_end=706
  _globals['_PROMPT_STATUS']._serialized_start=652
  _globals['_PROMPT_STATUS']._serialized_end=706
  _globals['_REPLY']._serialized_start=709
  _globals['_REPLY']._serialized_end=843
  _globals['_CREATEPROMPTREQUEST']._serialized_start=845
  _globals['_CREATEPROMPTREQUEST']._serialized_end=951
  _globals['_CREATEPROMPTRESPONSE']._serialized_start=953
  _globals['_CREATEPROMPTRESPONSE']._serialized_end=1029
  _globals['_LISTPROMPTSREQUEST']._serialized_start=1031
  _globals['_LISTPROMPTSREQUEST']._serialized_end=1082
  _globals['_LISTPROMPTSRESPONSE']._serialized_start=1084
  _globals['_LISTPROMPTSRESPONSE']._serialized_end=1161
  _globals['_GETPROMPTREQUEST']._serialized_start=1163
  _globals['_GETPROMPTREQUEST']._serialized_end=1232
  _globals['_GETPROMPTRESPONSE']._serialized_start=1234
  _globals['_GETPROMPTRESPONSE']._serialized_end=1307
  _globals['_CREATEREPLYREQUEST']._serialized_start=1310
  _globals['_CREATEREPLYREQUEST']._serialized_end=1445
  _globals['_CREATEREPLYRESPONSE']._serialized_start=1447
  _globals['_CREATEREPLYRESPONSE']._serialized_end=1519
# @@protoc_insertion_point(module_scope)
