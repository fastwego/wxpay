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
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/fastwego/wxpay/util"
)

func TestServer_PaymentResultNotify(t *testing.T) {

	mockRequest := httptest.NewRequest("POST", "http://example.com/foo", bytes.NewReader([]byte(`<xml>
  <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
  <attach><![CDATA[支付测试]]></attach>
  <bank_type><![CDATA[CFT]]></bank_type>
  <fee_type><![CDATA[CNY]]></fee_type>
  <is_subscribe><![CDATA[Y]]></is_subscribe>
  <mch_id><![CDATA[10000100]]></mch_id>
  <nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
  <openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
  <out_trade_no><![CDATA[1409811653]]></out_trade_no>
  <result_code><![CDATA[SUCCESS]]></result_code>
  <return_code><![CDATA[SUCCESS]]></return_code>
  <sign><![CDATA[C6F22BE0BD36EE6F7EF408F58A8D7D94]]></sign>
  <time_end><![CDATA[20140903131540]]></time_end>
  <total_fee>1</total_fee>
  <coupon_fee><![CDATA[10]]></coupon_fee>
  <coupon_count><![CDATA[1]]></coupon_count>
  <coupon_type><![CDATA[CASH]]></coupon_type>
  <coupon_id><![CDATA[10000]]></coupon_id>
  <trade_type><![CDATA[JSAPI]]></trade_type>
  <transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
</xml>`)))

	type args struct {
		request *http.Request
	}
	tests := []struct {
		name       string
		args       args
		wantParams map[string]string
		wantErr    bool
	}{
		{name: "case1", args: args{request: mockRequest}, wantParams: map[string]string{"transaction_id": "1004400740201409030005092168"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParams, err := MockWXPay.Server.PaymentNotify(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentNotify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotParams["transaction_id"], tt.wantParams["transaction_id"]) {
				t.Errorf("PaymentNotify() gotParams = %v, want %v", gotParams, tt.wantParams)
			}
		})
	}
}

func TestServer_RefundNotify(t *testing.T) {

	mockRequest := httptest.NewRequest("POST", "http://example.com/foo", bytes.NewReader([]byte(`<xml>
<return_code>SUCCESS</return_code>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[TeqClE3i0mvn3DrK]]></nonce_str>
   <req_info><![CDATA[d3nmAkAKIJxfqpROzQtEjzCQsD7Ekj2EKUBvSoAXZo/GTUYNvcCmFOhBazx2qs9dqjoSYAklfjw84CJqD0cnTuyNbPOIkLLgJ6coPz7ouVybFy5YjgNEY3FFaQ4mbxW62JfuxkkLNn3XcYWOdGc1wFgeONULDuxLwvynQUIpdQI0O1UOIPFZj1ItAm2PzsezDbSKJGknPAF7nfWmXcuiRWAs/FqhLrMR88LuF5sbGKe2ZJ6KnLm0QsB5BKJKRV6SRlNpBXMITmWk5dk8X34w1IqYehIK7lGW7M42wDpHQxgE5joJxbijKHIFtRaFa4emJxcIas0XrBXtqknJ0Zpb2+hBXLpGW+teJVmh7DQwG07vC5BahkyggH5j5K/TU/VI26UxJoHs7hHl4Iz51w7gOi6zwfge6x4IZNuKwhnfvot1ICtJDCAaVULoFAsEpHb13mGIL6l4t+2u2r27cvlaLPlDGna0aTBwy9fxMYqvwfBMwOECv5hbiCXaPMc3JSCi2M1wIZySY3/GZ8DAaMOeM2rBZvH9VPtfaMPj40tqtQaJguOGhAUVAwAD8dhXQ3Ueu+ymLVijRbCIKEqqX2WAVwcKF6nhXoSMIbMfjHd4Sy/NOJE7ABs11sBpiQ+nOdtHCa88MSINAHM0rV8tI+E+28wqHF3JAzTKs2sKMDmQpEraV0nuei8ytcc75c4Ywu2xBOY6CcW4oyhyBbUWhWuHpmCkUlh5teD6ktQNmwJcpiLZtWLmeLtVJzZKVCvuJ0ss0loQwylDIsQTw0vvqYOIwSgA5UoCDjNx/rt5p8HX3MYHVl3MEkFrkxQeRS165lr88ors3DyxVEwPc2dC3vZsRUWxQmsQFvyh1xDDoNz+hfAhNvfl3/oxv0ktKevSPAIXOb7phQ4DE6QTT66O4BgPBZMXaSbvvUF2azM1JoVfe0ooAOVKAg4zcf67eafB19zGbMYR7bxS3mkeHXYWLAnZlvu53Wl89ZNuxTTPMW3cECOic3ZEcEe3bQS2Bz6C1bEBim/fPnbJc3bbUomy21fOIiTkxjIL5UDu/AuXsTunHenKxLwUc1naBhRqoY+Sxo8L]]></req_info>
</xml>
`)))

	/*
		req_info解密后的示例：
		<root>
		<out_refund_no><![CDATA[131811191610442717309]]></out_refund_no>
		<out_trade_no><![CDATA[71106718111915575302817]]></out_trade_no>
		<refund_account><![CDATA[REFUND_SOURCE_RECHARGE_FUNDS]]></refund_account>
		<refund_fee><![CDATA[3960]]></refund_fee>
		<refund_id><![CDATA[50000408942018111907145868882]]></refund_id>
		<refund_recv_accout><![CDATA[支付用户零钱]]></refund_recv_accout>
		<refund_request_source><![CDATA[API]]></refund_request_source>
		<refund_status><![CDATA[SUCCESS]]></refund_status>
		<settlement_refund_fee><![CDATA[3960]]></settlement_refund_fee>
		<settlement_total_fee><![CDATA[3960]]></settlement_total_fee>
		<success_time><![CDATA[2018-11-19 16:24:13]]></success_time>
		<total_fee><![CDATA[3960]]></total_fee>
		<transaction_id><![CDATA[4200000215201811190261405420]]></transaction_id>
		</root>
	*/

	type args struct {
		request *http.Request
	}
	tests := []struct {
		name       string
		args       args
		wantParams map[string]string
		wantErr    bool
	}{
		{name: "case1", args: args{request: mockRequest}, wantParams: map[string]string{"transaction_id": "4200000215201811190261405420"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParams, err := MockWXPay.Server.RefundNotify(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentNotify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotParams["transaction_id"], tt.wantParams["transaction_id"]) {
				t.Errorf("PaymentNotify() gotParams = %v, want %v", gotParams, tt.wantParams)
			}
		})
	}
}

func TestServer_Response(t *testing.T) {
	mockRequest := httptest.NewRequest("POST", "http://example.com/foo", nil)
	recorder := httptest.NewRecorder()

	type args struct {
		writer  http.ResponseWriter
		request *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case1", args: args{writer: recorder, request: mockRequest}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MockWXPay.Server.ResponseSuccess(tt.args.writer, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResponseSuccess() error = %v, wantErr %v", err, tt.wantErr)
			}

			resp := recorder.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			response := struct {
				XMLName    xml.Name   `xml:"xml"`
				ReturnCode util.CDATA `xml:"return_code"`
				ReturnMsg  util.CDATA `xml:"return_msg"`
			}{}

			err = xml.Unmarshal(body, &response)

			if response.ReturnCode != "SUCCESS" {
				t.Errorf("ResponseSuccess() response.ReturnCode != \"SUCCESS\" but %s", response.ReturnCode)
			}
		})
	}
}
