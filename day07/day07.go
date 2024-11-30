package main

import "fmt"
import "os"

func main() {
	dat, err := os.ReadFile("day07/input07.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat[:])
	fmt.Println("Day07 solution1:", solvePuzzle1(input))
	fmt.Println("Day07 solution2:", solvePuzzle2(input))
}

func solvePuzzle1(input string) int64 {
	return 0
}

func solvePuzzle2(input string) int64 {
	return 0
}
