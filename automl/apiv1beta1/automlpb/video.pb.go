// Copyright 2020 Google LLC
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
// source: google/cloud/automl/v1beta1/video.proto

package automlpb

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

// Dataset metadata specific to video classification.
// All Video Classification datasets are treated as multi label.
type VideoClassificationDatasetMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VideoClassificationDatasetMetadata) Reset() {
	*x = VideoClassificationDatasetMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoClassificationDatasetMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoClassificationDatasetMetadata) ProtoMessage() {}

func (x *VideoClassificationDatasetMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoClassificationDatasetMetadata.ProtoReflect.Descriptor instead.
func (*VideoClassificationDatasetMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_video_proto_rawDescGZIP(), []int{0}
}

// Dataset metadata specific to video object tracking.
type VideoObjectTrackingDatasetMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VideoObjectTrackingDatasetMetadata) Reset() {
	*x = VideoObjectTrackingDatasetMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoObjectTrackingDatasetMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoObjectTrackingDatasetMetadata) ProtoMessage() {}

func (x *VideoObjectTrackingDatasetMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoObjectTrackingDatasetMetadata.ProtoReflect.Descriptor instead.
func (*VideoObjectTrackingDatasetMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_video_proto_rawDescGZIP(), []int{1}
}

// Model metadata specific to video classification.
type VideoClassificationModelMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VideoClassificationModelMetadata) Reset() {
	*x = VideoClassificationModelMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoClassificationModelMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoClassificationModelMetadata) ProtoMessage() {}

func (x *VideoClassificationModelMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoClassificationModelMetadata.ProtoReflect.Descriptor instead.
func (*VideoClassificationModelMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_video_proto_rawDescGZIP(), []int{2}
}

// Model metadata specific to video object tracking.
type VideoObjectTrackingModelMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VideoObjectTrackingModelMetadata) Reset() {
	*x = VideoObjectTrackingModelMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoObjectTrackingModelMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoObjectTrackingModelMetadata) ProtoMessage() {}

func (x *VideoObjectTrackingModelMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoObjectTrackingModelMetadata.ProtoReflect.Descriptor instead.
func (*VideoObjectTrackingModelMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_video_proto_rawDescGZIP(), []int{3}
}

var File_google_cloud_automl_v1beta1_video_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1beta1_video_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x22, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24,
	0x0a, 0x22, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x20, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x20, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xb1, 0x01, 0x0a,
	0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x0a, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1beta1_video_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1beta1_video_proto_rawDescData = file_google_cloud_automl_v1beta1_video_proto_rawDesc
)

func file_google_cloud_automl_v1beta1_video_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1beta1_video_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1beta1_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1beta1_video_proto_rawDescData)
	})
	return file_google_cloud_automl_v1beta1_video_proto_rawDescData
}

var file_google_cloud_automl_v1beta1_video_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_automl_v1beta1_video_proto_goTypes = []interface{}{
	(*VideoClassificationDatasetMetadata)(nil), // 0: google.cloud.automl.v1beta1.VideoClassificationDatasetMetadata
	(*VideoObjectTrackingDatasetMetadata)(nil), // 1: google.cloud.automl.v1beta1.VideoObjectTrackingDatasetMetadata
	(*VideoClassificationModelMetadata)(nil),   // 2: google.cloud.automl.v1beta1.VideoClassificationModelMetadata
	(*VideoObjectTrackingModelMetadata)(nil),   // 3: google.cloud.automl.v1beta1.VideoObjectTrackingModelMetadata
}
var file_google_cloud_automl_v1beta1_video_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1beta1_video_proto_init() }
func file_google_cloud_automl_v1beta1_video_proto_init() {
	if File_google_cloud_automl_v1beta1_video_proto != nil {
		return
	}
	file_google_cloud_automl_v1beta1_classification_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1beta1_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoClassificationDatasetMetadata); i {
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
		file_google_cloud_automl_v1beta1_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoObjectTrackingDatasetMetadata); i {
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
		file_google_cloud_automl_v1beta1_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoClassificationModelMetadata); i {
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
		file_google_cloud_automl_v1beta1_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoObjectTrackingModelMetadata); i {
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
			RawDescriptor: file_google_cloud_automl_v1beta1_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1beta1_video_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1beta1_video_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1beta1_video_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1beta1_video_proto = out.File
	file_google_cloud_automl_v1beta1_video_proto_rawDesc = nil
	file_google_cloud_automl_v1beta1_video_proto_goTypes = nil
	file_google_cloud_automl_v1beta1_video_proto_depIdxs = nil
}
