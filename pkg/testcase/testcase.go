package testcase

import (
	"fmt"
)

type TestCases []TestCase

type TestCase struct {
	Input string
	ExpectedLvl1, ExpectedLvl2 interface{}
}

func (tcs TestCases) Execute(solution func(string) (interface{}, interface{})) (total_pass bool) {
	total_pass = true
	for _, test := range tcs {
		resultLvl1, resultLvl2 := solution(test.Input)
		lvl1, expectedLvl1, lvl2, expectedLvl2 := fmt.Sprint(resultLvl1), fmt.Sprint(test.ExpectedLvl1), fmt.Sprint(resultLvl2), fmt.Sprint(test.ExpectedLvl2)
		passedLvl1 := lvl1 == expectedLvl1 || expectedLvl1 == ""
		passedLvl2 := lvl2 == expectedLvl2 || expectedLvl2 == ""
		passed := passedLvl1 && passedLvl2

		if !passed {
			total_pass = false
			fmt.Println("test case: ", test.Input)
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