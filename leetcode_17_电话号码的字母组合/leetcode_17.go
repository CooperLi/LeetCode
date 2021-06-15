package leetcode_17_电话号码的字母组合

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

1:!@#   2:abc   3:def
4:ghi   5:jkl   6:mno
7:pqrs  8:tuv   9:wxyz
*:+     0:      #:↑

示例:
输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/

/*
队列，先进先出，把已经拼接过的插到尾部，然后待拼接的从头部出去，直到所有的元素都拼接完了
*/
func letterCombinations(digits string) []string {
	if digits == "" || len(digits) == 0 {
		return nil
	}
	m := []string{" ", "*", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var res []string
	res = append(res, "")
	// 迭代每个数字
	for i := 0; i < len(digits); i++ {
		// 由当前遍历到的字符，取字典表中查找对应的字符串
		// byte类型的数字-'0' = int类型数字
		letters := m[digits[i]-'0']
		// 必须预先定义size，否则res在变化，会超时
		size := len(res)
		// 计算出队列长度后，将队列中的每个元素挨个拿出来,和letter里面的拼接
		for j := 0; j < size; j++ {
			// 每次都从队列中拿出第一个元素
			tmp := res[0]
			res = res[1:]
			// 然后跟"def"这样的字符串拼接，并再次放到队列尾部
			for k := 0; k < len(letters); k++ {
				res = append(res, tmp+string(letters[k]))
			}
		}
	}
	return res
}

/*
letters= hki
tmp= ad
inner res= [ae af bd be bf cd ce cf adh]
inner res= [ae af bd be bf cd ce cf adh adk]
inner res= [ae af bd be bf cd ce cf adh adk adi]
*/
