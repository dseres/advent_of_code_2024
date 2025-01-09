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

var adderImage = ` Wires of Ripple-carry adder #%v

 ---X:%3s---+---->+-----+
            |     | XOR |------+---E:%3s---->+-----+
 ---Y:%3s-----+-->+-----+      |             | XOR |-------------Z:%3s------>
            | |             |--------------->+-----+
            | |-->+-----+   |  |
            |     | AND |-------------------F:%3s------>+----+
            |---->+-----+   |  |                        | OR |--C out:%3s--->
                            |  |->+-----+          |--->+----+
                            |     | AND |---G:%3s--|
 ---C:%3s-------------------+---->+-----+

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

func (a adder) String() string {
	return fmt.Sprintf(adderImage, a.id, a.x, a.e, a.y, a.z, a.f, a.cout, a.g, a.c)
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
	depIndex := m.dependantIndex()
	m.checkAllXY(depIndex)
	fmt.Println(m.adders)
	return 0
}

func (m *machine) buildAdders() {
	m.adders = make([]adder, 0, m.bits)
	for i := range m.bits {
		a := adder{id: i}
		m.adders = append(m.adders, a)
	}
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

// dependantIndex will show which gates depends on a wire
func (m machine) dependantIndex() (index map[string][]*gate) {
	index = make(map[string][]*gate)
	for _, g := range m.gates {
		if g.left != nil && g.right != nil {
			insertIntoIndex(index, g, g.left.id)
			insertIntoIndex(index, g, g.right.id)
		}
	}
	return
}

func insertIntoIndex(index map[string][]*gate, g *gate, id string) {
	_, ok := index[id]
	if !ok {
		index[id] = []*gate{}
	}
	index[id] = append(index[id], g)
}

func (m *machine) checkAllXY(index map[string][]*gate) {
	for i, a := range m.adders {
		x := a.getXid()
		y := a.getYid()
		xor1, and1 := m.checkXY(index, x)
		xor2, and2 := m.checkXY(index, y)
		if xor1 != xor2 {
			panic(fmt.Sprintf("Adder of %v and %v is bad, the XOR gates of the wires are different.", x, y))
		}
		if and1 != and2 {
			panic(fmt.Sprintf("Adder of %v and %v is bad, the AND gates of the wires are different.", x, y))
		}
		a.x = x
		a.y = y
		a.e = xor1.id
		a.f = and1.id
		if i == 0 {
			a.z = xor1.id
			a.cout = and1.id
			a.c = " 0 "
			a.g = " 0 "
		}
		m.adders[i] = a
	}
}

func (m *machine) checkXY(index map[string][]*gate, id string) (xor, and *gate) {
	deps, ok := index[id]
	if !ok {
		panic(fmt.Sprintf("Wire %v is bad, there is no gate depending on this wire!", id))
	}
	if len(deps) != 2 {
		panic(fmt.Sprintf("Wire %v is bad, there are more than 2 gates are depending on this wire!", id))
	}
	if deps[0].opStr != "XOR" && deps[1].opStr != "XOR" {
		panic(fmt.Sprintf("Wire %v is bad, there is no XOR gate depending on this wire!", id))
	}
	if deps[0].opStr != "AND" && deps[1].opStr != "AND" {
		panic(fmt.Sprintf("Wire %v is bad, there is no XOR gate depending on this wire!", id))
	}
	if deps[0].opStr == "XOR" {
		return deps[0], deps[1]
	}
	return deps[1], deps[0]
}
