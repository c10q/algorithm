class Solution:
    def maximumUniqueSubarray(self, nums: List[int]) -> int:
        start_indexes = [-1 for _ in range(10001)]
        sums = [0 for _ in range(100001)]        
        
        start = 0
        max_subarray = 0
        for i in range(len(nums)):
            sums[i + 1] = sums[i] + nums[i]
            
            if start_indexes[nums[i]] != -1:
                max_subarray = max(max_subarray, sums[i] - sums[start])
                start = max(start, start_indexes[nums[i]] + 1)
            start_indexes[nums[i]] = i
            
        max_subarray = max(max_subarray, sum(nums[start:]))
        
        return max_subarray
