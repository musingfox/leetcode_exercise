func countTestedDevices(batteryPercentages []int) int {
	cnt := 0
	for _, device := range batteryPercentages {
		if device-cnt > 0 {
			cnt++
		}
		fmt.Println(device)
	}

	return cnt
}
