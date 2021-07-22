package linkedlistStack

import "testing"

func TestLinkedListStack(t *testing.T) {
	s := initStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Pop()
	s.Print()
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(8)
	s.Push(9)
	s.Print()
	s.Pop()
	s.Pop()
	s.Print()
}

