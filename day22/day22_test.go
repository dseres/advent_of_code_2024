package main

import "testing"
import "github.com/stretchr/testify/assert"

var testInput = `1
10
100
2024`

var testResults = []int{8685429, 4700978, 15273692, 8667524}

func TestNextSecret(t *testing.T) {
	a := assert.New(t)
	num := 123
	var secrets = []int{15887950, 16495136, 527345, 704524, 1553684, 12683156, 11100544, 12249484, 7753432, 5908254}
	for _, s := range secrets {
		num = nextSecret(num)
		a.Equal(num, s)
	}

	nums := parseInput(testInput)
	for i, n := range nums {
		for range 2000 {
			n = nextSecret(n)
		}
		a.Equal(testResults[i], n)
	}
}

func TestSolution1(t *testing.T) {
	result := solvePuzzle1(parseInput(testInput))
	a := assert.New(t)
	a.Equal(37327623, result)
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(parseInput(testInput))
	a := assert.New(t)
	a.Equal(0, result)
}
