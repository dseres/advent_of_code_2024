package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn
`

var interconnected = `aq,cg,yn
aq,vc,wq
co,de,ka
co,de,ta
co,ka,ta
de,ka,ta
kh,qp,ub
qp,td,wh
tb,vc,wq
tc,td,wh
td,wh,yn
ub,vc,wq`

func TestGraph(t *testing.T) {
	a := assert.New(t)
	a.Equal(int16(16), str2Int16("aq"))
	a.Equal(int16(58), str2Int16("cg"))
	g := parseInput(testInput)
	a.Equal(4, len(g[16]))
	for _, n := range g[16] {
		a.True(slices.Contains(g[n], 16))
	}

	a.True(startWithT(str2Int16("tc")))
	a.True(startWithT(str2Int16("td")))
}

func TestBronKerbosh2(t *testing.T) {
	g := graph{}
	insert(g, 1, 2)
	insert(g, 1, 5)
	insert(g, 2, 3)
	insert(g, 2, 5)
	insert(g, 3, 4)
	insert(g, 4, 5)
	insert(g, 4, 6)
	cliques := bronKerbosh2(g, []int16{}, []int16{1, 2, 3, 4, 5, 6}, []int16{})
	a := assert.New(t)
	req := [][]int16{{1, 2, 5}, {2, 3}, {3, 4}, {4, 5}, {4, 6}}
	a.Equal(req, cliques)
}

func TestSolution1(t *testing.T) {
	g := parseInput(testInput)
	fmt.Println('z' - 'a' + 1)
	fmt.Println(g)
	result := solvePuzzle1(g)
	a := assert.New(t)
	a.Equal(7, result)
}

func TestSolution2(t *testing.T) {
	result := solvePuzzle2(parseInput(testInput))
	a := assert.New(t)
	a.Equal(0, result)
}
