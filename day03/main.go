package main

import (
	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
	ip "github.com/IAmBullsaw/AOC-2020/pkg/inp"
)

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	inp := ip.ToSlice(puzzle)
	dAdd := 1
	rAdd := 3
	d := dAdd
	r := rAdd
	rMax := len(inp[0])
	tree := '#'
	for d < len(inp) {
		if inp[d][r] == byte(tree) {
			answer++
		}

		d = d + dAdd
		r = (r + rAdd) % rMax
	}
	return
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	inp := ip.ToSlice(puzzle)
	tree := '#'

	slope := func(rAdd, dAdd int) (result int) {
		d := dAdd
		r := rAdd
		rMax := len(inp[0])
		for d < len(inp) {
			if inp[d][r] == byte(tree) {
				result++
			}

			d = d + dAdd
			r = (r + rAdd) % rMax
		}
		return

	}

	answer = slope(1, 1) * slope(3, 1) * slope(5, 1) * slope(7, 1) * slope(1, 2)
	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
