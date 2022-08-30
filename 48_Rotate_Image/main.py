from typing import List


class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        for i in range(n // 2):
            # save top edge to temp
            temp = matrix[i][i:n - i - 1]
            temp.reverse()
            for j in range(i, n - i - 1):
                matrix[i][j] = matrix[n - 1 - j][i]
                matrix[n - 1 - j][i] = matrix[n - i - 1][n - 1 - j]
                matrix[n - i - 1][n - 1 - j] = matrix[j][n - i - 1]
                matrix[j][n - i - 1] = temp.pop()
