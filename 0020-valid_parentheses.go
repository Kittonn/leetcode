func isValid(s string) bool {
    mapParens := map[rune]rune {
        ')': '(',
        ']': '[',
        '}': '{',
    }

    var stack []rune

    for _, paren := range s {
        if open, exists := mapParens[paren]; exists {
            if len(stack) == 0 || stack[len(stack)-1] != open {
                return false
            }

            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, paren)
        }
    }

    return len(stack) == 0
}