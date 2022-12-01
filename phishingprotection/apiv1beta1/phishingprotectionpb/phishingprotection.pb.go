// Copyright 2019 Google LLC.
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: google/cloud/phishingprotection/v1beta1/phishingprotection.proto

package phishingprotectionpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The ReportPhishing request message.
type ReportPhishingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the project for which the report will be created,
	// in the format "projects/{project_number}".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The URI that is being reported for phishing content to be analyzed.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *ReportPhishingRequest) Reset() {
	*x = ReportPhishingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportPhishingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportPhishingRequest) ProtoMessage() {}

func (x *ReportPhishingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportPhishingRequest.ProtoReflect.Descriptor instead.
func (*ReportPhishingRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDescGZIP(), []int{0}
}

func (x *ReportPhishingRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ReportPhishingRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

// The ReportPhishing (empty) response message.
type ReportPhishingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportPhishingResponse) Reset() {
	*x = ReportPhishingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportPhishingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportPhishingResponse) ProtoMessage() {}

func (x *ReportPhishingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportPhishingResponse.ProtoReflect.Descriptor instead.
func (*ReportPhishingResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDescGZIP(), []int{1}
}

var File_google_cloud_phishingprotection_v1beta1_phishingprotection_proto protoreflect.FileDescriptor

var file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x68, 0x69, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b,
	0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a,
	0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x75, 0x72, 0x69, 0x22, 0x18, 0x0a, 0x16, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd3, 0x02, 0x0a, 0x20, 0x50, 0x68, 0x69, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0x12, 0xd7, 0x01, 0x0a, 0x0e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x3e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x68, 0x69,
	0x73, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x68,
	0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x68, 0x69,
	0x73, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x68,
	0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x22, 0x2c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x70, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x3a, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x2c, 0x75, 0x72, 0x69, 0x1a, 0x55, 0xca, 0x41, 0x21, 0x70, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xa5, 0x02, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x68, 0x69, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x17, 0x50, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x59, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x68,
	0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x70, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x04, 0x47, 0x43,
	0x50, 0x50, 0xaa, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x50, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x27, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x68, 0x69, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x2a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDescOnce sync.Once
	file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDescData = file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDesc
)

func file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDescGZIP() []byte {
	file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDescOnce.Do(func() {
		file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDescData)
	})
	return file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDescData
}

var file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_goTypes = []interface{}{
	(*ReportPhishingRequest)(nil),  // 0: google.cloud.phishingprotection.v1beta1.ReportPhishingRequest
	(*ReportPhishingResponse)(nil), // 1: google.cloud.phishingprotection.v1beta1.ReportPhishingResponse
}
var file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_depIdxs = []int32{
	0, // 0: google.cloud.phishingprotection.v1beta1.PhishingProtectionServiceV1Beta1.ReportPhishing:input_type -> google.cloud.phishingprotection.v1beta1.ReportPhishingRequest
	1, // 1: google.cloud.phishingprotection.v1beta1.PhishingProtectionServiceV1Beta1.ReportPhishing:output_type -> google.cloud.phishingprotection.v1beta1.ReportPhishingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_init() }
func file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_init() {
	if File_google_cloud_phishingprotection_v1beta1_phishingprotection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportPhishingRequest); i {
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
		file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportPhishingResponse); i {
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
			RawDescriptor: file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_goTypes,
		DependencyIndexes: file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_depIdxs,
		MessageInfos:      file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_msgTypes,
	}.Build()
	File_google_cloud_phishingprotection_v1beta1_phishingprotection_proto = out.File
	file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_rawDesc = nil
	file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_goTypes = nil
	file_google_cloud_phishingprotection_v1beta1_phishingprotection_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PhishingProtectionServiceV1Beta1Client is the client API for PhishingProtectionServiceV1Beta1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PhishingProtectionServiceV1Beta1Client interface {
	// Reports a URI suspected of containing phishing content to be reviewed. Once
	// the report review is complete, its result can be found in the Cloud
	// Security Command Center findings dashboard for Phishing Protection. If the
	// result verifies the existence of malicious phishing content, the site will
	// be added the to [Google's Social Engineering
	// lists](https://support.google.com/webmasters/answer/6350487/) in order to
	// protect users that could get exposed to this threat in the future.
	ReportPhishing(ctx context.Context, in *ReportPhishingRequest, opts ...grpc.CallOption) (*ReportPhishingResponse, error)
}

type phishingProtectionServiceV1Beta1Client struct {
	cc grpc.ClientConnInterface
}

func NewPhishingProtectionServiceV1Beta1Client(cc grpc.ClientConnInterface) PhishingProtectionServiceV1Beta1Client {
	return &phishingProtectionServiceV1Beta1Client{cc}
}

func (c *phishingProtectionServiceV1Beta1Client) ReportPhishing(ctx context.Context, in *ReportPhishingRequest, opts ...grpc.CallOption) (*ReportPhishingResponse, error) {
	out := new(ReportPhishingResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.phishingprotection.v1beta1.PhishingProtectionServiceV1Beta1/ReportPhishing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhishingProtectionServiceV1Beta1Server is the server API for PhishingProtectionServiceV1Beta1 service.
type PhishingProtectionServiceV1Beta1Server interface {
	// Reports a URI suspected of containing phishing content to be reviewed. Once
	// the report review is complete, its result can be found in the Cloud
	// Security Command Center findings dashboard for Phishing Protection. If the
	// result verifies the existence of malicious phishing content, the site will
	// be added the to [Google's Social Engineering
	// lists](https://support.google.com/webmasters/answer/6350487/) in order to
	// protect users that could get exposed to this threat in the future.
	ReportPhishing(context.Context, *ReportPhishingRequest) (*ReportPhishingResponse, error)
}

// UnimplementedPhishingProtectionServiceV1Beta1Server can be embedded to have forward compatible implementations.
type UnimplementedPhishingProtectionServiceV1Beta1Server struct {
}

func (*UnimplementedPhishingProtectionServiceV1Beta1Server) ReportPhishing(context.Context, *ReportPhishingRequest) (*ReportPhishingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportPhishing not implemented")
}

func RegisterPhishingProtectionServiceV1Beta1Server(s *grpc.Server, srv PhishingProtectionServiceV1Beta1Server) {
	s.RegisterService(&_PhishingProtectionServiceV1Beta1_serviceDesc, srv)
}

func _PhishingProtectionServiceV1Beta1_ReportPhishing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportPhishingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhishingProtectionServiceV1Beta1Server).ReportPhishing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.phishingprotection.v1beta1.PhishingProtectionServiceV1Beta1/ReportPhishing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhishingProtectionServiceV1Beta1Server).ReportPhishing(ctx, req.(*ReportPhishingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PhishingProtectionServiceV1Beta1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.phishingprotection.v1beta1.PhishingProtectionServiceV1Beta1",
	HandlerType: (*PhishingProtectionServiceV1Beta1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportPhishing",
			Handler:    _PhishingProtectionServiceV1Beta1_ReportPhishing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/phishingprotection/v1beta1/phishingprotection.proto",
}
