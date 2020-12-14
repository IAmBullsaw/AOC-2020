package main

import (
	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

var testcases = tc.TestCases{
	// 	{
	// 		Input: `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
	// mem[8] = 11
	// mem[7] = 101
	// mem[8] = 0
	// `,
	// 		Parameters:   map[string]int{},
	// 		ExpectedLvl1: `165`,
	// 		ExpectedLvl2: ``,
	// 	},
	{
		Input: `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`,
		Parameters:   map[string]int{},
		ExpectedLvl1: ``,
		ExpectedLvl2: `208`,
	},
}
