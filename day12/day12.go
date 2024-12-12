package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"
)

//go:embed input12.txt
var input string

var directions []image.Point = []image.Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type plot struct {
	Plant            rune
	X, Y, Id, Fences int
}

func main() {
	plots := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(plots))
	fmt.Println("Day07 solution2:", solvePuzzle2(plots))
}

func solvePuzzle1(plots [][]plot) int {
	return 0
}

func solvePuzzle2(plots [][]plot) int {
	return 0
}

func parseInput(input string) (plots [][]plot) {
	lines := strings.Split(input, "\n")
	for x, line := range lines {
		if len(line) > 0 {
			var plotLine []plot
			for y, r := range line {
				plotLine = append(plotLine, plot{r, x, y, -1, 0})
			}
			plots = append(plots, plotLine)
		}
	}
	return
}

func printPlots(plots [][]plot) {
	printPlants(plots)
	fmt.Println()
	printIds(plots)
	fmt.Println()
	printFences(plots)
}

func printPlants(plots [][]plot) {
	for _, l := range plots {
		for _, p := range l {
			fmt.Print(string(p.Plant))
		}
		fmt.Println()
	}
}

func printIds(plots [][]plot) {
	for _, l := range plots {
		for _, p := range l {
			fmt.Printf("%3d ", p.Id)
		}
		fmt.Println()
	}
}

func printFences(plots [][]plot) {
	for _, l := range plots {
		for _, p := range l {
			fmt.Print(p.Fences)
		}
		fmt.Println()
	}
}

func computeRegions(plots [][]plot) (regions [][]*plot) {
	for x, line := range plots {
		for y, p := range line {
			if p.Id == -1 {
				// New region
				plots[x][y].Id = len(regions)
				region := []*plot{&plots[x][y]}
				region = append(region, checkNeigbours(plots, plots[x][y])...)
				regions = append(regions, region)
			}
		}
	}
	return
}

func checkNeigbours(plots [][]plot, p plot) (region []*plot) {
	for _, n := range p.getNeighbours(plots) {
		if n.Id == -1 && n.Plant == p.Plant {
			n.Id = p.Id
			region = append(region, &plots[n.X][n.Y])
			region = append(region, checkNeigbours(plots, plots[n.X][n.Y])...)
		}
	}
	return
}

func (p *plot) getNeighbours(plots [][]plot) (neighbours []*plot) {
	for _, d := range directions {
		if n := p.getNeighbour(plots, d); n != nil {
			neighbours = append(neighbours, n)
		}
	}
	return
}

func (p *plot) getNeighbour(plots [][]plot, d image.Point) *plot {
	x := p.X + d.X
	y := p.Y + d.Y
	if valid(plots, x, y) {
		return &plots[x][y]
	}
	return nil
}

func valid(plots [][]plot, x, y int) bool {
	return 0 <= x && x < len(plots) && 0 <= y && y < len(plots[x])
}
