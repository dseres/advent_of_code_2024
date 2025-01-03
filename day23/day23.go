package main

import "fmt"
import _ "embed"
import "strings"

//go:embed input23.txt
var input string

func main() {
	g := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(g))
	fmt.Println("Day07 solution2:", solvePuzzle2(g))
}

func solvePuzzle1(g graph) int {
	return 0
}

func solvePuzzle2(g graph) int {
	return 0
}

type graph = map[int16][]int16

func parseInput(input string) (g graph) {
	g = make(graph)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		a, b := parseLine(line)
		insert(g, a, b)
		insert(g, b, a)
	}
	return
}

func parseLine(line string) (a, b int16) {
	s := strings.Split(line, "-")
	return str2Int16(s[0]), str2Int16(s[1])
}

func str2Int16(s string) int16 {
	return int16(s[0]-'a')*int16('z'-'a'+1) + int16(s[1]-'a')
}

func insert(g graph, from, to int16) {
	if _, ok := g[from]; ok {
		g[from] = append(g[from], to)
		return
	}
	g[from] = []int16{to}
}
