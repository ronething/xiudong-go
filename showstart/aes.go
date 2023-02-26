package showstart

import (
	"encoding/base64"

	"github.com/forgoer/openssl"
)

type AESCrypto struct {
	Key []byte
}

func (a *AESCrypto) Encrypt(src []byte) (string, error) {
	dst, err := openssl.AesECBEncrypt(src, a.Key, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(dst), nil
}

func (a *AESCrypto) Decrypt(src string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	data, err := openssl.AesECBDecrypt(b, a.Key, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
