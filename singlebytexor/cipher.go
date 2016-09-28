package main

import (
	"encoding/hex"
	"fmt"
)

const (
	hexStr = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
)

func main() {
	hexInBytes, _ := hex.DecodeString(hexStr)
	for i := 0; i < 256; i++ {
		str := xor(hexInBytes, byte(i))
		//fmt.Printf("bytes: %x \n", str)
		if score(str) > 4 {
			fmt.Printf("\n\ntext: %q\n\n", str)

		}
	}
}

func xor(hex []byte, key byte) []byte {
	res := make([]byte, len(hex))

	for i := 0; i < len(hex); i++ {
		res[i] = hex[i] ^ key
	}
	return res
}

func score(arr []byte) int {
	it := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == []byte(" ")[0] {
			it++
		}
	}
	return it
}
