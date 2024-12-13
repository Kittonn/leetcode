from typing import List

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        n = len(nums) // 2
        hashmap = {}

        for num in nums:
            if num not in hashmap:
                hashmap[num] = 1
            else:
                hashmap[num] += 1

            if hashmap[num] > n:
                return num