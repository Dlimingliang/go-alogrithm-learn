package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	node1 := ListNode{Val: 9}
	node2 := ListNode{Val: 9}
	node3 := ListNode{Val: 9}
	node4 := ListNode{Val: 9}
	node5 := ListNode{Val: 9}
	node6 := ListNode{Val: 9}
	node7 := ListNode{Val: 9}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7

	node11 := ListNode{Val: 9}
	node12 := ListNode{Val: 9}
	node13 := ListNode{Val: 9}
	node14 := ListNode{Val: 9}
	node11.Next = &node12
	node12.Next = &node13
	node13.Next = &node14

	//fmt.Println(hasCycle(&node1))
	//fmt.Println(detectCycle(&node1).Val)
	//fmt.Println(getIntersectionNode(&node1, &node4))
	//removeNthFromEnd(&node1, 1)
	//reverseList(&node1)
	//removeElements(&node1, 1)
	//oddEvenList(&node1)
	addTwoNumbers(&node1, &node11)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dump := &ListNode{Val: 0}
	temp := dump
	l1Val, l2Val, addNum := 0, 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		} else {
			l1Val = 0
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		} else {
			l2Val = 0
		}
		val := l1Val + l2Val + addNum
		next := val % 10
		addNum = val / 10
		temp.Next = &ListNode{Val: next}
		temp = temp.Next
	}

	if addNum != 0 {
		temp.Next = &ListNode{Val: addNum}
	}

	return dump.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	dump := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			dump.Next = list1
			dump = dump.Next
			list1 = list1.Next
		} else {
			dump.Next = list2
			dump = dump.Next
			list2 = list2.Next
		}
	}

	if list1 != nil {
		dump.Next = list1
	}

	if list2 != nil {
		dump.Next = list2
	}
	return head.Next
}

func oddEvenList(head *ListNode) *ListNode {

	//自己的想法
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//
	//temp := head.Next.Next
	//oddDump, evenDump := head, head.Next
	//tempOdd, tempEven := oddDump, evenDump
	//for index := 3; temp != nil; index++ {
	//	if index%2 == 0 {
	//		tempEven.Next = temp
	//		tempEven = tempEven.Next
	//	} else {
	//		tempOdd.Next = temp
	//		tempOdd = tempOdd.Next
	//	}
	//	temp = temp.Next
	//}
	//tempOdd.Next = nil
	//tempEven.Next = nil
	//tempOdd.Next = evenDump
	//return oddDump

	//官方的写法
	if head == nil {
		return head
	}

	evenHead := head.Next
	odd := head
	even := evenHead
	if even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
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
