package main

import (
	_ "embed"
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input18.txt
var input string

func main() {
	start := time.Now()
	m := new(input, 70, 1024)
	fmt.Println("Day18 solution1:", m.findPath())
	fmt.Println("Day18 solution2:", m.addMoreBarriers())
	fmt.Println("Day18 time:", time.Now().Sub(start))
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
	allBarriers        []point
	route              map[point]bool
}

func new(input string, size, count int) (m maze) {
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		m.allBarriers = append(m.allBarriers, parseLine(line))
	}
	m.barriers = make(map[image.Point]bool)
	m.visited = make(map[image.Point]node)
	m.reachable = make(map[image.Point]node)
	for _, b := range m.allBarriers[:count] {
		m.barriers[b] = true
	}
	m.start = point{0, 0}
	m.end = point{size, size}
	return
}

func parseLine(line string) point {
	parts := strings.Split(line, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic("Bad number")
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("Bad number")
	}
	return point{x, y}
}

func (m maze) String() string {
	size := m.end.X + 1
	buff := make([]byte, size*(size+1))
	l := 0
	for range size {
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

func (m maze) addMoreBarriers() string {
	m.computeRoute()
	for i := len(m.barriers); i < len(m.allBarriers); i++ {
		m.addBarrier(i)
		if !m.onRoute(i) {
			continue
		}
		if !m.hasRoute(m.end) {
			dist := m.reCalculate()
			if dist == 0 {
				bad := m.allBarriers[i]
				return fmt.Sprintf("%v,%v", bad.X, bad.Y)
			}
			m.computeRoute()
		}
	}
	panic("Route cannot be broken!")
}

func (m *maze) computeRoute() {
	m.route = map[image.Point]bool{}
	m.computeRouteImpl(m.end)
}

func (m *maze) computeRouteImpl(p point) {
	m.route[p] = true
	for _, prev := range m.visited[p].prev {
		m.computeRouteImpl(prev)
	}
}

func (m *maze) onRoute(i int) bool {
	_, ok := m.route[m.allBarriers[i]]
	return ok
}

func (m *maze) addBarrier(i int) {
	barrier := m.allBarriers[i]
	m.barriers[barrier] = true
	delete(m.visited, barrier)
}

func (m *maze) hasRoute(p point) bool {
	if p == m.start {
		return true
	}
	if node, ok := m.visited[p]; ok {
		for i, next := range node.prev {
			if m.hasRoute(next) {
				return true
			}
			// previous nodes are bad, we can remove this connection
			node.prev = node.prev[i:]
			m.visited[p] = node
		}
	}
	return false
}

func (m *maze) reCalculate() int {
	m.visited = map[image.Point]node{}
	m.reachable = map[image.Point]node{}
	return m.findPath()
}
