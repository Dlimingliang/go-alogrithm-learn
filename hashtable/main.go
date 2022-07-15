package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//mySet := hashset.Constructor()
	//mySet.Add(1)
	//mySet.Add(2)
	//fmt.Println(mySet.Contains(1))
	//fmt.Println(mySet.Contains(3))
	//mySet.Add(2)
	//fmt.Println(mySet.Contains(2))
	//mySet.Remove(2)
	//fmt.Println(mySet.Contains(2))
	//myMap := hashmap.Constructor()
	//myMap.Put(1, 1)
	//myMap.Put(2, 2)
	//fmt.Println(myMap.Get(1))
	//fmt.Println(myMap.Get(3))
	//myMap.Put(2, 1)
	//fmt.Println(myMap.Get(2))
	//myMap.Remove(2)
	//fmt.Println(myMap.Get(2))
	//fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
	//fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	//fmt.Println(isHappy(2))
	//fmt.Println(twoSum([]int{3, 2, 4}, 6))
	//fmt.Println(isIsomorphic("foo", "bar"))
	//fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
	//	[]string{"KFC", "Shogun", "Burger King"}))
	//fmt.Println(firstUniqChar("leetcode"))
	//fmt.Println(intersect([]int{9, 4, 9, 8, 4}, []int{4, 9, 5}))
	//fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	//fmt.Println(groupAnagrams([]string{""}))
	//fmt.Println(isValidSudoku([][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))
	//threeLeftOne := TreeNode{Val: 4}
	//twoLeftOne := TreeNode{Val: 2, Left: &threeLeftOne}
	//fourRightOne := TreeNode{Val: 4}
	//threeRightOne := TreeNode{Val: 2, Left: &fourRightOne}
	//threeRightTwo := TreeNode{Val: 4}
	//twoRightOne := TreeNode{Val: 3, Left: &threeRightOne, Right: &threeRightTwo}
	//root := TreeNode{Val: 1, Left: &twoLeftOne, Right: &twoRightOne}
	//fmt.Println(findDuplicateSubtrees(&root))
	//fmt.Println(numJewelsInStones("z", "ZZ"))
	fmt.Println(lengthOfLongestSubstring("asljlj"))
}

func lengthOfLongestSubstring(s string) int {
	//res := 0
	//subStrMap := map[byte]int{}
	//slice := make([]byte, 0)
	//for i := 0; i < len(s); i++ {
	//	if index, ok := subStrMap[s[i]]; ok {
	//		if len(slice) > res {
	//			res = len(slice)
	//		}
	//		slice = slice[index+1:]
	//		slice = append(slice, s[i])
	//		tempMap := map[byte]int{}
	//		for j := 0; j < len(slice); j++ {
	//			tempMap[slice[j]] = j
	//		}
	//		subStrMap = tempMap
	//	} else {
	//		subStrMap[s[i]] = len(slice)
	//		slice = append(slice, s[i])
	//	}
	//}
	//if len(slice) > res {
	//	res = len(slice)
	//}
	//return res

	subStrMap := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(subStrMap, s[i-1])
		}
		for rk+1 < n && subStrMap[s[rk+1]] == 0 {
			subStrMap[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[byte]bool, len(jewels))
	for i := 0; i < len(jewels); i++ {
		jewelsMap[jewels[i]] = true
	}
	res := 0
	for i := 0; i < len(stones); i++ {
		if jewelsMap[stones[i]] {
			res++
		}
	}
	return res
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	strMap := make(map[string]int)
	subTreeRecursion(root, &strMap, &res)
	return res
}

func subTreeRecursion(node *TreeNode, mapPoint *map[string]int, res *[]*TreeNode) string {
	if node == nil {
		return "#"
	}

	str := fmt.Sprintf("%d,%s,%s", node.Val, subTreeRecursion(node.Left, mapPoint, res), subTreeRecursion(node.Right, mapPoint, res))
	strMap := *mapPoint

	if count, ok := strMap[str]; ok {
		strMap[str] = count + 1
		if strMap[str] == 2 {
			*res = append(*res, node)
		}
	} else {
		strMap[str] = 1
	}

	return str
}

