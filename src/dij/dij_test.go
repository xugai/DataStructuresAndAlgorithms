package dij

import (
	"DataStructuresAndAlgorithms/src/dij/graph"
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	g := graph.InitialGraph(6)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 4, 15)
	g.AddEdge(1, 3, 2)
	g.AddEdge(1, 2, 15)
	g.AddEdge(3, 2, 1)
	g.AddEdge(2, 5, 5)
	g.AddEdge(3, 5, 12)
	g.AddEdge(4, 5, 10)
	Dijkstra(g, 0, 5)
}

func buildHeap(arr []int) {
	n := len(arr) - 1
	for i := n / 2; i > 0; i-- {
		heapSorting(arr, n, i)
	}
	k := n
	for k > 1 {
		tmp := arr[1]
		arr[1] = arr[k]
		arr[k] = tmp
		k--
		heapSorting(arr, k, 1)
	}
}

func heapSorting(arr []int, n, i int) {
	for true {
		minPosition := i
		if 2 * i <= n && arr[2 * i] < arr[i] {
			minPosition = 2 * i
		}
		if 2 * i + 1 <= n && arr[2 * i + 1] < arr[minPosition] {
			minPosition = 2 * i + 1
		}
		if minPosition == i {
			break
		}
		tmp := arr[minPosition]
		arr[minPosition] = arr[i]
		arr[i] = tmp
		i = minPosition
	}
}

func TestHeapSorting(t *testing.T) {
	arr := []int{-1, 2, 4, 9, 5, 6}
	buildHeap(arr)
	fmt.Println(arr)
}

