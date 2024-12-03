package main

import "testing"
import "github.com/stretchr/testify/assert"

var test_input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
var test_input2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestSolution1(t *testing.T) {
	result := solvePuzzle1(test_input)
	assert.New(t).Equal(int64(161), result)
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(test_input2)
	assert.New(t).Equal(int64(48), result)
}
