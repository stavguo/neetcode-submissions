func threeSum(nums []int) [][]int {
    res := [][]int{}
    // sort
    sort.Ints(nums)
    // for each num
    for i, a := range nums {
        if a > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        // left pointer, right pointer
        l, r := i + 1, len(nums) - 1
        // for l < r 
        for l < r {
            threeSum := a + nums[l] + nums[r]
            // if l + r > curNum, decrease r
            if threeSum > 0 {
                r--
            } else if threeSum < 0 {
                l++
            } else {
                res = append(res, []int{a, nums[l], nums[r]})
                l++
                r--
                for l < r && nums[l] == nums[l - 1] {
                    l++
                }
            }
        }
    }
    return res
}
