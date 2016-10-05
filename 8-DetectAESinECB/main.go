package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("8.txt")
	defer file.Close()
	//byteSlice := []byte{}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		check(scan.Text())
	}
}

func check(line string) {
	strArr := []string{}
	runeArr := make([]rune, 32)
	j := 0
	for i, char := range line {
		runeArr[j] = char
		if (i+1)%32 == 0 {
			//fmt.Printf("i: %v  char: %s \n", i, string(runeArr))
			strArr = append(strArr, string(runeArr))

			j = -1
		}
		j++
	}

	if findMatching(strArr) {
		fmt.Printf("Matching: %s", strArr)
	}
}

func findMatching(blocks []string) bool {
	for i, v1 := range blocks {
		for j, v2 := range blocks {
			if v1 == v2 && i != j {
				fmt.Printf("Block: %s\n", v1)
				return true
			}
		}
	}
	return false
}

func splitToBlocks(byteArray []byte, size int) [][]byte {
	retArr := [][]byte{}
	for i := 0; i < (len(byteArray) / size); i++ {
		block := byteArray[i*size : (i*size)+size]
		retArr = append(retArr, block)
	}
	return retArr
}
