package wxpay

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"net/url"
	"sort"
	"strings"
)

/*

请求参数签名

See: https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=4_3
*/
func (wxpay *WXPay) Sign(params url.Values, signType string) (sign string, err error) {

	var h hash.Hash
	if signType == SignTypeHMACSHA256 {
		h = hmac.New(sha256.New, []byte(wxpay.Config.Key))
	} else {
		h = md5.New()
	}

	kvs := []string{}
	for k, v := range params {
		if v[0] == "" {
			continue
		}
		kvs = append(kvs, fmt.Sprintf("%s=%s", k, v[0]))
	}

	sort.Strings(kvs)

	kvs = append(kvs, fmt.Sprintf("key=%s", wxpay.Config.Key))

	str := strings.Join(kvs, "&")

	if _, err = h.Write([]byte(str)); err != nil {
		return
	}

	sign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))

	return
}
