package main

import "testing"

var test_input = ""

func TestSolution1(t *testing.T) {
	result := solvePuzzle1(test_input)
	if result != 0 {
		t.Errorf("Day16 puzzle1 solution on test input failed. Result: %v", result)
	}
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(test_input)
	if result != 0 {
		t.Errorf("Day16 puzzle2 solution on test input failed. Result: %v", result)
	}
}
