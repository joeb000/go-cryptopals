package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("4.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//	fmt.Println(scanner.Text())
		hexInBytes, _ := hex.DecodeString(scanner.Text())
		for i := 0; i < 256; i++ {
			str := xor(hexInBytes, byte(i))
			//fmt.Printf("bytes: %x \n", str)
			if score(str) > 4 {
				fmt.Printf("\n\ntext: %q\n\n", str)

			}
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
