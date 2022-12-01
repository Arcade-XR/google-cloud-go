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
// source: google/cloud/recommendationengine/v1beta1/prediction_service.proto

package recommendationenginepb

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
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for Predict method.
type PredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Full resource name of the format:
	// `{name=projects/*/locations/global/catalogs/default_catalog/eventStores/default_event_store/placements/*}`
	// The id of the recommendation engine placement. This id is used to identify
	// the set of models that will be used to make the prediction.
	//
	// We currently support three placements with the following IDs by default:
	//
	//   - `shopping_cart`: Predicts items frequently bought together with one or
	//     more catalog items in the same shopping session. Commonly displayed after
	//     `add-to-cart` events, on product detail pages, or on the shopping cart
	//     page.
	//
	//   - `home_page`: Predicts the next product that a user will most likely
	//     engage with or purchase based on the shopping or viewing history of the
	//     specified `userId` or `visitorId`. For example - Recommendations for you.
	//
	//   - `product_detail`: Predicts the next product that a user will most likely
	//     engage with or purchase. The prediction is based on the shopping or
	//     viewing history of the specified `userId` or `visitorId` and its
	//     relevance to a specified `CatalogItem`. Typically used on product detail
	//     pages. For example - More items like this.
	//
	//   - `recently_viewed_default`: Returns up to 75 items recently viewed by the
	//     specified `userId` or `visitorId`, most recent ones first. Returns
	//     nothing if neither of them has viewed any items yet. For example -
	//     Recently viewed.
	//
	// The full list of available placements can be seen at
	// https://console.cloud.google.com/recommendation/datafeeds/default_catalog/dashboard
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Context about the user, what they are looking at and what action
	// they took to trigger the predict request. Note that this user event detail
	// won't be ingested to userEvent logs. Thus, a separate userEvent write
	// request is required for event logging.
	UserEvent *UserEvent `protobuf:"bytes,2,opt,name=user_event,json=userEvent,proto3" json:"user_event,omitempty"`
	// Optional. Maximum number of results to return per page. Set this property
	// to the number of prediction results required. If zero, the service will
	// choose a reasonable default.
	PageSize int32 `protobuf:"varint,7,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The previous PredictResponse.next_page_token.
	PageToken string `protobuf:"bytes,8,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. Filter for restricting prediction results. Accepts values for
	// tags and the `filterOutOfStockItems` flag.
	//
	//   - Tag expressions. Restricts predictions to items that match all of the
	//     specified tags. Boolean operators `OR` and `NOT` are supported if the
	//     expression is enclosed in parentheses, and must be separated from the
	//     tag values by a space. `-"tagA"` is also supported and is equivalent to
	//     `NOT "tagA"`. Tag values must be double quoted UTF-8 encoded strings
	//     with a size limit of 1 KiB.
	//
	//   - filterOutOfStockItems. Restricts predictions to items that do not have a
	//     stockState value of OUT_OF_STOCK.
	//
	// Examples:
	//
	//   - tag=("Red" OR "Blue") tag="New-Arrival" tag=(NOT "promotional")
	//   - filterOutOfStockItems  tag=(-"promotional")
	//   - filterOutOfStockItems
	Filter string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Use dryRun mode for this prediction query. If set to true, a
	// dummy model will be used that returns arbitrary catalog items.
	// Note that the dryRun mode should only be used for testing the API, or if
	// the model is not ready.
	DryRun bool `protobuf:"varint,4,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	// Optional. Additional domain specific parameters for the predictions.
	//
	// Allowed values:
	//
	//   - `returnCatalogItem`: Boolean. If set to true, the associated catalogItem
	//     object will be returned in the
	//     `PredictResponse.PredictionResult.itemMetadata` object in the method
	//     response.
	//   - `returnItemScore`: Boolean. If set to true, the prediction 'score'
	//     corresponding to each returned item will be set in the `metadata`
	//     field in the prediction response. The given 'score' indicates the
	//     probability of an item being clicked/purchased given the user's context
	//     and history.
	Params map[string]*structpb.Value `protobuf:"bytes,6,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. The labels for the predict request.
	//
	//   - Label keys can contain lowercase letters, digits and hyphens, must start
	//     with a letter, and must end with a letter or digit.
	//   - Non-zero label values can contain lowercase letters, digits and hyphens,
	//     must start with a letter, and must end with a letter or digit.
	//   - No more than 64 labels can be associated with a given request.
	//
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	Labels map[string]string `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PredictRequest) Reset() {
	*x = PredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictRequest) ProtoMessage() {}

func (x *PredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictRequest.ProtoReflect.Descriptor instead.
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDescGZIP(), []int{0}
}

func (x *PredictRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PredictRequest) GetUserEvent() *UserEvent {
	if x != nil {
		return x.UserEvent
	}
	return nil
}

func (x *PredictRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PredictRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *PredictRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *PredictRequest) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

func (x *PredictRequest) GetParams() map[string]*structpb.Value {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *PredictRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Response message for predict method.
type PredictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of recommended items. The order represents the ranking (from the
	// most relevant item to the least).
	Results []*PredictResponse_PredictionResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// A unique recommendation token. This should be included in the user event
	// logs resulting from this recommendation, which enables accurate attribution
	// of recommendation model performance.
	RecommendationToken string `protobuf:"bytes,2,opt,name=recommendation_token,json=recommendationToken,proto3" json:"recommendation_token,omitempty"`
	// IDs of items in the request that were missing from the catalog.
	ItemsMissingInCatalog []string `protobuf:"bytes,3,rep,name=items_missing_in_catalog,json=itemsMissingInCatalog,proto3" json:"items_missing_in_catalog,omitempty"`
	// True if the dryRun property was set in the request.
	DryRun bool `protobuf:"varint,4,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	// Additional domain specific prediction response metadata.
	Metadata map[string]*structpb.Value `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If empty, the list is complete. If nonempty, the token to pass to the next
	// request's PredictRequest.page_token.
	NextPageToken string `protobuf:"bytes,6,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *PredictResponse) Reset() {
	*x = PredictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResponse) ProtoMessage() {}

