class Solution:
    def makeEqual(self, words: List[str]) -> bool:
        chars = ''.join(words)
        charSet = set(chars)

        l = len(words)
        for c in charSet:
            if chars.count(c) % l != 0:
                return False
        return True
