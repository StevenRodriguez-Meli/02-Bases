package calculations

func AverageFunc(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum / len(nums)
}
