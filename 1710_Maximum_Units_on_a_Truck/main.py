from operator import itemgetter


class Solution:
    def maximumUnits(self, boxTypes: list[list[int]], truckSize: int) -> int:
        boxTypes.sort(key=itemgetter(1), reverse=True)
        result = 0
        boxTypeLen = len(boxTypes)

        index = 0
        while truckSize and index < boxTypeLen:
            boxCount, unit = boxTypes[index]
            boxCount = min(boxCount, truckSize)
            result += boxCount * unit
            truckSize -= boxCount
            index += 1

        return result
