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

// Package order 订单相关
package order

import (
	"github.com/fastwego/wxpay"
	"github.com/fastwego/wxpay/util"
)

/*
查询订单

该接口提供所有微信支付订单的查询，商户可以通过该接口主动查询订单状态，完成下一步的业务逻辑。

需要调用查询接口的情况：

◆ 当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知；

◆ 调用支付接口后，返回系统错误或未知交易状态情况；

◆ 调用被扫支付API，返回USERPAYING的状态；

◆ 调用关单或撤销接口API之前，需确认支付状态；

See: https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_2&index=4

POST https://api.mch.weixin.qq.com/pay/orderquery
*/
func OrderQuery(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("/pay/orderquery", params, false)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
