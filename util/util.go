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
	"reflect"
	"strings"
)

type CDATA string

func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

func StructToMapByXMLTag(item interface{}, result map[string]interface{}) {

	if item == nil {
		return
	}

	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("xml")

		if strings.Contains(tag, ",") {
			split := strings.Split(tag, ",")
			tag = strings.TrimSpace(split[0])
		}

		field := reflectValue.Field(i).Interface()

		if v.Field(i).Type.Kind() == reflect.Struct {
			StructToMapByXMLTag(field, result)
		} else if tag != "" && tag != "-" {
			result[tag] = field
		}
	}
}
