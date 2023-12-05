func imageSmoother(img [][]int) [][]int {
	w := len(img[0])
	h := len(img)

	ans := make([][]int, h)
	for i := range ans {
		ans[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			cellCnt := 0
			cellSum := 0
			for k := -1; k < 2; k++ {
				for l := -1; l < 2; l++ {
					if i+k >= 0 && i+k < h && j+l >= 0 && j+l < w {
						cellCnt++
						cellSum += img[i+k][j+l]
					}
				}
			}
			ans[i][j] = int(cellSum / cellCnt)
		}
	}

	return ans
}
