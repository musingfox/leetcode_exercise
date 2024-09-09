/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	floor := 1
	ceil := n
	for {
		ans := int((floor + ceil) / 2)
		g := guess(ans)
		if g == 1 {
			floor = ans + 1
		} else if g == -1 {
			ceil = ans
		} else {
			return ans
		}
	}
}
