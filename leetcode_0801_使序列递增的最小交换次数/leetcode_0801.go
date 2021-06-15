package leetcode_0801_使序列递增的最小交换次数

/*
我们有两个长度相等且不为空的整型数组 A 和 B 。
我们可以交换 A[i] 和 B[i] 的元素。注意这两个元素在各自的序列中应该处于相同的位置。

在交换过一些元素之后，数组 A 和 B 都应该是严格递增的（数组严格递增的条件仅为A[0] < A[1] < A[2] < ... < A[A.length - 1]）。
给定数组 A 和 B ，请返回使得两个数组均保持严格递增状态的最小交换次数。假设给定的输入总是有效的。

示例:
输入: A = [1,3,5,4], B = [1,2,3,7]
输出: 1
解释:
交换 A[3] 和 B[3] 后，两个数组如下:
A = [1, 3, 5, 7] ， B = [1, 2, 3, 4]
两个数组均为严格递增的。

注意:
A, B 两个数组的长度总是相等的，且长度的范围为 [1, 1000]。
A[i], B[i] 均为 [0, 2000]区间内的整数。
*/

func minSwap(a []int, b []int) int {
	n := len(a)
	r0, r1 := 0, 1

	for i := 1; i < n; i++ {
		newR0, newR1 := 1<<63-1, 1<<63-1
		if a[i] > a[i-1] && b[i] > b[i-1] {
			newR0 = min(newR0, r0)
			newR1 = min(newR1, r1+1)
		}
		if a[i] > b[i-1] && b[i] > a[i-1] {
			newR0 = min(newR0, r1)
			newR1 = min(newR1, r0+1)
		}
		r0, r1 = newR0, newR1
	}
	return min(r0, r1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
