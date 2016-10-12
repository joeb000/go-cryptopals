package main

import (
	"encoding/hex"
	"fmt"
)

const (
	feedHex  = "1c0111001f010100061a024b53535009181c"
	xorStr   = "686974207468652062756c6c277320657965"
	expected = "746865206b696420646f6e277420706c6179"
)

func main() {

	bytes, _ := hex.DecodeString(feedHex)
	bytes2, _ := hex.DecodeString(xorStr)

	dest := safeXORBytes(bytes, bytes2)

	fmt.Printf("XOR result: %x\n", dest)

	if fmt.Sprintf("%x", dest) == expected {
		fmt.Println("win")
	}
}

//stolen func from crypto/cipher/xor.go
func safeXORBytes(a, b []byte) []byte {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	dst := make([]byte, n)
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return dst
}
