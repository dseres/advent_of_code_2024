package main

import (
	_ "embed"
	"fmt"
	"image"
	"slices"
	"strings"
	"time"
)

//go:embed input12.txt
var input string

func main() {
	start := time.Now()
	garden := newGarden(input)
	// fmt.Println(garden)
	fmt.Println("Day12 solution1:", solvePuzzle1(garden))
	fmt.Println("Day12 solution2:", solvePuzzle2(garden))
	fmt.Println("Day12 time:", time.Now().Sub(start))
}

func solvePuzzle1(g garden) int {
	sum := 0
	for i := range g.regions {
		sum += g.getArea(i) * int(g.getPerimeter(i))
	}
	return sum
}

func solvePuzzle2(g garden) int {
	sum := 0
	for i := range g.regions {
		sum += g.getArea(i) * g.computeSidesFor(i)
	}
	return sum
}

var directions []image.Point = []image.Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type plot struct {
	plant  rune
	point  image.Point
	id     int
	fences int
}

type garden struct {
	plots   [][]plot
	regions [][]*plot
}

func newGarden(input string) (g garden) {
	g = parseInput(input)
	g.computeRegions()
	g.computeFences()
	return
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
			s += fmt.Sprintf("%4d ", p.id)
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
		s += "] "
		s += fmt.Sprintf("Area: %v, Fences: %v, Sides: %v\n", g.getArea(i), g.getPerimeter(i), g.computeSidesFor(i))
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
	// sort regions
	for _, region := range g.regions {
		slices.SortFunc(region, comparePlotPtr)
	}
}

func comparePlotPtr(a, b *plot) int {
	return slices.Compare([]int{a.point.X, a.point.Y}, []int{b.point.X, b.point.Y})
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

func (g *garden) getNeighbour(p *plot, dir image.Point) *plot {
	next := g.getNext(p, dir)
	if next == nil || next.id != p.id {
		return nil
	}
	return next
}

func (g *garden) computeFences() {
	for _, line := range g.plots {
		for y := range line {
			g.computeFencesFor(&line[y])
		}
	}
}

func (g *garden) computeFencesFor(p *plot) {
	for _, d := range directions {
		if g.getNeighbour(p, d) == nil {
			p.fences++
		}
	}
}

func (g *garden) computeSidesFor(i int) (count int) {
	region := g.regions[i]
	// Sides can be computed from slices
	rows, columns := g.getRowsAndColumns(region)
	count += g.countHorizontalSides(rows)
	count += g.countVerticalSides(columns)
	return
}

func (g *garden) getRowsAndColumns(region []*plot) (rows map[int][]*plot, columns map[int][]*plot) {
	rows, columns = map[int][]*plot{}, map[int][]*plot{}
	for _, p := range region {
		x, y := p.point.X, p.point.Y
		if _, ok := rows[x]; !ok {
			rows[x] = []*plot{}
		}
		rows[x] = append(rows[x], p)
		if _, ok := columns[y]; !ok {
			columns[y] = []*plot{}
		}
		columns[y] = append(columns[y], p)
	}
	return
}

func (g *garden) countHorizontalSides(rows map[int][]*plot) (count int) {
	for _, row := range rows {
		// create lists of indices, we have to count the distinct intervals in it
		xs1, xs2 := []int{}, []int{}
		for _, p := range row {
			if g.getNeighbour(p, directions[0]) == nil {
				xs1 = append(xs1, p.point.Y)
			}
			if g.getNeighbour(p, directions[2]) == nil {
				xs2 = append(xs2, p.point.Y)
			}
		}
		count += getDistinctIntervals(xs1)
		count += getDistinctIntervals(xs2)
	}
	return
}

func (g *garden) countVerticalSides(columns map[int][]*plot) (count int) {
	for _, column := range columns {
		ys1, ys2 := []int{}, []int{}
		for _, p := range column {
			if g.getNeighbour(p, directions[1]) == nil {
				ys1 = append(ys1, p.point.X)
			}
			if g.getNeighbour(p, directions[3]) == nil {
				ys2 = append(ys2, p.point.X)
			}
		}
		count += getDistinctIntervals(ys1)
		count += getDistinctIntervals(ys2)
	}
	return
}

func getDistinctIntervals(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 < nums[i] {
			count++
		}
	}
	return count
}

func (g *garden) getArea(i int) int {
	return len(g.regions[i])
}

func (g *garden) getPerimeter(i int) (perimeter int) {
	for _, p := range g.regions[i] {
		perimeter += p.fences
	}
	return
}
