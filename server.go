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
	"crypto/md5"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fastwego/wxpay/util"
)

/*
响应微信请求 或 推送消息/事件 的服务器
*/
type Server struct {
	Ctx *WXPay
}

/*
支付结果 回调

See: https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_7&index=3
*/
func (s *Server) PaymentNotify(request *http.Request) (params map[string]string, err error) {
	var body []byte
	body, err = ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}

	if s.Ctx.Logger != nil {
		s.Ctx.Logger.Println("PaymentNotify", string(body))
	}

	params, err = util.XML2Map(body)
	if err != nil {
		return
	}

	// 验证签名
	sign, err := s.Ctx.Sign(params, params["sign_type"])
	if err != nil {
		return
	}

	if params["sign"] != sign {
		err = fmt.Errorf(" params.Sign %s != Sign %s", params["sign"], sign)
		return
	}

	return
}

/*
退款结果 回调

See: https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_16&index=11
*/
func (s *Server) RefundNotify(request *http.Request) (params map[string]string, err error) {
	var body []byte
	body, err = ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}

	if s.Ctx.Logger != nil {
		s.Ctx.Logger.Println("RefundNotify", string(body))
	}

	encryptMsg := struct {
		XMLName  xml.Name `xml:"xml"`
		Appid    string   `xml:"appid"`
		Mchid    string   `xml:"mch_id"`
		NonceStr string   `xml:"nonce_str"`
		ReqInfo  string   `xml:"req_info"`
	}{}

	err = xml.Unmarshal(body, &encryptMsg)
	if err != nil {
		return
	}

	/*
		解密步骤如下：

		（1）对加密串A做 base64解码，得到加密串B

		（2）对商户 key做 md5，得到32位小写 key* ( key设置路径：微信商户平台(pay.weixin.qq.com)-->账户设置-->API安全-->密钥设置 )

		（3）用 key*对加密串B做AES-256-ECB解密（PKCS7Padding）`
	*/
	cipherText, err := base64.StdEncoding.DecodeString(encryptMsg.ReqInfo)
	if err != nil {
		return
	}

	key := []byte(fmt.Sprintf("%x", md5.Sum([]byte(s.Ctx.Config.ApiKey))))

	reqInfo, err := util.AESECBPKCS7Decrypt(cipherText, key)
	if err != nil {
		return
	}

	params, err = util.XML2Map(reqInfo)
	if err != nil {
		return
	}

	return
}

/*
ResponseSuccess 响应微信消息

<xml>
<return_code><![CDATA[SUCCESS]]></return_code>
<return_msg><![CDATA[OK]]></return_msg>
</xml>
*/
func (s *Server) ResponseSuccess(writer http.ResponseWriter, request *http.Request) (err error) {

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
		s.Ctx.Logger.Println("ResponseSuccess: ", string(output))
	}

	return
}
