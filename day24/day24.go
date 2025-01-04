package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input24.txt
var input string

type machine struct {
	startValues map[string]int
	values      map[string]int
	gates       map[string]gate
}

type gate struct {
	left, right, dest string
	operand           opFun
}

type opFun = func(a, b int) int

func main() {
	m := new(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(m))
	fmt.Println("Day07 solution2:", solvePuzzle2(m))
}

func new(input string) machine {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	m := machine{}
	m.values = parseValues(parts[0])
	m.gates = parseGates(parts[1])
	return m
}

func parseValues(input string) map[string]int {
	values := map[string]int{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		values[parts[0]] = v
	}
	return values
}

func parseGates(input string) map[string]gate {
	gates := map[string]gate{}
	re := regexp.MustCompile(`^(\w+) (\w+) (\w+) -> (\w+)$`)
	for _, line := range strings.Split(input, "\n") {
		sm := re.FindStringSubmatch(line)
		g := gate{left: sm[1], operand: parseOp(sm[2]), right: sm[3], dest: sm[4]}
		gates[sm[4]] = g
	}
	return gates
}

func parseOp(input string) opFun {
	var op opFun
	switch input {
	case "AND":
		op = func(a, b int) int {
			fmt.Printf("%v AND %v -> %v\n", a, b, a&b)
			return a & b
		}
	case "OR":
		op = func(a, b int) int {
			fmt.Printf("%v OR %v -> %v\n", a, b, a|b)
			return a | b
		}
	case "XOR":
		op = func(a, b int) int {
			fmt.Printf("%v XOR %v -> %v\n", a, b, a^b)
			return a ^ b
		}
	}
	return op
}

func solvePuzzle1(m machine) int {
	n := 0
	for k := range m.gates {
		if strings.HasPrefix(k, "z") {
			m.compute(k)
			n = n | m.intVal(k)
		}
	}
	for k, v := range m.values {
		fmt.Printf("%v = %v\n", k, v)
	}
	return n
}

func (m *machine) compute(wire string) {
	g, ok := m.gates[wire]
	fmt.Println(g)
	if !ok {
		panic(fmt.Sprintf("there is no operand for %v", wire))
	}
	if _, ok := m.values[g.left]; !ok {
		m.compute(g.left)
	}
	if _, ok := m.values[g.right]; !ok {
		m.compute(g.right)
	}
	value := g.operand(m.values[g.left], m.values[g.right])
	m.values[wire] = value
}

func (m machine) intVal(wire string) int {
	if !strings.HasPrefix(wire, "z") {
		panic(fmt.Sprintf("intVal function can be called with z wire instead of %v", wire))
	}
	bit, err := strconv.Atoi(wire[1:])
	if err != nil {
		panic(err)
	}
	return m.values[wire] << bit
}

func solvePuzzle2(m machine) int {
	return 0
}
