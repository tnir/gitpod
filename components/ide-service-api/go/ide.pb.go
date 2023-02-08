// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: ide.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkspaceType int32

const (
	WorkspaceType_REGULAR  WorkspaceType = 0
	WorkspaceType_PREBUILD WorkspaceType = 1
)

// Enum value maps for WorkspaceType.
var (
	WorkspaceType_name = map[int32]string{
		0: "REGULAR",
		1: "PREBUILD",
	}
	WorkspaceType_value = map[string]int32{
		"REGULAR":  0,
		"PREBUILD": 1,
	}
)

func (x WorkspaceType) Enum() *WorkspaceType {
	p := new(WorkspaceType)
	*p = x
	return p
}

func (x WorkspaceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkspaceType) Descriptor() protoreflect.EnumDescriptor {
	return file_ide_proto_enumTypes[0].Descriptor()
}

func (WorkspaceType) Type() protoreflect.EnumType {
	return &file_ide_proto_enumTypes[0]
}

func (x WorkspaceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkspaceType.Descriptor instead.
func (WorkspaceType) EnumDescriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{0}
}

type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{0}
}

type GetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetConfigResponse) Reset() {
	*x = GetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigResponse) ProtoMessage() {}

func (x *GetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigResponse.ProtoReflect.Descriptor instead.
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{1}
}

func (x *GetConfigResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// TODO: import type from other packages
// EnvironmentVariable describes an env var as key/value pair
type EnvironmentVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EnvironmentVariable) Reset() {
	*x = EnvironmentVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ide_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentVariable) ProtoMessage() {}

func (x *EnvironmentVariable) ProtoReflect() protoreflect.Message {
	mi := &file_ide_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentVariable.ProtoReflect.Descriptor instead.
func (*EnvironmentVariable) Descriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{2}
}

func (x *EnvironmentVariable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnvironmentVariable) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email *string `protobuf:"bytes,2,opt,name=email,proto3,oneof" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ide_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_ide_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

type ResolveWorkspaceConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type            WorkspaceType `protobuf:"varint,1,opt,name=type,proto3,enum=ide_service_api.WorkspaceType" json:"type,omitempty"`
	Context         string        `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	IdeSettings     string        `protobuf:"bytes,3,opt,name=ide_settings,json=ideSettings,proto3" json:"ide_settings,omitempty"`
	WorkspaceConfig string        `protobuf:"bytes,4,opt,name=workspace_config,json=workspaceConfig,proto3" json:"workspace_config,omitempty"`
	User            *User         `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ResolveWorkspaceConfigRequest) Reset() {
	*x = ResolveWorkspaceConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ide_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveWorkspaceConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveWorkspaceConfigRequest) ProtoMessage() {}

func (x *ResolveWorkspaceConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ide_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveWorkspaceConfigRequest.ProtoReflect.Descriptor instead.
func (*ResolveWorkspaceConfigRequest) Descriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{4}
}

func (x *ResolveWorkspaceConfigRequest) GetType() WorkspaceType {
	if x != nil {
		return x.Type
	}
	return WorkspaceType_REGULAR
}

func (x *ResolveWorkspaceConfigRequest) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *ResolveWorkspaceConfigRequest) GetIdeSettings() string {
	if x != nil {
		return x.IdeSettings
	}
	return ""
}

func (x *ResolveWorkspaceConfigRequest) GetWorkspaceConfig() string {
	if x != nil {
		return x.WorkspaceConfig
	}
	return ""
}

func (x *ResolveWorkspaceConfigRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type ResolveWorkspaceConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envvars         []*EnvironmentVariable `protobuf:"bytes,1,rep,name=envvars,proto3" json:"envvars,omitempty"`
	SupervisorImage string                 `protobuf:"bytes,2,opt,name=supervisor_image,json=supervisorImage,proto3" json:"supervisor_image,omitempty"`
	WebImage        string                 `protobuf:"bytes,3,opt,name=web_image,json=webImage,proto3" json:"web_image,omitempty"`
	IdeImageLayers  []string               `protobuf:"bytes,4,rep,name=ide_image_layers,json=ideImageLayers,proto3" json:"ide_image_layers,omitempty"`
	// control whether to configure default IDE for a user
	RefererIde string `protobuf:"bytes,5,opt,name=referer_ide,json=refererIde,proto3" json:"referer_ide,omitempty"`
	Tasks      string `protobuf:"bytes,6,opt,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *ResolveWorkspaceConfigResponse) Reset() {
	*x = ResolveWorkspaceConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ide_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveWorkspaceConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveWorkspaceConfigResponse) ProtoMessage() {}

