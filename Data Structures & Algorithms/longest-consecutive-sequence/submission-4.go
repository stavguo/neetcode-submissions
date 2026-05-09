func longestConsecutive(nums []int) int {
	mp := make(map[int]int)
	res := 0
	for _, num := range nums {
		if mp[num] == 0 { // if num is new
			left, right := mp[num - 1], mp[num + 1] // get sum data from immediate neighbors
			sum := left + 1 + right // get new total sum
			// setting mp[num - left]: if a seq meets us from the left, it'll work
			// setting mp[num]: if its the first in a seq, it'll initialize right
			// setting mp[num + right]: if a seq meets us from the right, it'll work
			mp[num - left], mp[num], mp[num + right] = sum, sum, sum
			if res < sum {
				res = sum
			}
		}
	}
	return res
}
