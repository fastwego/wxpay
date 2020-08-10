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

// Package corp_pay 企业付款
package corp_pay

import (
	"github.com/fastwego/wxpay"
	"github.com/fastwego/wxpay/util"
)

/*
企业付款到银行卡

企业付款业务是基于微信支付商户平台的资金管理能力，为了协助商户方便地实现企业向银行卡付款，针对部分有开发能力的商户，提供通过API完成企业付款到银行卡的功能。

See: https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2

POST https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank
*/
func PayBank(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/mmpaysptrans/pay_bank", params, true)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
