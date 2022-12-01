// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/cloud/resourcemanager/v3/tag_bindings.proto

package resourcemanagerpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A TagBinding represents a connection between a TagValue and a cloud
// resource (currently project, folder, or organization). Once a TagBinding is
// created, the TagValue is applied to all the descendants of the cloud
// resource.
type TagBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The name of the TagBinding. This is a String of the form:
	// `tagBindings/{full-resource-name}/{tag-value-name}` (e.g.
	// `tagBindings/%2F%2Fcloudresourcemanager.googleapis.com%2Fprojects%2F123/tagValues/456`).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The full resource name of the resource the TagValue is bound to.
	// E.g. `//cloudresourcemanager.googleapis.com/projects/123`
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	// The TagValue of the TagBinding.
	// Must be of the form `tagValues/456`.
	TagValue string `protobuf:"bytes,3,opt,name=tag_value,json=tagValue,proto3" json:"tag_value,omitempty"`
}

func (x *TagBinding) Reset() {
	*x = TagBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagBinding) ProtoMessage() {}

func (x *TagBinding) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagBinding.ProtoReflect.Descriptor instead.
func (*TagBinding) Descriptor() ([]byte, []int) {
	return file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescGZIP(), []int{0}
}

func (x *TagBinding) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TagBinding) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *TagBinding) GetTagValue() string {
	if x != nil {
		return x.TagValue
	}
	return ""
}

// Runtime operation information for creating a TagValue.
type CreateTagBindingMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTagBindingMetadata) Reset() {
	*x = CreateTagBindingMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagBindingMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagBindingMetadata) ProtoMessage() {}

