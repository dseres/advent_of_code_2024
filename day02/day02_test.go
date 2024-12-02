package main

import (
	"testing"
)

var test_input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestSolution1(t *testing.T) {
	reports := parseInput(test_input)
	result := solvePuzzle1(reports)
	if result != 2 {
		t.Errorf("Day02 puzzle1 solution on test input failed. Result: %v, Required: 2", result)
	}
}

func TestSolution2(t *testing.T) {
	reports := parseInput(test_input)
	result := solvePuzzle2(reports)
	if result != 4 {
		t.Errorf("Day02 puzzle2 solution on test input failed. Result: %v, Required: 4", result)
	}
}

func TestIsSafe(t *testing.T) {
	report := []int{7, 6, 4, 2, 1}
	safe := isSafe(report)
	if !safe {
		t.Errorf("Day02 isSafe gave bad results %v for %v", safe, report)
	}
	report = []int{1, 2, 7, 8, 9}
	safe = isSafe(report)
	if safe {
		t.Errorf("Day02 isSafe gave bad results %v for %v", safe, report)
	}
	report = []int{9, 7, 6, 2, 1}
	safe = isSafe(report)
	if safe {
		t.Errorf("Day02 isSafe gave bad results %v for %v", safe, report)
	}
	report = []int{1, 3, 2, 4, 5}
	safe = isSafe(report)
	if safe {
		t.Errorf("Day02 isSafe gave bad results %v for %v", safe, report)
	}
	report = []int{8, 6, 4, 4, 1}
	safe = isSafe(report)
	if safe {
		t.Errorf("Day02 isSafe gave bad results %v for %v", safe, report)
	}
	report = []int{1, 3, 6, 7, 9}
	safe = isSafe(report)
	if !safe {
		t.Errorf("Day02 isSafe gave bad results %v for %v", safe, report)
	}
}

func TestIsDampenedSafe(t *testing.T) {
	report := []int{7, 6, 4, 2, 1}
	safe := isDampenedSafe(report)
	if !safe {
		t.Errorf("Day02 isDampenedSafe gave bad results %v for %v", safe, report)
	}
	report = []int{1, 2, 7, 8, 9}
	safe = isDampenedSafe(report)
	if safe {
		t.Errorf("Day02 isDampenedSafe gave bad results %v for %v", safe, report)
	}
	report = []int{9, 7, 6, 2, 1}
	safe = isDampenedSafe(report)
	if safe {
		t.Errorf("Day02 isDampenedSafe gave bad results %v for %v", safe, report)
	}
	report = []int{1, 3, 2, 4, 5}
	safe = isDampenedSafe(report)
	if !safe {
		t.Errorf("Day02 isDampenedSafe gave bad results %v for %v", safe, report)
	}
	report = []int{8, 6, 4, 4, 1}
	safe = isDampenedSafe(report)
	if !safe {
		t.Errorf("Day02 isDampenedSafe gave bad results %v for %v", safe, report)
	}
	report = []int{1, 3, 6, 7, 9}
	safe = isDampenedSafe(report)
	if !safe {
		t.Errorf("Day02 isDampenedSafe gave bad results %v for %v", safe, report)
	}
}
