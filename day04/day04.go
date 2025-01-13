package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"time"
)

//go:embed input04.txt
var input []byte

func main() {
	start := time.Now()
	text := parseInput(input)
	fmt.Println("Day04 solution1:", solvePuzzle1(text))
	fmt.Println("Day04 solution2:", solvePuzzle2(text))
	fmt.Println("Day04 time:", time.Now().Sub(start))
}

func solvePuzzle1(text [][]byte) (counter int) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			counter += checkDirection(&text, dx, dy)
		}
	}
	return
}

func checkDirection(text *[][]byte, dx int, dy int) (counter int) {
	for i := range *text {
		for j := range (*text)[i] {
			if xmasInText(text, i, j, dx, dy) {
				counter++
			}
		}
	}
	return
}

func xmasInText(text *[][]byte, i, j, dx, dy int) bool {
	xmas := []byte("XMAS")
	for l := 0; l < len(xmas); l++ {
		m, n := i+l*dx, j+l*dy
		if m < 0 || len(*text) <= m || n < 0 || len((*text)[i]) <= n {
			return false
		}
		if (*text)[m][n] != xmas[l] {
			return false
		}
	}
	return true
}

func solvePuzzle2(text [][]byte) (counter int) {
	for i := 1; i < len(text)-1; i++ {
		for j := 1; j < len(text[i])-1; j++ {
			if isCrossMAS(&text, i, j) {
				counter++
			}
		}
	}
	return
}

func isCrossMAS(text *[][]byte, i int, j int) bool {
	return (*text)[i][j] == 'A' &&
		((*text)[i-1][j-1] == 'M' && (*text)[i+1][j+1] == 'S' ||
			(*text)[i-1][j-1] == 'S' && (*text)[i+1][j+1] == 'M') &&
		((*text)[i+1][j-1] == 'M' && (*text)[i-1][j+1] == 'S' ||
			(*text)[i+1][j-1] == 'S' && (*text)[i-1][j+1] == 'M')
}

func parseInput(input []byte) [][]byte {
	text := bytes.Split(input, []byte("\n"))
	// Remove last empty line if exists
	if len(text) > 0 && len(text[len(text)-1]) == 0 {
		text = text[:len(text)-1]
	}
	return text
}
