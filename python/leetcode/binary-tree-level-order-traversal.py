# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from collections import deque

class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        result = []
        if not root:
            return []
        
        root.level = 0
        queue = deque([root])
        
        while queue:
            node = queue.popleft()
            level = node.level
            
            if len(result) != level + 1:
                result.append([])
                
            result[node.level].append(node.val)
            
            if node.left:
                node.left.level = node.level + 1
                queue.append(node.left)
                
            if node.right:
                node.right.level = node.level + 1
                queue.append(node.right)
        
        return result
        
