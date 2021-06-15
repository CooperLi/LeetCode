package leetcode_20_有效的括号

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"
输出: true

示例 2:
输入: "()[]{}"
输出: true

示例 3:
输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false

示例 5:
输入: "{[]}"
输出: true
*/
func isValid(s string) bool {
	// 空字符串默认配对的
	if s == "" {
		return true
	}
	// 奇数个括号肯定不能全匹配
	if len(s)%2 == 1 {
		return false
	}
	// 建立映射关系
	// keyMap["}"] = {
	// keyMap["{"] = 空
	keyMap := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	var stack []string
	for i := 0; i < len(s); i++ {
		// 如果栈里有元素就在栈里查找
		if len(stack) > 0 {
			// 如果和map里面的有匹配，也就是，目前元素是右半边括号
			if tmp, ok := keyMap[string(s[i])]; ok {
				// 比较栈顶元素
				top := stack[len(stack)-1]
				// 【重要】如果当前遍历到的元素的映射和栈顶上的括号相同
				if top == tmp {
					// 栈顶元素pop掉
					stack = stack[:len(stack)-1]
					// 跳过这次循环
					continue
				}
			} // ↓
		}
		// 空栈或者没匹配上，就加入栈
		stack = append(stack, string(s[i]))
	}
	// 如果最后栈是空的，说明都匹配上了，返回true
	return len(stack) == 0
}
