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

package wxpay

import (
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/fastwego/wxpay/types"
)

var MockWXPay *WXPay
var MockSvr *httptest.Server
var MockSvrHandler *http.ServeMux
var onceSetup sync.Once

func TestMain(m *testing.M) {
	onceSetup.Do(func() {
		MockWXPay = New(Config{
			Appid:  "APPID",
			Mchid:  "MCHID",
			ApiKey: "KEY",
		})

		// Mock Server
		MockSvrHandler = http.NewServeMux()
		MockSvr = httptest.NewServer(MockSvrHandler)
		WXPayServerUrl = MockSvr.URL // 拦截发往微信服务器的请求
	})
	os.Exit(m.Run())
}

func TestWXPay_Sign(t *testing.T) {
	type args struct {
		params   map[string]string
		signType string
	}

	params := map[string]string{
		"appid":       "wxd930ea5d5a258f4f",
		"mch_id":      "10000100",
		"device_info": "1000",
		"nonce_str":   "ibuaiVcKdpRxkhJA",
		"body":        "test",
	}
	tests := []struct {
		name     string
		args     args
		wantSign string
		wantErr  bool
	}{
		{name: "case1", args: args{params: params, signType: types.SignTypeMD5}, wantSign: "9A0A8659F005D6984697E2CA0A9CF3B7", wantErr: false},
		{name: "case2", args: args{params: params, signType: types.SignTypeHMACSHA256}, wantSign: "6A9AE1657590FD6257D693A078E1C3E4BB6BA4DC30B23E0EE2496E54170DACD6", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wxpay := &WXPay{
				Config: Config{
					ApiKey: "192006250b4c09247ec02edce69f6a2d",
				},
			}
			gotSign, err := wxpay.Sign(tt.args.params, tt.args.signType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSign != tt.wantSign {
				t.Errorf("Sign() gotSign = %v, want %v", gotSign, tt.wantSign)
			}
		})
	}
}
