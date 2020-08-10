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

package profit_sharing

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/fastwego/wxpay/test"
	"github.com/fastwego/wxpay/types"
	"github.com/fastwego/wxpay/util"
)

func TestProfitSharing(t *testing.T) {
	test.Setup()

	mockResp := map[string][]byte{
		"case1": []byte("<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg></xml>"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc("/secapi/pay/profitsharing", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		params map[string]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		{name: "case1", args: args{params: map[string]string{
			"appid":     test.MockWXPay.Config.Appid,
			"mch_id":    test.MockWXPay.Config.Mchid,
			"nonce_str": util.GetRandString(16),
			"sign_type": types.SignTypeHMACSHA256}}, wantResult: "SUCCESS", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResult, err := ProfitSharing(test.MockWXPay, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfitSharing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult["return_code"], tt.wantResult) {
				t.Errorf("CloseOrder() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
