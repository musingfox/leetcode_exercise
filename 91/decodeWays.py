class Solution:
    alreadyCal = {}

    def numDecodings(self, s: str) -> int:
        if len(s) == 0 or s[0] == '0':
            return 0
        return self.dp(s)

    def dp(self, s) -> int:
        if s == '' or s[0] == '0':
            return 0
        if len(s) == 1:
            return 1
        if int(s[0]) >= 3:
            return self.getCalculate(s[1:])
        if int(s[0:2]) > 26:
            return self.getCalculate(s[1:])
        if int(s) < 27:
            if s[1] == '0':
                return 1
            return 2

        return self.getCalculate(s[1:]) + self.getCalculate(s[2:])

    def getCalculate(self, s) -> int:
        if (self.alreadyCal.get(s) is None):
            self.alreadyCal[s] = self.dp(s)
        return self.alreadyCal[s]
