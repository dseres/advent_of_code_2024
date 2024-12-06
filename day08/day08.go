package main

import "fmt"
import _ "embed"
import "strings"
import "strconv"

//go:embed input08.txt
var input string

func main() {
	nums := parseInput(input)
	fmt.Println("Day08 solution1:", solvePuzzle1(nums))
	fmt.Println("Day08 solution2:", solvePuzzle2(nums))
}

func solvePuzzle1(nums [][]int) int64 {
	return 0
}

func solvePuzzle2(nums [][]int) int64 {
	return 0
}

func parseInput(input string) (reports [][]int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			report := parseLine(line)
			reports = append(reports, report)
		}
	}
	return
}

func parseLine(line string) (report []int) {
	fields := strings.Fields(line)
	for _, field := range fields {
		val, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		report = append(report, val)
	}
	return
}
