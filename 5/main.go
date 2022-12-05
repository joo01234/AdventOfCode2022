package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	var stacks = [][]byte{
		[]byte("BPNQHDRT"),
		[]byte("WGBJTV"),
		[]byte("NRHDSVMQ"),
		[]byte("PZNMC"),
		[]byte("DZB"),
		[]byte("VCWZ"),
		[]byte("GZNCVQLS"),
		[]byte("LGJMDNV"),
		[]byte("TPMFZCG"),
	}
	var line []string
	var amount, start, end int
	for scanner.Scan() {
		line = strings.Fields(scanner.Text())
		amount, _ = strconv.Atoi(line[1])
		start, _ = strconv.Atoi(line[3])
		start--
		end, _ = strconv.Atoi(line[5])
		end--
		for i := 0; i < amount; i++ {
			stacks[end] = append(stacks[end], stacks[start][len(stacks[start])-amount+i])
		}
		for i := 0; i < amount; i++ {
			stacks[start] = stacks[start][:len(stacks[start])-1]
		}
	}
	for i := 0; i < 9; i++ {
		fmt.Print(string(stacks[i][len(stacks[i])-1]))
	}
}
