package main

import (
	"encoding/hex"
	"fmt"
)

func Prereq() bool {
	test := "this is a test"
	wokka := "wokka wokka!!!"
	expected := 37
	aa, _ := hex.DecodeString(hex.EncodeToString([]byte(test)))
	bb, _ := hex.DecodeString(hex.EncodeToString([]byte(wokka)))

	fmt.Printf("test:     %x\n", aa)
	fmt.Printf("woka:     %x\n", bb)
	fmt.Printf("Hamming Distance: %v\n", calculateHammingDistance(aa, bb))
	if calculateHammingDistance(aa, bb) == expected {
		return true
	} else {
		return false
	}
}

func calculateHammingDistance(a []byte, b []byte) int {
	prod := xor(a, b)

	return countSetBits(prod)
}

// Brian Kernighan's method of counting set bits
// See https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan
func countSetBits(byteArray []byte) int {
	setBits := 0
	for _, b := range byteArray {
		for b != 0 {
			b &= b - 1 //bitwise & operator against byte-1
			setBits++
		}
	}
	return setBits
}

func xor(hex []byte, key []byte) []byte {
	res := make([]byte, len(hex))
	for i := 0; i < len(hex); i++ {
		res[i] = hex[i] ^ key[i]
	}
	return res
}
