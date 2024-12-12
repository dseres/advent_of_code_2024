package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput1 = `AAAA
BBCD
BBCC
EEEC
`

var testInput2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

var testInput3 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

var testInput4 = `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

var testInput5 = `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

func TestSolution1(t *testing.T) {
	a := assert.New(t)

	g := newGarden(testInput1)
	a.Equal(140, solvePuzzle1(g))

	g = newGarden(testInput2)
	a.Equal(772, solvePuzzle1(g))

	g = newGarden(testInput3)
	a.Equal(1930, solvePuzzle1(g))
}

func TestSolution2(t *testing.T) {
	a := assert.New(t)

	g := newGarden(testInput1)
	a.Equal(80, solvePuzzle2(g))

	g = newGarden(testInput4)
	a.Equal(236, solvePuzzle2(g))

	g = newGarden(testInput5)
	a.Equal(368, solvePuzzle2(g))

	g = newGarden(testInput3)
	a.Equal(22, g.computeSidesFor(2))
	a.Equal(12, g.computeSidesFor(3))
	a.Equal(1206, solvePuzzle2(g))
}
