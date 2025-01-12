package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
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
	deps   map[string][]*gate
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
			return a & b
		}
	case "OR":
		op = func(a, b int) int {
			return a | b
		}
	case "XOR":
		op = func(a, b int) int {
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

func (g gate) String() string {
	if g.left != nil && g.right != nil && g.op != nil {
		return fmt.Sprintf("{%v (%v) %v %v (%v) -> %v (%v)}", g.left.id, g.left.value, g.opStr, g.right.id, g.right.value, g.id, g.value)
	}
	return fmt.Sprintf("{%v (%v)}", g.id, g.value)
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

func solvePuzzle2(m machine) string {
	m.buildAdders()
	m.createDependantIndex()
	badWires := m.checkAdders()
	slices.Sort(badWires)
	// for _, g := range m.gates {
	// 	g.computed = false
	// }
	// solvePuzzle1(m)
	return strings.Join(badWires, ",")
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

func (a adder) getXid() string {
	return fmt.Sprintf("x%02d", a.id)
}

func (a adder) getYid() string {
	return fmt.Sprintf("y%02d", a.id)
}

// createDependantIndex will show which gates depends on a wire
func (m *machine) createDependantIndex() {
	m.deps = make(map[string][]*gate)
	for _, g := range m.gates {
		if g.left != nil && g.right != nil {
			m.insertIntoIndex(g, g.left.id)
			m.insertIntoIndex(g, g.right.id)
		}
	}
}

func (m *machine) insertIntoIndex(g *gate, id string) {
	_, ok := m.deps[id]
	if !ok {
		m.deps[id] = []*gate{}
	}
	m.deps[id] = append(m.deps[id], g)
}

func (m *machine) checkAdders() (badWires []string) {
	bw := m.checkFirstAdder()
	badWires = append(badWires, bw...)
	if len(bw) == 2 {
		m.change(bw[0], bw[1])
		m.checkFirstAdder()
	}
	for _, a := range m.adders[1:] {
		bw := m.checkAdder(a)
		badWires = append(badWires, bw...)
		if len(bw) == 2 {
			m.change(bw[0], bw[1])
			m.checkAdder(a)
		}
	}
	return
}

func (m *machine) checkFirstAdder() (badWires []string) {
	// Load input fields
	a := m.adders[0]
	a.x = a.getXid()
	a.y = a.getYid()
	a.c = " 0 "
	a.g = a.c
	// Check XOR and AND gates from X, Y
	z1, cout1, _ := m.checkXORANDDeps(a.x)
	z2, cout2, _ := m.checkXORANDDeps(a.y)
	// Compare gates
	if z1 != z2 {
		panic(fmt.Sprintf("Wires %v or %v are bad, they have different dependants: %v, %v\n", a.x, a.y, z1, z2))
	}
	if cout1 != cout2 {
		panic(fmt.Sprintf("Wires %v or %v are bad, they have different dependants: %v, %v\n", a.x, a.y, cout1, cout2))
	}
	// Check there is no 'z' in E and F wires
	if !strings.HasPrefix(z1.id, "z") {
		badWires = append(badWires, z1.id)
	}
	if strings.HasPrefix(cout1.id, "z") {
		badWires = append(badWires, cout1.id)
	}
	a.e = z1.id
	a.f = cout1.id
	a.z = z1.id
	a.cout = cout1.id
	m.adders[a.id] = a
	return
}

func (m *machine) checkAdder(a adder) (badWires []string) {
	// Load input fields
	a.x = a.getXid()
	a.y = a.getYid()
	a.c = m.adders[a.id-1].cout
	var bw []string
	// Check XOR and AND gates from X, Y
	a.e, a.f, bw = m.computeEandFwires(a)
	badWires = append(badWires, bw...)

	// Check XOR and AND gates from E and C
	a.z, a.g, bw = m.computeZandGwires(a)
	badWires = append(badWires, bw...)
	if len(badWires) == 2 {
		m.adders[a.id] = a
		return badWires
	}

	// Check OR gates from F and G
	a.cout, bw = m.computeCOUTwire(a)
	badWires = append(badWires, bw...)
	m.adders[a.id] = a
	return
}

func (m *machine) computeEandFwires(a adder) (e, f string, bads []string) {
	e1, f1, ok := m.checkXORANDDeps(a.x)
	if !ok {
		panic(fmt.Sprintf("wire %v is bad", a.x))
	}
	e2, f2, ok := m.checkXORANDDeps(a.y)
	if !ok {
		panic(fmt.Sprintf("wire %v is bad", a.y))
	}
	// Compare gates
	if e1 != e2 {
		panic(fmt.Sprintf("Wires %v or %v are bad, they have different dependants: %v, %v\n", a.x, a.y, e1, e2))
	}
	if f1 != f2 {
		panic(fmt.Sprintf("Wires %v or %v are bad, they have different dependants: %v, %v\n", a.x, a.y, f1, f2))
	}
	// Check there is no 'z' in E and F wires
	if strings.HasPrefix(e1.id, "z") {
		bads = append(bads, e1.id)
	}
	if strings.HasPrefix(f1.id, "z") {
		bads = append(bads, f1.id)
	}
	return e1.id, f1.id, bads
}

func (m *machine) computeZandGwires(a adder) (z, g string, bads []string) {
	z1, g1, ok := m.checkXORANDDeps(a.e)
	if !ok {
		return "", "", []string{a.e}
	}
	z2, g2, ok := m.checkXORANDDeps(a.c)
	if !ok {
		return "", "", []string{a.c}
	}
	// Compare gates
	if z1 != z2 {
		panic(fmt.Sprintf("Wires %v or %v are bad, they have different dependants: %v, %v\n", a.e, a.c, z1, z2))
	}
	if g1 != g2 {
		panic(fmt.Sprintf("Wires %v or %v are bad, they have different dependants: %v, %v\n", a.e, a.c, g1, g2))
	}
	if !strings.HasPrefix(z1.id, "z") {
		bads = append(bads, z1.id)
	}
	if strings.HasPrefix(g1.id, "z") {
		bads = append(bads, g1.id)
	}
	return z1.id, g1.id, bads
}

func (m *machine) computeCOUTwire(a adder) (cout string, badWires []string) {
	cout1, ok := m.checkORDeps(a.f)
	if !ok {
		return "", []string{a.f}
	}
	cout2, ok := m.checkORDeps(a.g)
	if !ok {
		return "", []string{a.g}
	}
	// Compare gates
	if cout1 != cout2 {
		panic(fmt.Sprintf("Wires %v or %v are bad, they have different dependants: %v, %v\n", a.f, a.g, cout1, cout2))
	}
	// The carry bit of last gate will be the last z wire
	if strings.HasPrefix(cout1.id, "z") && a.id < m.bits-1 {
		badWires = append(badWires, cout1.id)
	}
	a.cout = cout1.id
	m.adders[a.id] = a
	return cout1.id, badWires
}

func (m *machine) checkXORANDDeps(w string) (xor, and *gate, ok bool) {
	deps, ok := m.deps[w]
	if !ok {
		return nil, nil, false
	}
	if len(deps) == 1 {
		return nil, nil, false
	}
	if len(deps) != 2 {
		return nil, nil, false
	}
	if deps[0].opStr != "XOR" && deps[1].opStr != "XOR" {
		return nil, nil, false
	}
	if deps[0].opStr != "AND" && deps[1].opStr != "AND" {
		return nil, nil, false
	}
	if deps[0].opStr == "XOR" {
		return deps[0], deps[1], true
	}
	return deps[1], deps[0], true
}

func (m *machine) checkORDeps(w string) (or *gate, ok bool) {
	deps, ok := m.deps[w]
	if !ok {
		return nil, false
	}
	if len(deps) != 1 {
		return nil, false
	}
	if deps[0].opStr != "OR" {
		return nil, false
	}
	return deps[0], true
}

func (m *machine) change(w1, w2 string) {
	g1, g2 := m.gates[w1], m.gates[w2]
	g1.id, g2.id = g2.id, g1.id
	m.gates[g1.id], m.gates[g2.id] = g1, g2
	for _, g := range m.deps[w1] {
		if g.left == g1 {
			g.left = g2
		}
		if g.right == g1 {
			g.right = g2
		}
	}
	for _, g := range m.deps[w2] {
		if g.left == g2 {
			g.left = g1
		}
		if g.right == g2 {
			g.right = g1
		}
	}
	g1.computed, g2.computed = false, false
	m.createDependantIndex()
}
