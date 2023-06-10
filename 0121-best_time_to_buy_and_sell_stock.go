func maxProfit(prices []int) int {
	profit, lowest := 0, prices[0]

	for _, price := range prices {
		if price-lowest > profit {
			profit = price - lowest
		}

		if price < lowest {
			lowest = price
		}

	}

	return profit
}