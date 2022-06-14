class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        temp = [[0 for _ in range(len(word2) + 1)] for _ in range(len(word1) + 1)]
        max_size_deleted_word = 0
        for i in range(len(word1)):
            for j in range(len(word2)):
                if word1[i] == word2[j]:
                    temp[i + 1][j + 1] += temp[i][j] + 1
                    max_size_deleted_word = max(max_size_deleted_word, temp[i + 1][j + 1])
                else:
                    temp[i + 1][j + 1] = max(temp[i + 1][j], temp[i][j + 1])
            
        return len(word1) + len(word2) - (max_size_deleted_word * 2)
                
