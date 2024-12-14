package main

import (
	_ "embed"
	"fmt"
	"image"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
	_ "time"
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
	w, h := getDimension(robots)
	changePostion(w, h, robots, 100)
	return getSafetyFactor(w, h, robots)
}

func solvePuzzle2(robots []robot) int {
	start := slices.Clone(robots)
	w, h := getDimension(robots)
	for i := 0; 0 <= i; i++ {
		step(w, h, robots)
		count := countRobotsHaving2Neighbours(robots)
		if 150 < count {
			// time.Sleep(100 * time.Millisecond)
			// fmt.Printf("\nStep %v, count:%v\n", i+1, count)
			// printRobots(w, h, robots)
			return i + 1
		}
		if slices.Equal(robots, start) {
			// fmt.Printf("Start position repeats in %v steps.\n", i)
			break
		}
	}
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

func printRobots(w, h int, robots []robot) {
	for i := range w {
		for j := range h {
			count := 0
			for _, r := range robots {
				if r.px == i && r.py == j {
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

func step(w, h int, robots []robot) {
	for i := range robots {
		r := &robots[i]
		r.px = ((r.px+r.vx)%w + w) % w
		r.py = ((r.py+r.vy)%h + h) % h
	}
}

func isSimetric(w, h int, robots []robot) bool {
	var c1, c2 int
	for _, r := range robots {
		if r.px < w/2 {
			c1++
		}
		if r.px > w/w {
			c2++
		}
	}
	return c1 == c2
}

func countRobotsHaving2Neighbours(robots []robot) int {
	positions := map[image.Point]bool{}
	for _, r := range robots {
		positions[image.Point{r.px, r.py}] = true
	}
	count := 0
	for p := range positions {
		neighbours := 0
		for p2 := range positions {
			d := p2.Sub(p)
			if p != p2 && -1 <= d.X && d.X <= 1 && -1 <= d.Y && d.Y <= 1 {
				neighbours++
			}
		}
		if neighbours >= 2 {
			count++
		}
	}
	return count
}
