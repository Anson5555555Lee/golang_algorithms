package snq

import "errors"

// Implement a MyQueue class which implements a queue using two stacks.
// Hints: #98, #114

type Queueable interface {
	Add(int)
	Remove() (int, error)
	Peek() int
	IsEmpty() bool
}

type QueueUsingStacks struct {
	new, old *Stack
}

func NewQueueUsingStacks() *QueueUsingStacks {
	new := NewStack()
	old := NewStack()
	return &QueueUsingStacks{new, old}
}

func (this *QueueUsingStacks) newToOld() {
	for len(this.new.stack) > 0 {
		val := this.new.Pop()
		this.old.Push(val)
	}
}

func (this *QueueUsingStacks) Add(x int) {
	this.new.Push(x)
}

func (this *QueueUsingStacks) Remove() (int, error) {
	if this.old.IsEmpty() {
		if this.new.IsEmpty() {
			return -1, errors.New("cannot remove, the queue is empty")
		}
		this.newToOld()
	}

	val := this.old.Pop()
	return val, nil
}

func (this *QueueUsingStacks) IsEmpty() bool {
	return this.new.IsEmpty() && this.old.IsEmpty()
}
