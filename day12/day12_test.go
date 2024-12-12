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
	p := parseInput(testInput1)
	computeRegions(p)
	fmt.Println(p)
	printPlots(p)

	result := solvePuzzle1(parseInput(testInput1))
	a := assert.New(t)
	a.Equal(0, result)
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(parseInput(testInput1))
	a := assert.New(t)
	a.Equal(0, result)
}
