package leetcode_0144_二叉树的前序遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// func preorderTraversal(root *TreeNode) []int {
//     var rightStack []*TreeNode
//     var res []int
//
//     for cur := root; cur != nil; {
//         res = append(res, cur.Val)
//         if cur.Left != nil {
//             if cur.Right != nil {
//                 rightStack = append(rightStack, cur.Right)
//             }
//             cur = cur.Left
//         } else { // cur.Left == nil
//             if cur.Right != nil {
//                 cur = cur.Right
//             } else {    // cur.Left == cur.Right == nil
//                 // stack 已空，说明完成遍历
//                 if len(rightStack) == 0 {
//                     break
//                 }
//                 // 否则
//                 // 取出最后放入的右侧节点，继续for循环
//                 cur = rightStack[len(rightStack)-1]
//                 rightStack = rightStack[:len(rightStack)-1]
//             }
//         }
//     }
//     return res
// }

/*
另一种，递归
 */
// 根左右
func preorderTraversal(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }
    // 根
    res = append(res, root.Val)
    // 左
    if root.Left != nil {
        res = append(res, preorderTraversal(root.Left)...)
    }
    // 右
    if root.Right != nil {
        res = append(res, preorderTraversal(root.Right)...)
    }
    return res
}
