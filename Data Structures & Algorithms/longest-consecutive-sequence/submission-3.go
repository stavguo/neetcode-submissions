func longestConsecutive(nums []int) int {
	mp := make(map[int]int)
	res := 0
	for _, num := range nums {
		if mp[num] == 0 {
			left, right := mp[num - 1], mp[num + 1]
			sum := left + 1 + right
			mp[num - left], mp[num], mp[num + right] = sum, sum, sum
			if res < sum {
				res = sum
			}
		}
	}
	return res
}
