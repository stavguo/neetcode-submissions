func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num] += 1
	}
	buckets := make([][]int, len(nums) + 1)
	for num, val := range count {
		buckets[val] = append(buckets[val], num)
	}
	res := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, num := range buckets[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
