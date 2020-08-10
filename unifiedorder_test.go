package wxpay

import (
	"net/http"
	"reflect"
	"testing"
)

func TestWXPay_UnifiedOrder(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`<xml>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str>
   <sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id>
   <trade_type><![CDATA[APP]]></trade_type>
</xml>`),
	}
	var resp []byte
	MockSvrHandler.HandleFunc(apiUnifiedOrder, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		params UnifiedOrderParams
	}
	tests := []struct {
		name       string
		args       args
		wantResult UnifiedOrderResult
		wantErr    bool
	}{
		{name: "case1", args: args{params: UnifiedOrderParams{
			DeviceInfo:     "DeviceInfo",
			Body:           "Body",
			OutTradeNo:     "NO.10086",
			TotalFee:       "100",
			SPBillCreateIP: "201.24.53.118",
			TradeType:      TradeTypeAPP,
		}}, wantResult: UnifiedOrderResult{
			PrePayid: "wx201411101639507cbf6ffd8b0779950874",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResult, err := MockWXPay.UnifiedOrder(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnifiedOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult.PrePayid, tt.wantResult.PrePayid) {
				t.Errorf("UnifiedOrder() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
