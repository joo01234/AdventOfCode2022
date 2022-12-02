package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	text1 := string(data)
	text1 = strings.Replace(text1, "X", "0", -1)
	text1 = strings.Replace(text1, "Y", "3", -1)
	text1 = strings.Replace(text1, "Z", "6", -1)
	text1 = strings.Replace(text1, "A", "1", -1)
	text1 = strings.Replace(text1, "B", "2", -1)
	text1 = strings.Replace(text1, "C", "3", -1)
	text := strings.Fields(text1)
	var total, add, value int
	for i, elem := range text {
		add = 0
		if i%2 == 1 {
			add, _ = strconv.Atoi(elem)
			value, _ = strconv.Atoi(text[i-1])
			if add == 0 {
				value -= 1
				if value == 0 {
					value = 3
				}
			} else if add == 6 {
				value += 1
				if value == 4 {
					value = 1
				}
			}
			add += value

		}
		total += add
	}
	fmt.Println(total)
}