func isValidSudoku(board [][]byte) bool {
	//n := len(board)
	//m := len(board[0])
	//rowSlice := make([]map[byte]bool, 9)
	//cloumnSlice := make([]map[byte]bool, 9)
	//Slice := make([]map[byte]bool, 9)
	//for y := 0; y < n; y++ {
	//	for x := 0; x < m; x++ {
	//		temp := board[y][x]
	//		if temp != '.' {
	//			if rowSlice[y][temp] {
	//				return false
	//			} else {
	//				if rowSlice[y] == nil {
	//					xx := make(map[byte]bool, 1)
	//					xx[temp] = true
	//					rowSlice[y] = xx
	//				} else {
	//					rowSlice[y][temp] = true
	//				}
	//			}
	//			if cloumnSlice[x][temp] {
	//				return false
	//			} else {
	//				if cloumnSlice[x] == nil {
	//					xx := make(map[byte]bool, 1)
	//					xx[temp] = true
	//					cloumnSlice[x] = xx
	//				} else {
	//					cloumnSlice[x][temp] = true
	//				}
	//			}
	//			index := (y/3)*3 + (x / 3)
	//			if Slice[index][temp] {
	//				return false
	//			} else {
	//				if Slice[index] == nil {
	//					xx := make(map[byte]bool, 1)
	//					xx[temp] = true
	//					Slice[index] = xx
	//				} else {
	//					Slice[index][temp] = true
	//				}
	//			}
	//		}
	//	}
	//}
	//return true

	//一次遍历
	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, str := range strs {
		key := getSortKey(str)
		strMap[key] = append(strMap[key], str)
	}

	res := make([][]string, 0)
	for _, value := range strMap {
		res = append(res, value)
	}
	return res
}

func getSortKey(str string) string {
	split := strings.Split(str, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, ok := numMap[num]; ok {
			temp := i - index
			if temp <= k {
				return true
			}
		}
		numMap[num] = i
	}
	return false
}

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	numMap := make(map[int]int)
	for _, num := range nums1 {
		if count, ok := numMap[num]; ok {
			numMap[num] = 1 + count
		} else {
			numMap[num] = 1
		}
	}

	for _, num := range nums2 {
		if count, ok := numMap[num]; ok && count > 0 {
			res = append(res, num)
			numMap[num] = count - 1
		}
	}
	return res
}

func firstUniqChar(s string) int {
	n := len(s)
	byteMap := make(map[byte]int, n)
	for i := 0; i < n; i++ {
		if _, ok := byteMap[s[i]]; ok {
			byteMap[s[i]] = n + 1
		} else {
			byteMap[s[i]] = i
		}
	}

	min := n
	for _, v := range byteMap {
		if v < min {
			min = v
		}
	}

	if min < n {
		return min
	}
	return -1
}

func findRestaurant(list1 []string, list2 []string) []string {
	var res []string
	strMap := make(map[string]int)
	for i, str := range list1 {
		strMap[str] = i
	}

	min := math.MaxInt64
	for i, str := range list2 {
		if value, ok := strMap[str]; ok {
			if i+value < min {
				min = i + value
				res = []string{str}
			} else if i+value == min {
				res = append(res, str)
			}
		}

	}
	return res
}

func isIsomorphic(s string, t string) bool {
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if _, ok := numMap[num]; ok {
			return []int{numMap[num], i}
		} else {
			numMap[target-num] = i
		}
	}
	return nil
}

func isHappy(n int) bool {
	numMap := make(map[int]bool, 0)
	for {
		value, temp := 0, n
		for temp > 0 {
			num := temp % 10
			num = num * num
			temp = temp / 10
			value += num
		}
		if value == 1 {
			return true
		}
		if numMap[value] {
			return false
		}
		numMap[value] = true
		n = value
	}
	return false
}

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	tempMap := make(map[int]bool, 0)
	numMap := make(map[int]bool)
	if len(nums1) > len(nums2) {
		for _, value := range nums1 {
			numMap[value] = true
		}
		for _, value := range nums2 {
			if numMap[value] {
				tempMap[value] = true
			}
		}
	} else {
		for _, value := range nums2 {
			numMap[value] = true
		}
		for _, value := range nums1 {
			if numMap[value] {
				tempMap[value] = true
			}
		}
	}
	for key, _ := range tempMap {
		res = append(res, key)
	}
	return res
}

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func containsDuplicate(nums []int) bool {
	repeatedMap := make(map[int]bool, 0)
	for _, num := range nums {
		if repeatedMap[num] {
			return true
		} else {
			repeatedMap[num] = true
		}
	}
	return false
}
