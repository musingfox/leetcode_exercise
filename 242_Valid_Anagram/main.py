class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        record = [0] * 26
        for char in s:
            ctoi = ord(char) - 97
            record[ctoi] += 1

        for char in t:
            ctoi = ord(char) - 97
            record[ctoi] -= 1
            if record[ctoi] < 0:
                return False

        return True
