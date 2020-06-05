package tree

import (
	"fmt"
	"testing"
)

func Test_InOrder(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1
	root.Right = new(TreeNode)
	root.Right.Val = 2
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 3
	fmt.Println(InOrder1(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InOrder(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, temp.Val)
			root = temp.Right
		}
	}
	return result
}

//递归
func InOrder1(root *TreeNode) []int {
	var result []int
	help(root, &result)
	return result
}

func help(root *TreeNode, result *[]int) {
	if root != nil {
		if root.Left != nil {
			help(root.Left, result)
		}
		*result = append(*result, root.Val)
		if root.Right != nil {
			help(root.Right, result)
		}
	}
}
