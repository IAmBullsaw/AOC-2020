package main

import (
	"math"
	"sort"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
	"github.com/IAmBullsaw/AOC-2020/pkg/inp"
)

type Pair struct {
	First, Second int
}

func getChoice(data []int, current int) (int, float64) {
	options := []Pair{}
	for i, d := range data {
		diff := current - d
		if diff <= 0 && diff >= -3 {
			options = append(options, Pair{First: i, Second: d})
		}
		if diff > -3 {
			break
		}
	}
	if len(options) == 0 {
		return 0, math.Inf(1)
	}
	return options[0].First, float64(options[0].Second)
}

func erase(data []int, i int) []int {
	// fmt.Println("erase", data, i)
	if len(data) == 0 {
		return data
	}
	if len(data) == i {
		return data[:i-1]
	}
	copy(data[i:], data[i+1:])
	data = data[:len(data)-1]
	return data
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	data := inp.ToInts(puzzle)
	outlet := 0
	diff1 := 0
	diff3 := 0
	sort.Ints(data)
	data = append(data, data[len(data)-1]+3)
	for {
		i, choice := getChoice(data, outlet)
		if math.IsInf(choice, 1) {
			break
		}
		diff := outlet - int(choice)
		if diff == -3 {
			diff3++
		} else if diff == -1 {
			diff1++
		}
		outlet = int(choice)
		data = erase(data, i)

	}

	return diff1 * diff3
}

func recurse(data []int, current []int, outlet int, result *[][]int) {
	options := []Pair{}
	for i, v := range data {
		diff := outlet - v
		if diff < 0 && diff >= -3 {
			options = append(options, Pair{First: i, Second: v})
		}
		if diff < -3 {
			break
		}
	}

	for _, v := range options {
		tmp := make([]int, len(data))
		copy(tmp, data)
		tmp = erase(tmp, v.First)
		tmp2 := make([]int, len(current))
		copy(tmp2, append(current, v.Second))
		recurse(tmp, tmp2, v.Second, result)
	}

	// Seeing no option
	// Santa watching over you
	// Writes neatly: Naughty
	if len(options) == 0 {
		*result = append(*result, current)
		return
	}
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {

	data := inp.ToInts(puzzle)
	sort.Ints(data)
	data = append(data, data[len(data)-1]+3) // add ours
	outlet := 0
	result := [][]int{}
	current := []int{}

	recurse(data, current, outlet, &result)
	return len(result)
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
