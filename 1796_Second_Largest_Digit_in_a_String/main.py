class Solution:
    def secondHighest(self, s: str) -> int:
        digits = set()
        for i in s:
            if i.isdigit():
                digits.add(int(i))

        digits = sorted(digits)
        if len(digits) > 1:
            return digits[-2]
        return -1
