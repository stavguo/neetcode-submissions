func isAlNum(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	l, r := 0, len(s) - 1
	for l < r {
		for l < r && !isAlNum(s[l]) {
			l++
		}
		for l < r && !isAlNum(s[r]) {
			r--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}
