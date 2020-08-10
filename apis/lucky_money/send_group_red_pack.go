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
发放裂变红包

裂变红包：一次可以发放一组红包。首先领取的用户为种子用户，种子用户领取一组红包当中的一个，并可以通过社交分享将剩下的红包给其他用户。裂变红包充分利用了人际传播的优势。

See: https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_5&index=4

POST https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack
*/
func SendGroupRedPack(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/mmpaymkttransfers/sendgroupredpack", params, true)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
