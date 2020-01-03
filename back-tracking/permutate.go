package backtracking

import "fmt"

func permutate(input []string) [][]string {
	result := [][]string{}
	chosen := []string{}

	var recursive func(input []string)

	recursive = func(input []string) {
		fmt.Println("input ", input, "chosen ", chosen)
		// base case
		if len(input) == 0 {
			// fmt.Println(chosen)
			result = append(result, append([]string{}, chosen...))
			return
		}

		for i := 0; i < len(input); i++ {
			// chose
			s := input[i]
			chosen = append(chosen, s)
			// explore
			input = append(input[:i], input[i+1:]...)
			fmt.Println("I choose ", s, "now input is ", input)
			recursive(input)

			fmt.Println("returned")
			fmt.Println()
			// unchose
			chosen = chosen[:len(chosen)-1]
			// input = append(input[:i], append([]string{s}, input[i:]...)...)
			input = append(input, "")
			copy(input[i+1:], input[i:])
			input[i] = s
			fmt.Println("I choose ", s, "now input is ", input)
		}

		return
	}

	recursive(input)

	for _, v := range result {
		fmt.Println(v)
	}
	return result
}
