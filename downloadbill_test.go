package wxpay

import (
	"net/http"
	"reflect"
	"testing"
)

func TestWXPay_DownloadBill(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte(`<xml>
   <return_code	>FAIL</return_code	>
   <return_msg>return_msg</return_msg>
   <error_code>error_code</error_code>
</xml>`),
	}
	var resp []byte
	MockSvrHandler.HandleFunc(apiDownloadBill, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		params DownloadBillParams
	}
	tests := []struct {
		name       string
		args       args
		wantResult []byte
		wantErr    bool
	}{
		{name: "case1", args: args{params: DownloadBillParams{
			BillDate: "20141110",
		}}, wantResult: []byte(``), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResult, err := MockWXPay.DownloadBill(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadBill() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				//t.Errorf("DownloadBill() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
