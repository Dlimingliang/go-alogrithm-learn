package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	comNode1 := ListNode{Val: 2}
	comNode2 := ListNode{Val: 4}
	comNode1.Next = &comNode2

	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 9}
	node3 := ListNode{Val: 1}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &comNode1

	node4 := ListNode{Val: 3}
	node4.Next = &comNode1

	//fmt.Println(hasCycle(&node1))
	//fmt.Println(detectCycle(&node1).Val)
	fmt.Println(getIntersectionNode(&node1, &node4))
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	slow := headA
	fast := headB
	for slow != fast {

		if slow == nil {
			slow = headB
		} else {
			slow = slow.Next
		}

		if fast == nil {
			fast = headA
		} else {
			fast = fast.Next
		}
	}
	return slow
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
