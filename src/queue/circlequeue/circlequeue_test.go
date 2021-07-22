package circlequeue

import (
	"fmt"
	"testing"
)

func TestCircleQueue(t *testing.T) {
	cq := initQueue(8)
	cq.Pop()
	cq.Push(1)
	cq.Push(2)
	cq.Push(3)
	cq.Push(4)
	cq.Push(5)
	cq.Push(6)
	cq.Push(7)
	cq.Print()
	fmt.Printf("pop %v\n", cq.Pop())
	cq.Push(8)
	cq.Print()
}

func fn(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return fn(n - 1) + fn(n - 2)
}
