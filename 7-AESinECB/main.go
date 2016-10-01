package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	text, _ := ioutil.ReadFile("7.txt")
	tb, _ := base64.StdEncoding.DecodeString(string(text))
	//	err := ioutil.WriteFile("unbase64.txt", tb, 0644)
	//	if err != nil {
	//		panic(err)

	//	}

	cipher, _ := aes.NewCipher([]byte("YELLOW SUBMARINE"))

	bs := cipher.BlockSize()
	dest := []byte{}
	ciphertext := make([]byte, bs)
	for len(tb) > 0 {
		cipher.Decrypt(ciphertext, tb[:bs])
		tb = tb[bs:]
		dest = append(dest, ciphertext...)
	}
	fmt.Printf("Decrypted: \n%s \n", dest)
}
