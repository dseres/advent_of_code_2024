package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input14.txt
var input string

type robot struct {
	px, py, vx, vy int
}

func main() {
	fmt.Println("Day07 solution1:", solvePuzzle1(parseInput(input)))
	fmt.Println("Day07 solution2:", solvePuzzle2(parseInput(input)))
}

func solvePuzzle1(robots []robot) int {
	rect := getRectangle(robots)
	fmt.Println(rect)
	for i := range robots {
		robots[i].px, robots[i].py = getPositionAfter(robots[i], 100, rect)
	}
	mx := (rect.vx - rect.px + 1) / 2
	my := (rect.vy - rect.py + 1) / 2
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
	fmt.Println(robots)
	fmt.Println(a, b, c, d)
	return a * b * c * d
}

func solvePuzzle2(robots []robot) int {
	rect := getRectangle(robots)
	fmt.Println(rect)
	printRobots(rect, robots)
	return 0
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
	strs := regexp.MustCompile("-?\\d+").FindAllString(line, -1)
	r.px, _ = strconv.Atoi(strs[0])
	r.py, _ = strconv.Atoi(strs[1])
	r.vx, _ = strconv.Atoi(strs[2])
	r.vy, _ = strconv.Atoi(strs[3])
	return
}

func getRectangle(robots []robot) robot {
	minx, miny := math.MaxInt, math.MaxInt
	maxx, maxy := math.MinInt, math.MinInt
	for _, r := range robots {
		minx, miny = min(minx, r.px), min(miny, r.py)
		maxx, maxy = max(maxx, r.px), max(maxy, r.py)
	}
	return robot{minx, miny, maxx, maxy}
}

func min[t ~int](a, b t) t {
	if a < b {
		return a
	}
	return b
}

func max[t ~int](a, b t) t {
	if a > b {
		return a
	}
	return b
}

func getPositionAfter(r robot, round int, rect robot) (x, y int) {
	return getCoordinateAfter(r.px, r.vx, rect.vx-rect.px+1, round),
		getCoordinateAfter(r.py, r.vy, rect.vy-rect.py+1, round)
}

func getCoordinateAfter(from, speed, width, round int) int {
	return (from + round*(speed+width)) % width
}

func printRobots(rect robot, robots []robot) {
	for i := range rect.vx - rect.px + 1 {
		for j := range rect.vy - rect.py + 1 {
			count := 0
			for _, r := range robots {
				if r.px == rect.px+i && r.py == rect.py+j {
					count++
				}
			}
			if count == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%v", count)
			}
		}
		fmt.Println()
	}
}
