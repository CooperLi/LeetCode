package leetcode_1360_日期之间隔几天

import (
	"strconv"
	"strings"
	"time"
)

/*
请你编写一个程序来计算两个日期之间隔了多少天。
日期以字符串形式给出，格式为 YYYY-MM-DD，如示例所示。

示例 1：
输入：date1 = "2019-06-29", date2 = "2019-06-30"
输出：1
示例 2：

输入：date1 = "2020-01-15", date2 = "2019-12-31"
输出：15


提示：
给定的日期是 1971 年到 2100 年之间的有效日期。
*/

// 非闰年每月天数固定
var months = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func daysBetweenDates(date1 string, date2 string) int {
	res := getDays(date1) - getDays(date2)
	if res < 0 {
		return -res
	}
	return res
}

// 是否闰年定义：
// 1、不能被100整除但能被4整除
// 2、能被400整除
func isLeap(year int) int {
	if (year%100 != 0 && year%4 == 0) || year%400 == 0 {
		return 1
	}
	return 0
}

// 计算当前日期和 1971-01-01 之间有多少天
func getDays(date string) int {
	fields := strings.Split(date, "-")
	year, _ := strconv.Atoi(fields[0])
	month, _ := strconv.Atoi(fields[1])
	day, _ := strconv.Atoi(fields[2])

	var days = 0
	// 计算每年要多出来几天
	for y := 1971; y < year; y++ {
		days += 365 + isLeap(y)
	}
	for m := 1; m < month; m++ {
		if m == 2 {
			days += isLeap(year) + 28
		} else {
			days += months[m]
		}
	}

	return days + day
}

// 偷懒解法，使用官方库
func daysBetweenDates2(date1 string, date2 string) int {
	t1, _ := time.Parse("2006-01-02", date1)
	t2, _ := time.Parse("2006-01-02", date2)
	ans := t1.Sub(t2).Hours()
	if ans < 0 {
		return -int(ans) / 24
	}
	return int(ans) / 24
}
