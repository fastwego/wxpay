package wxpay

import (
	"net/http"
	"reflect"
	"testing"
)

func TestWXPay_OrderQuery(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`<xml>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <device_info><![CDATA[1000]]></device_info>
   <nonce_str><![CDATA[TN55wO9Pba5yENl8]]></nonce_str>
   <sign><![CDATA[BDF0099C15FF7BC6B1585FBB110AB635]]></sign>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <openid><![CDATA[oUpF8uN95-Ptaags6E_roPHg7AG0]]></openid>
   <is_subscribe><![CDATA[Y]]></is_subscribe>
   <trade_type><![CDATA[APP]]></trade_type>
   <bank_type><![CDATA[CCB_DEBIT]]></bank_type>
   <total_fee>1</total_fee>
   <fee_type><![CDATA[CNY]]></fee_type>
   <transaction_id><![CDATA[1008450740201411110005820873]]></transaction_id>
   <out_trade_no><![CDATA[1415757673]]></out_trade_no>
   <attach><![CDATA[订单额外描述]]></attach>
   <time_end><![CDATA[20141111170043]]></time_end>
   <trade_state><![CDATA[SUCCESS]]></trade_state>
</xml>`),
	}
	var resp []byte
	MockSvrHandler.HandleFunc(apiOrderQuery, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		params OrderQueryParams
	}
	tests := []struct {
		name       string
		args       args
		wantResult OrderParams
		wantErr    bool
	}{
		{name: "case1", args: args{params: OrderQueryParams{
			TransactionID: "1008450740201411110005820873",
		}}, wantResult: OrderParams{
			TransactionID: "1008450740201411110005820873",
		}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResult, err := MockWXPay.OrderQuery(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult.TransactionID, tt.wantResult.TransactionID) {
				t.Errorf("OrderQuery() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
