function twoSum(nums: number[], target: number): number[] {
  const hashMap = new Map<number, number>();

  for (let i = 0; i < nums.length; i++) {
    const difference = target - nums[i];

    if (hashMap.has(nums[i])) {
      return [hashMap.get(nums[i]) as number, i];
    }

    hashMap.set(difference, i);
  }

  return [];
}