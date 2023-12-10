func getGoodIndices(variables [][]int, target int) []int {
	var ans []int

	for i, vars := range variables {
		a, b, c, m := int64(vars[0]), int64(vars[1]), int64(vars[2]), int64(vars[3])
		a = powAndMod(a, b, int64(10))
		a = powAndMod(a, c, m)
		if a == int64(target) {
			ans = append(ans, i)
		}
	}

	return ans
}

func powAndMod(base int64, exponent int64, mod int64) int64 {
	result := int64(1)
	for i := int64(0); i < exponent; i++ {
		result = result * base % mod
	}
	return result
}
