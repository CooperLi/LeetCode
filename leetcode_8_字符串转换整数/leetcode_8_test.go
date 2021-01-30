package leetcode_8_字符串转换整数

import (
	"math"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{"10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"},
			math.MaxInt32,
		},
		{
			"2",
			args{"10000a"},
			10000,
		},
		{
			"3",
			args{"123"},
			123,
		},
		{
			"4",
			args{"-123"},
			-123,
		},
		{
			"5",
			args{"     asdlkjf1"},
			0,
		},
		{
			"6",
			args{"4193 with words"},
			4193,
		},
		{
			"7",
			args{"-91283472332"},
			-2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
