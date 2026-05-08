func minWindow(s string, t string) string {
    // If t is empty string
    if t == "" {
        return ""
    }
    // Init res, resLen
    res, resLen := []int{-1,-1}, math.MaxInt32
    // Init countT
    countT := make(map[rune]int)
    for _, c := range t {
        countT[c]++
    }
    // Init have, need
    have, need := 0, len(countT)
    // Init window (l, window)
    l, window := 0, make(map[rune]int)
    // For r
    for r := 0; r < len(s); r++ {
        // get letter
        c := rune(s[r])
        // update window
        window[c]++
        // check if have gets incremented
        if countT[c] > 0 && window[c] == countT[c] {
            have++
        }
        // while have == need
        for have == need {
            if r - l + 1 < resLen {
                // update res and resLen
                res, resLen = []int{l,r}, r - l + 1
            }
            // reduce window
            window[rune(s[l])]--
            // check if decrement have
            if countT[rune(s[l])] > 0 && countT[rune(s[l])] > window[rune(s[l])] {
                have--
            }
            // increase l
            l++
        }
    }
    // If res not found, return empty string
    if res[0] == -1 {
        return ""
    }
    // Return res
    return s[res[0]:res[1]+1]
}
