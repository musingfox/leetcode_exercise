func minimumRounds(tasks []int) int {
	ans := 0

	var unique []int
	record := make(map[int]int)
	for i := 0; i < len(tasks); i++ {
		if _, ok := record[tasks[i]]; !ok {
			unique = append(unique, tasks[i])
		}
		record[tasks[i]]++
	}

	for i := 0; i < len(unique); i++ {
		if record[unique[i]] == 1 {
			return -1
		}
		r := record[unique[i]] % 3
		if r == 0 {
			ans += record[unique[i]] / 3
		} else {
			ans += int(record[unique[i]]/3) + 1
		}
	}

	return ans
}
