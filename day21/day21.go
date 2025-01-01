package main

import (
	"cmp"
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

func main() {
	codes := parseInput(input)
	fmt.Println(codes)
	fmt.Println("Day07 solution1:", solvePuzzle1(codes))
}

func solvePuzzle1(codes []string) int {
	sum := 0
	for _, code := range codes {
		sum += numOf(code) * convert(code, "", 'A', 0, 2)
		// sum += numOf(code) * getMin(code, 2)
	}
	return sum
}

func parseInput(input string) (parsed []string) {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func toNums(input []byte, level int) []byte {
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
		return toNums(output, level-1)
	}
	return output
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
		slices.SortFunc(strs, func(a, b string) int { return cmp.Compare(rank(b), rank(a)) })
		routes[to] = strs
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

func convert(code, prev string, from byte, level, maxLevel int) int {
	// fmt.Println(len(code), string(code), string(rune(from)), level, maxLevel, string(prev))
	if level > maxLevel {
		return len(code)
	}
	if len(code) == 0 {
		return convert(prev, "", 'A', level+1, maxLevel)
	}
	to := byte(code[0])
	routes := dirRoutes
	if level == 0 {
		routes = numRoutes
	}
	minLength := math.MaxInt
	for _, r := range routes[from][to] {
		next := prev + r
		length := convert(code[1:], next, to, level, maxLevel)
		if length < minLength {
			minLength = length
		}
	}
	return minLength
}

// func convertOneLevel(code string, from byte, level int) []string {
// 	if len(code) == 0 {
// 		return []string{}
// 	}
// 	to := code[0]
// 	routes := dirRoutes
// 	if level == 0 {
// 		routes = numRoutes
// 	}
// 	nextCodes := convertOneLevel(code[1:], to, level)
// 	if len(nextCodes) == 0 {
// 		return routes[from][to]
// 	}
// 	codes := []string{}
// 	for _, c := range routes[from][to] {
// 		for _, nc := range nextCodes {
// 			codes = append(codes, c+nc)
// 		}
// 	}
// 	return codes
// }

// func convertWithLevels(code string, maxLevel int) []string {
// 	codes := []string{code}
// 	for level := range maxLevel + 1 {
// 		// maxRank := math.MinInt
// 		// for _, nc := range next {
// 		// 	r := rank(nc)
// 		// 	if r > maxRank {
// 		// 		maxRank = r
// 		// 	}
// 		// }
// 		nextCodes := []string{}
// 		for _, c := range codes {
// 			nextCodes = append(nextCodes, convertOneLevel(c, 'A', level)...)
// 			// fmt.Println("Code: ", c, "length:", len(c), "next codes:", nextCodes)
// 		}
// 		codes = nextCodes
// 	}
// 	return codes
// }

// func getMin(code string, maxLevel int) int {
// 	codes := convertWithLevels(code, maxLevel)
// 	min := math.MaxInt
// 	for _, c := range codes {
// 		if len(c) < min {
// 			min = len(c)
// 		}
// 	}
// 	return min
// }

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
