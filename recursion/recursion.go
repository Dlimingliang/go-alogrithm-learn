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
	//fmt.Println(getRow(3))
	//fmt.Println(fib(3))
	//fmt.Println(climbStairs(44))
	node15 := TreeNode{Val: 15}
	node7 := TreeNode{Val: 7}
	node20 := TreeNode{Val: 20, Left: &node15, Right: &node7}
	node9 := TreeNode{Val: 9}
	root := TreeNode{Val: 3, Left: &node9, Right: &node20}
	fmt.Println(maxDepth(&root))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1
	if left > right {
		return left
	} else {
		return right
	}
}

func climbStairs(n int) int {
	//if n < 2 {
	//	return 1
	//}
	//return climbStairs(n-2) + climbStairs(n-1)

	if n < 2 {
		return 1
	}
	p, q, r := 1, 1, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

var fibMap = make(map[int]int)

func fib(n int) int {
	if value, ok := fibMap[n]; ok {
		return value
	}
	if n < 2 {
		return n
	}
	value := fib(n-1) + fib(n-2)
	fibMap[n] = value
	return value

	//迭代法
	//if n < 2 {
	//	return n
	//}
	//p, q, r := 0, 0, 1
	//for i := 2; i <= n; i++ {
	//	p = q
	//	q = r
	//	r = p + q
	//}
	//return r
}

func getRow(rowIndex int) []int {
	//var res, pre []int
	//for i := 0; i <= rowIndex; i++ {
	//	if i == 0 {
	//		res = []int{1}
	//	} else {
	//		res = make([]int, i+1)
	//		res[0] = 1
	//		res[i] = 1
	//		for j := 1; j < i; j++ {
	//			res[j] = pre[j] + pre[j-1]
	//		}
	//	}
	//	pre = res
	//}
	//return res
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	pre := getRow(rowIndex - 1)
	res := make([]int, rowIndex+1)
	res[0] = 1
	res[rowIndex] = 1
	for i := 1; i < rowIndex; i++ {
		res[i] = pre[i] + pre[i-1]
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
