package main

import (
	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

var testcases = tc.TestCases{
	{
		Input: `abc

		a
		b
		c

		ab
		ac

		a
		a
		a
		a

		b
		`,
		ExpectedLvl1: `11`,
		ExpectedLvl2: `6`,
	},
}
