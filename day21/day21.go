package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input21.txt
var input string

var numPad []byte = []byte("789456123X0A")
var keyPad []byte = []byte("X^A<v>")

func main() {
	codes := parseInput(input)
	fmt.Println(codes)
	fmt.Println("Day07 solution1:", solvePuzzle1(codes))
	fmt.Println("Day07 solution2:", solvePuzzle2(codes))
}

func solvePuzzle1(codes []string) int {
	return 0
}

func solvePuzzle2(codes []string) int {
	return 0
}

func parseInput(input string) (parsed []string) {
	return strings.Split(input, "\n")
}

func toNums(input []byte, level int) []byte {
	var pad []byte = keyPad
	if level == 0 {
		pad = numPad
	}
	pos := slices.Index(pad, 'A')
	output := []byte{}
	for _, d := range input {
		switch d {
		case '^':
			pos -= 3
		case '<':
			pos--
		case '>':
			pos++
		case 'v':
			pos += 3
		case 'A':
			if level == 0 {
				output = append(output, pad[pos])
			} else {
				output = append(output, pad[pos])
			}
		default:
			panic(fmt.Sprintf("bad instruction %v", d))
		}
	}
	if level > 0 {
		return toNums(output, level-1)
	}
	return output
}

// func routesOnNumPad() [][]byte {
// 	for i, line := range numPad {
// 		for j, v := range line {

// 		}
// 	}
// }
