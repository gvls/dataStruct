##  queue
队列/排队 have order. 先进先出.
only can add to tail and delete from head

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


###   link
insert and delete is fast

```shell
			first element
				|
				v
[头节点] -> [首元节点] -> [第二节点] -> ...

```

* insert
1. new node connect right node
2. left node connect new node

* delete
1. save element
2. left node connect right node 

