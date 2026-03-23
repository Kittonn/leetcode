func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if j, ok := seen[diff]; ok {
			return []int{j, i}
		}

		seen[num] = i
	}

	return []int{}
}