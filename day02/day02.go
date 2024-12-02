package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input02.txt
var input string

func main() {
	reports := parseInput(input)
	fmt.Println("Day02 solution1:", solvePuzzle1(reports))
	fmt.Println("Day02 solution2:", solvePuzzle2(reports))
}

func solvePuzzle1(reports [][]int) (count int64) {
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	return
}

func solvePuzzle2(reports [][]int) (count int64) {
	for _, report := range reports {
		if isDampenedSafe(report) {
			count++
		}
	}
	return
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

func isSafe(report []int) bool {
	return isDecreasing(report) || isIncreasing(report)
}

func isDecreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i] <= report[i+1] || report[i]-report[i+1] > 3 {
			return false
		}
	}
	return true
}

func isIncreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i] >= report[i+1] || report[i+1]-report[i] > 3 {
			return false
		}
	}
	return true
}

func checkWithDampener(report []int, checkFun func([]int) bool) bool {
	for i, _ := range report {
		dampenedReport := slices.Concat(report[:i], report[i+1:])
		if checkFun(dampenedReport) {
			return true
		}
	}
	return false
}

func isDampenedSafe(report []int) bool {
	return checkWithDampener(report, isIncreasing) || checkWithDampener(report, isDecreasing)
}
