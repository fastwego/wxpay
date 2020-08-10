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

// Package refund 退款 相关
package refund

import (
	"github.com/fastwego/wxpay"
	"github.com/fastwego/wxpay/util"
)

/*
退款

当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家，微信支付将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退到买家帐号上。


See: https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_4&index=6

POST https://api.mch.weixin.qq.com/secapi/pay/refund
*/
func Refund(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/secapi/pay/refund", params, true)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
