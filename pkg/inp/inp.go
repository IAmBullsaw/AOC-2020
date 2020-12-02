package inp

import (
	"strconv"
	"strings"
)

func ToSlice(text string) (result []string) {
	for _, a := range strings.Fields(text) {
		result = append(result, a)
	}
	return
}

// ToInts return a slice of ints from passed in text
func ToInts(text string) (result []int) {
	for _, a := range strings.Fields(text) {
		ai, _ := strconv.Atoi(a)
		result = append(result, ai)
	}
	return
}

// ToIntsFunc return a slice of ints from passed in text, split on f
func ToIntsFunc(text string, f func(rune) bool) (result []int) {
	for _, a := range strings.FieldsFunc(text, f) {
		ai, _ := strconv.Atoi(a)
		result = append(result, ai)
	}
	return
}
