package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
)

//go:embed input08.txt
var input string

func main() {
	f := newField(parseInput(input))
	fmt.Println("Day07 solution1:", solvePuzzle1(f))
	fmt.Println("Day07 solution2:", solvePuzzle2(f))
}

func solvePuzzle1(f field) int {
	return len(f.getAntinodes(f.computeAntinodes))
}

func solvePuzzle2(f field) int {
	return len(f.getAntinodes(f.computeAllAntinodes))
}

func parseInput(input string) (parsed [][]byte) {
	lines := bytes.Split([]byte(input), []byte{'\n'})
	for _, line := range lines {
		if len(line) > 0 {
			parsed = append(parsed, line)
		}
	}
	return
}

type point = image.Point
type field struct {
	input    [][]byte
	antennas map[byte][]point
}

func newField(input [][]byte) (f field) {
	f.input = input
	f.antennas = make(map[byte][]point, 0)
	for i, line := range input {
		for j, n := range line {
			if n != '.' {
				if _, ok := f.antennas[n]; !ok {
					f.antennas[n] = make([]point, 0)
				}
				f.antennas[n] = append(f.antennas[n], point{i, j})
			}
		}
	}
	return
}

func (f *field) getAntinodes(compute func([]point) []point) map[point]bool {
	antinodes := map[point]bool{}
	for _, positions := range f.antennas {
		for _, apos := range compute(positions) {
			antinodes[apos] = true
		}
	}
	return antinodes
}

func (f *field) computeAntinodes(positions []point) []point {
	antinodes := []point{}
	for i, p := range positions {
		for _, q := range positions[i+1:] {
			diff := p.Sub(q)
			s, r := p.Add(diff), q.Sub(diff)
			if f.valid(s) {
				antinodes = append(antinodes, s)
			}
			if f.valid(r) {
				antinodes = append(antinodes, r)
			}
		}
	}
	return antinodes
}

func (f *field) computeAllAntinodes(positions []point) []point {
	antinodes := []point{}
	for i, p := range positions {
		for _, q := range positions[i+1:] {
			diff := p.Sub(q)
			for s := p; f.valid(s); s = s.Add(diff) {
				antinodes = append(antinodes, s)
			}
			for s := p; f.valid(s); s = s.Sub(diff) {
				antinodes = append(antinodes, s)
			}
		}
	}
	return antinodes
}

func (f *field) valid(p point) bool {
	return 0 <= p.X && p.X < len(f.input) && 0 <= p.Y && p.Y < len(f.input[p.X])
}
