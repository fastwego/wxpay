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

/*
Package wxpay 微信支付
*/
package wxpay

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/fastwego/wxpay/types"
)

/*
微信支付 实例
*/
type WXPay struct {
	Config Config
	Client Client
	Server Server
	Logger *log.Logger
}

/*
微信支付配置
*/
type Config struct {
	Appid         string
	Mchid         string // 商户 id
	ApiKey        string // 商户 api key
	IsSandboxMode bool   // 是否开启沙箱模式
	Cert          string // 证书路径
}

/*
创建 实例
*/
func New(config Config) (wxpay *WXPay) {
	instance := WXPay{
		Config: config,
	}

	instance.Client = Client{Ctx: &instance}
	instance.Server = Server{Ctx: &instance}

	instance.Logger = log.New(os.Stdout, "[WXPay] ", log.LstdFlags|log.Llongfile)

	return &instance
}

/*

参数签名

See: https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=4_3
*/
func (wxpay *WXPay) Sign(params map[string]string, signType string) (sign string, err error) {

	var kvs []string
	for k, v := range params {
		if len(v) > 0 && strings.ToLower(k) != "sign" {
			kvs = append(kvs, fmt.Sprintf("%s=%s", k, v))
		}
	}

	sort.Strings(kvs)

	kvs = append(kvs, fmt.Sprintf("key=%s", wxpay.Config.ApiKey))

	str := strings.Join(kvs, "&")

	var h hash.Hash
	if signType == types.SignTypeHMACSHA256 {
		h = hmac.New(sha256.New, []byte(wxpay.Config.ApiKey))
	} else {
		h = md5.New()
	}
	if _, err = h.Write([]byte(str)); err != nil {
		return
	}

	sign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))

	if wxpay.Logger != nil {
		wxpay.Logger.Printf("Sign %s(%s) = %s", signType, str, sign)
	}

	return
}

/*
SetLogger 日志记录 默认输出到 os.Stdout

可以新建 logger 输出到指定文件

如果不想开启日志，可以 SetLogger(nil)
*/
func (wxpay *WXPay) SetLogger(logger *log.Logger) {
	wxpay.Logger = logger
}
