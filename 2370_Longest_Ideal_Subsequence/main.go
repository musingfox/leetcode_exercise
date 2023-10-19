func longestIdealString(s string, k int) int {
	ans := 0
	count := [26]int{}

	for i := 0; i < len(s); i++ {
		charInt := int(s[i] - 'a')
		count[charInt]++
		for j := max(charInt-k, 0); j <= min(25, charInt+k); j++ {
			if j != charInt {
				count[charInt] = max(count[charInt], count[j]+1)
			}
		}
		ans = max(ans, count[charInt])
	}

	return ans
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
