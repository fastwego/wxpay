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
查询 企业付款 零钱

用于商户的企业付款操作进行结果查询，返回付款操作详细结果。

查询企业付款API只支持查询30天内的订单，30天之前的订单请登录商户平台查询。

See: https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3

POST https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo
*/
func GetTransferInfo(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/mmpaymkttransfers/gettransferinfo", params, true)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
