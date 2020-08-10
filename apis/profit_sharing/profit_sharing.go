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
请求单次分账

单次分账请求按照传入的分账接收方账号和资金进行分账，同时会将订单剩余的待分账金额解冻给本商户。故操作成功后，订单不能再进行分账，也不能进行分账完结。

See: https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1

POST https://api.mch.weixin.qq.com/secapi/pay/profitsharing
*/
func ProfitSharing(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/secapi/pay/profitsharing", params, true)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
