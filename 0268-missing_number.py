from collections import Counter

class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        hashmap = Counter(nums)
        
        for i in range(len(nums) + 1):
            if i not in hashmap:
              return i
            
