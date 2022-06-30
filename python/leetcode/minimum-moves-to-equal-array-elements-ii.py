class Solution:
    def minMoves2(self, nums: List[int]) -> int:
        result = 0
        nums.sort()
        middle_idx = len(nums) // 2

        for num in nums:
            result += abs(num - nums[middle_idx])
            
        return result
