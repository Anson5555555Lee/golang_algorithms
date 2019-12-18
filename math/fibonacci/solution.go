package fibonacci

func fibonacci(i int) []int {
	fibSequence := []int{1}

	prevVal := 0
	currVal := 1

	if i == 1 {
		return fibSequence
	}

	counter := i - 1
	for counter > 0 {
		currVal = currVal + prevVal
		prevVal = currVal - prevVal

		fibSequence = append(fibSequence, currVal)

		counter--
	}

	return fibSequence
}

// An recursive implementation with O(2^N) time complexity
// func fibonacciNth(i int) int {
// 	if i == 0 {
// 		return 0
// 	}

// 	if i == 1 {
// 		return 1
// 	}

// 	return fibonacciNth(i-1) + fibonacciNth(i-2)
// }

// An DP implementation
func fibonacciNth(i int, memo []int) int {
	// Base case
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	if memo[i] == 0 {
		memo[i] = fibonacciNth(i-1, memo) + fibonacciNth(i-2, memo)
	}

	return memo[i]
}
