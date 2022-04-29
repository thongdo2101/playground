package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	// GeneratePublicKeyEncryption()
	Exec()
}

func Exec() {
	// id, err := machineid.ID()
	// CheckError(err)
	// fmt.Println(id)
	id := "godgroup2022"
	fmt.Println("---")
	pubFile, err := ioutil.ReadFile("./key.rsa.pub")
	CheckError(err)
	privateFile, err := ioutil.ReadFile("./key.rsa")
	CheckError(err)
	block, _ := pem.Decode(pubFile)
	pub, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	content := []byte(fmt.Sprintf("%s+%v", id, time.Now().Unix()))
	fmt.Println("----")
	cipherText := RSA_OAEP_Encrypt(content, *pub)
	tmp1 := base64.StdEncoding.EncodeToString(cipherText)
	tmp2, err := base64.StdEncoding.DecodeString(tmp1)
	CheckError(err)
	fmt.Println(tmp1)
	fmt.Println("----")
	priBlock, _ := pem.Decode(privateFile)
	key, _ := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	strByte := tmp2
	fmt.Println("---")
	decryptMess := RSA_OAEP_Decrypt(strByte, *key)
	fmt.Println(string(decryptMess))
}

func GeneratePublicKeyEncryption() {
	filename := "key"
	bitSize := 4096

	// Generate RSA key.
	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}

	// Extract public component.
	pub := key.Public()

	// Encode private key to PKCS#1 ASN.1 PEM.
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)

	// Encode public key to PKCS#1 ASN.1 PEM.
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
		},
	)

	// Write private key to file.
	if err := ioutil.WriteFile(filename+".rsa", keyPEM, 0700); err != nil {
		panic(err)
	}

	// Write public key to file.
	if err := ioutil.WriteFile(filename+".rsa.pub", pubPEM, 0755); err != nil {
		panic(err)
	}
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}

func RSA_OAEP_Encrypt(secretMessage []byte, key rsa.PublicKey) []byte {
	label := []byte("")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, secretMessage, label)
	CheckError(err)
	return ciphertext
}

func RSA_OAEP_Decrypt(cipherText []byte, privKey rsa.PrivateKey) []byte {
	label := []byte("")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, cipherText, label)
	CheckError(err)
	return plaintext
}
