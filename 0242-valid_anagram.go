func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sSeen := make(map[byte]int)
    tSeen := make(map[byte]int)

    for i := range len(s) {
        sSeen[s[i]]++
        tSeen[t[i]]++
    }

    for i, v := range sSeen {
        if tSeen[i] != v {
            return false
        }
    }

    return true
} 