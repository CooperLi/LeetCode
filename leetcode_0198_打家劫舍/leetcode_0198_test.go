package leetcode_0198_打家劫舍

import (
	"testing"
)

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{[]int{1, 2, 3, 1}},
			4,
		},
		{
			"2",
			args{[]int{1, 100, 3, 2}},
			102,
		},
		{
			"3",
			args{[]int{}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
