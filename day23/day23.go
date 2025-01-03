package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input23.txt
var input string

func main() {
	g := parseInput(input)
	fmt.Println(g)
	fmt.Println("Day07 solution1:", solvePuzzle1(g))
	fmt.Println("Day07 solution2:", solvePuzzle2(g))
}

type trio struct{ a, b, c int16 }

func solvePuzzle1(g graph) int {
	p := make([]int16, 0, len(g))
	for k := range g {
		p = append(p, k)
	}
	cliques := bronKerbosh2(g, []int16{}, p, []int16{})
	trios := make(map[trio]struct{})
	for _, clique := range cliques {
		getTrios(trios, clique)
	}
	return len(trios)
}

func solvePuzzle2(g graph) int {
	return 0
}

type graph = map[int16][]int16

func parseInput(input string) (g graph) {
	g = make(graph)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		a, b := parseLine(line)
		insert(g, a, b)
	}
	return
}

func parseLine(line string) (a, b int16) {
	s := strings.Split(line, "-")
	return str2Int16(s[0]), str2Int16(s[1])
}

func str2Int16(s string) int16 {
	return int16(s[0]-'a')*int16('z'-'a'+1) + int16(s[1]-'a')
}

func int162Str(i int16) string {
	return fmt.Sprintf("%c%c", i/('z'-'a'+1)+'a', i%('z'-'a'+1)+'a')
}

func startWithT(n int16) bool {
	return n/('z'-'a'+1) == 't'-'a'
}

func insert(g graph, a, b int16) {
	insertOneDir(g, a, b)
	insertOneDir(g, b, a)
}

func insertOneDir(g graph, from, to int16) {
	if _, ok := g[from]; ok {
		g[from] = append(g[from], to)
		return
	}
	g[from] = []int16{to}
}

func clique2String(clique []int16) string {
	b := strings.Builder{}
	b.WriteString("[")
	b.WriteString(int162Str(clique[0]))
	for _, i := range clique[1:] {
		b.WriteString(", ")
		b.WriteString(int162Str(i))
	}
	b.WriteString("]")
	return b.String()
}

// algorithm BronKerbosch2(R, P, X) is
//     if P and X are both empty then
//         report R as a maximal clique
//     choose a pivot vertex u in P ⋃ X
//     for each vertex v in P \ N(u) do
//         BronKerbosch2(R ⋃ {v}, P ⋂ N(v), X ⋂ N(v))
//         P := P \ {v}
//         X := X ⋃ {v}

func bronKerbosh2(n graph, r, p, x []int16) (found [][]int16) {
	// fmt.Println("BronKerbosh:", r, p, x)
	if len(p) == 0 && len(x) == 0 {
		slices.Sort(r)
		// fmt.Println("found: ", r)
		return [][]int16{r}
	}
	// pivot vertex has maximum neighbours
	pUx := slices.Concat(p, x)
	u := slices.MaxFunc(pUx, func(a, b int16) int {
		return cmp.Compare(len(n[a]), len(n[b]))
	})
	// fmt.Println("u, n[u]:", u, n[u])
	// for each vertex v in P \ N(u) do
	for _, v := range slices.Clone(p) {
		if slices.Contains(n[u], v) {
			continue
		}
		// fmt.Println("v, n[v]:", v, n[v])
		// R ⋃ {v}
		r2 := slices.Clone(r)
		r2 = append(r2, v)
		// P ⋂ N(v)
		p2 := []int16{}
		for _, v2 := range p {
			if slices.Contains(n[v], v2) {
				p2 = append(p2, v2)
			}
		}
		// X ⋂ N(v)
		x2 := []int16{}
		for _, v2 := range x {
			if slices.Contains(n[v], v2) {
				x2 = append(x2, v2)
			}
		}
		f := bronKerbosh2(n, r2, p2, x2)
		found = append(found, f...)
		// move v from p to x
		vi := slices.Index(p, v)
		p = slices.Delete(p, vi, vi+1)
		x = append(x, v)
		// fmt.Println(p, x)
	}
	slices.SortFunc(found, func(a, b []int16) int { return slices.Compare(a, b) })
	return
}

func getTrios(trios map[trio]struct{}, clique []int16) {
	if len(clique) < 3 {
		return
	}
	if len(clique) == 3 {
		t := trio{clique[0], clique[1], clique[2]}
		if hasT(t) {
			trios[t] = struct{}{}
		}
		return
	}
	for i, c1 := range clique {
		for j, c2 := range clique[i+1:] {
			for _, c3 := range clique[i+j+2:] {
				t := trio{c1, c2, c3}
				if hasT(t) {
					trios[t] = struct{}{}
				}
			}
		}
	}
}

func hasT(t trio) bool {
	return startWithT(t.a) || startWithT(t.b) || startWithT(t.c)
}
