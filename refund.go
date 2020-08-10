package wxpay

import (
	"bytes"
	"encoding/xml"
	"net/url"

	"github.com/fastwego/wxpay/util"
)

const (
	apiRefund = "/secapi/pay/refund"
)

/*
申请退款 参数

eg:
<xml>
   <appid>wx2421b1c4370ec43b</appid>
   <mch_id>10000100</mch_id>
   <nonce_str>6cefdb308e1e2e8aabd48cf79e546a02</nonce_str>
   <out_refund_no>1415701182</out_refund_no>
   <out_trade_no>1415757673</out_trade_no>
   <refund_fee>1</refund_fee>
   <total_fee>1</total_fee>
   <transaction_id>4006252001201705123297353072</transaction_id>
   <sign>FE56DD4AA85C0EECA82C35595A69E153</sign>
</xml>
*/
type RefundParams struct {
	OutTradeNo    string `xml:"out_trade_no,omitempty"`   // 商户系统内部的订单号
	TransactionId string `xml:"transaction_id,omitempty"` // 微信订单号
	OutRefundNo   string `xml:"out_refund_no"`            // 商户退款单号

	TotalFee  int `xml:"total_fee"`  // 订单金额
	RefundFee int `xml:"refund_fee"` // 退款金额

	RefundFeeType string `xml:"refund_fee_type,omitempty"` // 商户退款单号
	RefundDesc    string `xml:"refund_desc,omitempty"`     // 商户退款单号
	RefundAccount string `xml:"refund_account,omitempty"`  // 商户退款单号
	NotifyUrl     string `xml:"notify_url,omitempty"`      // 商户退款单号

}

/*
eg:
<xml>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[NfsMFbUFpdbEhPXP]]></nonce_str>
   <sign><![CDATA[B7274EB9F8925EB93100DD2085FA56C0]]></sign>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <transaction_id><![CDATA[4008450740201411110005820873]]></transaction_id>
   <out_trade_no><![CDATA[1415757673]]></out_trade_no>
   <out_refund_no><![CDATA[1415701182]]></out_refund_no>
   <refund_id><![CDATA[2008450740201411110000174436]]></refund_id>
   <refund_fee>1</refund_fee>
</xml>
*/
type RefundResult struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`

	ResultCode string `xml:"result_code,omitempty"`

	ErrCode    string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`

	AppID    string `xml:"appid,omitempty"`
	MchID    string `xml:"mch_id,omitempty"`
	NonceStr string `xml:"nonce_str,omitempty"`
	Sign     string `xml:"sign,omitempty"`

	TransactionID       string `xml:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty"`
	RefundID            string `xml:"refund_id,omitempty"`
	RefundFee           int    `xml:"refund_fee,omitempty"`
	SettlementRefundFee int    `xml:"settlement_refund_fee,omitempty"`
	TotalFee            int    `xml:"total_fee,omitempty"`
	SettlementTotalFee  int    `xml:"settlement_total_fee,omitempty"`
	FeeType             string `xml:"fee_type,omitempty"`
	CashFee             int    `xml:"cash_fee,omitempty"`
	CashFeeType         string `xml:"cash_fee_type,omitempty"`

	CashRefundFee     int `xml:"cash_refund_fee,omitempty"`
	CouponRefundFee   int `xml:"coupon_refund_fee,omitempty"`
	CouponRefundCount int `xml:"coupon_refund_count,omitempty"`

	CouponType0      string `xml:"coupon_type_0"`
	CouponRefundID0  string `xml:"coupon_refund_id_0"`
	CouponRefundFee0 int    `xml:"coupon_refund_fee_0"`
}

/*

退款

当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家，微信支付将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退到买家帐号上。


See: https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_4&index=6

POST https://api.mch.weixin.qq.com/secapi/pay/refund
*/
func (wxpay *WXPay) Refund(params RefundParams) (result RefundResult, err error) {

	allParams := struct {
		XMLName  xml.Name `xml:"xml"`
		Appid    string   `xml:"appid"`     // 应用ID
		MchID    string   `xml:"mch_id"`    // 商户号
		NonceStr string   `xml:"nonce_str"` // 随机字符串，不长于32位
		Sign     string   `xml:"sign"`      // 签名

		RefundParams
	}{}

	allParams.RefundParams = params
	allParams.Appid = wxpay.Config.Appid
	allParams.MchID = wxpay.Config.Mchid
	allParams.NonceStr = util.GetRandString(16)

	kvs := url.Values{}
	structToMap := util.StructToMap(params)
	for k, v := range structToMap {
		value, ok := v.(string)
		if ok && len(value) > 0 {
			kvs.Add(k, value)
		}
	}

	sign, err := wxpay.Sign(kvs, SignTypeMD5)
	if err != nil {
		return
	}
	allParams.Sign = sign

	body, err := xml.Marshal(allParams)
	if err != nil {
		return
	}

	resp, err := wxpay.Client.HTTPPost(apiRefund, bytes.NewReader(body), "application/xml;charset=utf-8")
	if err != nil {
		return
	}

	err = xml.Unmarshal(resp, &result)
	if err != nil {
		return
	}

	return
}
