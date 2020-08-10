package wxpay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/fastwego/wxpay/util"
)

const (
	apiDownloadBill = "/pay/downloadbill"
)

/*
ALL（默认值），返回当日所有订单信息（不含充值退款订单）

SUCCESS，返回当日成功支付的订单（不含充值退款订单）

REFUND，返回当日退款订单（不含充值退款订单）

RECHARGE_REFUND，返回当日充值退款订单
*/
const (
	BillTypeALL             = "ALL"
	BillTypeSUCCESS         = "SUCCESS"
	BillTypeREFUND          = "REFUND"
	BillTypeRECHARGE_REFUND = "RECHARGE_REFUND"
)

/*
下载交易账单

eg:
<xml>
   <appid>wx2421b1c4370ec43b</appid>
   <bill_date>20141110</bill_date>
   <bill_type>ALL</bill_type>
   <mch_id>10000100</mch_id>
   <nonce_str>21df7dc9cd8616b56919f20d9f679233</nonce_str>
   <sign>332F17B766FC787203EBE9D6E40457A1</sign>
</xml>
*/
type DownloadBillParams struct {
	BillDate string `xml:"bill_date"`           // 对账单日期
	BillType string `xml:"bill_type,omitempty"` // 账单类型
	TarType  string `xml:"tar_type,omitempty"`  // 压缩账单
}

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
func (wxpay *WXPay) DownloadBill(params DownloadBillParams) (result []byte, err error) {

	allParams := struct {
		XMLName  xml.Name `xml:"xml"`
		Appid    string   `xml:"appid"`     // 应用ID
		MchID    string   `xml:"mch_id"`    // 商户号
		NonceStr string   `xml:"nonce_str"` // 随机字符串，不长于32位
		Sign     string   `xml:"sign"`      // 签名

		DownloadBillParams
	}{}

	allParams.DownloadBillParams = params
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

	resp, err := wxpay.Client.HTTPPost(apiDownloadBill, bytes.NewReader(body), "application/xml;charset=utf-8")
	if err != nil {
		return
	}

	errorResp := struct {
		ReturnCode string `xml:"return_code"`
		ReturnMsg  string `xml:"return_msg"`
		ErrorCode  string `xml:"error_code"`
	}{}
	err = xml.Unmarshal(resp, &errorResp)
	if err != nil {
		return
	}

	if errorResp.ErrorCode != "" {
		err = fmt.Errorf("%v", errorResp)
		return
	}
	result = resp
	return
}
