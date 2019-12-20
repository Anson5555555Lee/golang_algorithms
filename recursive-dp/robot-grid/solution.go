package robot

type point struct {
	row    int
	column int
}

func NewPoint(c, r int) *point {
	return &point{c, r}
}

func createMatrix(r, c int) [][]int {
	matrix := make([][]int, r)
	for i := range matrix {
		matrix[i] = make([]int, c)
	}

	// place values (1 indicates position is valid, -1 indicates it is invalid)
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = 1
		}
	}

	return matrix
}

func bruteForce(matrix [][]int, rows, columns int) []*point {
	path := make([]*point, 0, 1000)

	// bounds checking
	if matrix != nil && rows != 0 {
		// start checking positions from bottom-right
		if ret := bruteForceGetPath(matrix, rows-1, columns-1, &path); ret == true {
			return path
		}
	}

	// otherwise return an empty vector indicating path does not exist
	return path
}

func bruteForceGetPath(matrix [][]int, currRow, currColumn int, path *[]*point) bool {
	//if out of bounds or current position is off limits, return false
	if currRow < 0 || currColumn < 0 || matrix[currRow][currColumn] == -1 {
		return false
	}

	atOrigin := currRow == 0 && currColumn == 0

	// everytime robot moves up or left and it is a valid position, add the point to the slice
	if atOrigin ||
		bruteForceGetPath(matrix, currRow-1, currColumn, path) ||
		bruteForceGetPath(matrix, currRow, currColumn-1, path) {
		currPos := NewPoint(currRow, currColumn)
		*path = append(*path, currPos)

		return true
	}

	return false
}

func f(n int) int {
	if n <= 1 {
		return 1
	}

	return f(n-1) + f(n-1)
}
