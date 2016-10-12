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
	hexToBytes, _ := hex.DecodeString(hexStr)
	encoded := base64.StdEncoding.EncodeToString(hexToBytes)
	fmt.Printf("hex: %x \nb64 encoded: %v \n\n", hexToBytes, encoded)

	if encoded == expected {
		fmt.Println("win")
	}
}
