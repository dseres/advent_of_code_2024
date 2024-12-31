package main

import (
	"slices"
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

	a.Equal([]byte("029A"), toNums([]byte("<A^A>^^AvvvA"), 0))
	a.Equal([]byte("029A"), toNums([]byte("<A^A^>^AvvvA"), 0))
	a.Equal([]byte("029A"), toNums([]byte("<A^A^^>AvvvA"), 0))

	a.Equal([]byte("029A"), toNums([]byte("v<<A>>^A<A>AvA<^AA>A<vAAA>^A"), 1))

	a.Equal([]byte("029A"), toNums([]byte("<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A"), 2))
	a.Equal([]byte("980A"), toNums([]byte("<v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A"), 2))
	a.Equal([]byte("179A"), toNums([]byte("<v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A"), 2))
	a.Equal([]byte("456A"), toNums([]byte("<v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A"), 2))
	a.Equal([]byte("379A"), toNums([]byte("<v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A"), 2))

	a.Equal([]byte("029A"), toNums([]byte("v<A<AA>^>AvA^<Av>A^Av<<A>^>AvA^Av<<A>^>AAv<A>A^A<A>Av<A<A>^>AAA<Av>A^A"), 2))
}

func TestRoutes(t *testing.T) {
	a := assert.New(t)
	routesFromA := searchRoutesFrom(numPad, slices.Index(numPad, 'A'))
	a.Equal(4, len(routesFromA[56].positions))
	a.Equal(9, len(routesFromA[55].positions))

	printRoutes(numRoutes)
	printRoutes(dirRoutes)
}

func TestConvert(t *testing.T) {
	a := assert.New(t)
	a.Equal(68, convert([]byte("029A"), 'A', 0, 2, []byte{}))
}

func TestSolution1(t *testing.T) {
	result := solvePuzzle1(parseInput(testInput))
	a := assert.New(t)
	a.Equal(126384, result)
}

// func TestSolution2(t *testing.T) {
// 	result := solvePuzzle2(parseInput(testInput))
// 	a := assert.New(t)
// 	a.Equal(0, result)
// }
