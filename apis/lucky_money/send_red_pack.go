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

// Package lucky_money 现金红包
package lucky_money

import (
	"github.com/fastwego/wxpay"
	"github.com/fastwego/wxpay/util"
)

/*
发放红包接口

1.发送频率限制------默认1800/min

2.发送个数上限------默认1800/min

3.场景金额限制------默认红包金额为1-200元，如有需要，可前往商户平台进行设置和申请

4.其他限制------商户单日出资金额上限--100万元；单用户单日收款金额上限--1000元；单用户可领取红包个数上限--10个；



See: https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3

POST https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack
*/
func SendRedPack(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/mmpaymkttransfers/sendredpack", params, true)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
