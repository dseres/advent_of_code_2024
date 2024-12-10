package main

import "testing"
import "github.com/stretchr/testify/assert"

var testInput = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

func TestSolution1(t *testing.T) {
	result1, result2 := solvePuzzle1(parseInput(testInput))
	a := assert.New(t)
	a.Equal(36, result1)
	a.Equal(81, result2)
}
