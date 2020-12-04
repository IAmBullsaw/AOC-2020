package main

import (
	"strconv"
	"strings"

	"regexp"

	cmp "github.com/IAmBullsaw/AOC-2020/pkg/compare"
	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

// getInput return a list of maps corresponding to one passport
func getInput(puzzle string) (pps []map[string]string) {
	pp := map[string]string{}
	for _, data := range strings.Split(puzzle, "\n\n") {

		for _, values := range strings.Fields(data) {
			kv := strings.Split(values, ":")
			pp[kv[0]] = kv[1]
		}
		pps = append(pps, pp)
		pp = map[string]string{}
	}
	return
}

// Fields should all be in a Passport
var Fields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string) (answer int) {
	pps := getInput(puzzle)

	for _, passport := range pps {
		if len(passport) < 7 || (len(passport) == 7 && passport["cid"] != "") {
			continue
		}
		for _, f := range Fields {
			if passport[f] == "" {
				goto end1
			}
		}
		answer++
	end1:
	}

	return
}

// Policy has a Rule pattern which match should adhere to the Validate function
type Policy struct {
	Rule     *regexp.Regexp
	Validate func(string) bool
}

// Predetermined Policies which each Passport field must follow (cid excluded)
var Policies = map[string]Policy{
	"byr": Policy{Rule: regexp.MustCompile(`^\d{4}$`), Validate: func(data string) bool {
		val, _ := strconv.Atoi(data)
		return cmp.InBounds(1920, val, 2002, true) == cmp.Within
	}},
	"iyr": Policy{Rule: regexp.MustCompile(`^\d{4}$`), Validate: func(data string) bool {
		val, _ := strconv.Atoi(data)
		return cmp.InBounds(2010, val, 2020, true) == cmp.Within
	}},
	"eyr": Policy{Rule: regexp.MustCompile(`^\d{4}$`), Validate: func(data string) bool {
		val, _ := strconv.Atoi(data)
		return cmp.InBounds(2020, val, 2030, true) == cmp.Within
	}},
	"hgt": Policy{Rule: regexp.MustCompile(`^\d+cm$|^\d+in$`), Validate: func(data string) bool {
		val, _ := strconv.Atoi(data[:len(data)-2])
		if data[len(data)-2:] == "cm" {
			return cmp.InBounds(150, val, 193, true) == cmp.Within
		} else {
			return cmp.InBounds(59, val, 76, true) == cmp.Within
		}
	}},
	"hcl": Policy{Rule: regexp.MustCompile(`^#\w{6}$`), Validate: func(data string) bool { return true }},
	"ecl": Policy{Rule: regexp.MustCompile(`^\w{3}$`), Validate: func(data string) bool {
		colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, color := range colors {
			if data == color {
				return true
			}
		}
		return false
	}},
	"pid": Policy{Rule: regexp.MustCompile(`^\d{9}$`), Validate: func(data string) bool { return true }},
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string) (answer int) {
	pps := getInput(puzzle)

	for _, passport := range pps {
		if len(passport) < 7 || (len(passport) == 7 && passport["cid"] != "") {
			continue
		} else {
			for key, policy := range Policies {
				if passport[key] == "" ||
					!(policy.Rule.MatchString(passport[key]) && policy.Validate(passport[key])) {
					goto end2
				}
			}
		}
		answer++
	end2:
	}

	return
}

func main() {
	execution.Run(solutionLvl1, solutionLvl2, testcases)
}
