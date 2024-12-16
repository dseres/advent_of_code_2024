package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"math"
)

//go:embed input16.txt
var input string

func main() {
	maze := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(maze))
	fmt.Println("Day07 solution2:", solvePuzzle2(maze))
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
	position            point
	direction, distance int
}

func solvePuzzle1(maze [][]byte) int {
	// TODO: write a Dijkstra algorythm
	//  * find reachable points
	n := node{startPoint(maze), toTheRight, 0}
	// fmt.Println(n)
	reachables := map[point]node{}
	visited := map[point]int{n.position: n.distance}
	for {
		addReachableNodes(maze, visited, n, &reachables)
		n.distance = math.MaxInt
		for _, r := range reachables {
			if r.distance < n.distance {
				n = r
			}
		}
		// fmt.Println(n, reachables, visited)
		visited[n.position] = n.distance
		delete(reachables, n.position)
		if maze[n.position.X][n.position.Y] == endByte {
			return n.distance
		}
	}
}

func solvePuzzle2(maze [][]byte) int {
	return 0
}

func startPoint(maze [][]byte) point {
	for x, l := range maze {
		for y, b := range l {
			if b == startByte {
				return point{x, y}
			}
		}
	}
	panic("Maze doesn't have a start point.")
}

func addReachableNodes(maze [][]byte, visited map[point]int, n node, reachables *map[point]node) {
	addReachable(maze, visited, n, 0, 1, reachables)
	addReachable(maze, visited, n, toTheRight, 1001, reachables)
	addReachable(maze, visited, n, toTheLeft, 1001, reachables)
}

func addReachable(maze [][]byte, visited map[point]int, n node, dir, dist int, reachables *map[point]node) {
	nextDir := (n.direction + dir) % 4
	nextPos := n.position.Add(dirPoints[nextDir])
	value := maze[nextPos.X][nextPos.Y]
	if value != wallByte {
		if _, ok := visited[nextPos]; !ok {
			(*reachables)[nextPos] = node{nextPos, nextDir, n.distance + dist}
		}
	}
}
