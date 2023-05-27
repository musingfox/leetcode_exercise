package main

import "math"

func minImpossibleOR(nums []int) int {
	power := 0
	for {
		n := int(math.Pow(2, float64(power)))
		isExist := false
		for _, num := range nums {
			if num == n {
				isExist = true
				break
			}
		}
		if !isExist {
			return n
		}
		power++
	}
}
