func maxProduct(nums []int) int {
	currMax, currMin := 1, 1
	ans := nums[0]

	for _, num := range nums {
		vals := []int{num, num * currMax, num * currMin}
		currMax, currMin = getMaxAndMin(vals)

		if currMax > ans {
			ans = currMax
		}
	}

	return ans
}

func getMaxAndMin(vals []int) (int, int) {
	max := vals[0]
	min := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		} else if val < min {
			min = val
		}
	}

	return max, min
}
