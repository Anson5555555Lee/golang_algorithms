package backtracking

import "fmt"

// Write a function sublists that finds every possible sub-list of a given vector.
// A sub-list of a vector V contains >= 0 of V's elements.
// What are "choices" in this problem?

func sublists(space []string) [][]string {
	result := [][]string{}
	chosen := []string{}

	var recursion func()

	recursion = func() {
		// base case
		if len(space) == 0 {
			temp := make([]string, len(chosen))
			copy(temp, chosen)
			result = append(result, temp)

			return
		}
		// recursive case
		// for each possible choice:
		s, space := space[0], space[1:]
		// - choose (try without s, ... try with s)

		// - explore
		recursion()
		chosen = append(chosen, s)
		recursion()

		// - un-choose
		space = append([]string{s}, space...)
		chosen = chosen[:len(chosen)-1]
	}

	recursion()
	for _, v := range result {
		fmt.Println(v)
	}
	return result
}
