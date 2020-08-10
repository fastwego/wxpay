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
拉取订单评价数据

商户可以通过该接口拉取用户在微信支付交易记录中针对你的支付记录进行的评价内容。商户可结合商户系统逻辑对该内容数据进行存储、分析、展示、客服回访以及其他使用。如商户业务对评价内容有依赖，可主动引导用户进入微信支付交易记录进行评价。


See: https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_99&index=12

POST https://api.mch.weixin.qq.com/billcommentsp/batchquerycomment
*/
func BatchQueryComment(ctx *wxpay.WXPay, params map[string]string) (result []byte, err error) {
	return ctx.Client.HTTPPost("/billcommentsp/batchquerycomment", params, true)
}
