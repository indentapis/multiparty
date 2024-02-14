"""Prompt allows external decisions to be introduced into LLM apps."""
from google.protobuf import struct_pb2

from multiparty.api.multiparty.prompt.v1.prompt_pb2 import Prompt # pylint: disable=no-name-in-module

form = struct_pb2.Value() # pylint: disable=no-member
form.struct_value.update({"color": "blue"})
p = Prompt(title="What is your favorite color?", form=form)
print(p)
