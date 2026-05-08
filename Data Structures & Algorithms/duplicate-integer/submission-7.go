func hasDuplicate(nums []int) bool {
    m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		_, found := m[nums[i]]
		if found {
			return true
		} else {
			m[nums[i]] = true
		}
	}
	return false
}
