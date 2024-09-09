# The guess API is already defined for you.
# @param num, your guess
# @return -1 if num is higher than the picked number
#          1 if num is lower than the picked number
#          otherwise return 0
# def guess(num: int) -> int:

class Solution:
    def guessNumber(self, n: int) -> int:
        floor = 1
        ceil = n
        while True:
            ans = (floor + ceil) // 2
            g = guess(ans)
            if g == 1:
                floor = ans + 1
            elif g == -1:
                ceil = ans
            else:
                return ans
