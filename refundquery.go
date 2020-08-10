package wxpay

import (
	"bytes"
	"encoding/xml"
	"net/url"

	"github.com/fastwego/wxpay/util"
)

const (
	apiRefundQuery = "/pay/refundquery"
)

/*
查询退款 参数

eg:
<xml>
   <appid>wx2421b1c4370ec43b</appid>
   <mch_id>10000100</mch_id>
   <nonce_str>0b9f35f484df17a732e537c37708d1d0</nonce_str>
   <out_refund_no></out_refund_no>
   <out_trade_no>1415757673</out_trade_no>
   <refund_id></refund_id>
   <transaction_id></transaction_id>
   <sign>66FFB727015F450D167EF38CCC549521</sign>
</xml>
*/
type RefundQueryParams struct {
	OutTradeNo    string `xml:"out_trade_no,omitempty"`   // 商户系统内部的订单号
	TransactionId string `xml:"transaction_id,omitempty"` // 微信订单号
	OutRefundNo   string `xml:"out_refund_no,omitempty"`  // 商户退款单号
	RefundId      string `xml:"refund_id,omitempty"`      // 微信退款单号
	Fffset        int    `xml:"offset,omitempty"`         // 偏移量

}

/*
eg:
<xml>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[TeqClE3i0mvn3DrK]]></nonce_str>
   <out_refund_no_0><![CDATA[1415701182]]></out_refund_no_0>
   <out_trade_no><![CDATA[1415757673]]></out_trade_no>
   <refund_count>1</refund_count>
   <refund_fee_0>1</refund_fee_0>
   <refund_id_0><![CDATA[2008450740201411110000174436]]></refund_id_0>
   <refund_status_0><![CDATA[PROCESSING]]></refund_status_0>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <sign><![CDATA[1F2841558E233C33ABA71A961D27561C]]></sign>
   <transaction_id><![CDATA[1008450740201411110005820873]]></transaction_id>
</xml>
*/
type RefundQueryResult struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`

	ResultCode string `xml:"result_code,omitempty"`

	ErrCode    string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`

	AppID    string `xml:"appid,omitempty"`
	MchID    string `xml:"mch_id,omitempty"`
	NonceStr string `xml:"nonce_str,omitempty"`
	Sign     string `xml:"sign,omitempty"`

	TransactionID      string `xml:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty"`
	TotalRefundCount   int    `xml:"total_refund_count,omitempty"`
	TotalFee           int    `xml:"total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty"`
	CashFee            int    `xml:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty"`
	SettlementTotalFee int    `xml:"settlement_total_fee,omitempty"`
	RefundCount        int    `xml:"refund_count"`

	OutRefundNo0       string `xml:"out_refund_no_0"`
	RefundID0          string `xml:"refund_id_0"`
	RefundChannel0     string `xml:"refund_channel_0,omitempty"`
	RefundFee0         int    `xml:"refund_fee_0,omitempty"`
	CouponRefundFee0   int    `xml:"coupon_refund_fee_0,omitempty"`
	CouponRefundCount0 int    `xml:"coupon_refund_count_0,omitempty"`
	CouponRefundID0_0  string `xml:"coupon_refund_id_0_0"`
	CouponType0_0      string `xml:"coupon_type_0_0"`
	CouponRefundFee0_0 int    `xml:"coupon_refund_fee_0_0,omitempty"`
	RefundStatus0      string `xml:"refund_status_0,omitempty"`
	RefundAccount0     string `xml:"refund_account_0,omitempty"`
	RefundRecvAccout0  string `xml:"refund_recv_accout_0,omitempty"`
	RefundSuccessTime0 string `xml:"refund_success_time_0,omitempty"`
}

/*

退款

当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家，微信支付将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退到买家帐号上。


See: https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_4&index=6

POST https://api.mch.weixin.qq.com/secapi/pay/refund
*/
func (wxpay *WXPay) RefundQuery(params RefundQueryParams) (result RefundQueryResult, err error) {

	allParams := struct {
		XMLName  xml.Name `xml:"xml"`
		Appid    string   `xml:"appid"`     // 应用ID
		MchID    string   `xml:"mch_id"`    // 商户号
		NonceStr string   `xml:"nonce_str"` // 随机字符串，不长于32位
		Sign     string   `xml:"sign"`      // 签名

		RefundQueryParams
	}{}

	allParams.RefundQueryParams = params
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

	resp, err := wxpay.Client.HTTPPost(apiRefundQuery, bytes.NewReader(body), "application/xml;charset=utf-8")
	if err != nil {
		return
	}

	err = xml.Unmarshal(resp, &result)
	if err != nil {
		return
	}

	return
}
