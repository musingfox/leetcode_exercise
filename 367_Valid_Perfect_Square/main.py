class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        return math.sqrt(num) * 10 % 10 == 0
