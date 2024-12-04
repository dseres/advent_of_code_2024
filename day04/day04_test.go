package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

var testOutput = `....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX`

func TestSolution1(t *testing.T) {
	text := parseInput([]byte(testInput))
	printText(text)
	rep, count := solvePuzzle1(text)
	fmt.Println()
	printText(rep)
	a := assert.New(t)
	a.Equal(int64(18), count)
}

func TestSolution2(t *testing.T) {
	text := parseInput([]byte(testInput))
	result := solvePuzzle2(text)
	a := assert.New(t)
	a.Equal(int64(0), result)
}
