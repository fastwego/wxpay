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

// Package profit_sharing 分账
package profit_sharing

import (
	"github.com/fastwego/wxpay"
	"github.com/fastwego/wxpay/util"
)

/*
请求多次分账

● 微信订单支付成功后，商户发起分账请求，将结算后的钱分到分账接收方。多次分账请求仅会按照传入的分账接收方进行分账，不会对剩余的金额进行任何操作。故操作成功后，在待分账金额不等于零时，订单依旧能够再次进行分账。

● 多次分账，可以将本商户作为分账接收方直接传入，实现释放资金给本商户的功能

● 对同一笔订单最多能发起20次多次分账请求



See: https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_6&index=2

POST https://api.mch.weixin.qq.com/secapi/pay/multiprofitsharing
*/
func MultiProfitSharing(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/secapi/pay/multiprofitsharing", params, true)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
