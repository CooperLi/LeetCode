package leetcode_2_两数相加

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
例子：
sum = l1.Val + l2.Val + carry = 10
carry = sum / 10 = 1    // 进位值变化
sum = sum % 10 = 0      // 实际存入链表的值
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 保存结果
	head := &ListNode{}
	// 当前坐标指针
	cur := head
	// 标识进位，初始进位为0
	carry := 0
	// 列表不为空或者存在进位
	// 存在进位依旧要循环，直到进位是0，即没有进位，就像加法，如果有进位，要一直加加加，到最后进位没了为止
	for l1 != nil || l2 != nil || carry > 0 {
		// 保存和
		sum := carry
		if l1 != nil {
			// 更新和
			sum += l1.Val
			// 链表往后走
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		// 计算是否进位
		carry = sum / 10
		// 更新当前表头的下一位是进位进了之后的余数，也就是个位数
		cur.Next = &ListNode{Val: sum % 10}
		// 更新表头到下一个位置
		cur = cur.Next
	}
	// 返回结果
	return head.Next
}
