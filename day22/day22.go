package main

import "fmt"
import _ "embed"
import "strings"
import "strconv"

//go:embed input22.txt
var input string

func main() {
	nums := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(nums))
	fmt.Println("Day07 solution2:", solvePuzzle2(nums))
}

func solvePuzzle1(nums []int) int {
	sum := 0
	for i := range nums {
		for range 2000 {
			nums[i] = nextSecret(nums[i])
		}
		sum += nums[i]
	}
	return sum
}

func solvePuzzle2(nums []int) int {
	return 0
}

func parseInput(input string) (parsed []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		if len(line) > 0 {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			parsed = append(parsed, num)
		}
	}
	return
}

func nextSecret(num int) int {
	num = (num ^ (num * 64)) % 16777216
	num = (num ^ int(float64(num)/32.0)) % 16777216
	num = (num ^ (num * 2048)) % 16777216
	return num
}
