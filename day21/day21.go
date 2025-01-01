package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input21.txt
var input string

var numPad []byte = []byte("789456123X0A")
var dirPad []byte = []byte("X^A<v>")
var numRoutes map[byte]map[byte][]string = getRoutes(numPad)
var dirRoutes map[byte]map[byte][]string = getRoutes(dirPad)

type key struct {
	level int
	code  string
}

var cache map[key]int = make(map[key]int)

func main() {
	codes := parseInput(input)
	fmt.Println(codes)
	fmt.Println("Day07 solution1:", solvePuzzle(codes, 2))
	fmt.Println("Day07 solution2:", solvePuzzle(codes, 25))
}

func solvePuzzle(codes []string, level int) int {
	cache = make(map[key]int)
	sum := 0
	for _, code := range codes {
		sum += numOf(code) * convert("A"+code, 0, level)
	}
	return sum
}

func parseInput(input string) (parsed []string) {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func toNums(input string, level int) string {
	var pad []byte = dirPad
	if level == 0 {
		pad = numPad
	}
	pos := slices.Index(pad, 'A')
	output := []byte{}
	for _, d := range input {
		switch d {
		case '^':
			pos -= 3
		case '<':
			pos--
		case '>':
			pos++
		case 'v':
			pos += 3
		case 'A':
			if level == 0 {
				output = append(output, pad[pos])
			} else {
				output = append(output, pad[pos])
			}
		default:
			panic(fmt.Sprintf("bad instruction %v", d))
		}
	}
	if level > 0 {
		return toNums(string(output), level-1)
	}
	return string(output)
}

type routes struct {
	from, to  byte
	positions [][]int
}

type node struct {
	to            byte
	index, length int
	prevs         []int
}

func getRoutes(pad []byte) map[byte]map[byte][]string {
	rs := make(map[byte]map[byte][]string)
	for i, from := range pad {
		if from == 'X' {
			continue
		}
		rs[from] = searchRoutesFrom(pad, i)
	}
	return rs
}

func searchRoutesFrom(pad []byte, from int) map[byte][]string {
	visited := make(map[byte]routes)
	visited[pad[from]] = routes{from: pad[from], to: pad[from], positions: [][]int{{from}}}
	reachables := make(map[byte]node)
	findReachables(pad, from, 0, visited, reachables)
	// fmt.Println(pad, from, visited, reachables)
	for len(reachables) > 0 {
		n, min := node{}, math.MaxInt
		for _, rn := range reachables {
			if rn.length < min {
				n = rn
				min = rn.length
			}
		}
		// fmt.Println(reachables, n)
		r := routes{from: pad[from], to: n.to}
		for _, prev := range n.prevs {
			v := visited[pad[n.index-prev]]
			for _, ps := range v.positions {
				ps2 := slices.Clone(ps)
				ps2 = append(ps2, prev)
				r.positions = append(r.positions, ps2)
			}
		}
		visited[n.to] = r
		delete(reachables, n.to)
		findReachables(pad, n.index, n.length, visited, reachables)
	}
	// Convert directions to string
	routes := make(map[byte][]string)
	for to, rs := range visited {
		strs := make([]string, 0, len(rs.positions))
		for _, r := range rs.positions {
			builder := strings.Builder{}
			for _, d := range r[1:] {
				switch d {
				case -3:
					builder.WriteRune('^')
				case 3:
					builder.WriteRune('v')
				case -1:
					builder.WriteRune('<')
				case 1:
					builder.WriteRune('>')
				}
			}
			builder.WriteRune('A')
			strs = append(strs, builder.String())
		}
		// filter only the best routes by rank
		maxRank := math.MinInt
		for _, s := range strs {
			r := rank(s)
			switch {
			case maxRank < r:
				routes[to] = []string{s}
				maxRank = r
			case r == maxRank:
				routes[to] = append(routes[to], s)
			}
		}
	}
	return routes
}

func findReachables(pad []byte, fromInd, fromLength int, visited map[byte]routes, reachables map[byte]node) {
	dirs := []int{-3, 3, -1, 1}
	for _, dir := range dirs {
		next := fromInd + dir
		if dir == -1 && fromInd%3 == 0 || dir == 1 && fromInd%3 == 2 {
			continue
		}
		if next < 0 || len(pad) <= next || pad[next] == 'X' {
			continue
		}
		if _, ok := visited[pad[next]]; ok {
			continue
		}
		r, ok := reachables[pad[next]]
		if ok && fromLength+1 < r.length {
			r.length = fromLength + 1
			r.prevs = []int{dir}
			reachables[pad[next]] = r
			continue
		}
		if ok && r.length == fromLength+1 {
			r.prevs = append(r.prevs, dir)
			reachables[pad[next]] = r
			continue
		}
		if !ok {
			r.to = pad[next]
			r.length = fromLength + 1
			r.prevs = []int{dir}
			r.index = next
			reachables[r.to] = r
		}
	}
}

func printRoutes(all map[byte]map[byte][]string) {
	for from, rf := range all {
		fmt.Printf("From %v:\n", string(from))
		for to, rs := range rf {
			fmt.Printf("\tto %v: %v\n", string(to), rs)
		}
	}
}

func numOf(code string) int {
	n, err := strconv.Atoi(code[:len(code)-1])
	if err != nil {
		panic(fmt.Sprintf("Bad number format in code %v", code))
	}
	return n
}

// Change BFS to DFS
func convert(code string, level, maxLevel int) int {
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
	for i := range code[:len(code)-1] {
		from := code[i]
		to := code[i+1]
		routes := routes[from][to]
		min := math.MaxInt
		for _, r := range routes {
			len := convert("A"+r, level+1, maxLevel)
			if len < min {
				min = len
			}
		}
		sum += min
	}
	cache[k] = sum
	fmt.Println("key:", k, "value:", sum)
	return sum
}

func rank(code string) (rank int) {
	if len(code) < 2 {
		return 0
	}
	for i := range code[:len(code)-1] {
		if code[i] == code[i+1] {
			rank++
		}
	}
	return
}
