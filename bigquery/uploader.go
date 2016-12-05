// Copyright 2015 Google Inc. All Rights Reserved.
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

package bigquery

import (
	"fmt"
	"reflect"

	"golang.org/x/net/context"
)

// An Uploader does streaming inserts into a BigQuery table.
// It is safe for concurrent use.
type Uploader struct {
	t *Table

	// SkipInvalidRows causes rows containing invalid data to be silently
	// ignored. The default value is false, which causes the entire request to
	// fail if there is an attempt to insert an invalid row.
	SkipInvalidRows bool

	// IgnoreUnknownValues causes values not matching the schema to be ignored.
	// The default value is false, which causes records containing such values
	// to be treated as invalid records.
	IgnoreUnknownValues bool

	// A TableTemplateSuffix allows Uploaders to create tables automatically.
	//
	// Experimental: this option is experimental and may be modified or removed in future versions,
	// regardless of any other documented package stability guarantees.
	//
	// When you specify a suffix, the table you upload data to
	// will be used as a template for creating a new table, with the same schema,
	// called <table> + <suffix>.
	//
	// More information is available at
	// https://cloud.google.com/bigquery/streaming-data-into-bigquery#template-tables
	TableTemplateSuffix string
}

// Uploader returns an Uploader that can be used to append rows to t.
// The returned Uploader may optionally be further configured before its Put method is called.
func (t *Table) Uploader() *Uploader {
	return &Uploader{t: t}
}

// Put uploads one or more rows to the BigQuery service.
//
// If src is ValueSaver, then its Save method is called to produce a row for uploading.
//
// If src is a struct or pointer to a struct, then its exported fields are used
// to produce a row for uploading. The table's schema field names must match
// the struct field names case-insensitively.
//
// If src is a slice of ValueSavers, structs, or struct pointers, then each
// element of the slice is treated as above, and multiple rows are uploaded.
//
// Put returns a PutMultiError if one or more rows failed to be uploaded.
// The PutMultiError contains a RowInsertionError for each failed row.
func (u *Uploader) Put(ctx context.Context, src interface{}) error {
	if saver, ok := src.(ValueSaver); ok {
		return u.putMulti(ctx, []ValueSaver{saver})
	}
	srcVal := reflect.ValueOf(src)
	if isStructOrStructPointer(srcVal) {
		return u.putMulti(ctx, []ValueSaver{structSaver{srcVal}})
	}
	if srcVal.Kind() != reflect.Slice {
		return fmt.Errorf("%T is not a ValueSaver, struct, struct pointer, or slice", src)
	}
	var savers []ValueSaver
	for i := 0; i < srcVal.Len(); i++ {
		vi := srcVal.Index(i)
		s := vi.Interface()
		if saver, ok := s.(ValueSaver); ok {
			savers = append(savers, saver)
		} else if isStructOrStructPointer(vi) {
			savers = append(savers, structSaver{vi})
		} else {
			return fmt.Errorf("element %d of src is of type %T, which is not a ValueSaver, struct, or struct pointer", i, s)
		}
	}
	return u.putMulti(ctx, savers)
}

func isStructOrStructPointer(v reflect.Value) bool {
	return v.Kind() == reflect.Struct || (v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct)
}

func (u *Uploader) putMulti(ctx context.Context, src []ValueSaver) error {
	var rows []*insertionRow
	for _, saver := range src {
		row, insertID, err := saver.Save()
		if err != nil {
			return err
		}
		rows = append(rows, &insertionRow{InsertID: insertID, Row: row})
	}

	return u.t.c.service.insertRows(ctx, u.t.ProjectID, u.t.DatasetID, u.t.TableID, rows, &insertRowsConf{
		skipInvalidRows:     u.SkipInvalidRows,
		ignoreUnknownValues: u.IgnoreUnknownValues,
		templateSuffix:      u.TableTemplateSuffix,
	})
}

// An insertionRow represents a row of data to be inserted into a table.
type insertionRow struct {
	// If InsertID is non-empty, BigQuery will use it to de-duplicate insertions of
	// this row on a best-effort basis.
	InsertID string
	// The data to be inserted, represented as a map from field name to Value.
	Row map[string]Value
}
