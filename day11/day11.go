package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var input string = "64599 31 674832 2659361 1 0 8867 321"

type elem struct {
	num, round int
}

var cache map[elem]int = make(map[elem]int, 0)

func main() {
	start := time.Now()
	nums := parseInput(input)
	fmt.Println("Day11 solution1:", solvePuzzle(nums, 25))
	fmt.Println("Day11 solution2:", solvePuzzle(nums, 75))
	fmt.Println("Day11 time:", time.Now().Sub(start))
}

func solvePuzzle(nums []int, round int) int {
	sum := 0
	for _, n := range nums {
		sum += count(n, round)
	}
	return sum
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

func count(num, round int) int {
	if round == 0 {
		return 1
	}
	if c, ok := cache[elem{num, round}]; ok {
		return c
	}
	count := blink(num, round)
	cache[elem{num, round}] = count
	return count
}

func blink(num, round int) int {
	if num == 0 {
		return count(1, round-1)
	}
	digits := len(strconv.Itoa(num))
	if digits&1 == 0 {
		dividor := int(math.Pow10(digits / 2))
		left := num % dividor
		right := num / dividor
		return count(left, round-1) + count(right, round-1)
	}
	return count(num*2024, round-1)
}
