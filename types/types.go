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

package types

const (
	AccountTypeBasic     = "Basic"
	AccountTypeOperation = "Operation"
	AccountTypeFees      = "Fees"
)

const (
	FeeTypeCNY = "CNY"
)

const (
	SignTypeHMACSHA256 = "HMAC-SHA256"
	SignTypeMD5        = "MD5"
)

const (
	TradeTypeJSAPI    = "JSAPI"    //JSAPI--JSAPI支付（或小程序支付）
	TradeTypeNATIVE   = "NATIVE"   //NATIVE--Native支付
	TradeTypeAPP      = "APP"      //APP--app支付
	TradeTypeMWEB     = "MWEB"     //MWEB--H5支付，不同trade_type决定了调起支付的方式，请根据支付产品正确上传
	TradeTypeMICROPAY = "MICROPAY" //MICROPAY--付款码支付，付款码支付有单独的支付接口，所以接口不需要上传，该字段在对账单中会出现
)

const (
	BillTypeALL             = "ALL"             // ALL（默认值），返回当日所有订单信息（不含充值退款订单）
	BillTypeSUCCESS         = "SUCCESS"         // SUCCESS，返回当日成功支付的订单（不含充值退款订单）
	BillTypeREFUND          = "REFUND"          // REFUND，返回当日退款订单（不含充值退款订单）
	BillTypeRECHARGE_REFUND = "RECHARGE_REFUND" //RECHARGE_REFUND，返回当日充值退款订单
)

const (
	TarType = "GZIP"
)
