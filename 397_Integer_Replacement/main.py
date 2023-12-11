class Solution:
    def integerReplacement(self, n: int) -> int:
        @cache
        def f(x):
            if x==1:
                return 0
            if x%2==0:
                return 1+f(x/2)
            else:
                return 1+min(f(x+1),f(x-1))

        return f(n)
