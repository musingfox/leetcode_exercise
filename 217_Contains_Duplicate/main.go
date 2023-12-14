func containsDuplicate(nums []int) bool {
	mapping := make(map[int]bool)
	for _, num := range nums {
		if _, ok := mapping[num]; ok {
			return true
		}
		mapping[num] = true
	}

	return false
}
