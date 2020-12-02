package count

// CountInt return the count of target in values
func CountInt(values []int, target int) (count int) {
	for _, v := range values {
		if v == target {
			count++
		}
	}
	return
}
