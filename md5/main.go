package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func main() {
	ciphertext := encrypt([]byte("45f7b07e-c309-4771-8a16-a028c62adb0b+dovanthong97@gmail.com"), "Chillgroup@123")
	fmt.Printf("Encrypted: %x\n", ciphertext)
	text, err := hex.DecodeString("33764e1f90a1fad2ed73ca33b9830a7de8ada7a27378b3877407bbfc622bf96ea91fe55dfa241355dcea1909699bfbf9a144518f7d4f8400b226ef83b562eee72c68a582e0748272183224980fda39d4f3bbf18ce0d3c3")
	if err != nil {
		panic(err)
	}
	plaintext := decrypt(text, "Chillgroup@123")
	fmt.Printf("Decrypted: %s\n", plaintext)
}
