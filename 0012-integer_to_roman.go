type Rule struct {
	value int
	rule  string
}

func intToRoman(num int) string {
	rules := []Rule{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""

	for _, rule := range rules {
		for num >= rule.value {
			roman += rule.rule
			num -= rule.value
		}
	}

	return roman
}