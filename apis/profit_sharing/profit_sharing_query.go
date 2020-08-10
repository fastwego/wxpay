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
查询分账结果

发起分账请求后，可调用此接口查询分账结果；发起分账完结请求后，可调用此接口查询分账完结的执行结果。

See: https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_2&index=3

POST https://api.mch.weixin.qq.com/pay/profitsharingquery
*/
func ProfitSharingQuery(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/pay/profitsharingquery", params, false)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
