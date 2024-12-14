package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`

func TestGetPosition(t *testing.T) {
	a := assert.New(t)

	rect := robot{0, 0, 10, 6}
	r := robot{2, 4, 2, -3}

	x, y := getPositionAfter(r, 4, rect)
	a.Equal(10, x)
	a.Equal(6, y)
	x, y = getPositionAfter(r, 5, rect)
	a.Equal(1, x)
	a.Equal(4, y)
}

func TestSolution1(t *testing.T) {
	a := assert.New(t)
	robots := parseInput(testInput)
	fmt.Println(robots)
	a.Equal(12, solvePuzzle1(robots))
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(parseInput(testInput))
	a := assert.New(t)
	a.Equal(0, result)
}
