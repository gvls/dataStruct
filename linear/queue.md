##  queue
队列/排队 have order. 先进先出.

###   array
search is fast

```shell
				add()
				  |
				  v
		 _		_____
		/		|	|
	   /		|	| <- tail
 size()		   /| x |
	  \	  len()	| x |
	   \	   \| x | <- head
		\-		-----
				  |
				  v
				  get()
```


###   one-way link
insert and delete is fast

```shell
						first element
							|
							v
[头指针] -> [头节点] -> [首元节点] -> ...

```

* insert
1. new node connect right node
2. left node connect new node

* delete
1. save element
2. left node connect right node 



###   two-way link
have next pointer and pre pointer
