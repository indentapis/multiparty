# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: indent/exec/v1/exec.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19indent/exec/v1/exec.proto\x12\x19indent.jsonschema.draft07\x1a\x1cgoogle/protobuf/struct.proto\x1a\x17google/rpc/status.proto\"u\n\nExecutable\x12\x1f\n\nimage_path\x18\x01 \x01(\tH\x00R\timagePath\x12\x1f\n\nimage_body\x18\x02 \x01(\tH\x00R\timageBody\x12\x18\n\x07version\x18\x03 \x01(\tR\x07versionB\x0b\n\tImageBody\"\x94\x01\n\tJsonPatch\x12\x31\n\x02op\x18\x01 \x01(\x0e\x32!.indent.jsonschema.draft07.OpTypeR\x02op\x12\x12\n\x04path\x18\x02 \x01(\tR\x04path\x12,\n\x05value\x18\x03 \x01(\x0b\x32\x16.google.protobuf.ValueR\x05value\x12\x12\n\x04\x66rom\x18\x04 \x01(\tR\x04\x66rom\"E\n\x05\x45vent\x12\x19\n\x08patch_id\x18\x01 \x01(\tR\x07patchId\x12!\n\x0c\x65xec_message\x18\x02 \x01(\tR\x0b\x65xecMessage\"I\n\x0cSetupRequest\x12\x39\n\x04\x65xec\x18\x01 \x01(\x0b\x32%.indent.jsonschema.draft07.ExecutableR\x04\x65xec\";\n\rSetupResponse\x12*\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.StatusR\x06status\"%\n\nCredential\x12\x17\n\x07\x61pi_key\x18\x01 \x01(\tR\x06\x61piKey\"\x14\n\x12\x43redentialResponse\"-\n\x05Query\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n\x05input\x18\x02 \x01(\tR\x05input\"e\n\rQueryResponse\x12*\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.StatusR\x06status\x12\x0e\n\x02id\x18\x02 \x01(\tR\x02id\x12\x18\n\x07results\x18\x03 \x01(\tR\x07results\"T\n\x06\x41\x63tion\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12:\n\x05patch\x18\x02 \x01(\x0b\x32$.indent.jsonschema.draft07.JsonPatchR\x05patch\"\x85\x01\n\x08Response\x12*\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.StatusR\x06status\x12\x1b\n\taction_id\x18\x02 \x01(\tR\x08\x61\x63tionId\x12\x30\n\x07results\x18\x03 \x03(\x0b\x32\x16.google.protobuf.ValueR\x07results\"\x0e\n\x0c\x43loseRequest\"\x0f\n\rCloseResponse*>\n\x06OpType\x12\x07\n\x03\x61\x64\x64\x10\x00\x12\n\n\x06remove\x10\x01\x12\x0b\n\x07replace\x10\x02\x12\x08\n\x04\x63opy\x10\x03\x12\x08\n\x04test\x10\x04\x42\x62\n\x12\x63om.indent.exec.v1B\tExecProtoP\x01Z\x17indent/exec/v1;execv1pb\xa2\x02\x03IJD\xaa\x02\x0eIndent.Exec.v1\xca\x02\x0eIndent\\Exec\\v1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'indent.exec.v1.exec_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\022com.indent.exec.v1B\tExecProtoP\001Z\027indent/exec/v1;execv1pb\242\002\003IJD\252\002\016Indent.Exec.v1\312\002\016Indent\\Exec\\v1'
  _OPTYPE._serialized_start=1054
  _OPTYPE._serialized_end=1116
  _EXECUTABLE._serialized_start=111
  _EXECUTABLE._serialized_end=228
  _JSONPATCH._serialized_start=231
  _JSONPATCH._serialized_end=379
  _EVENT._serialized_start=381
  _EVENT._serialized_end=450
  _SETUPREQUEST._serialized_start=452
  _SETUPREQUEST._serialized_end=525
  _SETUPRESPONSE._serialized_start=527
  _SETUPRESPONSE._serialized_end=586
  _CREDENTIAL._serialized_start=588
  _CREDENTIAL._serialized_end=625
  _CREDENTIALRESPONSE._serialized_start=627
  _CREDENTIALRESPONSE._serialized_end=647
  _QUERY._serialized_start=649
  _QUERY._serialized_end=694
  _QUERYRESPONSE._serialized_start=696
  _QUERYRESPONSE._serialized_end=797
  _ACTION._serialized_start=799
  _ACTION._serialized_end=883
  _RESPONSE._serialized_start=886
  _RESPONSE._serialized_end=1019
  _CLOSEREQUEST._serialized_start=1021
  _CLOSEREQUEST._serialized_end=1035
  _CLOSERESPONSE._serialized_start=1037
  _CLOSERESPONSE._serialized_end=1052
# @@protoc_insertion_point(module_scope)