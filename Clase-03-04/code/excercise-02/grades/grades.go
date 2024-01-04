package grades

func CalculateAverage(grades ...int) float64 {
	sum := 0
	count := 0

	for _, grade := range grades {
		if grade >= 0 {
			sum += grade
			count++
		}
	}

	if count > 0 {
		return float64(sum) / float64(count)
	}
	return 0
}
