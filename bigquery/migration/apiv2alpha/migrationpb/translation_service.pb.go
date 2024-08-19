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
// source: google/cloud/bigquery/migration/v2alpha/translation_service.proto

package migrationpb

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

// Supported SQL translation source dialects.
type TranslateQueryRequest_SqlTranslationSourceDialect int32

const (
	// SqlTranslationSourceDialect not specified.
	TranslateQueryRequest_SQL_TRANSLATION_SOURCE_DIALECT_UNSPECIFIED TranslateQueryRequest_SqlTranslationSourceDialect = 0
	// Teradata SQL.
	TranslateQueryRequest_TERADATA TranslateQueryRequest_SqlTranslationSourceDialect = 1
)

// Enum value maps for TranslateQueryRequest_SqlTranslationSourceDialect.
var (
	TranslateQueryRequest_SqlTranslationSourceDialect_name = map[int32]string{
		0: "SQL_TRANSLATION_SOURCE_DIALECT_UNSPECIFIED",
		1: "TERADATA",
	}
	TranslateQueryRequest_SqlTranslationSourceDialect_value = map[string]int32{
		"SQL_TRANSLATION_SOURCE_DIALECT_UNSPECIFIED": 0,
		"TERADATA": 1,
	}
)

func (x TranslateQueryRequest_SqlTranslationSourceDialect) Enum() *TranslateQueryRequest_SqlTranslationSourceDialect {
	p := new(TranslateQueryRequest_SqlTranslationSourceDialect)
	*p = x
	return p
}

func (x TranslateQueryRequest_SqlTranslationSourceDialect) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TranslateQueryRequest_SqlTranslationSourceDialect) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_enumTypes[0].Descriptor()
}

func (TranslateQueryRequest_SqlTranslationSourceDialect) Type() protoreflect.EnumType {
	return &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_enumTypes[0]
}

func (x TranslateQueryRequest_SqlTranslationSourceDialect) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TranslateQueryRequest_SqlTranslationSourceDialect.Descriptor instead.
func (TranslateQueryRequest_SqlTranslationSourceDialect) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescGZIP(), []int{0, 0}
}

// The error type of the SQL translation job.
type SqlTranslationError_SqlTranslationErrorType int32

const (
	// SqlTranslationErrorType not specified.
	SqlTranslationError_SQL_TRANSLATION_ERROR_TYPE_UNSPECIFIED SqlTranslationError_SqlTranslationErrorType = 0
	// Failed to parse the input text as a SQL query.
	SqlTranslationError_SQL_PARSE_ERROR SqlTranslationError_SqlTranslationErrorType = 1
	// Found unsupported functions in the input SQL query that are not able to
	// translate.
	SqlTranslationError_UNSUPPORTED_SQL_FUNCTION SqlTranslationError_SqlTranslationErrorType = 2
)

// Enum value maps for SqlTranslationError_SqlTranslationErrorType.
var (
	SqlTranslationError_SqlTranslationErrorType_name = map[int32]string{
		0: "SQL_TRANSLATION_ERROR_TYPE_UNSPECIFIED",
		1: "SQL_PARSE_ERROR",
		2: "UNSUPPORTED_SQL_FUNCTION",
	}
	SqlTranslationError_SqlTranslationErrorType_value = map[string]int32{
		"SQL_TRANSLATION_ERROR_TYPE_UNSPECIFIED": 0,
		"SQL_PARSE_ERROR":                        1,
		"UNSUPPORTED_SQL_FUNCTION":               2,
	}
)

func (x SqlTranslationError_SqlTranslationErrorType) Enum() *SqlTranslationError_SqlTranslationErrorType {
	p := new(SqlTranslationError_SqlTranslationErrorType)
	*p = x
	return p
}

func (x SqlTranslationError_SqlTranslationErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SqlTranslationError_SqlTranslationErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_enumTypes[1].Descriptor()
}

func (SqlTranslationError_SqlTranslationErrorType) Type() protoreflect.EnumType {
	return &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_enumTypes[1]
}

