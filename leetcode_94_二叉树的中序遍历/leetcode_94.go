package leetcode_94_二叉树的中序遍历
/*
给定一个二叉树，返回它的中序 遍历。

示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 */


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 左根右
func inorderTraversal(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }
    // 左
    if root.Left != nil {
        res = append(res, inorderTraversal(root.Left)...)
    }
    // 根
    res = append(res, root.Val)
    // 右
    if root.Right != nil {
        res = append(res, inorderTraversal(root.Right)...)
    }
    return res
}
