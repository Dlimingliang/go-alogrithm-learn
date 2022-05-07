package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := ListNode{Val: 3}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 0}
	node4 := ListNode{Val: -4}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node2
	//fmt.Println(hasCycle(&node1))
	fmt.Println(detectCycle(&node1).Val)
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if slow != fast {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
