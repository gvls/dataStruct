package main

import "fmt"

// package main 从一个点作为点集出发 for(不断 找最短的边 并入新点 到点集) O(v^2)
// suit less vertex

const (
	MAXVERTEX int = 9
	di        int = 99999
)

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
	G := [MAXVERTEX][MAXVERTEX]int{
		//v0 v1 v2  v3  v4  v5  v6  v7  v8
		{0, 10, di, di, di, 11, di, di, di},
		{10, 0, 18, di, di, di, 16, di, 12},
		{di, di, 0, 22, di, di, di, di, 8},
		{di, di, 22, 0, 20, di, di, 16, 21},
		{di, di, di, 20, 0, 26, di, 7, di},
		{11, di, di, di, 26, 0, 17, di, di},
		{di, 16, di, di, di, 17, 0, 19, di},
		{di, di, di, 16, 7, di, 19, 0, di},
		{di, 12, 8, 21, di, di, di, di, 0},
	}
	prim(G)
}

// prim
func prim(G [MAXVERTEX][MAXVERTEX]int) {
	// init
	var newVer int
	var edge [MAXVERTEX]int    // index-value : vx - vy : one edge
	var lowCost [MAXVERTEX]int // 某点/点集 到其他点的花费
	edge[0] = 0                // v0 - v0
	lowCost[0] = 0             // if value is 0, Vindex is ok
	for i := 1; i < MAXVERTEX; i++ {
		edge[0] = 0          //初始化为 v0 - v0 , v0 - v0 , ...
		lowCost[i] = G[0][i] // 初始化为 v0 这个点到其他点的花费
	}

	// get n-1 edge
	for i := 1; i < MAXVERTEX; i++ { // lowCost[v0]=0, so i!=0
		newVer = 0

		// get 1 edge : 某顶点/点集 至少被一条边链接，并且这条边一定是最短的
		min := di                        // min 的值必须大于点到点的最大距离
		for j := 1; j < MAXVERTEX; j++ { // lowCost[v0]=0, so i!=0
			if lowCost[j] != 0 && lowCost[j] < min { // 顶点未并入 && 最短的边
				min = lowCost[j] // min-value
				newVer = j       // mini-index
			}
		}
		lowCost[newVer] = 0
		fmt.Println(edge[newVer], " - ", newVer) // 点集合中的某点 - 新点

		// 合并新得到的点到已有的点集
		for j := 1; j < MAXVERTEX; j++ { // lowCost[v0]=0, so i!=0
			if lowCost[j] != 0 && G[newVer][j] < lowCost[j] {
				lowCost[j] = G[newVer][j] // 取点集和新点到某其他点的最小值
				edge[j] = newVer          // 暂时更新 其他点被 新点集 合并时，是链接新点还是旧点集
			}
		}
	}
}
