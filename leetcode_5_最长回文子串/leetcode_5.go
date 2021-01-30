package leetcode_5_最长回文子串

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：
输入: "cbbd"
输出: "bb"
*/

// 中心拓展法
// 怎么把算出来的最大和原string对应上？
func longestPalindrome(s string) string {
	if s == "" || len(s) < 1 {
		return ""
	}
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	size := len(s)
	for i := 0; i < size; i++ {
		odd := extendFromCenter(s, i, i)
		even := extendFromCenter(s, i, i+1)
		mx := max(odd, even)
		if mx > (end - start) { // 只有大于的才更新，小于等于的就不更新了
			start = i - (mx-1)/2 // i都在子串里面，所以这个确定了i在原字符串里的起始位置
			end = i + mx/2       // 确定了在源字符串的结束位置
		}
	}
	return s[start : end+1]
}

// 从中心开始，向两边扩散，返回子串长度，用于后面定位坐标
func extendFromCenter(s string, left, right int) int {
	runes := []rune(s)
	size := len(runes)
	for left >= 0 && right < size && runes[left] == runes[right] {
		left--
		right++
	}
	// 返回子串长度
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
