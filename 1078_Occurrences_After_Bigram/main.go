import "strings"

func findOcurrences(text string, first string, second string) []string {
	var ans []string
	textBlock := strings.Split(text, " ")
	for i := 0; i < len(textBlock)-2; i++ {
		if textBlock[i] == first && textBlock[i+1] == second {
			ans = append(ans, textBlock[i+2])
		}
	}

	return ans
}
