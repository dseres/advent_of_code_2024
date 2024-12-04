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

var testOutput2 = `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

func TestSolution1(t *testing.T) {
	text := parseInput([]byte(testInput))
	rep, count := solvePuzzle1(text)
	printText(rep)
	a := assert.New(t)
	a.Equal(int64(18), count)
}

func TestSolution2(t *testing.T) {
	text := parseInput([]byte(testInput))
	rep, count := solvePuzzle2(text)
	printText(rep)
	a := assert.New(t)
	a.Equal(int64(9), count)
}
