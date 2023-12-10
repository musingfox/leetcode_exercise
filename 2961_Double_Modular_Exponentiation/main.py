class Solution:
    def getGoodIndices(self, variables: List[List[int]], target: int) -> List[int]:
        result = []
        for index in range(len(variables)):
            vars = variables[index]
            if (vars[0] ** vars[1] % 10) ** vars[2] % vars[3] == target:
                result.append(index)

        return result
