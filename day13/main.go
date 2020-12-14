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
	// fmt.Println(bus)
	if time == 0 || bus == 0 {
		fmt.Println(time, bus)
	}
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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	data := strings.Split(puzzle, "\n")

	start := parameters["start"] //100000000000000
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
	for {
		// Start at
		// time     bus 7   bus 13  bus 59  bus 31  bus 19
		// 1068780    .       .       .       .       .
		// 1068781    D       .       .       .       .
		// 1068782    .       D       .       .       .
		if departsAt(time, busIds[0]) && departsAt(time+1, busIds[1]) {
			break
		}
		time++
	}
Solve2:
	for {
		for i, b := range busIds {
			if !departsAt(time+i, b) {
				time += LCM(busIds[0], busIds[1], busIds[2:i]...)
				continue Solve2
			}
		}
		answer = time
		break
	}

	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{"start": 100000000000000}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
