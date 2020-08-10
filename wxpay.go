/*
Package wxpay 微信支付
*/
package wxpay

import (
	"log"
	"os"
)

const (
	SandBox            = "sandboxnew"
	SignTypeHMACSHA256 = "HMAC-SHA256"
	SignTypeMD5        = "MD5"
	FeeTypeCNY         = "CNY"

	//JSAPI--JSAPI支付（或小程序支付）
	//NATIVE--Native支付
	//APP--app支付
	//MWEB--H5支付，不同trade_type决定了调起支付的方式，请根据支付产品正确上传
	//
	//MICROPAY--付款码支付，付款码支付有单独的支付接口，所以接口不需要上传，该字段在对账单中会出现
	TradeTypeJSAPI    = "JSAPI"
	TradeTypeNATIVE   = "NATIVE"
	TradeTypeAPP      = "APP"
	TradeTypeMWEB     = "MWEB"
	TradeTypeMICROPAY = "MICROPAY"
)

/*
微信支付 实例
*/
type WXPay struct {
	Config Config
	Client Client
	Server Server
	Logger *log.Logger
}

type Config struct {
	Appid     string `json:"app_id"`
	Mchid     string `json:"mch_id"`
	Key       string `json:"key"`
	NotifyURL string `json:"notify_url"`
	SignType  string `json:"sign_type"` // 	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

/*
创建 实例
*/
func New(config Config) (wxpay *WXPay) {
	instance := WXPay{
		Config: config,
	}

	instance.Client = Client{Ctx: &instance}
	instance.Server = Server{Ctx: &instance}

	instance.Logger = log.New(os.Stdout, "[WXPay] ", log.LstdFlags|log.Llongfile)

	return &instance
}
