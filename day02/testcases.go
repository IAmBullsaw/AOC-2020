package main

import (
	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

var testcases = tc.TestCases{
	{
		Input: `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`, // input
		ExpectedLvl1: `2`, // expected lvl 1
		ExpectedLvl2: `1`, // expected lvl 2
	},
}
