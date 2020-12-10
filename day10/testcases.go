package main

import (
	tc "github.com/IAmBullsaw/AOC-2020/pkg/testcase"
)

var testcases = tc.TestCases{
	{
		Input: `16
10
15
5
1
11
7
19
6
12
4`,
		Parameters:   map[string]int{},
		ExpectedLvl1: `35`,
		ExpectedLvl2: `8`,
	},
	{
		Input: `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`,
		Parameters:   map[string]int{},
		ExpectedLvl1: `220`,
		ExpectedLvl2: `19208`,
	},
	{
		Input: `59
134
159
125
95
92
169
43
154
46
110
79
117
151
141
56
87
10
65
170
89
32
40
118
36
94
124
173
164
166
113
67
76
102
107
52
144
119
2
72
86
73
66
13
15
38
47
109
103
128
165
148
116
146
18
135
68
83
133
171
145
48
31
106
161
6
21
22
77
172
28
78
96
55
132
39
100
108
33
23
54
157
80
153
9
62
26
147
1
27
131
88
138
93
14
123
122
158
152
71
49
101
37
99
160
53
3
`,
		Parameters:   map[string]int{},
		ExpectedLvl1: `2516`,
		ExpectedLvl2: `296196766695424`,
	},
}
