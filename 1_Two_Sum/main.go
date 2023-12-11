func twoSum(nums []int, target int) []int {
	arr := make(map[int]int)
	ans := []int{}

	for i, num := range nums {
		if _, ok := arr[num]; ok {
			ans = append(ans, i)
			ans = append(ans, arr[num])
			break
		}
		x := target - num
		arr[x] = i
	}

	return ans
}
