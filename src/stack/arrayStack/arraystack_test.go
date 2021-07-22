package arrayStack

import "testing"

func TestArrayStack(t *testing.T) {
	s := initStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()
	t.Log(s)
	s.Pop()
	s.Pop()
	s.Pop()
	t.Log(s)
}
