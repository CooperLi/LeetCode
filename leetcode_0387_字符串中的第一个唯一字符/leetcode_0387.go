package leetcode_0387_字符串中的第一个唯一字符

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

示例：
s = "leetcode"
返回 0

s = "loveleetcode"
返回 2
*/

/*
思考：
由于字母一共就26个，所以建立一个长度为26的数组
第一次遍历，记录每个字母最后一次出现的位置
第二次遍历，比较字母第一次出现的索引是不是最后一次出现的位置
如果是，就返回这个元素的index
如果不是，就在这个位置写入-1，标记这个元素已经是重复的了，不满足要求
如果遍历完没有找到，就返回-1
*/

func firstUniqChar(s string) int {
	lastIndex := [26]int{}

	for i, v := range s {
		lastIndex[v-'a'] = i
	}

	for i, v := range s {
		if lastIndex[v-'a'] == i {
			return i
		} else {
			lastIndex[v-'a'] = -1
		}
	}
	return -1
}
