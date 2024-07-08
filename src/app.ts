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
console.log(twoSum([2, 7, 11, 15], 9));

console.log(twoSum([3, 2, 4], 6));
