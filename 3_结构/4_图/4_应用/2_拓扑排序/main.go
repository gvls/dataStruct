package main

import "fmt"

// package main O(v+e)
// 使用 栈 就像 dfs
// 使用 队列 就像 bfs

// Edge
type Edge struct {
	HeadIndex int
	Weight    int
	Next      *Edge
}

// Vertex
type Vertex struct {
	In   int
	Data int
	Edge *Edge
}

const (
	MAXVERTEX int = 14
)

// main
func main() {
	G := [MAXVERTEX]Vertex{
		{0, 0, &Edge{11, 1, &Edge{5, 1, &Edge{4, 1, nil}}}},
		{0, 1, &Edge{8, 1, &Edge{4, 1, &Edge{2, 1, nil}}}},
		{2, 2, &Edge{9, 1, &Edge{6, 1, &Edge{5, 1, nil}}}},
		{0, 3, &Edge{13, 1, &Edge{2, 1, nil}}},
		{2, 4, &Edge{7, 1, nil}},
		{3, 5, &Edge{12, 1, &Edge{8, 1, nil}}},
		{1, 6, &Edge{5, 1, nil}},
		{2, 7, nil},
		{2, 8, &Edge{7, 1, nil}},
		{2, 9, &Edge{11, 1, &Edge{10, 1, nil}}},
		{1, 10, &Edge{13, 1, nil}},
		{2, 11, nil},
		{1, 12, &Edge{9, 1, nil}},
		{2, 13, nil},
	}
	Topu(G)

}

// Topu
func Topu(G [MAXVERTEX]Vertex) {
	// init
	stack := [MAXVERTEX]int{}
	top := 0
	count := 0
	// 入度为0 的顶点 入栈
	for k, v := range G {
		if v.In == 0 {
			stack[top] = k
			top++
		}
	}
	// 得到 拓扑排序
	for top != 0 {
		// pop
		top--
		t := stack[top]
		fmt.Println("pop : ", t, " ") // get one result
		count++

		// visit 相邻的点是否变成 入度为0 的点
		for i := G[t].Edge; i != nil; i = i.Next {
			k := i.HeadIndex
			//fmt.Println(k, " ", G[k].In)
			G[k].In--
			if G[k].In == 0 {
				stack[top] = k
				top++
			}
		}
	}

	// check if have 环
	if count < MAXVERTEX {
		fmt.Println("not exit 拓扑排序")
		return
	}
	fmt.Println("exit 拓扑排序")
}
