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
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

/*
微信 api 服务器地址
*/
var WXPayServerUrl = "https://api.mch.weixin.qq.com"

/*
HttpClient 用于向公众号接口发送请求
*/
type Client struct {
	Ctx *WXPay
}

// HTTPGet GET 请求
func (client *Client) HTTPGet(uri string) (resp []byte, err error) {
	if err != nil {
		return
	}
	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("GET %s", uri)
	}
	response, err := http.Get(WXPayServerUrl + uri)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return responseFilter(response)
}

//HTTPPost POST 请求
func (client *Client) HTTPPost(uri string, payload io.Reader, contentType string) (resp []byte, err error) {
	if err != nil {
		return
	}
	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("POST %s", uri)
	}
	response, err := http.Post(WXPayServerUrl+uri, contentType, payload)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return responseFilter(response)
}

/*
筛查微信 api 服务器响应，判断以下错误：

- http 状态码 不为 200

- 接口响应错误码 errcode 不为 0
*/
func responseFilter(response *http.Response) (resp []byte, err error) {
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("Status %s", response.Status)
		return
	}

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	errorResponse := struct {
		ReturnCode string `xml:"return_code"`
		ReturnMsg  string `xml:"return_msg"`
	}{}
	err = xml.Unmarshal(resp, &errorResponse)
	if err != nil {
		return
	}

	if errorResponse.ReturnCode != "SUCCESS" {
		err = errors.New(errorResponse.ReturnMsg)
		return
	}

	return
}
