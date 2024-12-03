package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input03.txt
var input string

func main() {
	fmt.Println("Day03 solution1:", solvePuzzle1(input))
	fmt.Println("Day03 solution2:", solvePuzzle2(input))
}

func solvePuzzle1(input string) int64 {
	re := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	matches := re.FindAllStringSubmatch(input, -1)
	sum := int64(0)
	for _, m := range matches {
		i, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		j, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		sum += int64(i * j)
	}
	return sum
}

func solvePuzzle2(input string) int64 {
	re := regexp.MustCompile("do\\(\\)|don't\\(\\)|mul\\((\\d+),(\\d+)\\)")
	matches := re.FindAllStringSubmatch(input, -1)
	sum := int64(0)
	do := true
	for _, m := range matches {
		if m[0] == "do()" {
			do = true
			continue
		}
		if m[0] == "don't()" {
			do = false
			continue
		}
		if do {
			i, err := strconv.Atoi(m[1])
			if err != nil {
				panic(err)
			}
			j, err := strconv.Atoi(m[2])
			if err != nil {
				panic(err)
			}
			sum += int64(i * j)
		}
	}
	return sum
}
