package wxpay

import (
	"bytes"
	"encoding/xml"
	"net/url"

	"github.com/fastwego/wxpay/util"
)

const (
	apiUnifiedOrder = "/pay/unifiedorder"
)

/*
统一下单参数
*/
type UnifiedOrderParams struct {
	DeviceInfo     string `xml:"device_info,omitempty"` // 设备号
	Body           string `xml:"body"`                  // 商品描述
	Detail         string `xml:"detail,omitempty"`      // detail
	Attach         string `xml:"attach,omitempty"`      // 附加数据
	OutTradeNo     string `xml:"out_trade_no"`          // 商户订单号
	TotalFee       string `xml:"total_fee"`             // 标价金额
	SPBillCreateIP string `xml:"spbill_create_ip"`      // 终端IP
	TimeStart      string `xml:"time_start,omitempty"`  // 交易起始时间
	TimeExpire     string `xml:"time_expire,omitempty"` // 交易结束时间
	GoodsTag       string `xml:"goods_tag,omitempty"`   // 订单优惠标记
	TradeType      string `xml:"trade_type"`            // 交易类型
	LimitPay       string `xml:"limit_pay,omitempty"`   // 指定支付方式
	Receipt        string `xml:"receipt,omitempty"`     // 开发票入口开放标识
}

/*
统一下单返回结果
*/
type UnifiedOrderResult struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid,omitempty"`
	Mchid      string `xml:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty"`
	TradeType  string `xml:"trade_type,omitempty"`
	PrePayid   string `xml:"prepay_id,omitempty"`
	CodeURL    string `xml:"code_url,omitempty"`
	ErrCode    string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`
}

/*
统一下单

商户系统先调用该接口在微信支付服务后台生成预支付交易单，返回正确的预支付交易会话标识后再在APP里面调起支付。

See: https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1

POST https://api.mch.weixin.qq.com/pay/unifiedorder
*/
func (wxpay *WXPay) UnifiedOrder(params UnifiedOrderParams) (result UnifiedOrderResult, err error) {

	allParams := struct {
		XMLName   xml.Name `xml:"xml"`
		Appid     string   `xml:"appid"`               // 应用ID
		Mchid     string   `xml:"mch_id"`              // 商户号
		NotifyURL string   `xml:"notify_url"`          // 通知地址
		NonceStr  string   `xml:"nonce_str"`           // 随机字符串
		Sign      string   `xml:"sign"`                // 签名
		SignType  string   `xml:"sign_type,omitempty"` // 签名类型
		FeeType   string   `xml:"fee_type,omitempty"`  // 标价币种

		UnifiedOrderParams
	}{}

	allParams.UnifiedOrderParams = params

	allParams.Appid = wxpay.Config.Appid
	allParams.Mchid = wxpay.Config.Mchid
	allParams.NonceStr = util.GetRandString(16)
	allParams.SignType = wxpay.Config.SignType
	allParams.FeeType = FeeTypeCNY
	allParams.NotifyURL = wxpay.Config.NotifyURL

	kvs := url.Values{}
	structToMap := util.StructToMap(allParams)
	for k, v := range structToMap {
		value, ok := v.(string)
		if ok && len(value) > 0 {
			kvs.Add(k, value)
		}
	}

	sign, err := wxpay.Sign(kvs, wxpay.Config.SignType)
	if err != nil {
		return
	}
	allParams.Sign = sign

	body, err := xml.Marshal(allParams)
	if err != nil {
		return
	}

	resp, err := wxpay.Client.HTTPPost(apiUnifiedOrder, bytes.NewReader(body), "application/xml;charset=utf-8")
	if err != nil {
		return
	}

	err = xml.Unmarshal(resp, &result)
	if err != nil {
		return
	}

	return
}
