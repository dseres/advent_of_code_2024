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
		t.Errorf("Day01 puzzle1 solution on test input failed. Result: %v, Required: 11", result)
	}
}

func TestSolution2(t *testing.T) {
	nums1, nums2 := ParseInput(testInput)
	result := solvePuzzle2(nums1, nums2)
	if result != 31 {
		t.Errorf("Day01 puzzle2 solution on test input failed. Result: %v, Required: 31", result)
	}
}
