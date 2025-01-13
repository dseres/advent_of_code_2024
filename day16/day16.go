package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"math"
	_ "slices"
	"time"
)

//go:embed input16.txt
var input string

func main() {
	start := time.Now()
	maze := parseInput(input)
	dist, count := solvePuzzle(maze)
	fmt.Println("Day16 solution1:", dist)
	fmt.Println("Day16 solution2:", count)
	fmt.Println("Day16 time:", time.Now().Sub(start))
}

func parseInput(input string) [][]byte {
	return bytes.Split([]byte(input), []byte("\n"))
}

const (
	emptyByte byte = '.'
	wallByte  byte = '#'
	startByte byte = 'S'
	endByte   byte = 'E'

	toTheRight int = 1
	toTheLeft  int = 3
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
	position  point
	direction int
}

type value struct {
	distance int
	prevs    []node
}

func solvePuzzle(maze [][]byte) (int, int) {
	// It is a dijkstra algorythm, x,y,direction represents a node
	sp, ep := startEndPoint(maze)
	sk := node{sp, toTheRight}
	sv := value{0, []node{}}
	reachables := map[node]value{}
	visited := map[node]value{sk: sv}
	addReachableNodes(maze, visited, reachables, sk)
	dijkstra(maze, visited, reachables)
	min, endNodes := minDist(visited, ep)
	count := countRoutes(visited, map[point]bool{}, endNodes)
	return min, count
}

func startEndPoint(maze [][]byte) (point, point) {
	sp, ep := point{}, point{}
	for x, l := range maze {
		for y, b := range l {
			if b == startByte {
				sp = point{x, y}
			}
			if b == endByte {
				ep = point{x, y}
			}
		}
	}
	return sp, ep
}

func dijkstra(maze [][]byte, visited map[node]value, reachables map[node]value) {
	for len(reachables) > 0 {
		dist := math.MaxInt
		n := node{}
		for k, v := range reachables {
			if v.distance < dist {
				dist = v.distance
				n = k
			}
		}
		visited[n] = reachables[n]
		delete(reachables, n)
		addReachableNodes(maze, visited, reachables, n)
	}
}

func addReachableNodes(maze [][]byte, visited map[node]value, reachables map[node]value, from node) {
	addReachable(maze, visited, reachables, from, node{from.position.Add(dirPoints[from.direction]), from.direction}, 1)
	addReachable(maze, visited, reachables, from, node{from.position, (from.direction + toTheLeft) % 4}, 1000)
	addReachable(maze, visited, reachables, from, node{from.position, (from.direction + toTheRight) % 4}, 1000)
}

func addReachable(maze [][]byte, visited map[node]value, reachables map[node]value, from, to node, dist int) {
	if b := maze[to.position.X][to.position.Y]; b != wallByte {
		dist = visited[from].distance + dist
		if _, ok := visited[to]; ok {
			return
		}
		if r, ok := reachables[to]; ok {
			if dist < r.distance {
				reachables[to] = value{dist, []node{from}}
			}
			if dist == r.distance {
				r.prevs = append(r.prevs, from)
				reachables[to] = r
			}
			return
		}
		reachables[to] = value{dist, []node{from}}
	}
}

func minDist(visited map[node]value, ep point) (int, []node) {
	min, nodes := math.MaxInt, []node{}
	for i := 0; i < 4; i++ {
		v, ok := visited[node{ep, i}]
		if ok && v.distance < min {
			min = v.distance
			nodes = append(nodes, node{ep, i})
		}
	}
	return min, nodes
}

func countRoutes(visited map[node]value, path map[point]bool, from []node) int {
	// fmt.Println(from)
	for _, k := range from {
		path[k.position] = true
		countRoutes(visited, path, visited[k].prevs)

	}
	return len(path)
}
