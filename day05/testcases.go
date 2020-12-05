package main

import (
	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

var testcases = tc.TestCases{
	{
		Input: `FBFBBFFRLR
		`,
		ExpectedLvl1: `357`,
		ExpectedLvl2: ``,
	},
	{
		Input: `FFFFBFBLLL
		FFFFBFBLLR
		FFFFBFBLRL
		FFFFBFBRLL
		FFFFBFBRLR
		FFFFBFBRRL
`,
		ExpectedLvl1: ``,
		ExpectedLvl2: ``,
	},
}
