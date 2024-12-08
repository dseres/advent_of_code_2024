package main

import "fmt"
import _ "embed"
import "strings"
import "strconv"

//go:embed input23.txt
var input string

func main() {
	nums := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(nums))
	fmt.Println("Day07 solution2:", solvePuzzle2(nums))
}

func solvePuzzle1(nums [][]int) int {
	return 0
}

func solvePuzzle2(nums [][]int) int {
	return 0
}

func parseInput(input string) (parsed [][]int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			report := parseLine(line)
			parsed = append(parsed, report)
		}
	}
	return
}

func parseLine(line string) (nums []int) {
	fields := strings.Fields(line)
	for _, field := range fields {
		val, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		nums = append(nums, val)
	}
	return
}
