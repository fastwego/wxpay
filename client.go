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
	"crypto/tls"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fastwego/wxpay/util"

	"golang.org/x/crypto/pkcs12"
)

/*
微信 api 服务器地址
*/
var WXPayServerUrl = "https://api.mch.weixin.qq.com"

// 获取RSA公钥API
var RSAServerUrl = "https://fraud.mch.weixin.qq.com"

/*
HttpClient 用于向接口发送请求
*/
type Client struct {
	Ctx *WXPay
}

//HTTPPost POST 请求
func (client *Client) HTTPPost(uri string, params map[string]string, isCertificate bool) (resp []byte, err error) {

	// 签名
	sign, err := client.Ctx.Sign(params, params["sign_type"])
	if err != nil {
		return
	}
	params["sign"] = sign

	body, err := util.Map2XML(params)
	if err != nil {
		return
	}

	// 特殊处理 rsa 接口服务器地址
	ServerUrl := WXPayServerUrl
	if uri == "/risk/getpublickey" {
		ServerUrl = RSAServerUrl
	}

	// 是否沙箱
	url := ServerUrl + uri
	if client.Ctx.Config.IsSandboxMode {
		url = ServerUrl + "/sandboxnew" + uri
	}
	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("POST %s %s\n", url, string(body))
	}

	// 是否需要证书
	var httpClient *http.Client
	if isCertificate && client.Ctx.Config.Cert != "" {
		httpClient, err = client.getHttpsClient()
		if err != nil {
			return
		}
	} else {
		httpClient = http.DefaultClient
	}

	response, err := httpClient.Post(url, "application/xml;charset=utf-8", bytes.NewReader(body))
	if err != nil {
		return
	}
	defer response.Body.Close()
	return client.responseFilter(response)
}

// 双向证书
func (client *Client) getHttpsClient() (c *http.Client, err error) {
	certfile, err := ioutil.ReadFile(client.Ctx.Config.Cert)
	if err != nil {
		return
	}

	blocks, err := pkcs12.ToPEM(certfile, client.Ctx.Config.Mchid)
	if err != nil {
		return
	}
	var pemfile []byte
	for _, b := range blocks {
		pemfile = append(pemfile, pem.EncodeToMemory(b)...)
	}
	cert, err := tls.X509KeyPair(pemfile, pemfile)
	if err != nil {
		return
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tr := &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
	}
	return &http.Client{
		Transport: tr,
	}, nil
}

/*
筛查微信 api 服务器响应，判断以下错误：

- http 状态码 不为 200

- 接口响应错误码 ReturnCode/ResultCode == "FAIL"
*/
func (client *Client) responseFilter(response *http.Response) (resp []byte, err error) {
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("Status %s", response.Status)
		return
	}

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("response %s", string(resp))
	}

	// 对于 下载类 接口 不返回 xml 内容
	if bytes.HasPrefix(resp, []byte("<xml>")) {
		errorResponse := struct {
			ReturnCode string `xml:"return_code"`
			ResultCode string `xml:"result_code"`
		}{}
		err = xml.Unmarshal(resp, &errorResponse)
		if err != nil {
			return
		}

		if errorResponse.ReturnCode == "FAIL" || errorResponse.ResultCode == "FAIL" {
			err = errors.New(string(resp))
			return
		}
	}
	return
}
