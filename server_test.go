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

func TestServer_ParseXML(t *testing.T) {

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
  <sign><![CDATA[93F6320A66286FEAA85E7BB1B551FA00]]></sign>
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
		wantParams OrderParams
		wantErr    bool
	}{
		{name: "case1", args: args{request: mockRequest}, wantParams: OrderParams{TransactionID: "1004400740201409030005092168"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParams, err := MockWXPay.Server.ParseXML(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotParams.TransactionID, tt.wantParams.TransactionID) {
				t.Errorf("ParseXML() gotParams = %v, want %v", gotParams, tt.wantParams)
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
			err := MockWXPay.Server.Response(tt.args.writer, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Response() error = %v, wantErr %v", err, tt.wantErr)
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
				t.Errorf("Response() response.ReturnCode != \"SUCCESS\" but %s", response.ReturnCode)
			}
		})
	}
}
