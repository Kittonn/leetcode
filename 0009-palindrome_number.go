func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y, temp := 0, x

	for temp != 0 {
		y = y*10 + temp%10
		temp = temp / 10
	}

	return y == x
}