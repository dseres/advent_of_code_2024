package main

import "testing"
import "github.com/stretchr/testify/assert"

var test_input = ``

func TestSolution1(t *testing.T) {
	result := solvePuzzle1(parseInput(test_input))
	a := assert.New(t)
	a.Equal(0, result)
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(parseInput(test_input))
	a := assert.New(t)
	a.Equal(0, result)
}
