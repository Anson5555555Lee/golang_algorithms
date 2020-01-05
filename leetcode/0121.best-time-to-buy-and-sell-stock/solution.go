package problem0121

// Definition:
// max_profit = max{price[j] - price[i]}, 0 <= i < j <= n-1

// Solution 1: Brute force
// Enumerate all pairs of (i, j)
// Time complexity: O(n^2)

// Solution 2: DP
// Keep tracking the min price so far

// Solution 3: DP

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	// we can also use array to trace the minPrice and maxProfit to show this is a dynamic programming
	minPrice, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i-1])
		maxProfit = max(maxProfit, prices[i]-minPrice)
	}

	return maxProfit
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
