package wxpay

import (
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
)

var MockWXPay *WXPay
var MockSvr *httptest.Server
var MockSvrHandler *http.ServeMux
var onceSetup sync.Once

func TestMain(m *testing.M) {
	onceSetup.Do(func() {
		MockWXPay = New(Config{
			Appid:     "APPID",
			Mchid:     "MCHID",
			Key:       "KEY",
			NotifyURL: "http://127.0.0.1/api/wxpay/notify",
			SignType:  SignTypeMD5,
		})

		// Mock Server
		MockSvrHandler = http.NewServeMux()
		MockSvr = httptest.NewServer(MockSvrHandler)
		WXPayServerUrl = MockSvr.URL // 拦截发往微信服务器的请求
	})
	os.Exit(m.Run())
}
