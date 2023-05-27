from typing import List

class Solution:
    def kthSmallestPrimeFraction(self, arr, k):
        def search(x):
            count = 0
            windowLeft = arr[0]
            windowRight = arr[-1]
            for
        left = 0
        right = arr[0] / arr[-1]
        x = (left + right) / 2
        count = search(x)
        if count == k:
            return [left, right]
        elif count < k:
            left = x
        else:
            right = x


sol = Solution()
print(sol.kthSmallestPrimeFraction([1,2,3,5], 3))
print(sol.kthSmallestPrimeFraction([1,7], 1))
