package testcase

import (
	"fmt"
)

type TestCases []TestCase

type TestCase struct {
	Input                      string
	ExpectedLvl1, ExpectedLvl2 interface{}
}

func (tcs TestCases) Execute(solution_lvl1 func(string) interface{}, solution_lvl2 func(string) interface{}) (total_pass bool) {
	total_pass = true
	for i, test := range tcs {
		fmt.Println("Test case:", i+1)
		resultLvl1 := solution_lvl1(test.Input)
		resultLvl2 := solution_lvl2(test.Input)
		lvl1, expectedLvl1, lvl2, expectedLvl2 := fmt.Sprint(resultLvl1), fmt.Sprint(test.ExpectedLvl1), fmt.Sprint(resultLvl2), fmt.Sprint(test.ExpectedLvl2)
		passedLvl1 := lvl1 == expectedLvl1 || expectedLvl1 == ""
		passedLvl2 := lvl2 == expectedLvl2 || expectedLvl2 == ""
		passed := passedLvl1 && passedLvl2

		if !passed {
			total_pass = false
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
