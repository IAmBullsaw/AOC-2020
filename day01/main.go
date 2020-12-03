package main

import (
	"sort"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
	"github.com/IAmBullsaw/AOC-2020/pkg/inp"
)

// solution_lvl1 return answers for level 1
func solution_lvl1(puzzle string) (answer int) {
	inp := inp.ToInts(puzzle)
	sort.Ints(inp)

	i := 0
	j := len(inp) - 1
	for i <= j {
		target := 2020
		sum := inp[i] + inp[j]
		if sum == target {
			answer = inp[i] * inp[j]
			break
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return
}

// solution_lvl2 return answers for level 2
func solution_lvl2(puzzle string) (answer int) {
	inp := inp.ToInts(puzzle)
	sort.Ints(inp)
	answer = -1
	for i := 0; i < len(inp); i++ {
		j := i + 1
		k := len(inp) - 1
		for j <= k {
			target := 2020
			sum := inp[i] + inp[j] + inp[k]
			if sum == target {
				answer = inp[i] * inp[j] * inp[k]
				break
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
		if answer != -1 {
			break
		}
	}

	return
}

func main() {
	execution.Run(solution_lvl1, solution_lvl2, testcases)
}