func (x *CreateTagBindingMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagBindingMetadata.ProtoReflect.Descriptor instead.
func (*CreateTagBindingMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescGZIP(), []int{1}
}

// The request message to create a TagBinding.
type CreateTagBindingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The TagBinding to be created.
	TagBinding *TagBinding `protobuf:"bytes,1,opt,name=tag_binding,json=tagBinding,proto3" json:"tag_binding,omitempty"`
	// Optional. Set to true to perform the validations necessary for creating the resource,
	// but not actually perform the action.
	ValidateOnly bool `protobuf:"varint,2,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *CreateTagBindingRequest) Reset() {
	*x = CreateTagBindingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagBindingRequest) ProtoMessage() {}

func (x *CreateTagBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagBindingRequest.ProtoReflect.Descriptor instead.
func (*CreateTagBindingRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTagBindingRequest) GetTagBinding() *TagBinding {
	if x != nil {
		return x.TagBinding
	}
	return nil
}

func (x *CreateTagBindingRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// Runtime operation information for deleting a TagBinding.
type DeleteTagBindingMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTagBindingMetadata) Reset() {
	*x = DeleteTagBindingMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagBindingMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagBindingMetadata) ProtoMessage() {}

func (x *DeleteTagBindingMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagBindingMetadata.ProtoReflect.Descriptor instead.
func (*DeleteTagBindingMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescGZIP(), []int{3}
}

// The request message to delete a TagBinding.
type DeleteTagBindingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the TagBinding. This is a String of the form:
	// `tagBindings/{id}` (e.g.
	// `tagBindings/%2F%2Fcloudresourcemanager.googleapis.com%2Fprojects%2F123/tagValues/456`).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteTagBindingRequest) Reset() {
	*x = DeleteTagBindingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagBindingRequest) ProtoMessage() {}

func (x *DeleteTagBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagBindingRequest.ProtoReflect.Descriptor instead.
func (*DeleteTagBindingRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTagBindingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request message to list all TagBindings for a parent.
type ListTagBindingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The full resource name of a resource for which you want to list existing
	// TagBindings.
	// E.g. "//cloudresourcemanager.googleapis.com/projects/123"
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of TagBindings to return in the response. The server
	// allows a maximum of 300 TagBindings to return. If unspecified, the server
	// will use 100 as the default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. A pagination token returned from a previous call to `ListTagBindings`
	// that indicates where this listing should continue from.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListTagBindingsRequest) Reset() {
	*x = ListTagBindingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagBindingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagBindingsRequest) ProtoMessage() {}

func (x *ListTagBindingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagBindingsRequest.ProtoReflect.Descriptor instead.
func (*ListTagBindingsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescGZIP(), []int{5}
}

func (x *ListTagBindingsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListTagBindingsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTagBindingsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The ListTagBindings response.
type ListTagBindingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A possibly paginated list of TagBindings for the specified TagValue or
	// resource.
	TagBindings []*TagBinding `protobuf:"bytes,1,rep,name=tag_bindings,json=tagBindings,proto3" json:"tag_bindings,omitempty"`
	// Pagination token.
	//
	// If the result set is too large to fit in a single response, this token
	// is returned. It encodes the position of the current result cursor.
	// Feeding this value into a new list request with the `page_token` parameter
	// gives the next page of the results.
	//
	// When `next_page_token` is not filled in, there is no next page and
	// the list returned is the last page in the result set.
	//
	// Pagination tokens have a limited lifetime.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTagBindingsResponse) Reset() {
	*x = ListTagBindingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagBindingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagBindingsResponse) ProtoMessage() {}

func (x *ListTagBindingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagBindingsResponse.ProtoReflect.Descriptor instead.
func (*ListTagBindingsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescGZIP(), []int{6}
}

func (x *ListTagBindingsResponse) GetTagBindings() []*TagBinding {
	if x != nil {
		return x.TagBindings
	}
	return nil
}

func (x *ListTagBindingsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_google_cloud_resourcemanager_v3_tag_bindings_proto protoreflect.FileDescriptor

var file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x2f, 0x74, 0x61, 0x67, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01,
	0x0a, 0x0a, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x61, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x4e, 0xea, 0x41, 0x4b, 0x0a,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x19, 0x74, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x74, 0x61,
	0x67, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x7d, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x96, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x51, 0x0a, 0x0b, 0x74, 0x61, 0x67, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x74, 0x61, 0x67, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22,
	0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x65, 0x0a, 0x17, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe0,
	0x41, 0x02, 0xfa, 0x41, 0x03, 0x12, 0x01, 0x2a, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x74, 0x61, 0x67, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x61, 0x67, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x74, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xe0, 0x05, 0x0a, 0x0b, 0x54,
	0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0xa6, 0x01, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x37,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x33, 0x2f, 0x74,
	0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0xc8, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x0f, 0x2f, 0x76, 0x33, 0x2f, 0x74,
	0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x3a, 0x0b, 0x74, 0x61, 0x67, 0x5f,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0xda, 0x41, 0x0b, 0x74, 0x61, 0x67, 0x5f, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0xca, 0x41, 0x26, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0xc9,
	0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x2a, 0x19, 0x2f, 0x76, 0x33, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x74, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2a, 0x2a, 0x7d, 0xda,
	0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xca, 0x41, 0x31, 0x0a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x90, 0x01, 0xca, 0x41, 0x23,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x67, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x2d, 0x6f, 0x6e, 0x6c, 0x79, 0x42, 0xf2, 0x01,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x10, 0x54, 0x61, 0x67, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x33, 0xea, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a,
	0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescOnce sync.Once
	file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescData = file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDesc
)

func file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescGZIP() []byte {
	file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescOnce.Do(func() {
		file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescData)
	})
	return file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDescData
}

var file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_cloud_resourcemanager_v3_tag_bindings_proto_goTypes = []interface{}{
	(*TagBinding)(nil),               // 0: google.cloud.resourcemanager.v3.TagBinding
	(*CreateTagBindingMetadata)(nil), // 1: google.cloud.resourcemanager.v3.CreateTagBindingMetadata
	(*CreateTagBindingRequest)(nil),  // 2: google.cloud.resourcemanager.v3.CreateTagBindingRequest
	(*DeleteTagBindingMetadata)(nil), // 3: google.cloud.resourcemanager.v3.DeleteTagBindingMetadata
	(*DeleteTagBindingRequest)(nil),  // 4: google.cloud.resourcemanager.v3.DeleteTagBindingRequest
	(*ListTagBindingsRequest)(nil),   // 5: google.cloud.resourcemanager.v3.ListTagBindingsRequest
	(*ListTagBindingsResponse)(nil),  // 6: google.cloud.resourcemanager.v3.ListTagBindingsResponse
	(*longrunning.Operation)(nil),    // 7: google.longrunning.Operation
}
var file_google_cloud_resourcemanager_v3_tag_bindings_proto_depIdxs = []int32{
	0, // 0: google.cloud.resourcemanager.v3.CreateTagBindingRequest.tag_binding:type_name -> google.cloud.resourcemanager.v3.TagBinding
	0, // 1: google.cloud.resourcemanager.v3.ListTagBindingsResponse.tag_bindings:type_name -> google.cloud.resourcemanager.v3.TagBinding
	5, // 2: google.cloud.resourcemanager.v3.TagBindings.ListTagBindings:input_type -> google.cloud.resourcemanager.v3.ListTagBindingsRequest
	2, // 3: google.cloud.resourcemanager.v3.TagBindings.CreateTagBinding:input_type -> google.cloud.resourcemanager.v3.CreateTagBindingRequest
	4, // 4: google.cloud.resourcemanager.v3.TagBindings.DeleteTagBinding:input_type -> google.cloud.resourcemanager.v3.DeleteTagBindingRequest
	6, // 5: google.cloud.resourcemanager.v3.TagBindings.ListTagBindings:output_type -> google.cloud.resourcemanager.v3.ListTagBindingsResponse
	7, // 6: google.cloud.resourcemanager.v3.TagBindings.CreateTagBinding:output_type -> google.longrunning.Operation
	7, // 7: google.cloud.resourcemanager.v3.TagBindings.DeleteTagBinding:output_type -> google.longrunning.Operation
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_resourcemanager_v3_tag_bindings_proto_init() }
func file_google_cloud_resourcemanager_v3_tag_bindings_proto_init() {
	if File_google_cloud_resourcemanager_v3_tag_bindings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagBinding); i {
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
		file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagBindingMetadata); i {
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
		file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagBindingRequest); i {
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
		file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagBindingMetadata); i {
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
		file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagBindingRequest); i {
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
		file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagBindingsRequest); i {
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
		file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagBindingsResponse); i {
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
			RawDescriptor: file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_resourcemanager_v3_tag_bindings_proto_goTypes,
		DependencyIndexes: file_google_cloud_resourcemanager_v3_tag_bindings_proto_depIdxs,
		MessageInfos:      file_google_cloud_resourcemanager_v3_tag_bindings_proto_msgTypes,
	}.Build()
	File_google_cloud_resourcemanager_v3_tag_bindings_proto = out.File
	file_google_cloud_resourcemanager_v3_tag_bindings_proto_rawDesc = nil
	file_google_cloud_resourcemanager_v3_tag_bindings_proto_goTypes = nil
	file_google_cloud_resourcemanager_v3_tag_bindings_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TagBindingsClient is the client API for TagBindings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TagBindingsClient interface {
	// Lists the TagBindings for the given cloud resource, as specified with
	// `parent`.
	//
	// NOTE: The `parent` field is expected to be a full resource name:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ListTagBindings(ctx context.Context, in *ListTagBindingsRequest, opts ...grpc.CallOption) (*ListTagBindingsResponse, error)
	// Creates a TagBinding between a TagValue and a cloud resource
	// (currently project, folder, or organization).
	CreateTagBinding(ctx context.Context, in *CreateTagBindingRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Deletes a TagBinding.
	DeleteTagBinding(ctx context.Context, in *DeleteTagBindingRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type tagBindingsClient struct {
	cc grpc.ClientConnInterface
}

func NewTagBindingsClient(cc grpc.ClientConnInterface) TagBindingsClient {
	return &tagBindingsClient{cc}
}

func (c *tagBindingsClient) ListTagBindings(ctx context.Context, in *ListTagBindingsRequest, opts ...grpc.CallOption) (*ListTagBindingsResponse, error) {
	out := new(ListTagBindingsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.resourcemanager.v3.TagBindings/ListTagBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagBindingsClient) CreateTagBinding(ctx context.Context, in *CreateTagBindingRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.resourcemanager.v3.TagBindings/CreateTagBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagBindingsClient) DeleteTagBinding(ctx context.Context, in *DeleteTagBindingRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.resourcemanager.v3.TagBindings/DeleteTagBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagBindingsServer is the server API for TagBindings service.
type TagBindingsServer interface {
	// Lists the TagBindings for the given cloud resource, as specified with
	// `parent`.
	//
	// NOTE: The `parent` field is expected to be a full resource name:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	ListTagBindings(context.Context, *ListTagBindingsRequest) (*ListTagBindingsResponse, error)
	// Creates a TagBinding between a TagValue and a cloud resource
	// (currently project, folder, or organization).
	CreateTagBinding(context.Context, *CreateTagBindingRequest) (*longrunning.Operation, error)
	// Deletes a TagBinding.
	DeleteTagBinding(context.Context, *DeleteTagBindingRequest) (*longrunning.Operation, error)
}

// UnimplementedTagBindingsServer can be embedded to have forward compatible implementations.
type UnimplementedTagBindingsServer struct {
}

func (*UnimplementedTagBindingsServer) ListTagBindings(context.Context, *ListTagBindingsRequest) (*ListTagBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTagBindings not implemented")
}
func (*UnimplementedTagBindingsServer) CreateTagBinding(context.Context, *CreateTagBindingRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTagBinding not implemented")
}
func (*UnimplementedTagBindingsServer) DeleteTagBinding(context.Context, *DeleteTagBindingRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTagBinding not implemented")
}

func RegisterTagBindingsServer(s *grpc.Server, srv TagBindingsServer) {
	s.RegisterService(&_TagBindings_serviceDesc, srv)
}

func _TagBindings_ListTagBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagBindingsServer).ListTagBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.resourcemanager.v3.TagBindings/ListTagBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagBindingsServer).ListTagBindings(ctx, req.(*ListTagBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagBindings_CreateTagBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagBindingsServer).CreateTagBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.resourcemanager.v3.TagBindings/CreateTagBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagBindingsServer).CreateTagBinding(ctx, req.(*CreateTagBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagBindings_DeleteTagBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagBindingsServer).DeleteTagBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.resourcemanager.v3.TagBindings/DeleteTagBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagBindingsServer).DeleteTagBinding(ctx, req.(*DeleteTagBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagBindings_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.resourcemanager.v3.TagBindings",
	HandlerType: (*TagBindingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTagBindings",
			Handler:    _TagBindings_ListTagBindings_Handler,
		},
		{
			MethodName: "CreateTagBinding",
			Handler:    _TagBindings_CreateTagBinding_Handler,
		},
		{
			MethodName: "DeleteTagBinding",
			Handler:    _TagBindings_DeleteTagBinding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/resourcemanager/v3/tag_bindings.proto",
}
