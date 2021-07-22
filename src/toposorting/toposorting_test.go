package toposorting

import (
	"fmt"
	"testing"
)

func TestTopoSorting(t *testing.T) {
	initGraph(4)
	adj[3].add(2)
	adj[2].add(1)
	adj[2].add(4)
	fmt.Println(adj)
}

