package main

// for internal test example
func CalcAvg(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}

	if sum == 0 {
		return 0
	}
	return sum / float64(len(values))
}

// for external test example
func IsString(str interface{}) bool {
	_, ok := str.(string)
	return ok
}
