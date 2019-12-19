package robot

import (
	"fmt"
	"testing"
)

func TestCreateMatrix(t *testing.T) {
	matrix := createMatrix(5, 7)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != 1 {
				t.Fatal("wrong")
			}
		}
	}
}

func TestBruceForce(t *testing.T) {
	// create a matrix
	matrix := createMatrix(5, 7)
	// set off-limits
	matrix[1][3] = -1
	matrix[2][5] = -1
	matrix[3][2] = -1
	matrix[0][1] = -1

	path := bruteForce(matrix, 5, 7)

	for _, v := range path {
		fmt.Printf("point is %v", v)
	}
	fmt.Println()
	fmt.Println(len(path))
}
