package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

//var encodedText []byte

func init() {
	loadRSAKey()
	//encodedText = encode("test@123")
	//fmt.Println(`encoded text: `, string(encodedText))
	//fmt.Println(`decoded text: `, decode(string(encodedText)))
}

var rsaKey *rsa.PrivateKey

func loadRSAKey() {
	bytes, err := ioutil.ReadFile("./config/keys/private.pem")
	if err != nil {
		fmt.Println(`loadRSAKey error: `, err)
		return
	}
	//fmt.Println(string(bytes))
	block, _ := pem.Decode(bytes)
	if block == nil {
		fmt.Println("failed to parse certificate PEM")
		return
	}
	rsaKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(`ParsePKCS1PrivateKey error: `, err)
		return
	}
	// fmt.Println(rsaKey)
}

func EncodeRSA(str string) []byte {
	h := sha256.New()
	bytes, err := rsa.EncryptOAEP(h, rand.Reader, &rsaKey.PublicKey, []byte(str), nil)
	if err != nil {
		fmt.Println(`encode error: `, err)
	}
	return bytes
}

func DecodeRSA(str string) string {
	strBytes, _ := base64.StdEncoding.DecodeString(str)
	// bytes, err := rsa.DecryptPKCS1v15(rand.Reader, rsaKey, strBytes)
	bytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, rsaKey, strBytes, nil)
	if err != nil {
		panic(`decode error: ` + err.Error())
	}

	return string(bytes)
}
