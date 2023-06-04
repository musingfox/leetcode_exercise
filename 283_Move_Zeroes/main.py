class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        lastZeroIndex = 0
        for i, _ in enumerate(nums):
            if nums[i] != 0:
                nums[i], nums[lastZeroIndex] = nums[lastZeroIndex], nums[i]
                lastZeroIndex += 1
