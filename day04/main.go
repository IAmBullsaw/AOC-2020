package main

import (
	"strconv"
	"strings"

	"regexp"

	cmp "github.com/IAmBullsaw/AOC-2020/pkg/compare"
	"github.com/IAmBullsaw/AOC-2020/pkg/execution"
)

// solutionLvl1 return answers for level 1
func solutionLvl1(puzzle string) (answer int) {
	pps := []map[string]string{}
	pp := map[string]string{}
	for _, data := range strings.Split(puzzle, "\n") {

		for _, values := range strings.Fields(data) {
			kv := strings.Split(values, ":")
			pp[kv[0]] = kv[1]
		}

		if len(data) == 0 {
			pps = append(pps, pp)
			pp = map[string]string{}
		}
	}

	fields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		"cid",
	}

	for _, passport := range pps {
		valid := true
		for _, f := range fields {
			if passport[f] == "" {
				if f != "cid" {
					valid = false
				}

			}
		}
		if valid {
			answer++
		}

	}

	return
}

type Policy struct {
	Rule     *regexp.Regexp
	Validate func(string) bool
}

// solutionLvl2 return answers for level 2
func solutionLvl2(puzzle string) (answer int) {
	pps := []map[string]string{}
	pp := map[string]string{}
	for _, data := range strings.Split(puzzle, "\n") {

		for _, values := range strings.Fields(data) {
			kv := strings.Split(values, ":")
			pp[kv[0]] = kv[1]
		}

		if len(data) == 0 {
			pps = append(pps, pp)
			pp = map[string]string{}
		}
	}

	fields := map[string]Policy{
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
		"cid": Policy{Rule: regexp.MustCompile(`.`), Validate: func(data string) bool { return true }},
	}

	for _, passport := range pps {
		valid := true
		for key, policy := range fields {
			if passport[key] == "" {
				if key != "cid" {
					valid = false
				}
			} else {
				if !(policy.Rule.MatchString(passport[key]) && policy.Validate(passport[key])) {
					valid = false
				}
			}

		}
		if valid {
			answer++
		}

	}

	return
}

func main() {
	execution.Run(solutionLvl1, solutionLvl2, testcases)
}
