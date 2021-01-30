package leetcode_206_反转链表

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

/*
我们可以申请两个指针:
第一个指针叫pre，最初是指向null的。
第二个指针cur指向head，然后不断遍历cur。
每次迭代到cur，都将cur的next指向pre，然后pre和cur前进一位。
都迭代完了(cur变成null了)，pre就是最后一个节点了,新的头部。
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    cur := head
    var pre *ListNode
    for cur != nil {
        // 记录当前节点的next， 就是要操作的那个节点
        tmp := cur.Next
        // 将当前节点指向pre
        cur.Next = pre
        // pre 和 cur 都前进一位
        pre = cur
        cur = tmp
    }
    return pre
}