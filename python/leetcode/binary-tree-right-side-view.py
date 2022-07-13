from collections import deque

class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []
        result = [root.val]
        root.layer = 1
        queue = deque([root])
        
        while queue:
            tree = queue.popleft()
            if tree.right:
                tree.right.layer = tree.layer + 1
                if len(result) < tree.right.layer:
                    result.append(tree.right.val)
                queue.append(tree.right)
            if tree.left:
                tree.left.layer = tree.layer + 1
                if len(result) < tree.left.layer:
                    result.append(tree.left.val)
                queue.append(tree.left)
                
        return result
