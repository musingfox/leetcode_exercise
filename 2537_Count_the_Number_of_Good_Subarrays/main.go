func countGood(nums []int, k int) int64 {
	begin, end := 0, 0
	pairCnt := 0
	l := len(nums)
	cntMap := make(map[int]int)
	ans := int64(0)

	for begin <= end {
		for pairCnt >= k {
			cntMap[nums[begin]]--
			if cntMap[nums[begin]] >= 1 {
				pairCnt -= cntMap[nums[begin]]
			}
			begin++
			if pairCnt >= k {
				ans += int64(l - end + 1)
			}
		}

		for pairCnt < k && end < l {
			if cntMap[nums[end]] >= 1 {
				pairCnt += cntMap[nums[end]]
			}
			cntMap[nums[end]]++
			end++
		}

		if pairCnt < k {
			return ans
		}
		ans += int64(l - end + 1)
	}

	return ans
}
