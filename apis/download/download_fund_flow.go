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
下载资金账单


商户可以通过该接口下载自2017年6月1日起 的历史资金流水账单。

说明：

1、资金账单中的数据反映的是商户微信账户资金变动情况；

2、当日账单在次日上午9点开始生成，建议商户在上午10点以后获取；

3、资金账单中涉及金额的字段单位为“元”。


See: https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_18&index=9

POST https://api.mch.weixin.qq.com/pay/downloadfundflow
*/
func DownloadFundFlow(ctx *wxpay.WXPay, params map[string]string) (result []byte, err error) {
	return ctx.Client.HTTPPost("/pay/downloadfundflow", params, true)
}
