// Copyright 2023 Indent Inc
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
// source: multiparty/v1/rule.proto

package multipartyv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Rule defines a set of Requirements that must be met.
type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// On sets when Rules apply.
	On *On `protobuf:"bytes,1,opt,name=on,proto3" json:"on,omitempty"`
	// Require defines what's needed to satisfy a Rule.
	Require []*Requirement `protobuf:"bytes,2,rep,name=require,proto3" json:"require,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_v1_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_v1_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_multiparty_v1_rule_proto_rawDescGZIP(), []int{0}
}

func (x *Rule) GetOn() *On {
	if x != nil {
		return x.On
	}
	return nil
}

func (x *Rule) GetRequire() []*Requirement {
	if x != nil {
		return x.Require
	}
	return nil
}

// On determines when a Rule is applied.
type On struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Events the Rule applies to.
	Events []string `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	// Match must be subset of event labels for Rule to apply.
	Match map[string]string `protobuf:"bytes,2,rep,name=match,proto3" json:"match,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *On) Reset() {
	*x = On{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_v1_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *On) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*On) ProtoMessage() {}

func (x *On) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_v1_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use On.ProtoReflect.Descriptor instead.
func (*On) Descriptor() ([]byte, []int) {
	return file_multiparty_v1_rule_proto_rawDescGZIP(), []int{1}
}

func (x *On) GetEvents() []string {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *On) GetMatch() map[string]string {
	if x != nil {
		return x.Match
	}
	return nil
}

// Requirement to satisfying a Rule.
type Requirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ensure specifies what meets the Requirement.
	Ensure string `protobuf:"bytes,1,opt,name=ensure,proto3" json:"ensure,omitempty"`
	// Title is a human readable name for the Requirement.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Description of the purpose of Rule.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Origin is the location of the Requirement.
	Origin *Origin `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`
	// Compiled version of Ensure set by Multiparty.
	Compiled *v1alpha1.CheckedExpr `protobuf:"bytes,5,opt,name=compiled,proto3" json:"compiled,omitempty"`
}

func (x *Requirement) Reset() {
	*x = Requirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_v1_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Requirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Requirement) ProtoMessage() {}

func (x *Requirement) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_v1_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Requirement.ProtoReflect.Descriptor instead.
func (*Requirement) Descriptor() ([]byte, []int) {
	return file_multiparty_v1_rule_proto_rawDescGZIP(), []int{2}
}

func (x *Requirement) GetEnsure() string {
	if x != nil {
		return x.Ensure
	}
	return ""
}

func (x *Requirement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Requirement) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Requirement) GetOrigin() *Origin {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *Requirement) GetCompiled() *v1alpha1.CheckedExpr {
	if x != nil {
		return x.Compiled
	}
	return nil
}

// Origin is the location of a Requirement's source.
type Origin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the source.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Line number of the source.
	Line uint32 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
}

func (x *Origin) Reset() {
	*x = Origin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_v1_rule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Origin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Origin) ProtoMessage() {}

func (x *Origin) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_v1_rule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Origin.ProtoReflect.Descriptor instead.
func (*Origin) Descriptor() ([]byte, []int) {
	return file_multiparty_v1_rule_proto_rawDescGZIP(), []int{3}
}

func (x *Origin) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Origin) GetLine() uint32 {
	if x != nil {
		return x.Line
	}
	return 0
}

// CreateRuleRequest is the request to create a Rule.
type CreateRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Space to create the Rule in.
	Space string `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`
	// Rule to create.
	Rule *Rule `protobuf:"bytes,2,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *CreateRuleRequest) Reset() {
	*x = CreateRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_v1_rule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRuleRequest) ProtoMessage() {}

func (x *CreateRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_v1_rule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRuleRequest.ProtoReflect.Descriptor instead.
func (*CreateRuleRequest) Descriptor() ([]byte, []int) {
	return file_multiparty_v1_rule_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRuleRequest) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *CreateRuleRequest) GetRule() *Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

// UpdateRuleRequest is the request to update a Rule.
type UpdateRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Space to update Rule.
	Space string `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`
	// Name of Rule to update.
	RuleName string `protobuf:"bytes,2,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
	// Rule to update.
	Rule *Rule `protobuf:"bytes,3,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *UpdateRuleRequest) Reset() {
	*x = UpdateRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_v1_rule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRuleRequest) ProtoMessage() {}

