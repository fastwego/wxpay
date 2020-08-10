package wxpay

import (
	"net/http"
	"reflect"
	"testing"
)

func TestWXPay_CloseOrder(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`<xml>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[BFK89FC6rxKCOjLX]]></nonce_str>
   <sign><![CDATA[72B321D92A7BFA0B2509F3D13C7B1631]]></sign>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <result_msg><![CDATA[OK]]></result_msg>
</xml>`),
	}
	var resp []byte
	MockSvrHandler.HandleFunc(apiCloseOrder, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		params CloseOrderParams
	}
	tests := []struct {
		name       string
		args       args
		wantResult CloseOrderResult
		wantErr    bool
	}{
		{name: "case1", args: args{params: CloseOrderParams{
			OutTradeNo: "1008450740201411110005820873",
		}}, wantResult: CloseOrderResult{
			ReturnCode: "SUCCESS",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResult, err := MockWXPay.CloseOrder(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CloseOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult.ReturnCode, tt.wantResult.ReturnCode) {
				t.Errorf("CloseOrder() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
