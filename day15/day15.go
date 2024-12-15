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
	wh = newWarehouse(input)
	fmt.Println("Day07 solution2:", solvePuzzle2(wh))
}

func solvePuzzle1(wh warehouse) int {
	wh.execInstructions()
	return wh.boxSum()
}

func solvePuzzle2(wh warehouse) int {
	wh.resize()
	wh.execInstructions()
	return wh.boxSum()
}

const (
	emptyByte byte = '.'
	wallByte  byte = '#'
	boxByte   byte = 'O'
	robotByte byte = '@'
	boxLeft   byte = '['
	boxRight  byte = ']'

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
		// fmt.Println(wh)
	}
}

func (wh *warehouse) execOn(instruction byte, actualPos point) bool {
	direction, ok := directionMap[instruction]
	if !ok {
		panic(fmt.Sprintf("Invalid instruction: %v", rune(instruction)))
	}
	nextPos := actualPos.Add(direction)
	nextValue := *wh.get(nextPos)
	switch nextValue {
	case wallByte:
		return false
	case emptyByte:
		wh.move(actualPos, nextPos)
		return true
	case boxByte, robotByte:
		if wh.execOn(instruction, nextPos) {
			wh.move(actualPos, nextPos)
			return true
		}
		return false
	case boxLeft, boxRight:
		return wh.execOnSpecialBox(instruction, direction, actualPos, nextPos)
	default:
		panic(fmt.Sprintf("Invalid point in warehouse representation[%v]: %v", nextPos, string(*wh.get(nextPos))))
	}
}

func (wh *warehouse) execOnSpecialBox(instruction byte, direction, actualPos, nextPos point) bool {
	if direction == leftPoint || direction == rightPoint {
		if wh.execOn(instruction, nextPos) {
			wh.move(actualPos, nextPos)
			return true
		}
		return false
	}
	// Moving up or down.
	if wh.canMoveUpOrDown(direction, actualPos) {
		otherPos := wh.getOtherPosOfBox(nextPos)
		wh.execOn(instruction, nextPos)
		wh.execOn(instruction, otherPos)
		wh.move(actualPos, nextPos)
		return true
	}
	return false
}

func (wh *warehouse) canMoveUpOrDown(direction point, p point) bool {
	if direction != upPoint && direction != downPoint {
		panic(fmt.Sprintf("Invalid direction: %v.", direction))
	}
	nextPos := p.Add(direction)
	nextValue := *wh.get(nextPos)
	switch nextValue {
	case wallByte:
		return false
	case emptyByte:
		return true
	case boxLeft, boxRight:
		// If moving up or down, we need the other part of the box
		otherPos := wh.getOtherPosOfBox(nextPos)
		return wh.canMoveUpOrDown(direction, nextPos) && wh.canMoveUpOrDown(direction, otherPos)
	default:
		panic(fmt.Sprintf("Invalid point in warehouse representation[%v]: %v", nextPos, string(*wh.get(nextPos))))
	}
}

func (wh *warehouse) getOtherPosOfBox(pos point) point {
	value := *wh.get(pos)
	if value == boxLeft {
		return pos.Add(rightPoint)
	}
	if value == boxRight {
		return pos.Add(leftPoint)
	}
	panic(fmt.Sprintf("Invalid position parameter for getOtherPartOfBox: %v", pos))
}

func (wh *warehouse) get(p point) *byte {
	return &wh.representation[p.X][p.Y]
}

func (wh *warehouse) move(from, to point) {
	if *wh.get(from) == emptyByte {
		return
	}
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
			p := *wh.get(point{x, y})
			if p == boxByte || p == boxLeft {
				sum += 100*x + y
			}
		}
	}
	return sum
}

func (wh *warehouse) resize() {
	newRep := make([][]byte, 0, len(wh.representation))
	for _, line := range wh.representation {
		newLine := make([]byte, 0, 2*len(line))
		for _, b := range line {
			switch b {
			case emptyByte, wallByte:
				newLine = append(newLine, b, b)
			case boxByte:
				newLine = append(newLine, boxLeft, boxRight)
			case robotByte:
				newLine = append(newLine, robotByte, emptyByte)
			default:
				panic(fmt.Sprintf("Invalid point in warehouse representation: %v", string(b)))
			}
		}
		newRep = append(newRep, newLine)
	}
	wh.robot.Y = 2 * wh.robot.Y
	wh.representation = newRep
}
