package leetcode_0108_将有序数组转换为二叉搜索树

/*
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // 二叉搜索树按中序遍历是有序的
// func sortedArrayToBST(nums []int) *TreeNode {
//     if len(nums) == 0 {
//         return nil
//     }
//     mid := len(nums) / 2
//     root := &TreeNode{Val:nums[mid]}
//     root.Left = sortedArrayToBST(nums[:mid])
//     root.Right = sortedArrayToBST(nums[mid+1:])   // mid+1确保是右半边
//     return root
// }

// 更优雅
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}
