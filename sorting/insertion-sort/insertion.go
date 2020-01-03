package insertion

// Insertion sort: orders a list of values by repetitively inserting a particular
// value into a sorted subset of the list
// Runtime: O(N^2)
// generally slightly faster than selection sort for most inputs

func sort(input []int) []int {
	for i := 1; i < len(input); i++ {
		// swap method
		// for j := 0; j < i; j++ {
		// 	if input[i] < input[j] {
		// 		input[j], input[i] = input[i], input[j]
		// 	}
		// }

		// slide elements right to make room for input[i]
		temp := input[i]
		j := i
		for j >= 1 && input[j-1] > temp {
			input[j] = input[j-1]
			j--
		}
		input[j] = temp

	}
	return input
}
