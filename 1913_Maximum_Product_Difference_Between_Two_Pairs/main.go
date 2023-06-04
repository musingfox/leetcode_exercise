func maxProductDifference(nums []int) int {
	w := 0
	x := 0
	y := 10000
	z := 10000

	for _, v := range nums {
		if v > w {
			x = w
			w = v
		} else if v > x {
			x = v
		}
		if v < z {
			y = z
			z = v
		} else if v < y {
			y = v
		}
	}

	return w*x - y*z
}
