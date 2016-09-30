package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if !Prereq() {
		log.Fatal("Get your shit together dawg!")
	}
	//bytes := readFileToBytes("6.txt")
	bytes, _ := ioutil.ReadFile("6.txt")

	strbytes := base64StringtoHex(string(bytes))

	fmt.Printf("data: %x\n", strbytes)

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

	cipherBytes := make([]byte, guessedKeySize)

	for i := 0; i < len(blockMultiArray[0]); i++ { //size
		slice := []byte{}

		for j := 0; j < len(blockMultiArray); j++ { //number of blocks
			//fmt.Printf("%x\n", blockMultiArray[j][i])
			slice = append(slice, blockMultiArray[j][i])
		}
		fmt.Println("--------------------")
		fmt.Printf("Slice %v: %x\n", i, slice)
		singleByteXOR(slice)
		cipherBytes[i] = singleByteXOR(slice)
	}

	fmt.Println(string(cipherBytes))

	fmt.Printf("\n\nCipherBytes: %x  Cipher %q\n", cipherBytes, cipherBytes)

}

func singleByteXOR(hexbytes []byte) byte {
	retByte := byte(0)
	bestScore := 0
	for i := 0; i < 256; i++ {
		str := singlexor(hexbytes, byte(i))
		//fmt.Printf("bytes: %x \n", str)

		if score(str) > bestScore {
			bestScore = score(str)
			retByte = byte(i)
		}
	}

	return retByte
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
	for i := 6; i <= 40; i++ {
		//fmt.Printf("slice1: %x   len=%v\n", strbytes[:i], len(strbytes[:i]))
		//fmt.Printf("slice2: %x   len=%v\n", strbytes[i:2*i], len(strbytes[i:2*i]))
		ham1 := calculateHammingDistance(input[:i], input[i:2*i])
		fmt.Printf("HAM1: %v\n", ham1)

		ham2 := calculateHammingDistance(input[i:2*i], input[2*i:3*i])
		ham3 := calculateHammingDistance(input[2*i:3*i], input[3*i:4*i])
		ham4 := calculateHammingDistance(input[3*i:4*i], input[4*i:5*i])
		ham5 := calculateHammingDistance(input[4*i:5*i], input[5*i:6*i])
		ham6 := calculateHammingDistance(input[5*i:6*i], input[6*i:7*i])

		avgHam := average(float64(ham1), float64(ham2), float64(ham3), float64(ham4), float64(ham5), float64(ham6))
		normalizedHam := float64(avgHam) / float64(i)
		fmt.Printf("i: %v avgHame: %v normalized %v\n", i, avgHam, normalizedHam)

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
func average(xs ...float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// convert base64 encoded string into hex byte array
func base64StringtoHex(b64Str string) []byte {
	b64, _ := base64.StdEncoding.DecodeString(b64Str)
	return b64
}
