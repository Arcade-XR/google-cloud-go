// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by cloud.google.com/go/internal/gapicgen/gensnippets. DO NOT EDIT.

// [START dialogflow_v2beta1_generated_Documents_DeleteDocument_sync]

package main

import (
	"context"

	dialogflow "cloud.google.com/go/dialogflow/apiv2beta1"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2beta1"
)

func main() {
	ctx := context.Background()
	c, err := dialogflow.NewDocumentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.DeleteDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/v2beta1#DeleteDocumentRequest.
	}
	op, err := c.DeleteDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END dialogflow_v2beta1_generated_Documents_DeleteDocument_sync]
