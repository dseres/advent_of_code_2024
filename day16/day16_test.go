package main

import "testing"
import "github.com/stretchr/testify/assert"

var testInput1 = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

var testInput2 = `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`

func TestSolution1(t *testing.T) {
	a := assert.New(t)
	a.Equal(7036, solvePuzzle1(parseInput(testInput1)))
	a.Equal(11048, solvePuzzle1(parseInput(testInput2)))
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(parseInput(testInput1))
	a := assert.New(t)
	a.Equal(0, result)
}
