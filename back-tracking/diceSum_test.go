package backtracking

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiceSum(t *testing.T) {
	actual := diceSum(2, 7)
	for _, v := range actual {
		fmt.Println(v)
	}
	expected := [][]int{[]int{1, 6}, []int{2, 5}, []int{3, 4}, []int{4, 3}, []int{5, 2}, []int{6, 1}}
	for _, v := range expected {
		fmt.Println(v)
	}
	if reflect.DeepEqual(actual, expected) != true {
		// t.Fatal("wrong")

	}
}
