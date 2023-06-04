package main

func firstMissingPositive(nums []int) int {
	max := len(nums)
	slice := make([]bool, len(nums)+1)
	hasNonPositive := false

	for _, v := range nums {
		if v < 1 || v > max {
			hasNonPositive = true
			continue
		}
		slice[v] = true
	}

	for i := 1; i <= max; i++ {
		if !slice[i] {
			return i
		}
	}
	if hasNonPositive {
		return max
	}

	return max + 1
}
