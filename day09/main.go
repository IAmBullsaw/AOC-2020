package main

import (
	"fmt"
	"sort"

	cmp "github.com/IAmBullsaw/AOC-2020/pkg/compare"
	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
	inp "github.com/IAmBullsaw/AOC-2020/pkg/inp"
)

func Match(numbers []int, target int) bool {
	sort.Ints(numbers)
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i] == numbers[j] {
			//          i     j
			// [1 2 3 4 5 5 5 5 5]
			return false
		}
		sum := numbers[i] + numbers[j]
		if sum == target {
			return true
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return false

}

func contigousEnds(numbers []int, target int) int {
	i := 0
	j := 1
	sum := numbers[i]
	for j < len(numbers) {
		if sum == target && i+1 < j {
			return cmp.MinInts(numbers[i:j]) + cmp.MaxInts(numbers[i:j])
		} else if sum < target && j < len(numbers) {
			sum += numbers[j]
			j++
		} else { // if sum > target {
			sum -= numbers[i]
			i++
		}

	}
	return -1
}

func simpleMatch(numbers []int, target int) bool {
	for _, a := range numbers {
		for _, b := range numbers {
			if a != b && a+b == target {
				return true
			}
		}
	}

	return false
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	data := inp.ToInts(puzzle)
	size := parameters["size"]
	for i := size + 1; i < len(data); i++ {
		tmp := make([]int, size)
		copy(tmp, data[i-size:i])
		if !Match(tmp, data[i]) {
			answer = data[i]
			break
		}
	}
	fmt.Println("lvl1", answer)
	return
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	target := solutionLvl1(puzzle, parameters)
	return contigousEnds(inp.ToInts(puzzle), target)
}

var lvl1Parameters = map[string]int{"size": 25}
var lvl2Parameters = map[string]int{"size": 25}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
