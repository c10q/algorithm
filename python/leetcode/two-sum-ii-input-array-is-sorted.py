class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        temp = [-1 for _ in range(0, 2001)]
        temp_duplicated = [0 for _ in range(0, 2001)]
        li = []
        
        l_n = len(numbers)
        for i in range(l_n):
            if temp[numbers[i] + 1000] == -1:
                temp[numbers[i] + 1000] = i + 1
                li.append(numbers[i])
            else:
                if numbers[i - 1] + numbers[i] == target:
                    return [i, i + 1]
        
        l = len(li)
        for i in range(l - 1):
            for j in range(i + 1, l):
                if li[i] + li[j] == target:
                    return [temp[li[i] + 1000], temp[li[j] + 1000]]
                    
