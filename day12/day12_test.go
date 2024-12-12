package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestSolution1(t *testing.T) {
	a := assert.New(t)

	g := parseInput(testInput1)
	fmt.Println(g)
	a.Equal(140, solvePuzzle1(g))

	g = parseInput(testInput2)
	fmt.Println(g)
	a.Equal(772, solvePuzzle1(g))

	g = parseInput(testInput3)
	fmt.Println(g)
	a.Equal(1930, solvePuzzle1(g))
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(parseInput(testInput1))
	a := assert.New(t)
	a.Equal(0, result)
}
