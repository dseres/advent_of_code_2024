package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput1 = `..........
..........
..........
....a.....
..........
.....a....
..........
..........
..........
..........
`

var testInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

func TestSolution1(t *testing.T) {
	a := assert.New(t)
	f := newField(parseInput(testInput1))
	a.Equal([]point{{1, 3}, {7, 6}}, f.computeAntinodes(f.antennas['a']))
	a.Equal(2, solvePuzzle1(f))

	f = newField(parseInput(testInput))
	a.Equal(14, solvePuzzle1(f))
}

func TestSolution2(t *testing.T) {
	a := assert.New(t)
	f := newField(parseInput(testInput))
	result := solvePuzzle2(f)
	a.Equal(34, result)
}
