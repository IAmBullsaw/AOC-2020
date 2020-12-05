package compare

type Bounds int

const (
	Within Bounds = iota
	Above
	Below
)

func (b Bounds) String() string {
	return [...]string{"Within", "Above", "Below"}[b]
}

// InBounds return -1 if under, 1 if within and 0 if above
func InBounds(x, v, y int, inclusive bool) Bounds {
	if inclusive {
		if x <= v && v <= y {
			return Within
		} else if x <= v {
			return Below
		} else {
			return Above
		}
	} else {
		if x < v && v < y {
			return Within
		} else if x < v {
			return Below
		} else {
			return Above
		}
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func MaxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
