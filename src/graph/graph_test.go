package graph

import (
	"testing"
)

func TestGraph(t *testing.T) {
	g := initialGraph(8)
	g.addEdge(0, 1)
	g.addEdge(0, 3)
	g.addEdge(1, 2)
	g.addEdge(1, 4)
	g.addEdge(2, 5)
	g.addEdge(3, 4)
	g.addEdge(4, 5)
	g.addEdge(4, 6)
	g.addEdge(5, 7)
	g.addEdge(6, 7)


	visited := make([]bool, g.v)
	prev := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		prev[i] = -1
	}
	//g.recurDFS(4, 5, prev, visited)
	g.NonRecurDFS(4, 7, prev, visited)
	printRoute(prev, 4, 7)
	// 以下代码是求三度顶点的操作
	//edgeRecord := make([]int, g.v)
	//for i := 0; i < g.v; i++ {
	//	edgeRecord[i] = g.bfs(0, i)
	//}
	//for i := 0; i < g.v; i++ {
	//	if edgeRecord[i] > 0 && edgeRecord[i] <= 3 {
	//		fmt.Printf("%v ", i)
	//	}
	//}
}
