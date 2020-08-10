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
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/fastwego/wxpay/util"
)

/*
响应微信请求 或 推送消息/事件 的服务器
*/
type Server struct {
	Ctx *WXPay
}

type OrderParams struct {
	ReturnCode         string `xml:"return_code"`
	ReturnMsg          string `xml:"return_msg"`
	Appid              string `xml:"appid" json:"appid"`
	Mchid              string `xml:"mch_id"`
	DeviceInfo         string `xml:"device_info"`
	NonceStr           string `xml:"nonce_str"`
	Sign               string `xml:"sign"`
	SignType           string `xml:"sign_type"`
	ResultCode         string `xml:"result_code"`
	ErrCode            string `xml:"err_code"`
	ErrCodeDes         string `xml:"err_code_des"`
	Openid             string `xml:"openid"`
	IsSubscribe        string `xml:"is_subscribe"`
	TradeType          string `xml:"trade_type"`
	BankType           string `xml:"bank_type"`
	TotalFee           int    `xml:"total_fee"`
	SettlementTotalFee int    `xml:"settlement_total_fee"`
	FeeType            string `xml:"fee_type"`
	CashFee            string `xml:"cash_fee"`
	CashFeeType        string `xml:"cash_fee_type"`

	CouponFee   int `xml:"coupon_fee"`
	CouponCount int `xml:"coupon_count"`

	CouponType0 string `xml:"coupon_type_0"`
	CouponID0   string `xml:"coupon_id_0"`
	CouponFee0  string `xml:"coupon_fee_0"`

	CouponType1 string `xml:"coupon_type_1"`
	CouponID1   string `xml:"coupon_id_1"`
	CouponFee1  string `xml:"coupon_fee_1"`

	CouponType2 string `xml:"coupon_type_2"`
	CouponID2   string `xml:"coupon_id_2"`
	CouponFee2  string `xml:"coupon_fee_2"`

	CouponType3 string `xml:"coupon_type_3"`
	CouponID3   string `xml:"coupon_id_3"`
	CouponFee3  string `xml:"coupon_fee_3"`

	CouponType4 string `xml:"coupon_type_4"`
	CouponID4   string `xml:"coupon_id_4"`
	CouponFee4  string `xml:"coupon_fee_4"`

	TransactionID string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}

func (s *Server) ParseXML(request *http.Request) (params OrderParams, err error) {
	var body []byte
	body, err = ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}

	if s.Ctx.Logger != nil {
		s.Ctx.Logger.Println(string(body))
	}

	err = xml.Unmarshal(body, &params)
	if err != nil {
		return
	}

	// 验证签名
	kvs := url.Values{}
	structToMap := util.StructToMap(params)
	for k, v := range structToMap {
		value, ok := v.(string)
		if ok && len(value) > 0 && k != "sign" {
			kvs.Add(k, value)
		}
	}

	sign, err := s.Ctx.Sign(kvs, kvs.Get("sign_type"))
	if err != nil {
		return
	}

	if params.Sign != sign {
		err = fmt.Errorf(" params.Sign %s != sign %s", params.Sign, sign)
		return
	}

	return
}

/*
Response 响应微信消息

<xml>
<return_code><![CDATA[SUCCESS]]></return_code>
<return_msg><![CDATA[OK]]></return_msg>
</xml>
*/
func (s *Server) Response(writer http.ResponseWriter, request *http.Request) (err error) {

	response := struct {
		XMLName    xml.Name   `xml:"xml"`
		ReturnCode util.CDATA `xml:"return_code"`
		ReturnMsg  util.CDATA `xml:"return_msg"`
	}{
		ReturnCode: "SUCCESS",
		ReturnMsg:  "OK",
	}

	output, err := xml.Marshal(response)
	if err != nil {
		return
	}

	_, err = writer.Write(output)

	if s.Ctx.Logger != nil {
		s.Ctx.Logger.Println("Response: ", string(output))
	}

	return
}
