func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_letters, t_letters := make(map[rune]int), make(map[rune]int)
	for i, ch := range s {
		s_letters[ch]++
		t_letters[rune(t[i])]++
	}
	for k, v := range s_letters {
		if t_letters[k] != v {
			return false
		}
	}
	return true
}
