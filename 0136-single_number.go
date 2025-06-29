func singleNumber(nums []int) int {
    seen := make(map[int]bool)

    for _, num := range nums {
        if _, exists := seen[num]; exists {
            delete(seen, num)
        } else {
            seen[num] = true
        }
    }

    for key, _ := range seen {
        return key
    }

    return 0
}