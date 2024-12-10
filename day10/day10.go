package main

import (
	_ "embed"
	"fmt"
	"image"
	"maps"
	"strings"
)

//go:embed input10.txt
var input string

type tile struct {
	height     uint8
	visited    bool
	trailheads map[point]bool
	routes     map[string]bool
}

type point = image.Point

var directions []point = []point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func main() {
	tiles := parseInput(input)
	solution1, solution2 := solvePuzzle1(tiles)
	fmt.Println("Day07 solution1:", solution1)
	fmt.Println("Day07 solution2:", solution2)
}

func solvePuzzle1(tiles [][]tile) (int, int) {
	starts := findStartPoints(tiles)
	fmt.Println(starts)
	sum_heads, sum_routes := 0, 0
	for _, sp := range starts {
		search4routes(sp, tiles)
		t := tiles[sp.X][sp.Y]
		sum_heads += len(t.trailheads)
		sum_routes += len(t.routes)
	}
	fmt.Println(tiles)
	return sum_heads, sum_routes
}

func parseInput(input string) (parsed [][]tile) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			report := parseLine(line)
			parsed = append(parsed, report)
		}
	}
	return
}

func parseLine(line string) (tiles []tile) {
	for _, c := range line {
		height := uint8(c - '0')
		tiles = append(tiles, tile{height: height, trailheads: map[point]bool{}, routes: map[string]bool{}})
	}
	return
}

func findStartPoints(tiles [][]tile) (starts []point) {
	fmt.Println(tiles)
	for x, line := range tiles {
		for y, t := range line {
			if t.height == 0 {
				starts = append(starts, point{X: x, Y: y})
			}
		}
	}
	return
}

func search4routes(sp point, tiles [][]tile) {
	st := &tiles[sp.X][sp.Y]
	st.visited = true
	if st.height == 9 {
		st.trailheads[sp] = true
		st.routes[appendRoute(sp, st, "")] = true
		return
	}
	for _, d := range directions {
		np := sp.Add(d)
		if valid(np, tiles) {
			nt := &tiles[np.X][np.Y]
			if nt.height == st.height+1 {
				if !nt.visited {
					search4routes(np, tiles)
				}
				maps.Copy(st.trailheads, nt.trailheads)
				copyRoutes(sp, st, nt)
			}

		}
	}
}

func valid(p point, tiles [][]tile) bool {
	return 0 <= p.X && p.X < len(tiles) && 0 <= p.Y && p.Y < len(tiles[p.X])
}

func appendRoute(p point, t *tile, r string) string {
	actual := fmt.Sprintf("{%v,%v,%v}", p.X, p.Y, string('0'+t.height))
	return actual + r
}

func copyRoutes(p point, t, nt *tile) {
	for r, _ := range nt.routes {
		t.routes[appendRoute(p, t, r)] = true
	}
}
