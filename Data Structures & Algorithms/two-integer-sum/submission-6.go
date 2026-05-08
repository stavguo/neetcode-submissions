func twoSum(nums []int, target int) []int {
    diffs := make(map[int]int)
    for i, num := range nums {
        if idx, ok := diffs[target - num]; ok {
            return []int{idx,i}
        }
        diffs[num] = i
    }
    return nil
}
