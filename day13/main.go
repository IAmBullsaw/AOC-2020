package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	data := strings.Split(puzzle, "\n")
	start, _ := strconv.Atoi(data[0])
	busIds := []int{}
	for _, b := range strings.Split(data[1], ",") {
		if b != "x" {
			id, _ := strconv.Atoi(b)
			busIds = append(busIds, id)
		}
	}
	time := start
Solved1:
	for {
		for _, b := range busIds {
			if time%b == 0 {
				answer = (time - start) * b
				break Solved1
			}
		}
		time++
	}

	return
}

func departsAt(time, bus int) bool {
	fmt.Println(bus)
	return bus == -1 || (time%bus) == 0
}

func all(l []bool) bool {
	for _, b := range l {
		if !b {
			return false
		}
	}
	return true
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	data := strings.Split(puzzle, "\n")
	start := 0 //100000000000000
	busIds := []int{}
	for _, b := range strings.Split(data[1], ",") {
		if b != "x" {
			id, _ := strconv.Atoi(b)
			busIds = append(busIds, id)
		} else {
			busIds = append(busIds, -1)
		}
	}
	time := start
Solve2:
	for {
		check := []bool{}
		for range busIds {
			check = append(check, false)
		}
		for i, b := range busIds {
			if departsAt(time, b) {
				check[i] = true
			}
		}
		if all(check) {
			break Solve2
		} else {
			for i := range busIds {
				check[i] = false
			}
		}
		time++
	}

	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
