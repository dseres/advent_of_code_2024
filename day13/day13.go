package main

import (
	_ "embed"
	"fmt"
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
	fmt.Println("Day07 solution1:", solvePuzzle1(cms))
	prepareForPuzzle2(cms)
	fmt.Println("Day07 solution2:", solvePuzzle1(cms))
}

func solvePuzzle1(cms []clawMachine) int {
	sum := 0
	for _, cm := range cms {
		if minCost, ok := getMinCost2(cm); ok {
			sum += minCost
		}
	}
	return sum
}

func solvePuzzle2(cms []clawMachine) int {
	sum := 0
	for _, cm := range cms {
		if minCost, ok := getMinCost2(cm); ok {
			sum += minCost
		}
	}
	return sum
}

type clawMachine struct {
	ax, ay, bx, by, px, py int
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
	return clawMachine{ax, ay, bx, by, px, py}
}

func getMinCost(cm clawMachine) (int, bool) {
	validMoves := []int{}
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			if cm.ax*i+cm.bx*j == cm.px && cm.ay*i+cm.by*j == cm.py {
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

func prepareForPuzzle2(cms []clawMachine) {
	for i := range cms {
		cm := &cms[i]
		cm.px += 10000000000000
		cm.py += 10000000000000
	}
}

func getMinCost2(cm clawMachine) (int, bool) {
	j1 := cm.ax*cm.py - cm.ay*cm.px
	j2 := cm.ax*cm.by - cm.ay*cm.bx
	fmt.Println(cm, j1, j2, j1%j2)
	if j1%j2 == 0 {
		j := j1 / j2
		i1 := (cm.px - j*cm.bx)
		i2 := cm.ax
		fmt.Println(i1, i2, i1%i2)
		if i1%i2 == 0 {
			i := i1 / i2
			fmt.Println(i, j, i*A_TOKEN+j*B_TOKEN)
			return i*A_TOKEN + j*B_TOKEN, true
		}
	}
	return 0, false
}
