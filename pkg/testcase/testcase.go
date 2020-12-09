package testcase

import (
	"fmt"
)

type TestCases []TestCase

type TestCase struct {
	Input                      string
	Parameters                 map[string]int
	ExpectedLvl1, ExpectedLvl2 interface{}
}

func (tcs TestCases) Execute(solutionLvl1 func(string, map[string]int) int, solutionLvl2 func(string, map[string]int) int) (totalPass bool) {
	totalPass = true
	for i, test := range tcs {
		fmt.Println("Test case:", i+1)
		resultLvl1 := solutionLvl1(test.Input, test.Parameters)
		resultLvl2 := solutionLvl2(test.Input, test.Parameters)
		lvl1, expectedLvl1, lvl2, expectedLvl2 := fmt.Sprint(resultLvl1), fmt.Sprint(test.ExpectedLvl1), fmt.Sprint(resultLvl2), fmt.Sprint(test.ExpectedLvl2)
		passedLvl1 := lvl1 == expectedLvl1 || expectedLvl1 == ""
		passedLvl2 := lvl2 == expectedLvl2 || expectedLvl2 == ""
		passed := passedLvl1 && passedLvl2

		if !passed {
			totalPass = false
			fmt.Println("Test case:", i+1, "Failed.")
		} else {
			fmt.Println("Test case:", i+1, "Passed.")
		}
		if !passedLvl1 {
			fmt.Println("  Level 1: ", lvl1, " expected: ", expectedLvl1)
		}
		if !passedLvl2 {
			fmt.Println("  Level 2: ", lvl2, " expected: ", expectedLvl2)
		}
	}
	return
}
