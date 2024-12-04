package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"slices"
)

//go:embed input04.txt
var input []byte

func main() {
	text := parseInput(input)
	printText(text)
	_, count := solvePuzzle1(text)
	fmt.Println("Day04 solution1:", count)
	_, count = solvePuzzle2(text)
	fmt.Println("Day04 solution2:", count)
}

func solvePuzzle1(text [][]byte) (representation [][]byte, counter int64) {
	representation = slices.Repeat([][]byte{slices.Repeat([]byte{'.'}, len(text[0]))}, len(text))
	xmas := []byte("XMAS")
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			for i := range text {
				for j := range text[i] {
					found := true
					for l := 0; l < len(xmas); l++ {
						m, n := i+l*dx, j+l*dy
						if m < 0 || len(text) <= m || n < 0 || len(text[i]) <= n {
							found = false
							break
						}
						if text[m][n] != xmas[l] {
							found = false
							break
						}
					}
					if found {
						for l := 0; l < len(xmas); l++ {
							m, n := i+l*dx, j+l*dy
							representation[m][n] = xmas[l]
						}
						counter++
					}
				}
			}
		}
	}
	return
}

func solvePuzzle2(text [][]byte) (representation [][]byte, counter int64) {
	representation = slices.Repeat([][]byte{slices.Repeat([]byte{'.'}, len(text[0]))}, len(text))
	for i := 1; i < len(text)-1; i++ {
		for j := 1; j < len(text[i])-1; j++ {
			if text[i][j] == 'A' &&
				(text[i-1][j-1] == 'M' && text[i+1][j+1] == 'S' ||
					text[i-1][j-1] == 'S' && text[i+1][j+1] == 'M') &&
				(text[i+1][j-1] == 'M' && text[i-1][j+1] == 'S' ||
					text[i+1][j-1] == 'S' && text[i-1][j+1] == 'M') {
				counter++
			}
		}
	}
	return
}

func parseInput(input []byte) [][]byte {
	text := bytes.Split(input, []byte("\n"))
	// Remove last empty line if exists
	if len(text) > 0 && len(text[len(text)-1]) == 0 {
		text = text[:len(text)-1]
	}
	return text
}

func printText(text [][]byte) {
	for _, line := range text {
		fmt.Println(string(line))
	}
}

func toString(text [][]byte) string {
	var s string
	for _, line := range text {
		s += fmt.Sprintln(line)
	}
	return s
}
