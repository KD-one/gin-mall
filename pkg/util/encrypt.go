package util

import (
	"crypto/aes"
	"encoding/hex"
)

const Key = "zheshiyigemiyao"

// EncryptAES aes 加密
func EncryptAES(text string) (string, error) {
	cipher, err := aes.NewCipher([]byte(Key))
	if err != nil {
		return "", err
	}
	out := make([]byte, len(text))
	cipher.Encrypt(out, []byte(text))

	// EncodeToString 返回 out 的十六进制编码。
	return hex.EncodeToString(out), nil
}

// DecryptAES aes 解密
func DecryptAES(encryptText string) (string, error) {
	decodeString, err := hex.DecodeString(encryptText)
	if err != nil {
		return "", err
	}

	cipher, err := aes.NewCipher([]byte(Key))
	if err != nil {
		return "", err
	}

	out := make([]byte, len(decodeString))
	cipher.Decrypt(out, decodeString)

	return string(out), nil

}
