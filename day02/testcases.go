package main

import (
	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

var testcases = tc.TestCases{
	{
		`1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`, // input
		`2`, // expected lvl 1
		`1`, // expected lvl 2
	},
}
