package main

import (
	"strconv"
	"strings"

	cmp "github.com/IAmBullsaw/AOC-2020/pkg/compare"
	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
	ip "github.com/IAmBullsaw/AOC-2020/pkg/inp"
)

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	inp := ip.ToSlice(puzzle)
	for i := 0; i < len(inp); i += 3 {
		lohi := strings.Split(inp[i], "-")

		lo, _ := strconv.Atoi(lohi[0])
		hi, _ := strconv.Atoi(lohi[1])
		b := string(inp[i+1][0])
		pwd := inp[i+2]

		count := strings.Count(pwd, b)
		if cmp.InBounds(lo, count, hi, true) == cmp.Within {
			answer++
		}
	}
	return
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	inp := ip.ToSlice(puzzle)
	for i := 0; i < len(inp); i += 3 {
		lohi := strings.Split(inp[i], "-")

		p1, _ := strconv.Atoi(lohi[0])
		p2, _ := strconv.Atoi(lohi[1])
		p1--
		p2--
		b := inp[i+1][0]
		pwd := inp[i+2]

		if (pwd[p1] == b || pwd[p2] == b) && pwd[p1] != pwd[p2] {
			answer++
		}

	}
	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
