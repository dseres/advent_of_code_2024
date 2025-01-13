package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

//go:embed input03.txt
var input string

func main() {
	start := time.Now()
	fmt.Println("Day03 solution1:", solvePuzzle1(input))
	fmt.Println("Day03 solution2:", solvePuzzle2(input))
	fmt.Println("Day03 time:", time.Now().Sub(start))
}

func solvePuzzle1(input string) (sum int64) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	for _, m := range matches {
		sum += int64(computeMul(m))
	}
	return sum
}

func solvePuzzle2(input string) (sum int64) {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
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
			sum += int64(computeMul(m))
		}
	}
	return sum
}

func computeMul(matches []string) int {
	i, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	j, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	return i * j
}
