class Solution:
    def countGood(self, nums: List[int], k: int) -> int:
        start, end = 0, 0
        l = len(nums)
        pairCnt = 0
        cntMap = {}
        ans = 0

        while start <= end:
            while pairCnt >= k:
                cntMap[nums[start]] -= 1
                if cntMap[nums[start]] >= 1:
                    pairCnt -= cntMap[nums[start]]
                start += 1
                if pairCnt >= k:
                    ans += l - end + 1

            while end < l and pairCnt < k:
                if nums[end] not in cntMap:
                    cntMap[nums[end]] = 0
                if cntMap[nums[end]] >= 1:
                    pairCnt += cntMap[nums[end]]
                cntMap[nums[end]] += 1
                end += 1

            if pairCnt < k:
                return ans

            ans += l - end + 1
        return ans
