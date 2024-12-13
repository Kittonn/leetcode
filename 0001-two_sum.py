from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:    
        hashmap = {}
        for i in range(len(nums)):
          diff = target - nums[i]

          if nums[i] in hashmap:
            return [hashmap[nums[i]], i]

          hashmap[diff] = i

        return []
