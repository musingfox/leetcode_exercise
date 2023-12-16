class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        currMax, currMin = 1, 1
        ans = nums[0]

        for num in nums:
            vals = (num, num * currMax, num * currMin)
            currMax, currMin = max(vals), min(vals)

            ans = max(ans, currMax)

        return ans
