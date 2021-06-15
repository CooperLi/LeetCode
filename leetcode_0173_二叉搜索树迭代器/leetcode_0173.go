package leetcode_0173_二叉搜索树迭代器

/*
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
调用 next() 将返回二叉搜索树中的下一个最小的数。

示例：
    7
   / \
  3   15
      / \
     9   20

BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false

提示：
next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。
*/

/*
思路：因为需要返回下一个最小的数，所以考虑到“二叉树的中序遍历可以得到递增的数列”这个特点
将中序遍历的结果存到数组nums中
访问nums的头部元素即可得到最小值
有没有下一个最小值可以判断nums是否为空
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 成员：
// nums: 保存递增的切片
// root: 根节点
type BSTIterator struct {
	nums []int
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	nums := make([]int, 0)
	// 得到二叉树中序遍历的切片形式
	inorder(root, &nums)
	return BSTIterator{
		nums: nums,
		root: root,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	// 数组第一个为最小的
	smallest := this.nums[0]
	// 删除第一个元素
	this.nums = this.nums[1:]
	return smallest
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.nums) > 0 {
		return true
	}
	return false
}

// 中序遍历：左中右
func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	// 左
	inorder(root.Left, nums)
	// 中
	*nums = append(*nums, root.Val)
	// 右
	inorder(root.Right, nums)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
