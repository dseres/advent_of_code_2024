package main

import "testing"
import "github.com/stretchr/testify/assert"

var testInput = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func TestSolution1(t *testing.T) {
	counter1, counter2 := solve(parseInput(testInput))
	a := assert.New(t)
	a.Equal(6, counter1)
	a.Equal(16, counter2)
}
