class Solution:
    def minImpossibleOR(self, nums: List[int]) -> int:
        if 1 not in nums:
            return 1

        power = 1
        while pow(2, power) in nums:
            power += 1

        return pow(2, power)
