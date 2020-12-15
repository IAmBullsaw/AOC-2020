package main

import (
	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

var testcases = tc.TestCases{
	{
		Input: `0,3,6
`,
		Parameters:   map[string]int{"target": 2020},
		ExpectedLvl1: `436`,
		ExpectedLvl2: ``,
	},
	{
		Input: `1,3,2
`,
		Parameters:   map[string]int{"target": 2020},
		ExpectedLvl1: `1`,
		ExpectedLvl2: ``,
	},
	{
		Input: `2,1,3
`,
		Parameters:   map[string]int{"target": 2020},
		ExpectedLvl1: `10`,
		ExpectedLvl2: ``,
	},
	{
		Input: `1,2,3
`,
		Parameters:   map[string]int{"target": 2020},
		ExpectedLvl1: `27`,
		ExpectedLvl2: ``,
	},
	{
		Input: `2,3,1
`,
		Parameters:   map[string]int{"target": 2020},
		ExpectedLvl1: `78`,
		ExpectedLvl2: ``,
	},
	{
		Input: `3,2,1
`,
		Parameters:   map[string]int{"target": 2020},
		ExpectedLvl1: `438`,
		ExpectedLvl2: ``,
	},
	{
		Input: `3,1,2
`,
		Parameters:   map[string]int{"target": 2020},
		ExpectedLvl1: `1836`,
		ExpectedLvl2: ``,
	},
}
