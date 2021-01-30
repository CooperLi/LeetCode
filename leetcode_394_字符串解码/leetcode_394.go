package leetcode_394_字符串解码

import (
    "strings"
)

/*
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:
s = "3[a]2[bc]", 返回 "aaabcbc".
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
*/

func decodeString(s string) string {
    stackNums := make([]int, 0)
    stackStr := make([]string, 0)
    var res string
    var num int
    for i := 0; i < len(s); i++ {
        switch {
        case s[i] >= '0' && s[i] <= '9':    // 是否数字
            // 这里的数字都是字符，所以用数字字符减去0的ASCII码就得到了真的数字大小
            num = 10*num + int(s[i]) - '0'
        case s[i] == '[':   // 入栈阶段
            // 入栈数字
            stackNums = append(stackNums, num)
            // 置空用于下次入栈
            num = 0
            // 字符入栈
            stackStr = append(stackStr, res)
            // 置空用于下次入栈
            res = ""
        case s[i] == ']':   // 出栈阶段
            // 获取最后入栈的字符
            tmp := stackStr[len(stackStr)-1]
            // 相当于pop
            stackStr = stackStr[:len(stackStr)-1]
            // 获取最后入栈的数字
            count := stackNums[len(stackNums)-1]
            // 相当于pop
            stackNums = stackNums[:len(stackNums)-1]
            // 拼接答案，重复完[]里面的，再接上外面的字符
            res = tmp + strings.Repeat(res, count)  // 重要
        default:    // 默认是字符的情况，不要忘记
            // 因为不会有3a这种情况，所以数字后面必定是'['，直接相接
            res += string(s[i])
        }
    }
    return res
}

/*
打印顺序：
输入：4[a3[b2[c2[d]]]]

case=isNum, num=4
case=[, stackNums=[4]
case=[, stackStr=[]
case=default, res=a
case=isNum, num=3
case=[, stackNums=[4 3]
case=[, stackStr=[ a]
case=default, res=b
case=isNum, num=2
case=[, stackNums=[4 3 2]
case=[, stackStr=[ a b]
case=default, res=c
case=isNum, num=2
case=[, stackNums=[4 3 2 2]
case=[, stackStr=[ a b c]
case=default, res=d
case=], tmp=c
case=], stackNums=[4 3 2 2]
case=], count=2
case=], stackNums=[4 3 2]
case=], res=cdd
case=], tmp=b
case=], stackNums=[4 3 2]
case=], count=2
case=], stackNums=[4 3]
case=], res=bcddcdd
case=], tmp=a
case=], stackNums=[4 3]
case=], count=3
case=], stackNums=[4]
case=], res=abcddcddbcddcddbcddcdd
case=], tmp=
case=], stackNums=[4]
case=], count=4
case=], stackNums=[]
case=], res=abcddcddbcddcddbcddcddabcddcddbcddcddbcddcddabcddcddbcddcddbcddcddabcddcddbcddcddbcddcdd
 */