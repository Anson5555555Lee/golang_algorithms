package problem0070

func bruteForce(n int) int {
	if n < 0 {
		return 0
	}

	// Base case.
	if n == 0 {
		return 1
	}

	return bruteForce(n-1) + bruteForce(n-2)
}

func memoization(n int, memo map[int]int) int {
	if n < 0 {
		return 0
	}

	// Base case.
	if n == 0 {
		return 1
	}

	if memo[n] == 0 {
		memo[n] = memoization(n-1, memo) + memoization(n-2, memo)
	}

	return memo[n]
}

func tripleSteps(n int, memo map[int]int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	if memo[n] == 0 {
		memo[n] = tripleSteps(n-1, memo) + tripleSteps(n-2, memo) + tripleSteps(n-3, memo)
	}

	return memo[n]
}
