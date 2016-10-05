package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const (
	hexStr   = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
)

func main() {
	theBytes, _ := hex.DecodeString(hexStr)
	tbytes := []byte(hexStr)
	encoded := base64.StdEncoding.EncodeToString(theBytes)
	fmt.Printf("bytes: %x, %v \n\n", theBytes, encoded)
	fmt.Printf("bytes: %x\n", tbytes)
	b64, _ := base64.StdEncoding.DecodeString(expected)

	fmt.Printf("B64: %x \n", b64)

	if encoded == expected {
		fmt.Println("win")
	}
}