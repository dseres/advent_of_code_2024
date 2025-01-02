package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input22.txt
var input string

type key struct{ a, b, c, d int }

func main() {
	nums := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(nums))
	fmt.Println("Day07 solution2:", solvePuzzle2(nums))
}

func solvePuzzle1(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += generate(n, 2000)
	}
	return sum
}

func solvePuzzle2(nums []int) int {
	all := map[key]int{}
	maxPrice := math.MinInt
	for _, n := range nums {
		for k, v := range prices(n) {
			price := all[k] + v
			all[k] = price
			if price > maxPrice {
				maxPrice = price
			}
		}
	}
	return maxPrice
}

func parseInput(input string) (parsed []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	parsed = make([]int, 0, len(lines))
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		parsed = append(parsed, num)
	}
	return
}

func nextSecret(num int) int {
	num = (num ^ num<<6) & 0xFFFFFF
	num = (num ^ num>>5) & 0xFFFFFF
	num = (num ^ num<<11) & 0xFFFFFF
	return num
}

func generate(n, iterations int) int {
	for range iterations {
		n = nextSecret(n)
	}
	return n
}

func prices(seed int) (prices map[key]int) {
	prices = make(map[key]int)
	diffs := make([]int, 0, 2000)
	n := seed
	prev := seed % 10
	for range 1999 {
		n = nextSecret(n)
		price := n % 10
		diffs = append(diffs, price-prev)
		prev = price
		if len(diffs) == 4 {
			k := key{diffs[0], diffs[1], diffs[2], diffs[3]}
			if _, ok := prices[k]; !ok {
				prices[k] = price
			}
			diffs = diffs[1:]
		}
	}
	return
}
