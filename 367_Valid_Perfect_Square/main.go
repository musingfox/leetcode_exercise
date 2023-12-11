func isPerfectSquare(num int) bool {
	left := 1
	right := num

	for {
		if left+1 == right {
			return false
		}
		sqrt := int((left + right) / 2)
		square := sqrt * sqrt
		if square > num {
			right = sqrt
		} else if square < num {
			left = sqrt
		} else {
			return true
		}
	}
}
