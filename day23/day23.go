package main

import (
	_ "embed"
	"fmt"
	"maps"
	"math"
	"slices"
	"strings"
)

//go:embed input23.txt
var input string

type graph = map[int16][]int16
type trio struct{ a, b, c int16 }
type set = map[int16]struct{}

func main() {
	g := parseInput(input)
	cliques := getSortedCliques(g)
	fmt.Println("Day07 solution1:", solvePuzzle1(cliques))
	fmt.Println("Day07 solution2:", solvePuzzle2(cliques))
}

func solvePuzzle1(cliques [][]int16) int {
	trios := make(map[trio]struct{})
	for _, clique := range cliques {
		getTrios(trios, clique)
	}
	return len(trios)
}

func solvePuzzle2(cliques [][]int16) string {
	mc, maxLen := []int16{}, math.MinInt
	for _, c := range cliques {
		if len(c) > maxLen {
			mc = c
			maxLen = len(c)
		}
	}
	return clique2String(mc)
}

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
	b.WriteString(int162Str(clique[0]))
	for _, i := range clique[1:] {
		b.WriteString(",")
		b.WriteString(int162Str(i))
	}
	return b.String()
}

// This function is implementation of the Bron-Kerbosh graph algorithm with pivot point
//
// algorithm BronKerbosch2(R, P, X) is
//
//	if P and X are both empty then
//	    report R as a maximal clique
//	choose a pivot vertex u in P ⋃ X
//	for each vertex v in P \ N(u) do
//	    BronKerbosch2(R ⋃ {v}, P ⋂ N(v), X ⋂ N(v))
//	    P := P \ {v}
//	    X := X ⋃ {v}
func bronKerbosh2(n graph, r, p, x set, found *[][]int16) {
	// fmt.Println("BronKerbosh:", r, p, x)
	if len(p) == 0 && len(x) == 0 {
		clique := slices.Collect(maps.Keys(r))
		slices.Sort(clique)
		*found = append(*found, clique)
	}
	u := pivot(n, p, x)
	// for each vertex v in P \ N(u) do
	for v := range minus(p, n[u]) {
		// R ⋃ {v}
		r2 := maps.Clone(r)
		r2[v] = struct{}{}
		// P ⋂ N(v)
		p2 := intersect(p, n[v])
		// X ⋂ N(v)
		x2 := intersect(x, n[v])
		bronKerbosh2(n, r2, p2, x2, found)
		// move v from p to x
		delete(p, v)
		x[v] = struct{}{}
	}
}

func union(a, b set) set {
	u := maps.Clone(a)
	maps.Insert(u, maps.All(b))
	return u
}

func pivot(n graph, p, x set) int16 {
	pUx := union(p, x)
	u, maxN := int16(0), math.MinInt16
	for k := range pUx {
		if len(n[k]) > maxN {
			maxN = len(n[k])
			u = k
		}
	}
	return u
}

func minus(a set, b []int16) set {
	m := maps.Clone(a)
	for _, k := range b {
		delete(m, k)
	}
	return m
}

func intersect(a set, b []int16) set {
	m := make(set)
	for _, k := range b {
		if _, ok := a[k]; ok {
			m[k] = struct{}{}
		}
	}
	return m
}

func getSortedCliques(g graph) [][]int16 {
	p := make(set)
	for k := range g {
		p[k] = struct{}{}
	}
	cliques := [][]int16{}
	bronKerbosh2(g, make(set), p, make(set), &cliques)
	slices.SortFunc(cliques, func(a, b []int16) int { return slices.Compare(a, b) })
	return cliques
}

func getTrios(trios map[trio]struct{}, clique []int16) {
	if len(clique) < 3 {
		return
	}
	for i, a := range clique {
		for j, b := range clique[i+1:] {
			for _, c := range clique[i+j+2:] {
				if startWithT(a) || startWithT(b) || startWithT(c) {
					t := trio{a, b, c}
					trios[t] = struct{}{}
				}
			}
		}
	}
}
