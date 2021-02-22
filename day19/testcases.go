package main

import (
	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

var testcases = tc.TestCases{
	{
		Input: `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb
`,
		Parameters:   map[string]int{},
		ExpectedLvl1: `2`,
		ExpectedLvl2: ``,
	},
}
