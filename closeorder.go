package wxpay

import (
	"bytes"
	"encoding/xml"
	"net/url"

	"github.com/fastwego/wxpay/util"
)

const (
	apiCloseOrder = "/pay/closeorder"
)

/*
关闭订单参数

<xml>
   <appid>wx2421b1c4370ec43b</appid>
   <mch_id>10000100</mch_id>
   <nonce_str>4ca93f17ddf3443ceabf72f26d64fe0e</nonce_str>
   <out_trade_no>1415983244</out_trade_no>
   <sign>59FF1DF214B2D279A0EA7077C54DD95D</sign>
</xml>
*/
type CloseOrderParams struct {
	OutTradeNo string `xml:"out_trade_no"` // 商户系统内部的订单号
}

/*
<xml>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[BFK89FC6rxKCOjLX]]></nonce_str>
   <sign><![CDATA[72B321D92A7BFA0B2509F3D13C7B1631]]></sign>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <result_msg><![CDATA[OK]]></result_msg>
</xml>
*/
type CloseOrderResult struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchID      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ResultMsg  string `xml:"result_msg"`
}

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
func (wxpay *WXPay) CloseOrder(params CloseOrderParams) (result CloseOrderResult, err error) {

	allParams := struct {
		XMLName  xml.Name `xml:"xml"`
		Appid    string   `xml:"appid"`     // 应用ID
		MchID    string   `xml:"mch_id"`    // 商户号
		NonceStr string   `xml:"nonce_str"` // 随机字符串，不长于32位
		Sign     string   `xml:"sign"`      // 签名

		CloseOrderParams
	}{}

	allParams.CloseOrderParams = params
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

	resp, err := wxpay.Client.HTTPPost(apiCloseOrder, bytes.NewReader(body), "application/xml;charset=utf-8")
	if err != nil {
		return
	}

	err = xml.Unmarshal(resp, &result)
	if err != nil {
		return
	}

	return
}
