package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = "125 17"

func TestSolution1(t *testing.T) {
	nums := parseInput(testInput)
	a := assert.New(t)
	a.Equal(22, solvePuzzle(nums, 6))
	a.Equal(55312, solvePuzzle(nums, 25))
}
