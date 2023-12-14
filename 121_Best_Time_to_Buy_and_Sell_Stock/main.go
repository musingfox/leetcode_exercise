func maxProfit(prices []int) int {
	low := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < low {
			low = prices[i]
		}
		if prices[i]-low > profit {
			profit = prices[i] - low
		}
	}

	return profit
}