func (x SqlTranslationError_SqlTranslationErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SqlTranslationError_SqlTranslationErrorType.Descriptor instead.
func (SqlTranslationError_SqlTranslationErrorType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescGZIP(), []int{3, 0}
}

// The request of translating a SQL query to Standard SQL.
type TranslateQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the project to which this translation request belongs.
	// Example: `projects/foo/locations/bar`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The source SQL dialect of `queries`.
	SourceDialect TranslateQueryRequest_SqlTranslationSourceDialect `protobuf:"varint,2,opt,name=source_dialect,json=sourceDialect,proto3,enum=google.cloud.bigquery.migration.v2alpha.TranslateQueryRequest_SqlTranslationSourceDialect" json:"source_dialect,omitempty"`
	// Required. The query to be translated.
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *TranslateQueryRequest) Reset() {
	*x = TranslateQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateQueryRequest) ProtoMessage() {}

func (x *TranslateQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateQueryRequest.ProtoReflect.Descriptor instead.
func (*TranslateQueryRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescGZIP(), []int{0}
}

func (x *TranslateQueryRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *TranslateQueryRequest) GetSourceDialect() TranslateQueryRequest_SqlTranslationSourceDialect {
	if x != nil {
		return x.SourceDialect
	}
	return TranslateQueryRequest_SQL_TRANSLATION_SOURCE_DIALECT_UNSPECIFIED
}

func (x *TranslateQueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// The response of translating a SQL query to Standard SQL.
type TranslateQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Immutable. The unique identifier for the SQL translation job.
	// Example: `projects/123/locations/us/translation/1234`
	TranslationJob string `protobuf:"bytes,4,opt,name=translation_job,json=translationJob,proto3" json:"translation_job,omitempty"`
	// The translated result. This will be empty if the translation fails.
	TranslatedQuery string `protobuf:"bytes,1,opt,name=translated_query,json=translatedQuery,proto3" json:"translated_query,omitempty"`
	// The list of errors encountered during the translation, if present.
	Errors []*SqlTranslationError `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	// The list of warnings encountered during the translation, if present,
	// indicates non-semantically correct translation.
	Warnings []*SqlTranslationWarning `protobuf:"bytes,3,rep,name=warnings,proto3" json:"warnings,omitempty"`
}

func (x *TranslateQueryResponse) Reset() {
	*x = TranslateQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateQueryResponse) ProtoMessage() {}

func (x *TranslateQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateQueryResponse.ProtoReflect.Descriptor instead.
func (*TranslateQueryResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescGZIP(), []int{1}
}

func (x *TranslateQueryResponse) GetTranslationJob() string {
	if x != nil {
		return x.TranslationJob
	}
	return ""
}

func (x *TranslateQueryResponse) GetTranslatedQuery() string {
	if x != nil {
		return x.TranslatedQuery
	}
	return ""
}

func (x *TranslateQueryResponse) GetErrors() []*SqlTranslationError {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *TranslateQueryResponse) GetWarnings() []*SqlTranslationWarning {
	if x != nil {
		return x.Warnings
	}
	return nil
}

// Structured error object capturing the error message and the location in the
// source text where the error occurs.
type SqlTranslationErrorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the row from the source text where the error occurred.
	Row int64 `protobuf:"varint,1,opt,name=row,proto3" json:"row,omitempty"`
	// Specifie the column from the source texts where the error occurred.
	Column int64 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	// A human-readable description of the error.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SqlTranslationErrorDetail) Reset() {
	*x = SqlTranslationErrorDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlTranslationErrorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlTranslationErrorDetail) ProtoMessage() {}

func (x *SqlTranslationErrorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlTranslationErrorDetail.ProtoReflect.Descriptor instead.
func (*SqlTranslationErrorDetail) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescGZIP(), []int{2}
}

func (x *SqlTranslationErrorDetail) GetRow() int64 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *SqlTranslationErrorDetail) GetColumn() int64 {
	if x != nil {
		return x.Column
	}
	return 0
}

func (x *SqlTranslationErrorDetail) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The detailed error object if the SQL translation job fails.
type SqlTranslationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of SQL translation error.
	ErrorType SqlTranslationError_SqlTranslationErrorType `protobuf:"varint,1,opt,name=error_type,json=errorType,proto3,enum=google.cloud.bigquery.migration.v2alpha.SqlTranslationError_SqlTranslationErrorType" json:"error_type,omitempty"`
	// Specifies the details of the error, including the error message and
	// location from the source text.
	ErrorDetail *SqlTranslationErrorDetail `protobuf:"bytes,2,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
}

func (x *SqlTranslationError) Reset() {
	*x = SqlTranslationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlTranslationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlTranslationError) ProtoMessage() {}

func (x *SqlTranslationError) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlTranslationError.ProtoReflect.Descriptor instead.
func (*SqlTranslationError) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescGZIP(), []int{3}
}

