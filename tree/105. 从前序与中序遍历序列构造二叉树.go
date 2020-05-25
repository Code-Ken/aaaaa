package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	mid := pos(preorder[0], inorder)
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}

func pos(value int, s []int) int {
	for p, v := range s {
		if v == value {
			return p
		}
	}
	return -1
}