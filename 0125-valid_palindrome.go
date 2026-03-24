import "strings"

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	lower := strings.ToLower(s)

	for l < r {
		if isNotAlphanumeric(lower[l]) {
			l++
			continue
		}

		if isNotAlphanumeric(lower[r]) {
			r--
			continue
		}

		if lower[l] != lower[r] {
			return false
		}

		l++
		r--
	}

	return true
}

func isNotAlphanumeric(c byte) bool {
	return !(c >= 'a' && c <= 'z') && !(c >= '0' && c <= '9')
}