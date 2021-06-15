package leetcode_98_验证二叉搜索树

import (
	"math"
)

/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：
节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

示例 1:
输入:
    2
   / \
  1   3
输出: true

示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false

解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
*/

/*
思路：
二叉树特性：根节点要大于左节点同时小于右节点；左右子树同时也要满足这个
所以可以通过递归，来遍历二叉树
*/
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解法
func isValidBST(root *TreeNode) bool {
	// 初始上下界为max和min
	return isBST(root, math.MinInt64, math.MaxInt64)
}

// root：根
// lower：下界
// upper：上界
func isBST(root *TreeNode, lower, upper int) bool {
	// 空树是二叉树
	if root == nil {
		return true
	}
	// 二叉树定义：根节点的值要大于左节点且小于右节点
	// 不满足定义就不是二叉树
	if root.Val <= lower || root.Val >= upper {
		return false
	}

	// 判断左子树是不是二叉树
	// 因为左子树里所有节点的值都要小于根节点
	// 所以上界 upper 改为 root.Val
	isLeft := isBST(root.Left, lower, root.Val)
	// 判断右子树是不是二叉树
	// 因为右子树里所有节点的值要大于根节点
	// 所以下界 lower 改为 root.Val
	isRight := isBST(root.Right, root.Val, upper)

	// 左右子树都是二叉树，才是二叉树
	return isLeft && isRight
}
