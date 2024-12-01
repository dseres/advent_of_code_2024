package main

import "testing"

var testInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestSolution1(t *testing.T) {
	nums1, nums2 := ParseInput(testInput)
	result := solvePuzzle1(nums1, nums2)
	if result != 11 {
		t.Errorf("Day01 puzzle1 solution on test input failed. Result: %v", result)
	}
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(testInput)
	if result != 0 {
		t.Errorf("Day01 puzzle2 solution on test input failed. Result: %v", result)
	}
}
