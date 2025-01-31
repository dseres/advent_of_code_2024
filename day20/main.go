package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"math"
	"slices"
	"strings"
	"time"
)

//go:embed input20.txt
var input string

func main() {
	start := time.Now()
	m := new(input)
	m.findPath()
	fmt.Println("Day20 solution1:", m.countShortcuts(100, 2))
	fmt.Println("Day20 solution2:", m.countShortcuts(100, 20))
	fmt.Println("Day20 time:", time.Now().Sub(start))
}

const (
	empty byte = '.'
	wall  byte = '#'
	start byte = 'S'
	end   byte = 'E'
)

type point = image.Point

var (
	upPoint    point = point{-1, 0}
	rightPoint point = point{0, 1}
	downPoint  point = point{1, 0}
	leftPoint  point = point{0, -1}

	dirPoints = []point{upPoint, rightPoint, downPoint, leftPoint}
)

type node struct {
	dist int
	prev []point
}

type maze struct {
	tiles              [][]byte
	start, end         point
	visited, reachable map[point]node
	route              []point
}

func new(input string) (m maze) {
	bs := []byte(strings.TrimSpace(input))
	for _, line := range bytes.Split(bs, []byte("\n")) {
		m.tiles = append(m.tiles, line)
	}
	m.start, m.end = m.startEndPoint()
	m.visited = make(map[image.Point]node)
	m.reachable = make(map[image.Point]node)
	return
}

func (m maze) startEndPoint() (point, point) {
	sp, ep := point{}, point{}
	for x, l := range m.tiles {
		for y, b := range l {
			if b == start {
				sp = point{x, y}
			}
			if b == end {
				ep = point{x, y}
			}
		}
	}
	return sp, ep
}

func (m maze) String() string {
	tiles := m.cloneTiles()
	if len(m.route) > 0 {
		for _, p := range m.route[1 : len(m.route)-1] {
			tiles[p.X][p.Y] = 'O'
		}
	}
	return string(bytes.Join(tiles, []byte("\n")))
}

func (m maze) cloneTiles() [][]byte {
	tiles := make([][]byte, 0, len(m.tiles))
	for _, l := range m.tiles {
		tiles = append(tiles, bytes.Clone(l))
	}
	return tiles
}

func (m *maze) findPath() int {
	if 0 < len(m.route) {
		return m.visited[m.end].dist
	}
	m.reachable[m.start] = node{dist: 0, prev: []point{}}
	for len(m.reachable) > 0 {
		p := m.closest()
		// move to visited
		m.visited[p] = m.reachable[p]
		delete(m.reachable, p)
		// check neigbours of last node
		m.checkNeigbours(p)
	}
	// build route
	for p := m.end; ; {
		m.route = append(m.route, p)
		if p == m.start {
			break
		}
		p = m.visited[p].prev[0]
	}
	slices.Reverse(m.route)
	// return distance
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
	for _, d := range dirPoints {
		next := p.Add(d)
		if next.X < 0 || len(m.tiles) <= next.X || next.Y < 0 || len(m.tiles[next.X]) <= next.Y {
			continue
		}
		// It is a wall
		if m.tiles[p.X][p.Y] == wall {
			continue
		}
		// It is visited
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

func (m maze) countShortcuts(length, jumps int) int {
	count := 0
	for i, p1 := range m.route {
		for j, p2 := range m.route[i+1:] {
			// You can shortcut route here
			dist := distance(p1, p2)
			if dist <= jumps && j+1-dist >= length {
				// This is a shortcut
				count++
			}
		}
	}
	return count
}

func distance(a, b point) int {
	d := a.Sub(b)
	return abs(d.X) + abs(d.Y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
