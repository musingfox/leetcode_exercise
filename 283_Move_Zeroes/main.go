package main

func moveZeroes(nums []int) {
	lastZeroIndex := 0

	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[lastZeroIndex] = nums[lastZeroIndex], nums[i]
			lastZeroIndex++
		}
	}
}
