package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input21.txt
var input string

var numPad string = "789\n456\n123\nX0A"
var dirPad string = "X^A\n<v>"
var numRoutes map[rune]map[rune][]string = getRoutes(numPad)
var dirRoutes map[rune]map[rune][]string = getRoutes(dirPad)

type key struct {
	level int
	code  string
}

var cache map[key]int = make(map[key]int)

func main() {
	codes := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle(codes, 2))
	fmt.Println("Day07 solution2:", solvePuzzle(codes, 25))
}

func solvePuzzle(codes []string, level int) int {
	cache = make(map[key]int)
	sum := 0
	for _, code := range codes {
		sum += numOf(code) * codeLength("A"+code, 0, level)
	}
	return sum
}

func parseInput(input string) (parsed []string) {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func getRoutes(pad string) (routes map[rune]map[rune][]string) {
	routes = make(map[rune]map[rune][]string)
	lines := strings.Split(pad, "\n")
	// Iterating over characters as `from`
	for y0, l0 := range lines {
		for x0, from := range l0 {
			if from == 'X' {
				continue
			}
			routes[from] = make(map[rune][]string)
			// Iterating over characters as `to`
			for y1, l1 := range lines {
				for x1, to := range l1 {
					if to == 'X' {
						continue
					}
					routes[from][to] = getRoutesFor(lines, x0, y0, x1, y1)
				}
			}
		}
	}
	return
}

func getRoutesFor(lines []string, x0, y0, x1, y1 int) []string {
	dx := x1 - x0
	xpart := ">>"[:abs(dx)]
	if dx < 0 {
		xpart = "<<"[:abs(dx)]
	}
	dy := y1 - y0
	ypart := "vvv"[:abs(dy)]
	if dy < 0 {
		ypart = "^^^"[:abs(dy)]
	}
	rs := make([]string, 0, 2)
	if lines[y1][x0] != 'X' {
		rs = append(rs, ypart+xpart+"A")
	}
	if lines[y0][x1] != 'X' {
		r := xpart + ypart + "A"
		if len(rs) == 0 || r != rs[0] {
			rs = append(rs, xpart+ypart+"A")
		}
	}
	return rs
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func numOf(code string) int {
	n, err := strconv.Atoi(code[:len(code)-1])
	if err != nil {
		panic(fmt.Sprintf("Bad number format in code %v", code))
	}
	return n
}

func codeLength(code string, level, maxLevel int) int {
	if level > maxLevel {
		return len(code) - 1
	}
	routes := dirRoutes
	if level == 0 {
		routes = numRoutes
	}
	k := key{level, code}
	if v, ok := cache[k]; ok {
		return v
	}
	sum := 0
	for i, from := range code[:len(code)-1] {
		to := rune(code[i+1])
		routes := routes[from][to]
		min := math.MaxInt
		for _, r := range routes {
			len := codeLength("A"+r, level+1, maxLevel)
			if len < min {
				min = len
			}
		}
		sum += min
	}
	cache[k] = sum
	return sum
}
