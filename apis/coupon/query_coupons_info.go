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
查询代金券信息



See: https://pay.weixin.qq.com/wiki/doc/api/tools/sp_coupon.php?chapter=12_5&index=6

POST https://api.mch.weixin.qq.com/mmpaymkttransfers/querycouponsinfo
*/
func QueryCouponsInfo(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/mmpaymkttransfers/querycouponsinfo", params, false)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
