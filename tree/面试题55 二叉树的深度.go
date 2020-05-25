package tree

var max int
func maxDepth(root *TreeNode) int {
	max = 0
	dfs(root,1)
	return max
}

func dfs(root *TreeNode,m int){
	if root != nil{
		if max < m {
			max = m
		}
		dfs(root.Left,m+1)
		dfs(root.Right,m+1)
	}
}
