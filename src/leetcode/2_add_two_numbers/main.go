package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	curr := &ListNode{}
	head := curr
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}
	return head.Next
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
	fmt.Println("\nDone")
}
