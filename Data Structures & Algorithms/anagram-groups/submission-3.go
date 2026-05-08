func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, s := range strs {
		runes := []rune(s)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedString := string(runes)
		anagrams[sortedString] = append(anagrams[sortedString], s)
	}
	output := make([][]string, 0, len(anagrams))
	for _, anagram := range anagrams {
		output = append(output, anagram)
	}
	return output
}
