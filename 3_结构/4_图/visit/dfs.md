##  dfs
深度优先
use recursion
like 前序遍历 of tree


###   process
连通图 : do once dfs
非连通图 : after do dfs, check `visited[]`, if have vertex which not be visited. repetition...

1. for : init `visited[]` 
2. for : do dfs()
3. mark `visited[]` 
4. do for to call dfs() 
