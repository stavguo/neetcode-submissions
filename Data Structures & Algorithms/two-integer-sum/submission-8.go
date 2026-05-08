func twoSum(nums []int, target int) []int {
	subtrahends := make(map[int]int)
    for i, num := range nums {
		s_idx, exists := subtrahends[target - num]
		if exists {
			return []int{s_idx, i}
		}
		subtrahends[num] = i
	}
	return []int{}
}
