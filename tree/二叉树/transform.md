##  transform
###   tree -> 二叉树
use 孩子兄弟表示法 of physics struct

1. add line between brother node
2. remove line of son node but save line of first son node 
3. first son node is left node
4. brother node is  right node



###   森林 -> 二叉树

1. all tree transform to 二叉树
2. first tree as left node + next tree as right node : get new left node
3. new left node + next tree as right node : get new left node
4. ... 



###   二叉树 -> tree
反向操作 of tree -> 二叉树



###   二叉树 -> 森林
if 二叉树 have right son node, it is 森林 rather than is tree

1. from root node, remove all line of right son node by recursion
2. 二叉树 -> tree 
