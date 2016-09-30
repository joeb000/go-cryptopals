package main

import (
	"fmt"
)

func main() {
	line := "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal"
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	byteLine := []byte(line)
	keyStr := []byte("ICE")
	fmt.Printf("Result: %x\n", repKeyXOR(byteLine, keyStr))
	fmt.Printf("Expect: %v\n", expected)
}

func repKeyXOR(b, key []byte) []byte {
	crypt := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		crypt[i] = b[i] ^ key[i%3]
	}
	return crypt
}
