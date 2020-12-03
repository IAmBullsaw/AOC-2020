package main

import (
	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
	ip "github.com/IAmBullsaw/AOC-2020/pkg/inp"
)

// solution_lvl1 return answers for level 1
func solution_lvl1(puzzle string) (answer interface{}) {
	inp := ip.ToSlice(puzzle)
	dAdd := 1
	rAdd := 3
	d := dAdd
	r := rAdd
	rMax := len(inp[0])
	tree := '#'
	trees := 0
	for d < len(inp) {
		if inp[d][r] == byte(tree) {
			trees++
		}

		d = d + dAdd
		r = (r + rAdd) % rMax
	}
	return trees
}

// solution _lvl2 return answers for level 2
func solution_lvl2(puzzle string) (answer interface{}) {
	inp := ip.ToSlice(puzzle)

	slope := func(rAdd, dAdd int) (result int) {
		d := dAdd
		r := rAdd
		rMax := len(inp[0])
		tree := '#'
		for d < len(inp) {
			if inp[d][r] == byte(tree) {
				result++
			}

			d = d + dAdd
			r = (r + rAdd) % rMax
		}
		return

	}

	prod := slope(1, 1) * slope(3, 1) * slope(5, 1) * slope(7, 1) * slope(1, 2)

	answer = prod
	return
}

func main() {
	execution.Run(solution_lvl1, solution_lvl2, testcases)
}
