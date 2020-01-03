package selection

func sort(input []int) {
	for i := 0; i < len(input); i++ {
		min := i
		for j := i + 1; j < len(input); j++ {
			// do the comparison
			if input[j] < input[min] {
				min = j
			}
		}

		// do the swap
		if min != i {
			input[min], input[i] = input[i], input[min]
		}
	}
}
