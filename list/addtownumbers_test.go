package list

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	header := new(ListNode)
	carry := 0
	curr := header
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum := x + y + carry

		curr.Next = new(ListNode)
		curr.Next.Val = sum % 10
		curr = curr.Next
		if sum > 9 {
			carry = 1
		} else {
			carry = 0
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		curr.Next = new(ListNode)
		curr.Next.Val = 1
	}
	return header.Next
}
