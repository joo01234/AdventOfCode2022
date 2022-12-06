package main

import (
	"fmt"
	"os"
)

func main() {
	dat, _ := os.ReadFile("data.txt")
	var list []byte
	for i := 0; i < 13; i++ {
		list = append(list, dat[i])
	}
	for i, char := range dat {
		if i >= 13 {
			if !repeated(char, list) {
				fmt.Println(i + 1)
				break
			} else {
				list = append(list, char)
				list = list[1:]
			}
		}
	}
}

func repeated(char byte, list []byte) bool {
	for i, c1 := range list {
		for _, c2 := range list[i+1:] {
			if c1 == c2 || c1 == char || c2 == char {
				return true
			}
		}
	}
	return false
}
