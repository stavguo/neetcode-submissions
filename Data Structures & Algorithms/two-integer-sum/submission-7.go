func twoSum(nums []int, target int) []int {
    diffs := make(map[int]int)
    for i, num := range nums {
        if j, ok := diffs[target - num]; ok {
            return []int{j,i}
        }
        diffs[num] = i
    }
    return []int{}
}
