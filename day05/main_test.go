package main

import (
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_input = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestSolution1(t *testing.T) {
	ordering, rules := parseInput(test_input)
	p := newPrinter(ordering)
	result := solvePuzzle1(p, rules)
	a := assert.New(t)
	a.Equal(143, result)
}

func TestSolution2(t *testing.T) {
	ordering, rules := parseInput(test_input)
	p := newPrinter(ordering)
	result := solvePuzzle2(p, rules)
	a := assert.New(t)
	a.Equal(123, result)
}
