func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int)

	for _, num := range nums {
		hashmap[num]++
	}

	keys := make([]int, 0, len(hashmap))
	for k := range hashmap {
		keys = append(keys, k)
	}

	slices.SortFunc(keys, func(a, b int) int {
		return hashmap[b] - hashmap[a]
	})

	return keys[:k]
}