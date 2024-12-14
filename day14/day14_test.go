package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

	w, h := 11, 7
	r := robot{2, 4, 2, -3}

	x, y := getPositionAfter(r, 4, w, h)
	a.Equal(10, x)
	a.Equal(6, y)
	x, y = getPositionAfter(r, 5, w, h)
	a.Equal(1, x)
	a.Equal(3, y)
}

func TestSolution1(t *testing.T) {
	a := assert.New(t)
	robots := parseInput(testInput)
	a.Equal(12, solvePuzzle1(robots))
}

func TestStep(t *testing.T) {
	a := assert.New(t)
	robots := parseInput(testInput)
	w, h := getDimension(robots)
	for range 100 {
		step(w, h, robots)
	}
	robots2 := parseInput(testInput)
	changePostion(w, h, robots2, 100)
	a.Equal(robots, robots2)
}
