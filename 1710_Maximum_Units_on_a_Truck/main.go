package main

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sortBoxTypes := make([]int, 1001)
	for _, boxType := range boxTypes {
		sortBoxTypes[boxType[1]] += boxType[0]
	}
	result := 0

	for i := 1000; i > 0; i-- {
		if truckSize >= sortBoxTypes[i] {
			result += sortBoxTypes[i] * i
			truckSize -= sortBoxTypes[i]
		} else {
			result += truckSize * i
			break
		}
	}

	return result
}

func main() {
	println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	println(maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
}
