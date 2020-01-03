package merge

// mergesort performs O(N) operations on each level (width)
// each level splits the data in 2, so there are log2N levels (height)
// product of these = N*log2N = O(Nlog2N)
// for example: N = 32. Performs ~log2(32) = 5 levels of N operations each

func sort(arr []int) []int {
	// base case
	// If array is empty or consists of one element then return this array since it is sorted.
	if len(arr) <= 1 {
		return arr
	}

	// Split array on two halves.
	middleIndex := len(arr) / 2
	leftArr := arr[0:middleIndex]
	rightArr := arr[middleIndex:]

	// Sort two halves of split array.
	leftSortedArr := sort(leftArr)
	rightSortedArr := sort(rightArr)

	// Merge two sorted array into one.
	return mergeSortedArrays(leftSortedArr, rightSortedArr)
}

func mergeSortedArrays(left, right []int) []int {
	sortedArray := []int{}

	for len(left) > 0 && len(right) > 0 {
		// Find minimum element of two arrays.
		var minimumElement int
		if left[0] < right[0] {
			minimumElement, left = left[0], left[1:]
		} else {
			minimumElement, right = right[0], right[1:]
		}

		// Push the minimum element of two arrays to the sorted array.
		sortedArray = append(sortedArray, minimumElement)
	}

	if len(left) > 0 {
		sortedArray = append(sortedArray, left...)
	}

	if len(right) > 0 {
		sortedArray = append(sortedArray, right...)
	}

	return sortedArray
}
