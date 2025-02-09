// Copyright © 2022 Meroxa, Inc.
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

package txfbuiltin

import (
	"testing"

	"github.com/conduitio/conduit/pkg/foundation/assert"
	"github.com/conduitio/conduit/pkg/foundation/cerrors"
	"github.com/conduitio/conduit/pkg/processor/transform"
	"github.com/conduitio/conduit/pkg/record"
	"github.com/google/go-cmp/cmp"
)

func TestFilterFieldKey_Build(t *testing.T) {
	type args struct {
		config transform.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "nil config returns error",
			args: args{
				config: nil,
			},
			wantErr: true,
		},
		{
			name: "empty config returns error",
			args: args{
				config: map[string]string{},
			},
			wantErr: true,
		},
		{
			name: "empty type returns error",
			args: args{
				config: transform.Config{
					"type":      "",
					"condition": "$[key]",
				},
			},
			wantErr: true,
		},
		{
			name: "empty condition returns error",
			args: args{
				config: transform.Config{
					"type":      "include",
					"condition": "",
					"fail":      "include",
				},
			},
			wantErr: true,
		},
		{
			name: "valid config should return transform",
			args: args{
				config: transform.Config{
					"type":      "include",
					"condition": ".key",
					"fail":      "include",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := FilterFieldKey(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestFilterFieldKey_Transform(t *testing.T) {
	type args struct {
		r record.Record
	}
	tests := []struct {
		name    string
		args    args
		config  transform.Config
		want    record.Record
		wantErr bool
		err     error
	}{
		{
			name: "should forward record on condition",
			config: map[string]string{
				"type":      "include",
				"condition": ".id",
				"fail":      "include",
			},
			args: args{r: record.Record{
				Key: record.StructuredData{
					"id": "foo",
				},
			}},
			want: record.Record{
				Key: record.StructuredData{
					"id": "foo",
				},
			},
			wantErr: false,
		},
		{
			name: "should drop record on condition",
			config: map[string]string{
				"type":      "exclude",
				"condition": ".id",
				"fail":      "include",
			},
			args: args{r: record.Record{
				Key: record.StructuredData{
					"id": "foo",
				},
			}},
			want:    record.Record{},
			wantErr: true,
			err:     ErrDropRecord,
		},
		{
			name: "should handle missing or null by failing",
			config: map[string]string{
				"type":          "include",
				"condition":     "@id",
				"missingornull": "fail",
				"exists":        "id",
			},
			args: args{r: record.Record{
				Key: record.StructuredData{
					"user": "foo",
				},
			}},
			want:    record.Record{},
			wantErr: true,
			err:     cerrors.New("field does not exist: id"),
		},
		{
			name: "should handle missing or null by including",
			config: map[string]string{
				"type":          "include",
				"condition":     "@id",
				"missingornull": "include",
				"exists":        "@id",
			},
			args: args{r: record.Record{
				Key: record.StructuredData{
					"user": "foo",
				},
			}},
			want: record.Record{
				Key: record.StructuredData{
					"user": "foo",
				},
			},
			wantErr: false,
		},
		{
			name: "should handle missing or null by excluding",
			config: map[string]string{
				"type":          "include",
				"condition":     "@id",
				"missingornull": "exclude",
			},
			args: args{r: record.Record{
				Key: record.StructuredData{
					"user": "foo",
				},
			}},
			want:    record.Record{},
			wantErr: true,
			err:     ErrDropRecord,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txf, err := FilterFieldKey(tt.config)
			assert.Ok(t, err)
			got, err := txf(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterFieldKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Logf("DIFF: %+s", diff)
				t.Fail()
			}
			if tt.err != nil {
				if diff := cmp.Diff(tt.err.Error(), err.Error()); diff != "" {
					t.Errorf("DIFF: %s", diff)
				}
			}
		})
	}
}

func TestFilterFieldPayload_Build(t *testing.T) {
	type args struct {
		config transform.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "nil config returns error",
			args:    args{config: nil},
			wantErr: true,
		},
		{
			name:    "empty config returns error",
			args:    args{config: map[string]string{}},
			wantErr: true,
		},
		{
			name: "empty condition returns error",
			args: args{config: map[string]string{
				"type":          "include",
				"missingornull": "fail",
				"condition":     "",
			}},
			wantErr: true,
		},
		{
			name: "empty type returns error",
			args: args{config: map[string]string{
				"type":          "",
				"condition":     "@id",
				"missingornull": "fail",
			}},
			wantErr: true,
		},
		{
			name: "valid config returns transform",
			args: args{config: map[string]string{
				"type":      "include",
				"condition": "@id",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := FilterFieldPayload(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterFieldPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestFilterFieldPayload_Transform(t *testing.T) {
	type args struct {
		config transform.Config
		r      record.Record
	}
	tests := []struct {
		name    string
		args    args
		want    record.Record
		wantErr bool
		err     error
	}{
		{
			name: "should forward record on condition",
			args: args{
				r: record.Record{
					Payload: record.StructuredData{
						"foo": "bar",
					},
				},
				config: transform.Config{
					"type":          "include",
					"condition":     "foo",
					"missingornull": "fail",
				}},
			want: record.Record{
				Payload: record.StructuredData{
					"foo": "bar",
				},
			},
			wantErr: false,
		},
		{
			name: "should drop record on condition",
			args: args{
				r: record.Record{
					Payload: record.StructuredData{
						"foo": "5",
					},
				},
				config: transform.Config{
					"type":      "exclude",
					"condition": "foo > 1",
				}},
			want:    record.Record{},
			wantErr: true,
			err:     ErrDropRecord,
		},
		{
			name: "should drop record on missing key",
			args: args{
				r: record.Record{
					Payload: record.StructuredData{
						"bar": "3",
					},
				},
				config: transform.Config{
					"type":          "exclude",
					"condition":     "foo > 1",
					"exists":        "foo",
					"missingornull": "exclude",
				}},
			want:    record.Record{},
			wantErr: true,
			err:     ErrDropRecord,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txf, err := FilterFieldPayload(tt.args.config)
			assert.Ok(t, err)
			got, err := txf(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterFieldPayload Error: %s - wanted: %s", err, tt.err)
				return
			}
			if tt.wantErr {
				if diff := cmp.Diff(tt.err.Error(), err.Error()); diff != "" {
					t.Errorf("FilterFieldPayload() failed: [DIFF] %s", diff)
				}
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("failed: %s", diff)
			}
		})
	}
}
