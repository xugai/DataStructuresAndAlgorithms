package dij

import (
	"DataStructuresAndAlgorithms/src/dij/graph"
	"fmt"
	"math"
)

type Vertex struct {
	id int
	dist int
	f int
	x int
	y int
}

func addVertex(id, x, y int) *Vertex{
	return &Vertex{
		id: id,
		dist: INT_MAX,
		f: INT_MAX,
		x: x,
		y: y,
	}
}

// 求两个顶点之间的曼哈顿距离
func hManhattan(s, t *Vertex) float64 {
	return math.Abs(float64(s.x) - float64(t.x)) + math.Abs(float64(s.y) - float64(t.y))
}

type PriorityQueue struct {
	elem []*Vertex
	capacity int
	count int
}

func initialPriorityQueue(capacity int) *PriorityQueue {
	pq := new(PriorityQueue)
	pq.elem = []*Vertex{}
	pq.capacity = capacity
	pq.elem = append(pq.elem, new(Vertex))
	return pq
}
// todo 因为现在是比较估价函数的值了，所以优先级队列里面的堆化要根据估价函数的值进行小顶堆构造
func (pq *PriorityQueue) poll() *Vertex {
	if pq.count == 0 {
		return nil
	}
	v := pq.elem[1]
	pq.elem[1] = pq.elem[len(pq.elem) - 1]
	heapify(pq.elem)
	pq.count--
	return v
}

func (pq *PriorityQueue) add(v *Vertex) {
	pq.count++
	pq.elem = append(pq.elem, v)
	n := len(pq.elem) - 1
	for n / 2 > 0 && pq.elem[n].f < pq.elem[n / 2].f {
		tmp := pq.elem[n]
		pq.elem[n] = pq.elem[n / 2]
		pq.elem[n / 2] = tmp
		n = n / 2
	}
}

func (pq *PriorityQueue) update(targetV *Vertex) {
	n := 0
	for idx, v := range pq.elem {
		if v.id == targetV.id {
			n = idx
			v.dist = targetV.dist
			break
		}
	}
	for n / 2 > 0 && pq.elem[n].f < pq.elem[n / 2].f {
		tmp := pq.elem[n]
		pq.elem[n] = pq.elem[n / 2]
		pq.elem[n / 2] = tmp
		n = n / 2
	}
}

func (pq *PriorityQueue) clear() {
	pq.count = 0
}

func (pq *PriorityQueue) isEmpty() bool {
	return pq.count == 0
}

func heapify(elem []*Vertex) {
	n := len(elem)
	minPosition, i := 1, 1
	for true {
		minPosition = i
		if 2 * i < n && elem[2 * i] != nil && elem[2 * i].f < elem[i].f {
			minPosition = 2 * i
		}
		if 2 * i + 1 < n && elem[2 * i + 1] != nil && elem[2 * i + 1].f < elem[minPosition].f {
			minPosition = 2 * i + 1
		}
		if minPosition == i {
			break
		}
		tmp := elem[minPosition]
		elem[minPosition] = elem[i]
		elem[i] = tmp
		i = minPosition
	}
}

const INT_MAX = int(^uint(0) >> 1)		// 首位为0，其余全是1
const INT_MIN = ^INT_MAX

func Dijkstra(g *graph.Graph, s, t int) {
	predecessor := make([]int, g.V)
	inQueue := make([]bool, g.V)
	vertexs := make([]*Vertex, g.V)
	for i := 0; i < len(vertexs); i++ {
		vertexs[i] = &Vertex{i, INT_MAX, INT_MAX, -1, -1}
	}
	pq := initialPriorityQueue(g.V)
	vertexs[s].dist = 0
	vertexs[s].f = 0
	pq.add(vertexs[s])
	inQueue[s] = true
	for !pq.isEmpty() {
		vertex := pq.poll()
		for i := 0; i < len(g.Adj[vertex.id]); i++ {
			e := g.Adj[vertex.id][i]
			nextVertex := vertexs[e.T]
			if vertex.dist + e.W < nextVertex.dist {
				// 由dij算法演变到A*算法，将原本构建小顶堆时根据dist值比较改为根据估价函数f值比较，每次都从堆顶中拿到估价函数最小的顶点
				nextVertex.dist = vertex.dist + e.W
				nextVertex.f = nextVertex.dist + int(hManhattan(nextVertex, vertexs[t]))
				predecessor[nextVertex.id] = vertex.id
				if inQueue[nextVertex.id] {
					pq.update(nextVertex)
				} else {
					pq.add(nextVertex)
					inQueue[nextVertex.id] = true
				}
			}
			if nextVertex.id == t {
				pq.clear()		// 一旦发现遍历到终点，则立即跳出循环逻辑
				break
			}
		}
	}
	fmt.Printf("%d ", s)
	print(s, t, predecessor)
}

func print(s, t int, predecessor []int) {
	if s == t {
		return
	}
	print(s, predecessor[t], predecessor)
	fmt.Printf("-> %d ", t)
}


