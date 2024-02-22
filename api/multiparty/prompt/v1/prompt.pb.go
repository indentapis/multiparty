// Copyright 2024 Indent Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: multiparty/prompt/v1/prompt.proto

package promptv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Prompt_Status int32

const (
	Prompt_UNKNOWN Prompt_Status = 0
	Prompt_OPEN    Prompt_Status = 1
	Prompt_CLOSED  Prompt_Status = 2
	Prompt_ERROR   Prompt_Status = 3
)

// Enum value maps for Prompt_Status.
var (
	Prompt_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "OPEN",
		2: "CLOSED",
		3: "ERROR",
	}
	Prompt_Status_value = map[string]int32{
		"UNKNOWN": 0,
		"OPEN":    1,
		"CLOSED":  2,
		"ERROR":   3,
	}
)

func (x Prompt_Status) Enum() *Prompt_Status {
	p := new(Prompt_Status)
	*p = x
	return p
}

func (x Prompt_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Prompt_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_multiparty_prompt_v1_prompt_proto_enumTypes[0].Descriptor()
}

func (Prompt_Status) Type() protoreflect.EnumType {
	return &file_multiparty_prompt_v1_prompt_proto_enumTypes[0]
}

func (x Prompt_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Prompt_Status.Descriptor instead.
func (Prompt_Status) EnumDescriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{1, 0}
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Space       string                 `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Created     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *Meta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Meta) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Meta) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Meta) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

type Prompt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *Meta           `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Title  string          `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Status Prompt_Status   `protobuf:"varint,3,opt,name=status,proto3,enum=multiparty.prompt.v1.Prompt_Status" json:"status,omitempty"`
	In     *structpb.Value `protobuf:"bytes,4,opt,name=in,proto3" json:"in,omitempty"`
	// jsonSchema = 5;
	Value *structpb.Value `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	// uiSchema = 7;
	Replies []*Reply `protobuf:"bytes,8,rep,name=replies,proto3" json:"replies,omitempty"`
}

func (x *Prompt) Reset() {
	*x = Prompt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prompt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prompt) ProtoMessage() {}

func (x *Prompt) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prompt.ProtoReflect.Descriptor instead.
func (*Prompt) Descriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{1}
}

func (x *Prompt) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Prompt) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Prompt) GetStatus() Prompt_Status {
	if x != nil {
		return x.Status
	}
	return Prompt_UNKNOWN
}

func (x *Prompt) GetIn() *structpb.Value {
	if x != nil {
		return x.In
	}
	return nil
}

func (x *Prompt) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Prompt) GetReplies() []*Reply {
	if x != nil {
		return x.Replies
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta       *Meta  `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	PromptName string `protobuf:"bytes,2,opt,name=prompt_name,json=promptName,proto3" json:"prompt_name,omitempty"`
	// value matches the schema of the prompt
	Value *structpb.Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{2}
}

func (x *Reply) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Reply) GetPromptName() string {
	if x != nil {
		return x.PromptName
	}
	return ""
}

func (x *Reply) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

// Prompts
type CreatePromptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceName string  `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	Prompt    *Prompt `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
}

func (x *CreatePromptRequest) Reset() {
	*x = CreatePromptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePromptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePromptRequest) ProtoMessage() {}

func (x *CreatePromptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePromptRequest.ProtoReflect.Descriptor instead.
func (*CreatePromptRequest) Descriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePromptRequest) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

func (x *CreatePromptRequest) GetPrompt() *Prompt {
	if x != nil {
		return x.Prompt
	}
	return nil
}

type CreatePromptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prompt *Prompt `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
}

func (x *CreatePromptResponse) Reset() {
	*x = CreatePromptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePromptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePromptResponse) ProtoMessage() {}

func (x *CreatePromptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePromptResponse.ProtoReflect.Descriptor instead.
func (*CreatePromptResponse) Descriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePromptResponse) GetPrompt() *Prompt {
	if x != nil {
		return x.Prompt
	}
	return nil
}

type ListPromptsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
}

func (x *ListPromptsRequest) Reset() {
	*x = ListPromptsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPromptsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPromptsRequest) ProtoMessage() {}

func (x *ListPromptsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPromptsRequest.ProtoReflect.Descriptor instead.
func (*ListPromptsRequest) Descriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{5}
}

func (x *ListPromptsRequest) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

type ListPromptsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prompts []*Prompt `protobuf:"bytes,1,rep,name=prompts,proto3" json:"prompts,omitempty"`
}

