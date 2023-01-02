from functools import cache
import math

class Solution:
    def leastOpsExpressTarget(self, x: int, target: int) -> int:

        @cache
        def recursion(t):
            if t < x:
                # add from 0 to t, or sub from x to t
                return min(2*t-1, 2*(x-t))
            log = math.log(t, x)
            left = int(log)
            right = left + 1
            opers = left + recursion(t - x**left)
            if x**right - t < t:
                opers = min(opers, right + recursion(x**right - t))

            return opers

        return recursion(target)
