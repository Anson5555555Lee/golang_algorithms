package snq

// How would you design a stack which, in addition to push and pop,
// has a function min which returns the minimum element?
// Push, pop and min should all operate in O(1) time.
// #27: observe that the minimum element doesn't change very often.
// It only changes when a smaller element is added, or when the smallest element is popped.
// #59:
// #78

type MinStack struct {
	stack []item
}

type item struct {
	min, x int
}

func NewMinStack() *MinStack {
	return &MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}
	prepend := []item{{min, x}}
	this.stack = append(prepend, this.stack...)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[1:]
}

func (this *MinStack) Top() int {
	return this.stack[0].x
}

func (this *MinStack) GetMin() int {
	return this.stack[0].min
}

// Solution #2: we can also solve this by using another stack
