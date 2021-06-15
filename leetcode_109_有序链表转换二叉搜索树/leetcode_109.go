package leetcode_109_有序链表转换二叉搜索树

/*
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一： 空间换时间，LinkedList --> Array --> TreeNode
// LinkedList --> Array
func listToArray(list *ListNode) []int {
    var nums []int
    for list != nil {
        nums = append(nums, list.Val)
        list = list.Next
    }
    return nums
}

// Array --> BST
func arrayToBST(nums []int, start, end int) *TreeNode {
    if start == end {
        return nil
    }
    mid := (start + end) >> 1
    root := &TreeNode{Val:nums[mid]}
    root.Left = arrayToBST(nums, start, mid)
    root.Right = arrayToBST(nums, mid+1, end)
    return root
}

// LinkedList --> Array --> BST
func sortedListToBST(head *ListNode) *TreeNode {
    nums := listToArray(head)
    return arrayToBST(nums, 0, len(nums))
}

// 方法二： 快慢指针，V(fast)=2V(slow), 当fast = nil的时候，slow刚好走到mid，主要是求mid
