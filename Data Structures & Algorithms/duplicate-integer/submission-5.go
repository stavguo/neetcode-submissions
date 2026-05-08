func hasDuplicate(nums []int) bool {
	if len(nums) < 1 {
		return false
	}
	// 1. sort
	sort.Ints(nums)
	// 2. create var for prev
	prev := nums[0]
	// 3. iterate through, check if equals prev
	for i := 1; i < len(nums); i++ {
		if prev == nums[i] {
			return true
		}
		prev = nums[i]
	}
	return false
}
