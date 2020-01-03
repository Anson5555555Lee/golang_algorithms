package backtracking

import (
	"fmt"
	"testing"
)

func TestDiceSum(t *testing.T) {
	diceSum(4, 11)
	// expected := [][]int{[]int{1, 6}, []int{2, 5}, []int{3, 4}, []int{4, 3}, []int{5, 2}, []int{6, 1}}
	// if reflect.DeepEqual(actual, expected) != true {
	// t.Fatal("wrong")
	// }
	fmt.Println(calls)
}
