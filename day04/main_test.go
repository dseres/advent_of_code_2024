package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestSolution1(t *testing.T) {
	text := parseInput([]byte(testInput))
	a := assert.New(t)
	a.Equal(18, solvePuzzle1(text))

}

func TestSolution2(t *testing.T) {
	text := parseInput([]byte(testInput))
	a := assert.New(t)
	a.Equal(9, solvePuzzle2(text))

}
