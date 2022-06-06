package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	//node3 := ListNode{Val: 3}
	//node4 := ListNode{Val: 4}
	//node5 := ListNode{Val: 5}
	node1.Next = &node2
	//node2.Next = &node3
	//node3.Next = &node4
	//node4.Next = &node5

	//fmt.Println(hasCycle(&node1))
	//fmt.Println(detectCycle(&node1).Val)
	//fmt.Println(getIntersectionNode(&node1, &node4))
	removeNthFromEnd(&node1, 1)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	temp := head
	for temp != nil {
		temp = temp.Next
		count++
	}

	removeIndex := count - n
	if removeIndex == 0 {
		head = head.Next
	} else {
		pre, current := head, head
		for removeIndex > 0 {
			pre, current = current, current.Next
			removeIndex--
		}
		pre.Next = current.Next
	}
	return head
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