func (x *UpdateRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_v1_rule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRuleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRuleRequest) Descriptor() ([]byte, []int) {
	return file_multiparty_v1_rule_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRuleRequest) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *UpdateRuleRequest) GetRuleName() string {
	if x != nil {
		return x.RuleName
	}
	return ""
}

func (x *UpdateRuleRequest) GetRule() *Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

// GetRuleRequest is the request to get a Rule.
type GetRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Space to get Rule from.
	Space string `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`
	// Name of Rule to get.
	RuleName string `protobuf:"bytes,2,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
}

func (x *GetRuleRequest) Reset() {
	*x = GetRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_v1_rule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleRequest) ProtoMessage() {}

func (x *GetRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_v1_rule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleRequest.ProtoReflect.Descriptor instead.
func (*GetRuleRequest) Descriptor() ([]byte, []int) {
	return file_multiparty_v1_rule_proto_rawDescGZIP(), []int{6}
}

func (x *GetRuleRequest) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *GetRuleRequest) GetRuleName() string {
	if x != nil {
		return x.RuleName
	}
	return ""
}

// ListRulesRequest is the request to list Rules.
type ListRulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Space to list Rules from.
	Space string `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`
}

func (x *ListRulesRequest) Reset() {
	*x = ListRulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_v1_rule_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRulesRequest) ProtoMessage() {}

func (x *ListRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_v1_rule_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRulesRequest.ProtoReflect.Descriptor instead.
func (*ListRulesRequest) Descriptor() ([]byte, []int) {
	return file_multiparty_v1_rule_proto_rawDescGZIP(), []int{7}
}

func (x *ListRulesRequest) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

// ListRulesResponse is the response to list Rules.
type ListRulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rules in the space.
	Rules []*Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *ListRulesResponse) Reset() {
	*x = ListRulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_v1_rule_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRulesResponse) ProtoMessage() {}

func (x *ListRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_v1_rule_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRulesResponse.ProtoReflect.Descriptor instead.
func (*ListRulesResponse) Descriptor() ([]byte, []int) {
	return file_multiparty_v1_rule_proto_rawDescGZIP(), []int{8}
}

func (x *ListRulesResponse) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// DeleteRuleRequest is the request to delete a Rule.
type DeleteRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Space to delete Rule from.
	Space string `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`
	// Name of Rule to delete.
	RuleName string `protobuf:"bytes,2,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
}

func (x *DeleteRuleRequest) Reset() {
	*x = DeleteRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multiparty_v1_rule_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRuleRequest) ProtoMessage() {}

func (x *DeleteRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_multiparty_v1_rule_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRuleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRuleRequest) Descriptor() ([]byte, []int) {
	return file_multiparty_v1_rule_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRuleRequest) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *DeleteRuleRequest) GetRuleName() string {
	if x != nil {
		return x.RuleName
	}
	return ""
}

var File_multiparty_v1_rule_proto protoreflect.FileDescriptor

var file_multiparty_v1_rule_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x04,
	0x52, 0x75, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x02, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x6e, 0x52, 0x02, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x22, 0x8a, 0x01,
	0x0a, 0x02, 0x4f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x05,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x2e, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x1a, 0x38, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xcf, 0x01, 0x0a, 0x0b, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e,
	0x73, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x73, 0x75,
	0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x41, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x45, 0x78,
	0x70, 0x72, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x22, 0x30, 0x0a, 0x06,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x52,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x72, 0x75, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x75,
	0x6c, 0x65, 0x22, 0x6f, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x72, 0x75,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72,
	0x75, 0x6c, 0x65, 0x22, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x22, 0x3e, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x22, 0x46, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x9f, 0x04, 0x0a, 0x05, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x64, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x20, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x3a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x7d, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x70, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x22,
	0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x32, 0x1d, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2f, 0x7b, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x64, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d,
	0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x12, 0x69, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12,
	0x1f, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x6d, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x2a, 0x1d, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2f, 0x7b, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0xaf, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x42, 0x09, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x61, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58,
	0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0d, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x19, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_multiparty_v1_rule_proto_rawDescOnce sync.Once
	file_multiparty_v1_rule_proto_rawDescData = file_multiparty_v1_rule_proto_rawDesc
)

func file_multiparty_v1_rule_proto_rawDescGZIP() []byte {
	file_multiparty_v1_rule_proto_rawDescOnce.Do(func() {
		file_multiparty_v1_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_multiparty_v1_rule_proto_rawDescData)
	})
	return file_multiparty_v1_rule_proto_rawDescData
}

