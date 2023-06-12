func lengthOfLastWord(s string) int {
	lastWordLength := 0
	for i := len(s) - 1; i >= 0; i-- {
		currentChar := s[i]
		if currentChar == ' ' && lastWordLength > 0 {
			break
		} else if currentChar != ' ' {
			lastWordLength++
		}
	}

	return lastWordLength
}