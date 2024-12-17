package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input17.txt
var input string

func main() {
	m := new(input)
	fmt.Println("Day07 solution1:", solvePuzzle1(m))
	fmt.Println("Day07 solution2:", solvePuzzle2(m))
}

func solvePuzzle1(m machine) string {
	m.run()
	output := make([]string, 0, len(m.output))
	for _, v := range m.output {
		output = append(output, strconv.Itoa(v))
	}
	return strings.Join(output, ",")
}

func solvePuzzle2(m machine) int {
	return 0
}

type machine struct {
	a, b, c, ip          int
	instructions, output []int
}

func new(input string) machine {
	parts := strings.Split(input, "\n\n")
	regs := regexp.MustCompile(`\d+`).FindAllString(parts[0], -1)
	a, _ := strconv.Atoi(regs[0])
	b, _ := strconv.Atoi(regs[1])
	c, _ := strconv.Atoi(regs[2])
	parts = strings.Fields(parts[1])
	instructions := []int{}
	for _, s := range strings.Split(parts[1], ",") {
		i, _ := strconv.Atoi(s)
		instructions = append(instructions, i)
	}
	return machine{a: a, b: b, c: c, instructions: instructions, ip: 0, output: []int{}}
}

func (m *machine) combo(i int) int {
	switch i {
	case 0, 1, 2, 3:
		return i
	case 4:
		return m.a
	case 5:
		return m.b
	case 6:
		return m.c
	default:
		panic(fmt.Sprintf("Invalid combo operand %v!", i))
	}
}

type computer interface {
	compute(m *machine, op int, combo int)
}

type adv struct{}

func (i adv) compute(m *machine, op int, combo int) {
	m.a = m.a / (1 << combo)
	m.ip += 2
}

type bxl struct{}

func (i bxl) compute(m *machine, op int, combo int) {
	m.b = m.b ^ op
	m.ip += 2
}

type bst struct{}

func (i bst) compute(m *machine, op int, combo int) {
	m.b = combo % 8
	m.ip += 2
}

type jnz struct{}

func (i jnz) compute(m *machine, op int, combo int) {
	if m.a == 0 {
		m.ip += 2
		return
	}
	m.ip = op
}

type bxc struct{}

func (i bxc) compute(m *machine, op int, combo int) {
	m.b = m.c ^ m.b
	m.ip += 2
}

type out struct{}

func (i out) compute(m *machine, op int, combo int) {
	v := combo % 8
	// fmt.Println("out: ", v)
	m.output = append(m.output, v)
	m.ip += 2
}

type bdv struct{}

func (i bdv) compute(m *machine, op int, combo int) {
	m.b = m.a / (1 << combo)
	m.ip += 2
}

type cdv struct{}

func (i cdv) compute(m *machine, op int, combo int) {
	m.c = m.a / (1 << combo)
	m.ip += 2
}

func (m *machine) run() {
	operations := []computer{adv{}, bxl{}, bst{}, jnz{}, bxc{}, out{}, bdv{}, cdv{}}
	for m.ip < len(m.instructions) {
		instruction := m.instructions[m.ip]
		operation := operations[instruction]
		operand := m.instructions[m.ip+1]
		// fmt.Printf("IP: %v, operation: %T %v, operand: %v, combo: %v\n", m.ip, operation, operation, operand, m.combo(operand))
		operation.compute(m, operand, m.combo(operand))
		// fmt.Println(m)
	}
}