func (x *ListPromptsResponse) Reset() {
	*x = ListPromptsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPromptsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPromptsResponse) ProtoMessage() {}

func (x *ListPromptsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPromptsResponse.ProtoReflect.Descriptor instead.
func (*ListPromptsResponse) Descriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{6}
}

func (x *ListPromptsResponse) GetPrompts() []*Prompt {
	if x != nil {
		return x.Prompts
	}
	return nil
}

type GetPromptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPromptRequest) Reset() {
	*x = GetPromptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPromptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPromptRequest) ProtoMessage() {}

func (x *GetPromptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPromptRequest.ProtoReflect.Descriptor instead.
func (*GetPromptRequest) Descriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{7}
}

func (x *GetPromptRequest) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

func (x *GetPromptRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPromptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prompt *Prompt `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
}

func (x *GetPromptResponse) Reset() {
	*x = GetPromptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPromptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPromptResponse) ProtoMessage() {}

func (x *GetPromptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPromptResponse.ProtoReflect.Descriptor instead.
func (*GetPromptResponse) Descriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{8}
}

func (x *GetPromptResponse) GetPrompt() *Prompt {
	if x != nil {
		return x.Prompt
	}
	return nil
}

// Replies
type CreateReplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceName  string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	PromptName string `protobuf:"bytes,2,opt,name=prompt_name,json=promptName,proto3" json:"prompt_name,omitempty"`
	Reply      *Reply `protobuf:"bytes,3,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *CreateReplyRequest) Reset() {
	*x = CreateReplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReplyRequest) ProtoMessage() {}

func (x *CreateReplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReplyRequest.ProtoReflect.Descriptor instead.
func (*CreateReplyRequest) Descriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{9}
}

func (x *CreateReplyRequest) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

func (x *CreateReplyRequest) GetPromptName() string {
	if x != nil {
		return x.PromptName
	}
	return ""
}

func (x *CreateReplyRequest) GetReply() *Reply {
	if x != nil {
		return x.Reply
	}
	return nil
}

type CreateReplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply *Reply `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *CreateReplyResponse) Reset() {
	*x = CreateReplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReplyResponse) ProtoMessage() {}

func (x *CreateReplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_prompt_v1_prompt_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReplyResponse.ProtoReflect.Descriptor instead.
func (*CreateReplyResponse) Descriptor() ([]byte, []int) {
	return file_multiparty_prompt_v1_prompt_proto_rawDescGZIP(), []int{10}
}

func (x *CreateReplyResponse) GetReply() *Reply {
	if x != nil {
		return x.Reply
	}
	return nil
}

var File_multiparty_prompt_v1_prompt_proto protoreflect.FileDescriptor

var file_multiparty_prompt_v1_prompt_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0xd0, 0x02, 0x0a, 0x06, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x6e, 0x12,
	0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x35, 0x0a,
	0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4f,
	0x50, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x22, 0x86, 0x01, 0x0a,
	0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x22, 0x4c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x22,
	0x33, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x73, 0x22, 0x45, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x06, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x05,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x48, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x42, 0xdf, 0x01, 0x0a, 0x18, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f,
	0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x50,
	0x58, 0xaa, 0x02, 0x14, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x14, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x5c, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x20, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5c, 0x50, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x16, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x3a,
	0x3a, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_multiparty_prompt_v1_prompt_proto_rawDescOnce sync.Once
	file_multiparty_prompt_v1_prompt_proto_rawDescData = file_multiparty_prompt_v1_prompt_proto_rawDesc
)

func file_multiparty_prompt_v1_prompt_proto_rawDescGZIP() []byte {
	file_multiparty_prompt_v1_prompt_proto_rawDescOnce.Do(func() {
		file_multiparty_prompt_v1_prompt_proto_rawDescData = protoimpl.X.CompressGZIP(file_multiparty_prompt_v1_prompt_proto_rawDescData)
	})
	return file_multiparty_prompt_v1_prompt_proto_rawDescData
}

var file_multiparty_prompt_v1_prompt_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_multiparty_prompt_v1_prompt_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_multiparty_prompt_v1_prompt_proto_goTypes = []interface{}{
	(Prompt_Status)(0),            // 0: multiparty.prompt.v1.Prompt.Status
	(*Meta)(nil),                  // 1: multiparty.prompt.v1.Meta
	(*Prompt)(nil),                // 2: multiparty.prompt.v1.Prompt
	(*Reply)(nil),                 // 3: multiparty.prompt.v1.Reply
	(*CreatePromptRequest)(nil),   // 4: multiparty.prompt.v1.CreatePromptRequest
	(*CreatePromptResponse)(nil),  // 5: multiparty.prompt.v1.CreatePromptResponse
	(*ListPromptsRequest)(nil),    // 6: multiparty.prompt.v1.ListPromptsRequest
	(*ListPromptsResponse)(nil),   // 7: multiparty.prompt.v1.ListPromptsResponse
	(*GetPromptRequest)(nil),      // 8: multiparty.prompt.v1.GetPromptRequest
	(*GetPromptResponse)(nil),     // 9: multiparty.prompt.v1.GetPromptResponse
	(*CreateReplyRequest)(nil),    // 10: multiparty.prompt.v1.CreateReplyRequest
	(*CreateReplyResponse)(nil),   // 11: multiparty.prompt.v1.CreateReplyResponse
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
	(*structpb.Value)(nil),        // 13: google.protobuf.Value
}
var file_multiparty_prompt_v1_prompt_proto_depIdxs = []int32{
	12, // 0: multiparty.prompt.v1.Meta.created:type_name -> google.protobuf.Timestamp
	12, // 1: multiparty.prompt.v1.Meta.updated:type_name -> google.protobuf.Timestamp
	1,  // 2: multiparty.prompt.v1.Prompt.meta:type_name -> multiparty.prompt.v1.Meta
	0,  // 3: multiparty.prompt.v1.Prompt.status:type_name -> multiparty.prompt.v1.Prompt.Status
	13, // 4: multiparty.prompt.v1.Prompt.in:type_name -> google.protobuf.Value
	13, // 5: multiparty.prompt.v1.Prompt.value:type_name -> google.protobuf.Value
	3,  // 6: multiparty.prompt.v1.Prompt.replies:type_name -> multiparty.prompt.v1.Reply
	1,  // 7: multiparty.prompt.v1.Reply.meta:type_name -> multiparty.prompt.v1.Meta
	13, // 8: multiparty.prompt.v1.Reply.value:type_name -> google.protobuf.Value
	2,  // 9: multiparty.prompt.v1.CreatePromptRequest.prompt:type_name -> multiparty.prompt.v1.Prompt
	2,  // 10: multiparty.prompt.v1.CreatePromptResponse.prompt:type_name -> multiparty.prompt.v1.Prompt
	2,  // 11: multiparty.prompt.v1.ListPromptsResponse.prompts:type_name -> multiparty.prompt.v1.Prompt
	2,  // 12: multiparty.prompt.v1.GetPromptResponse.prompt:type_name -> multiparty.prompt.v1.Prompt
	3,  // 13: multiparty.prompt.v1.CreateReplyRequest.reply:type_name -> multiparty.prompt.v1.Reply
	3,  // 14: multiparty.prompt.v1.CreateReplyResponse.reply:type_name -> multiparty.prompt.v1.Reply
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_multiparty_prompt_v1_prompt_proto_init() }
func file_multiparty_prompt_v1_prompt_proto_init() {
	if File_multiparty_prompt_v1_prompt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_multiparty_prompt_v1_prompt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_multiparty_prompt_v1_prompt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prompt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_multiparty_prompt_v1_prompt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_multiparty_prompt_v1_prompt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePromptRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_multiparty_prompt_v1_prompt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePromptResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_multiparty_prompt_v1_prompt_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPromptsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_multiparty_prompt_v1_prompt_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPromptsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_multiparty_prompt_v1_prompt_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPromptRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_multiparty_prompt_v1_prompt_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPromptResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_multiparty_prompt_v1_prompt_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReplyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_multiparty_prompt_v1_prompt_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReplyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_multiparty_prompt_v1_prompt_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_multiparty_prompt_v1_prompt_proto_goTypes,
		DependencyIndexes: file_multiparty_prompt_v1_prompt_proto_depIdxs,
		EnumInfos:         file_multiparty_prompt_v1_prompt_proto_enumTypes,
		MessageInfos:      file_multiparty_prompt_v1_prompt_proto_msgTypes,
	}.Build()
	File_multiparty_prompt_v1_prompt_proto = out.File
	file_multiparty_prompt_v1_prompt_proto_rawDesc = nil
	file_multiparty_prompt_v1_prompt_proto_goTypes = nil
	file_multiparty_prompt_v1_prompt_proto_depIdxs = nil
}
