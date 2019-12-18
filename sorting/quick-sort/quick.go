package quick

func sort(arr []int) []int {
	// If array is empty or consists of one element then return this array since it is sorted.
	if len(arr) <= 1 {
		return arr
	}

	// Init left and right arrays.
	leftArr := []int{}
	rightArr := []int{}

	// Take the first element of array as a pivot.
	pivot, arr := arr[0], arr[1:]
	centerArr := []int{pivot}

	for _, curr := range arr {
		if curr == pivot {
			centerArr = append(centerArr, curr)
		} else if curr < pivot {
			leftArr = append(leftArr, curr)
		} else {
			rightArr = append(rightArr, curr)
		}
	}

	leftSortedArr := sort(leftArr)
	rightSortedArr := sort(rightArr)

	sortedArr := []int{}
	sortedArr = append(leftSortedArr, centerArr...)
	sortedArr = append(sortedArr, rightSortedArr...)

	return sortedArr
}
