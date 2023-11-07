class Solution:
    def minimumRounds(self, tasks: List[int]) -> int:
        hashmap = Counter(tasks)
        res = 0
        for value in hashmap.values():
            if value <= 1:
                return -1
            res += ceil(value / 3)
        return res
