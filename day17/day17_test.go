package main

import (
	_ "fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput1 = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

func TestSmallerExamples(t *testing.T) {
	a := assert.New(t)

	m1 := machine{a: 0, b: 0, c: 9, instructions: []int{2, 6}, ip: 0, output: []int{}}
	m1.run()
	a.Equal(1, m1.b)

}

func TestSolution1(t *testing.T) {
	m := new(testInput1)
	result := solvePuzzle1(m)
	a := assert.New(t)
	a.Equal("4,6,3,5,6,3,5,2,1,0", result)
}

var testInput2 = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

// func TestSolution2(t *testing.T) {
// 	m := new(testInput2)
// 	result := solvePuzzle2(m)
// 	a := assert.New(t)
// 	a.Equal(117440, result)
// }

func Test_SliceComparison(t *testing.T) {
	a := assert.New(t)
	a.Equal(0, slices.Compare([]int{}, []int{}))
	a.Equal(0, slices.Compare([]int{1, 2, 3}, []int{1, 2, 3}))
	a.Equal(-1, slices.Compare([]int{1, 2, 3}, []int{1, 2, 3, 0}))
	a.Equal(-1, slices.Compare([]int{1, 2, 3}, []int{1, 3, 2}))
	a.Equal(0, slices.Compare([]int{}, []int{}))
	a.Equal(0, slices.Compare([]int{}, []int{}))
}

func TestSerie(t *testing.T) {
	m := new(input)
	a := assert.New(t)
	s := make([]int, 0, 16)
	serie(m.a, &s)
	a.Equal([]int{4, 1, 7, 6, 4, 1, 0, 2, 7}, s)
}

func BenchmarkSeries(b *testing.B) {
	s := make([]int, 0, 16)
	for i := 0; i < b.N; i++ {
		s = s[:0]
		serie(i, &s)
	}
}
