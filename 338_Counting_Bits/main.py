class Solution:
    def countBits(self, n: int) -> List[int]:
        ans = [None] * (n + 1)
        ans[0] = 0

        for i in range(1,n+1):
            ans[i] = ans[i >> 1] + (i % 2)

        return ans
