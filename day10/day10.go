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
}

type point = image.Point

var directions []point = []point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func main() {
	tiles := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(tiles))
	fmt.Println("Day07 solution2:", solvePuzzle2(tiles))
}

func solvePuzzle1(tiles [][]tile) int {
	starts := findStartPoints(tiles)
	fmt.Println(starts)
	sum := 0
	for _, sp := range starts {
		sum += len(search4routes(sp, tiles))
	}
	fmt.Println(tiles)
	return sum
}

func solvePuzzle2(tiles [][]tile) int {
	return 0
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
		tiles = append(tiles, tile{height: height, trailheads: map[point]bool{}})
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

func search4routes(sp point, tiles [][]tile) map[point]bool {
	if tiles[sp.X][sp.Y].height == 9 {
		tiles[sp.X][sp.Y].visited = true
		tiles[sp.X][sp.Y].trailheads[sp] = true
		return tiles[sp.X][sp.Y].trailheads
	}
	trailheads := map[point]bool{}
	for _, d := range directions {
		np := sp.Add(d)
		if valid(np, tiles) &&
			tiles[sp.X][sp.Y].height+1 == tiles[np.X][np.Y].height {
			if !tiles[np.X][np.Y].visited {
				maps.Copy(trailheads, search4routes(np, tiles))
			} else {
				maps.Copy(trailheads, tiles[np.X][np.Y].trailheads)
			}
		}
	}
	tiles[sp.X][sp.Y].visited = true
	tiles[sp.X][sp.Y].trailheads = trailheads
	return trailheads
}

func valid(p point, tiles [][]tile) bool {
	return 0 <= p.X && p.X < len(tiles) && 0 <= p.Y && p.Y < len(tiles[p.X])
}
