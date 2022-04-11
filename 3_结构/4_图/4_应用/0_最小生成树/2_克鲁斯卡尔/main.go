package main

import "fmt"

// package main 不断取最短边 若会形成回路就跳过该边 直到取完所有边 O(eloge)
// 适合少边
// 适用于无向图

const (
	MAXVERTEX int = 9
	di        int = 99999
)

// Edge
type Edge struct {
	Begin  int
	End    int
	Weight int
}

// main
func main() {
	// 					[V0]
	// 				  /      \
	// 				10 		  11
	// 			   / 		    \
	// 		   [V1]--16-[V6]-17--[V5]
	// 		  /   |  	 | \ 		|
	//      18    12 	 |  19 		|
	//	   / 	  | 	 | 	 \ 		|
	// 	[V2]--8--[V8] 	/ 	 [V7]   26
	// 		\ 	   |   24 	/   \ 	|
	// 		 22    21  /   / 	 7	|
	//   		\  |  /  16 	  \ |
	// 			  [V3]__/----20--[V4]
	//
	G := []Edge{ // from small to big
		{4, 7, 7},
		{2, 8, 8},
		{0, 1, 10},
		{0, 5, 11},
		{1, 8, 12},
		{3, 7, 16},
		{1, 6, 16},
		{5, 6, 17},
		{1, 2, 18},
		{6, 7, 19},
		{3, 4, 20},
		{3, 8, 21},
		{2, 3, 22},
		{3, 6, 24},
		{4, 5, 26},
	}
	kruskal(G)
}

// kruskal
func kruskal(G []Edge) {
	// value	1 2 3
	// 			^ ^ ^ ^
	// 			| | | |
	// index	0 1 2 3
	tail := [MAXVERTEX]int{} // 多个箭头的集合，有且只有一个尾部
	for k := range tail {    // 用Golang则可以忽略
		tail[k] = 0
	}

	// 找 n-1 条边
	for _, v := range G { // 连续取 边最短
		b := find(tail, v.Begin)
		e := find(tail, v.End)
		if b != e { // 不存在环; 已有箭头集合的头部 != 已有箭头集合的头部
			tail[b] = e // 已有箭头集合的头部 与 新点 链接
			fmt.Println(v.Begin, v.End, v.Weight)
		}

	}
}

// find 点在已有的箭头集合中能到的最远位置
// 已有的点 -> 已有箭头集合的头部
// 新点     -> 新点
func find(tail [MAXVERTEX]int, f int) int {
	for tail[f] > 0 {
		f = tail[f]
	}
	return f
}
