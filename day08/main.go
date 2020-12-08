package main

import (
	"strconv"
	"strings"

	com "github.com/IAmBullsaw/AOC-2020/pkg/computer"
	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

func Solve(instructions []com.Instruction) int {
	operations := make(map[interface{}]func(*com.Computer, *com.Instruction))
	operations["nop"] = func(c *com.Computer, _ *com.Instruction) {
		c.Pc++
	}
	operations["jmp"] = func(c *com.Computer, i *com.Instruction) {
		val := i.Value.(int)
		c.Pc += val
	}
	operations["acc"] = func(c *com.Computer, i *com.Instruction) {
		val := i.Value.(int)
		c.Acc += val
		c.Pc++
	}
	c := com.NewComputer(instructions, operations)
	req := func(c *com.Computer) bool {
		if c.Pc >= len(c.Instructions) || c.HasExecuted() {
			return true
		}
		return false
	}
	c.Solve(req)
	return c.Acc
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string) (answer int) {
	instructions := []com.Instruction{}
	data := strings.Fields(puzzle)
	for i := 0; i < len(data); i += 2 {
		val, _ := strconv.Atoi(data[i+1])
		instructions = append(instructions, com.Instruction{Operation: data[i], Value: val})
	}

	return Solve(instructions)
}

func Solve2(instructions []com.Instruction) int {
	operations := make(map[interface{}]func(*com.Computer, *com.Instruction))
	operations["nop"] = func(c *com.Computer, _ *com.Instruction) {
		c.Pc++
	}
	operations["jmp"] = func(c *com.Computer, i *com.Instruction) {
		val := i.Value.(int)
		c.Pc += val
	}
	operations["acc"] = func(c *com.Computer, i *com.Instruction) {
		val := i.Value.(int)
		c.Acc += val
		c.Pc++
	}
	c := com.NewComputer(instructions, operations)
	req := func(c *com.Computer) bool {
		if c.Pc >= len(c.Instructions) {
			c.Status["exit"] = 0
			return true
		}
		if c.HasExecuted() {
			c.Status["exit"] = 1
			return true
		}
		return false
	}
	for i, in := range c.Instructions {
		if in.Operation == "jmp" {
			c.ChangeInstructionOperation(i, "nop")
			c.Solve(req)
			if c.Status["exit"] == 0 {
				return c.Acc
			} else {
				c.Reset(true, true)
				continue
			}
		}

		if in.Operation == "nop" {
			c.ChangeInstructionOperation(i, "jmp")
			c.Solve(req)
			if c.Status["exit"] == 0 {
				return c.Acc
			} else {
				c.Reset(true, true)
				continue
			}
		}
	}

	return 0
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string) (answer int) {
	instructions := []com.Instruction{}
	data := strings.Fields(puzzle)
	for i := 0; i < len(data); i += 2 {
		val, _ := strconv.Atoi(data[i+1])
		instructions = append(instructions, com.Instruction{Operation: data[i], Value: val})
	}

	return Solve2(instructions)
}

func main() {
	execution.Run(solutionLvl1, solutionLvl2, testcases)
}
