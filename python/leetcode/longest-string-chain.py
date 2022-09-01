from collections import deque

class Solution:
    def longestStrChain(self, words: List[str]) -> int:
        temp = {}
        words.sort(key=len)
        
        for i in words:
            temp[i] = 1
            for j in range(len(i)):
                chr_deleted = i[:j] + i[j+1:]
                if chr_deleted in temp:
                    temp[ i ] = max(temp[i], temp[chr_deleted] + 1)
        
        return max(temp.values())
    
