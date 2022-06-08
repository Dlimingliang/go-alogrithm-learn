package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	//node4 := ListNode{Val: 4}
	//node5 := ListNode{Val: 5}
	node1.Next = &node2
	node2.Next = &node3
	//node3.Next = &node4
	//node4.Next = &node5

	//fmt.Println(hasCycle(&node1))
	//fmt.Println(detectCycle(&node1).Val)
	//fmt.Println(getIntersectionNode(&node1, &node4))
	//removeNthFromEnd(&node1, 1)
	//reverseList(&node1)
	//removeElements(&node1, 1)
	oddEvenList(&node1)
}

func oddEvenList(head *ListNode) *ListNode {
	temp := head
	oddDump, evenDump := &ListNode{Val: 0}, &ListNode{Val: 0}
	tempOdd, tempEven := oddDump, evenDump
	for index := 1; temp != nil; index++ {
		if index%2 == 0 {
			tempEven.Next = &ListNode{Val: temp.Val}
			tempEven = tempEven.Next
		} else {
			tempOdd.Next = &ListNode{Val: temp.Val}
			tempOdd = tempOdd.Next
		}
		temp = temp.Next
	}
	tempOdd.Next = evenDump.Next
	return oddDump.Next
}

func removeElements(head *ListNode, val int) *ListNode {
	dump := &ListNode{
		Next: head,
	}
	for temp := dump; temp.Next != nil; {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return dump.Next
}

func reverseList(head *ListNode) *ListNode {
	//向前移动法
	//if head == nil {
	//	return nil
	//}
	//tempHead := head
	//for head.Next != nil {
	//	temp := head.Next
	//	head.Next = temp.Next
	//	temp.Next = tempHead
	//	tempHead = temp
	//}
	//return tempHead

	//反转指针法
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//计算长度法
	//count := 0
	//temp := head
	//for temp != nil {
	//	temp = temp.Next
	//	count++
	//}
	//
	//removeIndex := count - n
	//if removeIndex == 0 {
	//	head = head.Next
	//} else {
	//	pre, current := head, head
	//	for removeIndex > 0 {
	//		pre, current = current, current.Next
	//		removeIndex--
	//	}
	//	pre.Next = current.Next
	//}
	//return head

	//双指针法
	dump := &ListNode{Val: 0, Next: head}
	fast, slow := head, dump
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dump.Next
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
