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
	fmt.Println(patterns, strs)
	fmt.Println("Day07 solution1:", solvePuzzle1(patterns, strs))
	fmt.Println("Day07 solution2:", solvePuzzle2(patterns, strs))
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

func solvePuzzle2(patterns []string, strs []string) int {
	return 0
}

func parseInput(input string) ([]string, []string) {
	parts := strings.Split(strings.TrimRight(input, " \t\r\n"), "\n\n")
	return strings.Split(parts[0], ", "), strings.Split(parts[1], "\n")
}
