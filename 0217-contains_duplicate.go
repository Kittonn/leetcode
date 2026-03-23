func containsDuplicate(nums []int) bool {
    seen := make(map[int]int)

    for _, num := range nums {
        if _, exist := seen[num]; exist {
            return true
        }

        seen[num] = 1
    }

    return false
}