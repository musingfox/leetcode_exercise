class Solution:
    def countTestedDevices(self, batteryPercentages: List[int]) -> int:
        cnt = 0

        for device in batteryPercentages:
            device -= cnt
            if device > 0:
                cnt += 1

        return cnt
