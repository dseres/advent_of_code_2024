package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("day01/input01.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat[:])
	nums1, nums2 := ParseInput(input)
	fmt.Println("Day01 solution1:", solvePuzzle1(nums1, nums2))
	fmt.Println("Day01 solution2:", solvePuzzle2(nums1, nums2))
}

func solvePuzzle1(nums1 []int, nums2 []int) (diff int64) {
	for i := range nums1 {
		if nums1[i] < nums2[i] {
			diff += int64(nums2[i] - nums1[i])
		} else {
			diff += int64(nums1[i] - nums2[i])
		}
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

func ParseInput(input string) (array1 []int, array2 []int) {
	for _, line := range strings.Split(input, "\n") {
		if len(line) > 0 {
			parts := strings.Fields(line)
			i, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			j, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			array1 = append(array1, i)
			array2 = append(array2, j)
		}
	}
	slices.Sort(array1)
	slices.Sort(array2)
	return
}
