package linkedlistStack

import "fmt"

type Node struct {
	val int
	next *Node
}

type Stack struct {
	items *Node
	capacity int
	count int
}

func (s *Stack) findTail() (tail *Node, prior *Node) {
	var tmp *Node
	head := s.items
	for true {
		if head.next == nil {
			break
		}
		tmp = head
		head = head.next
	}
	return head, tmp
}

func initStack(capacity int) *Stack {
	stack := Stack{
		capacity: capacity,
	}
	return &stack
}

func (s *Stack) Push(elem int) bool {
	if s.count == s.capacity {
		return false
	}
	if s.count == 0 {
		s.items = &Node{
			val: elem,
		}
		s.count++
		return true
	}
	tail, _ := s.findTail()
	tail.next = &Node{
		val: elem,
	}
	s.count++
	return true
}

func (s *Stack) Pop() int {
	if s.count == 0 {
		return -1
	}
	tail, prior := s.findTail()
	if prior == nil {
		value := s.items.val
		s.items = nil
		tail = nil
		s.count = 0
		return value
	}
	prior.next = nil
	s.count--
	return tail.val
}

func (s *Stack) Print() {
	if s.items == nil {
		fmt.Println("current stack is empty!")
		return
	}
	head := s.items
	for true {
		if head.next == nil {
			fmt.Printf("%v ", head.val)
			break
		}
		fmt.Printf("%v ", head.val)
		head = head.next
	}
	fmt.Println()
}
