package wxpay

import (
	"net/http"
	"reflect"
	"testing"
)

func TestWXPay_RefundQuery(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`<xml>
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
</xml>`),
	}
	var resp []byte
	MockSvrHandler.HandleFunc(apiRefundQuery, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		params RefundQueryParams
	}
	tests := []struct {
		name       string
		args       args
		wantResult RefundQueryResult
		wantErr    bool
	}{
		{name: "case1", args: args{params: RefundQueryParams{
			OutTradeNo: "1415701182",
		}}, wantResult: RefundQueryResult{
			ReturnCode: "SUCCESS",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResult, err := MockWXPay.RefundQuery(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefundQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult.ReturnCode, tt.wantResult.ReturnCode) {
				t.Errorf("RefundQuery() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
