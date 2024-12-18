package main

import (
	_ "embed"
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"
)

//go:embed input18.txt
var input string

func main() {
	m := new(input, 70, 1024)
	fmt.Println("Day07 solution1:", m.findPath())
	fmt.Println("Day07 solution2:", solvePuzzle2(m))
}

func solvePuzzle2(m maze) int {
	return 0
}

func parseInput(input string) (parsed [][]int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			report := parseLine(line)
			parsed = append(parsed, report)
		}
	}
	return
}

func parseLine(line string) (nums []int) {
	fields := strings.Fields(line)
	for _, field := range fields {
		val, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		nums = append(nums, val)
	}
	return
}

type point = image.Point

type node struct {
	dist int
	prev []point
}

type maze struct {
	barriers           map[point]bool
	start, end         point
	visited, reachable map[point]node
}

func new(input string, size, count int) (m maze) {
	m.barriers = make(map[image.Point]bool)
	m.visited = make(map[image.Point]node)
	m.reachable = make(map[image.Point]node)
	for i, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		if i >= count {
			break
		}
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("Bad number")
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("Bad number")
		}
		m.barriers[point{x, y}] = true
	}
	m.start = point{0, 0}
	m.end = point{size, size}
	return
}

func (m maze) String() string {
	size := m.end.X + 1
	buff := make([]byte, size*(size+1))
	l := 0
	for _ = range size {
		for j := range size + 1 {
			buff[l+j] = '.'
		}
		buff[l+size] = '\n'
		l += size + 1
	}
	for p := range m.barriers {
		buff[p.Y*(size+1)+p.X] = '#'
	}
	return string(buff)
}

func (m *maze) findPath() int {
	m.reachable[m.start] = node{dist: 0, prev: []point{}}
	for len(m.reachable) > 0 {
		mp := m.closest()
		// move to visited
		m.visited[mp] = m.reachable[mp]
		delete(m.reachable, mp)
		// check neigbours of last node
		m.checkNeigbours(mp)
	}
	return m.visited[m.end].dist
}

func (m maze) closest() point {
	mp, min := point{}, math.MaxInt
	for p, n := range m.reachable {
		if n.dist < min {
			mp, min = p, n.dist
		}
	}
	return mp
}

func (m maze) checkNeigbours(p point) {
	directions := []point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // up,left,down,right
	size := m.end.X
	for _, d := range directions {
		next := p.Add(d)
		// Not valid point
		if next.X < 0 || size < next.X || next.Y < 0 || size < next.Y {
			continue
		}
		// It is a barrier
		if _, ok := m.barriers[next]; ok {
			continue
		}
		// It is visted
		if _, ok := m.visited[next]; ok {
			continue
		}
		// Is next node already in reachables?
		nextDist := m.visited[p].dist + 1
		if nextNode, ok := m.reachable[next]; !ok || nextDist < nextNode.dist {
			m.reachable[next] = node{nextDist, []point{p}}
		} else if nextDist == nextNode.dist {
			nextNode.prev = append(nextNode.prev, p)
			m.reachable[next] = nextNode
		}
	}

}
