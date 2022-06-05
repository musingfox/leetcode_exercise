class Solution:
    def canThreePartsEqualSum(self, arr: list[int]) -> bool:
        if sum(arr) % 3 != 0 or len(arr) < 3:
            return False
        exceptPartialSum = sum(arr) // 3
        partialSum = 0
        cnt = 0
        for i in range(0, len(arr)):
            partialSum += arr[i]
            if partialSum == exceptPartialSum * (cnt + 1):
                cnt += 1
                if cnt == 2:
                    return i + 1 < len(arr) and sum(arr[i + 1:]) == exceptPartialSum

        return False
