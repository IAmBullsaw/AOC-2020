package main

import (
	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

var testcases = tc.TestCases{
	{
		Input: `1 + 2 * 3 + 4 * 5 + 6
`,
		Parameters:   map[string]int{},
		ExpectedLvl1: `71`,
		ExpectedLvl2: ``,
	},
	{
		Input: `1 + 2 * 3 + 4 * 5 + 6
2 * 3 + (4 * 5)
`,
		Parameters:   map[string]int{},
		ExpectedLvl1: `97`,
		ExpectedLvl2: ``,
	},
	{
		Input: `2 * 3 + (4 * 5)
`,
		Parameters:   map[string]int{},
		ExpectedLvl1: `26`,
		ExpectedLvl2: ``,
	},
	{
		Input: `5 + (8 * 3 + 9 + 3 * 4 * 3)
`,
		Parameters:   map[string]int{},
		ExpectedLvl1: `437`,
		ExpectedLvl2: ``,
	},
	{
		Input: `5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
`,
		Parameters:   map[string]int{},
		ExpectedLvl1: `12240`,
		ExpectedLvl2: ``,
	},
	{
		Input: `((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
`,
		Parameters:   map[string]int{},
		ExpectedLvl1: `13632`,
		ExpectedLvl2: ``,
	},
}
