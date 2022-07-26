package main

func main() {
	//reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	node4 := ListNode{Val: 4}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	//head := swapPairs(&node4)
	//fmt.Println(head)
	reverseList(&node1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//var pre *ListNode
	//node := head
	//for node != nil {
	//	temp := node.Next
	//	node.Next, pre = pre, node
	//	node = temp
	//}
	//return pre
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseListRecursion(pre, node *ListNode) *ListNode {
	if node == nil {
		return pre
	}
	head := reverseListRecursion(node, node.Next)
	node.Next = pre
	return head
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
