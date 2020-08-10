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

// Package download 下载 相关
package download

import "github.com/fastwego/wxpay"

/*
下载交易账单

商户可以通过该接口下载历史交易清单。比如掉单、系统错误等导致商户侧和微信侧数据不一致，通过对账单核对后可校正支付状态。

注意：

1、微信侧未成功下单的交易不会出现在对账单中。支付成功后撤销的交易会出现在对账单中，跟原支付单订单号一致；

2、微信在次日9点启动生成前一天的对账单，建议商户10点后再获取；

3、对账单中涉及金额的字段单位为“元”。

4、对账单接口只能下载三个月以内的账单。

5、对账单是以商户号纬度来生成的，如一个商户号与多个appid有绑定关系，则使用其中任何一个appid都可以请求下载对账单。对账单中的appid取自交易时候提交的appid，与请求下载对账单时使用的appid无关。

6、自2018年起入驻的商户默认是开通免充值券后的结算对账单。

See: https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_6&index=8

POST https://api.mch.weixin.qq.com/pay/downloadbill
*/
func DownloadBill(ctx *wxpay.WXPay, params map[string]string) (result []byte, err error) {
	return ctx.Client.HTTPPost("/pay/downloadbill", params, false)
}
