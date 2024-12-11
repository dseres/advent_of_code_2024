package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var input string = "64599 31 674832 2659361 1 0 8867 321"

func main() {
	nums := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(nums, 25))
	fmt.Println("Day07 solution2:", solvePuzzle2(nums, 25))
}

func solvePuzzle1(nums []int, round int) int {
	fmt.Println(nums)
	for i := 0; i < round; i++ {
		nums = blink(nums)
		fmt.Println(nums)
	}
	return len(nums)
}

func solvePuzzle2(nums []int, round int) int {
	return 0
}

func parseInput(input string) (nums []int) {
	fields := strings.Fields(input)
	for _, field := range fields {
		val, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		nums = append(nums, val)
	}
	return
}

func blink(nums []int) []int {
	next := make([]int, 0, len(nums)*2)
	for _, n := range nums {
		digits := len(strconv.Itoa(n))
		if n == 0 {
			next = append(next, 1)
		} else if digits&1 == 0 {
			dividor := int(math.Pow10(digits / 2))
			left := n % dividor
			right := n / dividor
			next = append(next, left, right)
		} else {
			next = append(next, n*2024)
		}
	}
	return next
}
