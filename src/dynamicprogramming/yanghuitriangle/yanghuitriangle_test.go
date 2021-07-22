package yanghuitriangle

import (
	"fmt"
	"testing"
)

func TestYanghuiTriangle(t *testing.T) {
	node_5_1 := Node{8, nil, nil}
	node_5_2 := Node{7, nil, nil}
	node_5_3 := Node{9, nil, nil}
	node_5_4 := Node{4, nil, nil}
	node_5_5 := Node{5, nil, nil}
	node_4_1 := Node{4, &node_5_1, &node_5_2}
	node_4_2 := Node{9, &node_5_2, &node_5_3}
	node_4_3 := Node{6, &node_5_3, &node_5_4}
	node_4_4 := Node{1, &node_5_4, &node_5_5}
	node_3_1 := Node{2, &node_4_1, &node_4_2}
	node_3_2 := Node{3, &node_4_2, &node_4_3}
	node_3_3 := Node{4, &node_4_3, &node_4_4}
	node_2_1 := Node{7, &node_3_1, &node_3_2}
	node_2_2 := Node{8, &node_3_2, &node_3_3}
	node_1_1 := Node{5, &node_2_1, &node_2_2}
	//printTriangle(&node_1_1)
	//shortestRoute(1, 5, 0, &node_1_1)
	//fmt.Println(minV)
	h := 5
	triangle := make([][]*Node, h)
	for i := 0; i < h; i++ {
		triangle[i] = make([]*Node, i+1)
	}
	triangle[0][0] = &node_1_1
	triangle[1][0] = &node_2_1
	triangle[1][1] = &node_2_2
	triangle[2][0] = &node_3_1
	triangle[2][1] = &node_3_2
	triangle[2][2] = &node_3_3
	triangle[3][0] = &node_4_1
	triangle[3][1] = &node_4_2
	triangle[3][2] = &node_4_3
	triangle[3][3] = &node_4_4
	triangle[4][0] = &node_5_1
	triangle[4][1] = &node_5_2
	triangle[4][2] = &node_5_3
	triangle[4][3] = &node_5_4
	triangle[4][4] = &node_5_5
	fmt.Println(yanghuiTriangle(triangle, h))
}
