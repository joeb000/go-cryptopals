package main

import (
	"encoding/hex"
	"fmt"
)

//Write a function to compute the edit distance/Hamming distance between two strings.
//The Hamming distance is just the number of differing bits. The distance between:
//this is a test
//and
//wokka wokka!!!
//is 37. Make sure your code agrees before you proceed

func main() {
	test := "this is a test"
	wokka := "wokka wokka!!!"
	calcHammingDistance(test, wokka)
}

func hamming(a []byte, b []byte) int {
	prod := xor(a, b)
	distance := 0

	for _, b := range prod {
		fmt.Printf("B: %x\n", b)
		for b != 0 {
			b &= b - 1
			fmt.Printf("b: %x\n", b)

			distance++
		}
	}

	return distance
}

func calcHammingDistance(a, b string) int {
	aa, _ := hex.DecodeString(hex.EncodeToString([]byte(a)))
	bb, _ := hex.DecodeString(hex.EncodeToString([]byte(b)))
	fmt.Printf("test: %x", aa)
	fmt.Printf("woka: %x", bb)
	fmt.Printf("Ham: %v\n", hamming(aa, bb))
	return 0
}

func xor(hex []byte, key []byte) []byte {
	res := make([]byte, len(hex))

	for i := 0; i < len(hex); i++ {
		res[i] = hex[i] ^ key[i]
	}
	return res
}
