package main

import (
	_ "embed"
	"fmt"
	"image"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input13.txt
var input string

const (
	A_TOKEN = 3
	B_TOKEN = 1
)

func main() {
	cms := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle(cms, getMinBruteForce))
	increasePrizes(cms)
	fmt.Println("Day07 solution2:", solvePuzzle(cms, getMinUsingEquations))
}

func solvePuzzle(cms []clawMachine, getMin func(clawMachine) (int, bool)) int {
	sum := 0
	for _, cm := range cms {
		if minCost, ok := getMin(cm); ok {
			sum += minCost
		}
	}
	return sum
}

type clawMachine struct {
	a, b, p image.Point
}

func parseInput(input string) (parsed []clawMachine) {
	blocks := strings.Split(input, "\n\n")
	for _, block := range blocks {
		cm := parseClawMachine(block)
		parsed = append(parsed, cm)
	}
	return
}

func parseClawMachine(block string) (cm clawMachine) {
	lines := strings.Split(block, "\n")
	rb := regexp.MustCompile(`^Button [AB]: X\+(\d+), Y\+(\d+)\s*$`)
	rm := regexp.MustCompile(`^Prize: X=(\d+), Y=(\d+)\s*$`)
	m1 := rb.FindAllStringSubmatch(lines[0], -1)
	m2 := rb.FindAllStringSubmatch(lines[1], -1)
	m3 := rm.FindAllStringSubmatch(lines[2], -1)
	ax, _ := strconv.Atoi(m1[0][1])
	ay, _ := strconv.Atoi(m1[0][2])
	bx, _ := strconv.Atoi(m2[0][1])
	by, _ := strconv.Atoi(m2[0][2])
	px, _ := strconv.Atoi(m3[0][1])
	py, _ := strconv.Atoi(m3[0][2])
	return clawMachine{image.Point{ax, ay}, image.Point{bx, by}, image.Point{px, py}}
}

func getMinBruteForce(c clawMachine) (int, bool) {
	validMoves := []int{}
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			if c.a.X*i+c.b.X*j == c.p.X && c.a.Y*i+c.b.Y*j == c.p.Y {
				cost := i*A_TOKEN + j*B_TOKEN
				validMoves = append(validMoves, cost)
			}
		}
	}
	if len(validMoves) == 0 {
		return 0, false
	}
	return slices.Min(validMoves), true
}

func increasePrizes(cms []clawMachine) {
	for i := range cms {
		cm := &cms[i]
		cm.p.X += 10000000000000
		cm.p.Y += 10000000000000
	}
}

// Mathematical equations can be formed to get i,j as a rational number
// if i,j will integer, those values can be used.
func getMinUsingEquations(c clawMachine) (int, bool) {
	j1 := c.a.X*c.p.Y - c.a.Y*c.p.X
	j2 := c.a.X*c.b.Y - c.a.Y*c.b.X
	if j1%j2 == 0 {
		j := j1 / j2
		i1 := (c.p.X - j*c.b.X)
		i2 := c.a.X
		if i1%i2 == 0 {
			i := i1 / i2
			return i*A_TOKEN + j*B_TOKEN, true
		}
	}
	return 0, false
}
