package main

import (
	"strconv"
	"strings"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
	ip "github.com/IAmBullsaw/AOC-2020/pkg/inp"
)

// solution_lvl1 return answers for level 1
func solution_lvl1(puzzle string) interface{} {
	answer := 0
	inp := ip.ToSlice(puzzle)
	for i := 0; i < len(inp); i += 3 {
		lohi := strings.Split(inp[i], "-")

		lo, _ := strconv.Atoi(lohi[0])
		hi, _ := strconv.Atoi(lohi[1])
		b := inp[i+1][0]
		pwd := inp[i+2]

		count := 0
		for i := 0; i < len(pwd); i++ {
			if pwd[i] == b {
				count++
			}
		}
		if lo <= count && count <= hi {
			answer++
		}
	}
	return answer
}

// solution _lvl2 return answers for level 2
func solution_lvl2(puzzle string) (answer interface{}) {
	valid := 0
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
			valid++
		}

	}
	answer = valid
	return
}

func main() {
	execution.Run(solution_lvl1, solution_lvl2, testcases)
}
