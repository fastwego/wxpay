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
添加分账接收方

商户发起添加分账接收方请求，后续可通过发起分账请求将结算后的钱分到该分账接收方。

See: https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_3&index=4

POST https://api.mch.weixin.qq.com/pay/profitsharingaddreceiver
*/
func ProfitSharingAddReceiver(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/pay/profitsharingaddreceiver", params, false)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
