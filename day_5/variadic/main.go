package variadic

func Sum(nums ...int) (sum int) {
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return
}
