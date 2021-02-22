package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

func toMap(puzzle string) map[string]string {
	data := strings.Split(puzzle, "\n\n")
	rules := map[string]string{}
	for _, line := range strings.Split(data[0], "\n") {
		dline := strings.Split(line, ": ")
		rules[dline[0]] = strings.TrimSpace(dline[1])
	}
	return rules
}

func expandRules(rules map[string]string, index string) string {
	rule := expandRule(rules, rules[index])
	return rule
}

func expandRule(rules map[string]string, expand string) string {
	expanded := strings.Fields(expand)
	for i, v := range expanded {
		if v[0] == '"' {
			// we got a literal
			literal := strings.Split(v, `"`)[1]
			return literal
		} else if unicode.IsDigit(rune(v[0])) {
			// we got a rule
			fmt.Println(i, v)
			v = expandRule(rules, v)
		}
	}
	return ""
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string, parameters map[string]int) (answer int) {
	rules := toMap(puzzle)
	expanded := expandRules(rules, "0")
	fmt.Println(expanded)

	return
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string, parameters map[string]int) (answer int) {
	return
}

var lvl1Parameters = map[string]int{}
var lvl2Parameters = map[string]int{}

func main() {
	execution.Run(solutionLvl1, lvl1Parameters, solutionLvl2, lvl2Parameters, testcases)
}
