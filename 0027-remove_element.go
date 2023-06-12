func removeElement(nums []int, val int) int {
	num := 0
	for _, v := range nums {
		if v != val {
			nums[num] = v
			num++
		}
	}

	return num
}