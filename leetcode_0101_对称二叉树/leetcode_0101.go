package leetcode_0101_对称二叉树

/*
给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3

说明:
如果你可以运用递归和迭代两种方法解决这个问题，会很加分
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路
递归结束条件：
都为空指针则返回 true
只有一个为空则返回 false

递归过程：
判断两个指针当前节点值是否相等
判断 A 的右子树与 B 的左子树是否对称
判断 A 的左子树与 B 的右子树是否对称

短路：
在递归判断过程中存在短路现象，也就是做 与 操作时，如果前面的值返回 false 则后面的不再进行计算
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(p, q *TreeNode) bool {
	// 空的肯定是对称的
	if p == nil && q == nil {
		return true
	}
	// 长度不一样的一定不对称
	if p == nil || q == nil {
		return false
	}
	// 判断是不是对称的
	// 值相同，而且各自的左等于对方的右，右等于左
	return p.Val == q.Val && symmetric(p.Left, q.Right) && symmetric(p.Right, q.Left)
}
