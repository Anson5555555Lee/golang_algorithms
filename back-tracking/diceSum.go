package backtracking

// Write a function diceSum that accepts two integer parameters: a number of dice to roll, and a desired sum of all die values.
// Output all combinations of die values that add up to exactly that sum.
// e.g.
// diceSum(2, 7)
// [1,6]
// [2,5]
// [3,4]
// [4,3]
// [5,2]
// [6,1]

// func diceSum(dice, sum int) [][]int {
// 	result := &[][]int{}

// 	diceSumHelper(dice, sum, &[]int{}, result)

// 	return *result
// }

// func diceSumHelper(dice, sum int, chosen *[]int, result *[][]int) {
// 	if dice == 0 {
// 		// base case
// 		if sum == 0 {
// 			(*result) = append(*result, *chosen)
// 		}
// 		return
// 	}

// 	// I will handle 1 die.
// 	// Try all possible values (1-6) to see if they can work
// 	for i := 1; i <= 6; i++ {
// 		// "choose" i
// 		(*chosen) = append(*chosen, i)

// 		// "explore" what could follow that
// 		diceSumHelper(dice-1, sum-i, chosen, result)

// 		// "un-choose" i
// 		(*chosen) = (*chosen)[0 : len(*(chosen))-1]
// 	}

// }

func diceSum(dice, sum int) [][]int {
	result := [][]int{}

	var diceSumHelper = func(dice, sum int, chosen []int, result [][]int) {
		if dice == 0 {
			// base case
			if sum == 0 {
				result = append(result, chosen)
			}
			return
		}

		// I will handle 1 die.
		// Try all possible values (1-6) to see if they can work
		for i := 1; i <= 6; i++ {
			// "choose" i
			chosen = append(chosen, i)

			// "explore" what could follow that
			diceSumHelper(dice-1, sum-i, chosen, result)

			// "un-choose" i
			chosen = chosen[0 : len(chosen)-1]
		}

	}

	diceSumHelper(dice, sum, []int{}, result)

	return result
}

// func diceSumHelper(dice, sum int, chosen []int, result [][]int) {
// 	if dice == 0 {
// 		// base case
// 		if sum == 0 {
// 			result = append(result, chosen)
// 		}
// 		return
// 	}

// 	// I will handle 1 die.
// 	// Try all possible values (1-6) to see if they can work
// 	for i := 1; i <= 6; i++ {
// 		// "choose" i
// 		chosen = append(chosen, i)

// 		// "explore" what could follow that
// 		diceSumHelper(dice-1, sum-i, chosen, result)

// 		// "un-choose" i
// 		chosen = chosen[0 : len(chosen)-1]
// 	}

// }
