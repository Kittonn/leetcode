func romanToInt(s string) int {
	num := 0

	rules := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	previous := 0

	for i := 0; i < len(s); i++ {
		current := rules[s[i]]
		num += current
		if previous < current {
			num -= previous * 2
		}
		previous = current
	}

	return num
}