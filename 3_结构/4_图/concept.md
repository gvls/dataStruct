##  concept

* 边
v1 <-> v2

* 无向边
edge which between v1 and v2 **haven't** direction
```shell
(v1,v2)
```

* 无向图
all edge **haven't** direction

* 有向边/弧
edge which from v1 to v2 **have** direction
v1 : 弧尾
v2 : 弧头(**箭头**)
```shell
<v1,v2>
```

* 有向图
**have** edge have direction

* 无向**完全**图
任意顶点之间有边
If have n vertex, count of edge is (nx(n-1))/2
n : n vertex
n-1 : edge of one vertex
/2 : repetition

* 有向**完全**图
任意顶点之间都有**2个方向**的边
If have n vertex, count of edge is nx(n-1)
n : n vertex
n-1 : edge of one vertex

* 稀疏图
have less vertex or edge

* 稠密图
have more vertex and edge

* 权weight
边/弧 have value

* 网
边/弧 have weight

* 子图
若图B的顶点和边都属于图A，那个B是A的子图

* 邻接点
v1到v2 **只** 通过一条边
v1与v2 关联



###   度
* 度 of vertex
count of edge which vertex 关联
```shell
TD(v) = OD(v) + ID(v)
```

* 出度 of vertex
从该定点指**出**的边的数量
```shell
OD(v)
```

* 入度 of vertex
箭头指向该定点的边的数量
```shell
ID(v)
```



###   路径
* 路径
v1到v2所经过所有顶点序列

* 路径的长度
count of 边/弧 in 路径

* 回路/环
end vertex is first vertex

* 简单路径
not have repetition vertex

* 简单回路/简单环
except first vertex and end vertex, not have repetition vertex



###   连通
* 连通
v1到v2 have 路径

* 连通图
无向图中 任意2个顶点连通 (不一定是完全图)

* 连通分量
无向图的 极大联通子图
子图是连通图的情况下 尽量多顶点
包含对应顶点所有的边

* 强连通图
有向图中 任何2个顶点都有2个方向的连通

* 强连通分量
有向图的 极大联通子图
子图是连通图的情况下 尽量多顶点
包含对应顶点所有的边



###   生成
* 连通图的生成树
极小的连通子图
包含图所有的顶点n，但仅有足以构成树的n-1条边

* 有向树
一个有向图，只有一个顶点入度为0，其余顶点入度为1

* 有向图的生成森林
consit of some 有向树
have all vertex of picture
有足以构成若干颗 不相交的有向树的 弧
