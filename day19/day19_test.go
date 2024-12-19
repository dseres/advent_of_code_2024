package main

import "testing"
import "github.com/stretchr/testify/assert"

var testInput = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func TestSolution1(t *testing.T) {
	result := solvePuzzle1(parseInput(testInput))
	a := assert.New(t)
	a.Equal(6, result)
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(parseInput(testInput))
	a := assert.New(t)
	a.Equal(0, result)
}
