package leetcode_0093_复原IP地址

import (
	"strconv"
)

/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:
输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
*/

func restoreIpAddresses(s string) []string {
	var res []string
	size := len(s)
	for i := 0; i < 3; i++ {
		for j := i + 1; j < i+4; j++ {
			for k := j + 1; k < j+4; k++ {
				if i < size && j < size && k < size {
					tmp1 := s[0 : i+1]
					tmp2 := s[i+1 : j+1]
					tmp3 := s[j+1 : k+1]
					tmp4 := s[k+1:]
					if helper(tmp1) && helper(tmp2) && helper(tmp3) && helper(tmp4) {
						ip := tmp1 + "." + tmp2 + "." + tmp3 + "." + tmp4
						res = append(res, ip)
					}
				}
			}
		}
	}
	return res
}

func helper(tmp string) bool {
	size := len(tmp)
	num, _ := strconv.Atoi(tmp)
	if tmp == "" || size == 0 || size > 3 || (tmp[0] == '0' && size > 1) || num > 255 {
		return false
	}
	return true
}
