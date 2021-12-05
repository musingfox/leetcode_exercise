# Definition for a binary tree node.
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def minDiffInBST(self, root: Optional[TreeNode]) -> int:
        self.min = float("inf")
        self.preNodeVal = None
        self.inorder(root)
        return self.min

    def inorder(self, node: Optional[TreeNode]):
        if node is None:
            return None
        self.inorder(node.left)
        self.min = min(self.min, node.val - self.preNodeVal)
        self.preNodeVal = node.val
        self.inorder(node.right)