func (x *ResolveWorkspaceConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ide_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveWorkspaceConfigResponse.ProtoReflect.Descriptor instead.
func (*ResolveWorkspaceConfigResponse) Descriptor() ([]byte, []int) {
	return file_ide_proto_rawDescGZIP(), []int{5}
}

func (x *ResolveWorkspaceConfigResponse) GetEnvvars() []*EnvironmentVariable {
	if x != nil {
		return x.Envvars
	}
	return nil
}

func (x *ResolveWorkspaceConfigResponse) GetSupervisorImage() string {
	if x != nil {
		return x.SupervisorImage
	}
	return ""
}

func (x *ResolveWorkspaceConfigResponse) GetWebImage() string {
	if x != nil {
		return x.WebImage
	}
	return ""
}

func (x *ResolveWorkspaceConfigResponse) GetIdeImageLayers() []string {
	if x != nil {
		return x.IdeImageLayers
	}
	return nil
}

func (x *ResolveWorkspaceConfigResponse) GetRefererIde() string {
	if x != nil {
		return x.RefererIde
	}
	return ""
}

func (x *ResolveWorkspaceConfigResponse) GetTasks() string {
	if x != nil {
		return x.Tasks
	}
	return ""
}

var File_ide_proto protoreflect.FileDescriptor

var file_ide_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x69, 0x64, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x12, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x3f, 0x0a, 0x13, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x3b, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xe6, 0x01,
	0x0a, 0x1d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x69, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x29, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x29, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x64, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x89, 0x02, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x65, 0x6e, 0x76,
	0x76, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x64, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x07, 0x65, 0x6e, 0x76, 0x76, 0x61, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x65, 0x62, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x65, 0x62, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x72, 0x49, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2a, 0x2a, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x45, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x10, 0x01, 0x32, 0xe5,
	0x01, 0x0a, 0x0a, 0x49, 0x44, 0x45, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x2e, 0x69, 0x64, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x69, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x03, 0x90, 0x02, 0x02, 0x12, 0x7e, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x2e, 0x2e, 0x69, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x69, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x03, 0x90, 0x02, 0x02, 0x42, 0x47, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74,
	0x70, 0x6f, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f,
	0x69, 0x64, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ide_proto_rawDescOnce sync.Once
	file_ide_proto_rawDescData = file_ide_proto_rawDesc
)

func file_ide_proto_rawDescGZIP() []byte {
	file_ide_proto_rawDescOnce.Do(func() {
		file_ide_proto_rawDescData = protoimpl.X.CompressGZIP(file_ide_proto_rawDescData)
	})
	return file_ide_proto_rawDescData
}

var file_ide_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ide_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ide_proto_goTypes = []interface{}{
	(WorkspaceType)(0),                     // 0: ide_service_api.WorkspaceType
	(*GetConfigRequest)(nil),               // 1: ide_service_api.GetConfigRequest
	(*GetConfigResponse)(nil),              // 2: ide_service_api.GetConfigResponse
	(*EnvironmentVariable)(nil),            // 3: ide_service_api.EnvironmentVariable
	(*User)(nil),                           // 4: ide_service_api.User
	(*ResolveWorkspaceConfigRequest)(nil),  // 5: ide_service_api.ResolveWorkspaceConfigRequest
	(*ResolveWorkspaceConfigResponse)(nil), // 6: ide_service_api.ResolveWorkspaceConfigResponse
}
var file_ide_proto_depIdxs = []int32{
	0, // 0: ide_service_api.ResolveWorkspaceConfigRequest.type:type_name -> ide_service_api.WorkspaceType
	4, // 1: ide_service_api.ResolveWorkspaceConfigRequest.user:type_name -> ide_service_api.User
	3, // 2: ide_service_api.ResolveWorkspaceConfigResponse.envvars:type_name -> ide_service_api.EnvironmentVariable
	1, // 3: ide_service_api.IDEService.GetConfig:input_type -> ide_service_api.GetConfigRequest
	5, // 4: ide_service_api.IDEService.ResolveWorkspaceConfig:input_type -> ide_service_api.ResolveWorkspaceConfigRequest
	2, // 5: ide_service_api.IDEService.GetConfig:output_type -> ide_service_api.GetConfigResponse
	6, // 6: ide_service_api.IDEService.ResolveWorkspaceConfig:output_type -> ide_service_api.ResolveWorkspaceConfigResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ide_proto_init() }
func file_ide_proto_init() {
	if File_ide_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigRequest); i {
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
		file_ide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigResponse); i {
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
		file_ide_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentVariable); i {
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
		file_ide_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_ide_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveWorkspaceConfigRequest); i {
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
		file_ide_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveWorkspaceConfigResponse); i {
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
	file_ide_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ide_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ide_proto_goTypes,
		DependencyIndexes: file_ide_proto_depIdxs,
		EnumInfos:         file_ide_proto_enumTypes,
		MessageInfos:      file_ide_proto_msgTypes,
	}.Build()
	File_ide_proto = out.File
	file_ide_proto_rawDesc = nil
	file_ide_proto_goTypes = nil
	file_ide_proto_depIdxs = nil
}