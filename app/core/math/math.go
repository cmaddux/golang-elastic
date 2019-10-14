package math

// MinInt returns the minimum integer given two values
func MinInt(v1 int, v2 int) int {
	if v2 < v1 {
		return v2
	}

	return v1
}
