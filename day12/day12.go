package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"
)

//go:embed input12.txt
var input string

func main() {
	plots := parseInput(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(plots))
	fmt.Println("Day07 solution2:", solvePuzzle2(plots))
}

func solvePuzzle1(g garden) int {
	return 0
}

func solvePuzzle2(g garden) int {
	return 0
}

var directions []image.Point = []image.Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type plot struct {
	plant  rune
	point  image.Point
	id     int
	fences uint
}

type garden struct {
	plots   [][]plot
	regions [][]*plot
}

func parseInput(input string) (g garden) {
	lines := strings.Split(input, "\n")
	for x, line := range lines {
		if len(line) > 0 {
			var plotLine []plot
			for y, r := range line {
				plotLine = append(plotLine, plot{r, image.Point{x, y}, -1, 0})
			}
			g.plots = append(g.plots, plotLine)
		}
	}
	return
}

// Implement Stringer
func (g garden) String() string {
	return g.printPlants() + "\n" + g.printIds() + "\n" + g.printFences() + "\n" + g.printRegions()
}

func (g *garden) printPlants() string {
	s := ""
	for _, l := range g.plots {
		for _, p := range l {
			s += fmt.Sprint(string(p.plant))
		}
		s += "\n"
	}
	return s
}

func (g *garden) printIds() string {
	s := ""
	for _, l := range g.plots {
		for _, p := range l {
			s += fmt.Sprintf("%3d ", p.id)
		}
		s += "\n"
	}
	return s
}

func (g *garden) printFences() string {
	s := ""
	for _, l := range g.plots {
		for _, p := range l {
			s += fmt.Sprint(p.fences)
		}
		s += "\n"
	}
	return s
}

func (g *garden) printRegions() string {
	s := ""
	for i, r := range g.regions {
		s += fmt.Sprintf("Region %v: [", i)
		for j, p := range r {
			s += fmt.Sprintf("{%v, %v}", p.point.X, p.point.Y)
			if j < len(r)-1 {
				s += ", "
			}
		}
		s += "]\n"
	}
	return s
}

func (g *garden) get(x, y int) *plot {
	return &g.plots[x][y]
}

func (g *garden) computeRegions() {
	for x, line := range g.plots {
		for y := range line {
			p := g.get(x, y)
			if p.id == -1 {
				// New region
				p.id = len(g.regions)
				g.regions = append(g.regions, []*plot{p})
				g.findPlotsInRegion(p)
			}
		}
	}
	return
}

func (g *garden) findPlotsInRegion(p *plot) {
	for _, dir := range directions {
		if next := g.getNext(p, dir); next != nil && next.id == -1 && next.plant == p.plant {
			// Add n to region
			next.id = p.id
			g.regions[p.id] = append(g.regions[p.id], next)
			g.findPlotsInRegion(next)
		}
	}
	return
}

func (g *garden) getNext(p *plot, dir image.Point) *plot {
	nextPos := p.point.Add(dir)
	if g.valid(nextPos) {
		return g.get(nextPos.X, nextPos.Y)
	}
	return nil
}

func (g *garden) valid(p image.Point) bool {
	return 0 <= p.X && p.X < len(g.plots) && 0 <= p.Y && p.Y < len(g.plots[p.X])
}
