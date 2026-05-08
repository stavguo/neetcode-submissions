func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms := make(map[rune]int)
	mt := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		ms[rune(s[i])] += 1
		mt[rune(t[i])] += 1
	}
	for k, v := range ms {
		if v != mt[k] {
			return false
		}
	}
	return true
}