func (x *PredictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResponse.ProtoReflect.Descriptor instead.
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDescGZIP(), []int{1}
}

func (x *PredictResponse) GetResults() []*PredictResponse_PredictionResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PredictResponse) GetRecommendationToken() string {
	if x != nil {
		return x.RecommendationToken
	}
	return ""
}

func (x *PredictResponse) GetItemsMissingInCatalog() []string {
	if x != nil {
		return x.ItemsMissingInCatalog
	}
	return nil
}

func (x *PredictResponse) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

func (x *PredictResponse) GetMetadata() map[string]*structpb.Value {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PredictResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// PredictionResult represents the recommendation prediction results.
type PredictResponse_PredictionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the recommended catalog item
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Additional item metadata / annotations.
	//
	// Possible values:
	//
	//   - `catalogItem`: JSON representation of the catalogItem. Will be set if
	//     `returnCatalogItem` is set to true in `PredictRequest.params`.
	//   - `score`: Prediction score in double value. Will be set if
	//     `returnItemScore` is set to true in `PredictRequest.params`.
	ItemMetadata map[string]*structpb.Value `protobuf:"bytes,2,rep,name=item_metadata,json=itemMetadata,proto3" json:"item_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PredictResponse_PredictionResult) Reset() {
	*x = PredictResponse_PredictionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResponse_PredictionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResponse_PredictionResult) ProtoMessage() {}

func (x *PredictResponse_PredictionResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResponse_PredictionResult.ProtoReflect.Descriptor instead.
func (*PredictResponse_PredictionResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PredictResponse_PredictionResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PredictResponse_PredictionResult) GetItemMetadata() map[string]*structpb.Value {
	if x != nil {
		return x.ItemMetadata
	}
	return nil
}

var File_google_cloud_recommendationengine_v1beta1_prediction_service_proto protoreflect.FileDescriptor

var file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x05, 0x0a,
	0x0e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x49, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xe0,
	0x41, 0x02, 0xfa, 0x41, 0x2f, 0x0a, 0x2d, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x58, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f, 0x72,
	0x75, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x64,
	0x72, 0x79, 0x52, 0x75, 0x6e, 0x12, 0x62, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x62, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x51, 0x0a,
	0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe3, 0x05, 0x0a, 0x0f,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x65, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x18, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x12, 0x64, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x80, 0x02, 0x0a, 0x10, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x82,
	0x01, 0x0a, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x5d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x57, 0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x53, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x32, 0xe2, 0x02, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xf3, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x71, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x59, 0x22, 0x54, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x2a, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x2f, 0x2a, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d,
	0x3a, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x0f, 0x6e, 0x61,
	0x6d, 0x65, 0x2c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x57, 0xca,
	0x41, 0x23, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x9f, 0x02, 0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x5d, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x3b, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0xa2, 0x02, 0x05, 0x52, 0x45, 0x43, 0x41,
	0x49, 0xaa, 0x02, 0x29, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x29,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x2c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDescOnce sync.Once
	file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDescData = file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDesc
)

func file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDescGZIP() []byte {
	file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDescData)
	})
	return file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDescData
}

var file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_goTypes = []interface{}{
	(*PredictRequest)(nil),                   // 0: google.cloud.recommendationengine.v1beta1.PredictRequest
	(*PredictResponse)(nil),                  // 1: google.cloud.recommendationengine.v1beta1.PredictResponse
	nil,                                      // 2: google.cloud.recommendationengine.v1beta1.PredictRequest.ParamsEntry
	nil,                                      // 3: google.cloud.recommendationengine.v1beta1.PredictRequest.LabelsEntry
	(*PredictResponse_PredictionResult)(nil), // 4: google.cloud.recommendationengine.v1beta1.PredictResponse.PredictionResult
	nil,                                      // 5: google.cloud.recommendationengine.v1beta1.PredictResponse.MetadataEntry
	nil,                                      // 6: google.cloud.recommendationengine.v1beta1.PredictResponse.PredictionResult.ItemMetadataEntry
	(*UserEvent)(nil),                        // 7: google.cloud.recommendationengine.v1beta1.UserEvent
	(*structpb.Value)(nil),                   // 8: google.protobuf.Value
}
var file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_depIdxs = []int32{
	7,  // 0: google.cloud.recommendationengine.v1beta1.PredictRequest.user_event:type_name -> google.cloud.recommendationengine.v1beta1.UserEvent
	2,  // 1: google.cloud.recommendationengine.v1beta1.PredictRequest.params:type_name -> google.cloud.recommendationengine.v1beta1.PredictRequest.ParamsEntry
	3,  // 2: google.cloud.recommendationengine.v1beta1.PredictRequest.labels:type_name -> google.cloud.recommendationengine.v1beta1.PredictRequest.LabelsEntry
	4,  // 3: google.cloud.recommendationengine.v1beta1.PredictResponse.results:type_name -> google.cloud.recommendationengine.v1beta1.PredictResponse.PredictionResult
	5,  // 4: google.cloud.recommendationengine.v1beta1.PredictResponse.metadata:type_name -> google.cloud.recommendationengine.v1beta1.PredictResponse.MetadataEntry
	8,  // 5: google.cloud.recommendationengine.v1beta1.PredictRequest.ParamsEntry.value:type_name -> google.protobuf.Value
	6,  // 6: google.cloud.recommendationengine.v1beta1.PredictResponse.PredictionResult.item_metadata:type_name -> google.cloud.recommendationengine.v1beta1.PredictResponse.PredictionResult.ItemMetadataEntry
	8,  // 7: google.cloud.recommendationengine.v1beta1.PredictResponse.MetadataEntry.value:type_name -> google.protobuf.Value
	8,  // 8: google.cloud.recommendationengine.v1beta1.PredictResponse.PredictionResult.ItemMetadataEntry.value:type_name -> google.protobuf.Value
	0,  // 9: google.cloud.recommendationengine.v1beta1.PredictionService.Predict:input_type -> google.cloud.recommendationengine.v1beta1.PredictRequest
	1,  // 10: google.cloud.recommendationengine.v1beta1.PredictionService.Predict:output_type -> google.cloud.recommendationengine.v1beta1.PredictResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_init() }
func file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_init() {
	if File_google_cloud_recommendationengine_v1beta1_prediction_service_proto != nil {
		return
	}
	file_google_cloud_recommendationengine_v1beta1_user_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictRequest); i {
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
		file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictResponse); i {
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
		file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictResponse_PredictionResult); i {
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
			RawDescriptor: file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_msgTypes,
	}.Build()
	File_google_cloud_recommendationengine_v1beta1_prediction_service_proto = out.File
	file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_rawDesc = nil
	file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_goTypes = nil
	file_google_cloud_recommendationengine_v1beta1_prediction_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Makes a recommendation prediction. If using API Key based authentication,
	// the API Key must be registered using the
	// [PredictionApiKeyRegistry][google.cloud.recommendationengine.v1beta1.PredictionApiKeyRegistry]
	// service. [Learn more](/recommendations-ai/docs/setting-up#register-key).
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
}

type predictionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictionServiceClient(cc grpc.ClientConnInterface) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Makes a recommendation prediction. If using API Key based authentication,
	// the API Key must be registered using the
	// [PredictionApiKeyRegistry][google.cloud.recommendationengine.v1beta1.PredictionApiKeyRegistry]
	// service. [Learn more](/recommendations-ai/docs/setting-up#register-key).
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
}

// UnimplementedPredictionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPredictionServiceServer struct {
}

func (*UnimplementedPredictionServiceServer) Predict(context.Context, *PredictRequest) (*PredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.recommendationengine.v1beta1.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/recommendationengine/v1beta1/prediction_service.proto",
}
