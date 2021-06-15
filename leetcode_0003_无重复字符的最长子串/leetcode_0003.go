package leetcode_0003_无重复字符的最长子串

import (
	"strings"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

// func lengthOfLongestSubstring(s string) int {
//     if s == "" {
//         return 0
//     }
//     if len(s) == 1 {
//         return 1
//     }
//     input := []rune(s)
//     var keyMap = make(map[string]int)
//     res := 0
//     start, end := 0, 0
//     for i:=0; i<len(input); i++ {
//         end=i
//         if _, ok := keyMap[string(input[i])];!ok {
//             res = max(res, end-start+1)
//             keyMap[string(input[i])] = i
//         } else {
//             res = max(res, end-start)
//             start = keyMap[string(input[i])]+1
//             keyMap[string(input[i])] = i
//         }
//     }
//     return max(res, end-start)
// }
//
// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

func lengthOfLongestSubstring(s string) int {
	res, start := 0, 0
	for i := 0; i < len(s); i++ {
		subStr := s[start:i]
		ss := string(s[i])
		isExist := strings.Index(subStr, ss)
		if isExist == -1 {
			// i-start+1=子串长度
			res = max(res, i-start+1)
		} else {
			// 滑动多少的窗口
			start = start + 1 + isExist
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
