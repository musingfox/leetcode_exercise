class Solution:
    def nextBeautifulNumber(self, n: int) -> int:
        while True:
            flag = True
            n = n + 1
            nString = str(n)
            nSet = set()
            for x in nString:
                if (x in nSet):
                    continue
                nSet.add(x)
                flag = flag and int(x) == nString.count(x)
                if flag is not True:
                    break
            if flag:
                return n
