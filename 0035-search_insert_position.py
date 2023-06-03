class Solution:
  def searchInsert(self, nums: list[int], target: int) -> int:
    left, right = 0, len(nums)

    while left < right:
      mid = (left + right) // 2

      if target == nums[mid]:
        return mid
      elif target > nums[mid]:
        left = mid + 1
      else:
        right = mid
    return left
