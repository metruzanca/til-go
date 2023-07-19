package main

import (
	"fmt"
	"os"
	"strconv"
)

// func (s *string) leftPad(count int, char rune) {}
// It appread that that ^, is not allowed. Rust can do that though :(

// Writing this in Go because I want to learn Go
func leftPad(str string, count int, char rune) string {
	for len(str) < count {
		str = string(char) + str
	}
	return str
}

func readInput(dayNum int) string {
	dat, err := os.ReadFile("inputs/" + leftPad(strconv.Itoa(dayNum), 2, '0') + ".txt")
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func main() {
	// https://adventofcode.com/2022/day/1
	input := readInput(1)
	fmt.Println(input)
	// Part 1

	// Part 2
}
