package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
)

//go:embed input15.txt
var input string

func main() {
	wh := newWarehouse(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(wh))
	fmt.Println("Day07 solution2:", solvePuzzle2(wh))
}

func solvePuzzle1(wh warehouse) int {
	fmt.Println(wh)
	wh.execInstructions()
	fmt.Println()
	fmt.Println(wh)
	return wh.boxSum()
}

func solvePuzzle2(wh warehouse) int {
	return 0
}

const (
	emptyByte byte = '.'
	wallByte  byte = '#'
	boxByte   byte = 'O'
	robotByte byte = '@'

	upByte    byte = '^'
	rightByte byte = '>'
	downByte  byte = 'v'
	leftByte  byte = '<'
)

type point = image.Point

var (
	upPoint    point = point{-1, 0}
	rightPoint point = point{0, 1}
	downPoint  point = point{1, 0}
	leftPoint  point = point{0, -1}

	directions   []point        = []point{upPoint, rightPoint, downPoint, leftPoint}
	directionMap map[byte]point = map[byte]point{
		upByte:    upPoint,
		rightByte: rightPoint,
		downByte:  downPoint,
		leftByte:  leftPoint,
	}
)

type warehouse struct {
	representation [][]byte
	robot          point
	instructions   []byte
}

func newWarehouse(input string) warehouse {
	parts := bytes.Split([]byte(input), []byte("\n\n"))
	var wh warehouse
	for x, line := range bytes.Split(parts[0], []byte("\n")) {
		if len(line) > 0 {
			for y, b := range line {
				if b == robotByte {
					wh.robot = point{x, y}
				}
			}
			wh.representation = append(wh.representation, line)
		}
	}

	wh.instructions = bytes.ReplaceAll(parts[1], []byte("\n"), []byte{})
	return wh
}

func (wh warehouse) String() string {
	str := ""
	for _, line := range wh.representation {
		str += string(line) + "\n"
	}
	return str
}

func (wh *warehouse) execInstructions() {
	for _, i := range wh.instructions {
		wh.execOn(i, wh.robot)
	}
}

func (wh *warehouse) execOn(i byte, p point) bool {
	d, ok := directionMap[i]
	if !ok {
		panic(fmt.Sprintf("Invalid instruction: %v", rune(i)))
	}
	next := p.Add(d)
	fmt.Println(string(i), p, next, string(*wh.get(next)))
	switch *wh.get(next) {
	case wallByte:
		return false
	case boxByte, robotByte:
		if wh.execOn(i, next) {
			wh.move(p, next)
			return true
		}
		return false
	case emptyByte:
		wh.move(p, next)
		return true
	default:
		panic(fmt.Sprintf("Invalid point in warehouse representation[%v]: %v", next, string(*wh.get(next))))
	}
}

func (wh *warehouse) get(p point) *byte {
	return &wh.representation[p.X][p.Y]
}

func (wh *warehouse) move(from, to point) {
	*wh.get(to) = *wh.get(from)
	*wh.get(from) = emptyByte
	if from == wh.robot {
		wh.robot = to
	}
}

func (wh *warehouse) boxSum() int {
	sum := 0
	for x := range wh.representation {
		for y := range wh.representation[x] {
			if *wh.get(point{x, y}) == boxByte {
				sum += 100*x + y
			}
		}
	}
	return sum
}
