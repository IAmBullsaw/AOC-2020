package main

import (
	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

var testcases = tc.TestCases{
	{
		Input: `939
7,13,x,x,59,x,31,19
`,
		Parameters:   map[string]int{"start": 1},
		ExpectedLvl1: `295`,
		ExpectedLvl2: `1068781`,
	},
	{
		Input: `939
67,7,59,61
	`,
		Parameters:   map[string]int{"start": 1},
		ExpectedLvl1: ``,
		ExpectedLvl2: `754018`,
	},
	{
		Input: `939
67,x,7,59,61
	`,
		Parameters:   map[string]int{"start": 1},
		ExpectedLvl1: ``,
		ExpectedLvl2: `779210`,
	},
	{
		Input: `939
67,7,x,59,61
	`,
		Parameters:   map[string]int{"start": 1},
		ExpectedLvl1: ``,
		ExpectedLvl2: `1261476`,
	},
	{
		Input: `939
1789,37,47,1889
	`,
		Parameters:   map[string]int{"start": 1000000000},
		ExpectedLvl1: ``,
		ExpectedLvl2: `1202161486`,
	},
}
