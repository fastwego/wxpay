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

// Package coupon 代金券
package coupon

import (
	"github.com/fastwego/wxpay"
	"github.com/fastwego/wxpay/util"
)

/*
发放代金券

用于商户主动调用接口给用户发放代金券的场景，已做防小号处理，给小号发放代金券将返回错误码。

注意：通过接口发放的代金券不会进入微信卡包

See: https://pay.weixin.qq.com/wiki/doc/api/tools/sp_coupon.php?chapter=12_3&index=4

POST https://api.mch.weixin.qq.com/mmpaymkttransfers/send_coupon
*/
func SendCoupon(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/mmpaymkttransfers/send_coupon", params, false)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
