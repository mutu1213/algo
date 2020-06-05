package tree

//很巧的是，将先序遍历改成 1-根 ，2-右子树 ，3-左子树的方式进行遍历后，结果刚好和后续遍历相反
func PostOrder(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Right
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = temp.Left
		}
	}
	if len(result) != 0 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[j], result[i] = result[i], result[j]
		}
	}
	return result
}

func PostOrder1(root *TreeNode) []int {
	var result []int
	postHelp(root, &result)
	return result
}

func postHelp(root *TreeNode, result *[]int) {
	if root != nil {
		if root.Left != nil {
			postHelp(root.Left, result)
		}
		if root.Right != nil {
			postHelp(root.Right, result)
		}
		*result = append(*result, root.Val)
	}
}
