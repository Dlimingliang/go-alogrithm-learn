package main

import "fmt"

func main() {
	//reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	node4 := ListNode{Val: 4}
	//node3 := ListNode{Val: 3, Next: &node4}
	//node2 := ListNode{Val: 2, Next: &node3}
	//node1 := ListNode{Val: 1, Next: &node2}
	head := swapPairs(&node4)
	fmt.Println(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[right], s[left] = s[left], s[right]
		left++
		right--
	}
}
