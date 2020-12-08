package main

import (
	"strconv"
	"strings"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
	ip "github.com/IAmBullsaw/AOC-2020/pkg/inp"
)

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string) (answer int) {
	data := ip.ToSlice(puzzle)
	set := make(map[int]interface{})
	i := 0
	for i < len(data) {
		if set[i] == true {
			break
		}
		set[i] = true
		instruction := data[i]
		val, _ := strconv.Atoi(data[i+1])

		if instruction == "nop" {
			i += 2
		} else if instruction == "jmp" {
			i += (val * 2)
		} else if instruction == "acc" {
			answer += val
			i += 2
		}

	}

	return
}

type Code int

const (
	nop Code = iota
	jmp
	acc
)

func (c Code) String() string {
	return [...]string{"nop", "jmp", "acc"}[c]
}

func ToCode(s string) Code {
	if s[0] == 'n' {
		return nop
	} else if s[0] == 'j' {
		return jmp
	} else if s[0] == 'a' {
		return acc
	}
	panic("aaaa")
}

type Instruction struct {
	Op    Code
	Value int
}

type Handheld struct {
	Pc, Acc      int
	Instructions []Instruction
	Executed     map[int]bool
}

func NewHandheld(instructions []Instruction) Handheld {
	return Handheld{Pc: 0, Acc: 0, Instructions: instructions, Executed: make(map[int]bool)}
}

func (h *Handheld) HasExecuted() bool {
	return h.Executed[h.Pc]

}

func (h *Handheld) Reset() {
	h.Pc = 0
	h.Acc = 0
	h.Executed = make(map[int]bool)
}

func (h *Handheld) Solve() int {
	for h.Pc < len(h.Instructions) {
		if h.HasExecuted() {
			break
		}
		h.Executed[h.Pc] = true
		if h.Instructions[h.Pc].Op == nop {
			h.Pc++
			continue
		}
		if h.Instructions[h.Pc].Op == jmp {
			h.Pc += h.Instructions[h.Pc].Value
			continue
		}
		if h.Instructions[h.Pc].Op == acc {
			h.Acc += h.Instructions[h.Pc].Value
			h.Pc++
		}
	}

	return h.Acc
}

func (h *Handheld) WillSolve() bool {
	h.Reset()
	for h.Pc < len(h.Instructions) {
		if h.HasExecuted() {
			return false
		}
		h.Executed[h.Pc] = true
		if h.Instructions[h.Pc].Op == nop {
			h.Pc++
			continue
		}
		if h.Instructions[h.Pc].Op == jmp {
			h.Pc += h.Instructions[h.Pc].Value
			continue
		}
		if h.Instructions[h.Pc].Op == acc {
			h.Pc++
		}
	}
	return true
}

func (h *Handheld) FixItFixItFixIt() {
	for {
		for i, instruction := range h.Instructions {
			if instruction.Op == jmp {
				h.Instructions[i].Op = nop
				if h.WillSolve() {
					h.Reset()
					return
				} else {
					h.Instructions[i].Op = jmp
				}
			}
			if instruction.Op == nop {
				h.Instructions[i].Op = jmp
				if h.WillSolve() {
					h.Reset()
					return
				} else {
					h.Instructions[i].Op = nop
				}
			}
		}
	}

}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string) (answer int) {
	instructions := []Instruction{}
	data := strings.Fields(puzzle)
	for i := 0; i < len(data); i += 2 {
		val, _ := strconv.Atoi(data[i+1])
		instructions = append(instructions, Instruction{Op: ToCode(data[i]), Value: val})
	}
	handheld := NewHandheld(instructions)
	handheld.FixItFixItFixIt()

	return handheld.Solve()
}

func main() {
	execution.Run(solutionLvl1, solutionLvl2, testcases)
}
