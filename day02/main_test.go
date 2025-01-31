package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var test_input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestSolution1(t *testing.T) {
	reports := parseInput(test_input)
	assert := assert.New(t)
	result := solvePuzzle1(reports)
	assert.Equal(int64(2), result)
}

func TestSolution2(t *testing.T) {
	reports := parseInput(test_input)
	result := solvePuzzle2(reports)
	assert := assert.New(t)
	assert.Equal(int64(4), result)
}

func TestIsSafe(t *testing.T) {
	assert := assert.New(t)
	assert.True(isSafe([]int{7, 6, 4, 2, 1}))
	assert.False(isSafe([]int{1, 2, 7, 8, 9}))
	assert.False(isSafe([]int{9, 7, 6, 2, 1}))
	assert.False(isSafe([]int{1, 3, 2, 4, 5}))
	assert.False(isSafe([]int{8, 6, 4, 4, 1}))
	assert.True(isSafe([]int{1, 3, 6, 7, 9}))
}

func TestIsDampenedSafe(t *testing.T) {
	assert := assert.New(t)
	assert.True(isDampenedSafe([]int{7, 6, 4, 2, 1}))
	assert.False(isDampenedSafe([]int{1, 2, 7, 8, 9}))
	assert.False(isDampenedSafe([]int{9, 7, 6, 2, 1}))
	assert.True(isDampenedSafe([]int{1, 3, 2, 4, 5}))
	assert.True(isDampenedSafe([]int{8, 6, 4, 4, 1}))
	assert.True(isDampenedSafe([]int{1, 3, 6, 7, 9}))
}
