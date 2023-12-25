func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := int((left + right) / 2)
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[left]
}
