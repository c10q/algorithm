class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        lastIndex = {}
        startIndex = 0
        result = 0
        for i in range(len(s)):
            c = s[i]
            if c in lastIndex and lastIndex[c] >= startIndex:
                startIndex = lastIndex[c] + 1
            else:
                result = max(result, i - startIndex + 1)
        
            lastIndex[c] = i
        
        return result
        
