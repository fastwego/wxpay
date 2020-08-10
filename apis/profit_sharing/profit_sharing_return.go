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
分账回退

● 对订单进行退款时，如果订单已经分账，可以先调用此接口将指定的金额从分账接收方（仅限商户类型的分账接收方）回退给本商户，然后再退款。

● 回退以原分账请求为依据，可以对分给分账接收方的金额进行多次回退，只要满足累计回退不超过该请求中分给接收方的金额。

● 此接口采用同步处理模式，即在接收到商户请求后，会实时返回处理结果

● 此功能需要接收方在商户平台-交易中心-分账-分账接收设置下，开启同意分账回退后，才能使用。

See: https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_7&index=7

POST https://api.mch.weixin.qq.com/secapi/pay/profitsharingreturn
*/
func ProfitSharingReturn(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/secapi/pay/profitsharingreturn", params, true)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
