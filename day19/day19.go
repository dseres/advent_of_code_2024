package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input19.txt
var input string

func main() {
	patterns, strs := parseInput(input)
	count1, count2 := solve(patterns, strs)
	fmt.Println("Day07 solution1:", count1)
	fmt.Println("Day07 solution2:", count2)
}

func solvePuzzle1(patterns []string, strs []string) int {
	count := 0
	re := regexp.MustCompile("^(" + strings.Join(patterns, "|") + ")*$")
	for _, str := range strs {
		if re.MatchString(str) {
			count++
		}
	}
	return count
}

func parseInput(input string) ([]string, []string) {
	parts := strings.Split(strings.TrimRight(input, " \t\r\n"), "\n\n")
	return strings.Split(parts[0], ", "), strings.Split(parts[1], "\n")
}

func solve(patterns []string, strs []string) (count1 int, count2 int) {
	pmap := map[string]bool{}
	for _, p := range patterns {
		pmap[p] = true
	}
	for _, str := range strs {
		c1, c2 := check(pmap, str)
		count2 += c2
		if c1 {
			count1++
		}
	}
	return
}

type counter struct {
	found bool
	count int
}

var memo map[string]counter = make(map[string]counter)

func check(patterns map[string]bool, str string) (found bool, count int) {
	if len(str) == 0 {
		return true, 1
	}
	for i := 1; i <= len(str); i++ {
		if _, ok := patterns[str[:i]]; !ok {
			continue
		}
		next := str[i:]
		// check memo
		if m, ok := memo[next]; ok {
			found = found || m.found
			count += m.count
		} else {
			f, c := check(patterns, next)
			memo[next] = counter{f, c}
			found = found || f
			count += c
		}
	}
	return
}
