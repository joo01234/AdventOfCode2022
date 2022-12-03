package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	items := strings.Fields(string(data))
	end := ""
	// item1 := ""
	// item2 := ""
	found := ""
	for i := 0; i < len(items); i++ {
		found = ""
		if (i+1)%3 == 0 {
			for a := 'a'; a <= 'z'; a++ {
				if strings.Contains(items[i], string(a)) && strings.Contains(items[i-1], string(a)) && strings.Contains(items[i-2], string(a)) {
					found += string(a)
				}
			}
			for a := 'A'; a <= 'Z'; a++ {
				if strings.Contains(items[i], string(a)) && strings.Contains(items[i-1], string(a)) && strings.Contains(items[i-2], string(a)) {
					found += string(a)
				}
			}
		}
		end += found
		// item1 = item[:len(item)/2]
		// item2 = item[len(item)/2:]
		// for i := 0; i < len(item1); i++ {
		// 	char := item1[i]
		// 	if strings.Contains(item2, string(char)) {
		// 		if !strings.Contains(found, string(char)) {
		// 			found += string(char)
		// 			end += string(char)
		// 		}
		// 	}
		// }

	}
	total := 0
	for _, char := range end {
		for a := 'a'; a <= 'z'; a++ {
			A := unicode.ToUpper(a)

			if char == a {
				total += int(char)
				total -= 96
			} else if char == A {
				total += int(char)
				total -= 38
			}
		}
	}
	fmt.Println(total)
}
