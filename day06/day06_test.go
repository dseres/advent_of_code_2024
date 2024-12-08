package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_input = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestSolution1(t *testing.T) {
	result := solvePuzzle1(parseInput(test_input))
	a := assert.New(t)
	a.Equal(41, result)
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(parseInput(test_input))
	a := assert.New(t)
	a.Equal(6, result)
}

func TestLoop0(t *testing.T) {
	var input = `.#.
.^#
...
#..
.#.`
	lab, pos := parseInput(input)
	loop, count := walk(0, pos, lab)
	a := assert.New(t)
	a.True(loop)
	a.Equal(3, count)
	printMap(lab)
}

func TestLoop1(t *testing.T) {
	var input = `.#.
.^#
#..
.#.`
	lab, pos := parseInput(input)
	loop, count := walk(0, pos, lab)
	a := assert.New(t)
	a.True(loop)
	a.Equal(2, count)
	printMap(lab)
}

func TestLoop2(t *testing.T) {
	a := assert.New(t)
	var input = `.#.
#^#
.#.`
	lab, pos := parseInput(input)
	printMap(lab)
	a.Panics(func() {
		walk(0, pos, lab)
	})
}

func TestLoop4(t *testing.T) {
	var input = `.#..
.^.#
#...
..#.`
	lab, pos := parseInput(input)
	loop, count := walk(0, pos, lab)
	a := assert.New(t)
	a.True(loop)
	a.Equal(4, count)
	printMap(lab)
}
