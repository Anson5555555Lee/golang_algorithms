package problem0121

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice, maxProfit := prices[0], 0
	for _, p := range prices {
		if p-minPrice >= maxProfit {
			maxProfit = p - minPrice
		}

		if p <= minPrice {
			minPrice = p
		}
	}

	return maxProfit
}