func (x *SqlTranslationError) GetErrorType() SqlTranslationError_SqlTranslationErrorType {
	if x != nil {
		return x.ErrorType
	}
	return SqlTranslationError_SQL_TRANSLATION_ERROR_TYPE_UNSPECIFIED
}

func (x *SqlTranslationError) GetErrorDetail() *SqlTranslationErrorDetail {
	if x != nil {
		return x.ErrorDetail
	}
	return nil
}

// The detailed warning object if the SQL translation job is completed but not
// semantically correct.
type SqlTranslationWarning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the details of the warning, including the warning message and
	// location from the source text.
	WarningDetail *SqlTranslationErrorDetail `protobuf:"bytes,1,opt,name=warning_detail,json=warningDetail,proto3" json:"warning_detail,omitempty"`
}

func (x *SqlTranslationWarning) Reset() {
	*x = SqlTranslationWarning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlTranslationWarning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlTranslationWarning) ProtoMessage() {}

func (x *SqlTranslationWarning) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlTranslationWarning.ProtoReflect.Descriptor instead.
func (*SqlTranslationWarning) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescGZIP(), []int{4}
}

func (x *SqlTranslationWarning) GetWarningDetail() *SqlTranslationErrorDetail {
	if x != nil {
		return x.WarningDetail
	}
	return nil
}

var File_google_cloud_bigquery_migration_v2alpha_translation_service_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdb, 0x02, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41,
	0x23, 0x0a, 0x21, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x86, 0x01, 0x0a,
	0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x71, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x65, 0x63,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x69,
	0x61, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x22, 0x5b, 0x0a, 0x1b, 0x53, 0x71, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x65, 0x63, 0x74, 0x12,
	0x2e, 0x0a, 0x2a, 0x53, 0x51, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x45, 0x43,
	0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x54, 0x45, 0x52, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x22, 0xa6, 0x02,
	0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xe0, 0x41, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x54, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53,
	0x71, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x5a, 0x0a, 0x08, 0x77, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x71, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x77, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x5f, 0x0a, 0x19, 0x53, 0x71, 0x6c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xeb, 0x02, 0x0a, 0x13, 0x53, 0x71, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x73, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x54, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x71,
	0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x53, 0x71, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x65, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x71, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0b,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x78, 0x0a, 0x17, 0x53,
	0x71, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x26, 0x53, 0x51, 0x4c, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x51, 0x4c, 0x5f, 0x50, 0x41, 0x52, 0x53, 0x45, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x4e, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x51, 0x4c, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x53, 0x71, 0x6c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x69, 0x0a, 0x0e, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x53, 0x71, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0d, 0x77, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x32, 0xe3, 0x02, 0x0a, 0x15, 0x53,
	0x71, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xf3, 0x01, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0xda, 0x41, 0x1b, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x2c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x65,
	0x63, 0x74, 0x2c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x3a, 0x01,
	0x2a, 0x22, 0x37, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x54, 0xca, 0x41, 0x20, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2,
	0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x42, 0xe7, 0x01, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x42, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0xaa, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x32, 0x41, 0x6c, 0x70, 0x68, 0x61,
	0xca, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescData = file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDesc
)

func file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescData)
	})
	return file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDescData
}

