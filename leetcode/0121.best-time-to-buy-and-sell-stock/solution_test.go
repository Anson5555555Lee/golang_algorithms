package problem0121

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	prices []int
}

type ans struct {
	one int
}

func Test_Problem0121(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{7, 2, 1, 5, 3, 6, 4},
			},
			ans{
				5,
			},
		},

		question{
			para{
				[]int{7, 1, 5, 3, 6, 4},
			},
			ans{
				5,
			},
		},

		question{
			para{
				[]int{7, 6, 5, 4, 3, 2, 1},
			},
			ans{
				0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, maxProfit(p.prices), "input:%v", p)
	}
}
