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
完结分账

1、不需要进行分账的订单，可直接调用本接口将订单的金额全部解冻给本商户
2、调用多次分账接口后，需要解冻剩余资金时，调用本接口将剩余的分账金额全部解冻给特约商户
3、已调用请求单次分账后，剩余待分账金额为零，不需要再调用此接口。

See: https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_5&index=6

POST https://api.mch.weixin.qq.com/secapi/pay/profitsharingfinish
*/
func ProfitSharingFinish(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/secapi/pay/profitsharingfinish", params, true)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
