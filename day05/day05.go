package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input05.txt
var input string

func main() {
	ordering, rules := parseInput(input)
	p := newPrinter(ordering)
	fmt.Println("Day05 solution1:", solvePuzzle1(p, rules))
	fmt.Println("Day05 solution2:", solvePuzzle2(p, rules))
}

func solvePuzzle1(p *printer, rules [][]int) (sum int) {
	for _, rule := range rules {
		if isValid(p, rule) {
			sum += rule[int(len(rule)/2)]
		}
	}
	return
}

func solvePuzzle2(p *printer, rules [][]int) (sum int) {
	for _, rule := range rules {
		if !isValid(p, rule) {
			valid, ok := p.impl(-1, rule)
			fmt.Println(rule, valid, ok)
			if ok {
				sum += valid[len(valid)/2]
			}
		}
	}
	return
}

type printer struct {
	m []bool
}

func newPrinter(ordering [][]int) *printer {
	p := printer{}
	p.m = make([]bool, 100*100, 100*100)
	for _, pair := range ordering {
		p.addOrdering(pair[0], pair[1])
	}
	return &p
}

func (p *printer) addOrdering(from, to int) {
	p.m[from*100+to] = true
}

func (p *printer) hasOrdering(from, to int) bool {
	return p.m[from*100+to]
}

func (p *printer) impl(prev int, nums []int) ([]int, bool) {
	if len(nums) == 0 {
		return nums, true
	}
	for _, n := range nums {
		if prev == -1 || p.hasOrdering(prev, n) {
			next := pop(nums, n)
			next, ok := p.impl(n, next)
			if ok {
				valid := make([]int, 0, len(nums))
				valid = append(valid, n)
				valid = append(valid, next...)
				return valid, true
			}
		}
	}
	return []int{}, false
}

func pop(nums []int, n int) (next []int) {

	next = make([]int, 0, len(nums))
	for _, m := range nums {
		if n != m {
			next = append(next, m)
		}
	}
	return
}

func isValid(p *printer, rule []int) (v bool) {
	v = true
	for i, page := range rule {
		for _, next := range rule[i+1:] {
			if !p.hasOrdering(page, next) {
				return false
			}

		}
	}
	return
}

func parseInput(input string) ([][]int, [][]int) {
	parts := strings.Split(input, "\n\n")
	return parseOrdering(parts[0]), parseRules(parts[1])
}

func parseOrdering(input string) (ordering [][]int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			pair := parseLine(line, "|")
			ordering = append(ordering, pair)
		}
	}
	return
}

func parseRules(input string) (rules [][]int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			pages := parseLine(line, ",")
			rules = append(rules, pages)
		}
	}
	return
}

func parseLine(line string, separator string) (report []int) {
	fields := strings.Split(line, separator)
	for _, field := range fields {
		val, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		report = append(report, val)
	}
	return
}
