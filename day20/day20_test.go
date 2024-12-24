package main

import (
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

	a.Equal(0, distance(point{1, 1}, point{1, 1}))
	a.Equal(1, distance(point{1, 1}, point{1, 0}))
	a.Equal(2, distance(point{3, 3}, point{1, 3}))
	a.Equal(2, distance(point{3, 3}, point{3, 1}))
	a.Equal(2, distance(point{3, 3}, point{2, 2}))

	a.Equal(1, m.countShortcuts(64, 2))
	a.Equal(2, m.countShortcuts(40, 2))
	a.Equal(3, m.countShortcuts(38, 2))
	a.Equal(44, m.countShortcuts(2, 2))

	a.Equal(3, m.countShortcuts(76, 20))
	a.Equal(7, m.countShortcuts(74, 20))
	a.Equal(29, m.countShortcuts(72, 20))
}
