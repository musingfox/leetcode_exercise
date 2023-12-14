class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        map = {}
        for num in nums:
            if num in map:
                return True
            map[num] = True

        return False
