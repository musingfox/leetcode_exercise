class Solution:
    def checkIfCanBreak(self, s1: str, s2: str) -> bool:
        a = "".join(sorted(s1))
        b = "".join(sorted(s2))
        smallString = min(a, b)
        bigString = max(a, b)

        for i in range(len(s1)):
            if (smallString[i] > bigString[i]):
                return False

        return True


ans = Solution()
assert ans.checkIfCanBreak('abc', 'xya') is True, "Should be True"
assert ans.checkIfCanBreak('abe', 'acd') is False, "Should be False"
assert ans.checkIfCanBreak('leetcodee', 'interview') is True, "Should be True"
