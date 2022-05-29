from inspect import stack


class Solution:
    def smallestSubsequence(self, s: str, k: int, letter: str, repetition: int) -> str:
        stack = []
        letterRequired = repetition
        letterRemain = s.count(letter)
        for i, char in enumerate(s):
            while stack and stack[-1] > char and len(stack) + len(s) - (i + 1) >= k and (stack[-1] != letter or letterRemain > letterRequired):
                if stack.pop() == letter:
                    letterRequired += 1
            if len(stack) < k:
                if char == letter:
                    stack.append(char)
                    letterRequired -= 1
                elif k - len(stack) > letterRequired:
                    stack.append(char)
            if char == letter:
                letterRemain -= 1

        return "".join(stack)


print("{} should be {}".format(Solution().smallestSubsequence("leet", 3, "e", 1), "eet"))
print("{} should be {}".format(Solution().smallestSubsequence("leetcode", 4, "e", 2), "ecde"))
print("{} should be {}".format(Solution().smallestSubsequence("bb", 2, "b", 2), "bb"))
print("{} should be {}".format(Solution().smallestSubsequence("aaabbbcccddd", 3, "b", 2), "abb"))
print("{} should be {}".format(Solution().smallestSubsequence("facfffkfnffoppfffzfz", 9, "f", 9), "fffffffff"))
print("{} should be {}".format(Solution().smallestSubsequence("mmmxmxymmm", 8, "m", 4), "mmmmxmmm"))
