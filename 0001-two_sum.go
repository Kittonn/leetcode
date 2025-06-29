func twoSum(nums []int, target int) []int {
	var tempMap = make(map[int]int)
	for idx, nums := range nums {
		diff := target - nums

		value, exists := tempMap[nums]
		if exists {
            return []int{value, idx}
		}

		tempMap[diff] = idx
	}

  return []int{}
}