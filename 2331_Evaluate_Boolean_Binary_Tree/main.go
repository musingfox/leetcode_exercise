/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
	return dfs(root)
}

func dfs(node *TreeNode) bool {
	if node.Val == 2 {
		return dfs(node.Left) || dfs(node.Right)
	}
	if node.Val == 3 {
		return dfs(node.Left) && dfs(node.Right)
	}

	return node.Val == 1
}
