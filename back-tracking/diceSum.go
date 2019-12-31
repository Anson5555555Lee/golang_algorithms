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

func diceSum(dice, sum int) {
	diceSumHelper(dice, sum, []int{})
}

func diceSumHelper(dice, sum int, chosen []int, result [][]int) {
	if dice == 0 {
		// base case

		return
	} else {
		// I will handle 1 die.
		// Try all possible values (1-6) to see if they can work
		for i := 1; i <= 6; i++ {
			// "choose" i
			chosen := append(chosse)
			// "explore" what could follow that
			diceSumHelper(dice-1, sum-i, chosen, result)

			// "un-choose" i
		}
	}
}
