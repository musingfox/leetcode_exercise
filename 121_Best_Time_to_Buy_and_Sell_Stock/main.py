class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        low = prices[0]
        profit = 0

        for price in prices:
            if price < low:
                low = price
            if price - low > profit:
                profit = price - low

        return profit
