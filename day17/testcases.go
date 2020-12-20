package main

import (
	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

var testcases = tc.TestCases{
	{
		Input: `.#.
..#
###
`,
		Parameters:   map[string]int{},
		ExpectedLvl1: `112`,
		ExpectedLvl2: ``,
	},
}
