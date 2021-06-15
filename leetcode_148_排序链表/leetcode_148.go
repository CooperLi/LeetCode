package leetcode_148_排序链表

/*
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:
输入: 4->2->1->3
输出: 1->2->3->4

示例 2:
输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// 先返回异常值：head就是空的或者head的next是空的
	// 也就是len(head) <=1 的情况
	if head == nil || head.Next == nil {
		return head
	}
	// 分成左右两个链表
	left, right := split(head)
	// 合并已经排过序的列表
	return merge(sortList(left), sortList(right))
}

// 从中间切成两个，所以要有快慢指针
// 慢指针走一步，快指针走两步，等快指针走到头，慢指针刚好走到中间
func split(head *ListNode) (left, right *ListNode) {
	slow, fast := head, head
	// 确保当len(head)==2，会均分成两个list
	var slowPre *ListNode
	for fast != nil && fast.Next != nil {
		// 保存slow
		slowPre = slow
		// V(fast) = 2V(slow)
		slow, fast = slow.Next, fast.Next.Next
	}
	// 切断list，slow的下一个设成nil就成两个list了
	slowPre.Next = nil
	// 赋值给左右两个list
	left, right = head, slow
	return left, right
}

// 排序+合并
func merge(left, right *ListNode) *ListNode {
	// 初始化地址
	cur := &ListNode{}
	// 同时创建头
	headPre := cur
	// 左右均不为空的情况
	for left != nil && right != nil {
		// 左边的小于右边的值
		if left.Val < right.Val {
			// cur总是存小的值，然后left指向下一个
			cur.Next, left = left, left.Next
		} else {
			// cur 总是存小的值，然后right指向下一个
			cur.Next, right = right, right.Next
		}
		// cur 指向下一个
		cur = cur.Next
	}
	// 把剩下的也接上
	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}
	// 返回head，也就是整个链表
	return headPre.Next
}
