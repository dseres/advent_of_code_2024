package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func TestSolution1(t *testing.T) {
	a := assert.New(t)
	m := new(testInput)
	m.findPath()
	fmt.Println(m)
	a.Equal(1, m.countShortcuts(64))
	a.Equal(2, m.countShortcuts(40))
	a.Equal(3, m.countShortcuts(38))
	a.Equal(44, m.countShortcuts(2))
}

// func TestSolution2(t *testing.T) {
// 	result := solvePuzzle2(parseInput(testInput))
// 	a := assert.New(t)
// 	a.Equal(0, result)
// }
