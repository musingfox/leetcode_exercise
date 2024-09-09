func maxProduct(words []string) int {
	n := len(words)
	bitmasks := make([]int, n)
	lengths := make([]int, n)

	for i := 0; i < n; i++ {
		mask := 0
		for _, c := range words[i] {
			mask |= 1 << (c - 'a')
		}
		bitmasks[i] = mask
		lengths[i] = len(words[i])
	}

	maxProduct := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if bitmasks[i]&bitmasks[j] == 0 {
				product := lengths[i] * lengths[j]
				if product > maxProduct {
					maxProduct = product
				}
			}
		}
	}

	return maxProduct
}
