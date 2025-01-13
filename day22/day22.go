package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input22.txt
var input string

func main() {
	start := time.Now()
	nums := parseInput(input)
	fmt.Println("Day22 solution1:", solvePuzzle1(nums))
	fmt.Println("Day22 solution2:", solvePuzzle2(nums))
	fmt.Println("Day22 time:", time.Now().Sub(start))
}

func solvePuzzle1(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += generate(n, 2000)
	}
	return sum
}

func solvePuzzle2(nums []int) int {
	var bananas = make([]int, 1<<20)

	for _, n := range nums {
		prices(bananas, n)
	}
	return slices.Max(bananas)
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

func prices(bananas []int, seed int) {
	var visited = make([]bool, 1<<20)
	k := int32(0)
	for i := range 2000 {
		next := nextSecret(seed)
		diff := next%10 - seed%10 + 10
		// diff is between 1 and 19. A 5 bit long integer is enough to store.
		k = ((k & (1<<15 - 1)) << 5) | int32(diff)
		if i >= 3 && !visited[k] {
			visited[k] = true
			bananas[k] += next % 10
		}
		seed = next
	}
}
