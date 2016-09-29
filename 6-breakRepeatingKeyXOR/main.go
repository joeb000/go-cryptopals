package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	if !Prereq() {
		log.Fatal("Get your shit together dawg!")
	}
	bytes := readFileToBytes("6.txt")
	strbytes := base64StringtoHex(string(bytes))

	guessedKeySize := guessKeySize(strbytes)
	fmt.Printf("Guessed Key Size: %v \n", guessedKeySize)

	blockMultiArray := splitToBlocks(strbytes, guessedKeySize)
	fmt.Println(len(blockMultiArray))

	//	arrayBound := len(blockMultiArray)

	//	arr := [size][arrayBound]byte{}

	//	for i := 0; i < len(blockMultiArray); i++ { //number of blocks
	//		for j := 0; j < len(blockMultiArray[0]); j++ { //size
	//			arr[j][i] = blockMultiArray[i][j]
	//		}
	//	}
	//	fmt.Println("revrsed")
	//	fmt.Println(arr)

	for i := 0; i < len(blockMultiArray[0]); i++ { //size
		slice := []byte{}

		for j := 0; j < len(blockMultiArray); j++ { //number of blocks
			//fmt.Printf("%x\n", blockMultiArray[j][i])
			slice = append(slice, blockMultiArray[j][i])
		}
		fmt.Println("--------------------")
		fmt.Printf("Slice %v: %x\n", i, slice)
		singleByteXOR(slice)
	}

}

func singleByteXOR(hexbytes []byte) {
	for i := 0; i < 256; i++ {
		str := singlexor(hexbytes, byte(i))
		//fmt.Printf("bytes: %x \n", str)
		if score_plaintext(str) > 560 {
			fmt.Printf("\n\ntext: %q\nscore: %v\n byte: %x\n", str, score_plaintext(str), byte(i))
		}
	}
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
		if arr[i] == []byte("a")[0] {
			it++
		}
		if arr[i] == []byte("t")[0] {
			it++
		}
		if arr[i] == []byte("s")[0] {
			it++
		}
		if arr[i] == []byte("o")[0] {
			it++
		}
	}
	return it
}

func score_plaintext(pt []byte) int {
	score := 0

	// iterate over our string and give points to nice chars
	for _, c := range pt {

		// if its a printable ascii char +1
		if c > 31 && c < 127 {
			score += 1

			// we love spaces and e's
			if c == 32 {
				score += 2
			} else if c == 69 || c == 101 {
				score += 1
			}
			if c == 47 {
				score -= 1
			}
		}

	}

	return score
}

func singlexor(hex []byte, key byte) []byte {
	res := make([]byte, len(hex))

	for i := 0; i < len(hex); i++ {
		res[i] = hex[i] ^ key
	}
	return res
}

func splitToBlocks(byteArray []byte, size int) [][]byte {
	retArr := [][]byte{}
	block1 := []byte{}
	for i := 0; i < (len(byteArray) / size); i++ {
		block1 = byteArray[i*size : (i*size)+size]
		retArr = append(retArr, block1)

	}
	fmt.Println("Values")
	fmt.Println((len(byteArray) / size))
	fmt.Printf("val: %x\n", retArr[6][4])

	return retArr

}

func readFileToBytes(filename string) []byte {

	file, _ := os.Open(filename)

	defer file.Close()

	//fmt.Printf("str: %v \n", string(fbytes))

	byteSlice := []byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		byteSlice = append(byteSlice, scanner.Bytes()...)
	}
	return byteSlice
}

// picks the lowest keysize based on normalized Hamming Distance
func guessKeySize(input []byte) int {
	lowestHam := 100.0 //large number guarunteed to be beaten
	retval := 0
	for i := 2; i <= 40; i++ {
		//fmt.Printf("slice1: %x   len=%v\n", strbytes[:i], len(strbytes[:i]))
		//fmt.Printf("slice2: %x   len=%v\n", strbytes[i:2*i], len(strbytes[i:2*i]))
		ham := calculateHammingDistance(input[:i], input[i:2*i])

		normalizedHam := float64(ham) / float64(i)
		if lowestHam > normalizedHam {
			lowestHam = normalizedHam
			retval = i
		}

		//		if normalizedHam < 2.8 {
		//			fmt.Printf("Index: %v", i)
		//			fmt.Printf("\nNormalizedHam: %v\n----------\n", normalizedHam)

		//		}
	}
	return retval
}

// convert base64 encoded string into hex byte array
func base64StringtoHex(b64Str string) []byte {
	b64, _ := base64.StdEncoding.DecodeString(b64Str)
	return b64
}
