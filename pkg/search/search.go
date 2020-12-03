package search

// Int returns nil or the index to the found int in values
func Int(values []int, target int) (found int) {
	for i, v := range values {
		if v == target {
			return i
		}
	}
	return
}

// IntFunc returns nil or the index to the int in values which satisfies f()
func IntFunc(values []int, f func(int) bool) (found int) {
	for i, v := range values {
		if f(v) {
			return i
		}
	}
	return
}

// IntAll returns nil or a list of indexes to the found int in values
func IntAll(values []int, target int) (found []int) {
	for i, v := range values {
		if v == target {
			found = append(found, i)
		}
	}
	return
}

// IntAllFunc returns nil or a list of indexes to ints in values which satisfies f()
func IntAllFunc(values []int, f func(int) bool) (found []int) {
	for i, v := range values {
		if f(v) {
			found = append(found, i)
		}
	}
	return
}