var file_multiparty_v1_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_multiparty_v1_rule_proto_goTypes = []interface{}{
	(*Rule)(nil),                 // 0: multiparty.v1.Rule
	(*On)(nil),                   // 1: multiparty.v1.On
	(*Requirement)(nil),          // 2: multiparty.v1.Requirement
	(*Origin)(nil),               // 3: multiparty.v1.Origin
	(*CreateRuleRequest)(nil),    // 4: multiparty.v1.CreateRuleRequest
	(*UpdateRuleRequest)(nil),    // 5: multiparty.v1.UpdateRuleRequest
	(*GetRuleRequest)(nil),       // 6: multiparty.v1.GetRuleRequest
	(*ListRulesRequest)(nil),     // 7: multiparty.v1.ListRulesRequest
	(*ListRulesResponse)(nil),    // 8: multiparty.v1.ListRulesResponse
	(*DeleteRuleRequest)(nil),    // 9: multiparty.v1.DeleteRuleRequest
	nil,                          // 10: multiparty.v1.On.MatchEntry
	(*v1alpha1.CheckedExpr)(nil), // 11: google.api.expr.v1alpha1.CheckedExpr
	(*emptypb.Empty)(nil),        // 12: google.protobuf.Empty
}
var file_multiparty_v1_rule_proto_depIdxs = []int32{
	1,  // 0: multiparty.v1.Rule.on:type_name -> multiparty.v1.On
	2,  // 1: multiparty.v1.Rule.require:type_name -> multiparty.v1.Requirement
	10, // 2: multiparty.v1.On.match:type_name -> multiparty.v1.On.MatchEntry
	3,  // 3: multiparty.v1.Requirement.origin:type_name -> multiparty.v1.Origin
	11, // 4: multiparty.v1.Requirement.compiled:type_name -> google.api.expr.v1alpha1.CheckedExpr
	0,  // 5: multiparty.v1.CreateRuleRequest.rule:type_name -> multiparty.v1.Rule
	0,  // 6: multiparty.v1.UpdateRuleRequest.rule:type_name -> multiparty.v1.Rule
	0,  // 7: multiparty.v1.ListRulesResponse.rules:type_name -> multiparty.v1.Rule
	4,  // 8: multiparty.v1.Rules.CreateRule:input_type -> multiparty.v1.CreateRuleRequest
	5,  // 9: multiparty.v1.Rules.UpdateRule:input_type -> multiparty.v1.UpdateRuleRequest
	6,  // 10: multiparty.v1.Rules.GetRule:input_type -> multiparty.v1.GetRuleRequest
	7,  // 11: multiparty.v1.Rules.ListRules:input_type -> multiparty.v1.ListRulesRequest
	9,  // 12: multiparty.v1.Rules.DeleteRule:input_type -> multiparty.v1.DeleteRuleRequest
	0,  // 13: multiparty.v1.Rules.CreateRule:output_type -> multiparty.v1.Rule
	0,  // 14: multiparty.v1.Rules.UpdateRule:output_type -> multiparty.v1.Rule
	0,  // 15: multiparty.v1.Rules.GetRule:output_type -> multiparty.v1.Rule
	8,  // 16: multiparty.v1.Rules.ListRules:output_type -> multiparty.v1.ListRulesResponse
	12, // 17: multiparty.v1.Rules.DeleteRule:output_type -> google.protobuf.Empty
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_multiparty_v1_rule_proto_init() }
func file_multiparty_v1_rule_proto_init() {
	if File_multiparty_v1_rule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_multiparty_v1_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
		file_multiparty_v1_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*On); i {
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
		file_multiparty_v1_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Requirement); i {
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
		file_multiparty_v1_rule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Origin); i {
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
		file_multiparty_v1_rule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRuleRequest); i {
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
		file_multiparty_v1_rule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRuleRequest); i {
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
		file_multiparty_v1_rule_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuleRequest); i {
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
		file_multiparty_v1_rule_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRulesRequest); i {
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
		file_multiparty_v1_rule_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRulesResponse); i {
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
		file_multiparty_v1_rule_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRuleRequest); i {
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
			RawDescriptor: file_multiparty_v1_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_multiparty_v1_rule_proto_goTypes,
		DependencyIndexes: file_multiparty_v1_rule_proto_depIdxs,
		MessageInfos:      file_multiparty_v1_rule_proto_msgTypes,
	}.Build()
	File_multiparty_v1_rule_proto = out.File
	file_multiparty_v1_rule_proto_rawDesc = nil
	file_multiparty_v1_rule_proto_goTypes = nil
	file_multiparty_v1_rule_proto_depIdxs = nil
}
