package yanghuitriangle

import (
	"fmt"
	"math"
)

type Node struct {
	val int
	left *Node
	right *Node
}

type Queue struct {
	nodes []*Node
	p int
}

func NewQueue() *Queue {
	return &Queue{
		[]*Node{},
		0,
	}
}

func (q *Queue) Push(n *Node) {
	for i := 0; i < len(q.nodes); i++ {
		if q.nodes[i] == n {
			return
		}
	}
	q.nodes = append(q.nodes, n)
}

func (q *Queue) Pop() *Node {
	n := q.nodes[q.p]
	q.p++
	return n
}

func (q *Queue) hasNext() bool {
	return q.p < len(q.nodes)
}

func printTriangle(n *Node) {
	if n != nil {
		queue := NewQueue()
		queue.Push(n)
		for queue.hasNext() {
			cur := queue.Pop()
			fmt.Printf("%d ", cur.val)
			if cur.left != nil {
				queue.Push(cur.left)
			}
			if cur.right != nil {
				queue.Push(cur.right)
			}
		}
	}
}

var minV = -1
func shortestRoute(i, h, cv int, node *Node) {
	cv += node.val
	if node.left != nil {
		shortestRoute(i + 1, h, cv, node.left)
	}
	if node.right != nil {
		shortestRoute(i + 1, h, cv, node.right)
	}
	if i == h {
		if minV == -1 {
			minV = cv
		} else if minV > cv {
			minV = cv
		}
	}
}

func yanghuiTriangle(triangle [][]*Node, h int) int {
	states := make([][]int, h)
	for i := 0; i < h; i++ {
		states[i] = make([]int, i + 1)
	}
	states[0][0] = triangle[0][0].val
	for i := 1; i < h; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				states[i][j] = states[i - 1][j] + triangle[i][j].val
			} else if j == len(triangle[i]) - 1 {
				states[i][j] = states[i - 1][j - 1] + triangle[i][j].val
			} else {
				left := states[i - 1][j - 1]
				right := states[i - 1][j]
				states[i][j] = int(math.Min(float64(left), float64(right))) + triangle[i][j].val
			}
		}
	}
	minV := states[h - 1][0]
	for i := 1; i < len(triangle[h - 1]); i++ {
		if minV > states[h - 1][i] {
			minV = states[h - 1][i]
		}
	}
	return minV
}


