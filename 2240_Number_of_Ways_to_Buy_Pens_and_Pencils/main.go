package main

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var ans int64
	var expensive int
	var cheap int
	if cost1 > cost2 {
		expensive = cost1
		cheap = cost2
	} else {
		expensive = cost2
		cheap = cost1
	}

	for total >= 0 {
		ans += int64(total/cheap) + 1
		total -= expensive
	}

	return ans
}
