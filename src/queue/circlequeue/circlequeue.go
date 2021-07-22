package circlequeue

import "fmt"

type CircleQueue struct {
	items []int
	capacity int
	head int
	tail int
}

func initQueue(capacity int) *CircleQueue {
	circleQueue := CircleQueue{
		items: make([]int, capacity),
		capacity: capacity,
	}
	return &circleQueue
}

func (cq *CircleQueue) Push(elem int) bool {
	// 判断循环队列是否已满
	if (cq.tail + 1) % cq.capacity == cq.head {
		fmt.Println("current circle queue is full!")
		return false
	}
	cq.items[cq.tail] = elem
	cq.tail = (cq.tail + 1) % cq.capacity
	return true
}

func (cq *CircleQueue) Pop() int {
	// 如果循环队列此时为空
	if cq.tail == cq.head {
		fmt.Println("current circle queue is empty!")
		return -1
	}
	elem := cq.items[cq.head]
	cq.items[cq.head] = 0
	cq.head = (cq.head + 1) % cq.capacity
	return elem
}

func (cq *CircleQueue) Print() {
	if cq.head == cq.tail {
		fmt.Println("current circle queue is empty!")
		return
	}
	for i := 0; i < cq.capacity; i++ {
		fmt.Printf("%v ", cq.items[i])
	}
	fmt.Println()
}
