package main

import "testing"
import "github.com/stretchr/testify/assert"

var testInput = "125 17"

func TestSolution1(t *testing.T) {
	nums := parseInput(testInput)
	a := assert.New(t)
	a.Equal(22, solvePuzzle1(nums, 6))
	a.Equal(55312, solvePuzzle1(nums, 25))
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(parseInput(testInput), 0)
	a := assert.New(t)
	a.Equal(0, result)
}
