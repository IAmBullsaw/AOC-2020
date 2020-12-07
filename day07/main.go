package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

type BagPairs []BagPair

type BagPair struct {
	Num  int
	Name string
}

func contains(rules map[string][]BagPair, name string, target string) bool {
	if name == target {
		return true
	}
	for _, bp := range rules[name] {
		if contains(rules, bp.Name, target) {
			return true
		}
	}
	return false
}

func countBFS(rules map[string]BagPairs, parents map[string]BagPairs, target string) (count int) {
	queue := BagPairs{BagPair{Name: target, Num: 1}}
	seen := make(map[string]interface{})

	for len(queue) > 0 {
		fmt.Println(queue)
		curr := queue[0]
		queue = queue[1:]
		if seen[curr.Name] == true {
			continue
		}
		seen[curr.Name] = true
		for _, y := range parents[curr.Name] {
			queue = append(queue, y)
		}
	}

	return
}

func count(rules map[string][]BagPair, name string) (count int) {
	for _, v := range rules {
		for _, bp := range v {
			if contains(rules, bp.Name, "shinygold") {
				count++
				break
			}
		}
	}
	return
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string) (answer int) {
	rules := make(map[string][]BagPair)
	for _, line := range strings.Split(puzzle, "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		bag := fields[0] + fields[1]
		contains := fields[4:]

		for i := 0; i < len(contains)-3; i += 4 {
			num, _ := strconv.Atoi(contains[i])
			name := contains[i+1] + contains[i+2]
			rules[bag] = append(rules[bag], BagPair{Num: num, Name: name})
		}
	}
	return count(rules, "shinygold")
}

func countAll(rules map[string]BagPairs, parents map[string]BagPairs, target string) (count int) {
	count = 1
	for _, v := range rules[target] {
		count += v.Num * countAll(rules, parents, v.Name)
	}
	return
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string) (answer int) {
	parents := make(map[string]BagPairs)
	rules := make(map[string]BagPairs)
	for _, line := range strings.Split(puzzle, "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		bag := fields[0] + fields[1]
		contains := fields[4:]
		if contains[0] == "no" {
			continue
		}

		for i := 0; i < len(contains)-3; i += 4 {
			num, _ := strconv.Atoi(contains[i])
			name := contains[i+1] + contains[i+2]
			rules[bag] = append(rules[bag], BagPair{Num: num, Name: name})
			parents[name] = append(parents[name], BagPair{Num: num, Name: bag})
		}
	}
	return countAll(rules, parents, "shinygold") - 1
}

func main() {
	execution.Run(solutionLvl1, solutionLvl2, testcases)
}
