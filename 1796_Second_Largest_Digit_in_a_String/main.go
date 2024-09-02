import "sort"

func secondHighest(s string) int {
	digits := []int{-1}
	for _, v := range s {
		num, err := strconv.Atoi(string(v))
		if err == nil {
			digits = append(digits, num)
		}
	}

	sort.Ints(digits)
	max := digits[len(digits)-1]
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < max {
			return digits[i]
		}
	}
	return -1
}
