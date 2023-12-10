class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        left, right = 0, 0
        l = len(nums)
        cnt = 0
        ans = 0
        maxNum = max(nums)

        total = nums.count(maxNum)
        if total < k:
            return 0

        while left <= right:
            while cnt >= k:
                if nums[left] == maxNum:
                    cnt -= 1
                left += 1
                if cnt >= k:
                    ans += l - right + 1

            while cnt < k and right < l:
                if nums[right] == maxNum:
                    cnt += 1
                right += 1

            if cnt < k:
                return ans

            ans += l - right + 1

        return ans
