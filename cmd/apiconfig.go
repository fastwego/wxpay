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

package main

type Api struct {
	Name        string
	Description string
	Request     string
	See         string
	FuncName    string
	IsCert      bool
}

type ApiGroup struct {
	Name    string
	Apis    []Api
	Package string
}

var apiConfig = []ApiGroup{
	{
		Name:    `订单相关`,
		Package: `order`,
		Apis: []Api{
			{
				Name:        "统一下单",
				Description: `商户系统先调用该接口在微信支付服务后台生成预支付交易单，返回正确的预支付交易会话标识后再在APP里面调起支付。`,
				Request:     "POST https://api.mch.weixin.qq.com/pay/unifiedorder",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1",
				FuncName:    "UnifiedOrder",
				IsCert:      false,
			},
			{
				Name: "查询订单",
				Description: `该接口提供所有微信支付订单的查询，商户可以通过该接口主动查询订单状态，完成下一步的业务逻辑。

需要调用查询接口的情况：

◆ 当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知；

◆ 调用支付接口后，返回系统错误或未知交易状态情况；

◆ 调用被扫支付API，返回USERPAYING的状态；

◆ 调用关单或撤销接口API之前，需确认支付状态；`,
				Request:  "POST https://api.mch.weixin.qq.com/pay/orderquery",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_2&index=4",
				FuncName: "OrderQuery",
				IsCert:   false,
			},
			{
				Name: "关闭订单",
				Description: `以下情况需要调用关单接口：商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。

注意：订单生成后不能马上调用关单接口，最短调用时间间隔为5分钟。`,
				Request:  "POST https://api.mch.weixin.qq.com/pay/closeorder",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_3&index=5",
				FuncName: "CloseOrder",
				IsCert:   false,
			},
		},
	},
	{
		Name:    `退款 相关`,
		Package: `refund`,
		Apis: []Api{
			{
				Name: "退款",
				Description: `当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家，微信支付将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退到买家帐号上。
`,
				Request:  "POST https://api.mch.weixin.qq.com/secapi/pay/refund",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_4&index=6",
				FuncName: "Refund",
				IsCert:   true,
			},
			{
				Name:        "查询退款",
				Description: `当提交退款申请后，通过调用该接口查询退款状态。退款有一定延时，用零钱支付的退款20分钟内到账，银行卡支付的退款3个工作日后重新查询退款状态。`,
				Request:     "POST https://api.mch.weixin.qq.com/pay/refundquery",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_5&index=7",
				FuncName:    "RefundQuery",
				IsCert:      false,
			},
		},
	},
	{
		Name:    `下载 相关`,
		Package: `download`,
		Apis: []Api{
			{
				Name: "下载交易账单",
				Description: `商户可以通过该接口下载历史交易清单。比如掉单、系统错误等导致商户侧和微信侧数据不一致，通过对账单核对后可校正支付状态。

注意：

1、微信侧未成功下单的交易不会出现在对账单中。支付成功后撤销的交易会出现在对账单中，跟原支付单订单号一致；

2、微信在次日9点启动生成前一天的对账单，建议商户10点后再获取；

3、对账单中涉及金额的字段单位为“元”。

4、对账单接口只能下载三个月以内的账单。

5、对账单是以商户号纬度来生成的，如一个商户号与多个appid有绑定关系，则使用其中任何一个appid都可以请求下载对账单。对账单中的appid取自交易时候提交的appid，与请求下载对账单时使用的appid无关。

6、自2018年起入驻的商户默认是开通免充值券后的结算对账单。`,
				Request:  "POST https://api.mch.weixin.qq.com/pay/downloadbill",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_6&index=8",
				FuncName: "DownloadBill",
				IsCert:   false,
			},
			{
				Name: "下载资金账单",
				Description: `
商户可以通过该接口下载自2017年6月1日起 的历史资金流水账单。

说明：

1、资金账单中的数据反映的是商户微信账户资金变动情况；

2、当日账单在次日上午9点开始生成，建议商户在上午10点以后获取；

3、资金账单中涉及金额的字段单位为“元”。
`,
				Request:  "POST https://api.mch.weixin.qq.com/pay/downloadfundflow",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_18&index=9",
				FuncName: "DownloadFundFlow",
				IsCert:   true,
			},
			{
				Name: "拉取订单评价数据",
				Description: `商户可以通过该接口拉取用户在微信支付交易记录中针对你的支付记录进行的评价内容。商户可结合商户系统逻辑对该内容数据进行存储、分析、展示、客服回访以及其他使用。如商户业务对评价内容有依赖，可主动引导用户进入微信支付交易记录进行评价。
`,
				Request:  "POST https://api.mch.weixin.qq.com/billcommentsp/batchquerycomment",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_99&index=12",
				FuncName: "BatchQueryComment",
				IsCert:   true,
			},
		},
	},
	{
		Name:    `付款码支付`,
		Package: `micropay`,
		Apis: []Api{
			{
				Name: "付款码支付",
				Description: `
收银员使用扫码设备读取微信用户付款码以后，二维码或条码信息会传送至商户收银台，由商户收银台或者商户后台调用该接口发起支付。

提醒1：提交支付请求后微信会同步返回支付结果。当返回结果为“系统错误”时，商户系统等待5秒后调用【查询订单API】，查询支付实际交易结果；当返回结果为“USERPAYING”时，商户系统可设置间隔时间(建议10秒)重新查询支付结果，直到支付成功或超时(建议30秒)；

提醒2：在调用查询接口返回后，如果交易状况不明晰，请调用【撤销订单API】，此时如果交易失败则关闭订单，该单不能再支付成功；如果交易成功，则将扣款退回到用户账户。当撤销无返回或错误时，请再次调用。注意：请勿扣款后立即调用【撤销订单API】,建议至少15秒后再调用。撤销订单API需要双向证书。
`,
				Request:  "POST https://api.mch.weixin.qq.com/pay/micropay",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1#",
				FuncName: "MicroPay",
				IsCert:   false,
			},
			{
				Name: "撤销订单",
				Description: `支付交易返回失败或支付系统超时，调用该接口撤销交易。如果此订单用户支付失败，微信支付系统会将此订单关闭；如果用户支付成功，微信支付系统会将此订单资金退还给用户。

注意：7天以内的交易单可调用撤销，其他正常支付的单如需实现相同功能请调用申请退款API。提交支付交易后调用【查询订单API】，没有明确的支付结果再调用【撤销订单API】。
`,
				Request:  "POST https://api.mch.weixin.qq.com/secapi/pay/reverse",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3",
				FuncName: "Reverse",
				IsCert:   true,
			},
			{
				Name:        "付款码查询openid",
				Description: `通过付款码查询公众号Openid，调用查询后，该付款码只能由此商户号发起扣款，直至付款码更新。`,
				Request:     "POST https://api.mch.weixin.qq.com/tools/authcodetoopenid",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_13&index=9",
				FuncName:    "AuthCodeToOpenId",
				IsCert:      false,
			},
		},
	},
	{
		Name:    `Native 支付`,
		Package: `native`,
		Apis: []Api{
			{
				Name:        "转换短链接",
				Description: `该接口主要用于Native支付模式一中的二维码链接转成短链接(weixin://wxpay/s/XXXXXX)，减小二维码数据量，提升扫描速度和精确度。`,
				Request:     "POST https://api.mch.weixin.qq.com/tools/shorturl",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=9_9&index=10",
				FuncName:    "ShortUrl",
				IsCert:      false,
			},
		},
	},
	{
		Name:    `开发辅助`,
		Package: `dev_util`,
		Apis: []Api{
			{
				Name:        "沙箱获取 signKey",
				Description: `验收仿真测试系统的API验签密钥需从API获取`,
				Request:     "POST https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=23_1&index=2",
				FuncName:    "GetSignKey",
				IsCert:      false,
			},
			{
				Name:        "交易保障 上报",
				Description: `商户在调用微信支付提供的相关接口时，会得到微信支付返回的相关信息以及获得整个接口的响应时间。为提高整体的服务水平，协助商户一起提高服务质量，微信支付提供了相关接口调用耗时和返回信息的主动上报接口，微信支付可以根据商户侧上报的数据进一步优化网络部署，完善服务监控，和商户更好的协作为用户提供更好的业务体验。`,
				Request:     "POST https://api.mch.weixin.qq.com/payitil/report",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_8&index=10",
				FuncName:    "Report",
				IsCert:      false,
			},
		},
	},
	{
		Name:    `代金券`,
		Package: `coupon`,
		Apis: []Api{
			{
				Name: "发放代金券",
				Description: `用于商户主动调用接口给用户发放代金券的场景，已做防小号处理，给小号发放代金券将返回错误码。

注意：通过接口发放的代金券不会进入微信卡包`,
				Request:  "POST https://api.mch.weixin.qq.com/mmpaymkttransfers/send_coupon",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/tools/sp_coupon.php?chapter=12_3&index=4",
				FuncName: "SendCoupon",
				IsCert:   false,
			},
			{
				Name:        "查询代金券批次",
				Description: ``,
				Request:     "POST https://api.mch.weixin.qq.com/mmpaymkttransfers/query_coupon_stock",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/tools/sp_coupon.php?chapter=12_4&index=5",
				FuncName:    "QueryCouponStock",
				IsCert:      false,
			},
			{
				Name:        "查询代金券信息",
				Description: ``,
				Request:     "POST https://api.mch.weixin.qq.com/mmpaymkttransfers/querycouponsinfo",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/tools/sp_coupon.php?chapter=12_5&index=6",
				FuncName:    "QueryCouponsInfo",
				IsCert:      false,
			},
		},
	},
	{
		Name:    `现金红包`,
		Package: `lucky_money`,
		Apis: []Api{
			{
				Name: "发放红包接口",
				Description: `1.发送频率限制------默认1800/min

2.发送个数上限------默认1800/min

3.场景金额限制------默认红包金额为1-200元，如有需要，可前往商户平台进行设置和申请

4.其他限制------商户单日出资金额上限--100万元；单用户单日收款金额上限--1000元；单用户可领取红包个数上限--10个；

`,
				Request:  "POST https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3",
				FuncName: "SendRedPack",
				IsCert:   true,
			},
			{
				Name:        "发放裂变红包",
				Description: `裂变红包：一次可以发放一组红包。首先领取的用户为种子用户，种子用户领取一组红包当中的一个，并可以通过社交分享将剩下的红包给其他用户。裂变红包充分利用了人际传播的优势。`,
				Request:     "POST https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_5&index=4",
				FuncName:    "SendGroupRedPack",
				IsCert:      true,
			},
			{
				Name:        "查询红包记录",
				Description: `用于商户对已发放的红包进行查询红包的具体信息，可支持普通红包和裂变包。`,
				Request:     "POST https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_6&index=5",
				FuncName:    "GetHBInfo",
				IsCert:      true,
			},
			{
				Name:        "小程序红包-发放",
				Description: ``,
				Request:     "POST hhttps://api.mch.weixin.qq.com/mmpaymkttransfers/sendminiprogramhb",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=18_2&index=3",
				FuncName:    "SendMiniprogramHB",
				IsCert:      true,
			},
		},
	},
	{
		Name:    `企业付款`,
		Package: `corp_pay`,
		Apis: []Api{
			{
				Name:        "企业付款 零钱",
				Description: ``,
				Request:     "POST https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2",
				FuncName:    "Transfers",
				IsCert:      true,
			},
			{
				Name: "查询 企业付款 零钱",
				Description: `用于商户的企业付款操作进行结果查询，返回付款操作详细结果。

查询企业付款API只支持查询30天内的订单，30天之前的订单请登录商户平台查询。`,
				Request:  "POST https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3",
				FuncName: "GetTransferInfo",
				IsCert:   true,
			},
			{
				Name:        "企业付款到银行卡",
				Description: `企业付款业务是基于微信支付商户平台的资金管理能力，为了协助商户方便地实现企业向银行卡付款，针对部分有开发能力的商户，提供通过API完成企业付款到银行卡的功能。`,
				Request:     "POST https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2",
				FuncName:    "PayBank",
				IsCert:      true,
			},
			{
				Name:        "查询 企业付款到银行卡",
				Description: `用于对商户企业付款到银行卡操作进行结果查询，返回付款操作详细结果。`,
				Request:     "POST https://api.mch.weixin.qq.com/mmpaysptrans/query_bank",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_3",
				FuncName:    "QueryBank",
				IsCert:      true,
			},
			{
				Name:        "获取 RSA 加密公钥",
				Description: ``,
				Request:     "POST https://fraud.mch.weixin.qq.com/risk/getpublickey",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7&index=4",
				FuncName:    "GetPublicKey",
				IsCert:      true,
			},
		},
	},
	{
		Name:    `分账`,
		Package: `profit_sharing`,
		Apis: []Api{
			{
				Name:        "请求单次分账",
				Description: `单次分账请求按照传入的分账接收方账号和资金进行分账，同时会将订单剩余的待分账金额解冻给本商户。故操作成功后，订单不能再进行分账，也不能进行分账完结。`,
				Request:     "POST https://api.mch.weixin.qq.com/secapi/pay/profitsharing",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1",
				FuncName:    "ProfitSharing",
				IsCert:      true,
			},
			{
				Name: "请求多次分账",
				Description: `● 微信订单支付成功后，商户发起分账请求，将结算后的钱分到分账接收方。多次分账请求仅会按照传入的分账接收方进行分账，不会对剩余的金额进行任何操作。故操作成功后，在待分账金额不等于零时，订单依旧能够再次进行分账。

● 多次分账，可以将本商户作为分账接收方直接传入，实现释放资金给本商户的功能

● 对同一笔订单最多能发起20次多次分账请求

`,
				Request:  "POST https://api.mch.weixin.qq.com/secapi/pay/multiprofitsharing",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_6&index=2",
				FuncName: "MultiProfitSharing",
				IsCert:   true,
			},
			{
				Name:        "查询分账结果",
				Description: `发起分账请求后，可调用此接口查询分账结果；发起分账完结请求后，可调用此接口查询分账完结的执行结果。`,
				Request:     "POST https://api.mch.weixin.qq.com/pay/profitsharingquery",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_2&index=3",
				FuncName:    "ProfitSharingQuery",
				IsCert:      false,
			},
			{
				Name:        "添加分账接收方",
				Description: `商户发起添加分账接收方请求，后续可通过发起分账请求将结算后的钱分到该分账接收方。`,
				Request:     "POST https://api.mch.weixin.qq.com/pay/profitsharingaddreceiver",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_3&index=4",
				FuncName:    "ProfitSharingAddReceiver",
				IsCert:      false,
			},
			{
				Name:        "删除分账接收方",
				Description: `商户发起删除分账接收方请求，删除后不支持将结算后的钱分到该分账接收方`,
				Request:     "POST https://api.mch.weixin.qq.com/pay/profitsharingremovereceiver",
				See:         "https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_4&index=5",
				FuncName:    "ProfitSharingRemoveReceiver",
				IsCert:      false,
			},
			{
				Name: "完结分账",
				Description: `1、不需要进行分账的订单，可直接调用本接口将订单的金额全部解冻给本商户
2、调用多次分账接口后，需要解冻剩余资金时，调用本接口将剩余的分账金额全部解冻给特约商户
3、已调用请求单次分账后，剩余待分账金额为零，不需要再调用此接口。`,
				Request:  "POST https://api.mch.weixin.qq.com/secapi/pay/profitsharingfinish",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_5&index=6",
				FuncName: "ProfitSharingFinish",
				IsCert:   true,
			},
			{
				Name: "分账回退",
				Description: `● 对订单进行退款时，如果订单已经分账，可以先调用此接口将指定的金额从分账接收方（仅限商户类型的分账接收方）回退给本商户，然后再退款。

● 回退以原分账请求为依据，可以对分给分账接收方的金额进行多次回退，只要满足累计回退不超过该请求中分给接收方的金额。

● 此接口采用同步处理模式，即在接收到商户请求后，会实时返回处理结果

● 此功能需要接收方在商户平台-交易中心-分账-分账接收设置下，开启同意分账回退后，才能使用。`,
				Request:  "POST https://api.mch.weixin.qq.com/secapi/pay/profitsharingreturn",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_7&index=7",
				FuncName: "ProfitSharingReturn",
				IsCert:   true,
			},
			{
				Name: "回退结果查询",
				Description: `● 商户需要核实回退结果，可调用此接口查询回退结果。

● 如果分账回退接口返回状态为处理中，可调用此接口查询回退结果`,
				Request:  "POST https://api.mch.weixin.qq.com/pay/profitsharingreturnquery",
				See:      "https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_8&index=8",
				FuncName: "ProfitSharingReturnQuery",
				IsCert:   false,
			},
		},
	},
}
