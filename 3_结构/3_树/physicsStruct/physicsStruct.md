##  physicsStruct

###   双亲表示
need save all node
```shell
[data][parent]
```



###   孩子表示
use 多重链表法 : one node have more pointer which point to son node

* count of pointer == 度 of tree
Use more space
```shell
[data][son][son]...[son]
```

* count of pointer == 度 of current node
Use more time
```shell
[data][countOfSon][son][son]...[son]
```

* array + link
include all node and son node
```shell
[nodeData][firstSonPointer] ->	[indexOfSon][nextSonPointer] ->
[nodeData][firstSonPointer] ->
...
[nodeData][firstSonPointer]
```



###   孩子兄弟表示
```shell
[data][firstSonPointer][rightBrotherPointer]
```
