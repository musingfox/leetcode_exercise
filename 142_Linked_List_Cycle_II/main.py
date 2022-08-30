# Definition for singly-linked list.

from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        ids = {id(head)}
        curr = head
        while curr and curr.next:
            curr = curr.next
            currId = id(curr)
            if currId in ids:
                return curr
            ids.add(currId)
        return None
