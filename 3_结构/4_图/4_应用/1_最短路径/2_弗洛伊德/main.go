package main

import "fmt"

// package main 计算所有顶点到其他顶点的距离 O(n^3)
// 适合 计算所有顶点到其他顶点的距离  or  稠密图
// 使用 动态规划, 不断添加中间点 优化最短路径
// 适合 有向图/无向图

const (
	MAXVERTEX int = 9
	d         int = 99999
)

// main
func main() {
	G := [MAXVERTEX][MAXVERTEX]int{
		{0, 1, 5, d, d, d, d, d, d},
		{1, 0, 3, 7, 5, d, d, d, d},
		{5, 3, 0, d, 1, 7, d, d, d},
		{d, 7, d, 0, 2, d, 3, d, d},
		{d, 5, 1, 2, 0, 3, 6, 9, d},
		{d, d, 7, d, 3, 0, d, 5, d},
		{d, d, d, 3, 6, d, 0, 2, 7},
		{d, d, d, d, 9, 5, 2, 0, 4},
		{d, d, d, d, d, d, 7, 4, 0},
	}
	floyd(G)
}

// floyd
func floyd(G [MAXVERTEX][MAXVERTEX]int) {
	N := [MAXVERTEX][MAXVERTEX]int{} // 最短路径下，该点的前驱
	D := [MAXVERTEX][MAXVERTEX]int{} // 到该点最短路径 权重和

	// init
	for i := 0; i < MAXVERTEX; i++ {
		for j := 0; j < MAXVERTEX; j++ {
			D[i][j] = G[i][j]
			//N[i][j] = i // i -> j 的上一个节点是 i
			N[i][j] = j // i -> j 的下一个节点是 j
		}
	}
	fmt.Println("")
	for _, v := range N {
		fmt.Println(v)
	}

	// 计算 任何顶点到其他点的 最短路径
	for newVer := 0; newVer < MAXVERTEX; newVer++ { // 不断添加中间点
		for b := 0; b < MAXVERTEX; b++ { // 起点
			for e := 0; e < MAXVERTEX; e++ { // 终点
				if D[b][e] > D[b][newVer]+D[newVer][e] { // 加中间点是否变短
					D[b][e] = D[b][newVer] + D[newVer][e]
					//P[b][e] = P[newVer][e] // A -> C 添加 C 的上一个结点 B
					N[b][e] = N[b][newVer] // A -> C 添加 A 的下一个结点 B
				}
			}
		}
	}
	fmt.Println("")
	for _, v := range N {
		fmt.Println(v)
	}
	fmt.Println("")

	//test
	for b := 0; b < MAXVERTEX; b++ {
		for e := 0; e < MAXVERTEX; e++ {
			fmt.Print(b, "-", e, ":", D[b][e])
			k := N[b][e]

			fmt.Print("path: ", b)
			for k != e {
				fmt.Print(" -> ", k)
				k = N[k][e]
			}
			fmt.Println(" -> ", e)
		}
		fmt.Println("")
	}

	//test
	//for b := 0; b < 1; b++ {
	//for e := 8; e < 9; e++ {
	//fmt.Print(b, "-", e, ":", D[b][e])
	//k := P[b][e]
	//fmt.Print("path: ", b)
	//for k != e {
	//fmt.Print(" -> ", k)
	//k = P[k][e]
	//}
	//fmt.Println(" -> ", e)
	//}
	//fmt.Println("")
	//}
}
