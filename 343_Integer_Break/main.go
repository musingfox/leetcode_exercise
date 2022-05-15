package main

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	result := 1

	for {
		if n <= 4 {
			break
		}
		result *= 3
		n -= 3
	}

	return result * n
}
