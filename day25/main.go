package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input25.txt
var input string

func main() {
	start := time.Now()
	locks, keys := parseInput(input)
	fmt.Println("Day25 solution1:", solvePuzzle1(locks, keys))
	fmt.Println("Day25 time:", time.Now().Sub(start))
}

func parseInput(input string) (locks [][]int, keys [][]int) {
	blocks := strings.Split(strings.TrimSpace(input), "\n\n")
	for _, block := range blocks {
		parsed, isLock := parseBlock(block)
		if isLock {
			locks = append(locks, parsed)
			continue
		}
		keys = append(keys, parsed)
	}
	return
}

func parseBlock(input string) (parsed []int, isLock bool) {
	lines := strings.Split(input, "\n")
	isLock = checkBlock(lines)
	parsed = make([]int, len(lines[0]))
	for _, line := range lines[1 : len(lines)-1] {
		for j, r := range line {
			if r == '#' {
				parsed[j]++
			}
		}
	}
	return
}

func checkBlock(lines []string) bool {
	isLock := false
	if lines[0] == "#####" {
		isLock = true
	} else if lines[0] == "....." {
		isLock = false
	} else {
		panic(fmt.Sprintf("bad block, it is not lock or key: %v", lines))
	}
	return isLock
}

func solvePuzzle1(locks [][]int, keys [][]int) int {
	// fmt.Println(len(locks), locks, len(keys), keys)
	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			// fmt.Println(lock, key, match(lock, key))
			if match(lock, key) {
				count++
			}
		}
	}
	return count
}

func match(lock, key []int) bool {
	for i, l := range lock {
		k := key[i]
		if 5 < l+k {
			return false
		}
	}
	return true
}
