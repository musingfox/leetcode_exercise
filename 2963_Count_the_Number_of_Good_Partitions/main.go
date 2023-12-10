func numberOfGoodPartitions(nums []int) int {
	last := make(map[int]int)

	for k, v := range nums {
		last[v] = k
	}

	f := -1
	p := -1

	for k, v := range nums {
		if f < k {
			p++
		}
		f = max(last[v], f)
	}

	return int(powAndMod(2, p, 1000000007))
}

func powAndMod(base int, exp int, mod int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
		result %= mod
	}

	return result
}
