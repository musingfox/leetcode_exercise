class Solution:
    def waysToBuyPensPencils(self, total: int, cost1: int, cost2: int) -> int:
        ans = 0
        expensive = max(cost1, cost2)
        cheap = min()
        if cost1 > cost2:
            expensive = cost1
            cheap = cost2
        else:
            expensive = cost2
            cheap = cost1

        while total >= 0:
            ans += total // cheap + 1
            total -= expensive
        return ans
