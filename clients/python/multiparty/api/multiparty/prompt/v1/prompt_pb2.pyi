"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
Copyright 2024 Indent Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.internal.enum_type_wrapper
import google.protobuf.message
import google.protobuf.struct_pb2
import google.protobuf.timestamp_pb2
import sys
import typing

if sys.version_info >= (3, 10):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

@typing_extensions.final
class Meta(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SPACE_FIELD_NUMBER: builtins.int
    NAME_FIELD_NUMBER: builtins.int
    DESCRIPTION_FIELD_NUMBER: builtins.int
    CREATED_FIELD_NUMBER: builtins.int
    UPDATED_FIELD_NUMBER: builtins.int
    EXPIRES_FIELD_NUMBER: builtins.int
    space: builtins.str
    name: builtins.str
    description: builtins.str
    @property
    def created(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    @property
    def updated(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    @property
    def expires(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    def __init__(
        self,
        *,
        space: builtins.str = ...,
        name: builtins.str = ...,
        description: builtins.str = ...,
        created: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        updated: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        expires: google.protobuf.timestamp_pb2.Timestamp | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["created", b"created", "expires", b"expires", "updated", b"updated"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["created", b"created", "description", b"description", "expires", b"expires", "name", b"name", "space", b"space", "updated", b"updated"]) -> None: ...

global___Meta = Meta

@typing_extensions.final
class Prompt(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class _Status:
        ValueType = typing.NewType("ValueType", builtins.int)
        V: typing_extensions.TypeAlias = ValueType

    class _StatusEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[Prompt._Status.ValueType], builtins.type):
        DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
        UNKNOWN: Prompt._Status.ValueType  # 0
        OPEN: Prompt._Status.ValueType  # 1
        CLOSED: Prompt._Status.ValueType  # 2
        ERROR: Prompt._Status.ValueType  # 3

    class Status(_Status, metaclass=_StatusEnumTypeWrapper): ...
    UNKNOWN: Prompt.Status.ValueType  # 0
    OPEN: Prompt.Status.ValueType  # 1
    CLOSED: Prompt.Status.ValueType  # 2
    ERROR: Prompt.Status.ValueType  # 3

    META_FIELD_NUMBER: builtins.int
    TITLE_FIELD_NUMBER: builtins.int
    STATUS_FIELD_NUMBER: builtins.int
    IN_FIELD_NUMBER: builtins.int
    VALUE_FIELD_NUMBER: builtins.int
    REPLIES_FIELD_NUMBER: builtins.int
    @property
    def meta(self) -> global___Meta: ...
    title: builtins.str
    status: global___Prompt.Status.ValueType
    @property
    def value(self) -> google.protobuf.struct_pb2.Value:
        """jsonSchema = 5;"""
    @property
    def replies(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___Reply]:
        """uiSchema = 7;"""
    def __init__(
        self,
        *,
        meta: global___Meta | None = ...,
        title: builtins.str = ...,
        status: global___Prompt.Status.ValueType = ...,
        value: google.protobuf.struct_pb2.Value | None = ...,
        replies: collections.abc.Iterable[global___Reply] | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["in", b"in", "meta", b"meta", "value", b"value"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["in", b"in", "meta", b"meta", "replies", b"replies", "status", b"status", "title", b"title", "value", b"value"]) -> None: ...

global___Prompt = Prompt

@typing_extensions.final
class Reply(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    META_FIELD_NUMBER: builtins.int
    PROMPT_NAME_FIELD_NUMBER: builtins.int
    VALUE_FIELD_NUMBER: builtins.int
    @property
    def meta(self) -> global___Meta: ...
    prompt_name: builtins.str
    @property
    def value(self) -> google.protobuf.struct_pb2.Value:
        """value matches the schema of the prompt"""
    def __init__(
        self,
        *,
        meta: global___Meta | None = ...,
        prompt_name: builtins.str = ...,
        value: google.protobuf.struct_pb2.Value | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["meta", b"meta", "value", b"value"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["meta", b"meta", "prompt_name", b"prompt_name", "value", b"value"]) -> None: ...

global___Reply = Reply

@typing_extensions.final
class CreatePromptRequest(google.protobuf.message.Message):
    """Prompts"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SPACE_NAME_FIELD_NUMBER: builtins.int
    PROMPT_FIELD_NUMBER: builtins.int
    space_name: builtins.str
    @property
    def prompt(self) -> global___Prompt: ...
    def __init__(
        self,
        *,
        space_name: builtins.str = ...,
        prompt: global___Prompt | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["prompt", b"prompt"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["prompt", b"prompt", "space_name", b"space_name"]) -> None: ...

global___CreatePromptRequest = CreatePromptRequest

@typing_extensions.final
class CreatePromptResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PROMPT_FIELD_NUMBER: builtins.int
    @property
    def prompt(self) -> global___Prompt: ...
    def __init__(
        self,
        *,
        prompt: global___Prompt | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["prompt", b"prompt"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["prompt", b"prompt"]) -> None: ...

global___CreatePromptResponse = CreatePromptResponse

@typing_extensions.final
class ListPromptsRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SPACE_NAME_FIELD_NUMBER: builtins.int
    space_name: builtins.str
    def __init__(
        self,
        *,
        space_name: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["space_name", b"space_name"]) -> None: ...

global___ListPromptsRequest = ListPromptsRequest

@typing_extensions.final
class ListPromptsResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PROMPTS_FIELD_NUMBER: builtins.int
    @property
    def prompts(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___Prompt]: ...
    def __init__(
        self,
        *,
        prompts: collections.abc.Iterable[global___Prompt] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["prompts", b"prompts"]) -> None: ...

global___ListPromptsResponse = ListPromptsResponse

@typing_extensions.final
class GetPromptRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SPACE_NAME_FIELD_NUMBER: builtins.int
    NAME_FIELD_NUMBER: builtins.int
    space_name: builtins.str
    name: builtins.str
    def __init__(
        self,
        *,
        space_name: builtins.str = ...,
        name: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["name", b"name", "space_name", b"space_name"]) -> None: ...

global___GetPromptRequest = GetPromptRequest

@typing_extensions.final
class GetPromptResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PROMPT_FIELD_NUMBER: builtins.int
    @property
    def prompt(self) -> global___Prompt: ...
    def __init__(
        self,
        *,
        prompt: global___Prompt | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["prompt", b"prompt"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["prompt", b"prompt"]) -> None: ...

global___GetPromptResponse = GetPromptResponse

@typing_extensions.final
class CreateReplyRequest(google.protobuf.message.Message):
    """Replies"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SPACE_NAME_FIELD_NUMBER: builtins.int
    PROMPT_NAME_FIELD_NUMBER: builtins.int
    REPLY_FIELD_NUMBER: builtins.int
    space_name: builtins.str
    prompt_name: builtins.str
    @property
    def reply(self) -> global___Reply: ...
    def __init__(
        self,
        *,
        space_name: builtins.str = ...,
        prompt_name: builtins.str = ...,
        reply: global___Reply | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["reply", b"reply"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["prompt_name", b"prompt_name", "reply", b"reply", "space_name", b"space_name"]) -> None: ...

global___CreateReplyRequest = CreateReplyRequest

@typing_extensions.final
class CreateReplyResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    REPLY_FIELD_NUMBER: builtins.int
    @property
    def reply(self) -> global___Reply: ...
    def __init__(
        self,
        *,
        reply: global___Reply | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["reply", b"reply"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["reply", b"reply"]) -> None: ...

global___CreateReplyResponse = CreateReplyResponse
