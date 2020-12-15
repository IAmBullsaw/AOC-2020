package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	target := parameters["target"]
	nums := make(map[int]int) // number: turn
	data := strings.Split(puzzle, ",")
	for i, v := range data {
		i++
		val, err := strconv.Atoi(v)
		if err != nil {
			v = strings.TrimSpace(v)
			val, err = strconv.Atoi(v)
			if err != nil {
				fmt.Println("HOLY SMOKES, BATMAN!")
			}
		}
		nums[val] = i
	}

	var curr int // start at 0
	var next int
	for i := len(data) + 1; i < target; i++ {
		if last, ok := nums[curr]; ok {
			next = i - last
		} else {
			next = 0
		}
		nums[curr] = i
		curr = next
	}

	return curr
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	target := parameters["target"]
	nums := make(map[int]int) // number: turn
	data := strings.Split(puzzle, ",")
	for i, v := range data {
		i++
		val, err := strconv.Atoi(v)
		if err != nil {
			v = strings.TrimSpace(v)
			val, err = strconv.Atoi(v)
			if err != nil {
				fmt.Println("HOLY SMOKES, BATMAN!")
			}
		}
		nums[val] = i
	}

	var curr int // start at 0
	var next int
	for i := len(data) + 1; i < target; i++ {
		if last, ok := nums[curr]; ok {
			next = i - last
		} else {
			next = 0
		}
		nums[curr] = i
		curr = next
	}

	return curr
}

var lvl1Parameters = map[string]int{"target": 2020}
var lvl2Parameters = map[string]int{"target": 30000000}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
