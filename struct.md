##  struct

###   logic struct
relation between data element and data element

* collection
all element have same character
all element is **equal**
```shell
 /--------------\
/	x		x	 \
|	   	x		 |
|	x  		x	 |
|x	   	x		 |
|	   			x|
\	 x 		 x	 /
 \--------------/
```

* linear
relation between element is **one face one** 
```shell
[X]-[X]-[X]-[X]-[X]-[X]-[X]
```

* tree
relation between element is **one face more** 
```shell
		[X]
	   /   \
	 [X]    [X]
	/ \	   / \	\
 [X] [X]  [X][X][X]
		...
```

* picture
relation between element is **more face more** 
```shell
		[X]
	   / | \
	  /  |  \
	 /  [X]  \
	/  / |	\ \
    [X]-[X]-[X]
```



###   physics struct
actual place for data

* order 顺序
elment is 一个挨着一个
```shell
[X][X][X][X][X][X][X]
```

* link 链式
element is **不一定**一个挨着一个
use pointer to location palce of other element
```shell
  ____   ____________   ______
 /    \ /            \ /	  \
[X][_][X][_][_][_][_][X][_][_][X]\  ... [_][_]...
[_][_][_][_][_][_][_][_][_][_][_] | ... [_][_]...
[_][_][_][_][_][_][_][_][_][_][_] | ... [_][_]...
[_][_][_][_][_][_][_][_][_][_][X]/	... [_][_]...
```

