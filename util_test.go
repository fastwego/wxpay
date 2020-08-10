package wxpay

import (
	"net/url"
	"testing"
)

func TestWXPay_Sign(t *testing.T) {
	type args struct {
		params   url.Values
		signType string
	}

	params := map[string][]string{
		"appid":       []string{"wxd930ea5d5a258f4f"},
		"mch_id":      []string{"10000100"},
		"device_info": []string{"1000"},
		"body":        []string{"test"},
		"nonce_str":   []string{"ibuaiVcKdpRxkhJA"},
	}
	tests := []struct {
		name     string
		args     args
		wantSign string
		wantErr  bool
	}{
		{name: "case1", args: args{params: params, signType: SignTypeMD5}, wantSign: "9A0A8659F005D6984697E2CA0A9CF3B7", wantErr: false},
		{name: "case2", args: args{params: params, signType: SignTypeHMACSHA256}, wantSign: "6A9AE1657590FD6257D693A078E1C3E4BB6BA4DC30B23E0EE2496E54170DACD6", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wxpay := &WXPay{
				Config: Config{
					Key: "192006250b4c09247ec02edce69f6a2d",
				},
			}
			gotSign, err := wxpay.Sign(tt.args.params, tt.args.signType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSign != tt.wantSign {
				t.Errorf("Sign() gotSign = %v, want %v", gotSign, tt.wantSign)
			}
		})
	}
}
