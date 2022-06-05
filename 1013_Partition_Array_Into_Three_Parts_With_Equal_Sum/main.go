package main

func canThreePartsEqualSum(arr []int) bool {
	if len(arr) < 4 {
		return false
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	exceptPartialSum := sum / 3
	cnt := 0
	partialSum := 0
	for i := 0; i < len(arr); i++ {
		partialSum += arr[i]
		if partialSum == exceptPartialSum*(cnt+1) {
			cnt++
			if cnt == 2 {
				return i+1 < len(arr)
			}
		}
	}

	return false
}
