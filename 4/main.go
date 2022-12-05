package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	list := strings.Fields(string(data))
	var sections [][]int
	var split1, split2 []string
	var line []int
	var n int
	for _, elem := range list {
		split1 = strings.Split(elem, "-")
		for _, item := range split1 {
			split2 = strings.Split(item, ",")
			for _, item2 := range split2 {
				n, _ = strconv.Atoi(item2)
				line = append(line, n)
			}
		}
		sections = append(sections, line)
		line = make([]int, 0)
	}
	var count = 0
	for i := 0; i < len(sections); i++ {
		if (sections[i][0] <= sections[i][3] && sections[i][0] >= sections[i][2]) || (sections[i][1] >= sections[i][2] && sections[i][1] <= sections[i][3]) || (sections[i][2] <= sections[i][1] && sections[i][2] >= sections[i][0]) || (sections[i][3] >= sections[i][0] && sections[i][3] <= sections[i][1]) {
			count++
		}
	}
	fmt.Println(count)
}
