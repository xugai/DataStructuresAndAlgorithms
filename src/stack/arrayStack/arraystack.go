package arrayStack

type Stack struct {
	items []int
	capacity int
	count int
}

func initStack(n int) *Stack {
	stack := Stack{
		items: make([]int, n),
		capacity: n,
		count: 0,
	}
	return &stack
}

func (s *Stack) Push(elem int) bool {
	if s.count == s.capacity {
		return false
	}
	s.items[s.count] = elem
	s.count++
	return true
}

func (s *Stack) Pop() int {
	if s.count == 0 {
		return -1
	}
	elem := s.items[s.count - 1]
	s.items[s.count - 1] = 0
	s.count--
	return elem
}


