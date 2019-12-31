package backtracking

import (
	"fmt"
	"strconv"
)

func printBinary(digits int) {
	printBinaryHelper(digits, "")
}

func printBinaryHelper(digits int, output string) {
	if digits == 0 {
		fmt.Println(output)
		// we have to add this return here, because in golang the function won't terminate until a explicit return
		return
	}

	// Obervation: when the set of digit choices available is large, using a loop to enumerate them aviods redundancy.
	// print decimal
	for i := 0; i < 10; i++ {
		printBinaryHelper(digits-1, output+strconv.Itoa(i))
	}
}
