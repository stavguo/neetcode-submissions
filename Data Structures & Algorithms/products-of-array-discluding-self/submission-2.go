func productExceptSelf(nums []int) []int {
    // [1, 2, 4, 6]
    // [1, 1, 2, 8] first pass
    // [48, 24, 6, 1] second pass
    // [48, 24, 12, 8] final pass
    products := make([]int, len(nums))
    for i := range products {
        products[i] = 1
    }
    for i := 1; i < len(nums); i++ {
        products[i] = products[i - 1] * nums[i - 1]
    }
    curprod := nums[len(nums) - 1]
    for i := len(nums) - 2; i >= 0; i-- {
        products[i] = products[i] * curprod
        curprod *= nums[i]
    }
    return products
}
