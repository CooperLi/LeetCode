package leetcode_0145_二叉树的后续遍历
/*
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 */

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 左右根
func postorderTraversal(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }
    // 左
    if root.Left != nil {
        res = append(res, postorderTraversal(root.Left)...)
    }
    // 右
    if root.Right != nil {
        res = append(res, postorderTraversal(root.Right)...)
    }
    // 根
    res = append(res, root.Val)
    return res
}
