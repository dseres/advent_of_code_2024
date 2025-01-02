package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var nums = []int{1, 10, 100, 2024}

func TestNextSecret(t *testing.T) {
	a := assert.New(t)
	num := 123
	var secrets = []int{15887950, 16495136, 527345, 704524, 1553684, 12683156, 11100544, 12249484, 7753432, 5908254}
	for _, s := range secrets {
		num = nextSecret(num)
		a.Equal(num, s)
	}
}

func TestGenerate(t *testing.T) {
	a := assert.New(t)
	var requirements = []int{8685429, 4700978, 15273692, 8667524}
	for i, n := range nums {
		a.Equal(requirements[i], generate(n, 2000))
	}
}

func TestSolution1(t *testing.T) {
	result := solvePuzzle1(nums)
	a := assert.New(t)
	a.Equal(37327623, result)
	a.Equal(21147129593, solvePuzzle1(parseInput(input)))
}

func TestSequences(t *testing.T) {
	a := assert.New(t)
	a.Equal(7, prices(1)[key{-2, 1, -1, 3}])
	a.Equal(7, prices(2)[key{-2, 1, -1, 3}])
	a.Equal(0, prices(3)[key{-2, 1, -1, 3}])
	a.Equal(9, prices(2024)[key{-2, 1, -1, 3}])
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2([]int{1, 2, 3, 2024})
	a := assert.New(t)
	a.Equal(23, result)
}
