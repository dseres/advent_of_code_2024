package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `029A
980A
179A
456A
379A`

/*
029A: <vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A
980A: <v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A
179A: <v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
456A: <v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A
379A: <v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
*/

func TestToNum(t *testing.T) {
	a := assert.New(t)

	a.Equal("029A", toNums("<A^A>^^AvvvA", 0))
	a.Equal("029A", toNums("<A^A^>^AvvvA", 0))
	a.Equal("029A", toNums("<A^A^^>AvvvA", 0))

	a.Equal("029A", toNums("v<<A>>^A<A>AvA<^AA>A<vAAA>^A", 1))

	a.Equal("029A", toNums("<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A", 2))
	a.Equal("980A", toNums("<v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A", 2))
	a.Equal("179A", toNums("<v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A", 2))
	a.Equal("456A", toNums("<v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A", 2))
	a.Equal("379A", toNums("<v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A", 2))

	a.Equal("029A", toNums("v<A<AA>^>AvA^<Av>A^Av<<A>^>AvA^Av<<A>^>AAv<A>A^A<A>Av<A<A>^>AAA<Av>A^A", 2))
}

func TestRoutes(t *testing.T) {
	a := assert.New(t)
	routesFromA := searchRoutesFrom(numPad, strings.Index(numPad, "A"))
	a.Equal(2, len(routesFromA['5']))
	a.Equal(1, len(routesFromA['7']))
}

func TestConvert(t *testing.T) {
	a := assert.New(t)
	a.Equal(68, codeLength("A029A", 0, 2))
}

func TestSolution1(t *testing.T) {
	a := assert.New(t)
	a.Equal(126384, solvePuzzle(parseInput(testInput), 2))
	a.Equal(248684, solvePuzzle(parseInput(input), 2))
	a.Equal(307055584161760, solvePuzzle(parseInput(input), 25))
}
