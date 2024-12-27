package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
)

//go:embed input21.txt
var input string

var numPad []byte = []byte("789456123X0A")
var dirPad []byte = []byte("X^A<v>")

func main() {
	codes := parseInput(input)
	fmt.Println(codes)
	fmt.Println("Day07 solution1:", solvePuzzle1(codes))
	fmt.Println("Day07 solution2:", solvePuzzle2(codes))
}

func solvePuzzle1(codes []string) int {
	return 0
}

func solvePuzzle2(codes []string) int {
	return 0
}

func parseInput(input string) (parsed []string) {
	return strings.Split(input, "\n")
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

// func routesOnNumPad() [][]byte {
// 	for i, line := range numPad {
// 		for j, v := range line {

// 		}
// 	}
// }

type routes struct {
	from, to  byte
	positions [][]int
}

type node struct {
	to            byte
	index, length int
	prevs         []int
}

func getRoutes(pad []byte) map[byte]map[byte]routes {
	rs := make(map[byte]map[byte]routes)
	for i, from := range pad {
		if from == 'X' {
			continue
		}
		rs[from] = searchRoutesFrom(pad, i)
	}
	return rs
}

func searchRoutesFrom(pad []byte, from int) map[byte]routes {
	visited := make(map[byte]routes)
	visited[pad[from]] = routes{from: pad[from], to: pad[from], positions: [][]int{[]int{from}}}
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
			v, _ := visited[pad[n.index-prev]]
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
	for _, rs := range visited {
		for _, r := range rs.positions {
			for i, d := range r[1:] {
				switch d {
				case -3:
					r[i] = '^'
				case 3:
					r[i] = 'v'
				case -1:
					r[i] = '<'
				case 1:
					r[i] = '>'
				}
			}
			r[len(r)-1] = 'A'
		}
	}
	return visited
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

func printRoutes(all map[byte]map[byte]routes) {
	for from, rf := range all {
		fmt.Printf("From %v:\n", string(rune(from)))
		for to, rs := range rf {
			fmt.Printf("\tto %v: [", string(rune(to)))
			for i, r := range rs.positions {
				for _, d := range r {
					fmt.Print(string(rune(d)))
				}
				if i < len(rs.positions)-1 {
					fmt.Print(", ")
				}
			}
			fmt.Printf("]\n")
		}
	}

}
