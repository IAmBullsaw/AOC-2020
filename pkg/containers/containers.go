package containers

type MinMax struct {
	min int
	max int
}

func NewMinMax(s [2]int) (mm MinMax) {
	if s[0] < s[1] {
		mm.min = s[0]
		mm.max = s[1]
	} else {
		mm.min = s[1]
		mm.max = s[0]
	}
	return
}
