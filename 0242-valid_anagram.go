func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    count := make(map[byte]int)

    for i := range s {
        count[s[i]]++
        count[t[i]]--
    }

    for _, v := range count {
        if v != 0 {
            return false
        }
    }

    return true
} 