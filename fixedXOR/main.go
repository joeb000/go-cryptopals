package main

import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

const (
	feedHex  = "1c0111001f010100061a024b53535009181c"
	xorStr   = "686974207468652062756c6c277320657965"
	wordSize = int(unsafe.Sizeof(uintptr(0)))
)

func main() {

	byteArr, _ := hex.DecodeString(feedHex)
	byteArr2, _ := hex.DecodeString(xorStr)

	dest := safeXORBytes(byteArr, byteArr2)

	fmt.Printf("XOR result: %x\n", dest)
	//fmt.Printf("UNsafe: %v \n\n", wordSize)
	fmt.Printf("Bytes %x\n\n", []byte(" "))
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
