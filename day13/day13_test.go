package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
`

func TestSolution(t *testing.T) {
	cms := parseInput(testInput)
	a := assert.New(t)

	a.Equal(480, solvePuzzle(cms, getMinBruteForce))
	a.Equal(480, solvePuzzle(cms, getMinUsingEquations))

	increasePrizes(cms)
	_, ok0 := getMinUsingEquations(cms[0])
	a.False(ok0)
	_, ok1 := getMinUsingEquations(cms[1])
	a.True(ok1)
	_, ok2 := getMinUsingEquations(cms[2])
	a.False(ok2)
	_, ok3 := getMinUsingEquations(cms[3])
	a.True(ok3)
}
