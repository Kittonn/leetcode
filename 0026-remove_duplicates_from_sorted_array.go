func removeDuplicates(nums []int) int {
	count := 0
	previous := -101

	for _, num := range nums {
		nums[count] = num
		if num != previous {
			count++
			previous = num
		}
	}

	return count
}