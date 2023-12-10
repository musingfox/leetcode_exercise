func countSubarrays(nums []int, k int) int64 {
	left, right := 0, 0
	l := len(nums)
	cnt := 0
	ans := 0
	maxNum := 0
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}

	for left <= right {
		for cnt >= k {
			if nums[left] == maxNum {
				cnt--
			}
			left++
			if cnt >= k {
				ans += l - right + 1
			}
		}
		for right < l && cnt < k {
			if nums[right] == maxNum {
				cnt++
			}
			right++
		}

		if cnt < k {
			return int64(ans)
		}
		ans += l - right + 1
	}

	return int64(ans)
}