var file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_goTypes = []any{
	(TranslateQueryRequest_SqlTranslationSourceDialect)(0), // 0: google.cloud.bigquery.migration.v2alpha.TranslateQueryRequest.SqlTranslationSourceDialect
	(SqlTranslationError_SqlTranslationErrorType)(0),       // 1: google.cloud.bigquery.migration.v2alpha.SqlTranslationError.SqlTranslationErrorType
	(*TranslateQueryRequest)(nil),                          // 2: google.cloud.bigquery.migration.v2alpha.TranslateQueryRequest
	(*TranslateQueryResponse)(nil),                         // 3: google.cloud.bigquery.migration.v2alpha.TranslateQueryResponse
	(*SqlTranslationErrorDetail)(nil),                      // 4: google.cloud.bigquery.migration.v2alpha.SqlTranslationErrorDetail
	(*SqlTranslationError)(nil),                            // 5: google.cloud.bigquery.migration.v2alpha.SqlTranslationError
	(*SqlTranslationWarning)(nil),                          // 6: google.cloud.bigquery.migration.v2alpha.SqlTranslationWarning
}
var file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_depIdxs = []int32{
	0, // 0: google.cloud.bigquery.migration.v2alpha.TranslateQueryRequest.source_dialect:type_name -> google.cloud.bigquery.migration.v2alpha.TranslateQueryRequest.SqlTranslationSourceDialect
	5, // 1: google.cloud.bigquery.migration.v2alpha.TranslateQueryResponse.errors:type_name -> google.cloud.bigquery.migration.v2alpha.SqlTranslationError
	6, // 2: google.cloud.bigquery.migration.v2alpha.TranslateQueryResponse.warnings:type_name -> google.cloud.bigquery.migration.v2alpha.SqlTranslationWarning
	1, // 3: google.cloud.bigquery.migration.v2alpha.SqlTranslationError.error_type:type_name -> google.cloud.bigquery.migration.v2alpha.SqlTranslationError.SqlTranslationErrorType
	4, // 4: google.cloud.bigquery.migration.v2alpha.SqlTranslationError.error_detail:type_name -> google.cloud.bigquery.migration.v2alpha.SqlTranslationErrorDetail
	4, // 5: google.cloud.bigquery.migration.v2alpha.SqlTranslationWarning.warning_detail:type_name -> google.cloud.bigquery.migration.v2alpha.SqlTranslationErrorDetail
	2, // 6: google.cloud.bigquery.migration.v2alpha.SqlTranslationService.TranslateQuery:input_type -> google.cloud.bigquery.migration.v2alpha.TranslateQueryRequest
	3, // 7: google.cloud.bigquery.migration.v2alpha.SqlTranslationService.TranslateQuery:output_type -> google.cloud.bigquery.migration.v2alpha.TranslateQueryResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_init() }
func file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_init() {
	if File_google_cloud_bigquery_migration_v2alpha_translation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TranslateQueryRequest); i {
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
		file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TranslateQueryResponse); i {
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
		file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SqlTranslationErrorDetail); i {
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
		file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SqlTranslationError); i {
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
		file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SqlTranslationWarning); i {
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
			RawDescriptor: file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_depIdxs,
		EnumInfos:         file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_enumTypes,
		MessageInfos:      file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_msgTypes,
	}.Build()
	File_google_cloud_bigquery_migration_v2alpha_translation_service_proto = out.File
	file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_rawDesc = nil
	file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_goTypes = nil
	file_google_cloud_bigquery_migration_v2alpha_translation_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SqlTranslationServiceClient is the client API for SqlTranslationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SqlTranslationServiceClient interface {
	// Translates input queries from source dialects to GoogleSQL.
	TranslateQuery(ctx context.Context, in *TranslateQueryRequest, opts ...grpc.CallOption) (*TranslateQueryResponse, error)
}

type sqlTranslationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSqlTranslationServiceClient(cc grpc.ClientConnInterface) SqlTranslationServiceClient {
	return &sqlTranslationServiceClient{cc}
}

func (c *sqlTranslationServiceClient) TranslateQuery(ctx context.Context, in *TranslateQueryRequest, opts ...grpc.CallOption) (*TranslateQueryResponse, error) {
	out := new(TranslateQueryResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.bigquery.migration.v2alpha.SqlTranslationService/TranslateQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqlTranslationServiceServer is the server API for SqlTranslationService service.
type SqlTranslationServiceServer interface {
	// Translates input queries from source dialects to GoogleSQL.
	TranslateQuery(context.Context, *TranslateQueryRequest) (*TranslateQueryResponse, error)
}

// UnimplementedSqlTranslationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSqlTranslationServiceServer struct {
}

func (*UnimplementedSqlTranslationServiceServer) TranslateQuery(context.Context, *TranslateQueryRequest) (*TranslateQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranslateQuery not implemented")
}

func RegisterSqlTranslationServiceServer(s *grpc.Server, srv SqlTranslationServiceServer) {
	s.RegisterService(&_SqlTranslationService_serviceDesc, srv)
}

func _SqlTranslationService_TranslateQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqlTranslationServiceServer).TranslateQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.bigquery.migration.v2alpha.SqlTranslationService/TranslateQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqlTranslationServiceServer).TranslateQuery(ctx, req.(*TranslateQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SqlTranslationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.bigquery.migration.v2alpha.SqlTranslationService",
	HandlerType: (*SqlTranslationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TranslateQuery",
			Handler:    _SqlTranslationService_TranslateQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/bigquery/migration/v2alpha/translation_service.proto",
}
