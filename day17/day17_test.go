package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

func TestSmallerExamples(t *testing.T) {
	a := assert.New(t)

	m1 := machine{a: 0, b: 0, c: 9, instructions: []int{2, 6}, ip: 0, output: []int{}}
	m1.run()
	fmt.Println(m1)
	a.Equal(1, m1.b)

}

func TestSolution1(t *testing.T) {
	m := new(testInput)
	result := solvePuzzle1(m)
	a := assert.New(t)
	a.Equal("4,6,3,5,6,3,5,2,1,0", result)
}

func TestSolution2(t *testing.T) {
	m := new(testInput)
	result := solvePuzzle2(m)
	a := assert.New(t)
	a.Equal(0, result)
}
