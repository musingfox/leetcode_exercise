func integerReplacement(n int) int {
	cnt := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			if n == 3 || n%4 == 1 {
				n -= 1
			} else {
				n += 1
			}
		}
		cnt++
	}

	return cnt
}
