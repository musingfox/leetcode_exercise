func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i]+currSum {
			currSum = nums[i]
		} else {
			currSum += nums[i]
		}

		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}
