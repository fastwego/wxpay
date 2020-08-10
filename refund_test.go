package wxpay

import (
	"net/http"
	"reflect"
	"testing"
)

func TestWXPay_Refund(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`<xml>
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
</xml>`),
	}
	var resp []byte
	MockSvrHandler.HandleFunc(apiRefund, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		params RefundParams
	}
	tests := []struct {
		name       string
		args       args
		wantResult RefundResult
		wantErr    bool
	}{
		{name: "case1", args: args{params: RefundParams{
			OutTradeNo: "1415757673",
		}}, wantResult: RefundResult{
			ReturnCode: "SUCCESS",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResult, err := MockWXPay.Refund(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Refund() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult.ReturnCode, tt.wantResult.ReturnCode) {
				t.Errorf("Refund() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
