package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	//node4 := ListNode{Val: 4}
	//node3 := ListNode{Val: 3, Next: &node4}
	//node2 := ListNode{Val: 2, Next: &node3}
	//node1 := ListNode{Val: 1, Next: &node2}
	////head := swapPairs(&node4)
	////fmt.Println(head)
	//reverseList(&node1)
	//one := TreeNode{Val: 1}
	//three := TreeNode{Val: 3}
	//two := TreeNode{Val: 2, Left: &one, Right: &three}
	//seven := TreeNode{Val: 7}
	//root := TreeNode{Val: 4, Left: &two, Right: &seven}
	//node := searchBST(&root, 5)
	//fmt.Println(node)
	fmt.Println(getRow(2))
}

func getRow(rowIndex int) []int {
	var res, pre []int
	for i := 0; i <= rowIndex; i++ {
		if i == 0 {
			res = []int{1}
		} else {
			res = make([]int, i+1)
			res[0] = 1
			res[i] = 1
			for j := 1; j < i; j++ {
				res[j] = pre[j] + pre[j-1]
			}
		}
		pre = res
	}
	return res
}

func searchBST(root *TreeNode, val int) *TreeNode {
	var res *TreeNode
	if root == nil || root.Val == val {
		res = root
	} else if root.Val < val {
		res = searchBST(root.Right, val)
	} else {
		res = searchBST(root.Left, val)
	}
	return res
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
