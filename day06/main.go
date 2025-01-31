package main

import (
	_ "bufio"
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "os"
	"slices"
	"time"
)

//go:embed input06.txt
var input string

var directions []image.Point = []image.Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func main() {
	start := time.Now()
	fmt.Println("Day06 solution1:", solvePuzzle1(parseInput(input)))
	fmt.Println("Day06 solution2:", solvePuzzle2(parseInput(input)))
	fmt.Println("Day06 time:", time.Now().Sub(start))
}

func solvePuzzle1(lab [][]byte, pos image.Point) int {
	_, counter := walk(0, pos, lab)
	return counter
}

func solvePuzzle2(lab [][]byte, pos image.Point) int {
	dir := 0
	obstacles := map[image.Point]bool{}
	start := pos
	for {
		valid := true
		dir, valid = computeNextDirection(dir, pos, lab)
		if !valid {
			return len(obstacles)
		}
		nextPos := pos.Add(directions[dir])
		// Checking loop on a cloned map starting with initial position and direction
		newLab := clone(lab)
		newLab[nextPos.X][nextPos.Y] = '#'
		loop, _ := walk(0, start, newLab)
		if loop {
			obstacles[nextPos] = true
		}
		pos = nextPos
	}
}

func walk(dir int, pos image.Point, lab [][]byte) (bool, int) {
	counter := 0
	for {
		if lab[pos.X][pos.Y] == '0'+byte(dir) {
			// This is a loop
			return true, counter
		}
		if lab[pos.X][pos.Y] == '.' {
			counter++
		}
		lab[pos.X][pos.Y] = '0' + byte(dir)
		valid := true
		dir, valid = computeNextDirection(dir, pos, lab)
		if !valid {
			return false, counter
		}
		pos = pos.Add(directions[dir])
	}
}

func computeNextDirection(dir int, pos image.Point, lab [][]byte) (new_dir int, valid bool) {
	for i := 0; i < 4; i++ {
		nextPos := pos.Add(directions[dir])
		if !inRange(lab, nextPos) {
			return dir, false
		}
		if lab[nextPos.X][nextPos.Y] != '#' {
			return dir, true
		}
		dir = (dir + 1) % 4
	}
	panic("Position is invalid. It is between four obstacles.")
}

func inRange(m [][]byte, pos image.Point) bool {
	return 0 <= pos.X && pos.X < len(m) && 0 <= pos.Y && pos.Y < len(m[pos.X])
}

func printMap(lab [][]byte) {
	for _, line := range lab {
		fmt.Println(string(line))
	}
}

func clone(lab [][]byte) [][]byte {
	newLab := [][]byte{}
	for _, line := range lab {
		newLab = append(newLab, slices.Clone(line))
	}
	return newLab
}

func parseInput(input string) (lab [][]byte, startPos image.Point) {
	lab = bytes.Split([]byte(input), []byte{'\n'})
	startPos = findStartPos(lab)
	lab[startPos.X][startPos.Y] = '.'
	return lab, startPos
}

func findStartPos(lab [][]byte) (position image.Point) {
	for position.X = range lab {
		for position.Y = range lab[position.X] {
			if lab[position.X][position.Y] == '^' {
				return
			}
		}
	}
	panic("Cannot find start position.")
}
