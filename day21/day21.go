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

var numPad = "789456123X0A"
var dirPad = "X^A<v>"
var numRoutes map[byte]map[byte][]string = getRoutes(numPad)
var dirRoutes map[byte]map[byte][]string = getRoutes(dirPad)

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

func toNums(input string, level int) string {
	pad := padOf(level)
	pos := strings.Index(pad, "A")
	output := strings.Builder{}
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
			output.WriteByte(pad[pos])
		default:
			panic(fmt.Sprintf("bad instruction %v", d))
		}
	}
	if level > 0 {
		return toNums(output.String(), level-1)
	}
	return output.String()
}

func padOf(level int) string {
	pad := dirPad
	if level == 0 {
		pad = numPad
	}
	return pad
}

type node struct {
	from, to    byte
	routes      [][]int // Routes are stored as sequence of index offsets. E.g.: [-3 -3 1 1]
	index, dist int
	prevs       []int
}

func getRoutes(pad string) map[byte]map[byte][]string {
	rs := make(map[byte]map[byte][]string)
	for i := range pad {
		if pad[i] == 'X' {
			continue
		}
		rs[pad[i]] = searchRoutesFrom(pad, i)
	}
	return rs
}

func searchRoutesFrom(pad string, startInd int) map[byte][]string {
	visited := make(map[byte]node)
	start := pad[startInd]
	visited[pad[startInd]] = node{from: start, to: start, index: startInd, routes: [][]int{{}}}
	reachables := make(map[byte]node)
	findReachables(pad, visited, reachables, visited[start])
	for len(reachables) > 0 {
		n := getNextReachable(reachables)
		n = computeRoutesOfNode(pad, visited, n)
		visited[n.to] = n
		delete(reachables, n.to)
		findReachables(pad, visited, reachables, n)
	}
	return convertVisiteds(visited)
}

func findReachables(pad string, visited, reachables map[byte]node, n node) {
	dirs := []int{-3, 3, -1, 1}
	for _, dir := range dirs {
		nextInd := n.index + dir
		// Is next position is valid?
		if dir == -1 && n.index%3 == 0 || dir == 1 && n.index%3 == 2 {
			continue
		}
		if nextInd < 0 || len(pad) <= nextInd || pad[nextInd] == 'X' {
			continue
		}
		to := pad[nextInd]
		if _, ok := visited[to]; ok {
			continue
		}
		// Check next node in reachables
		switch r, ok := reachables[to]; {
		case ok && n.dist+1 < r.dist: // A better route found to node
			r.dist = n.dist + 1
			r.prevs = []int{dir}
			reachables[to] = r
		case ok && r.dist == n.dist+1: // A new route found to node
			r.prevs = append(r.prevs, dir)
			reachables[to] = r
		case !ok: // Add new node to reachables
			r = node{from: n.from, to: to, index: nextInd, dist: n.dist + 1, prevs: []int{dir}}
			reachables[to] = r
		}
	}
}

func getNextReachable(reachables map[byte]node) node {
	n, min := node{}, math.MaxInt
	for _, rn := range reachables {
		if rn.dist < min {
			n = rn
			min = rn.dist
		}
	}
	return n
}

func computeRoutesOfNode(pad string, visited map[byte]node, n node) node {
	for _, dir := range n.prevs {
		prevNode := visited[pad[n.index-dir]]
		for _, prevRoute := range prevNode.routes {
			route := slices.Clone(prevRoute)
			route = append(route, dir)
			n.routes = append(n.routes, route)
		}
	}
	return n
}

func convertVisiteds(visited map[byte]node) map[byte][]string {
	// Convert directions to string
	routes := make(map[byte][]string)
	for to, v := range visited {
		strs := routesAsString(v.routes)
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

func routesAsString(directions [][]int) []string {
	strs := []string{}
	for _, route := range directions {
		builder := strings.Builder{}
		for _, dir := range route {
			switch dir {
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
	return strs
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
	for i := range code[:len(code)-1] {
		from := code[i]
		to := code[i+1]
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
