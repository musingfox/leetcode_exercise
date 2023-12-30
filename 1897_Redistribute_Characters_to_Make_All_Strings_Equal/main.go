func makeEqual(words []string) bool {
	m := make(map[rune]int)

	for _, word := range words {
		for _, char := range word {
			if _, ok := m[char]; !ok {
				m[char] = 0
			}
			m[char]++
		}
	}

	l := len(words)
	for _, v := range m {
		if v%l != 0 {
			return false
		}
	}
	return true
}
