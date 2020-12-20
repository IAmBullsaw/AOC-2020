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

func (r Rule) ValidateBoth(i int) bool {
	return r.ValidateLeft(i) && r.ValidateRight(i)
}

func (r Rule) String() string {
	return fmt.Sprintf("%s: %d-%d or %d-%d", r.Title, r.LeftLo, r.LeftHi, r.RightLo, r.RightHi)
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

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	rules, _, tickets := data2(puzzle)
	matches := map[int][]string{}

	for _, t := range tickets {
		for i, num := range t {
			any := false
			rule := rules[0]
			for _, r := range rules {
				if r.ValidateLeft(num) || r.ValidateRight(num) {
					any = true
					rule = r
					break
				}
			}
			if !any {
				matches[i] = append(matches[i], rule.Title)
			}
		}
	}
	fmt.Println(matches)
	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
