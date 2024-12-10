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
	result := solvePuzzle1(parseInput(testInput))
	a := assert.New(t)
	a.Equal(36, result)
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(parseInput(testInput))
	a := assert.New(t)
	a.Equal(0, result)
}
