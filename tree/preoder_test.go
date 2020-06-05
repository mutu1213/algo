package tree

import (
	"fmt"
	"testing"
)

func Test_PreOrder(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1
	root.Right = new(TreeNode)
	root.Right.Val = 2
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 3
	fmt.Println(PreOrder(root))
}

func PreOrder(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = temp.Right
		}
	}
	return result
}

//递归
func PreOrder1(root *TreeNode) []int {
	var result []int
	preHelp(root, &result)
	return result
}

func preHelp(root *TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
	}
	if root.Left != nil {
		preHelp(root.Left, result)
	}
	if root.Right != nil {
		preHelp(root.Right, result)
	}
}
