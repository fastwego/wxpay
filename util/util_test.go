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
	"testing"

	"github.com/fastwego/wxpay/types"
)

func TestStructToMapByXMLTag(t *testing.T) {

	type T1 struct {
		B string `xml:"b"`
	}

	type T struct {
		XMLName  xml.Name `xml:"xml"`
		Appid    string   `xml:"appid"`               // 应用ID
		Mchid    string   `xml:"mch_id"`              // 商户号
		NonceStr string   `xml:"nonce_str"`           // 随机字符串
		Sign     string   `xml:"sign"`                // 签名
		SignType string   `xml:"sign_type,omitempty"` // 签名类型
		T1
	}

	item1 := T{
		Appid:    "100",
		SignType: types.SignTypeMD5,
	}
	item1.T1 = T1{B: "200"}

	type args struct {
		item   interface{}
		result map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{
			item:   item1,
			result: map[string]interface{}{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StructToMapByXMLTag(tt.args.item, tt.args.result)

			s, ok := tt.args.result["appid"].(string)
			if !ok {
				t.Error("Not OK")
				return
			}

			if s != "100" {
				t.Error("not equal")
				return
			}

			sign_type, ok := tt.args.result["sign_type"].(string)
			if !ok {
				t.Error("Not OK")
				return
			}

			if sign_type != types.SignTypeMD5 {
				t.Error("sign_type not equal")
				return
			}
		})
	}
}
