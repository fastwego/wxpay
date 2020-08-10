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

package util

import (
	"bytes"
	"crypto/aes"
	"errors"
	"fmt"
)

/*
AES encryption with ECB and PKCS7 padding

AES算法有AES-128、AES-192、AES-256三种，分别对应的key是 16、24、32字节长度

对应的加密解密区块长度BlockSize也是16、24、32字节长度
*/
func AESECBPKCS7Encrypt(pt, key []byte) (encrypted []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := NewECBEncrypter(block)

	bufLen := len(pt)
	padLen := mode.BlockSize() - (bufLen % mode.BlockSize())
	padText := bytes.Repeat([]byte{byte(padLen)}, padLen)
	pt = append(pt, padText...)
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)

	encrypted = pt
	return
}

/*
   AES decryption with ECB and PKCS7 padding

   AES算法有AES-128、AES-192、AES-256三种，分别对应的key是 16、24、32字节长度

   对应的加密解密区块长度BlockSize也是16、24、32字节长度
*/
func AESECBPKCS7Decrypt(cipherData, key []byte) (decrypted []byte, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
			return
		}
	}()

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := NewECBDecrypter(block)
	pt := make([]byte, len(cipherData))
	mode.CryptBlocks(pt, cipherData)

	bufLen := len(pt)
	if bufLen == 0 {
		err = errors.New("invalid padding size")
		return
	}

	pad := pt[bufLen-1]
	padLen := int(pad)
	if padLen > bufLen || padLen > mode.BlockSize() {
		err = errors.New("invalid padding size")
		return
	}

	for _, v := range pt[bufLen-padLen : bufLen-1] {
		if v != pad {
			err = errors.New("invalid padding")
			return
		}
	}

	decrypted = pt[:bufLen-padLen]

	return
}
