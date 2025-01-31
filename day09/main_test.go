package main

import "testing"
import "github.com/stretchr/testify/assert"

var testInput = "2333133121414131402"

func TestSolution1(t *testing.T) {
	blocks := toBlocks(parseInput(testInput))
	result := solvePuzzle1(blocks)
	a := assert.New(t)
	a.Equal(1928, result)
}

func TestSolution2(t *testing.T) {
	blocks := toBlocks(parseInput(testInput))
	result := solvePuzzle2(blocks)
	a := assert.New(t)
	a.Equal(2858, result)
}
