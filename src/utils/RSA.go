package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var publicKey string
var privateKey string

// bits 生成的公私钥对的位数，一般为 1024 或 2048
// privateKey 生成的私钥
// publicKey 生成的公钥
func GenRsaKey(bits int) (privateKey, publicKey string) {
	priKey, err2 := rsa.GenerateKey(rand.Reader, bits)
	if err2 != nil {
		panic(err2)
	}

	derStream := x509.MarshalPKCS1PrivateKey(priKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	prvKey := pem.EncodeToMemory(block)
	puKey := &priKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(puKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubKey := pem.EncodeToMemory(block)

	privateKey = string(prvKey)
	publicKey = string(pubKey)
	return
}

// RsaEncryptBase64 使用 RSA 公钥加密数据, 返回加密后并编码为 base64 的数据
func RsaEncryptBase64(originalData string) (string, error) {
	block, _ := pem.Decode([]byte(publicKey))
	pubKey, parseErr := x509.ParsePKIXPublicKey(block.Bytes)
	if parseErr != nil {
		fmt.Println(parseErr)
		return "", errors.New("解析公钥失败")
	}
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(originalData))
	return base64.StdEncoding.EncodeToString(encryptedData), err
}

func ReadKey() (err error) {
	publicKeyByte, err := os.ReadFile(viper.GetString("File.PublicKeyFileName"))
	if err != nil {
		return err
	}

	privateKeyByte, err := os.ReadFile(viper.GetString("File.PrivateKeyFileName"))
	if err != nil {
		return err
	}

	publicKey = string(publicKeyByte)
	privateKey = string(privateKeyByte)
	return nil
}

// RsaDecryptBase64 使用 RSA 私钥解密数据
func RsaDecryptBase64(encryptedData string) (string, error) {
	encryptedDecodeBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode([]byte(privateKey))
	priKey, parseErr := x509.ParsePKCS1PrivateKey(block.Bytes)
	if parseErr != nil {
		fmt.Println(parseErr)
		return "", errors.New("解析私钥失败")
	}

	originalData, encryptErr := rsa.DecryptPKCS1v15(rand.Reader, priKey, encryptedDecodeBytes)
	return string(originalData), encryptErr
}
