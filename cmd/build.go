// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/iancoleman/strcase"
)

var rootPath = "./apis/"

func main() {
	var pkgFlag string
	flag.StringVar(&pkgFlag, "package", "default", "")
	flag.Parse()
	for _, group := range apiConfig {

		if group.Package == pkgFlag {
			build(group)
		}
	}

	if pkgFlag == "apilist" {
		apilist()
	}

}

func apilist() {
	for _, group := range apiConfig {
		fmt.Printf("- %s(%s)\n", group.Name, group.Package)
		for _, api := range group.Apis {
			split := strings.Split(api.Request, " ")
			parse, _ := url.Parse(split[1])

			if api.FuncName == "" {
				api.FuncName = strcase.ToCamel(path.Base(parse.Path))
			}

			godocLink := fmt.Sprintf("https://pkg.go.dev/github.com/fastwego/wxpay/apis/%s?tab=doc#%s", group.Package, api.FuncName)
			fmt.Printf("\t- [%s](%s) \n\t\t- [%s (%s)](%s)\n", api.Name, api.See, api.FuncName, parse.Path, godocLink)
		}
	}
}

func build(group ApiGroup) {

	for _, api := range group.Apis {

		_FUNC_NAME_ := api.FuncName

		split := strings.Split(api.Request, " ")
		parseUrl, _ := url.Parse(split[1])
		parseUrl.Path = strings.ReplaceAll(parseUrl.Path, "/sandboxnew", "")

		tpl := postFuncTpl
		if group.Package == "download" {
			tpl = postFuncTpl2
		}
		tpl = strings.ReplaceAll(tpl, "_TITLE_", api.Name)
		tpl = strings.ReplaceAll(tpl, "_DESCRIPTION_", api.Description)
		tpl = strings.ReplaceAll(tpl, "_REQUEST_", api.Request)
		tpl = strings.ReplaceAll(tpl, "_SEE_", api.See)
		tpl = strings.ReplaceAll(tpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_API_PATH_", parseUrl.Path)

		cert := "false"
		if api.IsCert {
			cert = "true"
		}
		tpl = strings.ReplaceAll(tpl, "_IS_CERT_", cert)

		// output func
		fileContent := fmt.Sprintf(`// Package %s %s
package %s

%s`, path.Base(group.Package), group.Name, path.Base(group.Package), tpl)
		filename := rootPath + group.Package + "/" + strcase.ToSnake(api.FuncName) + ".go"
		_ = os.MkdirAll(path.Dir(filename), 0644)
		ioutil.WriteFile(filename, []byte(fileContent), 0644)

		// TestFunc
		tpl = testFuncTpl
		if group.Package == "download" {
			tpl = testFuncTpl2
		}
		tpl = strings.ReplaceAll(tpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_API_PATH_", parseUrl.Path)

		// output Test
		testFileContent := fmt.Sprintf(`package %s

%s
`, path.Base(group.Package), tpl)
		//fmt.Println(testFileContent)
		ioutil.WriteFile(rootPath+group.Package+"/"+strcase.ToSnake(api.FuncName)+"_test.go", []byte(testFileContent), 0644)

		//Example
		tpl = strings.ReplaceAll(exampleFuncTpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_PACKAGE_", path.Base(group.Package))

		// output example
		exampleFileContent := fmt.Sprintf(`package %s_test

%s
`, path.Base(group.Package), tpl)
		//fmt.Println(testFileContent)
		ioutil.WriteFile(rootPath+group.Package+"/example_"+strcase.ToSnake(api.FuncName)+"_test.go", []byte(exampleFileContent), 0644)

	}

}

var postFuncTpl = `/*
_TITLE_

_DESCRIPTION_

See: _SEE_

_REQUEST_
*/
func _FUNC_NAME_(ctx *wxpay.WXPay, params map[string]string) (result map[string]string, err error) {

	resp, err := ctx.Client.HTTPPost("_API_PATH_", params, _IS_CERT_)
	if err != nil {
		return
	}

	result, err = util.XML2Map(resp)
	if err != nil {
		return
	}
	return
}
`

var postFuncTpl2 = `/*
_TITLE_

_DESCRIPTION_

See: _SEE_

_REQUEST_
*/
func _FUNC_NAME_(ctx *wxpay.WXPay, params map[string]string) (result []byte, err error) {
	return ctx.Client.HTTPPost("_API_PATH_", params, _IS_CERT_)
}
`

var testFuncTpl = `
func Test_FUNC_NAME_(t *testing.T) {
	test.Setup()

	mockResp := map[string][]byte{
		"case1": []byte("<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg></xml>"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc("_API_PATH_", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		params map[string]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		{name: "case1", args: args{params: map[string]string{
			"appid":     test.MockWXPay.Config.Appid,
			"mch_id":    test.MockWXPay.Config.Mchid,
			"nonce_str": util.GetRandString(16),
			"sign_type": types.SignTypeHMACSHA256}}, wantResult: "SUCCESS", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResult, err := _FUNC_NAME_(test.MockWXPay, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("_FUNC_NAME_() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult["return_code"], tt.wantResult) {
				t.Errorf("CloseOrder() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
`

var testFuncTpl2 = `
func Test_FUNC_NAME_(t *testing.T) {
	test.Setup()

	mockResp := map[string][]byte{
		"case1": []byte("<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg></xml>"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc("_API_PATH_", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		params map[string]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []byte
		wantErr    bool
	}{
		{name: "case1", args: args{params: map[string]string{
			"appid":     test.MockWXPay.Config.Appid,
			"mch_id":    test.MockWXPay.Config.Mchid,
			"nonce_str": util.GetRandString(16),
			"sign_type": types.SignTypeHMACSHA256}}, wantResult: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResult, err := _FUNC_NAME_(test.MockWXPay, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchQueryComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("CloseOrder() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
`

var exampleFuncTpl = `
func Example_FUNC_NAME_() {
	var ctx *wxpay.WXPay

	params := map[string]string{
		"appid":"APPID",
		// ...
	}
	resp, err := _PACKAGE_._FUNC_NAME_(ctx, params)

	fmt.Println(resp, err)
}
`
