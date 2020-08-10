package wxpay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/fastwego/wxpay/util"
)

const (
	apiDownloadFundFlow = "/pay/downloadfundflow"
)

/*
账单的资金来源账户：

Basic  基本账户

Operation 运营账户

Fees 手续费账户
*/
const (
	AccountTypeBasic     = "Basic"
	AccountTypeOperation = "Operation"
	AccountTypeFees      = "Fees"
)

/*
下载资金账单

eg:
<xml>
  <appid>wx2421b1c4370ec43b</appid>
  <bill_date>20141110</bill_date>
  <account_type>Basic</account_type>
  <mch_id>10000100</mch_id>
  <nonce_str>21df7dc9cd8616b56919f20d9f679233</nonce_str>
  <sign>332F17B766FC787203EBE9D6E40457A1</sign>
</xml>
*/
type DownloadFundFlowParams struct {
	BillDate    string `xml:"bill_date"`          // 对账单日期
	AccountType string `xml:"account_type"`       // 账单类型
	TarType     string `xml:"tar_type,omitempty"` // 压缩账单
}

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
func (wxpay *WXPay) DownloadFundFlow(params DownloadFundFlowParams) (result []byte, err error) {

	allParams := struct {
		XMLName  xml.Name `xml:"xml"`
		Appid    string   `xml:"appid"`     // 应用ID
		MchID    string   `xml:"mch_id"`    // 商户号
		NonceStr string   `xml:"nonce_str"` // 随机字符串，不长于32位
		Sign     string   `xml:"sign"`      // 签名

		DownloadFundFlowParams
	}{}

	allParams.DownloadFundFlowParams = params
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

	sign, err := wxpay.Sign(kvs, SignTypeHMACSHA256)
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
