package leetcode_739_每日温度

/*
根据每日气温列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。
例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
*/

/*
其实就是输出比当前数大的数和当前数的距离，没有的就是0
立刻想到暴力法，O(n^2)
*/

// T[i] 和 T[i+1:]里面的每个元素比较，有比T[i]大的，就返回两者坐标差
// 记得break
// func dailyTemperatures(T []int) []int {
// 	var res = make([]int, len(T))
// 	for i := 0; i < len(T); i++ {
// 		for j := i + 1; j < len(T); j++ {
// 			if T[j] > T[i] {
// 				res[i] = j - i
// 				break
// 			}
// 		}
// 	}
// 	return res
// }

/*
单调栈解法
新建一个列表tmp，里面存的每个都是第一个比当前元素大的值的索引
答案就是这个列表tmp里的每个值减去自己所在的索引,如果没有就是-1
T  =  [73, 74, 75, 71, 69, 72, 76, 73]
ans = [1,  2,  6,  5,  5,  6,  -1, -1]
怎么遍历一次就找到呢？
维护一个栈stack（单调递减），遍历所有元素
入栈前，比较当前元素和栈顶元素的大小，如果大于栈顶元素，就弹出栈顶，同时弹出栈顶的元素在ans里的值就是这个大的元素的index
如果小于栈顶元素，就入栈
比如目前栈：
stack = [73]， 要比较的元素是74，大于73，那么在ans[0] = indexOf(74) = 1
*/
func dailyTemperatures(T []int) []int {
	var res = make([]int, len(T))
	var stack []int
	for i := 0; i < len(T); i++ {
		// stack[len(stack)-1] ==== stack.Top, 栈顶元素和当前遍历到的T[i]相比较
		// 如果T[i] 大于栈顶元素，那么就弹栈，而
		for len(stack) != 0 && T[stack[len(stack)-1]] < T[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[index] = i - index
		}
		stack = append(stack, i)
	}
	for len(stack) != 0 {
		tmp := stack[len(stack)-1]
		res[tmp] = 0
		stack = stack[:len(stack)-1]
	}
	return res
}
