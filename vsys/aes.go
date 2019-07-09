package vsys

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"errors"
)

// AesEncrypt is aes encrypt function
// key can be any length
// data can be any length
// iv is random,multi encrypt will get different result
// PKCS5Padding , CBC pattern
func AesEncrypt(key []byte, data []byte) ([]byte, error) {
	keyHash := sha512.Sum512(key)
	aesKey := keyHash[:32]
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	cbcIv := make([]byte, 16)
	rand.Read(cbcIv)
	data = pKCS5Padding(data, block.BlockSize())
	crypted := make([]byte, len(data)+16)
	copy(crypted[:16], cbcIv)
	blockMode := cipher.NewCBCEncrypter(block, cbcIv)
	blockMode.CryptBlocks(crypted[16:], data)
	return []byte(base64.StdEncoding.EncodeToString(crypted)), nil
}

// AesDecrypt aes cbc PKCS5Padding
func AesDecrypt(key []byte, data []byte) (origData []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("AES descrypt fail")
		}
	}()
	keyHash := sha512.Sum512(key)
	data, _ = base64.StdEncoding.DecodeString(string(data))
	cbcIv := data[:16]
	aesKey := keyHash[:32]
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, cbcIv)
	origData = make([]byte, len(data[16:]))
	blockMode.CryptBlocks(origData, data[16:])
	origData = pKCS5UnPadding(origData)
	return origData, nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
