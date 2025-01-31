package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
	_ "time"
)

//go:embed input14.txt
var input string

type robot struct {
	px, py, vx, vy int
}

func main() {
	start := time.Now()
	robots := parseInput(input)
	width, height := getDimension(robots)
	robots2 := slices.Clone(robots)
	fmt.Println("Day14 solution1:", solvePuzzle1(width, height, robots))
	fmt.Println("Day14 solution2:", solvePuzzle2(width, height, robots2))
	fmt.Println("Day14 time:", time.Now().Sub(start))
}

func solvePuzzle1(w, h int, robots []robot) int {
	changePostion(w, h, robots, 100)
	return getSafetyFactor(w, h, robots)
}

func solvePuzzle2(w, h int, robots []robot) int {
	start := slices.Clone(robots)
	min, ind := math.MaxInt, -1
	for i := 1; true; i++ {
		step(w, h, robots)
		dist := getEntropy(w, h, robots)
		if dist < min {
			min, ind = dist, i
		}
		if slices.Equal(robots, start) {
			break
		}
	}
	return ind
}

func parseInput(input string) (robots []robot) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			robots = append(robots, parseLine(line))
		}
	}
	return
}

func parseLine(line string) (r robot) {
	strs := regexp.MustCompile(`-?\d+`).FindAllString(line, -1)
	r.px, _ = strconv.Atoi(strs[0])
	r.py, _ = strconv.Atoi(strs[1])
	r.vx, _ = strconv.Atoi(strs[2])
	r.vy, _ = strconv.Atoi(strs[3])
	return
}

func getDimension(robots []robot) (int, int) {
	w, h := math.MinInt, math.MinInt
	for _, r := range robots {
		w, h = max(w, r.px), max(h, r.py)
	}
	return w + 1, h + 1
}

func max[t ~int](a, b t) t {
	if a > b {
		return a
	}
	return b
}

func changePostion(w, h int, robots []robot, steps int) {
	for i := range robots {
		robots[i].px, robots[i].py = getPositionAfter(robots[i], steps, w, h)
	}
}

func getPositionAfter(r robot, round, width, height int) (x, y int) {
	return getCoordinateAfter(r.px, r.vx, width, round),
		getCoordinateAfter(r.py, r.vy, height, round)
}

func getCoordinateAfter(from, speed, width, round int) int {
	return (from + round*(speed+width)) % width
}

func getSafetyFactor(width, height int, robots []robot) int {
	mx := width / 2
	my := height / 2
	a, b, c, d := 0, 0, 0, 0
	for _, r := range robots {
		if r.px < mx && r.py < my {
			a++
		}
		if r.px > mx && r.py < my {
			b++
		}
		if r.px < mx && r.py > my {
			c++
		}
		if r.px > mx && r.py > my {
			d++
		}
	}
	return a * b * c * d
}

func step(w, h int, robots []robot) {
	for i := range robots {
		r := &robots[i]
		r.px = ((r.px+r.vx)%w + w) % w
		r.py = ((r.py+r.vy)%h + h) % h
	}
}

func getEntropy(w, h int, robots []robot) int {
	dist := 0
	for _, r := range robots {
		dist += abs(r.px-w/2) + abs(r.py-h/2)
	}
	return dist
}

func abs[T ~int](t T) T {
	if t < 0 {
		return -t
	}
	return t
}
