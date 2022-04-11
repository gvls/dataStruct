package main

import "fmt"

// package main 计算一个顶点到其他顶点的最短路径
// 从一个点作为点集出发 for(不断 找最短的边 并入新点 到点集) O(v^2)
// 像 普里姆 上面最短的边就是那点的最短路径
// 不能有负边
// 适合计算 一个顶点到其他顶点 or  稀疏图
// 适用 有向图/无向图
// 若计算所有顶点到其他顶点的距离 O(n^3)

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
	dijkstra(G)
}

// dijkstra
func dijkstra(G [MAXVERTEX][MAXVERTEX]int) {
	P := [MAXVERTEX]int{}     // 最短路径下，该点的前驱
	D := [MAXVERTEX]int{}     // 到该点最短路径 权重和
	final := [MAXVERTEX]int{} // 是否求得 到该点的最短路径
	var newVer int

	// init
	for i := 0; i < MAXVERTEX; i++ {
		final[i] = 0   // 用Golng则可以忽略
		D[i] = G[0][i] // value : v0 -> vx
		P[i] = 0       // 用Golang则可以忽略
	}
	D[0] = 0     // v0 - v0 is 0
	final[0] = 1 // v0

	// 得到n-1个 确定最短路径 的点
	for i := 1; i < MAXVERTEX; i++ {
		min := d

		// 得到一个 已经确定最短路径 的点
		for j := 0; j < MAXVERTEX; j++ { // 找出离v0最近的点，v0到该点的最短距离一定是当前这个
			if final[j] != 1 && D[j] < min {
				newVer = j
				min = D[j]
			}
		}
		final[newVer] = 1 // 该点求得最短路径

		// 合并 v0到该点再到其他点的距离 和 v0直接到其他点的距离
		for j := 0; j < MAXVERTEX; j++ {
			if final[j] != 1 && (min+G[newVer][j] < D[j]) { // 取最小值
				D[j] = min + G[newVer][j]
				P[j] = newVer // 用新前驱更新
			}
		}
	}
	for k, v := range P {
		fmt.Println(k, " ", v)
	}
}
