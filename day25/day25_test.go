package main

import "testing"
import "github.com/stretchr/testify/assert"

var testInput = `#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`

func TestSolution1(t *testing.T) {
	a := assert.New(t)
	locks, keys := parseInput(testInput)
	a.Equal([][]int{[]int{0, 5, 3, 4, 3}, []int{1, 2, 0, 5, 3}}, locks)
	a.Equal([][]int{[]int{5, 0, 2, 1, 3}, []int{4, 3, 4, 0, 2}, []int{3, 0, 2, 0, 1}}, keys)
	result := solvePuzzle1(locks, keys)
	a.Equal(3, result)
}
