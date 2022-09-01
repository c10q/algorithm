class Solution:
    def longestPalindrome(self, s: str) -> str:
        max_left = 0
        max_right = 0
        
        if len(s) == 1:
            return s
        if len(s) == 2:
            if s[0] == s[1]:
                return s
            else:
                return s[0]
        
        for i in range(1, len(s) - 1):
            print('i', i)
            char = s[i]
            depth = 0
            while depth < i and depth < len(s) - i - 1:
                depth += 1
                left = i - depth
                right = i + depth
                if s[left] == s[right]:
                    if right - left > max_right - max_left:
                        print(right, left)
                        max_right = right
                        max_left = left
                else:
                    break
                        
            depth = 0        
            while depth < i and depth < len(s) - i:
                depth += 1
                left = i - depth
                right = i + depth  - 1
                if s[left] == s[right]:
                    if (right) - left > max_right - max_left:
                        max_left = left
                        max_right = right
                else:
                    break
                    
            depth = 0   
            while depth < i and depth < len(s) - i - 1:          
                depth += 1
                left = i - depth + 1
                right = i + depth
                if s[left] == s[right]:
                    if right - left > max_right - max_left:
                        max_left = left
                        max_right = right
                else:
                    break
                        
        return s[max_left: max_right + 1]
            
            
