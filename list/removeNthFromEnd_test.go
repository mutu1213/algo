package list

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
//先走N+1步，判断条件比先走N步少一些，因为判断p1是否为nil结束，否则结束条件需要判断p1.next,则需要先判断p1
//head不是空节点，有val
//加入dummy节点，保证p2.next不会空，减少边界判断
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	p1 := dummy
	p2 := dummy
	i := 0
	for ; i < n+1 && p1 != nil; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	if i >= n { //max(i)== len(list);i>=n 说明 len(list)>=n;
		p2.Next = p2.Next.Next
	}
	return dummy.Next
}

func Test_removeNthFromEnd(t *testing.T) {
	head := new(ListNode)
	head.Val = 1
	//head.Next = new(ListNode)
	//head.Next.Val = 2
	//head.Next.Next = new(ListNode)
	//head.Next.Next.Val = 3
	//head.Next.Next.Next = new(ListNode)
	//head.Next.Next.Next.Val = 4
	//head.Next.Next.Next.Next = new(ListNode)
	//head.Next.Next.Next.Next.Val = 5
	fmt.Println(removeNthFromEnd(head, 1))

}
