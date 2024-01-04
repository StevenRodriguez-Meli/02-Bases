package calculations

func MinimumFunc(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}
