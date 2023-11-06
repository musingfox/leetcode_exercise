func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := make(map[rune]int)

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		record[runes[i]]++
	}

	runes = []rune(t)
	for i := 0; i < len(runes); i++ {
		record[runes[i]]--
		if record[runes[i]] < 0 {
			return false
		}
	}

	return true
}
