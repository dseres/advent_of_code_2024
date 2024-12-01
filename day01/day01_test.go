package main

import "testing"

var testInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestSolution1(t *testing.T) {
	nums1, nums2 := parseInput(testInput)
	result := solvePuzzle1(nums1, nums2)
	if result != 11 {
		t.Errorf("Day01 puzzle1 solution on test input failed. Result: %v, Required: 11", result)
	}
}

func TestSolution2(t *testing.T) {
	nums1, nums2 := parseInput(testInput)
	result := solvePuzzle2WithMap(nums1, nums2)
	if result != 31 {
		t.Errorf("Day01 puzzle2 solution on test input failed. Result: %v, Required: 31", result)
	}
}

func BenchmarkSolvePuzzle2(b *testing.B) {
	nums1, nums2 := parseInput(input)
	for i := 0; i < b.N; i += 1 {
		solvePuzzle2(nums1, nums2)
	}
}

func BenchmarkSolvePuzzle2WithMap(b *testing.B) {
	nums1, nums2 := parseInput(input)
	for i := 0; i < b.N; i += 1 {
		solvePuzzle2WithMap(nums1, nums2)
	}
}
