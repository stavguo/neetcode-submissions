func lengthOfLongestSubstring(s string) int {
    mp := make(map[byte]int)
    l, maxLen := 0, 0
    for r := 0; r < len(s); r++ {
        // check if exists, set l
        if idx, found := mp[s[r]]; found {
            l = max(idx + 1, l)
        }
        // add to char's last seen pos to mp
        mp[s[r]] = r
        // update maxLen
        if r - l + 1 > maxLen {
            maxLen = r - l + 1
        }
    }
    return maxLen
}
