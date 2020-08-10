// Copyright 2020 FastWeGo
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

package util

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"
)

func TestStringMap_MarshalXML(t *testing.T) {

	tests := []struct {
		name string
		m    stringMap
		want []byte
	}{
		{name: "case1", m: stringMap{
			"key_1": "Value One",
			"key_2": "Value Two",
		}, want: []byte(`<stringMap><key_1>Value One</key_1><key_2>Value Two</key_2></stringMap>`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, _ := xml.Marshal(tt.m)
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("MarshalXML() data = %v, want %v", string(data), string(tt.want))
				return
			}

			fmt.Println(string(data))
		})
	}
}

func TestStringMap_UnmarshalXML(t *testing.T) {

	tests := []struct {
		name string
		m    []byte
		want stringMap
	}{
		{name: "case1", m: []byte(`<xml><key_1>Value One</key_1><key_2>Value Two</key_2></xml>`), want: stringMap{
			"key_1": "Value One",
			"key_2": "Value Two",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data stringMap
			_ = xml.Unmarshal(tt.m, &data)
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("MarshalXML() data = %v, want %v", data, tt.want)
				return
			}

			fmt.Println(data)
		})
	}
}

func TestMap2XML(t *testing.T) {
	type args struct {
		kvs map[string]string
	}
	tests := []struct {
		name     string
		args     args
		wantText []byte
	}{
		{name: "case1", args: args{kvs: map[string]string{
			"key_1": "Value One",
			"key_2": "Value Two",
		}}, wantText: []byte(`<xml><key_1>Value One</key_1><key_2>Value Two</key_2></xml>`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotText, _ := Map2XML(tt.args.kvs)
			if !reflect.DeepEqual(gotText, tt.wantText) {
				t.Errorf("Map2XML() = %v, want %v", string(gotText), string(tt.wantText))
			}
		})
	}
}

func TestXML2Map(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[string]string
	}{
		{name: "case1", args: args{text: []byte(`<xml><key_1>Value One</key_1><key_2>Value Two</key_2></xml>`)}, wantResult: map[string]string{
			"key_1": "Value One",
			"key_2": "Value Two",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, _ := XML2Map(tt.args.text)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("XML2Map() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
