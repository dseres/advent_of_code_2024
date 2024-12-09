package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input09.txt
var input string

type group struct {
	start, len int
}

const SPACE = int16(-1)

func main() {
	blocks := toBlocks(parseInput(input))
	fmt.Println("Day07 solution1:", solvePuzzle1(blocks))
	blocks = toBlocks(parseInput(input))
	fmt.Println("Day07 solution2:", solvePuzzle2(blocks))
}

func solvePuzzle1(blocks []int16) int {
	blocks = defragment(blocks)
	return computeSum(blocks)
}

func solvePuzzle2(blocks []int16) int {
	defragment2(blocks)
	return computeSum(blocks)
}

func parseInput(input string) []uint8 {
	if input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}
	nums := make([]uint8, 0, len(input))
	for i := range input {
		if n, err := strconv.ParseUint(input[i:i+1], 10, 8); err == nil {
			nums = append(nums, uint8(n))
		} else {
			panic(err)
		}
	}
	return nums
}

func toBlocks(nums []uint8) []int16 {
	blocks := make([]int16, 0, len(nums)*10)
	empty := false
	id := int16(0)
	for _, n := range nums {
		for i := uint8(0); i < n; i++ {
			if empty {
				blocks = append(blocks, -1)
			} else {
				blocks = append(blocks, id)
			}
		}
		if !empty {
			id++
		}
		empty = !empty
	}
	return blocks
}

func defragment(blocks []int16) []int16 {
	l := len(blocks) - 1
	for i, id := range blocks {
		for blocks[l] == SPACE {
			l--
		}
		if i >= l {
			return blocks[0 : i+1]
		}
		if id == SPACE {
			blocks[i], blocks[l] = blocks[l], SPACE
		}
	}
	panic("Unreachable code")
}

func computeSum(blocks []int16) int {
	sum := 0
	for i, id := range blocks {
		if id != -1 {
			sum += i * int(id)
		}
	}
	return sum
}

func defragment2(blocks []int16) {
	files, spaces := getGroups(blocks)
	for j := len(files) - 1; j >= 0; j-- {
		file := files[j]
		for i := 0; i < len(spaces) && spaces[i].start < file.start; i++ {
			space := spaces[i]
			if space.len >= file.len {
				// Move file
				for k := 0; k < file.len; k++ {
					blocks[space.start+k] = blocks[file.start+k]
					blocks[file.start+k] = SPACE
				}
				spaces[i] = group{space.start + file.len, space.len - file.len}
				if i == 0 && spaces[i].len == 0 {
					spaces = spaces[1:]
				}
				break
			}
		}
	}
}

func getGroups(blocks []int16) ([]group, []group) {
	files := make([]group, 0, 10000)
	spaces := make([]group, 0, 10000)
	inFile, start, prev_id := true, 0, SPACE
	for i, id := range blocks {
		if inFile && id != prev_id {
			// New file or space
			files = append(files, group{start, i - start})
			start = i
			if id == SPACE {
				inFile = false
			}
		} else if !inFile && id != SPACE {
			// New file
			spaces = append(spaces, group{start, i - start})
			start = i
			inFile = !inFile
		}
		prev_id = id
	}
	if inFile {
		// New file
		files = append(files, group{start, len(blocks) - start})
	} else {
		// New space
		spaces = append(spaces, group{start, len(blocks) - start})
	}
	return files, spaces
}
