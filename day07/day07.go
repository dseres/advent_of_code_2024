package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input07.txt
var input string

type operand int

const (
	ADD operand = iota
	MUL
	CONCAT
)

func main() {
	nums := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(nums))
	fmt.Println("Day07 solution2:", solvePuzzle2(nums))
}

func solvePuzzle1(nums [][]int) int {
	sum := 0
	for _, row := range nums {
		if valid(row, false) {
			sum += row[0]
		}
	}
	return sum
}

func solvePuzzle2(nums [][]int) int {
	sum := 0
	for _, row := range nums {
		if valid(row, true) {
			sum += row[0]
		}
	}
	return sum
}

func valid(nums []int, concat bool) bool {
	if testOp(nums[0], concat, nums[2:], nums[1], ADD) {
		return true
	}
	if testOp(nums[0], concat, nums[2:], nums[1], MUL) {
		return true
	}
	if concat && testOp(nums[0], concat, nums[2:], nums[1], CONCAT) {
		return true
	}
	return false
}

func testOp(required int, concat bool, nums []int, computed int, op operand) bool {
	if len(nums) == 0 {
		return computed == required
	}
	next := doOp(computed, nums[0], op)
	if testOp(required, concat, nums[1:], next, ADD) {
		return true
	}
	if testOp(required, concat, nums[1:], next, MUL) {
		return true
	}
	if concat && testOp(required, concat, nums[1:], next, CONCAT) {
		return true
	}
	return false
}

func doOp(i, j int, op operand) int {
	next := 0
	switch op {
	case ADD:
		next = i + j
	case MUL:
		next = i * j
	case CONCAT:
		nextStr := fmt.Sprintf("%v%v", i, j)
		next, _ = strconv.Atoi(nextStr)
	}
	return next
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
	fields := regexp.MustCompile(`:?\s+`).Split(line, -1)
	for _, field := range fields {
		val, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		report = append(report, val)
	}
	return
}
