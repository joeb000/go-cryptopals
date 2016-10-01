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

	cipherBytes := make([]byte, guessedKeySize)

	for i := 0; i < len(blockMultiArray[0]); i++ { //size
		slice := []byte{}

		for j := 0; j < len(blockMultiArray); j++ { //number of blocks
			slice = append(slice, blockMultiArray[j][i])
		}
		singleByteXOR(slice)
		cipherBytes[i] = singleByteXOR(slice)
	}

	fmt.Printf("\nCipherBytes: %x  Cipher %q\n", cipherBytes, cipherBytes)

	resultString := repKeyXOR(strbytes, cipherBytes)
	fmt.Printf("Decrypted Text:\n\n %s", resultString)
}

func repKeyXOR(b, key []byte) []byte {
	crypt := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		crypt[i] = b[i] ^ key[i%len(key)]
	}
	return crypt
}

func singleByteXOR(hexbytes []byte) byte {
	retByte := byte(0)
	bestScore := 0
	for i := 0; i < 256; i++ {
		str := singlexor(hexbytes, byte(i))
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
	return retArr
}

func readFileToBytes(filename string) []byte {

	file, _ := os.Open(filename)

	defer file.Close()
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
		ham1 := calculateHammingDistance(input[:i], input[i:2*i])
		ham2 := calculateHammingDistance(input[i:2*i], input[2*i:3*i])
		ham3 := calculateHammingDistance(input[2*i:3*i], input[3*i:4*i])
		ham4 := calculateHammingDistance(input[3*i:4*i], input[4*i:5*i])
		ham5 := calculateHammingDistance(input[4*i:5*i], input[5*i:6*i])
		ham6 := calculateHammingDistance(input[5*i:6*i], input[6*i:7*i])

		avgHam := average(float64(ham1), float64(ham2), float64(ham3), float64(ham4), float64(ham5), float64(ham6))
		normalizedHam := float64(avgHam) / float64(i)

		if lowestHam > normalizedHam {
			lowestHam = normalizedHam
			retval = i
		}
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
