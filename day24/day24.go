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
	bits   int
	values map[string]int
	gates  map[string]*gate
	adders []adder
}

type gate struct {
	id          string
	value       int
	op          opFun
	opStr       string
	left, right *gate
	computed    bool
}

type opFun = func(a, b int) int

var adderImage = ` Wires of Ripple-carry adder

 ---X---+---->+-----+
        |     | XOR |--E---+------->+-----+
 ---Y-----+-->+-----+      |        | XOR |--------------Z------>
        | |             |---------->+-----+
        | |-->+-----+   |  |
        |     | AND |-------------------F------>+----+
        |---->+-----+   |  |                    | OR |--C out -->
                        |  |->+-----+      |--->+----+
                        |     | AND |---G--|
 ---C-------------------+---->+-----+
`

type adder struct {
	id      int
	x, y, c string //input
	z, cout string // output
	e, f, g string // internal
}

func main() {
	m := new(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(m))
	fmt.Println("Day07 solution2:", solvePuzzle2(m))
}

func new(input string) machine {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	m := machine{}
	m.values, m.bits = parseValues(parts[0])
	m.gates = parseGates(parts[1])
	m.loadInput()
	return m
}

func parseValues(input string) (values map[string]int, bits int) {
	values = map[string]int{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		values[parts[0]] = v
		if strings.HasPrefix(parts[0], "x") {
			bits++
		}
	}
	return
}

func parseGates(input string) map[string]*gate {
	gates := map[string]*gate{}
	re := regexp.MustCompile(`^(\w+) (\w+) (\w+) -> (\w+)$`)
	for _, line := range strings.Split(input, "\n") {
		sm := re.FindStringSubmatch(line)
		leftId, op, rightId, id := sm[1], sm[2], sm[3], sm[4]
		if _, ok := gates[id]; !ok {
			gates[id] = &gate{}
		}
		g := gates[id]
		g.id = id
		g.op = parseOp(op)
		g.opStr = op
		if l, ok := gates[leftId]; ok {
			g.left = l
		} else {
			g.left = &gate{id: leftId}
			gates[leftId] = g.left
		}
		if l, ok := gates[rightId]; ok {
			g.right = l
		} else {
			g.right = &gate{id: rightId}
			gates[rightId] = g.right
		}
	}
	return gates
}

func parseOp(input string) opFun {
	var op opFun
	switch input {
	case "AND":
		op = func(a, b int) int {
			// fmt.Printf("%v AND %v -> %v\n", a, b, a&b)
			return a & b
		}
	case "OR":
		op = func(a, b int) int {
			// fmt.Printf("%v OR %v -> %v\n", a, b, a|b)
			return a | b
		}
	case "XOR":
		op = func(a, b int) int {
			// fmt.Printf("%v XOR %v -> %v\n", a, b, a^b)
			return a ^ b
		}
	}
	return op
}

func (m *machine) loadInput() {
	for k, v := range m.values {
		g := m.gates[k]
		g.value = v
		g.computed = true
	}
}

func (m machine) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("X: %*b\n", m.bits+1, m.getX()))
	b.WriteString(fmt.Sprintf("Y: %*b\n", m.bits+1, m.getY()))
	b.WriteString(fmt.Sprintf("Z: %-*b\n", m.bits+1, m.getZ()))
	b.WriteString("Rules:")
	for _, g := range m.gates {
		if g.op == nil {
			b.WriteString(fmt.Sprintf("\t%v (%v) %v\n", g.id, g.value, g))
			continue
		}
		b.WriteString(fmt.Sprintf("\t%v (%v) %v %v (%v) -> %v (%v) %v\n", g.left.id, g.left.value, g.opStr, g.right.id, g.right.value, g.id, g.value, g))
	}
	return b.String()
}

func solvePuzzle1(m machine) int {
	n := 0
	for k := range m.gates {
		if strings.HasPrefix(k, "z") {
			g := m.gates[k]
			m.compute(g)
			n = n | (g.value << g.bitNo())
		}
	}
	// fmt.Println(m)
	return n
}

func (m *machine) compute(g *gate) {
	// fmt.Printf("Computing %v from %v and %v\n", g, g.left, g.right)
	if !g.computed && g.left != nil && g.right != nil && g.op != nil {
		m.compute(g.left)
		m.compute(g.right)
		g.value = g.op(g.left.value, g.right.value)
	}
	g.computed = true
}

func (g gate) bitNo() int {
	if !strings.HasPrefix(g.id, "z") {
		panic(fmt.Sprintf("intVal function can be called with z wire instead of %v", g.id))
	}
	bit, err := strconv.Atoi(g.id[1:])
	if err != nil {
		panic(err)
	}
	return bit
}

func solvePuzzle2(m machine) int {
	fmt.Println("bits: ", m.bits)
	m.buildAdders()
	return 0
}

func (m *machine) buildAdders() {
	// x00, y00 -> z00
}

func (m machine) getX() int {
	return m.getValue("x")
}

func (m machine) getY() int {
	return m.getValue("x")
}

func (m machine) getZ() int {
	return m.getValue("z")
}

func (m machine) getValue(prefix string) int {
	v := 0
	for i := range len(m.gates) {
		id := fmt.Sprintf("%v%02d", prefix, i)
		if g, ok := m.gates[id]; ok {
			v = v | (g.value << i)
		}
	}
	return v
}

// func (m machine) clone() (other machine) {
// 	other.bits = m.bits
// 	other.gates = maps.Clone(m.gates)
// 	other.values = maps.Clone(m.values)
// 	return
// }

func (a adder) getXid() string {
	return fmt.Sprintf("x%02d", a.id)
}
func (a adder) getYid() string {
	return fmt.Sprintf("y%02d", a.id)
}
func (a adder) getZid() string {
	return fmt.Sprintf("z%02d", a.id)
}
