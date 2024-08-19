// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/securitycenter/v1/group_membership.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible types of groups.
type GroupMembership_GroupType int32

const (
	// Default value.
	GroupMembership_GROUP_TYPE_UNSPECIFIED GroupMembership_GroupType = 0
	// Group represents a toxic combination.
	GroupMembership_GROUP_TYPE_TOXIC_COMBINATION GroupMembership_GroupType = 1
)

// Enum value maps for GroupMembership_GroupType.
var (
	GroupMembership_GroupType_name = map[int32]string{
		0: "GROUP_TYPE_UNSPECIFIED",
		1: "GROUP_TYPE_TOXIC_COMBINATION",
	}
	GroupMembership_GroupType_value = map[string]int32{
		"GROUP_TYPE_UNSPECIFIED":       0,
		"GROUP_TYPE_TOXIC_COMBINATION": 1,
	}
)

func (x GroupMembership_GroupType) Enum() *GroupMembership_GroupType {
	p := new(GroupMembership_GroupType)
	*p = x
	return p
}

func (x GroupMembership_GroupType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupMembership_GroupType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securitycenter_v1_group_membership_proto_enumTypes[0].Descriptor()
}

func (GroupMembership_GroupType) Type() protoreflect.EnumType {
	return &file_google_cloud_securitycenter_v1_group_membership_proto_enumTypes[0]
}

func (x GroupMembership_GroupType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupMembership_GroupType.Descriptor instead.
func (GroupMembership_GroupType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_group_membership_proto_rawDescGZIP(), []int{0, 0}
}

// Contains details about groups of which this finding is a member. A group is a
// collection of findings that are related in some way.
type GroupMembership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of group.
	GroupType GroupMembership_GroupType `protobuf:"varint,1,opt,name=group_type,json=groupType,proto3,enum=google.cloud.securitycenter.v1.GroupMembership_GroupType" json:"group_type,omitempty"`
	// ID of the group.
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GroupMembership) Reset() {
	*x = GroupMembership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_group_membership_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMembership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMembership) ProtoMessage() {}

func (x *GroupMembership) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_group_membership_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMembership.ProtoReflect.Descriptor instead.
func (*GroupMembership) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_group_membership_proto_rawDescGZIP(), []int{0}
}

func (x *GroupMembership) GetGroupType() GroupMembership_GroupType {
	if x != nil {
		return x.GroupType
	}
	return GroupMembership_GROUP_TYPE_UNSPECIFIED
}

func (x *GroupMembership) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

var File_google_cloud_securitycenter_v1_group_membership_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_group_membership_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xd1, 0x01, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x58, 0x0a, 0x0a, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x22, 0x49, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x16, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x4f, 0x58, 0x49, 0x43, 0x5f, 0x43, 0x4f,
	0x4d, 0x42, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x42, 0xee, 0x01, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_group_membership_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_group_membership_proto_rawDescData = file_google_cloud_securitycenter_v1_group_membership_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_group_membership_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_group_membership_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_group_membership_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_group_membership_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_group_membership_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_group_membership_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_securitycenter_v1_group_membership_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v1_group_membership_proto_goTypes = []any{
	(GroupMembership_GroupType)(0), // 0: google.cloud.securitycenter.v1.GroupMembership.GroupType
	(*GroupMembership)(nil),        // 1: google.cloud.securitycenter.v1.GroupMembership
}
var file_google_cloud_securitycenter_v1_group_membership_proto_depIdxs = []int32{
	0, // 0: google.cloud.securitycenter.v1.GroupMembership.group_type:type_name -> google.cloud.securitycenter.v1.GroupMembership.GroupType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_group_membership_proto_init() }
func file_google_cloud_securitycenter_v1_group_membership_proto_init() {
	if File_google_cloud_securitycenter_v1_group_membership_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_group_membership_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GroupMembership); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v1_group_membership_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_group_membership_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_group_membership_proto_depIdxs,
		EnumInfos:         file_google_cloud_securitycenter_v1_group_membership_proto_enumTypes,
		MessageInfos:      file_google_cloud_securitycenter_v1_group_membership_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_group_membership_proto = out.File
	file_google_cloud_securitycenter_v1_group_membership_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_group_membership_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_group_membership_proto_depIdxs = nil
}
