package main

import (
	"strings"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string) (answer int) {
	for _, data := range strings.Split(puzzle, "\n\n") {
		set := map[rune]bool{}
		for _, values := range strings.Fields(data) {
			for _, c := range values {
				set[c] = true
			}
		}
		answer += len(set)
	}

	return
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string) (answer int) {
	for _, data := range strings.Split(puzzle, "\n\n") {
		set := map[rune]int{}
		size := 0
		for _, values := range strings.Fields(data) {
			for _, c := range values {
				set[c]++
			}
			size++
		}
		for _, v := range set {
			if v == size {
				answer++
			}
		}
	}
	return
}

func main() {
	execution.Run(solutionLvl1, solutionLvl2, testcases)
}
