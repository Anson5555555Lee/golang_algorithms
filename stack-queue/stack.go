package snq

import "sync"

type Stackable interface {
	Push(int)
	Pop() int
	Peek() int
	IsEmpty() bool
}

type Stack struct {
	stack []int
	lock  sync.Mutex
}

func NewStack() *Stack {
	s := &Stack{}
	s.stack = []int{}

	return s
}

func (s *Stack) Push(value int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	prepend := []int{value}
	s.stack = append(prepend, s.stack...)
}

func (s *Stack) Pop() (el int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	el, s.stack = s.stack[0], s.stack[1:]

	return el
}

func (s *Stack) Peek() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.stack[0]
}

func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return (len(s.stack)) == 0
}
