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

// Package micropay 付款码支付
package micropay

import (
	"github.com/fastwego/wxpay"
	"github.com/fastwego/wxpay/util"
)

/*
付款码查询openid

通过付款码查询公众号Openid，调用查询后，该付款码只能由此商户号发起扣款，直至付款码更新。

See: https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_13&index=9

POST https://api.mch.weixin.qq.com/tools/authcodetoopenid
*/
func AuthCodeToOpenId(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/tools/authcodetoopenid", params, false)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
