package main

import "fmt"

// package main 找到事件的最早才能发生的时间 和事件由于后面的截止时间还没到从而自己最晚发生的时间，推导出 活动最早发生时间 和最晚发生时间 O(v+e)
// 假如有多条关键路径，那么实际生产环境中需要同时处理多条关键路径

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
	MAXVERTEX int = 4
)

// main
func main() {
	// 			[v1]
	// 	   	   /   \
	// 		  3     5
	// 		 /       \
	//   [v0] 		 [v3]
	// 		\       /
	// 		 4     8
	// 		  \   /
	// 		  [v2]
	G := [MAXVERTEX]Vertex{
		//{0, 0, &Edge{2, 4, &Edge{1, 3, nil}}},
		//{1, 1, &Edge{4, 6, &Edge{3, 5, nil}}},
		//{1, 2, &Edge{5, 7, &Edge{3, 8, nil}}},
		//{2, 3, &Edge{4, 3, nil}},
		//{2, 4, &Edge{7, 4, &Edge{6, 9, nil}}},
		//{1, 5, &Edge{7, 6, nil}},
		//{1, 6, &Edge{9, 2, nil}},
		//{2, 7, &Edge{8, 5, nil}},
		//{1, 8, &Edge{9, 3, nil}},
		//{2, 9, nil},

		{0, 0, &Edge{2, 4, &Edge{1, 3, nil}}},
		{1, 1, &Edge{3, 5, nil}},
		{1, 2, &Edge{3, 8, nil}},
		{2, 3, nil},
	}
	CriticalPath(G)
}

// Topo
func Topo(G [MAXVERTEX]Vertex) (*[MAXVERTEX]int, int, *[MAXVERTEX]int) {
	// init
	top := 0
	count := 0
	stack := [MAXVERTEX]int{}
	for i := 0; i < MAXVERTEX; i++ { // 找 入度为0 的入栈
		if G[i].In == 0 {
			stack[top] = i
			top++
		}
	}

	// 事件最早发生时间
	etv := [MAXVERTEX]int{}
	for i := 0; i < MAXVERTEX; i++ {
		etv[i] = 0
	}
	top2 := 0
	stack2 := [MAXVERTEX]int{}

	// 得到整个流程的顶点序列 和对应顶点序列在所有到达它的路径中最长的时间
	// 当该点 前面花费最长时间的 活动加事件 做完，该点才能发生。也就是最早发生时间
	for top != 0 {
		top--
		t := stack[top]
		count++

		stack2[top2] = t
		top2++

		for i := G[t].Edge; i != nil; i = i.Next {
			// 处理边及弧头点
			k := i.HeadIndex  // 得到弧头
			G[k].In--         // 移除弧
			if G[k].In == 0 { // 入度为0 的点入栈
				stack[top] = k
				top++
			}
			if (etv[t] + i.Weight) > etv[k] { // 比较 各个 k的前一个顶点的权重加上到k的弧的权重
				fmt.Println(t, " -> ", k, " now time : ", etv[t]+i.Weight)
				etv[k] = etv[t] + i.Weight // 找出到 k 这个点的最大权重
			}
		}
	}

	if count < MAXVERTEX {
		fmt.Println("not is AOV")
		return nil, 0, nil
	}
	fmt.Println("is AOV")
	return &stack2, top2, &etv
}

// CriticalPath
func CriticalPath(G [MAXVERTEX]Vertex) {
	stack2, top2, etv := Topo(G)
	fmt.Println("")
	for k, v := range stack2 {
		fmt.Printf("事件 %d 最早在时间 %d 才能发生\n", v, etv[k])
	}
	fmt.Println("")
	// 事件最晚发生时间(比这个晚的话，就不行了)
	ltv := [MAXVERTEX]int{}
	for i := 0; i < MAXVERTEX; i++ {
		ltv[i] = (*etv)[MAXVERTEX-1] //初始化为一个非常非常晚的时间
	}
	fmt.Println(ltv)

	// 只有 该事件的后面的事件 才能决定该时间是否能晚点发生。假如该事件的后面的事件 在等待其他事件和活动，并且 该事件和到达后面事件之前的活动已经完成，那个该事件可以晚点做
	for top2 != 0 {
		top2--
		b := stack2[top2]

		for i := G[b].Edge; i != nil; i = i.Next {
			e := i.HeadIndex

			if ltv[e]-i.Weight < ltv[b] {
				ltv[b] = ltv[e] - i.Weight
				fmt.Printf("事件 %d 最早在时间 %d 才能发生。事件 %d 最晚在时间 %d 瞬间后经过活动时间 %d 刚好让事件 %d 发生\n", e, ltv[e], b, ltv[b], i.Weight, e)
			}
		}
	}
	fmt.Println(ltv)

	for i := 0; i < MAXVERTEX; i++ {
		for e := G[i].Edge; e != nil; e = e.Next {
			j := e.HeadIndex

			etc := etv[i]            //活动最早
			lte := ltv[j] - e.Weight //活动最晚

			if etc == lte { // 假如 lte 比 etc 大，也就是 etc 可以再晚一点
				fmt.Println("<", G[i].Data, ",", G[j].Data, ">", e.Weight)
			}
		}
	}
}
