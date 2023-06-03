class Solution:
  def missingNumber(self, nums: list[int]) -> int:
    n = len(nums)
    sum_nums = (n * (n + 1)) // 2
    return sum_nums - sum(nums)