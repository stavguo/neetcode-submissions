func characterReplacement(s string, k int) int {
    count := make(map[byte]int)
    res, l, maxf := 0, 0, 0
    for r := 0; r < len(s); r++ {
        // 1. take count of new char, r
        count[s[r]]++
        // 2. determine what is the max char
        if count[s[r]] > maxf {
            maxf = count[s[r]]
        }
        // 3. If window len - maxf > k, window is too big so decrement count[s[l]]
        // and move l forwards
        if r - l + 1 - maxf > k {
            count[s[l]]--
            l++
        }
        // 4. update res using window len
        if r - l + 1 > res {
            res = r - l + 1
        }
    }
    return res
}
