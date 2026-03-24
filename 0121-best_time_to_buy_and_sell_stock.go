func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}