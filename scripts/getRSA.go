/*
 * Genarate rsa keys.
todo:！！！ 不能用 ！！！
ssh 生成的可以
openssl genrsa -out rsa_2048_priv.pem 2048
openssl rsa -pubout -in rsa_2048_priv.pem -out rsa_2048_pub.pem
*/

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := key.PublicKey

	savePEMKey("./config/keys/private.pem", key)
	savePublicPEMKey("./config/keys/public.pem", &publicKey)
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	bytes := x509.MarshalPKCS1PrivateKey(key)
	checkError(err)

	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: bytes,
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func savePublicPEMKey(fileName string, pubkey *rsa.PublicKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	bytes := x509.MarshalPKCS1PublicKey(pubkey)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: bytes,
	}

	err = pem.Encode(outFile, pemkey)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
