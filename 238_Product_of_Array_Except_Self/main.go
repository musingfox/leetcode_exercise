func productExceptSelf(nums []int) []int {
	l := len(nums)
	ans := []int{1}

	for i := 1; i < l; i++ {
		ans = append(ans, ans[i-1]*nums[i-1])
	}
	fmt.Println(ans)
	postfixProduct := 1
	for i := l - 1; i >= 0; i-- {
		ans[i] *= postfixProduct
		postfixProduct *= nums[i]
	}

	return ans
}
