package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"
)

//go:embed input01.txt
var input string

func main() {
	start := time.Now()
	nums1, nums2 := parseInput(input)
	fmt.Println("Day01 solution1:", solvePuzzle1(nums1, nums2))
	fmt.Println("Day01 solution2:", solvePuzzle2WithMap(nums1, nums2))
	fmt.Println("Day01 time:", time.Now().Sub(start))
}

func solvePuzzle1(nums1 []int, nums2 []int) (diff int64) {
	for i := range nums1 {
		diff += int64(abs(nums1[i] - nums2[i]))
	}
	return
}

func solvePuzzle2(nums1 []int, nums2 []int) (similarity int64) {
	for _, n := range nums1 {
		var count int64 = 0
		for _, m := range nums2 {
			if n == m {
				count += 1
			}
		}
		similarity += int64(n) * count
	}
	return
}

func solvePuzzle2WithMap(nums1 []int, nums2 []int) (similarity int64) {
	m := createCountMap(nums2)
	for _, n := range nums1 {
		similarity += int64(n * m[n])
	}
	return
}

func createCountMap(nums2 []int) map[int]int {
	m := make(map[int]int)
	for _, n := range nums2 {
		m[n] += 1
	}
	return m
}

func parseInput(input string) (array1 []int, array2 []int) {
	for _, line := range strings.Split(input, "\n") {
		if len(line) > 0 {
			var i, j int
			fmt.Sscanf(line, "%d %d", &i, &j)
			array1 = append(array1, i)
			array2 = append(array2, j)
		}
	}
	slices.Sort(array1)
	slices.Sort(array2)
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
