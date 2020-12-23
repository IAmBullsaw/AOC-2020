package main

import (
	"fmt"
	"strconv"
	"strings"

	cmp "github.com/IAmBullsaw/AOC-2020/pkg/compare"
	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

type Rule struct {
	Title                            string
	LeftLo, LeftHi, RightLo, RightHi int
}

func (r Rule) ValidateLeft(i int) bool {
	return cmp.InBounds(r.LeftLo, i, r.LeftHi, true) == cmp.Within
}
func (r Rule) ValidateRight(i int) bool {
	return cmp.InBounds(r.RightLo, i, r.RightHi, true) == cmp.Within
}
func (r Rule) Validate(i int) bool {
	return r.ValidateLeft(i) || r.ValidateRight(i)
}

func (r Rule) String() string {
	return fmt.Sprintf("%s", r.Title)
}

func data(puzzle string) ([]Rule, []int, [][]int) {
	rules := []Rule{}
	ticket := []int{}
	tickets := [][]int{}
	var readRules, yourTicket, nearbyTickets bool
	readRules = true
	for _, i := range strings.Split(puzzle, "\n") {
		if i == "" {
			continue
		}
		if i == "your ticket:" {
			yourTicket = true
			readRules = false
			continue
		}
		if yourTicket {
			// rules are done
			for _, snum := range strings.Split(i, ",") {
				num, _ := strconv.Atoi(snum)
				ticket = append(ticket, num)
			}
			yourTicket = false
			continue
		}
		if i == "nearby tickets:" {
			nearbyTickets = true
			continue
		}
		if nearbyTickets {
			// your ticket is done
			t := []int{}
			for _, snum := range strings.Split(i, ",") {
				num, _ := strconv.Atoi(snum)
				t = append(t, num)
			}
			tickets = append(tickets, t)
			continue
		}
		if readRules {
			data := strings.Split(i, ":")
			if len(data) == 2 {
				// parse rules
				bothLimits := strings.Split(data[1], " or ")
				lelimit := strings.Split(bothLimits[0], "-")
				rilimit := strings.Split(bothLimits[1], "-")
				lelo, _ := strconv.Atoi(lelimit[0][1:])
				lehi, _ := strconv.Atoi(lelimit[1])
				rilo, _ := strconv.Atoi(rilimit[0])
				rihi, _ := strconv.Atoi(rilimit[1])
				rules = append(rules, Rule{Title: data[0],
					LeftLo: lelo, LeftHi: lehi,
					RightLo: rilo, RightHi: rihi})
			}
		}
	}
	return rules, ticket, tickets
}

func data2(puzzle string) ([]Rule, []int, [][]int) {
	rules := []Rule{}
	ticket := []int{}
	tickets := [][]int{}
	var readRules, yourTicket, nearbyTickets bool
	readRules = true
	for _, i := range strings.Split(puzzle, "\n") {
		if i == "" {
			continue
		}
		if i == "your ticket:" {
			yourTicket = true
			readRules = false
			continue
		}
		if yourTicket {
			// rules are done
			for _, snum := range strings.Split(i, ",") {
				num, _ := strconv.Atoi(snum)
				ticket = append(ticket, num)
			}
			yourTicket = false
			continue
		}
		if i == "nearby tickets:" {
			nearbyTickets = true
			continue
		}
		if nearbyTickets {
			// your ticket is done
			t := []int{}
			for _, snum := range strings.Split(i, ",") {
				num, _ := strconv.Atoi(snum)
				t = append(t, num)
			}
			tickets = append(tickets, t)
			continue
		}
		if readRules {
			data := strings.Split(i, ":")
			if len(data) == 2 {
				// parse rules
				bothLimits := strings.Split(data[1], " or ")
				lelimit := strings.Split(bothLimits[0], "-")
				rilimit := strings.Split(bothLimits[1], "-")
				lelo, _ := strconv.Atoi(lelimit[0][1:])
				lehi, _ := strconv.Atoi(lelimit[1])
				rilo, _ := strconv.Atoi(rilimit[0])
				rihi, _ := strconv.Atoi(rilimit[1])
				rules = append(rules, Rule{Title: data[0],
					LeftLo: lelo, LeftHi: lehi,
					RightLo: rilo, RightHi: rihi})
			}
		}
	}
	return rules, ticket, tickets
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	rules, _, tickets := data(puzzle)

	for _, t := range tickets {
		for _, num := range t {
			any := false
			for _, r := range rules {
				if r.ValidateLeft(num) || r.ValidateRight(num) {
					any = true
					break
				}
			}
			if !any {
				answer += num
			}
		}
	}
	return
}

func isValid(rules []Rule, ticket []int) bool {
	for _, value := range ticket {
		for _, r := range rules {
			if r.Validate(value) {
				return true
			}
		}
	}
	return false
}

func setAdd(possible *map[int][]Rule, i int, r Rule) bool {
	for _, rule := range (*possible)[i] {
		if r == rule {
			return false
		}
	}
	(*possible)[i] = append((*possible)[i], r)
	return true
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	rules, ourTicket, tickets := data2(puzzle)
	valid := [][]int{ourTicket}
	for _, ticket := range tickets {
		if isValid(rules, ticket) {
			valid = append(valid, ticket)
		}
	}

	matches := map[int][]Rule{}
	used := map[Rule]bool{}
	for {
		for _, r := range rules {
			_, ok := used[r]
			if ok {
				continue
			}
			for i := 0; i < len(ourTicket); i++ {
				count := 0
				for _, ticket := range valid {
					if r.Validate(ticket[i]) {
						count++
					}
				}
				if count == len(valid) {
					setAdd(&matches, i, r)
					used[r] = true
				}
			}
		}
		if len(used) == len(rules) {
			break
		}

	}

	chosen := map[Rule]int{}
	for {
		clean := false
		rule := Rule{}
		for col, candidates := range matches {
			if len(candidates) == 1 {
				chosen[candidates[0]] = col
				clean = true
				rule = candidates[0]
				candidates = []Rule{}
				delete(matches, col)
			}
		}
		if clean {
			for i, candidates := range matches {
				for j, c := range candidates {
					if c == rule {
						// TODO remove rule from candidates
						candidates[j] = candidates[len(candidates)-1]
						candidates = candidates[:len(candidates)-1]
						matches[i] = candidates
						break
					}
				}
			}
		}
		if len(matches) == 0 {
			break
		}
	}
	answer = 1
	for k, v := range chosen {
		if strings.HasPrefix(k.Title, "departure") {
			answer *= ourTicket[v]
		}
	}
	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
