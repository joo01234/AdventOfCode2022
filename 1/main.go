package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.Split(string(data), "\n")
	max1, max2, max3 := 0, 0, 0
	sum := 0
	for _, word := range input {
		if word == "" {
			if sum > max1 {
				max3 = max2
				max2 = max1
				max1 = sum
			} else if sum > max2 {
				max3 = max2
				max2 = sum
			} else if sum > max3 {
				max3 = sum
			}
			sum = 0
			continue
		}
		add, _ := strconv.Atoi(word)
		sum += add
	}
	fmt.Println(max1 + max2 + max3)
}
