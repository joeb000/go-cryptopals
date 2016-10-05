package main

import (
	"bytes"
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
		if altscore(str) > 4 {
			fmt.Printf("\n\ntext: %q\nscore: %v\n", str, altscore(str))

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
		if arr[i] == []byte("e")[0] {
			it++
		}

	}
	return it
}

//see https://picoctf.com/crypto_mats/#substitution
func altscore(arr []byte) int {
	it := 0
	if bytes.Contains(arr, []byte("er")) {
		it += 2
	}
	if bytes.Contains(arr, []byte("of")) {
		it += 2
	}
	if bytes.Contains(arr, []byte("in")) {
		it += 2
	}
	if bytes.Contains(arr, []byte("th")) {
		it += 2
	}
	if bytes.Contains(arr, []byte(" ")) {
		it++
	}
	if bytes.Contains(arr, []byte("e")) {
		it++
	}
	if bytes.Contains(arr, []byte("t")) {
		it++
	}
	return it
}
