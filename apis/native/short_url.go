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

// Package native Native 支付
package native

import (
	"github.com/fastwego/wxpay"
	"github.com/fastwego/wxpay/util"
)

/*
转换短链接

该接口主要用于Native支付模式一中的二维码链接转成短链接(weixin://wxpay/s/XXXXXX)，减小二维码数据量，提升扫描速度和精确度。

See: https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=9_9&index=10

POST https://api.mch.weixin.qq.com/tools/shorturl
*/
func ShortUrl(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/tools/shorturl", params, false)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
